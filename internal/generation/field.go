package generation

import (
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

// Field Field export
type Field struct {
	name     string
	typ      string
	Tag      string
	nullable bool
	list     bool
}

func (f *Field) jsonTag(key string, omitempty bool) {
	if omitempty {
		f.Tag = fmt.Sprintf("`json:\"%s,omitempty\"`", key)
	} else {
		f.Tag = fmt.Sprintf("`json:\"%s\"`", key)
	}
}

func (f Field) Name() string {
	return strings.Title(f.name)
}

func (f Field) GetType() string {
	b := strings.Builder{}

	if f.list {
		b.WriteString("[]")
	}

	if f.nullable {
		b.WriteString("*")
	}

	b.WriteString(mapType(f.typ))

	return b.String()
}

func mapType(namedType string) string {
	switch namedType {
	case "ID":
		return "string"
	case "Boolean":
		return "bool"
	case "Int":
		return "int"
	case "String":
		return "string"
	default:
		return namedType
	}
}

func newField(name string, typ *ast.Type) *Field {
	f := Field{
		name: name,
	}

	f.nullable = true

	if typ.NonNull {
		f.nullable = false
	}

	// List Type
	if typ.NamedType == "" {
		f.list = true
	}

	// TODO: find better way to do this
	// This is just to avoid lint warnings
	if f.name != "__typename" {
		f.jsonTag(f.name, !typ.NonNull)
	}

	return &f
}
