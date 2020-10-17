package xhtml5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("quote blocks", func() {

	Context("delimited blocks", func() {

		It("single-line quote with author and title ", func() {
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
&#8212; john doe<br/>
<cite>quote title</cite>
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("single-line quote with author and title, and ID and title ", func() {
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
&#8212; john doe<br/>
<cite>quote title</cite>
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line quote with author and title", func() {
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
&#8212; john doe<br/>
<cite>quote title</cite>
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line quote with author only and nested listing", func() {
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
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("single-line quote with title only", func() {
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
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line quote without author and title", func() {
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
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("empty quote without author and title", func() {
			source := `[quote]
____
____`
			// asciidoctor will include an empty line in the `blockquote` element, I'm not sure why.
			expected := `<div class="quoteblock">
<blockquote>
</blockquote>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))

		})
	})
})
