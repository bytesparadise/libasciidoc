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

// Image roles add float and alignment attributes -- we turn these into roles.
func (r *sgmlRenderer) renderImageRoles(attrs types.Attributes) sanitized {
	var roles []string

	if val, ok := attrs.GetAsString(types.AttrImageFloat); ok {
		roles = append(roles, val)
	}
	if val, ok := attrs.GetAsString(types.AttrImageAlign); ok {
		roles = append(roles, "text-"+val)
	}
	switch r := attrs[types.AttrRole].(type) {
	case []string:
		roles = append(roles, r...)
	case string:
		roles = append(roles, r)
	}
	return sanitized(template.HTMLEscapeString(strings.Join(roles, " ")))
}
