package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("markdown-style quote blocks", func() {

	Context("as delimited blocks", func() {

		It("with single marker without author", func() {
			source := `> some text
on *multiple lines*`

			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some text
on <strong>multiple lines</strong></p>
</div>
</blockquote>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with marker on each line without author", func() {
			source := `> some text
> on *multiple lines*`

			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some text
on <strong>multiple lines</strong></p>
</div>
</blockquote>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with marker on each line with author", func() {
			source := `> some text
> on *multiple lines*
> -- John Doe`
			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some text
on <strong>multiple lines</strong></p>
</div>
</blockquote>
<div class="attribution">
&#8212; John Doe
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with marker on each line with author and title", func() {
			source := `.title
> some text
> on *multiple lines*
> -- John Doe`
			expected := `<div class="quoteblock">
<div class="title">title</div>
<blockquote>
<div class="paragraph">
<p>some text
on <strong>multiple lines</strong></p>
</div>
</blockquote>
<div class="attribution">
&#8212; John Doe
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with author only", func() {
			source := `> -- John Doe`
			expected := `<div class="quoteblock">
<blockquote>
</blockquote>
<div class="attribution">
&#8212; John Doe
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
