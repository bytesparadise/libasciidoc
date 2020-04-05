package html5

import (
	"bytes"
	"html/template"
	"strconv"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var tocRootTmpl texttemplate.Template
var tocSectionTmpl texttemplate.Template

func init() {
	tocRootTmpl = newTextTemplate("toc", `<div id="toc" class="toc">
<div id="toctitle">Table of Contents</div>
{{ . }}
</div>`)
	tocSectionTmpl = newTextTemplate("toc section", `{{ $ctx := .Context }}{{ with .Data }}<ul class="sectlevel{{ .Level }}">
{{ range .Sections }}<li><a href="#{{ .ID }}">{{ .Title }}</a>{{ if .Children }}
{{ renderChildren $ctx .Children }}
</li>{{else}}</li>{{end}}
{{end}}{{end}}</ul>`,
		texttemplate.FuncMap{
			"renderChildren": renderTableOfContentsSections,
		},
	)
}

// TableOfContents the structure of the table of contents
type TableOfContents struct {
	Content template.HTML
}

// TableOfContentsSectionGroup a group of sections in the table of contents
type TableOfContentsSectionGroup struct {
	Level    int
	Elements []TableOfContentsSection
}

// TableOfContentsSection a section in the table of contents
type TableOfContentsSection struct {
	Level    int
	Href     string
	Title    template.HTML
	Elements template.HTML
}

func renderTableOfContents(ctx renderer.Context, toc types.TableOfContents) ([]byte, error) {
	log.Debug("rendering table of contents...")
	renderedSections, err := renderTableOfContentsSections(ctx, toc.Sections)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering table of contents")
	}
	if renderedSections == template.HTML("") {
		// nothing to render (document has no section)
		return []byte{}, nil
	}
	result := bytes.NewBuffer(nil)
	err = tocRootTmpl.Execute(result, renderedSections)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering table of contents")
	}
	// log.Debugf("rendered ToC: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderTableOfContentsSections(ctx renderer.Context, sections []types.ToCSection) (template.HTML, error) {
	if len(sections) == 0 {
		return template.HTML(""), nil
	}
	resultBuf := bytes.NewBuffer(nil)
	err := tocSectionTmpl.Execute(resultBuf, ContextualPipeline{
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
		return template.HTML(""), errors.Wrap(err, "failed to render document ToC")
	}
	log.Debugf("retrieved sections for ToC: %+v", sections)
	return template.HTML(resultBuf.String()), nil
}

// NewTableOfContents initializes a TableOfContents from the sections
// of the given document
func NewTableOfContents(ctx renderer.Context) (types.TableOfContents, error) {
	sections := make([]types.ToCSection, 0, len(ctx.Document.Elements))
	for _, e := range ctx.Document.Elements {
		if s, ok := e.(types.Section); ok {
			tocs, err := visitSection(ctx, s, 1)
			if err != nil {
				return types.TableOfContents{}, err
			}
			sections = append(sections, tocs...) // cqn be 1 or more (for the root section, we immediatly get its children)
		}
	}
	return types.TableOfContents{
		Sections: sections,
	}, nil
}

func visitSection(ctx renderer.Context, section types.Section, currentLevel int) ([]types.ToCSection, error) {
	tocLevels, err := getTableOfContentsLevels(ctx.Document)
	if err != nil {
		return []types.ToCSection{}, err
	}
	children := make([]types.ToCSection, 0, len(section.Elements))
	log.Debugf("visiting children section: %t (%d < %d)", currentLevel < tocLevels, currentLevel, tocLevels)
	if currentLevel <= tocLevels {
		for _, e := range section.Elements {
			if s, ok := e.(types.Section); ok {
				tocs, err := visitSection(ctx, s, currentLevel+1)
				if err != nil {
					return []types.ToCSection{}, err
				}
				children = append(children, tocs...)
			}
		}
	}
	if section.Level == 0 {
		return children, nil // for the root section, immediatly return its children)
	}

	renderedTitle, err := renderPlainText(ctx, section.Title)
	if err != nil {
		return []types.ToCSection{}, err
	}

	return []types.ToCSection{
		{
			ID:       section.Attributes.GetAsString(types.AttrID),
			Level:    section.Level,
			Title:    string(renderedTitle),
			Children: children,
		},
	}, nil

}

func getTableOfContentsLevels(doc types.Document) (int, error) {
	log.Debugf("doc attributes: %v", doc.Attributes)
	if l, found := doc.Attributes.GetAsString(types.AttrTableOfContentsLevels); found {
		log.Debugf("ToC levels: '%s'", l)
		return strconv.Atoi(l)
	}
	log.Debug("ToC levels: '2' (default)")
	return 2, nil
}
