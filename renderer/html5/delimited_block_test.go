package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Delimited Blocks", func() {

	Context("Fenced blocks", func() {

		It("Fenced block with multiple lines", func() {
			actualContent := "```\nsome source code\n\nhere\n```"
			expected := `<div class="listingblock">
			<div class="content">
			<pre class="highlight"><code>some source code
			
			here</code></pre>
			</div>
			</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})

	Context("Listing blocks", func() {

		It("Listing block with multiple lines", func() {
			actualContent := `----
some source code

here
----`
			expected := `<div class="listingblock">
			<div class="content">
			<pre class="highlight"><code>some source code
			
			here</code></pre>
			</div>
			</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})

	Context("Literal blocks", func() {

		It("literal block with multiple lines", func() {
			actualContent := ` some source code
here`
			expected := `<div class="literalblock">
<div class="content">
<pre> some source code
here</pre>
</div>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})
})
