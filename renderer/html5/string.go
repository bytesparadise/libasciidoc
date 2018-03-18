package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

var stringElementTmpl template.Template

// initializes the templates
func init() {
	stringElementTmpl = newHTMLTemplate("string element", "{{.}}")
}

func renderStringElement(ctx *renderer.Context, str types.StringElement) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := stringElementTmpl.Execute(result, str.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render string element")
	}
	// log.Debugf("rendered string: %s", result.Bytes())
	return result.Bytes(), nil
}
