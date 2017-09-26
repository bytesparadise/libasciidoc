package html5

import (
	"bytes"
	"html/template"

	asciidoc "github.com/bytesparadise/libasciidoc/context"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var literalBlockTmpl *template.Template

// initializes the templates
func init() {
	literalBlockTmpl = newHTMLTemplate("literal block", `<div class="literalblock">
<div class="content">
<pre>{{.Content}}</pre>
</div>
</div>`)
}

func renderLiteralBlock(ctx asciidoc.Context, block types.LiteralBlock) ([]byte, error) {
	log.Debugf("rendering delimited block with content: %s", block.Content)
	result := bytes.NewBuffer(nil)
	err := literalBlockTmpl.Execute(result, block)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	// log.Debugf("rendered delimited block: %s", result.Bytes())
	return result.Bytes(), nil
}
