package html5_test

import (
	"bufio"
	"bytes"
	"context"
	"os"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/require"
)

var _ = Describe("article.adoc", func() {

	It("should render without failure", func() {
		f, err := os.Open("article.adoc")
		require.NoError(GinkgoT(), err, "Error found while opening the document")
		reader := bufio.NewReader(f)
		doc, err := parser.ParseDocument("", reader)
		require.NoError(GinkgoT(), err, "Error found while parsing the document")
		GinkgoT().Logf("actual document: `%s`", spew.Sdump(doc))
		buff := bytes.NewBuffer(nil)
		rendererCtx := renderer.Wrap(context.Background(), doc)
		_, err = html5.Render(rendererCtx, buff)
		require.NoError(GinkgoT(), err)
	})
})
