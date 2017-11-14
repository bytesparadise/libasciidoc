package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var verbatimBlockTmpl *template.Template

// initializes the templates
func init() {
	verbatimBlockTmpl = newHTMLTemplate("listing block", `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>{{.Content}}</code></pre>
</div>
</div>`)
}

func renderDelimitedBlock(ctx *renderer.Context, block types.DelimitedBlock) ([]byte, error) {
	log.Debugf("rendering delimited block with content: %s", block.Content)
	result := bytes.NewBuffer(nil)
	tmpl, err := selectDelimitedBlockTemplate(block)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	err = tmpl.Execute(result, block)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	// log.Debugf("rendered delimited block: %s", result.Bytes())
	return result.Bytes(), nil
}

func selectDelimitedBlockTemplate(block types.DelimitedBlock) (*template.Template, error) {
	switch block.Kind {
	case types.FencedBlock, types.ListingBlock:
		return verbatimBlockTmpl, nil
	default:
		return nil, errors.Errorf("no template for block of kind %v", block.Kind)
	}
}
