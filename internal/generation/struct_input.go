package generation

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/ast"
)

func (i *InputStruct) generateInput(name string, vars []*ast.VariableDefinition) {
	s := Struct{
		key: name,
		typ: inputStruct,
	}

	for _, v := range vars {
		s.fields = append(s.fields, i.processVariable(v))
	}

	i.structs = append(i.structs, s)
}

func (i *InputStruct) processVariable(v *ast.VariableDefinition) Field {
	f := newField(v.Variable, v.Type)

	if v.Type.NamedType != "" {
		f.typ = v.Type.NamedType
	} else {
		panic(fmt.Errorf("Unexpected named type: %s", v.Type))
	}

	f.jsonTag(f.name, !v.Type.NonNull)

	return *f
}

func processInputType(def *ast.Definition) Struct {
	s := Struct{
		key: def.Name,
		typ: fragmentStruct,
	}

	for _, field := range def.Fields {
		name := field.Name
		f := newField(name, field.Type)
		f.typ = field.Type.NamedType

		s.fields = append(s.fields, *f)
	}

	return s
}
