package sgml

import (
	"html/template"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementRoles(attrs types.Attributes) sanitized {
	switch r := attrs[types.AttrRole].(type) {
	case []string:
		return sanitized(template.HTMLEscapeString(strings.Join(r, " ")))
	case string:
		return sanitized(template.HTMLEscapeString(r))
	default:
		return ""
	}
}
