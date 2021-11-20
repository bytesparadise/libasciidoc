package sgml

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderDelimitedBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	switch b.Kind {
	case types.Example:
		if b.Attributes.Has(types.AttrStyle) {
			return r.renderAdmonitionBlock(ctx, b)
		}
		return r.renderExampleBlock(ctx, b)
	case types.Fenced:
		return r.renderFencedBlock(ctx, b)
	case types.Literal:
		return r.renderLiteralBlock(ctx, b)
	case types.Listing:
		return r.renderListingBlock(ctx, b)
	case types.MarkdownQuote:
		return r.renderMarkdownQuoteBlock(ctx, b)
	case types.Passthrough:
		return r.renderPassthroughBlock(ctx, b)
	case types.Quote:
		return r.renderQuoteBlock(ctx, b)
	case types.Verse:
		return r.renderVerseBlock(ctx, b)
	case types.Sidebar:
		return r.renderSidebarBlock(ctx, b)
	default:
		return "", fmt.Errorf("unsupported kind of delimited block: '%s'", b.Kind)
	}

}
