package html5

import (
	"bytes"
	"context"
	"html/template"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var sourceBlockTmpl *template.Template

// initializes the templates
func init() {
	sourceBlockTmpl = newTemplate("delimited source block", `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>{{.Content}}</code></pre>
</div>
</div>`)
}

func renderDelimitedBlock(ctx context.Context, block types.DelimitedBlock) ([]byte, error) {
	log.Debugf("rendering delimited block with content: %s", block.Content)
	result := bytes.NewBuffer(make([]byte, 0))
	err := sourceBlockTmpl.Execute(result, block)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	// log.Debugf("rendered delimited block: %s", result.Bytes())
	return result.Bytes(), nil
}
