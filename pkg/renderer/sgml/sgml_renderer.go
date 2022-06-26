package sgml

import (
	"strings"
	"sync"
	texttemplate "text/template"

	"github.com/pkg/errors"
)

type sgmlRenderer struct {
	templates Templates
	functions texttemplate.FuncMap

	admonitionBlockOnce sync.Once
	admonitionBlockTmpl *texttemplate.Template

	admonitionParagraphOnce sync.Once
	admonitionParagraphTmpl *texttemplate.Template

	articleOnce sync.Once
	articleTmpl *texttemplate.Template

	articleHeaderOnce sync.Once
	articleHeaderTmpl *texttemplate.Template

	blockImageOnce sync.Once
	blockImageTmpl *texttemplate.Template

	boldTextOnce sync.Once
	boldTextTmpl *texttemplate.Template

	calloutListOnce sync.Once
	calloutListTmpl *texttemplate.Template

	calloutListElementOnce sync.Once
	calloutListElementTmpl *texttemplate.Template

	calloutRefOnce sync.Once
	calloutRefTmpl *texttemplate.Template

	embeddedParagraphOnce sync.Once
	embeddedParagraphTmpl *texttemplate.Template

	documentDetailsOnce sync.Once
	documentDetailsTmpl *texttemplate.Template

	documentAuthorDetailsOnce sync.Once
	documentAuthorDetailsTmpl *texttemplate.Template

	exampleBlockOnce sync.Once
	exampleBlockTmpl *texttemplate.Template

	externalCrossReferenceOnce sync.Once
	externalCrossReferenceTmpl *texttemplate.Template

	fencedBlockOnce sync.Once
	fencedBlockTmpl *texttemplate.Template

	footnoteOnce sync.Once
	footnoteTmpl *texttemplate.Template

	footnoteElementOnce sync.Once
	footnoteElementTmpl *texttemplate.Template

	footnoteRefOnce sync.Once
	footnoteRefTmpl *texttemplate.Template

	footnoteRefPlainOnce sync.Once
	footnoteRefPlainTmpl *texttemplate.Template

	footnotesOnce sync.Once
	footnotesTmpl *texttemplate.Template

	iconFontOnce sync.Once
	iconFontTmpl *texttemplate.Template

	iconImageOnce sync.Once
	iconImageTmpl *texttemplate.Template

	iconTextOnce sync.Once
	iconTextTmpl *texttemplate.Template

	inlineButtonOnce sync.Once
	inlineButtonTmpl *texttemplate.Template

	inlineIconOnce sync.Once
	inlineIconTmpl *texttemplate.Template

	inlineImageOnce sync.Once
	inlineImageTmpl *texttemplate.Template

	inlineMenuOnce sync.Once
	inlineMenuTmpl *texttemplate.Template

	internalCrossReferenceOnce sync.Once
	internalCrossReferenceTmpl *texttemplate.Template

	invalidFootnoteOnce sync.Once
	invalidFootnoteTmpl *texttemplate.Template

	italicTextOnce sync.Once
	italicTextTmpl *texttemplate.Template

	labeledListOnce sync.Once
	labeledListTmpl *texttemplate.Template

	labeledListElementOnce sync.Once
	labeledListElementTmpl *texttemplate.Template

	labeledListHorizontalOnce sync.Once
	labeledListHorizontalTmpl *texttemplate.Template

	labeledListHorizontalElementOnce sync.Once
	labeledListHorizontalElementTmpl *texttemplate.Template

	lineBreakOnce sync.Once
	lineBreakTmpl *texttemplate.Template

	linkOnce sync.Once
	linkTmpl *texttemplate.Template

	listingBlockOnce sync.Once
	listingBlockTmpl *texttemplate.Template

	literalBlockOnce sync.Once
	literalBlockTmpl *texttemplate.Template

	manpageHeaderOnce sync.Once
	manpageHeaderTmpl *texttemplate.Template

	manpageNameParagraphOnce sync.Once
	manpageNameParagraphTmpl *texttemplate.Template

	markdownQuoteBlockOnce sync.Once
	markdownQuoteBlockTmpl *texttemplate.Template

	markedTextOnce sync.Once
	markedTextTmpl *texttemplate.Template

	monospaceTextOnce sync.Once
	monospaceTextTmpl *texttemplate.Template

	openBlockOnce sync.Once
	openBlockTmpl *texttemplate.Template

	orderedListOnce sync.Once
	orderedListTmpl *texttemplate.Template

	orderedListElementOnce sync.Once
	orderedListElementTmpl *texttemplate.Template

	paragraphOnce sync.Once
	paragraphTmpl *texttemplate.Template

	passthroughBlockOnce sync.Once
	passthroughBlockTmpl *texttemplate.Template

	preambleOnce sync.Once
	preambleTmpl *texttemplate.Template

	qAndAListOnce sync.Once
	qAndAListTmpl *texttemplate.Template

	qAndAListElementOnce sync.Once
	qAndAListElementTmpl *texttemplate.Template

	quoteBlockOnce sync.Once
	quoteBlockTmpl *texttemplate.Template

	quoteParagraphOnce sync.Once
	quoteParagraphTmpl *texttemplate.Template

	sectionContentOnce sync.Once
	sectionContentTmpl *texttemplate.Template

	sectionTitleOnce sync.Once
	sectionTitleTmpl *texttemplate.Template

	sidebarBlockOnce sync.Once
	sidebarBlockTmpl *texttemplate.Template

	sourceBlockOnce sync.Once
	sourceBlockTmpl *texttemplate.Template

	subscriptTextOnce sync.Once
	subscriptTextTmpl *texttemplate.Template

	superscriptTextOnce sync.Once
	superscriptTextTmpl *texttemplate.Template

	tableOnce sync.Once
	tableTmpl *texttemplate.Template

	tableBodyOnce sync.Once
	tableBodyTmpl *texttemplate.Template

	tableCellOnce sync.Once
	tableCellTmpl *texttemplate.Template

	tableCellBlockOnce sync.Once
	tableCellBlockTmpl *texttemplate.Template

	tableHeaderOnce sync.Once
	tableHeaderTmpl *texttemplate.Template

	tableHeaderCellOnce sync.Once
	tableHeaderCellTmpl *texttemplate.Template

	tableFooterOnce sync.Once
	tableFooterTmpl *texttemplate.Template

	tableFooterCellOnce sync.Once
	tableFooterCellTmpl *texttemplate.Template

	tableRowOnce sync.Once
	tableRowTmpl *texttemplate.Template

	thematicBreakOnce sync.Once
	thematicBreakTmpl *texttemplate.Template

	tocEntryOnce sync.Once
	tocEntryTmpl *texttemplate.Template

	tocRootOnce sync.Once
	tocRootTmpl *texttemplate.Template

	tocSectionOnce sync.Once
	tocSectionTmpl *texttemplate.Template

	unorderedListOnce sync.Once
	unorderedListTmpl *texttemplate.Template

	unorderedListElementOnce sync.Once
	unorderedListElementTmpl *texttemplate.Template

	verseBlockOnce sync.Once
	verseBlockTmpl *texttemplate.Template

	verseParagraphOnce sync.Once
	verseParagraphTmpl *texttemplate.Template
}

