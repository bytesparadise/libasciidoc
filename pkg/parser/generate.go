package parser

//go:generate pigeon -optimize-parser -alternate-entrypoints AsciidocRawDocument,RawFile,TextDocument,DocumentRawBlock,FileLocation,IncludedFileLine,InlineLinks,LabeledListItemTerm,NormalBlockContent,NormalParagraphContent,VerseBlockContent,MarkdownQuoteBlockAttribution,InlineElements -o parser.go parser.peg
