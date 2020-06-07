package sgml

import "sync"

// sgmlRenderer a generic renderer for all SGML backends
type sgmlRenderer struct {
	functions   funcMap
	templates   Templates
	prepareOnce sync.Once

	// Processed templates
	admonitionBlock         *textTemplate
	admonitionParagraph     *textTemplate
	article                 *textTemplate
	articleHeader           *textTemplate
	blankLine               *textTemplate
	blockImage              *textTemplate
	boldText                *textTemplate
	calloutList             *textTemplate
	delimitedBlockParagraph *textTemplate
	documentDetails         *textTemplate
	documentAuthorDetails   *textTemplate
	externalCrossReference  *textTemplate
	exampleBlock            *textTemplate
	fencedBlock             *textTemplate
	footnote                *textTemplate
	footnoteRef             *textTemplate
	footnoteRefPlain        *textTemplate
	footnotes               *textTemplate
	inlineImage             *textTemplate
	internalCrossReference  *textTemplate
	invalidFootnote         *textTemplate
	italicText              *textTemplate
	labeledList             *textTemplate
	labeledListHorizontal   *textTemplate
	lineBreak               *textTemplate
	link                    *textTemplate
	listingBlock            *textTemplate
	literalBlock            *textTemplate
	manpageHeader           *textTemplate
	manpageNameParagraph    *textTemplate
	monospaceText           *textTemplate
	orderedList             *textTemplate
	paragraph               *textTemplate
	passthroughBlock        *textTemplate
	preamble                *textTemplate
	qAndAList               *textTemplate
	quoteBlock              *textTemplate
	quoteParagraph          *textTemplate
	sectionContent          *textTemplate
	sectionHeader           *textTemplate
	sectionOne              *textTemplate
	sidebarBlock            *textTemplate
	sourceBlock             *textTemplate
	sourceBlockContent      *textTemplate
	sourceParagraph         *textTemplate
	stringElement           *textTemplate
	subscriptText           *textTemplate
	superscriptText         *textTemplate
	table                   *textTemplate
	tocRoot                 *textTemplate
	tocSection              *textTemplate
	unorderedList           *textTemplate
	verbatimLine            *textTemplate
	verseBlock              *textTemplate
	verseBlockParagraph     *textTemplate
	verseParagraph          *textTemplate
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
		r.delimitedBlockParagraph, err = r.newTemplate("delimited-block-paragraph", tmpls.DelimitedBlockParagraph, err)
		r.documentDetails, err = r.newTemplate("document-details", tmpls.DocumentDetails, err)
		r.documentAuthorDetails, err = r.newTemplate("document-author-details", tmpls.DocumentAuthorDetails, err)
		r.exampleBlock, err = r.newTemplate("example-block", tmpls.ExampleBlock, err)
		r.externalCrossReference, err = r.newTemplate("external-xref", tmpls.ExternalCrossReference, err)
		r.fencedBlock, err = r.newTemplate("fenced-block", tmpls.FencedBlock, err)
		r.footnote, err = r.newTemplate("footnote", tmpls.Footnote, err)
		r.footnoteRef, err = r.newTemplate("footnote-ref", tmpls.FootnoteRef, err)
		r.footnoteRefPlain, err = r.newTemplate("footnote-ref-plain", tmpls.FootnoteRefPlain, err)
		r.footnotes, err = r.newTemplate("footnotes", tmpls.Footnotes, err)
		r.inlineImage, err = r.newTemplate("inline-image", tmpls.InlineImage, err)
		r.internalCrossReference, err = r.newTemplate("internal-xref", tmpls.InternalCrossReference, err)
		r.invalidFootnote, err = r.newTemplate("invalid-footnote", tmpls.InvalidFootnote, err)
		r.italicText, err = r.newTemplate("italic-text", tmpls.ItalicText, err)
		r.labeledList, err = r.newTemplate("labeled-list", tmpls.LabeledList, err)
		r.labeledListHorizontal, err = r.newTemplate("labeled-list-horizontal", tmpls.LabeledListHorizontal, err)
		r.lineBreak, err = r.newTemplate("line-break", tmpls.LineBreak, err)
		r.link, err = r.newTemplate("link", tmpls.Link, err)
		r.listingBlock, err = r.newTemplate("listing", tmpls.ListingBlock, err)
		r.literalBlock, err = r.newTemplate("literal-block", tmpls.LiteralBlock, err)
		r.manpageHeader, err = r.newTemplate("manpage-header", tmpls.ManpageHeader, err)
		r.manpageNameParagraph, err = r.newTemplate("manpage-name-paragraph", tmpls.ManpageNameParagraph, err)
		r.monospaceText, err = r.newTemplate("monospace-text", tmpls.MonospaceText, err)
		r.orderedList, err = r.newTemplate("ordered-list", tmpls.OrderedList, err)
		r.paragraph, err = r.newTemplate("paragraph", tmpls.Paragraph, err)
		r.passthroughBlock, err = r.newTemplate("passthrough", tmpls.PassthroughBlock, err)
		r.preamble, err = r.newTemplate("preamble", tmpls.Preamble, err)
		r.qAndAList, err = r.newTemplate("qanda-block", tmpls.QAndAList, err)
		r.quoteBlock, err = r.newTemplate("quote-block", tmpls.QuoteBlock, err)
		r.quoteParagraph, err = r.newTemplate("quote-paragraph", tmpls.QuoteParagraph, err)
		r.sectionContent, err = r.newTemplate("section-content", tmpls.SectionContent, err)
		r.sectionHeader, err = r.newTemplate("section-header", tmpls.SectionHeader, err)
		r.sectionOne, err = r.newTemplate("section-one", tmpls.SectionOne, err)
		r.stringElement, err = r.newTemplate("string-element", tmpls.StringElement, err)
		r.sidebarBlock, err = r.newTemplate("sidebar-block", tmpls.SidebarBlock, err)
		r.sourceBlock, err = r.newTemplate("source-block", tmpls.SourceBlock, err)
		r.sourceBlockContent, err = r.newTemplate("source-block-content", tmpls.SourceBlockContent, err)
		r.sourceParagraph, err = r.newTemplate("source-paragraph", tmpls.SourceParagraph, err)
		r.subscriptText, err = r.newTemplate("subscript", tmpls.SubscriptText, err)
		r.superscriptText, err = r.newTemplate("superscript", tmpls.SuperscriptText, err)
		r.table, err = r.newTemplate("table", tmpls.Table, err)
		r.tocRoot, err = r.newTemplate("toc-root", tmpls.TocRoot, err)
		r.tocSection, err = r.newTemplate("toc-section", tmpls.TocSection, err)
		r.unorderedList, err = r.newTemplate("unordered-list", tmpls.UnorderedList, err)
		r.verbatimLine, err = r.newTemplate("verbatim-line", tmpls.VerbatimLine, err)
		r.verseBlock, err = r.newTemplate("verse", tmpls.VerseBlock, err)
		r.verseBlockParagraph, err = r.newTemplate("verse-block-paragraph", tmpls.VerseBlockParagraph, err)
		r.verseParagraph, err = r.newTemplate("verse-paragraph", tmpls.VerseParagraph, err)
	})
	return err
}