type template func() (*texttemplate.Template, error)

func (s *sgmlRenderer) execute(loadTmpl template, data interface{}) (string, error) {
	tmpl, err := loadTmpl()
	result := &strings.Builder{}
	if err != nil {
		return "", errors.Wrap(err, "unable to load template")
	}
	if err := tmpl.Execute(result, data); err != nil {
		return "", err
	}
	return result.String(), nil

}

func (r *sgmlRenderer) admonitionBlock() (*texttemplate.Template, error) {
	var err error
	r.admonitionBlockOnce.Do(func() {
		r.admonitionBlockTmpl, err = r.newTemplate("AdmonitionBlock", r.templates.AdmonitionBlock, err)
	})
	return r.admonitionBlockTmpl, err
}

func (r *sgmlRenderer) admonitionParagraph() (*texttemplate.Template, error) {
	var err error
	r.admonitionParagraphOnce.Do(func() {
		r.admonitionParagraphTmpl, err = r.newTemplate("AdmonitionParagraph", r.templates.AdmonitionParagraph, err)
	})
	return r.admonitionParagraphTmpl, err
}

func (r *sgmlRenderer) article() (*texttemplate.Template, error) {
	var err error
	r.articleOnce.Do(func() {
		r.articleTmpl, err = r.newTemplate("Article", r.templates.Article, err)
	})
	return r.articleTmpl, err
}

