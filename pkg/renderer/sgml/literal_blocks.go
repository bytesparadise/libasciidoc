package sgml

import (
	"math"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderLiteralBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	log.Debugf("rendering literal block")
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", err
	}

	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block roles")
	}
	title, err := r.renderElementTitle(b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
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
		Title:   title,
		Roles:   roles,
		Content: strings.Trim(content, "\n"),
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderLiteralParagraph(ctx *renderer.Context, b *types.Paragraph) (string, error) {
	log.Debugf("rendering literal paragraph")
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", err
	}
	if b.Attributes.GetAsStringWithDefault(types.AttrLiteralBlockType, "") == types.LiteralBlockWithSpacesOnFirstLine {
		content = trimHeadingSpaces(content)
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block roles")
	}
	title, err := r.renderElementTitle(b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
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
		Title:   title,
		Roles:   roles,
		Content: content,
	})
	return result.String(), err
}

func trimHeadingSpaces(content string) string {
	lines := strings.Split(content, "\n")
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
	return strings.Join(lines, "\n")
}
