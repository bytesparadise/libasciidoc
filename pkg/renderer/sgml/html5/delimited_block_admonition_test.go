package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("admonition blocks", func() {

	Context("as delimited blocks", func() {

		It("admonition block with multiple elements alone", func() {
			source := `[NOTE]
====
some listing code
with *bold content*

* and a list item

====`
			expected := `<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
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
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition block with ID and title", func() {
			source := `[NOTE]
[#id-for-admonition-block]
.title for admonition block
====
some listing code
with *bold content*

* and a list item
====`
			expected := `<div id="id-for-admonition-block" class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="title">title for admonition block</div>
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
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
		It("admonition block with ID, title and icon", func() {
			source := `:icons: font
			
[NOTE]
[#id-for-admonition-block]
.title for admonition block
====
some listing code
with *bold content*

* and a list item

====`
			expected := `<div id="id-for-admonition-block" class="admonitionblock note">
<table>
<tr>
<td class="icon">
<i class="fa icon-note" title="Note"></i>
</td>
<td class="content">
<div class="title">title for admonition block</div>
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
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition block with ID, title and SVG icon", func() {
			source := `:icons:
:icontype: svg
			
[NOTE]
[#id-for-admonition-block]
.title for admonition block
====
some listing code
with *bold content*

* and a list item

====`
			expected := `<div id="id-for-admonition-block" class="admonitionblock note">
<table>
<tr>
<td class="icon">
<img src="images/icons/note.svg" alt="Note">
</td>
<td class="content">
<div class="title">title for admonition block</div>
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
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition paragraph and admonition block with multiple elements", func() {
			source := `[CAUTION]                      
this is an admonition paragraph.
								
								
[NOTE]                         
.Title2                        
====                           
This is an admonition block
								
with another paragraph    
====`
			expected := `<div class="admonitionblock caution">
<table>
<tr>
<td class="icon">
<div class="title">Caution</div>
</td>
<td class="content">
this is an admonition paragraph.
</td>
</tr>
</table>
</div>
<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="title">Title2</div>
<div class="paragraph">
<p>This is an admonition block</p>
</div>
<div class="paragraph">
<p>with another paragraph</p>
</div>
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition paragraph with an icon", func() {
			source := `:icons: font

TIP: an admonition text on
2 lines.`
			expected := `<div class="admonitionblock tip">
<table>
<tr>
<td class="icon">
<i class="fa icon-tip" title="Tip"></i>
</td>
<td class="content">
an admonition text on
2 lines.
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition paragraph with ID, title and icon", func() {
			source := `:icons: font

[#id-for-admonition-block]
.title for the admonition block
TIP: an admonition text on 1 line.
`
			expected := `<div id="id-for-admonition-block" class="admonitionblock tip">
<table>
<tr>
<td class="icon">
<i class="fa icon-tip" title="Tip"></i>
</td>
<td class="content">
<div class="title">title for the admonition block</div>
an admonition text on 1 line.
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
