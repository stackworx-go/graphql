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
	Input     *Struct
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

// Name Name export
func (s Struct) Name() string {
	if s.typ == inputStruct || s.typ == payloadStruct {
		return fmt.Sprintf("%s%s", s.key, s.typ)
	}

	return s.key
}

// PayloadStruct PayloadStruct export
type PayloadStruct struct {
	structs []Struct
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

		if len(operation.VariableDefinitions) > 0 {
			q.Input = &Struct{
				key: operation.Name,
				typ: inputStruct,
			}
			q.processArguments(operation.VariableDefinitions)
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
func (s Struct) Print() string {
	b := strings.Builder{}

	b.WriteString(fmt.Sprintf("type %s struct {\n", s.Name()))

	for _, f := range s.fields {
		b.WriteString(fmt.Sprintf("\t%s %s %s\n", strings.Title(f.name), f.getType(), f.tag))
	}

	b.WriteString("}")

	return b.String()
}
