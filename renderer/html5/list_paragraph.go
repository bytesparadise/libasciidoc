package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const listParagraphTmpl = `{{ $ctx := .Context }}{{ with .Data }}<p>{{ $lines := .Lines }}{{ range $index, $line := $lines }}{{ renderElement $ctx $line | printf "%s" }}{{ if notLastItem $index $lines }}{{ print "\n" }}{{ end }}{{ end }}</p>{{ end }}`

func renderListParagraph(ctx *renderer.Context, p *types.ListParagraph) ([]byte, error) {
	// TODO: move this to init
	t := texttemplate.New("list paragraph")
	t.Funcs(texttemplate.FuncMap{
		"renderElement": renderElement,
		"notLastItem":   notLastItem,
	})
	var err error
	t, err = t.Parse(listParagraphTmpl)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse list paragraph template")
	}

	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err = t.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    p,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render list paragraph")
	}

	log.Debugf("rendered list paragraph: %s", result.Bytes())
	return result.Bytes(), nil
}
