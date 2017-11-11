package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Blank lines", func() {
	It("blank line between 2 paragraphs", func() {
		actualContent := `first paragraph

second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		verify(GinkgoT(), expected, actualContent)
	})

	It("blank line with spaces and tabs between 2 paragraphs", func() {
		actualContent := `first paragraph
		  
second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		verify(GinkgoT(), expected, actualContent)
	})

	It("blank lines (tabs) at end of document", func() {
		actualContent := `first paragraph
		
		
		`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>`
		verify(GinkgoT(), expected, actualContent)
	})

	It("blank lines (spaces) at end of document", func() {
		actualContent := `first paragraph
		
		
        `
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>`
		verify(GinkgoT(), expected, actualContent)
	})
})
