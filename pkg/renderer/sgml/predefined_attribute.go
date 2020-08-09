package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderPredefinedAttribute(a types.PredefinedAttribute) (string, error) {
	result := &strings.Builder{}
	if err := r.predefinedAttribute.Execute(result, struct {
		Name string
	}{
		Name: a.Name,
	}); err != nil {
		return "", errors.Wrap(err, "error while rendering predefined attribute")
	}
	log.Debugf("rendered predefined attribute for '%s': '%s'", a.Name, result.String())
	return result.String(), nil
}
