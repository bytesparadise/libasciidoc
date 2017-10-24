package html5

import (
	"bytes"
	"fmt"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
)

func processAttributeDeclaration(ctx *renderer.Context, attribute *types.DocumentAttributeDeclaration) error {
	ctx.Document.Attributes.AddAttribute(attribute)
	return nil
}

func processAttributeReset(ctx *renderer.Context, attribute types.DocumentAttributeReset) error {
	ctx.Document.Attributes.Reset(attribute)
	return nil
}

func renderAttributeSubstitution(ctx *renderer.Context, attribute types.DocumentAttributeSubstitution) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	if value, found := ctx.Document.Attributes[attribute.Name]; found {
		result.WriteString(fmt.Sprintf("%v", value))
	} else {
		result.WriteString(fmt.Sprintf("{%s}", attribute.Name))
	}
	return result.Bytes(), nil
}
