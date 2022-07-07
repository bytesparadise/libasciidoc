package xhtml5

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// Render renders the document to the output, using a default instance
// of the renderer, with default templates.
func Render(doc *types.Document, config *configuration.Configuration, output io.Writer) (types.Metadata, error) {
	templates := html5.Templates()
	// XHTML5 overrides of HTML5.
	templates.Article = articleTmpl
	templates.BlockImage = blockImageTmpl
	templates.LineBreak = lineBreakTmpl
	templates.DocumentAuthorDetails = documentAuthorDetailsTmpl
	templates.DocumentDetails = documentDetailsTmpl
	templates.Footnotes = footnotesTmpl
	templates.IconImage = iconImageTmpl
	templates.InlineImage = inlineImageTmpl
	templates.LabeledListHorizontalElement = labeledListHorizontalItemTmpl
	templates.Table = tableTmpl
	templates.ThematicBreak = thematicBreakTmpl
	templates.QuoteBlock = quoteBlockTmpl
	templates.QuoteParagraph = quoteParagraphTmpl
	templates.VerseBlock = verseBlockTmpl
	templates.VerseParagraph = verseParagraphTmpl

	return sgml.Render(doc, config, output, templates)
}
