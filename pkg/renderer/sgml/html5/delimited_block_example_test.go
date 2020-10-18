package html5_test

import (
	"strings"

	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("example blocks", func() {

	Context("delimited blocks", func() {

		It("example block with multiple elements - case 1", func() {
			source := `====
some listing code
with *bold content*

* and a list item

====`
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>some listing code
with <strong>bold content</strong></p>
</div>
<div class="ulist">
<ul>
<li>
<p>and a list item</p>
</li>
</ul>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with multiple elements - case 2", func() {
			source := `====
*bold content*

and more content
====`
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p><strong>bold content</strong></p>
</div>
<div class="paragraph">
<p>and more content</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with multiple elements - case 3", func() {
			source := `====
*bold content*

and "more" content
====`
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p><strong>bold content</strong></p>
</div>
<div class="paragraph">
<p>and "more" content</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with ID and title", func() {
			source := `[#id-for-example-block]
.example block title
====
foo

====`
			expected := `<div id="id-for-example-block" class="exampleblock">
<div class="title">Example 1. example block title</div>
<div class="content">
<div class="paragraph">
<p>foo</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with custom caption and title", func() {
			source := `[caption="Caption A. "]
.example block title
====
foo

====`
			expected := `<div class="exampleblock">
<div class="title">Caption A. example block title</div>
<div class="content">
<div class="paragraph">
<p>foo</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with custom global caption and title", func() {
			source := `:example-caption: Caption

.example block title
====
foo

====`
			expected := `<div class="exampleblock">
<div class="title">Caption 1. example block title</div>
<div class="content">
<div class="paragraph">
<p>foo</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with suppressed caption and title", func() {
			source := `:example-caption!:

.example block title
====
foo

====`
			expected := `<div class="exampleblock">
<div class="title">example block title</div>
<div class="content">
<div class="paragraph">
<p>foo</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("paragraph blocks", func() {

		It("with single plaintext line", func() {
			source := `[example]
some *example* content`
			expected := `<div class="exampleblock">
<div class="content">
some <strong>example</strong> content
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("with custom substitutions", func() {

		source := `:github-url: https://github.com
			
[subs="$SUBS"]
====
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* a list item
====

<1> a callout
`

		It("should apply the default substitution", func() {
			s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]", "")
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a> &lt;1&gt;
and &lt;more text&gt; on the<br>
<strong>next</strong> lines with a link to <a href="https://github.com" class="bare">https://github.com</a></p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a> &lt;1&gt;
and &lt;more text&gt; on the<br>
<strong>next</strong> lines with a link to <a href="https://github.com" class="bare">https://github.com</a></p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to https://example.com[] <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
*next* lines with a link to {github-url}[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to https://github.com[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
*next* lines with a link to <a href="https://github.com" class="bare">https://github.com</a></p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to https://example.com[] &lt;1&gt;
and &lt;more text&gt; on the +
*next* lines with a link to {github-url}[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to https://example.com[] <1>
and <more text> on the<br>
*next* lines with a link to {github-url}[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]</p>
</div>
<div class="ulist">
<ul>
<li>
<p>a list item</p>
</li>
</ul>
</div>
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
