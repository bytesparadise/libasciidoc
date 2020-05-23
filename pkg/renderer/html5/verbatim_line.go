package html5

import (
	"bytes"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

var verbatimLineTmpl texttemplate.Template

// initializes the templates
func init() {
	verbatimLineTmpl = newTextTemplate("verbatim line", `{{ if .Callouts}}{{ escape .Content }}{{ else }}{{ .Content | escape | trim }}{{ end }}{{ range $i, $c := .Callouts }}<b class="conum">({{ $c.Ref }})</b>{{ end }}`,
		texttemplate.FuncMap{
			"escape": EscapeString,
			"trim": func(s string) string {
				return strings.TrimRight(s, " ")
			},
		})
}

func renderVerbatimLine(l types.VerbatimLine) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	if err := verbatimLineTmpl.Execute(result, l); err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}
