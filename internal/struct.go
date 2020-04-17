package internal

import (
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

// TODO: generate a struct to hold the 'method'
// query<OperationName>(<OperationName>Input) <OperationName>Payload

// Query Query export
type Query struct {
	query   string
	name    string
	input   *Struct
	payload *Struct
}

// Struct Struct export
type Struct struct {
	builder strings.Builder
	depth   int
}

func (s *Struct) write(str string) error {
	for i := 1; i <= s.depth; i++ {
		// spaces for better console output
		s.builder.WriteString("  ")
	}

	_, err := s.builder.WriteString(str)
	return err
}

func (s *Struct) writeNoIdent(str string) error {
	_, err := s.builder.WriteString(str)
	return err
}

func (s *Struct) indent() {
	s.depth++
}

func (s *Struct) deindent() {
	s.depth--
}

// TODO: probably need schema
func generateStruct(query *ast.QueryDocument) ([]Query, error) {
	var queries []Query

	for _, operation := range query.Operations {
		q := Query{
			name: operation.Name,
			input: &Struct{
				builder: strings.Builder{},
			},
			payload: &Struct{
				builder: strings.Builder{},
			},
		}

		q.processPayload(operation)

		if len(operation.VariableDefinitions) > 0 {
			q.input.write("\n")
			q.processArguments(operation.VariableDefinitions)
			q.input.write("\n")
		}

		if len(operation.VariableDefinitions) > 0 {
			q.payload.write(fmt.Sprintf(`func %s(input %sInput) (*%sPayload, error) {
	return nil, nil
}`, q.name, q.name, q.name))
		} else {
			q.payload.write(fmt.Sprintf(`func %s() (*%sPayload, error) {
	return nil, nil
}`, q.name, q.name))
		}

		// fmt.Printf("%s\n", s.builder.String())

		queries = append(queries, q)
	}

	return queries, nil
}

func (q *Query) processArguments(vars []*ast.VariableDefinition) {
	q.input.write(fmt.Sprintf("type %sInput struct {\n", q.name))

	q.input.indent()
	for _, v := range vars {
		q.processVariable(v)
	}
	q.input.deindent()

	q.input.write("}\n")
}

func (q *Query) processPayload(operation *ast.OperationDefinition) {
	q.payload.write(fmt.Sprintf("type %sPayload struct {\n", q.name))

	q.payload.indent()
	for _, selection := range operation.SelectionSet {
		q.processField(selection)
	}
	q.payload.deindent()

	q.payload.write("}\n")
}

func (q *Query) processVariable(v *ast.VariableDefinition) {
	name := v.Variable
	q.input.write(fmt.Sprintf("%s ", strings.Title(name)))

	if v.Type.NamedType != "" {
		q.input.writeNoIdent(fmt.Sprintf(" %s", v.Type.NamedType))
	} else {
		// TODO
		fmt.Printf("Variable Type: %s", v.Type)
	}

	if !v.Type.NonNull {
		q.input.writeNoIdent(fmt.Sprintf(" `json:\"%s,omitempty\"`", name))
	} else {
		q.input.writeNoIdent(fmt.Sprintf(" `json:\"%s\"`", name))
	}

	q.input.writeNoIdent("\n")
}

func (q *Query) processField(selection ast.Selection) {
	field, ok := selection.(*ast.Field)

	if ok {
		fieldType := field.Definition.Type
		name := field.Name

		if field.Alias != "" {
			name = field.Alias
		}

		q.payload.write(fmt.Sprintf("%s ", strings.Title(name)))

		// List Type
		if fieldType.NamedType == "" {
			q.payload.writeNoIdent("[]")

			if !fieldType.NonNull {
				q.payload.writeNoIdent("*")
			}
		}

		// Nested
		if len(field.SelectionSet) > 0 {
			q.payload.indent()
			q.payload.writeNoIdent("struct {\n")
			for _, selection := range field.SelectionSet {
				q.processField(selection)
			}
			q.payload.deindent()
			q.payload.write("}\n")
		} else {
			q.payload.writeNoIdent(fieldType.NamedType)

			if !fieldType.NonNull {
				q.payload.writeNoIdent(fmt.Sprintf(" `json:\"%s,omitempty\"`", name))
			} else {
				q.payload.writeNoIdent(fmt.Sprintf(" `json:\"%s\"`", name))
			}

			q.payload.writeNoIdent("\n")
		}

	} else {
		// TODO: remove after more testing
		panic(fmt.Errorf("unexpected Selection Type: %v\n", selection))
	}
}

func (q Query) String() string {
	b := strings.Builder{}

	b.WriteString(q.payload.builder.String())

	if q.input != nil {
		b.WriteString("\n")
		b.WriteString(q.input.builder.String())
	}

	// TODO: write query and method

	return b.String()
}
