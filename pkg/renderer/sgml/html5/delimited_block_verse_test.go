package html5_test

import (
	"strings"

	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("verse blocks", func() {

	Context("delimited blocks", func() {

		It("single-line verse with author and title ", func() {
			source := `[verse, john doe, verse title]
____
some *verse* content

____`
			expected := `<div class="verseblock">
<pre class="content">some <strong>verse</strong> content</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("single-line verse with author, id and title ", func() {
			source := `[verse, john doe, verse title]
[#id-for-verse-block]
.title for verse block
____
some *verse* content
____`
			expected := `<div id="id-for-verse-block" class="verseblock">
<div class="title">title for verse block</div>
<pre class="content">some <strong>verse</strong> content</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line verse with author and title", func() {
			source := `[verse, john doe, verse title]
____
- some 
- verse 
- content

and more!

____`
			expected := `<div class="verseblock">
<pre class="content">- some
- verse
- content

and more!</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("single-line verse with author only", func() {
			source := `[verse, john doe]
____
some verse content
____`
			expected := `<div class="verseblock">
<pre class="content">some verse content</pre>
<div class="attribution">
&#8212; john doe
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("single-line verse with title only", func() {
			source := `[verse, , verse title]
____
some verse content
____`
			expected := `<div class="verseblock">
<pre class="content">some verse content</pre>
<div class="attribution">
&#8212; verse title
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line verse without author and title", func() {
			source := `[verse]
____
lines 
	and tabs 
are preserved

____`

			expected := `<div class="verseblock">
<pre class="content">lines
	and tabs
are preserved</pre>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("empty verse without author and title", func() {
			source := `[verse]
____
____`
			expected := `<div class="verseblock">
<pre class="content"></pre>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})

		Context("with custom substitutions", func() {

			source := `:github-url: https://github.com
			
[subs="$SUBS"]
[verse, john doe, verse title]
____
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item
____

<1> a callout
`

			It("should apply the default substitution", func() {
				s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]", "")
				expected := `<div class="verseblock">
<pre class="content">a link to <a href="https://example.com" class="bare">https://example.com</a> &lt;1&gt;
and &lt;more text&gt; on the<br>
<strong>next</strong> lines with a link to <a href="https://github.com" class="bare">https://github.com</a>

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'normal' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "normal")
				expected := `<div class="verseblock">
<pre class="content">a link to <a href="https://example.com" class="bare">https://example.com</a> &lt;1&gt;
and &lt;more text&gt; on the<br>
<strong>next</strong> lines with a link to <a href="https://github.com" class="bare">https://github.com</a>

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'quotes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "quotes")
				expected := `<div class="verseblock">
<pre class="content">a link to https://example.com[] <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "macros")
				expected := `<div class="verseblock">
<pre class="content">a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'attributes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "attributes")
				expected := `<div class="verseblock">
<pre class="content">a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to https://github.com[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'attributes,macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
				expected := `<div class="verseblock">
<pre class="content">a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
*next* lines with a link to <a href="https://github.com" class="bare">https://github.com</a>

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'specialchars' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "specialchars")
				expected := `<div class="verseblock">
<pre class="content">a link to https://example.com[] &lt;1&gt;
and &lt;more text&gt; on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'replacements' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "replacements")
				expected := `<div class="verseblock">
<pre class="content">a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'post_replacements' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "post_replacements")
				expected := `<div class="verseblock">
<pre class="content">a link to https://example.com[] <1>
and <more text> on the<br>
*next* lines with a link to {github-url}[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'quotes,macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
				expected := `<div class="verseblock">
<pre class="content">a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'macros,quotes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
				expected := `<div class="verseblock">
<pre class="content">a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})

			It("should apply the 'none' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "none")
				expected := `<div class="verseblock">
<pre class="content">a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(s)).To(MatchHTML(expected))
			})
		})
	})
})
