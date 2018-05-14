package html5

import (
	"bytes"
	"html/template"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/renderer"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var preambleTmpl template.Template
var sectionHeaderTmpl template.Template
var section1ContentTmpl template.Template
var otherSectionContentTmpl template.Template

// initializes the templates
func init() {
	preambleTmpl = newHTMLTemplate("preamble",
		`<div id="preamble">
<div class="sectionbody">
{{.}}
</div>
</div>`)
	section1ContentTmpl = newHTMLTemplate("section 1",
		`<div class="{{.Class}}">
{{.SectionTitle}}
<div class="sectionbody">{{ if .Elements }}
{{.Elements}}{{end}}
</div>
</div>`)
	otherSectionContentTmpl = newHTMLTemplate("other section",
		`<div class="{{.Class}}">
{{.SectionTitle}}{{ if .Elements }}
{{.Elements}}{{end}}
</div>`)
	sectionHeaderTmpl = newHTMLTemplate("other sectionTitle",
		`<h{{.Level}} id="{{.ID}}">{{.Content}}</h{{.Level}}>`)
}

func renderPreamble(ctx *renderer.Context, p types.Preamble) ([]byte, error) {
	log.Debugf("Rendering preamble...")
	renderedElementsBuff := bytes.NewBuffer(nil)
	for i, element := range p.Elements {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render preamble")
		}
		renderedElementsBuff.Write(renderedElement)
		if i < len(p.Elements)-1 {
			renderedElementsBuff.WriteString("\n")
		}
	}
	result := bytes.NewBuffer(nil)
	err := preambleTmpl.Execute(result, template.HTML(renderedElementsBuff.String()))
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section")
	}
	log.Debugf("rendered p: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderSection(ctx *renderer.Context, s types.Section) ([]byte, error) {
	log.Debugf("Rendering section level %d", s.Level)
	renderedSectionTitle, err := renderSectionTitle(ctx, s.Level, s.Title)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section")
	}
	renderedSectionElements, err := renderSectionElements(ctx, s.Elements)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section")
	}
	result := bytes.NewBuffer(nil)
	// select the appropriate template for the section
	var tmpl template.Template
	if s.Level == 1 {
		tmpl = section1ContentTmpl
	} else {
		tmpl = otherSectionContentTmpl
	}
	renderedHTMLSectionTitle := template.HTML(renderedSectionTitle)
	renderedHTMLElements := template.HTML(renderedSectionElements)
	err = tmpl.Execute(result, struct {
		Class        string
		SectionTitle template.HTML
		Elements     template.HTML
	}{
		Class:        "sect" + strconv.Itoa(s.Level),
		SectionTitle: renderedHTMLSectionTitle,
		Elements:     renderedHTMLElements,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section")
	}
	log.Debugf("rendered section: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderSectionTitle(ctx *renderer.Context, level int, sectionTitle types.SectionTitle) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	renderedContent, err := renderElement(ctx, sectionTitle.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering sectionTitle content")
	}
	content := template.HTML(strings.TrimSpace(string(renderedContent)))
	var id string
	if i, ok := sectionTitle.Attributes[types.AttrID].(string); ok {
		id = i
	}
	err = sectionHeaderTmpl.Execute(result, struct {
		Level   int
		ID      string
		Content template.HTML
	}{
		Level:   level + 1,
		ID:      id,
		Content: content,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering sectionTitle")
	}
	// log.Debugf("rendered sectionTitle: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderSectionElements(ctx *renderer.Context, elements []interface{}) ([]byte, error) {
	renderedElementsBuff := bytes.NewBuffer(nil)
	for i, element := range elements {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render section element")
		}
		renderedElementsBuff.Write(renderedElement)
		if i < len(elements)-1 {
			renderedElementsBuff.WriteString("\n")
		}
	}
	return renderedElementsBuff.Bytes(), nil
}
