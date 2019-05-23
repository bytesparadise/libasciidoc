package html5_test

import (
	. "github.com/onsi/ginkgo"
)

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

	Context("subscript content", func() {

		It("subscript content alone", func() {
			actualContent := "~subscriptcontent~"
			expectedResult := `<div class="paragraph">
<p><sub>subscriptcontent</sub></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("subscript content in sentence", func() {

			actualContent := "some ~subscriptcontent~."
			expectedResult := `<div class="paragraph">
<p>some <sub>subscriptcontent</sub>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("superscript content", func() {

		It("superscript content alone", func() {
			actualContent := "^superscriptcontent^"
			expectedResult := `<div class="paragraph">
<p><sup>superscriptcontent</sup></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("superscript content in sentence", func() {

			actualContent := "some ^superscriptcontent^."
			expectedResult := `<div class="paragraph">
<p>some <sup>superscriptcontent</sup>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("nested content", func() {

		It("nested bold quote within bold quote with same punctuation", func() {

			actualContent := "*some *nested bold* content*."
			expectedResult := `<div class="paragraph">
<p><strong>some *nested bold</strong> content*.</p>
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

		It("escaped bold content in sentence", func() {
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

	Context("mixed content", func() {

		It("unbalanced bold in monospace - case 1", func() {
			actualContent := "`*a`"
			expectedResult := `<div class="paragraph">
<p><code>*a</code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("unbalanced bold in monospace - case 2", func() {
			actualContent := "`a*b`"
			expectedResult := `<div class="paragraph">
<p><code>a*b</code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("italic in monospace", func() {
			actualContent := "`_a_`"
			expectedResult := `<div class="paragraph">
<p><code><em>a</em></code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("unbalanced italic in monospace", func() {
			actualContent := "`a_b`"
			expectedResult := `<div class="paragraph">
<p><code>a_b</code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("unparsed bold in monospace", func() {
			actualContent := "`a*b*`"
			expectedResult := `<div class="paragraph">
<p><code>a*b*</code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("parsed subscript in monospace", func() {
			actualContent := "`a~b~`"
			expectedResult := `<div class="paragraph">
<p><code>a<sub>b</sub></code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiline in monospace - case 1", func() {
			actualContent := "`a\nb`"
			expectedResult := `<div class="paragraph">
<p><code>a
b</code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiline in monospace - case 2", func() {
			actualContent := "`a\n*b*`"
			expectedResult := `<div class="paragraph">
<p><code>a
<strong>b</strong></code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("link in bold", func() {
			actualContent := "*a link:/[b]*"
			expectedResult := `<div class="paragraph">
<p><strong>a <a href="/">b</a></strong></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("image in bold", func() {
			actualContent := "*a image:foo.png[]*"
			expectedResult := `<div class="paragraph">
<p><strong>a <span class="image"><img src="foo.png" alt="foo"></span></strong></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("singleplus passthrough in bold", func() {
			actualContent := "*a +image:foo.png[]+*"
			expectedResult := `<div class="paragraph">
<p><strong>a image:foo.png[]</strong></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("tripleplus passthrough in bold", func() {
			actualContent := "*a +++image:foo.png[]+++*"
			expectedResult := `<div class="paragraph">
<p><strong>a image:foo.png[]</strong></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("link in italic", func() {
			actualContent := "_a link:/[b]_"
			expectedResult := `<div class="paragraph">
<p><em>a <a href="/">b</a></em></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("image in italic", func() {
			actualContent := "_a image:foo.png[]_"
			expectedResult := `<div class="paragraph">
<p><em>a <span class="image"><img src="foo.png" alt="foo"></span></em></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("singleplus passthrough in italic", func() {
			actualContent := "_a +image:foo.png[]+_"
			expectedResult := `<div class="paragraph">
<p><em>a image:foo.png[]</em></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("tripleplus passthrough in italic", func() {
			actualContent := "_a +++image:foo.png[]+++_"
			expectedResult := `<div class="paragraph">
<p><em>a image:foo.png[]</em></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("link in monospace", func() {
			actualContent := "`a link:/[b]`"
			expectedResult := `<div class="paragraph">
<p><code>a <a href="/">b</a></code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("image in monospace", func() {
			actualContent := "`a image:foo.png[]`"
			expectedResult := `<div class="paragraph">
<p><code>a <span class="image"><img src="foo.png" alt="foo"></span></code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("singleplus passthrough in monospace", func() {
			actualContent := "`a +image:foo.png[]+`"
			expectedResult := `<div class="paragraph">
<p><code>a image:foo.png[]</code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("tripleplus passthrough in monospace", func() {
			actualContent := "`a +++image:foo.png[]+++`"
			expectedResult := `<div class="paragraph">
<p><code>a image:foo.png[]</code></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})

})
