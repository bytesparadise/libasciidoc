package sgml

import (
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderExampleBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}
	caption := &strings.Builder{}

	// default, example block
	number := 0
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block")
	}
	c, found, err := b.Attributes.GetAsString(types.AttrCaption)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block")
	}
	if !found {
		c, found, err = ctx.Attributes.GetAsString(types.AttrExampleCaption)
		if err != nil {
			return "", errors.Wrap(err, "unable to render example block")
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
	title, err := r.renderElementTitle(b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	caption.WriteString(c)
	err = r.exampleBlock.Execute(result, struct {
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
	return result.String(), err
}

func (r *sgmlRenderer) renderExampleParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	log.Debug("rendering example paragraph...")
	result := &strings.Builder{}
	content, err := r.renderElements(ctx, p.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote paragraph lines")
	}
	roles, err := r.renderElementRoles(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	title, err := r.renderElementTitle(p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	err = r.exampleBlock.Execute(result, struct {
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

	return result.String(), err
}
