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
	Input     *InputStruct
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
	key    string
	fields []Field
	typ    structType
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

// InputStruct InputStruct export
type InputStruct struct {
	structs Structs
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
		var buff bytes.Buffer
		queryFmt := formatter.NewFormatter(&buff)

		queryFmt.FormatQueryDocument(&ast.QueryDocument{
			Operations: ast.OperationList{operation},
			// TODO: fragments
		})

		q := Query{
			Name:      operation.Name,
			Query:     buff.String(),
			operation: operation,
			Payload:   &PayloadStruct{},
			Input:     nil,
		}

		q.Payload.generatePayload(operation)

		// TODO: rather make each variable its own argument
		if len(operation.VariableDefinitions) > 0 {
			q.Input = &InputStruct{}
			q.Input.generateInput(operation.Name, operation.VariableDefinitions)
		}

		queries = append(queries, q)
	}

	return queries, nil
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

// Print Print export
func (p InputStruct) Print() string {
	b := strings.Builder{}

	for _, s := range p.structs {
		b.WriteString(s.Print())
		b.WriteString("\n")
	}

	return b.String()
}

// Print Print export
func (s Struct) Print() string {
	b := strings.Builder{}

	b.WriteString(fmt.Sprintf("type %s struct {\n", s.Name()))

	for _, f := range s.fields {
		b.WriteString(fmt.Sprintf("\t%s %s %s\n", strings.Title(f.name), f.getType(), f.tag))
	}

	b.WriteString("}")

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
