package sgml

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderVerbatimLine(l types.VerbatimLine) ([]byte, error) {
	result := &bytes.Buffer{}
	if err := r.verbatimLine.Execute(result, l); err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}
