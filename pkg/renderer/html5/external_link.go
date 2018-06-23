package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var linkTmpl template.Template

// initializes the templates
func init() {
	linkTmpl = newHTMLTemplate("external link", `<a href="{{ .URL }}"{{if .Class}} class="{{ .Class }}"{{ end }}>{{ .Text }}</a>`)
}

func renderLink(ctx *renderer.Context, l types.Link) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	// text := l.Text
	text := l.Text()
	class := ""
	if text == "" {
		text = l.URL
		class = "bare"
	}
	err := linkTmpl.Execute(result, struct {
		URL   string
		Text  string
		Class string
	}{
		URL:   l.URL,
		Text:  text,
		Class: class,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render external link")
	}
	log.Debugf("rendered external link: %s", result.Bytes())
	return result.Bytes(), nil
}