func (r *sgmlRenderer) articleHeader() (*texttemplate.Template, error) {
	var err error
	r.articleHeaderOnce.Do(func() {
		r.articleHeaderTmpl, err = r.newTemplate("ArticleHeader", r.templates.ArticleHeader, err)
	})
	return r.articleHeaderTmpl, err
}

func (r *sgmlRenderer) blockImage() (*texttemplate.Template, error) {
	var err error
	r.blockImageOnce.Do(func() {
		r.blockImageTmpl, err = r.newTemplate("BlockImage", r.templates.BlockImage, err)
	})
	return r.blockImageTmpl, err
}

func (r *sgmlRenderer) boldText() (*texttemplate.Template, error) {
	var err error
	r.boldTextOnce.Do(func() {
		r.boldTextTmpl, err = r.newTemplate("BoldText", r.templates.BoldText, err)
	})
	return r.boldTextTmpl, err
}

func (r *sgmlRenderer) calloutList() (*texttemplate.Template, error) {
	var err error
	r.calloutListOnce.Do(func() {
		r.calloutListTmpl, err = r.newTemplate("CalloutList", r.templates.CalloutList, err)
	})
	return r.calloutListTmpl, err
}

func (r *sgmlRenderer) calloutListElement() (*texttemplate.Template, error) {
	var err error
	r.calloutListElementOnce.Do(func() {
		r.calloutListElementTmpl, err = r.newTemplate("CalloutListElement", r.templates.CalloutListElement, err)
	})
	return r.calloutListElementTmpl, err
}

func (r *sgmlRenderer) calloutRef() (*texttemplate.Template, error) {
	var err error
	r.calloutRefOnce.Do(func() {
		r.calloutRefTmpl, err = r.newTemplate("CalloutRef", r.templates.CalloutRef, err)
	})
	return r.calloutRefTmpl, err
}

func (r *sgmlRenderer) embeddedParagraph() (*texttemplate.Template, error) {
	var err error
	r.embeddedParagraphOnce.Do(func() {
		r.embeddedParagraphTmpl, err = r.newTemplate("EmbeddedParagraph", r.templates.EmbeddedParagraph, err)
	})
	return r.embeddedParagraphTmpl, err
}

func (r *sgmlRenderer) documentDetails() (*texttemplate.Template, error) {
	var err error
	r.documentDetailsOnce.Do(func() {
		r.documentDetailsTmpl, err = r.newTemplate("DocumentDetails", r.templates.DocumentDetails, err)
	})
	return r.documentDetailsTmpl, err
}

func (r *sgmlRenderer) documentAuthorDetails() (*texttemplate.Template, error) {
	var err error
	r.documentAuthorDetailsOnce.Do(func() {
		r.documentAuthorDetailsTmpl, err = r.newTemplate("DocumentAuthorDetails", r.templates.DocumentAuthorDetails, err)
	})
	return r.documentAuthorDetailsTmpl, err
}

func (r *sgmlRenderer) exampleBlock() (*texttemplate.Template, error) {
	var err error
	r.exampleBlockOnce.Do(func() {
		r.exampleBlockTmpl, err = r.newTemplate("ExampleBlock", r.templates.ExampleBlock, err)
	})
	return r.exampleBlockTmpl, err
}

