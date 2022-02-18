package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderTableOfContents(ctx *renderer.Context, toc *types.TableOfContents) (string, error) {
	if toc == nil || toc.Sections == nil {
		return "", nil
	}

	title, err := r.renderTableOfContentsTitle(ctx)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendering ToC %s", spew.Sdump(toc))
	}
	renderedSections, err := r.renderTableOfContentsSections(ctx, toc.Sections)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}
	if renderedSections == "" {
		// nothing to render (document has no section)
		return "", nil
	}
	result := &strings.Builder{}
	err = r.tocRoot.Execute(result, struct {
		Title    string
		Sections string
	}{
		Title:    title,
		Sections: renderedSections,
	})
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}
	// log.Debugf("rendered ToC: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) renderTableOfContentsTitle(ctx *renderer.Context) (string, error) {
	title, found, err := ctx.Attributes.GetAsString(types.AttrTableOfContentsTitle)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}
	if !found {
		return "Table of Contents", nil // default value // TODO: use a constant?
	}
	// parse
	value, err := parser.ParseAttributeValue(title)
	if err != nil {
		return "", err
	}
	return r.renderElements(ctx, value)

}

func (r *sgmlRenderer) renderTableOfContentsSections(ctx *renderer.Context, sections []*types.ToCSection) (string, error) {
	if len(sections) == 0 {
		return "", nil
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendering sections (in toc): '%s'", spew.Sdump(sections))
	}

	resultBuf := &strings.Builder{}
	contents := &strings.Builder{}
	for _, entry := range sections {
		buf, err := r.renderTableOfContentsEntry(ctx, entry)
		if err != nil {
			return "", errors.Wrap(err, "unable to render ToC section")
		}
		contents.WriteString(buf)
	}

	err := r.tocSection.Execute(resultBuf, struct {
		Level   int
		Content string
	}{
		Level:   sections[0].Level,
		Content: contents.String(),
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to render document ToC")
	}
	// log.Debugf("retrieved sections for ToC: %+v", sections)
	return resultBuf.String(), nil // nolint:gosec
}

func (r *sgmlRenderer) renderTableOfContentsEntry(ctx *renderer.Context, entry *types.ToCSection) (string, error) {
	resultBuf := &strings.Builder{}

	content, err := r.renderTableOfContentsSections(ctx, entry.Children)
	if err != nil {
		return "", errors.Wrap(err, "unable to render ToC entry children")
	}
	var number string
	if ctx.SectionNumbering != nil {
		number = ctx.SectionNumbering[entry.ID]
	}
	s, found := ctx.ElementReferences[entry.ID]
	if !found {
		return "", errors.New("unable to render ToC entry title (missing element reference")
	}
	entry.Title, err = r.renderPlainText(ctx, s)
	if err != nil {
		return "", errors.Wrap(err, "unable to render ToC entry title (missing element reference")
	}
	err = r.tocEntry.Execute(resultBuf, struct {
		Number  string
		ID      string
		Title   string
		Content string
	}{
		Number:  number,
		ID:      entry.ID,
		Title:   entry.Title,
		Content: content,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to render document ToC")
	}
	return resultBuf.String(), nil // nolint:gosec
}
