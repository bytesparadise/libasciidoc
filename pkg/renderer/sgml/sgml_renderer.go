package sgml

import (
	"sync"
	text "text/template"
)

// sgmlRenderer a generic renderer for all SGML backends
type sgmlRenderer struct {
	functions   text.FuncMap
	templates   Templates
	prepareOnce sync.Once

	// Processed templates
	admonitionBlock           *text.Template
	admonitionParagraph       *text.Template
	article                   *text.Template
	articleHeader             *text.Template
	blankLine                 *text.Template
	blockImage                *text.Template
	boldText                  *text.Template
	calloutList               *text.Template
	calloutListItem           *text.Template
	calloutRef                *text.Template
	embeddedParagraph         *text.Template
	documentDetails           *text.Template
	documentAuthorDetails     *text.Template
	externalCrossReference    *text.Template
	exampleBlock              *text.Template
	fencedBlock               *text.Template
	footnote                  *text.Template
	footnoteItem              *text.Template
	footnoteRef               *text.Template
	footnoteRefPlain          *text.Template
	footnotes                 *text.Template
	iconFont                  *text.Template
	iconImage                 *text.Template
	iconText                  *text.Template
	inlineButton              *text.Template
	inlineIcon                *text.Template
	inlineImage               *text.Template
	inlineMenu                *text.Template
	internalCrossReference    *text.Template
	invalidFootnote           *text.Template
	italicText                *text.Template
	labeledList               *text.Template
	labeledListItem           *text.Template
	labeledListHorizontal     *text.Template
	labeledListHorizontalItem *text.Template
	lineBreak                 *text.Template
	link                      *text.Template
	listingBlock              *text.Template
	literalBlock              *text.Template
	manpageHeader             *text.Template
	manpageNameParagraph      *text.Template
	markdownQuoteBlock        *text.Template
	markedText                *text.Template
	monospaceText             *text.Template
	openBlock                 *text.Template
	orderedList               *text.Template
	orderedListItem           *text.Template
	paragraph                 *text.Template
	passthroughBlock          *text.Template
	preamble                  *text.Template
	qAndAList                 *text.Template
	qAndAListItem             *text.Template
	quoteBlock                *text.Template
	quoteParagraph            *text.Template
	sectionContent            *text.Template
	sectionTitle              *text.Template
	sidebarBlock              *text.Template
	sourceBlock               *text.Template
	subscriptText             *text.Template
	superscriptText           *text.Template
	table                     *text.Template
	tableBody                 *text.Template
	tableCell                 *text.Template
	tableCellBlock            *text.Template
	tableHeader               *text.Template
	tableHeaderCell           *text.Template
	tableFooter               *text.Template
	tableFooterCell           *text.Template
	tableRow                  *text.Template
	thematicBreak             *text.Template
	tocEntry                  *text.Template
	tocRoot                   *text.Template
	tocSection                *text.Template
	unorderedList             *text.Template
	unorderedListItem         *text.Template
	verseBlock                *text.Template
	verseParagraph            *text.Template
}

