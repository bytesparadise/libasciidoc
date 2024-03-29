package html5_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("paragraphs", func() {

	Context("regular paragraphs", func() {

		It("a standalone paragraph with special character", func() {
			source := `*bold content* 
& more content afterwards`
			expected := `<div class="paragraph">
<p><strong>bold content</strong>
&amp; more content afterwards</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("a standalone paragraph with trailing spaces", func() {
			source := `*bold content*    
   & more content afterwards...`
			expected := `<div class="paragraph">
<p><strong>bold content</strong>
   &amp; more content afterwards&#8230;&#8203;</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("a standalone paragraph with an ID and a title", func() {
			source := `[#foo]
.a title
*bold content* with more content afterwards...`
			expected := `<div id="foo" class="paragraph">
<div class="title">a title</div>
<p><strong>bold content</strong> with more content afterwards&#8230;&#8203;</p>
</div>
`
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
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph with single quotes", func() {
			source := `a 'subsection' paragraph.`
			expected := `<div class="paragraph">
<p>a 'subsection' paragraph.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("empty paragraph", func() {
			source := `{blank}`
			expected := `<div class="paragraph">
<p></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})

		It("paragraph with role", func() {
			source := `[.text-left]
some content`
			expected := `<div class="paragraph text-left">
<p>some content</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph with predefined attribute", func() {
			source := "hello{nbsp}{plus}{nbsp}world"
			expected := `<div class="paragraph">
<p>hello&#160;&#43;&#160;world</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with custom title attribute - explicit and unquoted", func() {
			source := `:title: cookies
			
[title=my {title}]
foo
baz`
			expected := `<div class="paragraph">
<div class="title">my cookies</div>
<p>foo
baz</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with custom title attribute - explicit and single quoted", func() {
			source := `:title: cookies
			
[title='my {title}']
foo
baz`
			expected := `<div class="paragraph">
<div class="title">my cookies</div>
<p>foo
baz</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with custom title attribute - explicit and double quoted", func() {
			source := `:title: cookies
			
[title="my {title}"]
foo
baz`
			expected := `<div class="paragraph">
<div class="title">my cookies</div>
<p>foo
baz</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with custom title attribute - implicit", func() {
			source := `:title: cookies
			
.my {title}
foo
baz`
			expected := `<div class="paragraph">
<div class="title">my cookies</div>
<p>foo
baz</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with multiple substitutions in attributes", func() {
			source := `:role1: ROLE1
:role2: ROLE2

[.{role1}.{role2}]
some content`

			expected := `<div class="paragraph ROLE1 ROLE2">
<p>some content</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with block attributes splitting 2 paragraphs", func() {
			source := `a paragraph
[.left.text-center]
another paragraph with an image image:cookie.jpg[cookie]
`
			expected := `<div class="paragraph">
<p>a paragraph</p>
</div>
<div class="paragraph left text-center">
<p>another paragraph with an image <span class="image"><img src="cookie.jpg" alt="cookie"></span></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with block attributes splitting paragraph and block image", func() {
			source := `a paragraph
[.left.text-center]
image::cookie.jpg[cookie]
`
			expected := `<div class="paragraph">
<p>a paragraph</p>
</div>
<div class="imageblock left text-center">
<div class="content">
<img src="cookie.jpg" alt="cookie">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with non-ASCII characters with unicode", func() {
			source := `Привет!`
			expected := `<div class="paragraph">
<p>Привет!</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})
		It("with non-ASCII characters without unicode", func() {
			source := `Привет!`
			// non-ascii characters are "html escaped"
			expected := `<div class="paragraph">
<p>&#1055;&#1088;&#1080;&#1074;&#1077;&#1090;!</p>
</div>
`
			Expect(RenderHTML(source, configuration.WithAttribute(types.AttrUnicode, false))).To(MatchHTML(expected))

		})

		Context("with custom substitutions", func() {

			// using the same input for all substitution tests
			source := `:github-url: https://github.com
:github-title: GitHub

[subs="$SUBS"]
links to {github-title}: https://github.com[{github-title}] and *<https://github.com[_{github-title}_]>*
and another one using attribute substitution: {github-url}[{github-title}]...
// a single-line comment.`

			It("should apply the 'attributes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "attributes")
				expected := `<div class="paragraph">
<p>links to GitHub: https://github.com[GitHub] and *<https://github.com[_GitHub_]>*
and another one using attribute substitution: https://github.com[GitHub]...</p>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'attributes,macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "macros,attributes")
				// differs from Asciidoctor
				expected := `<div class="paragraph">
<p>links to GitHub: <a href="https://github.com">GitHub</a> and *<<a href="https://github.com">_GitHub_</a>>*
and another one using attribute substitution: https://github.com[GitHub]...</p>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})
		})

		Context("with line break", func() {

			It("at end of line", func() {
				source := `foo +
bar
baz`
				expected := `<div class="paragraph">
<p>foo<br>
bar
baz</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("as paragraph attribute", func() {

				source := `[%hardbreaks]
foo
bar
baz`
				expected := `<div class="paragraph">
<p>foo<br>
bar<br>
baz</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("as document attribute", func() {
				source := `:hardbreaks:
foo
bar
baz`
				expected := `<div class="paragraph">
<p>foo<br>
bar<br>
baz</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("paragraph with document attribute resets", func() {
				source := `:author: Xavier
						
:!author1:
:author2!:
a paragraph written by {author}.`
				expected := `<div class="paragraph">
<p>a paragraph written by Xavier.</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
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
</div>
`
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
</div>
`
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
</div>
`
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
</div>
`
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
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition note replace caption", func() {
			source := `[caption=Aviso]
NOTE: this is a note.`
			expected := `<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Aviso</div>
</td>
<td class="content">
this is a note.
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition warning default caption", func() {
			source := `:warning-caption: Red Alert!
WARNING: Missiles inbound.`
			expected := `<div class="admonitionblock warning">
<table>
<tr>
<td class="icon">
<div class="title">Red Alert!</div>
</td>
<td class="content">
Missiles inbound.
</td>
</tr>
</table>
</div>
`
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
</div>
`
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
</div>
`
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
</div>
`
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
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a verse with empty author", func() {
			source := `[verse,  ]
I am a verse paragraph.`
			expected := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a verse without author", func() {
			source := `[verse]
I am a verse paragraph.`
			expected := `<div class="verseblock">
<pre class="content">I am a verse paragraph.</pre>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		// 		It("image block as a verse", func() {
		// 			source := `[verse, john doe, verse title]
		// image::foo.png[]`
		// 			expected := `<div class="verseblock">
		// <pre class="content">image::foo.png[]</pre>
		// <div class="attribution">
		// &#8212; john doe<br>
		// <cite>verse title</cite>
		// </div>
		// </div>
		// `
		// 			Expect(RenderHTML(source)).To(MatchHTML(expected))
		// 		})
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
</div>
`
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
</div>
`
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
</div>
`
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
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a quote with empty author", func() {
			source := `[quote,  ]
I am a quote paragraph.`
			expected := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("paragraph as a quote without author", func() {
			source := `[quote]
I am a quote paragraph.`
			expected := `<div class="quoteblock">
<blockquote>
I am a quote paragraph.
</blockquote>
</div>
`
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
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("image block is NOT a quote", Pending, func() {
			// needs clarification on how to interpret an image block with such block attributes
			source := `[quote, john doe, quote title]
image::foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="quote" width="john doe" height="quote title">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("thematic break", func() {
			source := "- - -"
			expected := "<hr>\n"
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("source paragraphs", func() {

		It("with source and languages attributes", func() {
			source := `:source-highlighter: chroma

[source,c]
int main(int argc, char **argv);
`
			expected := `<div class="listingblock">
<div class="content">
<pre class="chroma highlight"><code data-lang="c"><span class="tok-kt">int</span> <span class="tok-nf">main</span><span class="tok-p">(</span><span class="tok-kt">int</span> <span class="tok-n">argc</span><span class="tok-p">,</span> <span class="tok-kt">char</span> <span class="tok-o">**</span><span class="tok-n">argv</span><span class="tok-p">);</span></code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
