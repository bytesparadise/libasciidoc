package html5

import "github.com/bytesparadise/libasciidoc/pkg/types"

func renderElementID(attrs types.ElementAttributes) string {
	if id, ok := attrs[types.AttrID].(string); ok {
		return id
	}
	return ""
}
