package generation

import (
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

func (q *Query) processArguments(vars []*ast.VariableDefinition) {
	for _, v := range vars {
		q.Input.fields = append(q.Input.fields, q.processVariable(v))
	}
}

func (q *Query) processVariable(v *ast.VariableDefinition) Field {
	name := v.Variable
	f := Field{
		name: fmt.Sprintf("%s ", strings.Title(name)),
	}

	if v.Type.NamedType != "" {
		f.typ = v.Type.NamedType
	} else {
		panic(fmt.Errorf("Unexpected named type: %s", v.Type))
	}

	f.jsonTag(name, !v.Type.NonNull)

	return f
}
