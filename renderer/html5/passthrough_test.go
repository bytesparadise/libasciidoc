package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("passthroughs", func() {

	Context("tripleplus passthrough", func() {

		It("an empty standalone tripleplus passthrough", func() {
			actualContent := `++++++`
			expectedResult := ``
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("an empty tripleplus passthrough in a paragraph", func() {
			actualContent := `++++++ with more content afterwards...`
			expectedResult := `<div class="paragraph">
<p> with more content afterwards...</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("a standalone tripleplus passthrough", func() {
			actualContent := `+++*bold content*+++`
			expectedResult := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("tripleplus passthrough in paragraph", func() {
			actualContent := `The text +++<u>underline & me</u>+++ is underlined.`
			expectedResult := `<div class="paragraph">
<p>The text <u>underline & me</u> is underlined.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("singleplus Passthrough", func() {

		It("an empty standalone singleplus passthrough", func() {
			actualContent := `++`
			expectedResult := ``
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("an empty singleplus passthrough in a paragraph", func() {
			actualContent := `++ with more content afterwards...`
			expectedResult := `<div class="paragraph">
<p> with more content afterwards...</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("a singleplus passthrough", func() {
			actualContent := `+*bold content*+`
			expectedResult := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("singleplus passthrough in paragraph", func() {
			actualContent := `The text +<u>underline me</u>+ is not underlined.`
			expectedResult := `<div class="paragraph">
<p>The text &lt;u&gt;underline me&lt;/u&gt; is not underlined.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("passthrough Macro", func() {

		It("passthrough macro with single word", func() {
			actualContent := `pass:[hello]`
			expectedResult := `<div class="paragraph">
<p>hello</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("passthrough macro with words", func() {
			actualContent := `pass:[hello, world]`
			expectedResult := `<div class="paragraph">
<p>hello, world</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("empty passthrough macro", func() {
			actualContent := `pass:[]`
			expectedResult := ``
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("passthrough macro with spaces", func() {
			actualContent := `pass:[ *hello*, world ]`
			expectedResult := `<div class="paragraph">
<p> *hello*, world </p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("passthrough macro with line break", func() {
			actualContent := "pass:[hello,\nworld]"
			expectedResult := `<div class="paragraph">
<p>hello,
world</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("passthrough Macro with Quoted Text", func() {

		It("passthrough macro with single quoted word", func() {
			actualContent := `pass:q[*hello*]`
			expectedResult := `<div class="paragraph">
<p><strong>hello</strong></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("passthrough macro with quoted word in sentence", func() {
			actualContent := `pass:q[ a *hello*, world ]`
			expectedResult := `<div class="paragraph">
<p> a <strong>hello</strong>, world </p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
