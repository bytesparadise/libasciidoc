package sgml

import "github.com/bytesparadise/libasciidoc/pkg/types"

// Attribution a document block attribution
type Attribution struct { // TODO: unexport this type?
	First  string
	Second string
}

// paragraphAttribution return a new attribution for the given Paragraph.
// Can be empty if no attribution was specified.
func paragraphAttribution(p types.Paragraph) (Attribution, error) {
	return newAttribution(p.Attributes)
}

// quoteBlockAttribution return a new attribution for the given QuoteBlock.
// Can be empty if no attribution was specified.
func quoteBlockAttribution(b types.QuoteBlock) (Attribution, error) {
	return newAttribution(b.Attributes)
}

// verseBlockAttribution return a new attribution for the given VerseBlock.
// Can be empty if no attribution was specified.
func verseBlockAttribution(b types.VerseBlock) (Attribution, error) {
	return newAttribution(b.Attributes)
}

// markdownQuoteBlockAttribution return a new attribution for the given MarkdownQuoteBlock.
// Can be empty if no attribution was specified.
func markdownQuoteBlockAttribution(b types.MarkdownQuoteBlock) (Attribution, error) {
	return newAttribution(b.Attributes)
}

func newAttribution(attrs types.Attributes) (Attribution, error) {
	result := Attribution{}
	if author, found, err := attrs.GetAsString(types.AttrQuoteAuthor); err != nil {
		return Attribution{}, err
	} else if found {
		result.First = author
		if title, found, err := attrs.GetAsString(types.AttrQuoteTitle); err != nil {
			return Attribution{}, err
		} else if found {
			result.Second = title
		}
	} else if title, found, err := attrs.GetAsString(types.AttrQuoteTitle); err != nil {
		return Attribution{}, err
	} else if found {
		result.First = title
	}
	return result, nil
}
