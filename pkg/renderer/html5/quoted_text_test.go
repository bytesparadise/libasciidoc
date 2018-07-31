package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("quoted texts", func() {

	Context("bold content", func() {
		It("bold content alone", func() {
			actualContent := "*bold content*"
			expectedResult := `<div class="paragraph">
<p><strong>bold content</strong></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("bold content in sentence", func() {
			actualContent := "some *bold content*."
			expectedResult := `<div class="paragraph">
<p>some <strong>bold content</strong>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("italic content", func() {
		It("italic content alone", func() {
			actualContent := "_italic content_"
			expectedResult := `<div class="paragraph">
<p><em>italic content</em></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("italic content in sentence", func() {

			actualContent := "some _italic content_."
			expectedResult := `<div class="paragraph">
<p>some <em>italic content</em>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("monospace content", func() {
		It("monospace content alone", func() {
			actualContent := "`monospace content`"
			expectedResult := `<div class="paragraph">
<p><code>monospace content</code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("monospace content in sentence", func() {

			actualContent := "some `monospace content`."
			expectedResult := `<div class="paragraph">
<p>some <code>monospace content</code>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("Nested content", func() {

		It("nested bold quote within bold quote with same punctuation", func() {

			actualContent := "*some *nested bold* content*."
			expectedResult := `<div class="paragraph">
<p><strong>some <strong>nested bold</strong> content</strong>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("italic content within bold quote in sentence", func() {
			actualContent := "some *bold and _italic content_* together."
			expectedResult := `<div class="paragraph">
<p>some <strong>bold and <em>italic content</em></strong> together.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("invalid  content", func() {

		It("italic content within invalid bold quote in sentence", func() {
			actualContent := "some *bold and _italic content_ * together."
			expectedResult := `<div class="paragraph">
<p>some *bold and <em>italic content</em> * together.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("invalid italic content within bold quote in sentence", func() {

			actualContent := "some *bold and _italic content _ together*."
			expectedResult := `<div class="paragraph">
<p>some <strong>bold and _italic content _ together</strong>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("prevented substitution", func() {

		It("esacped bold content in sentence", func() {
			actualContent := "some \\*bold content*."
			expectedResult := `<div class="paragraph">
<p>some *bold content*.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("italic content within escaped bold quote in sentence", func() {
			actualContent := "some \\*bold and _italic content_* together."
			expectedResult := `<div class="paragraph">
<p>some *bold and <em>italic content</em>* together.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})

})
