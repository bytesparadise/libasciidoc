package sgml

import "github.com/bytesparadise/libasciidoc/pkg/types"

func (r *sgmlRenderer) renderElementID(attrs types.Attributes) string {
	if id, ok := attrs[types.AttrID].(string); ok {
		return id
	}
	return ""
}
