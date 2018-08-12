package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var literalBlockTmpl texttemplate.Template

// initializes the templates
func init() {
	literalBlockTmpl = newTextTemplate("literal block", `<div class="literalblock">
<div class="content">
<pre>{{.Content}}</pre>
</div>
</div>`)
}

func renderLiteralBlock(ctx *renderer.Context, b types.LiteralBlock) ([]byte, error) {
	log.Debugf("rendering delimited block with content: %s", b.Content)
	result := bytes.NewBuffer(nil)
	err := literalBlockTmpl.Execute(result, b)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	return result.Bytes(), nil
}
