package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var listParagraphTmpl texttemplate.Template

// initializes the templates
func init() {
	listParagraphTmpl = newTextTemplate("list paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}<p>{{ $lines := .Lines }}{{ range $index, $line := $lines }}{{ renderElement $ctx $line | printf "%s" }}{{ if includeNewline $ctx $index $lines }}{{ print "\n" }}{{ end }}{{ end }}</p>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"includeNewline": includeNewline,
		})
}

func renderListParagraph(ctx *renderer.Context, p types.ListParagraph) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err := listParagraphTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    p,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render list paragraph")
	}

	log.Debugf("rendered list paragraph: %s", result.Bytes())
	return result.Bytes(), nil
}