func (r *sgmlRenderer) externalCrossReference() (*texttemplate.Template, error) {
	var err error
	r.externalCrossReferenceOnce.Do(func() {
		r.externalCrossReferenceTmpl, err = r.newTemplate("ExternalCrossReference", r.templates.ExternalCrossReference, err)
	})
	return r.externalCrossReferenceTmpl, err
}

func (r *sgmlRenderer) fencedBlock() (*texttemplate.Template, error) {
	var err error
	r.fencedBlockOnce.Do(func() {
		r.fencedBlockTmpl, err = r.newTemplate("FencedBlock", r.templates.FencedBlock, err)
	})
	return r.fencedBlockTmpl, err
}

func (r *sgmlRenderer) footnote() (*texttemplate.Template, error) {
	var err error
	r.footnoteOnce.Do(func() {
		r.footnoteTmpl, err = r.newTemplate("Footnote", r.templates.Footnote, err)
	})
	return r.footnoteTmpl, err
}

func (r *sgmlRenderer) footnoteElement() (*texttemplate.Template, error) {
	var err error
	r.footnoteElementOnce.Do(func() {
		r.footnoteElementTmpl, err = r.newTemplate("FootnoteElement", r.templates.FootnoteElement, err)
	})
	return r.footnoteElementTmpl, err
}

func (r *sgmlRenderer) footnoteRef() (*texttemplate.Template, error) {
	var err error
	r.footnoteRefOnce.Do(func() {
		r.footnoteRefTmpl, err = r.newTemplate("FootnoteRef", r.templates.FootnoteRef, err)
	})
	return r.footnoteRefTmpl, err
}

func (r *sgmlRenderer) footnoteRefPlain() (*texttemplate.Template, error) {
	var err error
	r.footnoteRefPlainOnce.Do(func() {
		r.footnoteRefPlainTmpl, err = r.newTemplate("FootnoteRefPlain", r.templates.FootnoteRefPlain, err)
	})
	return r.footnoteRefPlainTmpl, err
}

func (r *sgmlRenderer) footnotes() (*texttemplate.Template, error) {
	var err error
	r.footnotesOnce.Do(func() {
		r.footnotesTmpl, err = r.newTemplate("Footnotes", r.templates.Footnotes, err)
	})
	return r.footnotesTmpl, err
}

func (r *sgmlRenderer) iconFont() (*texttemplate.Template, error) {
	var err error
	r.iconFontOnce.Do(func() {
		r.iconFontTmpl, err = r.newTemplate("IconFont", r.templates.IconFont, err)
	})
	return r.iconFontTmpl, err
}

func (r *sgmlRenderer) iconImage() (*texttemplate.Template, error) {
	var err error
	r.iconImageOnce.Do(func() {
		r.iconImageTmpl, err = r.newTemplate("IconImage", r.templates.IconImage, err)
	})
	return r.iconImageTmpl, err
}

func (r *sgmlRenderer) iconText() (*texttemplate.Template, error) {
	var err error
	r.iconTextOnce.Do(func() {
		r.iconTextTmpl, err = r.newTemplate("IconText", r.templates.IconText, err)
	})
	return r.iconTextTmpl, err
}

func (r *sgmlRenderer) inlineButton() (*texttemplate.Template, error) {
	var err error
	r.inlineButtonOnce.Do(func() {
		r.inlineButtonTmpl, err = r.newTemplate("InlineButton", r.templates.InlineButton, err)
	})
	return r.inlineButtonTmpl, err
}

func (r *sgmlRenderer) inlineIcon() (*texttemplate.Template, error) {
	var err error
	r.inlineIconOnce.Do(func() {
		r.inlineIconTmpl, err = r.newTemplate("InlineIcon", r.templates.InlineIcon, err)
	})
	return r.inlineIconTmpl, err
}

func (r *sgmlRenderer) inlineImage() (*texttemplate.Template, error) {
	var err error
	r.inlineImageOnce.Do(func() {
		r.inlineImageTmpl, err = r.newTemplate("InlineImage", r.templates.InlineImage, err)
	})
	return r.inlineImageTmpl, err
}

