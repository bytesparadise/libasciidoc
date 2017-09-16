package html5

import (
	"bytes"
	"fmt"

	asciidoc "github.com/bytesparadise/libasciidoc/context"
	"github.com/bytesparadise/libasciidoc/types"
)

func processAttributeDeclaration(ctx asciidoc.Context, attribute types.DocumentAttributeDeclaration) error {
	ctx.Document.Attributes.Add(attribute)
	return nil
}

func processAttributeReset(ctx asciidoc.Context, attribute types.DocumentAttributeReset) error {
	ctx.Document.Attributes.Reset(attribute)
	return nil
}

func renderAttributeSubstitution(ctx asciidoc.Context, attribute types.DocumentAttributeSubstitution) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	value := ctx.Document.Attributes.Get(attribute)
	if value == nil {
		result.WriteString(fmt.Sprintf("{%s}", attribute.Name))
	} else {
		result.WriteString(*value)
	}
	return result.Bytes(), nil
}
