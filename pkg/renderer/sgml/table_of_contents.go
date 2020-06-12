package sgml

import (
	"bytes"
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderTableOfContents(ctx *renderer.Context, toc types.TableOfContents) ([]byte, error) {
	log.Debug("rendering table of contents...")
	renderedSections, err := r.renderTableOfContentsSections(ctx, toc.Sections)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering table of contents")
	}
	if renderedSections == "" {
		// nothing to render (document has no section)
		return []byte{}, nil
	}
	result := &bytes.Buffer{}
	err = r.tocRoot.Execute(result, renderedSections)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering table of contents")
	}
	// log.Debugf("rendered ToC: %s", result.Bytes())
	return result.Bytes(), nil
}

func (r *sgmlRenderer) renderTableOfContentsSections(ctx *renderer.Context, sections []types.ToCSection) (sanitized, error) {
	if len(sections) == 0 {
		return "", nil
	}
	resultBuf := &bytes.Buffer{}
	err := r.tocSection.Execute(resultBuf, ContextualPipeline{
		Context: ctx,
		Data: struct {
			Level    int
			Sections []types.ToCSection
		}{
			Level:    sections[0].Level,
			Sections: sections,
		},
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to render document ToC")
	}
	log.Debugf("retrieved sections for ToC: %+v", sections)
	return sanitized(resultBuf.String()), nil //nolint: gosec
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
	log.Debugf("visiting children section: %t (%d < %d)", currentLevel < tocLevels, currentLevel, tocLevels)
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
			Title:    string(renderedTitle),
			Children: children,
		},
	}, nil

}

func getTableOfContentsLevels(ctx *renderer.Context) (int, error) {
	log.Debugf("doc attributes: %v", ctx.Attributes)
	if l, found := ctx.Attributes.GetAsString(types.AttrTableOfContentsLevels); found {
		log.Debugf("ToC levels: '%s'", l)
		return strconv.Atoi(l)
	}
	log.Debug("ToC levels: '2' (default)")
	return 2, nil
}
