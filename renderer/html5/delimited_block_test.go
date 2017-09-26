package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering Delimited Blocks", func() {
	Context("Listing blocks", func() {

		It("source block with multiple lines", func() {
			content := "```\nsome source code\n\nhere\n```"
			expected := `<div class="listingblock">
			<div class="content">
			<pre class="highlight"><code>some source code
			
			here</code></pre>
			</div>
			</div>`
			verify(GinkgoT(), expected, content)
		})
	})

	Context("Literal blocks", func() {

		It("literal block with multiple lines", func() {
			content := " some source code\nhere\n"
			expected := `<div class="literalblock">
<div class="content">
<pre> some source code
here</pre>
</div>
</div>`
			verify(GinkgoT(), expected, content)
		})
	})
})
