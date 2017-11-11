package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Paragraphs", func() {

	It("a standalone paragraph with special character", func() {
		actualContent := `*bold content* 
& more content afterwards...`
		expected := `<div class="paragraph">
<p><strong>bold content</strong> 
&amp; more content afterwards...</p>
</div>`
		verify(GinkgoT(), expected, actualContent)
	})

	It("a standalone paragraph with an ID and a title", func() {
		actualContent := `[#foo]
.a title
*bold content* with more content afterwards...`
		expected := `<div id="foo" class="paragraph">
<div class="doctitle">a title</div>
<p><strong>bold content</strong> with more content afterwards...</p>
</div>`
		verify(GinkgoT(), expected, actualContent)
	})

	It("2 paragraphs and blank line", func() {
		actualContent := `
*bold content* with more content afterwards...

and here another paragraph

`
		expected := `<div class="paragraph">
<p><strong>bold content</strong> with more content afterwards...</p>
</div>
<div class="paragraph">
<p>and here another paragraph</p>
</div>`
		verify(GinkgoT(), expected, actualContent)
	})
})
