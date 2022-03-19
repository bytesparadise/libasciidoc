package xhtml5_test

import (
	"strings"

	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("delimited blocks", func() {

	Context("listing blocks", func() {

		It("with multiple lines", func() {
			source := `----
some source code

here

----`
			expected := `<div class="listingblock">
<div class="content">
<pre>some source code

here</pre>
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("with ID and title", func() {
			source := `[#id-for-listing-block]
.listing block title
----
some source code
----`
			expected := `<div id="id-for-listing-block" class="listingblock">
<div class="title">listing block title</div>
<div class="content">
<pre>some source code</pre>
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("with ID and title and empty trailing line", func() {
			source := `[#id-for-listing-block]
.listing block title
----
some source code

----`
			expected := `<div id="id-for-listing-block" class="listingblock">
<div class="title">listing block title</div>
<div class="content">
<pre>some source code</pre>
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("with html content", func() {
			source := `----
<a>link</a>
----`
			expected := `<div class="listingblock">
<div class="content">
<pre>&lt;a&gt;link&lt;/a&gt;</pre>
</div>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("with single callout", func() {
			source := `----
import <1>
----
<1> an import`
			expected := `<div class="listingblock">
<div class="content">
<pre>import <b class="conum">(1)</b></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>an import</p>
</li>
</ol>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("with multiple callouts and blankline between calloutitems", func() {
			source := `----
import <1>

func foo() {} <2>
----
<1> an import

<2> a func`
			expected := `<div class="listingblock">
<div class="content">
<pre>import <b class="conum">(1)</b>

func foo() {} <b class="conum">(2)</b></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>an import</p>
</li>
<li>
<p>a func</p>
</li>
</ol>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("with multiple callouts on same line", func() {
			source := `----
import <1> <2><3>

func foo() {} <4>
----
<1> an import
<2> a single import
<3> a single basic import
<4> a func`
			expected := `<div class="listingblock">
<div class="content">
<pre>import <b class="conum">(1)</b><b class="conum">(2)</b><b class="conum">(3)</b>

func foo() {} <b class="conum">(4)</b></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>an import</p>
</li>
<li>
<p>a single import</p>
</li>
<li>
<p>a single basic import</p>
</li>
<li>
<p>a func</p>
</li>
</ol>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("with invalid callout", func() {
			source := `----
import <a>
----
<a> an import`
			expected := `<div class="listingblock">
<div class="content">
<pre>import &lt;a&gt;</pre>
</div>
</div>
<div class="paragraph">
<p>&lt;a&gt; an import</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("with custom substitutions", func() {

		// testing custom substitutions on listing blocks only, as
		// other verbatim blocks (fenced, literal, source, passthrough)
		// share the same implementation

		source := `:github-url: https://github.com
			
[subs="$SUBS"]
----
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item
----

<1> a callout
`

		It("should apply the default substitution", func() {
			s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]", "")
			expected := `<div class="listingblock">
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'normal' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "normal")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> &lt;1&gt;
and &lt;more text&gt; on the<br/>
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'quotes' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "quotes")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'macros' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "macros")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'attributes' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "attributes")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to https://github.com[]

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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'attributes,macros' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
*next* lines with a link to <a href="https://github.com" class="bare">https://github.com</a>

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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'specialchars' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "specialchars")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] &lt;1&gt;
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'replacements' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "replacements")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the +
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'post_replacements' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "post_replacements")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the<br/>
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'quotes,macros' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
			expected := `<div class="listingblock">
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'macros,quotes' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
			expected := `<div class="listingblock">
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'none' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "none")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the +
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
			Expect(RenderXHTML(s)).To(MatchHTML(expected))
		})
	})

})
