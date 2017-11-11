package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Passthroughs", func() {

	Context("Tripleplus Passthrough", func() {

		It("an empty standalone tripleplus passthrough", func() {
			actualContent := `++++++`
			expected := ``
			verify(GinkgoT(), expected, actualContent)
		})

		It("an empty tripleplus passthrough in a paragraph", func() {
			actualContent := `++++++ with more content afterwards...`
			expected := `<div class="paragraph">
<p> with more content afterwards...</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("a standalone tripleplus passthrough", func() {
			actualContent := `+++*bold content*+++`
			expected := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("tripleplus passthrough in paragraph", func() {
			actualContent := `The text +++<u>underline & me</u>+++ is underlined.`
			expected := `<div class="paragraph">
<p>The text <u>underline & me</u> is underlined.</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})

	Context("Singleplus Passthrough", func() {

		It("an empty standalone singleplus passthrough", func() {
			actualContent := `++`
			expected := ``
			verify(GinkgoT(), expected, actualContent)
		})

		It("an empty singleplus passthrough in a paragraph", func() {
			actualContent := `++ with more content afterwards...`
			expected := `<div class="paragraph">
<p> with more content afterwards...</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("a singleplus passthrough", func() {
			actualContent := `+*bold content*+`
			expected := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("singleplus passthrough in paragraph", func() {
			actualContent := `The text +<u>underline me</u>+ is not underlined.`
			expected := `<div class="paragraph">
<p>The text &lt;u&gt;underline me&lt;/u&gt; is not underlined.</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})

	Context("Passthrough Macro", func() {

		It("passthrough macro with single word", func() {
			actualContent := `pass:[hello]`
			expected := `<div class="paragraph">
<p>hello</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("passthrough macro with words", func() {
			actualContent := `pass:[hello, world]`
			expected := `<div class="paragraph">
<p>hello, world</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("empty passthrough macro", func() {
			actualContent := `pass:[]`
			expected := ``
			verify(GinkgoT(), expected, actualContent)
		})

		It("passthrough macro with spaces", func() {
			actualContent := `pass:[ *hello*, world ]`
			expected := `<div class="paragraph">
<p> *hello*, world </p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("passthrough macro with line break", func() {
			actualContent := "pass:[hello,\nworld]"
			expected := `<div class="paragraph">
<p>hello,
world</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})

	Context("Passthrough Macro with Quoted Text", func() {

		It("passthrough macro with single quoted word", func() {
			actualContent := `pass:q[*hello*]`
			expected := `<div class="paragraph">
<p><strong>hello</strong></p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("passthrough macro with quoted word in sentence", func() {
			actualContent := `pass:q[ a *hello*, world ]`
			expected := `<div class="paragraph">
<p> a <strong>hello</strong>, world </p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})
})
