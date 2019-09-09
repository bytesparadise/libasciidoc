package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("blank lines", func() {

	It("blank line between 2 paragraphs", func() {
		source := `first paragraph 

second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		verify("test.adoc", expected, source)
	})

	It("blank line with spaces and tabs between 2 paragraphs", func() {
		source := `first paragraph
		  
second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		verify("test.adoc", expected, source)
	})

	It("blank lines (tabs) at end of document", func() {
		source := `first paragraph
		
		
		`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>`
		verify("test.adoc", expected, source)
	})

	It("blank lines (spaces) at end of document", func() {
		source := `first paragraph
		
		
        `
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>`
		verify("test.adoc", expected, source)
	})
})
