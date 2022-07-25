package sgml

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderSpecialCharacter(s *types.SpecialCharacter) (string, error) {
	switch s.Name {
	case `&`:
		return "&amp;", nil
	case `<`:
		return "&lt;", nil
	case `>`:
		return "&gt;", nil
	default:
		return "", fmt.Errorf("unknown special character: '%s'", s.Name)
	}
}
