package sgml

import (
	"math"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderLiteralBlock(ctx *renderer.Context, b types.LiteralBlock) (string, error) {
	log.Debugf("rendering delimited block with content: %s", b.Lines)
	var lines []string
	if t, found := b.Attributes.GetAsString(types.AttrLiteralBlockType); found && t == types.LiteralBlockWithSpacesOnFirstLine {
		if len(b.Lines) == 1 {
			lines = []string{strings.TrimLeft(b.Lines[0], " ")}
		} else {
			lines = make([]string, len(b.Lines))
			// remove as many spaces as needed on each line
			spaceCount := float64(0)
			// first pass to determine the minimum number of spaces to remove
			for i, line := range b.Lines {
				l := strings.TrimLeft(line, " ")
				if i == 0 {
					spaceCount = float64(len(line) - len(l))
				} else {
					spaceCount = math.Min(spaceCount, float64(len(line)-len(l)))
				}
			}
			log.Debugf("trimming %d space(s) on each line", int(spaceCount))
			// then remove the same number of spaces on each line
			spaces := strings.Repeat(" ", int(spaceCount))
			for i, line := range b.Lines {
				lines[i] = strings.TrimPrefix(line, spaces)
			}
		}
	} else {
		lines = b.Lines
	}
	result := &strings.Builder{}
	err := r.literalBlock.Execute(result, struct {
		Context *renderer.Context
		ID      sanitized
		Title   sanitized
		Roles   sanitized
		Content string
		Lines   []string
	}{

		Context: ctx,
		ID:      r.renderElementID(b.Attributes),
		Title:   r.renderElementTitle(b.Attributes),
		Roles:   r.renderElementRoles(b.Attributes),
		Lines:   lines,
		Content: strings.Join(lines, "\n"),
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render delimited block")
	}
	return result.String(), nil
}
