package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) prerenderTableOfContents(ctx *context, toc *types.TableOfContents) error {
	if toc == nil || toc.Sections == nil {
		return nil
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("pre-rendering table of contents: %s", spew.Sdump(toc))
	}
	if err := r.prerenderTableOfContentsSections(ctx, toc.Sections); err != nil {
		return errors.Wrap(err, "error while rendering table of contents")
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("pre-rendered table of contents: %s", spew.Sdump(toc))
	// }
	return nil
}

func (r *sgmlRenderer) prerenderTableOfContentsSections(ctx *context, sections []*types.ToCSection) error {
	for _, entry := range sections {
		if err := r.prerenderTableOfContentsEntry(ctx, entry); err != nil {
			return errors.Wrap(err, "unable to render table of contents section")
		}
	}
	// log.Debugf("retrieved sections for table of contents: %+v", sections)
	return nil
}

func (r *sgmlRenderer) prerenderTableOfContentsEntry(ctx *context, entry *types.ToCSection) error {
	if err := r.prerenderTableOfContentsSections(ctx, entry.Children); err != nil {
		return errors.Wrap(err, "unable to render table of contents entry children")
	}
	if ctx.sectionNumbering != nil {
		entry.Number = ctx.sectionNumbering[entry.ID]
	}
	s, found := ctx.elementReferences[entry.ID]
	if !found {
		return errors.New("unable to render table of contents entry title (missing element reference")
	}
	title, err := RenderPlainText(s)
	if err != nil {
		return errors.Wrap(err, "unable to render table of contents entry title (missing element reference")
	}
	entry.Title = title
	return nil
}

func (r *sgmlRenderer) renderTableOfContents(ctx *context, toc *types.TableOfContents) (string, error) {
	if toc == nil || toc.Sections == nil {
		return "", nil
	}

	title, err := r.renderTableOfContentsTitle(ctx)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}

	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("rendering table of contents %s", spew.Sdump(toc))
	// }
	renderedSections, err := r.renderTableOfContentsSections(ctx, toc.Sections)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}
	if renderedSections == "" {
		// nothing to render (document has no section)
		return "", nil
	}
	return r.execute(r.tocRoot, struct {
		Title    string
		Sections string
	}{
		Title:    title,
		Sections: renderedSections,
	})
}

func (r *sgmlRenderer) renderTableOfContentsTitle(ctx *context) (string, error) {
	title, found := ctx.attributes[types.AttrTableOfContentsTitle]

	if !found {
		return "Table of Contents", nil // default value // TODO: use a constant?
	}
	switch title := title.(type) {
	case string:
		return title, nil
	case []interface{}:
		return r.renderElements(ctx, title)
	default:
		return r.renderElement(ctx, title)
	}

}

func (r *sgmlRenderer) renderTableOfContentsSections(ctx *context, sections []*types.ToCSection) (string, error) {
	if len(sections) == 0 {
		return "", nil
	}
	contents := &strings.Builder{}
	for _, entry := range sections {
		buf, err := r.renderTableOfContentsEntry(ctx, entry)
		if err != nil {
			return "", errors.Wrap(err, "unable to render table of contents section")
		}
		contents.WriteString(buf)
	}
	return r.execute(r.tocSection, struct {
		Level   int
		Content string
	}{
		Level:   sections[0].Level,
		Content: contents.String(),
	})
}

func (r *sgmlRenderer) renderTableOfContentsEntry(ctx *context, entry *types.ToCSection) (string, error) {
	content, err := r.renderTableOfContentsSections(ctx, entry.Children)
	if err != nil {
		return "", errors.Wrap(err, "unable to render table of contents entry children")
	}
	return r.execute(r.tocEntry, struct {
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
}
