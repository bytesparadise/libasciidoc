package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var tableOfContentTmpl template.Template
var tableOfContentSectionSetTmpl template.Template

func init() {
	tableOfContentTmpl = newHTMLTemplate("toc", `<div id="toc" class="toc">
<div id="toctitle">Table of Contents</div>
{{.Content}}
</div>`)
	tableOfContentSectionSetTmpl = newHTMLTemplate("toc section", `<ul class="sectlevel{{.Level}}">
{{ range .Elements }}<li><a href="#{{.Href}}">{{.Title}}</a>{{ if .Subelements }}
{{.Subelements}}
</li>{{else}}</li>{{end}}
{{end}}</ul>`)
}

type TableOfContent struct {
	Content template.HTML
}
type TableOfContentSectionGroup struct {
	Level    int
	Elements []TableOfContentSection
}
type TableOfContentSection struct {
	Level       int
	Href        string
	Title       template.HTML
	Subelements *template.HTML
}

func renderTableOfContent(ctx *renderer.Context, m types.TableOfContentsMacro) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	renderedSections, err := renderTableOfContentSections(ctx, ctx.Document.Elements, 1)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering table of content")
	}
	err = tableOfContentTmpl.Execute(result, TableOfContent{
		Content: *renderedSections,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering table of content")
	}
	log.Debugf("rendered TOC: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderTableOfContentSections(ctx *renderer.Context, elements []types.DocElement, currentLevel int) (*template.HTML, error) {
	sections := make([]TableOfContentSection, 0)
	for _, element := range elements {
		log.Debugf("traversing document element of type %T", element)
		switch section := element.(type) {
		case types.Section:
			renderedTitle, err := renderElement(ctx, section.Title.Content)
			if err != nil {
				return nil, errors.Wrapf(err, "error while rendering table of content section")
			}
			tocLevels, err := ctx.Document.Attributes.GetTOCLevels()
			if err != nil {
				return nil, errors.Wrapf(err, "error while rendering table of content section")
			}
			var renderedChildSections *template.HTML
			if currentLevel < *tocLevels {
				renderedChildSections, err = renderTableOfContentSections(ctx, section.Elements, currentLevel+1)
				if err != nil {
					return nil, errors.Wrapf(err, "error while rendering table of content section")
				}
			}
			sections = append(sections, TableOfContentSection{
				Level:       section.Level,
				Href:        section.Title.ID.Value,
				Title:       template.HTML(string(renderedTitle)),
				Subelements: renderedChildSections,
			})
		}
	}
	if len(sections) == 0 {
		return nil, nil
	}
	resultBuf := bytes.NewBuffer(nil)
	tableOfContentSectionSetTmpl.Execute(resultBuf, TableOfContentSectionGroup{
		Level:    sections[0].Level,
		Elements: sections,
	})
	log.Debugf("retrieved sections for TOC: %+v", sections)
	result := template.HTML(resultBuf.String())
	return &result, nil
}
