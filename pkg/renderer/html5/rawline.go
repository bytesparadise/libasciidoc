package html5

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderVerbatimLine(l types.VerbatimLine) ([]byte, error) {
	return []byte(strings.TrimRight(l.Content, " ")), nil
}
