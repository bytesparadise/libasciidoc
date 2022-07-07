package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderPreamble(ctx *context, p *types.Preamble) (string, error) {
	// log.Debugf("rendering preamble...")
	// the <div id="preamble"> wrapper is only necessary
	// if the document has a section 0

	content, err := r.renderElements(ctx, p.Elements)
	if err != nil {
		return "", errors.Wrap(err, "error rendering preamble elements")
	}
	toc, err := r.renderTableOfContents(ctx, p.TableOfContents)
	if err != nil {
		return "", errors.Wrap(err, "error rendering preamble elements")
	}
	return r.execute(r.preamble, struct {
		Context *context
		Wrapper bool
		Content string
		ToC     string
	}{
		Context: ctx,
		Wrapper: ctx.hasHeader,
		Content: string(content),
		ToC:     string(toc),
	})
}
