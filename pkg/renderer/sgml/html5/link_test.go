package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("links", func() {

	Context("external links", func() {

		It("external link without text", func() {

			source := "a link to https://foo.com[]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://foo.com" class="bare">https://foo.com</a>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("external link with quoted text", func() {
			source := "https://foo.com[_a_ *b* `c`]"
			expected := `<div class="paragraph">
<p><a href="https://foo.com"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("external link with unquoted text having comma", func() {
			source := "https://foo.com[A, B, and C]"
			expected := `<div class="paragraph">
<p><a href="https://foo.com">A, B, and C</a></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		// 		It("email link with unquoted text having comma", func() {
		// 			source := "mailto:foo@example.com[A, B, and C]"
		// 			expected := `<div class="paragraph">
		// <p><a href="mailto:foo@example.com?subject=B&amp;body=and+C">A</a></p>
		// </div>`
		// 			Expect(RenderHTML(source)).To(MatchHTML(expected))
		// 		})

		It("email link with quoted text having comma", func() {
			source := `mailto:foo@example.com["A, B, and C"]`
			expected := `<div class="paragraph">
<p><a href="mailto:foo@example.com">A, B, and C</a></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		Context("with document attribute substitutions", func() {

			It("external link with a document attribute substitution for the whole URL", func() {
				source := `:url: https://foo.bar
	
:url: https://foo2.bar
	
a link to {url}`
				expected := `<div class="paragraph">
<p>a link to <a href="https://foo2.bar" class="bare">https://foo2.bar</a></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("external link with two document attribute substitutions only", func() {
				source := `:scheme: https
:path: foo.bar
	
a link to {scheme}://{path}`
				expected := `<div class="paragraph">
<p>a link to <a href="https://foo.bar" class="bare">https://foo.bar</a></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("external link with two document attribute substitutions and a reset", func() {
				source := `:scheme: https
:path: foo.bar

:!path:
	
a link to {scheme}://{path}`
				expected := `<div class="paragraph">
<p>a link to <a href="https://{path}" class="bare">https://{path}</a></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("external link with document attribute in section 0 title", func() {
				source := `= a title to {scheme}://{path} and https://foo.baz
:scheme: https
:path: foo.bar`
				expected := `a title to https://foo.bar and https://foo.baz`
				Expect(RenderHTML5Title(source)).To(Equal(expected))
			})

			It("external link with document attribute in section 1 title", func() {
				source := `:scheme: https
:path: foo.bar
	
== a title to {scheme}://{path} and https://foo.baz`
				expected := `<div class="sect1">
<h2 id="_a_title_to_https_foo_bar_and_https_foo_baz">a title to <a href="https://foo.bar" class="bare">https://foo.bar</a> and <a href="https://foo.baz" class="bare">https://foo.baz</a></h2>
<div class="sectionbody">
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("external link with two document attribute substitutions and a reset", func() {
				source := `:scheme: https
:path: foo.bar

:!path:

a link to {scheme}://{path} and https://foo.baz`
				expected := `<div class="paragraph">
<p>a link to <a href="https://{path}" class="bare">https://{path}</a> and <a href="https://foo.baz" class="bare">https://foo.baz</a></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			source := "a link to link:foo.adoc[]."
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc" class="bare">foo.adoc</a>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link to doc with text", func() {
			source := "a link to link:foo.adoc[foo doc]."
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc">foo doc</a>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link with text having comma", func() {
			source := "a link to link:foo.adoc[A, B, and C]"
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc">A, B, and C</a></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link to external URL with text", func() {
			source := "a link to link:https://foo.bar[foo doc]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://foo.bar">foo doc</a>.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("invalid relative link to doc", func() {
			source := "a link to link:foo.adoc."
			expected := `<div class="paragraph">
<p>a link to link:foo.adoc.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link with quoted text", func() {
			source := "link:/[_a_ *b* `c`]"
			expected := `<div class="paragraph">
<p><a href="/"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		Context("with document attribute substitutions", func() {

			It("relative link with two document attribute substitutions and a reset", func() {
				source := `:scheme: link
:path: foo.bar

:!path:

a link to {scheme}:{path}[] and https://foo.baz`
				expected := `<div class="paragraph">
<p>a link to <a href="{path}" class="bare">{path}</a> and <a href="https://foo.baz" class="bare">https://foo.baz</a></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})
})
