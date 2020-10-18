package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("listing blocks", func() {

	Context("delimited blocks", func() {

		It("with no line", func() {
			source := `----
----`
			expected := `<div class="listingblock">
<div class="content">
<pre></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
