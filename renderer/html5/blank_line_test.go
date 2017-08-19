package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering Blank lines", func() {
	It("blank line between 2 paragraphs", func() {
		content := `first paragraph

second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		verify(GinkgoT(), expected, content)
	})
	It("blank line with spaces and tabs between 2 paragraphs", func() {
		content := `first paragraph
		  
second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		verify(GinkgoT(), expected, content)
	})
})
