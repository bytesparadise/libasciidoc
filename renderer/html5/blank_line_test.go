package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("blank lines", func() {

	It("blank line between 2 paragraphs", func() {
		actualContent := `first paragraph

second paragraph`
		expectedResult := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("blank line with spaces and tabs between 2 paragraphs", func() {
		actualContent := `first paragraph
		  
second paragraph`
		expectedResult := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("blank lines (tabs) at end of document", func() {
		actualContent := `first paragraph
		
		
		`
		expectedResult := `<div class="paragraph">
<p>first paragraph</p>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("blank lines (spaces) at end of document", func() {
		actualContent := `first paragraph
		
		
        `
		expectedResult := `<div class="paragraph">
<p>first paragraph</p>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})
})
