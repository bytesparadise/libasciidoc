package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderInlineMenu(m *types.InlineMenu) (string, error) {
	buf := &strings.Builder{}
	if err := r.inlineMenu.Execute(buf, struct {
		Path []string
	}{
		Path: m.Path,
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}
