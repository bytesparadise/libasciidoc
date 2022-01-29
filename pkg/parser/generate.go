package parser

//go:generate pigeon -optimize-parser -optimize-grammar -alternate-entrypoints DocumentRawLine,DocumentFragment,Substitutions,DelimitedBlockElements,HeaderGroup,FileLocation,IncludedFileLine,MarkdownQuoteAttribution,BlockAttributes,InlineAttributes,TableColumnsAttribute,LineRanges,TagRanges -o parser.go parser.peg
