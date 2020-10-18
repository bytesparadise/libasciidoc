package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderVerseBlock(ctx *renderer.Context, b types.VerseBlock) (string, error) {
	result := &strings.Builder{}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render verser block")
	}
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
	}()
	ctx.WithinDelimitedBlock = true
	content, err := r.renderLines(ctx, discardEmptyLines(b.Lines))
	if err != nil {
		return "", errors.Wrap(err, "unable to render verse block")
	}

	err = r.verseBlock.Execute(result, struct {
		Context     *renderer.Context
		ID          string
		Title       string
		Roles       string
		Attribution Attribution
		Content     string
	}{
		Context:     ctx,
		ID:          r.renderElementID(b.Attributes),
		Title:       r.renderElementTitle(b.Attributes),
		Roles:       roles,
		Attribution: verseBlockAttribution(b),
		Content:     string(content),
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderVerseParagraph(ctx *renderer.Context, p types.Paragraph) (string, error) {
	log.Debug("rendering verse paragraph...")
	result := &strings.Builder{}

	content, err := r.renderLines(ctx, p.Lines, r.withPlainText())
	if err != nil {
		return "", errors.Wrap(err, "unable to render verse paragraph lines")
	}
	err = r.verseParagraph.Execute(result, struct {
		Context     *renderer.Context
		ID          string
		Title       string
		Attribution Attribution
		Content     string
		Lines       [][]interface{}
	}{
		Context:     ctx,
		ID:          r.renderElementID(p.Attributes),
		Title:       r.renderElementTitle(p.Attributes),
		Attribution: paragraphAttribution(p),
		Content:     string(content),
		Lines:       p.Lines,
	})

	return result.String(), err
}
