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
	externalLinkTmpl = newHTMLTemplate("external link", `<a href="{{ .URL }}">{{ .Text }}</a>`)
}

func renderExternalLink(ctx *renderer.Context, l *types.ExternalLink) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := externalLinkTmpl.Execute(result, l)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render external link")
	}
	log.Debugf("rendered external link: %s", result.Bytes())
	return result.Bytes(), nil
}
