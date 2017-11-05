package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering Quoted Texts", func() {

	Context("Bold content", func() {
		It("bold content alone", func() {
			content := "*bold content*"
			expected := `<div class="paragraph">
	<p><strong>bold content</strong></p>
	</div>`
			verify(GinkgoT(), expected, content)
		})

		It("bold content in sentence", func() {
			content := "some *bold content*."
			expected := `<div class="paragraph">
	<p>some <strong>bold content</strong>.</p>
	</div>`
			verify(GinkgoT(), expected, content)
		})
	})

	Context("Italic content", func() {
		It("italic content alone", func() {
			content := "_italic content_"
			expected := `<div class="paragraph">
<p><em>italic content</em></p>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("italic content in sentence", func() {

			content := "some _italic content_."
			expected := `<div class="paragraph">
<p>some <em>italic content</em>.</p>
</div>`
			verify(GinkgoT(), expected, content)
		})
	})

	Context("Monospace content", func() {
		It("monospace content alone", func() {
			content := "`monospace content`"
			expected := `<div class="paragraph">
<p><code>monospace content</code></p>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("monospace content in sentence", func() {

			content := "some `monospace content`."
			expected := `<div class="paragraph">
<p>some <code>monospace content</code>.</p>
</div>`
			verify(GinkgoT(), expected, content)
		})
	})

	Context("Nested content", func() {

		It("nested bold quote within bold quote with same punctuation", func() {

			content := "*some *nested bold* content*."
			expected := `<div class="paragraph">
<p><strong>some <strong>nested bold</strong> content</strong>.</p>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("italic content within bold quote in sentence", func() {
			content := "some *bold and _italic content_* together."
			expected := `<div class="paragraph">
<p>some <strong>bold and <em>italic content</em></strong> together.</p>
</div>`
			verify(GinkgoT(), expected, content)
		})
	})

	Context("Invalid  content", func() {

		It("italic content within invalid bold quote in sentence", func() {
			content := "some *bold and _italic content_ * together."
			expected := `<div class="paragraph">
	<p>some *bold and <em>italic content</em> * together.</p>
	</div>`
			verify(GinkgoT(), expected, content)
		})

		It("invalid italic content within bold quote in sentence", func() {

			content := "some *bold and _italic content _ together*."
			expected := `<div class="paragraph">
<p>some <strong>bold and _italic content _ together</strong>.</p>
</div>`
			verify(GinkgoT(), expected, content)
		})
	})

	Context("Prevented substitution", func() {

		It("esacped bold content in sentence", func() {
			content := "some \\*bold content*."
			expected := `<div class="paragraph">
<p>some *bold content*.</p>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("italic content within escaped bold quote in sentence", func() {
			content := "some \\*bold and _italic content_* together."
			expected := `<div class="paragraph">
<p>some *bold and <em>italic content</em>* together.</p>
</div>`
			verify(GinkgoT(), expected, content)
		})

	})

})
