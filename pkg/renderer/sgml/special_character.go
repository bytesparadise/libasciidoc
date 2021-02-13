package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderSpecialCharacter(ctx *Context, s types.SpecialCharacter) (string, error) {
	// log.Debugf("rendering special character...")
	if !ctx.EncodeSpecialChars {
		// just return the special character as-is
		return s.Name, nil
	}
	// TODO: no need for a template here, just a map
	result := &strings.Builder{}
	if err := r.specialCharacter.Execute(result, struct {
		Name string
	}{
		Name: s.Name,
	}); err != nil {
		return "", errors.Wrap(err, "error while rendering special character")
	}
	return result.String(), nil
}