func (r *sgmlRenderer) inlineMenu() (*texttemplate.Template, error) {
	var err error
	r.inlineMenuOnce.Do(func() {
		r.inlineMenuTmpl, err = r.newTemplate("InlineMenu", r.templates.InlineMenu, err)
	})
	return r.inlineMenuTmpl, err
}

func (r *sgmlRenderer) internalCrossReference() (*texttemplate.Template, error) {
	var err error
	r.internalCrossReferenceOnce.Do(func() {
		r.internalCrossReferenceTmpl, err = r.newTemplate("InternalCrossReference", r.templates.InternalCrossReference, err)
	})
	return r.internalCrossReferenceTmpl, err
}

func (r *sgmlRenderer) invalidFootnote() (*texttemplate.Template, error) {
	var err error
	r.invalidFootnoteOnce.Do(func() {
		r.invalidFootnoteTmpl, err = r.newTemplate("InvalidFootnote", r.templates.InvalidFootnote, err)
	})
	return r.invalidFootnoteTmpl, err
}

func (r *sgmlRenderer) italicText() (*texttemplate.Template, error) {
	var err error
	r.italicTextOnce.Do(func() {
		r.italicTextTmpl, err = r.newTemplate("ItalicText", r.templates.ItalicText, err)
	})
	return r.italicTextTmpl, err
}

func (r *sgmlRenderer) labeledList() (*texttemplate.Template, error) {
	var err error
	r.labeledListOnce.Do(func() {
		r.labeledListTmpl, err = r.newTemplate("LabeledList", r.templates.LabeledList, err)
	})
	return r.labeledListTmpl, err
}

func (r *sgmlRenderer) labeledListElement() (*texttemplate.Template, error) {
	var err error
	r.labeledListElementOnce.Do(func() {
		r.labeledListElementTmpl, err = r.newTemplate("LabeledListElement", r.templates.LabeledListElement, err)
	})
	return r.labeledListElementTmpl, err
}

func (r *sgmlRenderer) labeledListHorizontal() (*texttemplate.Template, error) {
	var err error
	r.labeledListHorizontalOnce.Do(func() {
		r.labeledListHorizontalTmpl, err = r.newTemplate("LabeledListHorizontal", r.templates.LabeledListHorizontal, err)
	})
	return r.labeledListHorizontalTmpl, err
}

func (r *sgmlRenderer) labeledListHorizontalElement() (*texttemplate.Template, error) {
	var err error
	r.labeledListHorizontalElementOnce.Do(func() {
		r.labeledListHorizontalElementTmpl, err = r.newTemplate("LabeledListHorizontalElement", r.templates.LabeledListHorizontalElement, err)
	})
	return r.labeledListHorizontalElementTmpl, err
}

func (r *sgmlRenderer) lineBreak() (*texttemplate.Template, error) {
	var err error
	r.lineBreakOnce.Do(func() {
		r.lineBreakTmpl, err = r.newTemplate("LineBreak", r.templates.LineBreak, err)
	})
	return r.lineBreakTmpl, err
}

func (r *sgmlRenderer) link() (*texttemplate.Template, error) {
	var err error
	r.linkOnce.Do(func() {
		r.linkTmpl, err = r.newTemplate("Link", r.templates.Link, err)
	})
	return r.linkTmpl, err
}

func (r *sgmlRenderer) listingBlock() (*texttemplate.Template, error) {
	var err error
	r.listingBlockOnce.Do(func() {
		r.listingBlockTmpl, err = r.newTemplate("ListingBlock", r.templates.ListingBlock, err)
	})
	return r.listingBlockTmpl, err
}

func (r *sgmlRenderer) literalBlock() (*texttemplate.Template, error) {
	var err error
	r.literalBlockOnce.Do(func() {
		r.literalBlockTmpl, err = r.newTemplate("LiteralBlock", r.templates.LiteralBlock, err)
	})
	return r.literalBlockTmpl, err
}

