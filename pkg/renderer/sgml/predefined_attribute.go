package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderPredefinedAttribute(a *types.PredefinedAttribute) (string, error) {
	return predefinedAttribute(a.Name), nil
}
