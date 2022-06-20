package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cross references", func() {

	Context("using shorthand syntax", func() {

		It("with custom id to section above with rich title", func() {

			source := `[[thetitle]]
== a *title*

with some content linked to <<thetitle>>!`
			expected := `<div class="sect1">
<h2 id="thetitle">a <strong>title</strong></h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some content linked to <a href="#thetitle">a <strong>title</strong></a>!</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with custom id to section afterwards", func() {

			source := `see <<thetitle>>
			
[#thetitle]
== a *title*
`
			expected := `<div class="paragraph">
<p>see <a href="#thetitle">a <strong>title</strong></a></p>
</div>
<div class="sect1">
<h2 id="thetitle">a <strong>title</strong></h2>
<div class="sectionbody">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with custom id and label", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expected := `<div class="sect1">
<h2 id="thetitle">a title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some content linked to <a href="#thetitle">a label to the title</a>!</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("to paragraph defined later in the document", func() {
			source := `a reference to <<a-paragraph>>

[#a-paragraph]
.another paragraph
some content`
			expected := `<div class="paragraph">
<p>a reference to <a href="#a-paragraph">another paragraph</a></p>
</div>
<div id="a-paragraph" class="paragraph">
<div class="title">another paragraph</div>
<p>some content</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("to section defined later in the document", func() {
			source := `a reference to <<section>>

[#section]
== A section with a link to https://example.com

some content`
			expected := `<div class="paragraph">
<p>a reference to <a href="#section">A section with a link to https://example.com</a></p>
</div>
<div class="sect1">
<h2 id="section">A section with a link to <a href="https://example.com" class="bare">https://example.com</a></h2>
<div class="sectionbody">
<div class="paragraph">
<p>some content</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("to term in labeled list", func() {
			source := `[[a_term]]term::
// a comment

Here's a reference to the definition of <<a_term>>.`
			expected := `<div class="dlist">
<dl>
<dt class="hdlist1"><a id="a_term"></a>term</dt>
</dl>
</div>
<div class="paragraph">
<p>Here&#8217;s a reference to the definition of <a href="#a_term">term</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("invalid section reference", func() {

			source := `[[thetitle]]
== a title

with some content linked to <<thewrongtitle>>!`
			expected := `<div class="sect1">
<h2 id="thetitle">a title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some content linked to <a href="#thewrongtitle">[thewrongtitle]</a>!</p>
</div>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("natural ref to section with plaintext title", func() {
			source := `see <<Section 1>>.

== Section 1`
			expected := `<div class="paragraph">
<p>see <a href="#_section_1">Section 1</a>.</p>
</div>
<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("natural ref to section with rich title", func() {
			source := `see <<Section *1*>>.

== Section *1*`
			expected := `<div class="paragraph">
<p>see <a href="#_section_1">Section <strong>1</strong></a>.</p>
</div>
<div class="sect1">
<h2 id="_section_1">Section <strong>1</strong></h2>
<div class="sectionbody">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("to image with title", func() {
			source := `see <<thecookie>>

[#thecookie]
.A cookie
image::cookie.jpg[]`
			expected := `<div class="paragraph">
<p>see <a href="#thecookie">A cookie</a></p>
</div>
<div id="thecookie" class="imageblock">
<div class="content">
<img src="cookie.jpg" alt="cookie">
</div>
<div class="title">Figure 1. A cookie</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("to image in table cell", func() {
			source := `a reference to <<cookie>>

|===
a|
[#cookie]
.A cookie
image::cookie.png[Cookie]
|===`
			expected := `<div class="paragraph">
<p>a reference to <a href="#cookie">A cookie</a></p>
</div>
<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 100%;">
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><div class="content"><div id="cookie" class="imageblock">
<div class="content">
<img src="cookie.png" alt="Cookie">
</div>
<div class="title">Figure 1. A cookie</div>
</div></div></td>
</tr>
</tbody>
</table>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("using macro syntax", func() {

		It("to other doc with plain text location and rich label", func() {
			source := `some content linked to xref:another-doc.adoc[*another doc*]!`
			expected := `<div class="paragraph">
<p>some content linked to <a href="another-doc.html"><strong>another doc</strong></a>!</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("to other doc with document attribute in location and label", func() {
			source := `:foo: foo-doc
some content linked to xref:{foo}.adoc[another_doc()]!`
			expected := `<div class="paragraph">
<p>some content linked to <a href="foo-doc.html">another_doc()</a>!</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("to other section with empty label", func() {
			source := `some content linked to xref:section_a[]!`
			expected := `<div class="paragraph">
<p>some content linked to <a href="#section_a">[section_a]</a>!</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("to other doc with empty label", func() {
			source := `some content linked to xref:foo.adoc[]!`
			expected := `<div class="paragraph">
<p>some content linked to <a href="foo.html">foo.html</a>!</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
