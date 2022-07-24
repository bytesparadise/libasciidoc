package sgml

import "github.com/bytesparadise/libasciidoc/pkg/types"

// Attribution a document block attribution
type Attribution struct { // TODO: unexport this type?
	First  string
	Second string
}

func newAttribution(b types.WithAttributes) Attribution {
	result := Attribution{}
	if author, found := b.GetAttributes().GetAsString(types.AttrQuoteAuthor); found {
		result.First = author
		if title, found := b.GetAttributes().GetAsString(types.AttrQuoteTitle); found {
			result.Second = title
		}
	} else if title, found := b.GetAttributes().GetAsString(types.AttrQuoteTitle); found {
		result.First = title
	}
	return result
}
