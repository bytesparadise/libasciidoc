package sgml

import (
	"text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementID(attrs types.Attributes) sanitized {
	if id, ok := attrs[types.AttrID].(string); ok {
		return sanitized(template.HTMLEscapeString(id))
	}
	return ""
}
