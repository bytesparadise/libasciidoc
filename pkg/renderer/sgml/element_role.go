package sgml

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderElementRoles(ctx *Context, attrs types.Attributes) (string, error) {
	var result []string
	if roles, ok := attrs[types.AttrRoles].([]interface{}); ok {
		for _, e := range roles {
			s, err := r.renderElementRole(ctx, e)
			if err != nil {
				return "", err
			}
			result = append(result, s)
		}
	}
	log.Debugf("rendered roles: '%s'", result)
	return strings.Join(result, " "), nil
}

// Image roles add float and alignment attributes -- we turn these into roles.
func (r *sgmlRenderer) renderImageRoles(ctx *Context, attrs types.Attributes) (string, error) {
	var result []string
	if val, ok := attrs.GetAsString(types.AttrFloat); ok {
		result = append(result, val)
	}
	if val, ok := attrs.GetAsString(types.AttrImageAlign); ok {
		result = append(result, "text-"+val)
	}
	if roles, ok := attrs[types.AttrRoles].([]interface{}); ok {
		for _, e := range roles {
			s, err := r.renderElementRole(ctx, e)
			if err != nil {
				return "", err
			}
			result = append(result, s)
		}
	}
	log.Debugf("rendered image roles: '%s'", result)
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
	log.Debugf("rendered role: '%s'", result.String())
	return result.String(), nil
}
