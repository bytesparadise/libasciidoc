package html5

import (
	"bytes"
	"html/template"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var tableOfContentTmpl texttemplate.Template
var tableOfContentSectionSetTmpl texttemplate.Template

func init() {
	tableOfContentTmpl = newTextTemplate("toc", `<div id="toc" class="toc">
<div id="toctitle">Table of Contents</div>
{{ .Content }}
</div>`)
	tableOfContentSectionSetTmpl = newTextTemplate("toc section", `<ul class="sectlevel{{ .Level }}">
{{ range .Elements }}<li><a href="#{{ .Href }}">{{ .Title }}</a>{{ if .Subelements }}
{{ .Subelements }}
</li>{{else}}</li>{{end}}
{{end}}</ul>`)
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
	Level       int
	Href        string
	Title       template.HTML
	Subelements *template.HTML
}

func renderTableOfContents(ctx *renderer.Context, m types.TableOfContentsMacro) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	renderedSections, err := renderTableOfContentsSections(ctx, ctx.Document.Elements, 1)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering table of content")
	}
	err = tableOfContentTmpl.Execute(result, TableOfContents{
		Content: *renderedSections,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering table of content")
	}
	log.Debugf("rendered TOC: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderTableOfContentsSections(ctx *renderer.Context, elements []interface{}, currentLevel int) (*template.HTML, error) {
	sections := make([]TableOfContentsSection, 0)
	for _, element := range elements {
		log.Debugf("traversing document element of type %T", element)
		switch section := element.(type) {
		case types.Section:
			renderedTitle, err := renderElement(ctx, section.Title.Elements)
			if err != nil {
				return nil, errors.Wrapf(err, "error while rendering table of content section")
			}
			tocLevels, err := ctx.Document.Attributes.GetTOCLevels()
			if err != nil {
				return nil, errors.Wrapf(err, "error while rendering table of content section")
			}
			var renderedChildSections *template.HTML
			if currentLevel < *tocLevels {
				renderedChildSections, err = renderTableOfContentsSections(ctx, section.Elements, currentLevel+1)
				if err != nil {
					return nil, errors.Wrapf(err, "error while rendering table of content section")
				}
			}
			var id string
			if i, ok := section.Title.Attributes[types.AttrID].(string); ok {
				id = i
			}
			renderedTitleStr := strings.TrimSpace(string(renderedTitle))
			sections = append(sections, TableOfContentsSection{
				Level:       section.Level,
				Href:        id,
				Title:       template.HTML(renderedTitleStr),
				Subelements: renderedChildSections,
			})
		}
	}
	if len(sections) == 0 {
		return nil, nil
	}
	resultBuf := bytes.NewBuffer(nil)
	err := tableOfContentSectionSetTmpl.Execute(resultBuf, TableOfContentsSectionGroup{
		Level:    sections[0].Level,
		Elements: sections,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to render document TOC")
	}
	log.Debugf("retrieved sections for TOC: %+v", sections)
	result := template.HTML(resultBuf.String())
	return &result, nil
}
