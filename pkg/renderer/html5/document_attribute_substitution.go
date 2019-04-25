package html5

import (
	"bytes"
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func processAttributeDeclaration(ctx *renderer.Context, attr types.DocumentAttributeDeclaration) []byte {
	ctx.Document.Attributes.AddDeclaration(attr)
	return []byte{}
}

func processAttributeReset(ctx *renderer.Context, attr types.DocumentAttributeReset) []byte {
	ctx.Document.Attributes.Reset(attr)
	return []byte{}
}

func renderAttributeSubstitution(ctx *renderer.Context, attr types.DocumentAttributeSubstitution) []byte {
	result := bytes.NewBuffer(nil)
	if value, found := ctx.Document.Attributes[attr.Name]; found {
		result.WriteString(fmt.Sprintf("%v", value))
	} else if value, found := predefined[attr.Name]; found {
		result.WriteString(fmt.Sprintf("%v", value))
	} else {
		result.WriteString(fmt.Sprintf("{%s}", attr.Name))
	}
	return result.Bytes()
}
