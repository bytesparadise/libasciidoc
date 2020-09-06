package sgml

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderElementRoles(attrs types.Attributes) (string, error) {
	var roles []string
	switch role := attrs[types.AttrRole].(type) {
	case []interface{}:
		for _, er := range role {
			switch er := er.(type) {
			case types.ElementRole:
				s, err := r.renderElementRole(er)
				if err != nil {
					return "", err
				}
				roles = append(roles, s)
			default:
				return "", fmt.Errorf("unpected type of element in role: '%T'", er)
			}
		}
	case types.ElementRole:
		s, err := r.renderElementRole(role)
		if err != nil {
			return "", err
		}
		roles = append(roles, s)
	}
	return strings.Join(roles, " "), nil
}

// Image roles add float and alignment attributes -- we turn these into roles.
func (r *sgmlRenderer) renderImageRoles(attrs types.Attributes) (string, error) {
	var roles []string
	if val, ok := attrs.GetAsString(types.AttrFloat); ok {
		roles = append(roles, val)
	}
	if val, ok := attrs.GetAsString(types.AttrImageAlign); ok {
		roles = append(roles, "text-"+val)
	}
	switch role := attrs[types.AttrRole].(type) {
	case []interface{}:
		for _, er := range role {
			switch er := er.(type) {
			case types.ElementRole:
				s, err := r.renderElementRole(er)
				if err != nil {
					return "", err
				}
				roles = append(roles, s)
			default:
				return "", fmt.Errorf("unpected type of element in role: '%T'", er)
			}
		}
	case types.ElementRole:
		s, err := r.renderElementRole(role)
		if err != nil {
			return "", err
		}
		roles = append(roles, s)
	}
	return strings.Join(roles, " "), nil
}

func (r *sgmlRenderer) renderElementRole(role types.ElementRole) (string, error) {
	result := strings.Builder{}
	for _, e := range role {
		switch e := e.(type) {
		case string:
			result.WriteString(e)
		case types.StringElement:
			result.WriteString(e.Content)
		case types.SpecialCharacter:
			s, err := r.renderSpecialCharacter(e)
			if err != nil {
				return "", err
			}
			result.WriteString(s)
		default:
			return "", fmt.Errorf("unexpected type of element while rendering elemenr role: '%T'", e)
		}
	}
	return result.String(), nil
}
