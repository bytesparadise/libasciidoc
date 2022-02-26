package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderSpecialCharacter(s *types.SpecialCharacter) (string, error) {
	// log.Debugf("rendering special character '%s'", s.Name)
	return EscapeString(s.Name), nil
}
