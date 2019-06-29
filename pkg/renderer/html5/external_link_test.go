package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("links", func() {

	Context("external links", func() {

		It("external link with text only", func() {

			source := "a link to https://foo.com[the website]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://foo.com">the website</a>.</p>
</div>`
			verify(expected, source)
		})

		It("external link without text", func() {

			source := "a link to https://foo.com[]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://foo.com" class="bare">https://foo.com</a>.</p>
</div>`
			verify(expected, source)
		})

		It("external link with quoted text", func() {
			source := "https://foo.com[_a_ *b* `c`]"
			expected := `<div class="paragraph">
<p><a href="https://foo.com"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>`
			verify(expected, source)
		})

		It("external link inside a multiline paragraph", func() {
			source := `a http://website.com
and more text on the
next lines`

			expected := `<div class="paragraph">
<p>a <a href="http://website.com" class="bare">http://website.com</a>
and more text on the
next lines</p>
</div>`
			verify(expected, source)
		})
	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			source := "a link to link:foo.adoc[]."
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc" class="bare">foo.adoc</a>.</p>
</div>`
			verify(expected, source)
		})

		It("relative link to doc with text", func() {
			source := "a link to link:foo.adoc[foo doc]."
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc">foo doc</a>.</p>
</div>`
			verify(expected, source)
		})

		It("relative link to external URL with text", func() {
			source := "a link to link:https://foo.bar[foo doc]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://foo.bar">foo doc</a>.</p>
</div>`
			verify(expected, source)
		})

		It("invalid relative link to doc", func() {
			source := "a link to link:foo.adoc."
			expected := `<div class="paragraph">
<p>a link to link:foo.adoc.</p>
</div>`
			verify(expected, source)
		})

		It("relative link with quoted text", func() {
			source := "link:/[_a_ *b* `c`]"
			expected := `<div class="paragraph">
<p><a href="/"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>`
			verify(expected, source)
		})
	})

})
