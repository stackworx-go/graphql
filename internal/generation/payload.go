package generation

import (
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

func (p *PayloadStruct) generatePayload(operation *ast.OperationDefinition) {
	p.processSelectionSet(operation.Name, payloadStruct, operation.SelectionSet)
}

func (p *PayloadStruct) processSelectionSet(key string, structType structType, selectionSet ast.SelectionSet) {
	s := Struct{
		key: key,
		typ: structType,
	}
	for _, selection := range selectionSet {
		s.fields = append(s.fields, p.processField(&s, selection))
	}

	p.structs = append(p.structs, s)
}

func (p *PayloadStruct) processField(parent *Struct, selection ast.Selection) Field {
	switch s := selection.(type) {
	case *ast.Field:
		field := s

		name := field.Name

		if field.Alias != "" {
			name = field.Alias
		}

		f := newField(name, field.Definition.Type)

		// Nested
		if len(field.SelectionSet) > 0 {
			f.typ = fmt.Sprintf("%s%s", parent.Name(), strings.Title(field.Name))
			p.processSelectionSet(f.typ, fragmentStruct, field.SelectionSet)
		} else {
			f.typ = field.Definition.Type.NamedType
		}

		return *f
	case *ast.InlineFragment:
		inlineFragment := s

		key := fmt.Sprintf("%sFragment", inlineFragment.TypeCondition)

		f := Fragment{
			name:          fmt.Sprintf("%s%s", parent.Name(), inlineFragment.TypeCondition),
			reference:     key,
			typeCondition: inlineFragment.TypeCondition,
		}

		p.processSelectionSet(f.Name(), fragmentStruct, inlineFragment.SelectionSet)

		parent.fragments = append(parent.fragments, f)

		return Field{
			name: key,
			typ:  "*" + f.Name(),
		}
	default:
		// TODO: remove after more testing
		panic(fmt.Errorf("unexpected Selection Type: %v", selection))
	}

}
