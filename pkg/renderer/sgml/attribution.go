package sgml

import "github.com/bytesparadise/libasciidoc/pkg/types"

// Attribution a document block attribution
type Attribution struct { // TODO: unexport this type?
	First  string
	Second string
}

// paragraphAttribution return a new attribution for the given Paragraph.
// Can be empty if no attribution was specified.
func paragraphAttribution(p types.Paragraph) Attribution {
	return newAttribution(p.Attributes)
}

// quoteBlockAttribution return a new attribution for the given QuoteBlock.
// Can be empty if no attribution was specified.
func quoteBlockAttribution(b types.QuoteBlock) Attribution {
	return newAttribution(b.Attributes)
}

// verseBlockAttribution return a new attribution for the given VerseBlock.
// Can be empty if no attribution was specified.
func verseBlockAttribution(b types.VerseBlock) Attribution {
	return newAttribution(b.Attributes)
}

// markdownQuoteBlockAttribution return a new attribution for the given MarkdownQuoteBlock.
// Can be empty if no attribution was specified.
func markdownQuoteBlockAttribution(b types.MarkdownQuoteBlock) Attribution {
	return newAttribution(b.Attributes)
}

func newAttribution(attrs types.Attributes) Attribution {
	result := Attribution{}
	if author, found := attrs.GetAsString(types.AttrQuoteAuthor); found {
		result.First = author
		if title, found := attrs.GetAsString(types.AttrQuoteTitle); found {
			result.Second = title
		}
	} else if title, found := attrs.GetAsString(types.AttrQuoteTitle); found {
		result.First = title
	}
	return result
}
