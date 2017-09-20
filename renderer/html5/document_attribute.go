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
		switch value := value.(type) {
		case *interface{}:
			result.WriteString(fmt.Sprintf("%v", *value))
		default:
			return nil, fmt.Errorf("unsupported type of attribute to substitute: %T", value)
		}
	}
	return result.Bytes(), nil
}
