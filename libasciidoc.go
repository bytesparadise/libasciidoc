package libasciidoc

import (
	"context"
	"io"

	"github.com/bytesparadise/libasciidoc/parser"
	htmlrenderer "github.com/bytesparadise/libasciidoc/renderer/html5"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ConvertToHTML converts the content of the given reader `r` into an HTML document that is written in the given writer `w`. Returns an error if a problem occurred
func ConvertToHTML(r io.Reader, w io.Writer) error {
	doc, err := parser.ParseReader("", r)
	if err != nil {
		return errors.Wrapf(err, "error while parsing the document")
	}
	document := doc.(*types.Document)
	err = htmlrenderer.Render(context.Background(), *document, w)
	if err != nil {
		return errors.Wrapf(err, "error while rendering the document")
	}
	log.Debugf("Done processing document")
	return nil
}
