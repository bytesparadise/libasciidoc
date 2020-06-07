package sgml

import "github.com/bytesparadise/libasciidoc/pkg/types"

// Attribution a document block attribution
type Attribution struct { // TODO: unexport this type?
	First  string
	Second string
}

// newParagraphAttribution return a new attribution for the given paragraph.
// Can be empty if no attribution was specified.
func newParagraphAttribution(p types.Paragraph) Attribution {
	return newAttribution(p.Attributes)
}

// newDelimitedBlockAttribution return a new attribution for the given delimited block.
// Can be empty if no attribution was specified.
func newDelimitedBlockAttribution(b types.DelimitedBlock) Attribution {
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
