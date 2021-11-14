package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderPreamble(ctx *renderer.Context, p *types.Preamble) (string, error) {
	// log.Debugf("rendering preamble...")
	result := &strings.Builder{}
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
	err = r.preamble.Execute(result, struct {
		Context *renderer.Context
		Wrapper bool
		Content string
		ToC     string
	}{
		Context: ctx,
		Wrapper: ctx.HasHeader,
		Content: string(content),
		ToC:     string(toc),
	})
	if err != nil {
		return "", errors.Wrap(err, "error while rendering preamble")
	}
	// log.Debugf("rendered preamble: %s", result.Bytes())
	return result.String(), nil
}
