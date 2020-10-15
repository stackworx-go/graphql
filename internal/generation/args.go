package generation

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/ast"
)

func generateArgs(vars []*ast.VariableDefinition) Arguments {
	args := Arguments{}

	for _, v := range vars {
		args = append(args, processVariable(v))
	}

	return args
}

func processVariable(v *ast.VariableDefinition) argument {
	f := newField(v.Variable, v.Type)

	if v.Type.NamedType != "" {
		f.typ = v.Type.NamedType
	} else {
		panic(fmt.Errorf("Unexpected named type: %s", v.Type))
	}

	// TODO: refactor this out of field
	return argument{
		Name: f.name,
		Type: f.GetType(),
	}
}

func processInputType(def *ast.Definition, scalarUpload string) Struct {
	s := Struct{
		key: def.Name,
		typ: fragmentStruct,
	}

	hasFileUpload := false

	for _, field := range def.Fields {
		name := field.Name
		f := newField(name, field.Type)
		f.typ = field.Type.NamedType

		if f.typ == scalarUpload {
			hasFileUpload = true
		}

		s.Fields = append(s.Fields, *f)
	}

	s.hasFileUpload = hasFileUpload
	return s
}
