package html5

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

var templates = sgml.Templates{
	AdmonitionBlock:           admonitionBlockTmpl,
	AdmonitionParagraph:       admonitionParagraphTmpl,
	Article:                   articleTmpl,
	ArticleHeader:             articleHeaderTmpl,
	BlankLine:                 blankLineTmpl,
	BlockImage:                blockImageTmpl,
	BoldText:                  boldTextTmpl,
	CalloutList:               calloutListTmpl,
	CalloutListItem:           calloutListItemTmpl,
	CalloutRef:                calloutRefTmpl,
	DelimitedBlockParagraph:   delimitedBlockParagraphTmpl,
	DocumentDetails:           documentDetailsTmpl,
	DocumentAuthorDetails:     documentAuthorDetailsTmpl,
	ExternalCrossReference:    externalCrossReferenceTmpl,
	ExampleBlock:              exampleBlockTmpl,
	FencedBlock:               fencedBlockTmpl,
	Footnote:                  footnoteTmpl,
	FootnoteItem:              footnoteItemTmpl,
	FootnoteRef:               footnoteRefTmpl,
	FootnoteRefPlain:          footnoteRefPlainTmpl,
	Footnotes:                 footnotesTmpl,
	IconFont:                  iconFontTmpl,
	IconImage:                 iconImageTmpl,
	IconText:                  iconTextTmpl,
	ImageCaption:              imageCaptionTmpl,
	InlineIcon:                inlineIconTmpl,
	InlineImage:               inlineImageTmpl,
	InternalCrossReference:    internalCrossReferenceTmpl,
	InvalidFootnote:           invalidFootnoteTmpl,
	ItalicText:                italicTextTmpl,
	LabeledList:               labeledListTmpl,
	LabeledListItem:           labeledListItemTmpl,
	LabeledListHorizontal:     labeledListHorizontalTmpl,
	LabeledListHorizontalItem: labeledListHorizontalItemTmpl,
	LineBreak:                 lineBreakTmpl,
	Link:                      linkTmpl,
	ListingBlock:              listingBlockTmpl,
	LiteralBlock:              literalBlockTmpl,
	ManpageHeader:             manpageHeaderTmpl,
	ManpageNameParagraph:      manpageNameParagraphTmpl,
	MarkedText:                markedTextTmpl,
	MonospaceText:             monospaceTextTmpl,
	OrderedList:               orderedListTmpl,
	OrderedListItem:           orderedListItemTmpl,
	PassthroughBlock:          pssThroughBlock,
	Paragraph:                 paragraphTmpl,
	Preamble:                  preambleTmpl,
	QAndAList:                 qAndAListTmpl,
	QAndAListItem:             qAndAListItemTmpl,
	QuoteBlock:                quoteBlockTmpl,
	QuoteParagraph:            quoteParagraphTmpl,
	SectionContent:            sectionContentTmpl,
	SectionHeader:             sectionHeaderTmpl,
	SidebarBlock:              sidebarBlockTmpl,
	SourceBlock:               sourceBlockTmpl,
	StringElement:             stringTmpl,
	SubscriptText:             subscriptTextTmpl,
	SuperscriptText:           superscriptTextTmpl,
	Table:                     tableTmpl,
	TableBody:                 tableBodyTmpl,
	TableCell:                 tableCellTmpl,
	TableHeader:               tableHeaderTmpl,
	TableHeaderCell:           tableHeaderCellTmpl,
	TableRow:                  tableRowTmpl,
	ThematicBreak:             thematicBreakTmpl,
	TocRoot:                   tocRootTmpl,
	TocEntry:                  tocEntryTmpl,
	TocSection:                tocSectionTmpl,
	UnorderedList:             unorderedListTmpl,
	UnorderedListItem:         unorderedListItemTmpl,
	VerbatimLine:              verbatimLineTmpl,
	VerseBlock:                verseBlockTmpl,
	VerseParagraph:            verseParagraphTmpl,
}

var defaultRenderer sgml.Renderer

func init() {
	// NB: This is fast, and doesn't including parsing.
	defaultRenderer = sgml.NewRenderer(templates)
}

// Render renders the document to the output, using a default instance
// of the renderer, with default templates.
func Render(ctx *renderer.Context, doc types.Document, output io.Writer) (types.Metadata, error) {
	return defaultRenderer.Render(ctx, doc, output)
}

// Templates returns the default Templates use for HTML5.  It may be useful
// for derived implementations.
func Templates() sgml.Templates {
	return templates
}
