package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementRoles(attrs types.Attributes) string {
	switch r := attrs[types.AttrRole].(type) {
	case []string:
		return strings.Join(r, " ")
	case string:
		return r
	default:
		return ""
	}
}
