package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("comments", func() {

	Context("single line comments", func() {

		It("single line comment alone", func() {
			doc := `// A single-line comment.`
			expected := ""
			verify(expected, doc)
		})

		It("single line comment at end of line", func() {
			doc := `foo // A single-line comment.`
			expected := `<div class="paragraph">
<p>foo // A single-line comment.</p>
</div>`
			verify(expected, doc)
		})

		It("single line comment within a paragraph", func() {
			doc := `a first line
// A single-line comment.
another line`
			expected := `<div class="paragraph">
<p>a first line
another line</p>
</div>`
			verify(expected, doc)
		})
	})

	Context("comment blocks", func() {

		It("comment block alone", func() {
			doc := `//// 
a *comment* block
with multiple lines
////`
			expected := ""
			verify(expected, doc)
		})

		It("comment block with paragraphs around", func() {
			doc := `a first paragraph
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
			verify(expected, doc)
		})
	})

})
