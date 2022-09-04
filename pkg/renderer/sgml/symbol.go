package sgml

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

var symbols = map[string]string{
	"(C)":  "&#169;",
	"(R)":  "&#174;",
	"(TM)": "&#8482;",
	"...":  "&#8230;&#8203;", // include the 'zero width' character (`&#8203;`) to prevent increased letter spacing in justification
	"'":    "&#8217;",
	"'`":   "&#8216;",
	"`'":   "&#8217;",
	"\"`":  "&#8220;",
	"`\"":  "&#8221;",
	"->":   "&#8594;",
	"<-":   "&#8592;",
	"=>":   "&#8658;",
	"<=":   "&#8656;",
	"--":   "&#8212;&#8203;",        // include the 'zero width' character (`&#8203;`) to prevent increased letter spacing in justification
	" -- ": "&#8201;&#8212;&#8201;", // surrounded by thin spaces
}

func (r *sgmlRenderer) renderSymbol(s *types.Symbol) (string, error) {
	if str, found := symbols[s.Name]; found {
		// if s.Prefix != "" {
		// 	return s.Prefix + str, nil
		// }
		return str, nil
	}
	return "", fmt.Errorf("symbol '%s' is not defined", s.Name)
}
