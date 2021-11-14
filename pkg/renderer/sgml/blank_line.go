package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderBlankLine(ctx *Context, _ *types.BlankLine) (string, error) {
	return "", nil
}

func (r *sgmlRenderer) renderLineBreak() (string, error) {
	buf := &strings.Builder{}
	if err := r.lineBreak.Execute(buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (r *sgmlRenderer) renderThematicBreak() (string, error) {
	buf := &strings.Builder{}
	if err := r.thematicBreak.Execute(buf, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}
