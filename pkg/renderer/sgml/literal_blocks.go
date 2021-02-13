package sgml

import (
	"math"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderLiteralBlock(ctx *renderer.Context, b types.LiteralBlock) (string, error) {
	// log.Debugf("rendering literal block with content: %s", b.Lines)
	lines := make([]string, len(b.Lines))
	var err error
	for i, line := range b.Lines {
		if lines[i], err = r.renderLine(ctx, line); err != nil {
			return "", errors.Wrap(err, "unable to render literal block")
		}
	}
	if t, found := b.Attributes[types.AttrLiteralBlockType]; found && t == types.LiteralBlockWithSpacesOnFirstLine {
		if len(lines) == 1 {
			lines = []string{strings.TrimLeft(lines[0], " ")}
		} else {
			// remove as many spaces as needed on each line
			spaceCount := 0
			// first pass to determine the minimum number of spaces to remove
			for i, line := range lines {
				l := strings.TrimLeft(line, " ")
				if i == 0 {
					spaceCount = len(line) - len(l)
				} else {
					spaceCount = int(math.Min(float64(spaceCount), float64(len(line)-len(l))))
				}
			}
			// log.Debugf("trimming %d space(s) on each line", int(spaceCount))
			// then remove the same number of spaces on each line
			spaces := strings.Repeat(" ", spaceCount)
			for i, line := range lines {
				lines[i] = strings.TrimPrefix(line, spaces)
			}
		}
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block roles")
	}
	result := &strings.Builder{}
	err = r.literalBlock.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Roles   string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(b.Attributes),
		Title:   r.renderElementTitle(b.Attributes),
		Roles:   roles,
		Content: strings.Join(lines, "\n"),
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block")
	}
	return result.String(), nil
}
