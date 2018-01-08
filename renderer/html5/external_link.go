package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var externalLinkTmpl *template.Template

// initializes the templates
func init() {
	externalLinkTmpl = newHTMLTemplate("external link", `<a href="{{ .URL }}"{{if .Class}} class="{{ .Class }}"{{ end }}>{{ .Text }}</a>`)
}

func renderExternalLink(ctx *renderer.Context, l *types.ExternalLink) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	text := l.Text
	class := ""
	if l.Text == "" {
		text = l.URL
		class = "bare"
	}
	err := externalLinkTmpl.Execute(result, struct {
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
