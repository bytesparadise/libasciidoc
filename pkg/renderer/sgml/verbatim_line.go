package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderVerbatimLine(l types.VerbatimLine) (string, error) {
	result := &strings.Builder{}
	if err := r.verbatimLine.Execute(result, l); err != nil {
		return "", err
	}
	return result.String(), nil
}
