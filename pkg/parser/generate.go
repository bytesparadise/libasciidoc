package parser

//go:generate pigeon -optimize-parser -alternate-entrypoints RawSource,RawDocument,DocumentRawBlock,FileLocation,IncludedFileLine,InlineLinks,LabeledListItemTerm,MarkdownQuoteAttribution,QuotedTextSubs,NoneSubs,AttributeSubs,ReplacementSubs,PostReplacementSubs,InlinePassthroughSubs,CalloutSubs,RawDocumentBlocks,InlineMacroSubs,NormalBlocks,VerseMacroSubs,MarkdownQuoteMacroSubs,MarkdownQuoteLine -o parser.go parser.peg
