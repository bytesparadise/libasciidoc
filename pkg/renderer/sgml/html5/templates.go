package html5

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

var templates = sgml.Templates{
	AdmonitionBlock:              admonitionBlockTmpl,
	AdmonitionParagraph:          admonitionParagraphTmpl,
	Article:                      articleTmpl,
	ArticleHeader:                articleHeaderTmpl,
	BlockImage:                   blockImageTmpl,
	BoldText:                     boldTextTmpl,
	CalloutList:                  calloutListTmpl,
	CalloutListElement:           calloutListElementTmpl,
	CalloutRef:                   calloutRefTmpl,
	DocumentDetails:              documentDetailsTmpl,
	DocumentAuthorDetails:        documentAuthorDetailsTmpl,
	EmbeddedParagraph:            embeddedParagraphTmpl,
	ExternalCrossReference:       externalCrossReferenceTmpl,
	ExampleBlock:                 exampleBlockTmpl,
	FencedBlock:                  fencedBlockTmpl,
	Footnote:                     footnoteTmpl,
	FootnoteElement:              footnoteElementTmpl,
	FootnoteRef:                  footnoteRefTmpl,
	FootnoteRefPlain:             footnoteRefPlainTmpl,
	Footnotes:                    footnotesTmpl,
	IconFont:                     iconFontTmpl,
	IconImage:                    iconImageTmpl,
	IconText:                     iconTextTmpl,
	InlineButton:                 inlineButtonTmpl,
	InlineIcon:                   inlineIconTmpl,
	InlineImage:                  inlineImageTmpl,
	InlineMenu:                   inlineMenuTmpl,
	InternalCrossReference:       internalCrossReferenceTmpl,
	InvalidFootnote:              invalidFootnoteTmpl,
	ItalicText:                   italicTextTmpl,
	LabeledList:                  labeledListTmpl,
	LabeledListElement:           labeledListElementTmpl,
	LabeledListHorizontal:        labeledListHorizontalTmpl,
	LabeledListHorizontalElement: labeledListHorizontalElementTmpl,
	LineBreak:                    lineBreakTmpl,
	Link:                         linkTmpl,
	ListingBlock:                 listingBlockTmpl,
	LiteralBlock:                 literalBlockTmpl,
	ManpageHeader:                manpageHeaderTmpl,
	ManpageNameParagraph:         manpageNameParagraphTmpl,
	MarkdownQuoteBlock:           markdownQuoteBlockTmpl,
	MarkedText:                   markedTextTmpl,
	MonospaceText:                monospaceTextTmpl,
	OpenBlock:                    openBlockTmpl,
	OrderedList:                  orderedListTmpl,
	OrderedListElement:           orderedListElementTmpl,
	PassthroughBlock:             passthroughBlock,
	Paragraph:                    paragraphTmpl,
	Preamble:                     preambleTmpl,
	QAndAList:                    qAndAListTmpl,
	QAndAListElement:             qAndAListElementTmpl,
	QuoteBlock:                   quoteBlockTmpl,
	QuoteParagraph:               quoteParagraphTmpl,
	SectionContent:               sectionContentTmpl,
	SectionTitle:                 sectionTitleTmpl,
	SidebarBlock:                 sidebarBlockTmpl,
	SourceBlock:                  sourceBlockTmpl,
	SubscriptText:                subscriptTextTmpl,
	SuperscriptText:              superscriptTextTmpl,
	Table:                        tableTmpl,
	TableBody:                    tableBodyTmpl,
	TableCell:                    tableCellTmpl,
	TableCellBlock:               tableCellBlockTmpl,
	TableHeader:                  tableHeaderTmpl,
	TableHeaderCell:              tableHeaderCellTmpl,
	TableFooter:                  tableFooterTmpl,
	TableFooterCell:              tableFooterCellTmpl,
	TableRow:                     tableRowTmpl,
	ThematicBreak:                thematicBreakTmpl,
	TocRoot:                      tocRootTmpl,
	TocEntry:                     tocEntryTmpl,
	TocSection:                   tocSectionTmpl,
	UnorderedList:                unorderedListTmpl,
	UnorderedListElement:         unorderedListElementTmpl,
	VerseBlock:                   verseBlockTmpl,
	VerseParagraph:               verseParagraphTmpl,
}

var defaultRenderer sgml.Renderer

func init() {
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
