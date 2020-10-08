package sgml

import "sync"

// sgmlRenderer a generic renderer for all SGML backends
type sgmlRenderer struct {
	functions   funcMap
	templates   Templates
	prepareOnce sync.Once

	// Processed templates
	admonitionBlock           *textTemplate
	admonitionParagraph       *textTemplate
	article                   *textTemplate
	articleHeader             *textTemplate
	blankLine                 *textTemplate
	blockImage                *textTemplate
	boldText                  *textTemplate
	calloutList               *textTemplate
	calloutListItem           *textTemplate
	calloutRef                *textTemplate
	delimitedBlockParagraph   *textTemplate
	documentDetails           *textTemplate
	documentAuthorDetails     *textTemplate
	externalCrossReference    *textTemplate
	exampleBlock              *textTemplate
	fencedBlock               *textTemplate
	footnote                  *textTemplate
	footnoteItem              *textTemplate
	footnoteRef               *textTemplate
	footnoteRefPlain          *textTemplate
	footnotes                 *textTemplate
	iconFont                  *textTemplate
	iconImage                 *textTemplate
	iconText                  *textTemplate
	inlineIcon                *textTemplate
	inlineImage               *textTemplate
	internalCrossReference    *textTemplate
	invalidFootnote           *textTemplate
	italicText                *textTemplate
	labeledList               *textTemplate
	labeledListItem           *textTemplate
	labeledListHorizontal     *textTemplate
	labeledListHorizontalItem *textTemplate
	lineBreak                 *textTemplate
	link                      *textTemplate
	listingBlock              *textTemplate
	literalBlock              *textTemplate
	manpageHeader             *textTemplate
	manpageNameParagraph      *textTemplate
	markdownQuoteBlock        *textTemplate
	markedText                *textTemplate
	monospaceText             *textTemplate
	orderedList               *textTemplate
	orderedListItem           *textTemplate
	paragraph                 *textTemplate
	passthroughBlock          *textTemplate
	preamble                  *textTemplate
	predefinedAttribute       *textTemplate
	qAndAList                 *textTemplate
	qAndAListItem             *textTemplate
	quoteBlock                *textTemplate
	quoteParagraph            *textTemplate
	sectionContent            *textTemplate
	sectionHeader             *textTemplate
	sidebarBlock              *textTemplate
	sourceBlock               *textTemplate
	specialCharacter          *textTemplate
	stringElement             *textTemplate
	subscriptText             *textTemplate
	superscriptText           *textTemplate
	table                     *textTemplate
	tableBody                 *textTemplate
	tableCell                 *textTemplate
	tableHeader               *textTemplate
	tableHeaderCell           *textTemplate
	tableRow                  *textTemplate
	thematicBreak             *textTemplate
	tocEntry                  *textTemplate
	tocRoot                   *textTemplate
	tocSection                *textTemplate
	unorderedList             *textTemplate
	unorderedListItem         *textTemplate
	verbatimLine              *textTemplate
	verseBlock                *textTemplate
	verseParagraph            *textTemplate
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
		r.delimitedBlockParagraph, err = r.newTemplate("delimited-block-paragraph", tmpls.DelimitedBlockParagraph, err)
		r.documentDetails, err = r.newTemplate("document-details", tmpls.DocumentDetails, err)
		r.documentAuthorDetails, err = r.newTemplate("document-author-details", tmpls.DocumentAuthorDetails, err)
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
		r.inlineIcon, err = r.newTemplate("inline-icon", tmpls.InlineIcon, err)
		r.inlineImage, err = r.newTemplate("inline-image", tmpls.InlineImage, err)
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
		r.orderedList, err = r.newTemplate("ordered-list", tmpls.OrderedList, err)
		r.orderedListItem, err = r.newTemplate("ordered-list-item", tmpls.OrderedListItem, err)
		r.paragraph, err = r.newTemplate("paragraph", tmpls.Paragraph, err)
		r.passthroughBlock, err = r.newTemplate("passthrough", tmpls.PassthroughBlock, err)
		r.preamble, err = r.newTemplate("preamble", tmpls.Preamble, err)
		r.predefinedAttribute, err = r.newTemplate("predefined attribute", tmpls.PredefinedAttribute, err)
		r.qAndAList, err = r.newTemplate("qanda-list", tmpls.QAndAList, err)
		r.qAndAListItem, err = r.newTemplate("qanda-list-item", tmpls.QAndAListItem, err)
		r.quoteBlock, err = r.newTemplate("quote-block", tmpls.QuoteBlock, err)
		r.quoteParagraph, err = r.newTemplate("quote-paragraph", tmpls.QuoteParagraph, err)
		r.sectionContent, err = r.newTemplate("section-content", tmpls.SectionContent, err)
		r.sectionHeader, err = r.newTemplate("section-header", tmpls.SectionHeader, err)
		r.stringElement, err = r.newTemplate("string-element", tmpls.StringElement, err)
		r.sidebarBlock, err = r.newTemplate("sidebar-block", tmpls.SidebarBlock, err)
		r.sourceBlock, err = r.newTemplate("source-block", tmpls.SourceBlock, err)
		r.specialCharacter, err = r.newTemplate("special-character", tmpls.SpecialCharacter, err)
		r.subscriptText, err = r.newTemplate("subscript", tmpls.SubscriptText, err)
		r.superscriptText, err = r.newTemplate("superscript", tmpls.SuperscriptText, err)
		r.table, err = r.newTemplate("table", tmpls.Table, err)
		r.tableBody, err = r.newTemplate("table-body", tmpls.TableBody, err)
		r.tableCell, err = r.newTemplate("table-cell", tmpls.TableCell, err)
		r.tableHeader, err = r.newTemplate("table-header", tmpls.TableHeader, err)
		r.tableHeaderCell, err = r.newTemplate("table-header-cell", tmpls.TableHeaderCell, err)
		r.tableRow, err = r.newTemplate("table-row", tmpls.TableRow, err)
		r.thematicBreak, err = r.newTemplate("thematic-break", tmpls.ThematicBreak, err)
		r.tocEntry, err = r.newTemplate("toc-entry", tmpls.TocEntry, err)
		r.tocRoot, err = r.newTemplate("toc-root", tmpls.TocRoot, err)
		r.tocSection, err = r.newTemplate("toc-section", tmpls.TocSection, err)
		r.unorderedList, err = r.newTemplate("unordered-list", tmpls.UnorderedList, err)
		r.unorderedListItem, err = r.newTemplate("unordered-list-item", tmpls.UnorderedListItem, err)
		r.verbatimLine, err = r.newTemplate("verbatim-line", tmpls.VerbatimLine, err)
		r.verseBlock, err = r.newTemplate("verse", tmpls.VerseBlock, err)
		r.verseParagraph, err = r.newTemplate("verse-paragraph", tmpls.VerseParagraph, err)
	})
	return err
}
