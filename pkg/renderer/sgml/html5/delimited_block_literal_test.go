package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("literal blocks", func() {

	Context("with spaces indentation", func() {

		It("literal block from 1-line paragraph with single space", func() {
			source := ` some literal content`
			expected := `<div class="literalblock">
<div class="content">
<pre>some literal content</pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("literal block from paragraph with single space on first line", func() {
			source := ` some literal content
on 3
lines.`
			expected := `<div class="literalblock">
<div class="content">
<pre> some literal content
on 3
lines.</pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("literal block from paragraph with same spaces on each line", func() {
			source := `  some literal content
  on 3
  lines.`
			expected := `<div class="literalblock">
<div class="content">
<pre>some literal content
on 3
lines.</pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("literal block from paragraph with single spaces on each line", func() {
			source := ` literal content
   on many lines  
     has some leading spaces preserved.`
			// note: trailing spaces are removed
			expected := `<div class="literalblock">
<div class="content">
<pre>literal content
  on many lines
    has some leading spaces preserved.</pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("mixing literal block with attributes followed by a paragraph ", func() {
			source := `.title
[#ID]
  some literal content

a normal paragraph.`
			expected := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre>some literal content</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("as block delimiter", func() {

		It("literal block with delimited and attributes followed by 1-line paragraph", func() {
			source := `[#ID]
.title
....
 some literal content with space preserved
....
a normal paragraph.`
			expected := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre> some literal content with space preserved</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with empty blank line around content", func() {
			source := `....

some content

....`
			expected := `<div class="literalblock">
<div class="content">
<pre>some content</pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with preserved spaces and followed by 1-line paragraph", func() {
			source := `[#ID]
.title
....
   some literal 
   content 
....
a normal paragraph.`
			expected := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre>   some literal
   content</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("with literal attribute", func() {

		It("literal block from 1-line paragraph with attribute", func() {
			source := `[literal]   
 literal content
  on many lines 
 has its leading spaces preserved.

a normal paragraph.`
			expected := `<div class="literalblock">
<div class="content">
<pre> literal content
  on many lines
 has its leading spaces preserved.</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("literal block from 2-lines paragraph with attribute", func() {
			source := `[#ID]
[literal]   
.title
some literal content
on two lines.

a normal paragraph.`
			expected := `<div id="ID" class="literalblock">
<div class="title">title</div>
<div class="content">
<pre>some literal content
on two lines.</pre>
</div>
</div>
<div class="paragraph">
<p>a normal paragraph.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("with custom substitutions", func() {

		// testing custom substitutions on a literal block only

		It("should apply the default substitution on block with delimiter", func() {
			source := `:github-url: https://github.com
			
....
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item
....

<1> a callout
`
			expected := `<div class="literalblock">
<div class="content">
<pre>a link to https://example.com[] <b class="conum">(1)</b>
and &lt;more text&gt; on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
		It("should apply the 'normal' substitution on block with delimiter", func() {
			source := `:github-url: https://github.com
			
[subs="normal"]
....
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item
....

<1> a callout
`
			expected := `<div class="literalblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> &lt;1&gt;
and &lt;more text&gt; on the<br>
<strong>next</strong> lines with a link to <a href="https://github.com" class="bare">https://github.com</a>

* not a list item</pre>
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("should apply the 'quotes,macros' substitution on block with delimiter", func() {
			source := `:github-url: https://github.com
			
[subs="quotes,macros"]
....
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item
....

<1> a callout
`
			expected := `<div class="literalblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]

* not a list item</pre>
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("should apply the 'quotes,macros' substitution on block with spaces", func() {
			source := `:github-url: https://github.com
			
[subs="quotes,macros"]
  a link to https://example.com[] <1> 
  and <more text> on the +
  *next* lines with a link to {github-url}[]

<1> a callout
`
			expected := `<div class="literalblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]</pre>
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("should apply the 'quotes,macros' substitution on block with attribute", func() {
			source := `:github-url: https://github.com
			
[subs="quotes,macros"]
[literal]
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

<1> a callout
`
			expected := `<div class="literalblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]</pre>
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
