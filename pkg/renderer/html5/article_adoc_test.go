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
)

var _ = Describe("article.adoc", func() {

	It("should render without failure", func() {
		f, err := os.Open("article.adoc")
		Expect(err).ToNot(HaveOccurred())
		reader := bufio.NewReader(f)
		doc, err := parser.ParseDocument("", reader)
		Expect(err).ToNot(HaveOccurred())
		GinkgoT().Logf("actual document: `%s`", spew.Sdump(doc))
		buff := bytes.NewBuffer(nil)
		rendererCtx := renderer.Wrap(context.Background(), doc)
		_, err = html5.Render(rendererCtx, buff)
		Expect(err).ToNot(HaveOccurred())
	})
})
