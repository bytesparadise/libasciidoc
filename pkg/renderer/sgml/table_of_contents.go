package sgml

import (
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderTableOfContents(ctx *renderer.Context, toc types.TableOfContents) (string, error) {
	log.Debug("rendering table of contents...")
	renderedSections, err := r.renderTableOfContentsSections(ctx, toc.Sections)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}
	if renderedSections == "" {
		// nothing to render (document has no section)
		return "", nil
	}
	result := &strings.Builder{}
	err = r.tocRoot.Execute(result, renderedSections)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering table of contents")
	}
	// log.Debugf("rendered ToC: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) renderTableOfContentsSections(ctx *renderer.Context, sections []types.ToCSection) (string, error) {
	if len(sections) == 0 {
		return "", nil
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
		Context  *renderer.Context
		Level    int
		Content  string
		Sections []types.ToCSection
	}{
		Context:  ctx,
		Level:    sections[0].Level,
		Content:  contents.String(),
		Sections: sections,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to render document ToC")
	}
	// log.Debugf("retrieved sections for ToC: %+v", sections)
	return resultBuf.String(), nil //nolint: gosec
}

func (r *sgmlRenderer) renderTableOfContentsEntry(ctx *renderer.Context, entry types.ToCSection) (string, error) {
	resultBuf := &strings.Builder{}

	content, err := r.renderTableOfContentsSections(ctx, entry.Children)
	if err != nil {
		return "", errors.Wrap(err, "unable to render ToC entry children")
	}

	err = r.tocEntry.Execute(resultBuf, struct {
		Context  *renderer.Context
		Level    int
		ID       string
		Title    string
		Content  string
		Children []types.ToCSection
	}{
		Context:  ctx,
		Level:    entry.Level,
		ID:       string(entry.ID),
		Title:    entry.Title,
		Content:  content,
		Children: entry.Children,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to render document ToC")
	}
	return resultBuf.String(), nil //nolint: gosec
}

// newTableOfContents initializes a TableOfContents from the sections
// of the given document
func (r *sgmlRenderer) newTableOfContents(ctx *renderer.Context, doc types.Document) (types.TableOfContents, error) {
	sections := make([]types.ToCSection, 0, len(doc.Elements))
	for _, e := range doc.Elements {
		if s, ok := e.(types.Section); ok {
			tocs, err := r.visitSection(ctx, s, 1)
			if err != nil {
				return types.TableOfContents{}, err
			}
			sections = append(sections, tocs...) // cqn be 1 or more (for the root section, we immediately get its children)
		}
	}
	return types.TableOfContents{
		Sections: sections,
	}, nil
}

func (r *sgmlRenderer) visitSection(ctx *renderer.Context, section types.Section, currentLevel int) ([]types.ToCSection, error) {
	tocLevels, err := getTableOfContentsLevels(ctx)
	if err != nil {
		return []types.ToCSection{}, err
	}
	children := make([]types.ToCSection, 0, len(section.Elements))
	// log.Debugf("visiting children section: %t (%d < %d)", currentLevel < tocLevels, currentLevel, tocLevels)
	if currentLevel <= tocLevels {
		for _, e := range section.Elements {
			if s, ok := e.(types.Section); ok {
				tocs, err := r.visitSection(ctx, s, currentLevel+1)
				if err != nil {
					return []types.ToCSection{}, err
				}
				children = append(children, tocs...)
			}
		}
	}
	if section.Level == 0 {
		return children, nil // for the root section, immediately return its children)
	}

	renderedTitle, err := r.renderPlainText(ctx, section.Title)
	if err != nil {
		return []types.ToCSection{}, err
	}

	return []types.ToCSection{
		{
			ID:       section.Attributes.GetAsStringWithDefault(types.AttrID, ""),
			Level:    section.Level,
			Title:    renderedTitle,
			Children: children,
		},
	}, nil

}

func getTableOfContentsLevels(ctx *renderer.Context) (int, error) {
	// log.Debugf("doc attributes: %v", ctx.Attributes)
	if l, found, err := ctx.Attributes.GetAsString(types.AttrTableOfContentsLevels); err != nil {
		return -1, err
	} else if found {
		// log.Debugf("ToC levels: '%s'", l)
		return strconv.Atoi(l)
	}
	log.Debug("ToC levels: '2' (default)")
	return 2, nil
}
