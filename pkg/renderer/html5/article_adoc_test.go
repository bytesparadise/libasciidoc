package html5_test

import (
	"bufio"
	"bytes"
	"os"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
		ctx := renderer.NewContext(doc)
		_, err = html5.Render(ctx, buff)
		Expect(err).ToNot(HaveOccurred())
	})
})
