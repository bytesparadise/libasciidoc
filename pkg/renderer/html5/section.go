package html5

import (
	"bytes"
	"strconv"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var preambleTmpl texttemplate.Template
var sectionHeaderTmpl texttemplate.Template
var section1ContentTmpl texttemplate.Template
var otherSectionContentTmpl texttemplate.Template

// initializes the templates
func init() {
	preambleTmpl = newTextTemplate("preamble",
		`{{ $ctx := .Context }}{{ with .Data }}{{ if .Wrapper }}<div id="preamble">
<div class="sectionbody">
{{ end }}{{ renderElements $ctx .Elements | printf "%s" }}{{ if .Wrapper }}
</div>
</div>{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
		})
	section1ContentTmpl = newTextTemplate("section 1",
		`{{ $ctx := .Context }}{{ with .Data }}<div class="{{ .Class }}">
{{ .SectionTitle }}
<div class="sectionbody">{{ $elements := renderElements $ctx .Elements | printf "%s" }}{{ if $elements }}
{{ $elements }}{{ end }}
</div>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
		})
	otherSectionContentTmpl = newTextTemplate("other section",
		`{{ $ctx := .Context }}{{ with .Data }}<div class="{{ .Class }}">
{{ .SectionTitle }}{{ $elements := renderElements $ctx .Elements | printf "%s" }}{{ if $elements }}
{{ $elements }}{{ end }}
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
		})
	sectionHeaderTmpl = newTextTemplate("other sectionTitle",
		`<h{{ .Level }} id="{{ .ID }}">{{ .Content }}</h{{ .Level }}>`)
}

func renderPreamble(ctx *renderer.Context, p types.Preamble) ([]byte, error) {
	log.Debugf("rendering preamble...")
	result := bytes.NewBuffer(nil)
	// the <div id="preamble"> wrapper is only necessary
	// if the document has a section 0
	wrapper := false
	if _, ok := ctx.Document.Title(); ok {
		wrapper = true
	}
	err := preambleTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			Wrapper  bool
			Elements []interface{}
		}{
			Wrapper:  wrapper,
			Elements: p.Elements,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering preamble")
	}
	log.Debugf("rendered preamble: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderSection(ctx *renderer.Context, s types.Section) ([]byte, error) {
	log.Debugf("rendering section level %d", s.Level)
	renderedSectionTitle, err := renderSectionTitle(ctx, s.Level, s.Title)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section")
	}
	// renderedSectionElements, err := renderElements(ctx, s.Elements)
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "error while rendering section")
	// }
	result := bytes.NewBuffer(nil)
	// select the appropriate template for the section
	var tmpl texttemplate.Template
	if s.Level == 1 {
		tmpl = section1ContentTmpl
	} else {
		tmpl = otherSectionContentTmpl
	}
	err = tmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			Class        string
			SectionTitle string
			Elements     []interface{}
		}{
			Class:        "sect" + strconv.Itoa(s.Level),
			SectionTitle: renderedSectionTitle,
			Elements:     s.Elements,
		}})
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering section")
	}
	log.Debugf("rendered section: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderSectionTitle(ctx *renderer.Context, level int, sectionTitle types.SectionTitle) (string, error) {
	result := bytes.NewBuffer(nil)
	renderedContent, err := renderElement(ctx, sectionTitle.Elements)
	if err != nil {
		return "", errors.Wrapf(err, "error while rendering sectionTitle content")
	}
	renderedContentStr := strings.TrimSpace(string(renderedContent))
	id := generateID(ctx, sectionTitle.Attributes)
	err = sectionHeaderTmpl.Execute(result, struct {
		Level   int
		ID      string
		Content string
	}{
		Level:   level + 1,
		ID:      id,
		Content: renderedContentStr,
	})
	if err != nil {
		return "", errors.Wrapf(err, "error while rendering sectionTitle")
	}
	// log.Debugf("rendered sectionTitle: %s", result.Bytes())
	return result.String(), nil
}
