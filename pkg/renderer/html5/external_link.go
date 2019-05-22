package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var linkTmpl texttemplate.Template

// initializes the templates
func init() {
	linkTmpl = newTextTemplate("external link", `<a href="{{ .URL }}"{{if .Class}} class="{{ .Class }}"{{ end }}>{{ .Text }}</a>`)
}

func renderLink(ctx *renderer.Context, l types.InlineLink) ([]byte, error) { //nolint: unparam
	result := bytes.NewBuffer(nil)
	location := l.Location.Resolve(ctx.Document.Attributes)
	var text []byte
	class := ""
	var err error
	if t, ok := l.Attributes[types.AttrInlineLinkText].(types.InlineElements); ok {
		text, err = renderElement(ctx, t)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render external link")
		}
	} else {
		class = "bare"
		text = []byte(location)
	}
	err = linkTmpl.Execute(result, struct {
		URL   string
		Text  string
		Class string
	}{
		URL:   location,
		Text:  string(text),
		Class: class,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render external link")
	}
	log.Debugf("rendered external link: %s", result.Bytes())
	return result.Bytes(), nil
}
