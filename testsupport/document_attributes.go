package testsupport

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// DocumentAttributes returns the attributes in the document header
func DocumentAttributes(actual types.Document) types.DocumentAttributes {
	ctx := renderer.NewContext(actual)
	ctx = renderer.ProcessDocumentHeader(ctx)
	return ctx.Document.Attributes
}
