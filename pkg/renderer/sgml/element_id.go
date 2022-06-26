package sgml

import (
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementID(attrs types.Attributes) string {
	if id, ok := attrs[types.AttrID].(string); ok {
		return string(texttemplate.HTMLEscapeString(id))
	}
	return ""
}
