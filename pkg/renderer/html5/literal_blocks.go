package html5

import (
	"bytes"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var literalBlockTmpl texttemplate.Template

// initializes the templates
func init() {
	literalBlockTmpl = newTextTemplate("literal block", `<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="literalblock">
{{ if .Title }}<div class="title">{{ .Title }}</div>
{{ end }}<div class="content">
<pre>{{.Content}}</pre>
</div>
</div>`)
}

func renderLiteralBlock(ctx *renderer.Context, b types.LiteralBlock) ([]byte, error) {
	log.Debugf("rendering delimited block with content: %s", b.Content)
	result := bytes.NewBuffer(nil)
	err := literalBlockTmpl.Execute(result, struct {
		ID      string
		Title   string
		Content string
	}{
		ID:      b.Attributes.GetAsString(types.AttrID),
		Title:   b.Attributes.GetAsString(types.AttrTitle),
		Content: strings.TrimSpace(b.Content),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	return result.Bytes(), nil
}
