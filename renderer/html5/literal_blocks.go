package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var literalBlockTmpl template.Template

// initializes the templates
func init() {
	literalBlockTmpl = newHTMLTemplate("literal block", `<div class="literalblock">
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
