package html5

import (
	"bytes"
	"context"
	"html/template"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

var stringElementTmpl *template.Template

// initializes the templates
func init() {
	stringElementTmpl = newHTMLTemplate("string element", "{{.}}")
}

func renderStringElement(ctx context.Context, str types.StringElement) ([]byte, error) {
	result := bytes.NewBuffer(make([]byte, 0))
	err := stringElementTmpl.Execute(result, str.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render string element")
	}
	// log.Debugf("rendered string: %s", result.Bytes())
	return result.Bytes(), nil
}
