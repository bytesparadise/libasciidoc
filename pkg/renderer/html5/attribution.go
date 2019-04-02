package html5

import "github.com/bytesparadise/libasciidoc/pkg/types"

// Attribution a document block attribution
type Attribution struct {
	First  string
	Second string
}

// NewParagraphAttribution return a new attribution for the given paragraph.
// Can be empty if no attribution was specified.
func NewParagraphAttribution(p types.Paragraph) Attribution {
	return newAttribution(p.Attributes)
}

// NewDelimitedBlockAttribution return a new attribution for the given delimited block.
// Can be empty if no attribution was specified.
func NewDelimitedBlockAttribution(b types.DelimitedBlock) Attribution {
	return newAttribution(b.Attributes)
}

func newAttribution(attrbs types.ElementAttributes) Attribution {
	result := Attribution{}
	if author := attrbs.GetAsString(types.AttrQuoteAuthor); author != "" {
		result.First = author
		if title := attrbs.GetAsString(types.AttrQuoteTitle); title != "" {
			result.Second = title
		}
	} else if title := attrbs.GetAsString(types.AttrQuoteTitle); title != "" {
		result.First = title
	}
	return result
}
