package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("quote blocks", func() {

	Context("as delimited blocks", func() {

		It("with single-line quote and author and title ", func() {
			source := `[quote, john doe, quote title]
____
some *quote* content

____`
			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some <strong>quote</strong> content</p>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with single-line quote and author and title, and ID and title ", func() {
			source := `[#id-for-quote-block]
[quote, john doe, quote title]
.title for quote block
____
some *quote* content
____`
			expected := `<div id="id-for-quote-block" class="quoteblock">
<div class="title">title for quote block</div>
<blockquote>
<div class="paragraph">
<p>some <strong>quote</strong> content</p>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with multi-line quote and author and title", func() {
			source := `[quote, john doe, quote title]
____

- some 
- quote 
- content

____`
			expected := `<div class="quoteblock">
<blockquote>
<div class="ulist">
<ul>
<li>
<p>some</p>
</li>
<li>
<p>quote</p>
</li>
<li>
<p>content</p>
</li>
</ul>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with multi-line quote and author only and nested listing", func() {
			source := `[quote, john doe]
____
* some
----
* quote 
----
* content
____`
			expected := `<div class="quoteblock">
<blockquote>
<div class="ulist">
<ul>
<li>
<p>some</p>
</li>
</ul>
</div>
<div class="listingblock">
<div class="content">
<pre>* quote</pre>
</div>
</div>
<div class="ulist">
<ul>
<li>
<p>content</p>
</li>
</ul>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with single-line quote and title only", func() {
			source := `[quote, , quote title]
____
some quote content
____`
			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some quote content</p>
</div>
</blockquote>
<div class="attribution">
&#8212; quote title
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with multi-line quote without author and title", func() {
			source := `[quote]
____
lines 
	and tabs 
are preserved, but not trailing spaces   

____`

			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>lines
	and tabs
are preserved, but not trailing spaces</p>
</div>
</blockquote>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with empty quote without author and title", func() {
			source := `[quote]
____
____`
			// asciidoctor will include an empty line in the `blockquote` element, I'm not sure why.
			expected := `<div class="quoteblock">
<blockquote>
</blockquote>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})
	})
})
