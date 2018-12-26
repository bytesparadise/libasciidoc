package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("paragraphs", func() {

	Context("paragraphs", func() {

		It("a standalone paragraph with special character", func() {
			actualContent := `*bold content* 
& more content afterwards`
			expectedResult := `<div class="paragraph">
<p><strong>bold content</strong>
&amp; more content afterwards</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("a standalone paragraph with trailing spaces", func() {
			actualContent := `*bold content*    
   & more content afterwards...`
			expectedResult := `<div class="paragraph">
<p><strong>bold content</strong>
   &amp; more content afterwards&#8230;&#8203;</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("a standalone paragraph with an ID and a title", func() {
			actualContent := `[#foo]
.a title
*bold content* with more content afterwards...`
			expectedResult := `<div id="foo" class="paragraph">
<div class="doctitle">a title</div>
<p><strong>bold content</strong> with more content afterwards&#8230;&#8203;</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("2 paragraphs and blank line", func() {
			actualContent := `
*bold content* with more content afterwards...

and here another paragraph

`
			expectedResult := `<div class="paragraph">
<p><strong>bold content</strong> with more content afterwards&#8230;&#8203;</p>
</div>
<div class="paragraph">
<p>and here another paragraph</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("paragraphs with line break", func() {

		It("with explicit line break", func() {
			actualContent := `foo +
bar
baz`
			expectedResult := `<div class="paragraph">
<p>foo<br>
bar
baz</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("with paragraph attribute", func() {

			actualContent := `[%hardbreaks]
foo
bar
baz`
			expectedResult := `<div class="paragraph">
<p>foo<br>
bar<br>
baz</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("with document attribute", func() {
			actualContent := `:hardbreaks:
foo
bar
baz`
			expectedResult := `<div class="paragraph">
<p>foo<br>
bar<br>
baz</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("admonition paragraphs", func() {

		It("note admonition paragraph", func() {
			actualContent := `NOTE: this is a note.`
			expectedResult := `<div class="admonitionblock note">
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
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiline warning admonition paragraph", func() {
			actualContent := `WARNING: this is a multiline
warning!`
			expectedResult := `<div class="admonitionblock warning">
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
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("admonition note paragraph with id and title", func() {
			actualContent := `[[foo]]
.bar
NOTE: this is a note.`
			expectedResult := `<div id="foo" class="admonitionblock note">
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
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("admonition paragraphs", func() {

		It("simple caution admonition paragraph", func() {
			actualContent := `[CAUTION] 
this is a caution!`
			expectedResult := `<div class="admonitionblock caution">
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
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiline caution admonition paragraph with title and id", func() {
			actualContent := `[[foo]]
[CAUTION] 
.bar
this is a
*caution*!`
			expectedResult := `<div id="foo" class="admonitionblock caution">
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
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("verse paragraphs", func() {

		It("paragraph as a verse with author and title", func() {
			actualContent := `[verse, john doe, verse title]
I am a verse paragraph.`
			expectedResult := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a verse with author, title and other attributes", func() {
			actualContent := `[[universal]]
[verse, john doe, verse title]
.universe
I am a verse paragraph.`
			expectedResult := `<div id="universal" class="verseblock">
<div class="title">universe</div>
<pre class="content">I am a verse paragraph.</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a verse with empty title", func() {
			actualContent := `[verse, john doe, ]
I am a verse paragraph.`
			expectedResult := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a verse without title", func() {
			actualContent := `[verse, john doe ]
I am a verse paragraph.`
			expectedResult := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a verse with empty author", func() {
			actualContent := `[verse,  ]
I am a verse paragraph.`
			expectedResult := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a verse without author", func() {
			actualContent := `[verse]
I am a verse paragraph.`
			expectedResult := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("image block as a verse", func() {
			actualContent := `[verse, john doe, verse title]
image::foo.png[]`
			expectedResult := `<div class="verseblock">
<pre class="content">image::foo.png[]</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("quote paragraphs", func() {

		It("single-line quote paragraph with author and title", func() {
			actualContent := `[quote, john doe, quote title]
some *quote* content`
			expectedResult := `<div class="quoteblock">
<blockquote>
some <strong>quote</strong> content
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a quote with author, title and other attributes", func() {
			actualContent := `[[universal]]
[quote, john doe, quote title]
.universe
I am a quote paragraph.`
			expectedResult := `<div id="universal" class="quoteblock">
<div class="title">universe</div>
<blockquote>
I am a quote paragraph.
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a quote with empty title", func() {
			actualContent := `[quote, john doe, ]
I am a quote paragraph.`
			expectedResult := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a quote without title", func() {
			actualContent := `[quote, john doe ]
I am a quote paragraph.`
			expectedResult := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a quote with empty author", func() {
			actualContent := `[quote,  ]
I am a quote paragraph.`
			expectedResult := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("paragraph as a quote without author", func() {
			actualContent := `[quote]
I am a quote paragraph.`
			expectedResult := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("inline image within a quote", func() {
			actualContent := `[quote, john doe, quote title]
a foo image:foo.png[]`
			expectedResult := `<div class="quoteblock">
<blockquote>
a foo <span class="image"><img src="foo.png" alt="foo"></span>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("image block is NOT a quote", func() {
			actualContent := `[quote, john doe, quote title]
image::foo.png[]`
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})

})