func (r *sgmlRenderer) manpageHeader() (*texttemplate.Template, error) {
	var err error
	r.manpageHeaderOnce.Do(func() {
		r.manpageHeaderTmpl, err = r.newTemplate("ManpageHeader", r.templates.ManpageHeader, err)
	})
	return r.manpageHeaderTmpl, err
}

func (r *sgmlRenderer) manpageNameParagraph() (*texttemplate.Template, error) {
	var err error
	r.manpageNameParagraphOnce.Do(func() {
		r.manpageNameParagraphTmpl, err = r.newTemplate("ManpageNameParagraph", r.templates.ManpageNameParagraph, err)
	})
	return r.manpageNameParagraphTmpl, err
}

func (r *sgmlRenderer) markdownQuoteBlock() (*texttemplate.Template, error) {
	var err error
	r.markdownQuoteBlockOnce.Do(func() {
		r.markdownQuoteBlockTmpl, err = r.newTemplate("MarkdownQuoteBlock", r.templates.MarkdownQuoteBlock, err)
	})
	return r.markdownQuoteBlockTmpl, err
}

func (r *sgmlRenderer) markedText() (*texttemplate.Template, error) {
	var err error
	r.markedTextOnce.Do(func() {
		r.markedTextTmpl, err = r.newTemplate("MarkedText", r.templates.MarkedText, err)
	})
	return r.markedTextTmpl, err
}

func (r *sgmlRenderer) monospaceText() (*texttemplate.Template, error) {
	var err error
	r.monospaceTextOnce.Do(func() {
		r.monospaceTextTmpl, err = r.newTemplate("MonospaceText", r.templates.MonospaceText, err)
	})
	return r.monospaceTextTmpl, err
}

func (r *sgmlRenderer) openBlock() (*texttemplate.Template, error) {
	var err error
	r.openBlockOnce.Do(func() {
		r.openBlockTmpl, err = r.newTemplate("OpenBlock", r.templates.OpenBlock, err)
	})
	return r.openBlockTmpl, err
}

func (r *sgmlRenderer) orderedList() (*texttemplate.Template, error) {
	var err error
	r.orderedListOnce.Do(func() {
		r.orderedListTmpl, err = r.newTemplate("OrderedList", r.templates.OrderedList, err)
	})
	return r.orderedListTmpl, err
}

func (r *sgmlRenderer) orderedListElement() (*texttemplate.Template, error) {
	var err error
	r.orderedListElementOnce.Do(func() {
		r.orderedListElementTmpl, err = r.newTemplate("OrderedListElement", r.templates.OrderedListElement, err)
	})
	return r.orderedListElementTmpl, err
}

func (r *sgmlRenderer) paragraph() (*texttemplate.Template, error) {
	var err error
	r.paragraphOnce.Do(func() {
		r.paragraphTmpl, err = r.newTemplate("Paragraph", r.templates.Paragraph, err)
	})
	return r.paragraphTmpl, err
}

func (r *sgmlRenderer) passthroughBlock() (*texttemplate.Template, error) {
	var err error
	r.passthroughBlockOnce.Do(func() {
		r.passthroughBlockTmpl, err = r.newTemplate("PassthroughBlock", r.templates.PassthroughBlock, err)
	})
	return r.passthroughBlockTmpl, err
}

func (r *sgmlRenderer) preamble() (*texttemplate.Template, error) {
	var err error
	r.preambleOnce.Do(func() {
		r.preambleTmpl, err = r.newTemplate("Preamble", r.templates.Preamble, err)
	})
	return r.preambleTmpl, err
}

func (r *sgmlRenderer) qAndAList() (*texttemplate.Template, error) {
	var err error
	r.qAndAListOnce.Do(func() {
		r.qAndAListTmpl, err = r.newTemplate("QAndAList", r.templates.QAndAList, err)
	})
	return r.qAndAListTmpl, err
}

func (r *sgmlRenderer) qAndAListElement() (*texttemplate.Template, error) {
	var err error
	r.qAndAListElementOnce.Do(func() {
		r.qAndAListElementTmpl, err = r.newTemplate("QAndAListElement", r.templates.QAndAListElement, err)
	})
	return r.qAndAListElementTmpl, err
}

