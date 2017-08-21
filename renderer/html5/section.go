package html5

import (
	"bytes"
	"context"
	"html/template"
	"reflect"
	"strconv"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// var section1HeaderTmpl *template.Template
var otherSectionHeaderTmpl *template.Template
var section1ContentTmpl *template.Template
var section2ContentTmpl *template.Template
var otherSectionContentTmpl *template.Template

// initializes the templates
func init() {
	section1ContentTmpl = newTemplate("section 1",
		`{{ if .Preamble }}<div id="preamble">
<div class="sectionbody">
{{.Preamble}}
</div>
</div>
{{end}}{{ if .Elements }}{{.Elements}}{{end}}`)
	section2ContentTmpl = newTemplate("section 2",
		`<div class="{{.Class}}">
{{.Heading}}
<div class="sectionbody">{{ if .Elements }}
{{.Elements}}{{end}}
</div>
</div>`)
	otherSectionContentTmpl = newTemplate("other section",
		`<div class="{{.Class}}">
{{.Heading}}{{ if .Elements }}
{{.Elements}}{{end}}
</div>`)
	// 	section1HeaderTmpl = newTemplate("section 1 heading",
	// 		`<div id="header">
	// <h1>{{.Content}}</h1>
	// </div>`)
	otherSectionHeaderTmpl = newTemplate("other heading",
		`<h{{.Level}} id="{{.ID}}">{{.Content}}</h{{.Level}}>`)
}

func renderSection(ctx context.Context, section types.Section) ([]byte, error) {
	switch section.Heading.Level {
	case 1:
		return renderSectionLevel1(ctx, section)
	default:
		return renderOtherSection(ctx, section)
	}
}

func renderSectionLevel1(ctx context.Context, section types.Section) ([]byte, error) {
	// only applies if the first element (if exists) is not a nested section
	var preambleElements []types.DocElement
	var otherElements []types.DocElement
	// log.Debugf("Preparing Preamble elements...")
	for i, element := range section.Elements {
		log.Debugf(" %v", reflect.TypeOf(element))

		if _, ok := element.(*types.Section); ok {
			if i > 0 {
				preambleElements = section.Elements[:i]
			} else {
				preambleElements = make([]types.DocElement, 0)
			}
			otherElements = section.Elements[i:]
			break
		}
	}
	// log.Debugf("Preamble elements: %d", len(preambleElements))
	renderedPreambleElementsBuff := bytes.NewBuffer(make([]byte, 0))
	for i, element := range preambleElements {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render preamble element")
		}
		renderedPreambleElementsBuff.Write(renderedElement)
		if i < len(preambleElements)-1 {
			renderedPreambleElementsBuff.WriteString("\n")
		}
	}
	renderedHTMLPreamble := template.HTML(renderedPreambleElementsBuff.String())
	renderedElementsBuff := bytes.NewBuffer(make([]byte, 0))
	for i, element := range otherElements {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render section element")
		}
		renderedElementsBuff.Write(renderedElement)
		if i < len(otherElements)-1 {
			renderedElementsBuff.WriteString("\n")
		}
	}
	renderedHTMLElements := template.HTML(renderedElementsBuff.String())
	result := bytes.NewBuffer(make([]byte, 0))
	err := section1ContentTmpl.Execute(result, struct {
		Class    string
		Preamble template.HTML
		Elements template.HTML
	}{
		Class:    "sect" + strconv.Itoa(section.Heading.Level-1),
		Preamble: renderedHTMLPreamble,
		Elements: renderedHTMLElements,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section")
	}
	// log.Debugf("rendered section: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderOtherSection(ctx context.Context, section types.Section) ([]byte, error) {
	renderedHeading, err := renderHeading(ctx, section.Heading)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section heading")
	}
	renderedElementsBuff := bytes.NewBuffer(make([]byte, 0))
	for i, element := range section.Elements {
		renderedElement, err := renderElement(ctx, element)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render section element")
		}
		renderedElementsBuff.Write(renderedElement)
		if i < len(section.Elements)-1 {
			renderedElementsBuff.WriteString("\n")
		}
	}
	result := bytes.NewBuffer(make([]byte, 0))
	// select the appropriate template for the section
	var tmpl *template.Template
	if section.Heading.Level == 1 {
		tmpl = section1ContentTmpl
	} else if section.Heading.Level == 2 {
		tmpl = section2ContentTmpl
	} else {
		tmpl = otherSectionContentTmpl
	}
	renderedHTMLHeading := template.HTML(renderedHeading)
	renderedHTMLElements := template.HTML(renderedElementsBuff.String())
	err = tmpl.Execute(result, struct {
		Class    string
		Heading  template.HTML
		Elements template.HTML
	}{
		Class:    "sect" + strconv.Itoa(section.Heading.Level-1),
		Heading:  renderedHTMLHeading,
		Elements: renderedHTMLElements,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section")
	}
	// log.Debugf("rendered section: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderHeading(ctx context.Context, heading types.Heading) ([]byte, error) {
	result := bytes.NewBuffer(make([]byte, 0))
	renderedContent, err := renderElement(ctx, heading.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering heading content")
	}
	content := template.HTML(string(renderedContent))
	err = otherSectionHeaderTmpl.Execute(result, struct {
		Level   int
		ID      string
		Content template.HTML
	}{
		Level:   heading.Level,
		ID:      heading.ID.Value,
		Content: content,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering heading")
	}
	// log.Debugf("rendered heading: %s", result.Bytes())
	return result.Bytes(), nil
}
