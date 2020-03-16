package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("index terms", func() {

	It("index term in existing paragraph line", func() {
		source := `a paragraph with an ((index)) term.`
		expected := `<div class="paragraph">
<p>a paragraph with an index term.</p>
</div>`
		Expect(RenderHTML5Body(source)).To(Equal(expected))
	})

	It("index term in separate paragraph line", func() {
		source := `((foo_bar_baz _italic_))
a paragraph with an index term.`
		expected := `<div class="paragraph">
<p>foo_bar_baz <em>italic</em>
a paragraph with an index term.</p>
</div>`
		Expect(RenderHTML5Body(source)).To(Equal(expected))
	})
})

var _ = Describe("concealed index terms", func() {

	It("index term in existing paragraph line", func() {
		source := `a paragraph with an index term (((index, term, here))).`
		expected := `<div class="paragraph">
<p>a paragraph with an index term .</p>
</div>`
		Expect(RenderHTML5Body(source)).To(Equal(expected))
	})

	It("index term in single paragraph line", func() {
		source := `(((index, term)))
a paragraph with an index term.`
		expected := `<div class="paragraph">
<p>a paragraph with an index term.</p>
</div>`
		Expect(RenderHTML5Body(source)).To(Equal(expected))
	})
})
