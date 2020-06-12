package html5

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

var templates = sgml.Templates{
	AdmonitionBlock:         admonitionBlockTmpl,
	AdmonitionParagraph:     admonitionParagraphTmpl,
	Article:                 articleTmpl,
	ArticleHeader:           articleHeaderTmpl,
	BlankLine:               blankLineTmpl,
	BlockImage:              blockImageTmpl,
	BoldText:                boldTextTmpl,
	CalloutList:             calloutListTmpl,
	DelimitedBlockParagraph: delimitedBlockParagraphTmpl,
	DocumentDetails:         documentDetailsTmpl,
	DocumentAuthorDetails:   documentAuthorDetailsTmpl,
	ExternalCrossReference:  externalCrossReferenceTmpl,
	ExampleBlock:            exampleBlockTmpl,
	FencedBlock:             fencedBlockTmpl,
	Footnote:                footnoteTmpl,
	FootnoteRef:             footnoteRefTmpl,
	FootnoteRefPlain:        footnoteRefPlainTmpl,
	Footnotes:               footnotesTmpl,
	InlineImage:             inlineImageTmpl,
	InternalCrossReference:  internalCrossReferenceTmpl,
	InvalidFootnote:         invalidFootnoteTmpl,
	ItalicText:              italicTextTmpl,
	LabeledList:             labeledListTmpl,
	LabeledListHorizontal:   labeledListHorizontalTmpl,
	LineBreak:               lineBreakTmpl,
	Link:                    linkTmpl,
	ListingBlock:            listingBlockTmpl,
	LiteralBlock:            literalBlockTmpl,
	ManpageHeader:           manpageHeaderTmpl,
	ManpageNameParagraph:    manpageNameParagraphTmpl,
	MonospaceText:           monospaceTextTmpl,
	OrderedList:             orderedListTmpl,
	PassthroughBlock:        pssThroughBlock,
	Paragraph:               paragraphTmpl,
	Preamble:                preambleTmpl,
	QAndAList:               qAndAListTmpl,
	QuoteBlock:              quoteBlockTmpl,
	QuoteParagraph:          quoteParagraphTmpl,
	SectionContent:          sectionContentTmpl,
	SectionHeader:           sectionHeaderTmpl,
	SectionOne:              sectionOneTmpl,
	SidebarBlock:            sidebarBlockTmpl,
	SourceBlock:             sourceBlockTmpl,
	SourceBlockContent:      sourceBlockContentTmpl,
	SourceParagraph:         sourceParagraphTmpl,
	StringElement:           stringTmpl,
	SubscriptText:           subscriptTextTmpl,
	SuperscriptText:         superscriptTextTmpl,
	Table:                   tableTmpl,
	TocRoot:                 tocRootTmpl,
	TocSection:              tocSectionTmpl,
	UnorderedList:           unorderedListTmpl,
	VerbatimLine:            verbatimLineTmpl,
	VerseBlock:              verseBlockTmpl,
	VerseBlockParagraph:     verseBlockParagraphTmpl,
	VerseParagraph:          verseParagraphTmpl,
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
