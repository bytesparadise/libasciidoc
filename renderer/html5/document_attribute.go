package html5

import (
	"bytes"
	"fmt"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
)

func processAttributeDeclaration(ctx *renderer.Context, attr types.DocumentAttributeDeclaration) error {
	ctx.Document.Attributes.AddAttribute(attr)
	return nil
}

func processAttributeReset(ctx *renderer.Context, attr types.DocumentAttributeReset) error {
	ctx.Document.Attributes.Reset(attr)
	return nil
}

func renderAttributeSubstitution(ctx *renderer.Context, attr types.DocumentAttributeSubstitution) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	if value, found := ctx.Document.Attributes[attr.Name]; found {
		result.WriteString(fmt.Sprintf("%v", value))
	} else {
		result.WriteString(fmt.Sprintf("{%s}", attr.Name))
	}
	return result.Bytes(), nil
}
