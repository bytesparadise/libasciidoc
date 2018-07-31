package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var stringElementTmpl template.Template

// initializes the templates
func init() {
	stringElementTmpl = newHTMLTemplate("string element", "{{.}}")
}

func renderStringElement(ctx *renderer.Context, str types.StringElement) ([]byte, error) {
	// ctx.SetTrimTrailingSpaces(true) // trailing spaces can be trimmed if the last element of a line is a StringElement
	result := bytes.NewBuffer(nil)
	// content := strings.Replace(str.Content, "\t", "    ", -1)
	err := stringElementTmpl.Execute(result, str.Content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render string element")
	}
	// log.Debugf("rendered string: %s", result.Bytes())
	return result.Bytes(), nil
}
