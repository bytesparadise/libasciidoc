package sgml

import (
	"text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementStyle(attrs types.Attributes) string {
	if id, ok := attrs[types.AttrStyle].(string); ok {
		return string(template.HTMLEscapeString(id))
	}
	return ""
}
