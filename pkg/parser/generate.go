package parser

//go:generate pigeon -optimize-parser -optimize-grammar -alternate-entrypoints RawSource,RawDocument,DocumentBlock,FileLocation,LabeledListItemTerm,SpecialCharacterSubs,MarkdownQuoteAttribution,QuotedTextSubs,NoneSubs,AttributeSubs,ReplacementSubs,PostReplacementSubs,InlinePassthroughSubs,CalloutSubs,InlineMacroSubs,MarkdownQuoteMacroSubs,BlockAttributes,LineRanges,TagRanges -o parser.go parser.peg
