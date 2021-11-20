package xhtml5

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

var templates = html5.Templates()

var defaultRenderer sgml.Renderer

func init() {
	templates = html5.Templates()

	// XHTML5 overrides of HTML5.
	templates.Article = articleTmpl
	templates.BlankLine = blankLineTmpl
	templates.BlockImage = blockImageTmpl
	templates.LineBreak = lineBreakTmpl
	templates.DocumentAuthorDetails = documentAuthorDetailsTmpl
	templates.DocumentDetails = documentDetailsTmpl
	templates.Footnotes = footnotesTmpl
	templates.IconImage = iconImageTmpl
	templates.InlineImage = inlineImageTmpl
	templates.LabeledListHorizontalItem = labeledListHorizontalItemTmpl
	templates.Table = tableTmpl
	templates.ThematicBreak = thematicBreakTmpl
	templates.QuoteBlock = quoteBlockTmpl
	templates.QuoteParagraph = quoteParagraphTmpl
	templates.VerseBlock = verseBlockTmpl
	templates.VerseParagraph = verseParagraphTmpl

	// NB: This is fast, and doesn't including parsing.
	defaultRenderer = sgml.NewRenderer(templates)
}

// Render renders the document to the output, using a default instance
// of the renderer, with default templates.
func Render(ctx *renderer.Context, doc *types.Document, output io.Writer) (types.Metadata, error) {
	return defaultRenderer.Render(ctx, doc, output)
}

// Templates returns the default Templates use for HTML5.  It may be useful
// for derived implementations.
func Templates() sgml.Templates {
	return templates
}
