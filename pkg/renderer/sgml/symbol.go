package sgml

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

var symbols = map[string]string{
	"(C)":  "&#169;",
	"(R)":  "&#174;",
	"(TM)": "&#8482;",
	"...":  "&#8230;&#8203;",
	"'":    "&#8217;",
	"'`":   "&#8216;",
	"`'":   "&#8217;",
	"\"`":  "&#8220;",
	"`\"":  "&#8221;",
}

func (r *sgmlRenderer) renderSymbol(s *types.Symbol) (string, error) {
	if str, found := symbols[s.Name]; found {
		if s.Prefix != "" {
			return s.Prefix + str, nil
		}
		return str, nil
	}
	return "", fmt.Errorf("symbol '%s' is not defined", s.Name)
}
