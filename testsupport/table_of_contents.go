package testsupport

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"io/ioutil"
)

// TableOfContents returns the Table of Contents for the given document
func TableOfContents(doc types.Document) (types.TableOfContents, error) {
	ctx := renderer.NewContext(doc, configuration.NewConfiguration())
	if _, err := html5.Render(ctx, doc, ioutil.Discard); err != nil {
		return ctx.TableOfContents, err
	}
	return ctx.TableOfContents, nil
}
