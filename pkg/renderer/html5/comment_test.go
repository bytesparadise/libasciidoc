package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("comments", func() {

	Context("single line comments", func() {

		It("single line comment alone", func() {
			actualDocument := `// A single-line comment.`
			expectedResult := ""
			verify(GinkgoT(), expectedResult, actualDocument)
		})

		It("single line comment at end of line", func() {
			actualDocument := `foo // A single-line comment.`
			expectedResult := `<div class="paragraph">
<p>foo // A single-line comment.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualDocument)
		})

		It("single line comment within a paragraph", func() {
			actualDocument := `a first line
// A single-line comment.
another line`
			expectedResult := `<div class="paragraph">
<p>a first line
another line</p>
</div>`
			verify(GinkgoT(), expectedResult, actualDocument)
		})
	})

	Context("comment blocks", func() {

		It("comment block alone", func() {
			actualDocument := `//// 
a *comment* block
with multiple lines
////`
			expectedResult := ""
			verify(GinkgoT(), expectedResult, actualDocument)
		})

		It("comment block with paragraphs around", func() {
			actualDocument := `a first paragraph
//// 
a *comment* block
with multiple lines
////
a second paragraph`
			expectedResult := `<div class="paragraph">
<p>a first paragraph</p>
</div>
<div class="paragraph">
<p>a second paragraph</p>
</div>`
			verify(GinkgoT(), expectedResult, actualDocument)
		})
	})

})
