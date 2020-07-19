package sgml

import (
	"html/template"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementRoles(attrs types.Attributes) string {
	switch r := attrs[types.AttrRole].(type) {
	case []string:
		return string(template.HTMLEscapeString(strings.Join(r, " ")))
	case string:
		return string(template.HTMLEscapeString(r))
	default:
		return ""
	}
}

// Image roles add float and alignment attributes -- we turn these into roles.
func (r *sgmlRenderer) renderImageRoles(attrs types.Attributes) string {
	var roles []string

	if val, ok := attrs.GetAsString(types.AttrFloat); ok {
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
	return string(template.HTMLEscapeString(strings.Join(roles, " ")))
}
