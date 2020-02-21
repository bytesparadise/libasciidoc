package testsupport

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// IncludeTableOfContentsPlaceHolder includes the table or contents placeholder in the given document
func IncludeTableOfContentsPlaceHolder(actual types.Document) types.Document {
	ctx := renderer.NewContext(actual)
	ctx = renderer.IncludeTableOfContentsPlaceHolder(ctx)
	return ctx.Document
}
