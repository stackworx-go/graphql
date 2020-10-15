package generation

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
)

// Query Query export
type Query struct {
	Query     string
	operation *ast.OperationDefinition
	Name      string
	Arguments Arguments
	Payload   *PayloadStruct
	FileMap   map[string]string
}

func (q Query) HasUploadFiles() bool {
	return len(q.FileMap) > 0
}

func (q Query) HasInput() bool {
	return q.Arguments != nil
}

type structType string

const (
	inputStruct    structType = "Input"
	payloadStruct  structType = "Payload"
	fragmentStruct structType = "Fragment"
)

// Struct Struct export
type Struct struct {
	key       string
	Fields    []Field
	typ       structType
	fragments []Fragment
}

type Structs []Struct

// Name Name export
func (s Struct) Name() string {
	if s.typ == inputStruct || s.typ == payloadStruct {
		return fmt.Sprintf("%s%s", s.key, s.typ)
	}

	return s.key
}

// PayloadStruct PayloadStruct export
type PayloadStruct struct {
	Structs Structs
}

// Fragment Fragment export
type Fragment struct {
	name          string
	reference     string
	typeCondition string
	Struct
}

// Name Name export
func (f Fragment) Name() string {
	return fmt.Sprintf("%s%s", f.name, "Fragment")
}

// Arguments Arguments export
type Arguments []argument

type argument struct {
	Name  string
	Field *Field
}

// GenerateInputStructs GenerateInputStructs export
func GenerateInputStructs(schema *ast.Schema) (Structs, error) {
	var inputStructs []Struct

	for _, t := range schema.Types {
		if !t.BuiltIn && t.Kind == ast.InputObject {
			inputStructs = append(inputStructs, processInputType(t))
		}
	}

	return inputStructs, nil
}

// GenerateQueries GenerateQueries export
func GenerateQueries(query *ast.QueryDocument, schema *ast.Schema, scalarUpload string) ([]Query, error) {
	var queries []Query

	for _, operation := range query.Operations {
		// Insert extra fields similar to relay
		insertQueryTypeFields(&operation.SelectionSet)

		var buff bytes.Buffer
		queryFmt := formatter.NewFormatter(&buff)

		singleOperationQuery := &ast.QueryDocument{
			Operations: ast.OperationList{operation},
			// TODO: fragments
		}

		queryFmt.FormatQueryDocument(singleOperationQuery)

		q := Query{
			Name:      operation.Name,
			Query:     buff.String(),
			operation: operation,
			Payload:   &PayloadStruct{},
		}

		q.Payload.generatePayload(operation)

		if len(operation.VariableDefinitions) > 0 {
			q.Arguments = generateArgs(operation.VariableDefinitions)
		}

		q.FileMap = processFileUploads(q, schema, scalarUpload)

		queries = append(queries, q)
	}

	return queries, nil
}

func processFileUploads(q Query, schema *ast.Schema, scalarUpload string) map[string]string {
	fileMap := make(map[string]string)
	for _, arg := range q.Arguments {
		// TODO: check file list
		def := schema.Types[arg.Field.typ]

		path := []string{arg.Name}
		processQueryArguments(def, schema, fileMap, path, scalarUpload)
	}

	return fileMap
}

func processQueryArguments(def *ast.Definition, schema *ast.Schema, fileMap map[string]string, path []string, scalarUpload string) {
	// TODO: check file list
	if def.Kind == ast.Scalar {
		if def.Name == scalarUpload {
			fileMap[strings.Join(path, ".")] = ""
		}
	} else if def.Kind == ast.InputObject {
		for _, field := range def.Fields {
			subPath := append(path, field.Name)
			processQueryArguments(schema.Types[field.Type.Name()], schema, fileMap, subPath, scalarUpload)
		}
	}
}

func insertQueryTypeFields(selectionSet *ast.SelectionSet) {
	// insertTypename := false
	var addTypenameSelectionSets []*ast.SelectionSet

	// https://stackoverflow.com/a/51106195/614371
	for idx := range *selectionSet {
		if inlineFragment, ok := (*selectionSet)[idx].(*ast.InlineFragment); ok {
			// insertTypename = true
			insertQueryTypeFields(&inlineFragment.SelectionSet)

			addTypenameSelectionSets = append(addTypenameSelectionSets, &inlineFragment.SelectionSet)
		} else if field, ok := (*selectionSet)[idx].(*ast.Field); ok {
			insertQueryTypeFields(&field.SelectionSet)
		}
	}

	// TODO: check if typename already exists
	for _, set := range addTypenameSelectionSets {
		*set = append(*set, &ast.Field{
			Name: "__typename",
			Definition: &ast.FieldDefinition{
				Name: "__typename",
				Type: &ast.Type{
					NamedType: "String",
					NonNull:   true,
				},
			},
		})
	}

	// if insertTypename {
	// 	*selectionSet = append(*selectionSet, &ast.Field{
	// 		Name: "__typename",
	// 		Definition: &ast.FieldDefinition{
	// 			Name: "__typename",
	// 			Type: &ast.Type{
	// 				NamedType: "String",
	// 				NonNull:   true,
	// 			},
	// 		},
	// 	})
	// }
}

// TODO: move to template
func (s Struct) FragmentUnmarshaler() string {
	b := strings.Builder{}

	for _, f := range s.fragments {
		b.WriteString("\t")
		b.WriteString(fmt.Sprintf(`case "%s":`, f.typeCondition))
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf("\t\terr = json.Unmarshal(data, &f.%s)\n", f.reference))
		b.WriteString(`		if err != nil {
			return err
		}`)
		b.WriteString("\n")
	}

	b.WriteString("\tdefault:\n")
	b.WriteString("\t\t")
	b.WriteString(`panic(fmt.Errorf("unexpected object type: %s", typename.Typename))`)

	return b.String()
}

func (s Struct) HasFragments() bool {
	return len(s.fragments) > 0
}
