package sgml

import (
	"text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementStyle(attrs types.Attributes) sanitized {
	if id, ok := attrs[types.AttrStyle].(string); ok {
		return sanitized(template.HTMLEscapeString(id))
	}
	return ""
}
