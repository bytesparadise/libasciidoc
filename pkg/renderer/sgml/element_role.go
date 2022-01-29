package sgml

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementRoles(ctx *Context, attrs types.Attributes) (string, error) {
	if roles, ok := attrs[types.AttrRoles].(types.Roles); ok {
		result := make([]string, len(roles))
		for i, e := range roles {
			s, err := r.renderElementRole(ctx, e)
			if err != nil {
				return "", err
			}
			result[i] = s
		}
		return strings.Join(result, " "), nil
	}
	return "", nil
}

// Image roles add float and alignment attributes -- we turn these into roles.
func (r *sgmlRenderer) renderImageRoles(ctx *Context, attrs types.Attributes) (string, error) {
	var result []string
	if val, ok, err := attrs.GetAsString(types.AttrFloat); err != nil {
		return "", err
	} else if ok {
		result = append(result, val)
	}
	if val, ok, err := attrs.GetAsString(types.AttrImageAlign); err != nil {
		return "", err
	} else if ok {
		result = append(result, "text-"+val)
	}
	if roles, ok := attrs[types.AttrRoles].(types.Roles); ok {
		for _, e := range roles {
			s, err := r.renderElementRole(ctx, e)
			if err != nil {
				return "", err
			}
			result = append(result, s)
		}
	}
	// log.Debugf("rendered image roles: '%s'", result)
	return strings.Join(result, " "), nil
}

func (r *sgmlRenderer) renderElementRole(ctx *Context, role interface{}) (string, error) {
	result := strings.Builder{}
	switch role := role.(type) {
	case string:
		result.WriteString(role)
	case []interface{}:
		// when the role is made of strings and special characters
		for _, e := range role {
			s, err := r.renderElement(ctx, e)
			if err != nil {
				return "", err
			}
			result.WriteString(s)
		}
	default:
		return "", fmt.Errorf("unexpected type of element while rendering element role: '%T'", role)
	}
	// log.Debugf("rendered role: '%s'", result.String())
	return result.String(), nil
}
