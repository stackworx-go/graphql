package generation

import (
	"fmt"
	"strings"
)

// Field Field export
type Field struct {
	name     string
	typ      string
	tag      string
	nullable bool
	list     bool
}

func (f *Field) jsonTag(key string, omitempty bool) {
	if omitempty {
		f.tag = fmt.Sprintf("`json:\"%s,omitempty\"`", key)
	} else {
		f.tag = fmt.Sprintf("`json:\"%s\"`", key)
	}
}

func (f Field) getType() string {
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
