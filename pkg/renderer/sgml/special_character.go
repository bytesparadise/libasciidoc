package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderSpecialCharacter(s types.SpecialCharacter) (string, error) {
	log.Debugf("rendering special character...")
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
