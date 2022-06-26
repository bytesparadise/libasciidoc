package sgml

import (
	"math"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderExampleBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	// default, example block
	number := 0
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block roles")
	}
	c, found, err := b.Attributes.GetAsString(types.AttrCaption)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block caption")
	}
	if !found {
		c, found, err = ctx.Attributes.GetAsString(types.AttrExampleCaption)
		if err != nil {
			return "", errors.Wrap(err, "unable to render example block caption")
		}
		if found && c != "" {
			c += " {counter:example-number}. "
		}
	}
	// TODO: Replace this hack when we have attribute substitution fully working
	if strings.Contains(c, "{counter:example-number}") {
		number = ctx.GetAndIncrementExampleBlockCounter()
		c = strings.ReplaceAll(c, "{counter:example-number}", strconv.Itoa(number))
	}
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block title")
	}
	caption := &strings.Builder{}
	caption.WriteString(c)
	return r.execute(r.exampleBlock, struct {
		Context       *renderer.Context
		ID            string
		Title         string
		Caption       string
		Roles         string
		ExampleNumber int
		Content       string
	}{
		Context:       ctx,
		ID:            r.renderElementID(b.Attributes),
		Title:         title,
		Caption:       caption.String(),
		Roles:         roles,
		ExampleNumber: number,
		Content:       content,
	})
}

func (r *sgmlRenderer) renderExampleParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	log.Debug("rendering example paragraph...")
	content, err := r.renderElements(ctx, p.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example paragraph content")
	}
	roles, err := r.renderElementRoles(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example paragraph roles")
	}
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example paragraph title")
	}
	return r.execute(r.exampleBlock, struct {
		Context       *renderer.Context
		ID            string
		Title         string
		Caption       string
		Roles         string
		ExampleNumber int
		Content       string
	}{
		Context: ctx,
		Roles:   roles,
		ID:      r.renderElementID(p.Attributes),
		Title:   title,
		Content: string(content) + "\n",
	})
}

func (r *sgmlRenderer) renderLiteralParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	log.Debugf("rendering literal paragraph")
	content, err := r.renderElements(ctx, p.Elements)
	if err != nil {
		return "", err
	}
	// only adjust heading spaces if the paragraph has the `LiteralParagraph` attribute
	if p.Attributes[types.AttrStyle] == types.LiteralParagraph {
		content = adjustHeadingSpaces(content)
	}
	roles, err := r.renderElementRoles(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block roles")
	}
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block roles")
	}
	return r.execute(r.literalBlock, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Roles   string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(p.Attributes),
		Title:   title,
		Roles:   roles,
		Content: content,
	})
}

// adjustHeadingSpaces removes the same number of heading spaces on each line, based on the
// number of heading spaces of the first line
func adjustHeadingSpaces(content string) string {
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
