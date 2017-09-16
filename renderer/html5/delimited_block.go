package html5

import (
	"bytes"
	"html/template"

	asciidoc "github.com/bytesparadise/libasciidoc/context"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var sourceBlockTmpl *template.Template

// initializes the templates
func init() {
	sourceBlockTmpl = newHTMLTemplate("delimited source block", `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>{{.Content}}</code></pre>
</div>
</div>`)
}

func renderDelimitedBlock(ctx asciidoc.Context, block types.DelimitedBlock) ([]byte, error) {
	log.Debugf("rendering delimited block with content: %s", block.Content)
	result := bytes.NewBuffer(nil)
	err := sourceBlockTmpl.Execute(result, block)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	// log.Debugf("rendered delimited block: %s", result.Bytes())
	return result.Bytes(), nil
}