func (r *sgmlRenderer) quoteBlock() (*texttemplate.Template, error) {
	var err error
	r.quoteBlockOnce.Do(func() {
		r.quoteBlockTmpl, err = r.newTemplate("QuoteBlock", r.templates.QuoteBlock, err)
	})
	return r.quoteBlockTmpl, err
}

func (r *sgmlRenderer) quoteParagraph() (*texttemplate.Template, error) {
	var err error
	r.quoteParagraphOnce.Do(func() {
		r.quoteParagraphTmpl, err = r.newTemplate("QuoteParagraph", r.templates.QuoteParagraph, err)
	})
	return r.quoteParagraphTmpl, err
}

func (r *sgmlRenderer) sectionContent() (*texttemplate.Template, error) {
	var err error
	r.sectionContentOnce.Do(func() {
		r.sectionContentTmpl, err = r.newTemplate("SectionContent", r.templates.SectionContent, err)
	})
	return r.sectionContentTmpl, err
}

func (r *sgmlRenderer) sectionTitle() (*texttemplate.Template, error) {
	var err error
	r.sectionTitleOnce.Do(func() {
		r.sectionTitleTmpl, err = r.newTemplate("SectionTitle", r.templates.SectionTitle, err)
	})
	return r.sectionTitleTmpl, err
}

func (r *sgmlRenderer) sidebarBlock() (*texttemplate.Template, error) {
	var err error
	r.sidebarBlockOnce.Do(func() {
		r.sidebarBlockTmpl, err = r.newTemplate("SidebarBlock", r.templates.SidebarBlock, err)
	})
	return r.sidebarBlockTmpl, err
}

func (r *sgmlRenderer) sourceBlock() (*texttemplate.Template, error) {
	var err error
	r.sourceBlockOnce.Do(func() {
		r.sourceBlockTmpl, err = r.newTemplate("SourceBlock", r.templates.SourceBlock, err)
	})
	return r.sourceBlockTmpl, err
}

func (r *sgmlRenderer) subscriptText() (*texttemplate.Template, error) {
	var err error
	r.subscriptTextOnce.Do(func() {
		r.subscriptTextTmpl, err = r.newTemplate("SubscriptText", r.templates.SubscriptText, err)
	})
	return r.subscriptTextTmpl, err
}

func (r *sgmlRenderer) superscriptText() (*texttemplate.Template, error) {
	var err error
	r.superscriptTextOnce.Do(func() {
		r.superscriptTextTmpl, err = r.newTemplate("SuperscriptText", r.templates.SuperscriptText, err)
	})
	return r.superscriptTextTmpl, err
}

func (r *sgmlRenderer) table() (*texttemplate.Template, error) {
	var err error
	r.tableOnce.Do(func() {
		r.tableTmpl, err = r.newTemplate("Table", r.templates.Table, err)
	})
	return r.tableTmpl, err
}

func (r *sgmlRenderer) tableBody() (*texttemplate.Template, error) {
	var err error
	r.tableBodyOnce.Do(func() {
		r.tableBodyTmpl, err = r.newTemplate("TableBody", r.templates.TableBody, err)
	})
	return r.tableBodyTmpl, err
}

func (r *sgmlRenderer) tableCell() (*texttemplate.Template, error) {
	var err error
	r.tableCellOnce.Do(func() {
		r.tableCellTmpl, err = r.newTemplate("TableCell", r.templates.TableCell, err)
	})
	return r.tableCellTmpl, err
}

func (r *sgmlRenderer) tableCellBlock() (*texttemplate.Template, error) {
	var err error
	r.tableCellBlockOnce.Do(func() {
		r.tableCellBlockTmpl, err = r.newTemplate("TableCellBlock", r.templates.TableCellBlock, err)
	})
	return r.tableCellBlockTmpl, err
}

