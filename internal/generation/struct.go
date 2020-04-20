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
	fields    []Field
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
	structs Structs
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
	Name string
	Type string
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

// GenerateStruct GenerateStruct export
func GenerateStruct(query *ast.QueryDocument) ([]Query, error) {
	var queries []Query

	for _, operation := range query.Operations {
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
			Arguments: nil,
		}

		q.Payload.generatePayload(operation)

		// TODO: rather make each variable its own argument
		if len(operation.VariableDefinitions) > 0 {
			q.Arguments = generateArgs(operation.VariableDefinitions)
		}

		queries = append(queries, q)
	}

	return queries, nil
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

// Print Print export
func (p PayloadStruct) Print() string {
	b := strings.Builder{}

	for _, s := range p.structs {
		b.WriteString(s.Print())
		b.WriteString("\n")
	}

	return b.String()
}

func (s Struct) fragmentUnmarshaler() string {
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

// Print Print export
func (s Struct) Print() string {
	b := strings.Builder{}

	b.WriteString(fmt.Sprintf("type %s struct {\n", s.Name()))

	for _, f := range s.fields {
		b.WriteString(fmt.Sprintf("\t%s %s %s\n", strings.Title(f.name), f.getType(), f.tag))
	}

	b.WriteString("}\n")

	// Custom json serializer
	if len(s.fragments) > 0 {
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf(`func (f *%s) UnmarshalJSON(data []byte) error {
	var typename typename
	err := json.Unmarshal(data, &typename)

	if err != nil {
		return err
	}

	// Extract local Fields if any
	// Causes circular loop
	// Will need second struct 
	// err = json.Unmarshal(data, &f)

	// f.__typename = typename.Typename

	switch(typename.Typename) {
%s
	}

	return nil
}`, s.Name(), s.fragmentUnmarshaler()))
		// 		b.WriteString("\n\n")

		// 		b.WriteString(fmt.Sprintf(`func (f *%[1]s) Typename() string {
		// 	return f.__typename
		// }`, s.Name()))
		// 		b.WriteString("\n")
	}

	return b.String()
}

// Print Print export
func (ss Structs) Print() string {
	b := strings.Builder{}

	for _, s := range ss {
		b.WriteString(s.Print())
		b.WriteString("\n")
	}

	return b.String()
}
