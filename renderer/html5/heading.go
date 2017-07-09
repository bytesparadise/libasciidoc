package html5

import (
	"bytes"
	"context"
	"html/template"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var mainHeaderTmpl *template.Template
var sectionHeaderTmpl *template.Template

// initializes the templates
func init() {
	mainHeaderTmpl = newTemplate("heading", `<div id="header">
<h1>{{.Content}}</h1>
</div>`)
	sectionHeaderTmpl = newTemplate("heading", `<h{{.Level}} id="{{.ID}}">{{.Content}}</h{{.Level}}>`)
}

func renderHeading(ctx context.Context, heading types.Heading) ([]byte, error) {
	result := bytes.NewBuffer(make([]byte, 0))
	var tmpl *template.Template
	switch heading.Level {
	case 1:
		tmpl = mainHeaderTmpl
	default:
		tmpl = sectionHeaderTmpl
	}
	renderedContent, err := renderElement(ctx, heading.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering heading content")
	}
	content := template.HTML(string(renderedContent))
	err = tmpl.Execute(result, struct {
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
	log.Debugf("rendered heading: %s", result.Bytes())
	return result.Bytes(), nil
}
