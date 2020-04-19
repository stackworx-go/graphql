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
		s.fields = append(s.fields, p.processField(s, selection))
	}

	p.structs = append(p.structs, s)
}

func (p *PayloadStruct) processField(parent Struct, selection ast.Selection) Field {
	field, ok := selection.(*ast.Field)

	f := Field{}

	f.nullable = true

	if ok {
		fieldType := field.Definition.Type
		f.name = field.Name

		if field.Alias != "" {
			f.name = field.Alias
		}

		if fieldType.NonNull {
			f.nullable = false
		}

		// List Type
		if fieldType.NamedType == "" {
			f.list = true
		}

		// Nested
		if len(field.SelectionSet) > 0 {
			f.typ = fmt.Sprintf("%s%s", parent.Name(), strings.Title(field.Name))
			p.processSelectionSet(f.typ, fragmentStruct, field.SelectionSet)
		} else {
			f.typ = fieldType.NamedType
		}

		f.jsonTag(f.name, !fieldType.NonNull)

	} else {
		// TODO: remove after more testing
		panic(fmt.Errorf("unexpected Selection Type: %v", selection))
	}

	return f
}
