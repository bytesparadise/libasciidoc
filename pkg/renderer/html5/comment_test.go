package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("comments", func() {

	Context("single line comments", func() {

		It("single line comment alone", func() {
			source := `// A single-line comment.`
			expected := ""
			Expect(RenderHTML5Body(source)).To(Equal(expected))
		})

		It("single line comment at end of line", func() {
			source := `foo // A single-line comment.`
			expected := `<div class="paragraph">
<p>foo // A single-line comment.</p>
</div>`
			Expect(RenderHTML5Body(source)).To(Equal(expected))
		})

		It("single line comment within a paragraph", func() {
			source := `a first line
// A single-line comment.
another line`
			expected := `<div class="paragraph">
<p>a first line
another line</p>
</div>`
			Expect(RenderHTML5Body(source)).To(Equal(expected))
		})
	})

	Context("comment blocks", func() {

		It("comment block alone", func() {
			source := `//// 
a *comment* block
with multiple lines
////`
			expected := ""
			Expect(RenderHTML5Body(source)).To(Equal(expected))
		})

		It("comment block with paragraphs around", func() {
			source := `a first paragraph
//// 
a *comment* block
with multiple lines
////
a second paragraph`
			expected := `<div class="paragraph">
<p>a first paragraph</p>
</div>
<div class="paragraph">
<p>a second paragraph</p>
</div>`
			Expect(RenderHTML5Body(source)).To(Equal(expected))
		})
	})

})
