package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("quoted texts", func() {

	Context("bold content", func() {

		It("bold content alone", func() {
			source := "*bold content*"
			expected := `<div class="paragraph">
<p><strong>bold content</strong></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("bold content in sentence", func() {
			source := "some *bold content*."
			expected := `<div class="paragraph">
<p>some <strong>bold content</strong>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("italic content", func() {

		It("italic content alone", func() {
			source := "_italic content_"
			expected := `<div class="paragraph">
<p><em>italic content</em></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic content in sentence", func() {

			source := "some _italic content_."
			expected := `<div class="paragraph">
<p>some <em>italic content</em>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("monospace content", func() {

		It("monospace content alone", func() {
			source := "`monospace content`"
			expected := `<div class="paragraph">
<p><code>monospace content</code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("monospace content in sentence", func() {

			source := "some `monospace content`."
			expected := `<div class="paragraph">
<p>some <code>monospace content</code>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("subscript content", func() {

		It("subscript content alone", func() {
			source := "~subscriptcontent~"
			expected := `<div class="paragraph">
<p><sub>subscriptcontent</sub></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("subscript content in sentence", func() {

			source := "some ~subscriptcontent~."
			expected := `<div class="paragraph">
<p>some <sub>subscriptcontent</sub>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("superscript content", func() {

		It("superscript content alone", func() {
			source := "^superscriptcontent^"
			expected := `<div class="paragraph">
<p><sup>superscriptcontent</sup></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("superscript content in sentence", func() {

			source := "some ^superscriptcontent^."
			expected := `<div class="paragraph">
<p>some <sup>superscriptcontent</sup>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("nested content", func() {

		It("nested bold quote within bold quote with same punctuation", func() {

			source := "*some *nested bold* content*."
			expected := `<div class="paragraph">
<p><strong>some *nested bold</strong> content*.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic content within bold quote in sentence", func() {
			source := "some *bold and _italic content_* together."
			expected := `<div class="paragraph">
<p>some <strong>bold and <em>italic content</em></strong> together.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("invalid  content", func() {

		It("italic content within invalid bold quote in sentence", func() {
			source := "some *bold and _italic content_ * together."
			expected := `<div class="paragraph">
<p>some *bold and <em>italic content</em> * together.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("invalid italic content within bold quote in sentence", func() {

			source := "some *bold and _italic content _ together*."
			expected := `<div class="paragraph">
<p>some <strong>bold and _italic content _ together</strong>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("prevented substitution", func() {

		It("escaped bold content in sentence", func() {
			source := "some \\*bold content*."
			expected := `<div class="paragraph">
<p>some *bold content*.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic content within escaped bold quote in sentence", func() {
			source := "some \\*bold and _italic content_* together."
			expected := `<div class="paragraph">
<p>some *bold and <em>italic content</em>* together.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("mixed content", func() {

		It("unbalanced bold in monospace - case 1", func() {
			source := "`*a`"
			expected := `<div class="paragraph">
<p><code>*a</code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("unbalanced bold in monospace - case 2", func() {
			source := "`a*b`"
			expected := `<div class="paragraph">
<p><code>a*b</code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("italic in monospace", func() {
			source := "`_a_`"
			expected := `<div class="paragraph">
<p><code><em>a</em></code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("unbalanced italic in monospace", func() {
			source := "`a_b`"
			expected := `<div class="paragraph">
<p><code>a_b</code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("unparsed bold in monospace", func() {
			source := "`a*b*`"
			expected := `<div class="paragraph">
<p><code>a*b*</code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("parsed subscript in monospace", func() {
			source := "`a~b~`"
			expected := `<div class="paragraph">
<p><code>a<sub>b</sub></code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multiline in monospace - case 1", func() {
			source := "`a\nb`"
			expected := `<div class="paragraph">
<p><code>a
b</code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multiline in monospace - case 2", func() {
			source := "`a\n*b*`"
			expected := `<div class="paragraph">
<p><code>a
<strong>b</strong></code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("link in bold", func() {
			source := "*a link:/[b]*"
			expected := `<div class="paragraph">
<p><strong>a <a href="/">b</a></strong></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image in bold", func() {
			source := "*a image:foo.png[]*"
			expected := `<div class="paragraph">
<p><strong>a <span class="image"><img src="foo.png" alt="foo"></span></strong></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("singleplus passthrough in bold", func() {
			source := "*a +image:foo.png[]+*"
			expected := `<div class="paragraph">
<p><strong>a image:foo.png[]</strong></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("tripleplus passthrough in bold", func() {
			source := "*a +++image:foo.png[]+++*"
			expected := `<div class="paragraph">
<p><strong>a image:foo.png[]</strong></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("link in italic", func() {
			source := "_a link:/[b]_"
			expected := `<div class="paragraph">
<p><em>a <a href="/">b</a></em></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image in italic", func() {
			source := "_a image:foo.png[]_"
			expected := `<div class="paragraph">
<p><em>a <span class="image"><img src="foo.png" alt="foo"></span></em></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("singleplus passthrough in italic", func() {
			source := "_a +image:foo.png[]+_"
			expected := `<div class="paragraph">
<p><em>a image:foo.png[]</em></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("tripleplus passthrough in italic", func() {
			source := "_a +++image:foo.png[]+++_"
			expected := `<div class="paragraph">
<p><em>a image:foo.png[]</em></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("link in monospace", func() {
			source := "`a link:/[b]`"
			expected := `<div class="paragraph">
<p><code>a <a href="/">b</a></code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image in monospace", func() {
			source := "`a image:foo.png[]`"
			expected := `<div class="paragraph">
<p><code>a <span class="image"><img src="foo.png" alt="foo"></span></code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("singleplus passthrough in monospace", func() {
			source := "`a +image:foo.png[]+`"
			expected := `<div class="paragraph">
<p><code>a image:foo.png[]</code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("tripleplus passthrough in monospace", func() {
			source := "`a +++image:foo.png[]+++`"
			expected := `<div class="paragraph">
<p><code>a image:foo.png[]</code></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

})
