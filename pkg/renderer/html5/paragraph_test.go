package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("paragraphs", func() {

	Context("paragraphs", func() {

		It("a standalone paragraph with special character", func() {
			source := `*bold content* 
& more content afterwards`
			expected := `<div class="paragraph">
<p><strong>bold content</strong>
&amp; more content afterwards</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("a standalone paragraph with trailing spaces", func() {
			source := `*bold content*    
   & more content afterwards...`
			expected := `<div class="paragraph">
<p><strong>bold content</strong>
   &amp; more content afterwards&#8230;&#8203;</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("a standalone paragraph with an ID and a title", func() {
			source := `[#foo]
.a title
*bold content* with more content afterwards...`
			expected := `<div id="foo" class="paragraph">
<div class="doctitle">a title</div>
<p><strong>bold content</strong> with more content afterwards&#8230;&#8203;</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("2 paragraphs and blank line", func() {
			source := `
*bold content* with more content afterwards...

and here another paragraph

`
			expected := `<div class="paragraph">
<p><strong>bold content</strong> with more content afterwards&#8230;&#8203;</p>
</div>
<div class="paragraph">
<p>and here another paragraph</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph with single quotes", func() {
			source := `a 'subsection' paragraph.`
			expected := `<div class="paragraph">
<p>a 'subsection' paragraph.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("empty paragraph", func() {
			source := `{blank}`
			expected := `<div class="paragraph">
<p></p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})

		It("paragraph with role", func() {
			source := `[.text-left]
some content`
			expected := `<div class="paragraph text-left">
<p>some content</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})
	})

	Context("paragraphs with line break", func() {

		It("with explicit line break", func() {
			source := `foo +
bar
baz`
			expected := `<div class="paragraph">
<p>foo<br>
bar
baz</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with paragraph attribute", func() {

			source := `[%hardbreaks]
foo
bar
baz`
			expected := `<div class="paragraph">
<p>foo<br>
bar<br>
baz</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with document attribute", func() {
			source := `:hardbreaks:
foo
bar
baz`
			expected := `<div class="paragraph">
<p>foo<br>
bar<br>
baz</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph with document attribute resets", func() {
			source := `:author: Xavier
						
:!author1:
:author2!:
a paragraph written by {author}.`
			expected := `<div class="paragraph">
<p>a paragraph written by Xavier.</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("admonition paragraphs", func() {

		It("note admonition paragraph", func() {
			source := `NOTE: this is a note.`
			expected := `<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
this is a note.
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multiline warning admonition paragraph", func() {
			source := `WARNING: this is a multiline
warning!`
			expected := `<div class="admonitionblock warning">
<table>
<tr>
<td class="icon">
<div class="title">Warning</div>
</td>
<td class="content">
this is a multiline
warning!
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition note paragraph with id and title", func() {
			source := `[[foo]]
.bar
NOTE: this is a note.`
			expected := `<div id="foo" class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="title">bar</div>
this is a note.
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("admonition paragraphs", func() {

		It("simple caution admonition paragraph", func() {
			source := `[CAUTION] 
this is a caution!`
			expected := `<div class="admonitionblock caution">
<table>
<tr>
<td class="icon">
<div class="title">Caution</div>
</td>
<td class="content">
this is a caution!
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multiline caution admonition paragraph with title and id", func() {
			source := `[[foo]]
[CAUTION] 
.bar
this is a
*caution*!`
			expected := `<div id="foo" class="admonitionblock caution">
<table>
<tr>
<td class="icon">
<div class="title">Caution</div>
</td>
<td class="content">
<div class="title">bar</div>
this is a
<strong>caution</strong>!
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("verse paragraphs", func() {

		It("paragraph as a verse with author and title", func() {
			source := `[verse, john doe, verse title]
I am a verse paragraph.`
			expected := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a verse with author, title and other attributes", func() {
			source := `[[universal]]
[verse, john doe, verse title]
.universe
I am a verse paragraph.`
			expected := `<div id="universal" class="verseblock">
<div class="title">universe</div>
<pre class="content">I am a verse paragraph.</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a verse with empty title", func() {
			source := `[verse, john doe, ]
I am a verse paragraph.`
			expected := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a verse without title", func() {
			source := `[verse, john doe ]
I am a verse paragraph.`
			expected := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a verse with empty author", func() {
			source := `[verse,  ]
I am a verse paragraph.`
			expected := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a verse without author", func() {
			source := `[verse]
I am a verse paragraph.`
			expected := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image block as a verse", func() {
			source := `[verse, john doe, verse title]
image::foo.png[]`
			expected := `<div class="verseblock">
<pre class="content">image::foo.png[]</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("quote paragraphs", func() {

		It("single-line quote paragraph with author and title", func() {
			source := `[quote, john doe, quote title]
some *quote* content`
			expected := `<div class="quoteblock">
<blockquote>
some <strong>quote</strong> content
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a quote with author, title and other attributes", func() {
			source := `[[universal]]
[quote, john doe, quote title]
.universe
I am a quote paragraph.`
			expected := `<div id="universal" class="quoteblock">
<div class="title">universe</div>
<blockquote>
I am a quote paragraph.
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a quote with empty title", func() {
			source := `[quote, john doe, ]
I am a quote paragraph.`
			expected := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a quote without title", func() {
			source := `[quote, john doe ]
I am a quote paragraph.`
			expected := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a quote with empty author", func() {
			source := `[quote,  ]
I am a quote paragraph.`
			expected := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a quote without author", func() {
			source := `[quote]
I am a quote paragraph.`
			expected := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("inline image within a quote", func() {
			source := `[quote, john doe, quote title]
a foo image:foo.png[]`
			expected := `<div class="quoteblock">
<blockquote>
a foo <span class="image"><img src="foo.png" alt="foo"></span>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image block is NOT a quote", func() {
			source := `[quote, john doe, quote title]
image::foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

})
