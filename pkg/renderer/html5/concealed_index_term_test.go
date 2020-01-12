package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("concealed index terms", func() {

	Context("draft document", func() {

		It("index term in existing paragraph line", func() {
			source := `a paragraph with an index term (((index, term, here))).`
			expected := `<div class="paragraph">
<p>a paragraph with an index term .</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("index term in single paragraph line", func() {
			source := `(((index, term)))
a paragraph with an index term.`
			expected := `<div class="paragraph">
<p>a paragraph with an index term.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})
})
