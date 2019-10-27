package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("links", func() {

	Context("external links", func() {

		It("external link without text", func() {

			source := "a link to https://foo.com[]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://foo.com" class="bare">https://foo.com</a>.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("external link with quoted text", func() {
			source := "https://foo.com[_a_ *b* `c`]"
			expected := `<div class="paragraph">
<p><a href="https://foo.com"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("external link with text having comma", func() {
			source := "https://foo.com[A, B, and C]"
			expected := `<div class="paragraph">
<p><a href="https://foo.com">A, B, and C</a></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
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
			Expect(source).To(RenderHTML5Body(expected))
		})
	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			source := "a link to link:foo.adoc[]."
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc" class="bare">foo.adoc</a>.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("relative link to doc with text", func() {
			source := "a link to link:foo.adoc[foo doc]."
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc">foo doc</a>.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("relative link with text having comma", func() {
			source := "a link to link:foo.adoc[A, B, and C]"
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc">A, B, and C</a></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("relative link to external URL with text", func() {
			source := "a link to link:https://foo.bar[foo doc]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://foo.bar">foo doc</a>.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("invalid relative link to doc", func() {
			source := "a link to link:foo.adoc."
			expected := `<div class="paragraph">
<p>a link to link:foo.adoc.</p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})

		It("relative link with quoted text", func() {
			source := "link:/[_a_ *b* `c`]"
			expected := `<div class="paragraph">
<p><a href="/"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>`
			Expect(source).To(RenderHTML5Body(expected))
		})
	})

})