func (r *sgmlRenderer) prepareTemplates() error {
	tmpls := r.templates
	var err error
	r.prepareOnce.Do(func() {
		r.admonitionBlock, err = r.newTemplate("admonition-block", tmpls.AdmonitionBlock, err)
		r.admonitionParagraph, err = r.newTemplate("admonition-paragraph", tmpls.AdmonitionParagraph, err)
		r.article, err = r.newTemplate("article", tmpls.Article, err)
		r.articleHeader, err = r.newTemplate("article-header", tmpls.ArticleHeader, err)
		r.blankLine, err = r.newTemplate("blank-line", tmpls.BlankLine, err)
		r.blockImage, err = r.newTemplate("block-image", tmpls.BlockImage, err)
		r.boldText, err = r.newTemplate("bold-text", tmpls.BoldText, err)
		r.calloutList, err = r.newTemplate("callout-list", tmpls.CalloutList, err)
		r.calloutListItem, err = r.newTemplate("callout-list-item", tmpls.CalloutListItem, err)
		r.calloutRef, err = r.newTemplate("callout-ref", tmpls.CalloutRef, err)
		r.documentDetails, err = r.newTemplate("document-details", tmpls.DocumentDetails, err)
		r.documentAuthorDetails, err = r.newTemplate("document-author-details", tmpls.DocumentAuthorDetails, err)
		r.embeddedParagraph, err = r.newTemplate("embedded-paragraph", tmpls.EmbeddedParagraph, err)
		r.exampleBlock, err = r.newTemplate("example-block", tmpls.ExampleBlock, err)
		r.externalCrossReference, err = r.newTemplate("external-xref", tmpls.ExternalCrossReference, err)
		r.fencedBlock, err = r.newTemplate("fenced-block", tmpls.FencedBlock, err)
		r.footnote, err = r.newTemplate("footnote", tmpls.Footnote, err)
		r.footnoteItem, err = r.newTemplate("footnote-item", tmpls.FootnoteItem, err)
		r.footnoteRef, err = r.newTemplate("footnote-ref", tmpls.FootnoteRef, err)
		r.footnoteRefPlain, err = r.newTemplate("footnote-ref-plain", tmpls.FootnoteRefPlain, err)
		r.footnotes, err = r.newTemplate("footnotes", tmpls.Footnotes, err)
		r.iconFont, err = r.newTemplate("icon-font", tmpls.IconFont, err)
		r.iconImage, err = r.newTemplate("icon-image", tmpls.IconImage, err)
		r.iconText, err = r.newTemplate("icon-text", tmpls.IconText, err)
		r.inlineButton, err = r.newTemplate("inline-button", tmpls.InlineButton, err)
		r.inlineIcon, err = r.newTemplate("inline-icon", tmpls.InlineIcon, err)
		r.inlineImage, err = r.newTemplate("inline-image", tmpls.InlineImage, err)
		r.inlineMenu, err = r.newTemplate("inline-menu", tmpls.InlineMenu, err)
		r.internalCrossReference, err = r.newTemplate("internal-xref", tmpls.InternalCrossReference, err)
		r.invalidFootnote, err = r.newTemplate("invalid-footnote", tmpls.InvalidFootnote, err)
		r.italicText, err = r.newTemplate("italic-text", tmpls.ItalicText, err)
		r.labeledList, err = r.newTemplate("labeled-list", tmpls.LabeledList, err)
		r.labeledListItem, err = r.newTemplate("labeled-list-item", tmpls.LabeledListItem, err)
		r.labeledListHorizontal, err = r.newTemplate("labeled-list-horizontal", tmpls.LabeledListHorizontal, err)
		r.labeledListHorizontalItem, err = r.newTemplate("labeled-list-horizontal-item", tmpls.LabeledListHorizontalItem, err)
		r.lineBreak, err = r.newTemplate("line-break", tmpls.LineBreak, err)
		r.link, err = r.newTemplate("link", tmpls.Link, err)
		r.listingBlock, err = r.newTemplate("listing", tmpls.ListingBlock, err)
		r.literalBlock, err = r.newTemplate("literal-block", tmpls.LiteralBlock, err)
		r.manpageHeader, err = r.newTemplate("manpage-header", tmpls.ManpageHeader, err)
		r.manpageNameParagraph, err = r.newTemplate("manpage-name-paragraph", tmpls.ManpageNameParagraph, err)
		r.markdownQuoteBlock, err = r.newTemplate("markdown-quote-block", tmpls.MarkdownQuoteBlock, err)
		r.markedText, err = r.newTemplate("marked-text", tmpls.MarkedText, err)
		r.monospaceText, err = r.newTemplate("monospace-text", tmpls.MonospaceText, err)
		r.openBlock, err = r.newTemplate("open-block", tmpls.OpenBlock, err)
		r.orderedList, err = r.newTemplate("ordered-list", tmpls.OrderedList, err)
		r.orderedListItem, err = r.newTemplate("ordered-list-item", tmpls.OrderedListElement, err)
		r.paragraph, err = r.newTemplate("paragraph", tmpls.Paragraph, err)
		r.passthroughBlock, err = r.newTemplate("passthrough", tmpls.PassthroughBlock, err)
		r.preamble, err = r.newTemplate("preamble", tmpls.Preamble, err)
		r.qAndAList, err = r.newTemplate("qanda-list", tmpls.QAndAList, err)
		r.qAndAListItem, err = r.newTemplate("qanda-list-item", tmpls.QAndAListItem, err)
		r.quoteBlock, err = r.newTemplate("quote-block", tmpls.QuoteBlock, err)
		r.quoteParagraph, err = r.newTemplate("quote-paragraph", tmpls.QuoteParagraph, err)
		r.sectionContent, err = r.newTemplate("section-content", tmpls.SectionContent, err)
		r.sectionTitle, err = r.newTemplate("section-header", tmpls.SectionHeader, err)
		r.sidebarBlock, err = r.newTemplate("sidebar-block", tmpls.SidebarBlock, err)
		r.sourceBlock, err = r.newTemplate("source-block", tmpls.SourceBlock, err)
		r.subscriptText, err = r.newTemplate("subscript", tmpls.SubscriptText, err)
		r.superscriptText, err = r.newTemplate("superscript", tmpls.SuperscriptText, err)
		r.table, err = r.newTemplate("table", tmpls.Table, err)
		r.tableBody, err = r.newTemplate("table-body", tmpls.TableBody, err)
		r.tableCell, err = r.newTemplate("table-cell", tmpls.TableCell, err)
		r.tableCellBlock, err = r.newTemplate("table-cell-block", tmpls.TableCellBlock, err)
		r.tableHeader, err = r.newTemplate("table-header", tmpls.TableHeader, err)
		r.tableHeaderCell, err = r.newTemplate("table-header-cell", tmpls.TableHeaderCell, err)
		r.tableFooter, err = r.newTemplate("table-header", tmpls.TableFooter, err)
		r.tableFooterCell, err = r.newTemplate("table-header-cell", tmpls.TableFooterCell, err)
		r.tableRow, err = r.newTemplate("table-row", tmpls.TableRow, err)
		r.thematicBreak, err = r.newTemplate("thematic-break", tmpls.ThematicBreak, err)
		r.tocEntry, err = r.newTemplate("toc-entry", tmpls.TocEntry, err)
		r.tocRoot, err = r.newTemplate("toc-root", tmpls.TocRoot, err)
		r.tocSection, err = r.newTemplate("toc-section", tmpls.TocSection, err)
		r.unorderedList, err = r.newTemplate("unordered-list", tmpls.UnorderedList, err)
		r.unorderedListItem, err = r.newTemplate("unordered-list-item", tmpls.UnorderedListItem, err)
		r.verseBlock, err = r.newTemplate("verse", tmpls.VerseBlock, err)
		r.verseParagraph, err = r.newTemplate("verse-paragraph", tmpls.VerseParagraph, err)
	})
	return err
}