func (r *sgmlRenderer) tableHeader() (*texttemplate.Template, error) {
	var err error
	r.tableHeaderOnce.Do(func() {
		r.tableHeaderTmpl, err = r.newTemplate("TableHeader", r.templates.TableHeader, err)
	})
	return r.tableHeaderTmpl, err
}

func (r *sgmlRenderer) tableHeaderCell() (*texttemplate.Template, error) {
	var err error
	r.tableHeaderCellOnce.Do(func() {
		r.tableHeaderCellTmpl, err = r.newTemplate("TableHeaderCell", r.templates.TableHeaderCell, err)
	})
	return r.tableHeaderCellTmpl, err
}

func (r *sgmlRenderer) tableFooter() (*texttemplate.Template, error) {
	var err error
	r.tableFooterOnce.Do(func() {
		r.tableFooterTmpl, err = r.newTemplate("TableFooter", r.templates.TableFooter, err)
	})
	return r.tableFooterTmpl, err
}

func (r *sgmlRenderer) tableFooterCell() (*texttemplate.Template, error) {
	var err error
	r.tableFooterCellOnce.Do(func() {
		r.tableFooterCellTmpl, err = r.newTemplate("TableFooterCell", r.templates.TableFooterCell, err)
	})
	return r.tableFooterCellTmpl, err
}

func (r *sgmlRenderer) tableRow() (*texttemplate.Template, error) {
	var err error
	r.tableRowOnce.Do(func() {
		r.tableRowTmpl, err = r.newTemplate("TableRow", r.templates.TableRow, err)
	})
	return r.tableRowTmpl, err
}

func (r *sgmlRenderer) thematicBreak() (*texttemplate.Template, error) {
	var err error
	r.thematicBreakOnce.Do(func() {
		r.thematicBreakTmpl, err = r.newTemplate("ThematicBreak", r.templates.ThematicBreak, err)
	})
	return r.thematicBreakTmpl, err
}

func (r *sgmlRenderer) tocEntry() (*texttemplate.Template, error) {
	var err error
	r.tocEntryOnce.Do(func() {
		r.tocEntryTmpl, err = r.newTemplate("TocEntry", r.templates.TocEntry, err)
	})
	return r.tocEntryTmpl, err
}

func (r *sgmlRenderer) tocRoot() (*texttemplate.Template, error) {
	var err error
	r.tocRootOnce.Do(func() {
		r.tocRootTmpl, err = r.newTemplate("TocRoot", r.templates.TocRoot, err)
	})
	return r.tocRootTmpl, err
}

func (r *sgmlRenderer) tocSection() (*texttemplate.Template, error) {
	var err error
	r.tocSectionOnce.Do(func() {
		r.tocSectionTmpl, err = r.newTemplate("TocSection", r.templates.TocSection, err)
	})
	return r.tocSectionTmpl, err
}

func (r *sgmlRenderer) unorderedList() (*texttemplate.Template, error) {
	var err error
	r.unorderedListOnce.Do(func() {
		r.unorderedListTmpl, err = r.newTemplate("UnorderedList", r.templates.UnorderedList, err)
	})
	return r.unorderedListTmpl, err
}

func (r *sgmlRenderer) unorderedListElement() (*texttemplate.Template, error) {
	var err error
	r.unorderedListElementOnce.Do(func() {
		r.unorderedListElementTmpl, err = r.newTemplate("UnorderedListElement", r.templates.UnorderedListElement, err)
	})
	return r.unorderedListElementTmpl, err
}

func (r *sgmlRenderer) verseBlock() (*texttemplate.Template, error) {
	var err error
	r.verseBlockOnce.Do(func() {
		r.verseBlockTmpl, err = r.newTemplate("VerseBlock", r.templates.VerseBlock, err)
	})
	return r.verseBlockTmpl, err
}

func (r *sgmlRenderer) verseParagraph() (*texttemplate.Template, error) {
	var err error
	r.verseParagraphOnce.Do(func() {
		r.verseParagraphTmpl, err = r.newTemplate("VerseParagraph", r.templates.VerseParagraph, err)
	})
	return r.verseParagraphTmpl, err
}
