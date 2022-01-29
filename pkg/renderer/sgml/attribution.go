package sgml

import "github.com/bytesparadise/libasciidoc/pkg/types"

// Attribution a document block attribution
type Attribution struct { // TODO: unexport this type?
	First  string
	Second string
}

func newAttribution(b types.WithAttributes) (Attribution, error) {
	result := Attribution{}
	if author, found, err := b.GetAttributes().GetAsString(types.AttrQuoteAuthor); err != nil {
		return Attribution{}, err
	} else if found {
		result.First = author
		if title, found, err := b.GetAttributes().GetAsString(types.AttrQuoteTitle); err != nil {
			return Attribution{}, err
		} else if found {
			result.Second = title
		}
	} else if title, found, err := b.GetAttributes().GetAsString(types.AttrQuoteTitle); err != nil {
		return Attribution{}, err
	} else if found {
		result.First = title
	}
	return result, nil
}
