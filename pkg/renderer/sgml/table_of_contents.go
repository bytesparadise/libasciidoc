package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) prerenderTableOfContents(ctx *renderer.Context, toc *types.TableOfContents) error {
	if toc == nil || toc.Sections == nil {
		return nil
	}

	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("pre-rendering ToC: %s", spew.Sdump(toc))
	// }
	if err := r.prerenderTableOfContentsSections(ctx, toc.Sections); err != nil {
		return errors.Wrap(err, "error while rendering table of contents")
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("pre-rendered ToC: %s", spew.Sdump(toc))
	// }
	return nil
}

func (r *sgmlRenderer) prerenderTableOfContentsSections(ctx *renderer.Context, sections []*types.ToCSection) error {
	for _, entry := range sections {
		if err := r.prerenderTableOfContentsEntry(ctx, entry); err != nil {
			return errors.Wrap(err, "unable to render ToC section")
		}
	}
	// log.Debugf("retrieved sections for ToC: %+v", sections)
	return nil // nolint:gosec
}

func (r *sgmlRenderer) prerenderTableOfContentsEntry(ctx *renderer.Context, entry *types.ToCSection) error {
	if err := r.prerenderTableOfContentsSections(ctx, entry.Children); err != nil {
		return errors.Wrap(err, "unable to render ToC entry children")
	}
	if ctx.SectionNumbering != nil {
		entry.Number = ctx.SectionNumbering[entry.ID]
	}
	s, found := ctx.ElementReferences[entry.ID]
	if !found {
		return errors.New("unable to render ToC entry title (missing element reference")
	}
	title, err := r.renderPlainText(ctx, s)
	if err != nil {
		return errors.Wrap(err, "unable to render ToC entry title (missing element reference")
	}
	entry.Title = title
	return nil
}

func (r *sgmlRenderer) renderTableOfContents(ctx *renderer.Context, toc *types.TableOfContents) (string, error) {
	if toc == nil || toc.Sections == nil {
		return "", nil
	}

	title, err := r.renderTableOfContentsTitle(ctx)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}

	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("rendering ToC %s", spew.Sdump(toc))
	// }
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
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("rendering sections (in toc): '%s'", spew.Sdump(sections))
	// }

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
	content, err := r.renderTableOfContentsSections(ctx, entry.Children)
	if err != nil {
		return "", errors.Wrap(err, "unable to render ToC entry children")
	}
	resultBuf := &strings.Builder{}
	err = r.tocEntry.Execute(resultBuf, struct {
		Number  string
		ID      string
		Title   string
		Content string
	}{
		Number:  entry.Number,
		ID:      entry.ID,
		Title:   entry.Title,
		Content: content,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to render document ToC")
	}
	return resultBuf.String(), nil // nolint:gosec
}
