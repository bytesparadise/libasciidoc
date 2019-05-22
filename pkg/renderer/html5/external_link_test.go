package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("links", func() {

	Context("external links", func() {

		It("external link with text", func() {

			actualContent := "a link to https://foo.com[the website]."
			expectedResult := `<div class="paragraph">
<p>a link to <a href="https://foo.com">the website</a>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("external link without text", func() {

			actualContent := "a link to https://foo.com[]."
			expectedResult := `<div class="paragraph">
<p>a link to <a href="https://foo.com" class="bare">https://foo.com</a>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("external link with quoted text", func() {
			actualContent := "https://foo.com[_a_ *b* `c`]"
			expectedResult := `<div class="paragraph">
<p><a href="https://foo.com"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			actualContent := "a link to link:foo.adoc[]."
			expectedResult := `<div class="paragraph">
<p>a link to <a href="foo.adoc" class="bare">foo.adoc</a>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("relative link to doc with text", func() {
			actualContent := "a link to link:foo.adoc[foo doc]."
			expectedResult := `<div class="paragraph">
<p>a link to <a href="foo.adoc">foo doc</a>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("relative link to external URL with text", func() {
			actualContent := "a link to link:https://foo.bar[foo doc]."
			expectedResult := `<div class="paragraph">
<p>a link to <a href="https://foo.bar">foo doc</a>.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("invalid relative link to doc", func() {
			actualContent := "a link to link:foo.adoc."
			expectedResult := `<div class="paragraph">
<p>a link to link:foo.adoc.</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("relative link with quoted text", func() {
			actualContent := "link:/[_a_ *b* `c`]"
			expectedResult := `<div class="paragraph">
<p><a href="/"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

})
