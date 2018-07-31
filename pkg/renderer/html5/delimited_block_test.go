package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("delimited Blocks", func() {

	Context("fenced blocks", func() {

		It("fenced block with multiple lines", func() {
			actualContent := "```\nsome source code\n\nhere\n\n\n\n```"
			expectedResult := `<div class="listingblock">
<div class="content">
<pre>some source code

here</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("listing blocks", func() {

		It("listing block with multiple lines", func() {
			actualContent := `----
some source code

here
----`
			expectedResult := `<div class="listingblock">
<div class="content">
<pre>some source code

here</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("literal blocks", func() {

		It("literal block with multiple lines", func() {
			actualContent := ` some source code
here`
			expectedResult := `<div class="literalblock">
<div class="content">
<pre> some source code
here</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("example blocks", func() {

		It("example block with multiple elements", func() {
			actualContent := `====
some listing code
with *bold content*

* and a list item
====`
			expectedResult := `<div class="exampleblock">
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
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("admonition blocks", func() {

		It("admonition block with multiple elements alone", func() {
			actualContent := `[NOTE]
[#ID]
====
some listing code
with *bold content*

* and a list item
====`
			expectedResult := `<div id="ID" class="admonitionblock note">
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
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("admonition paragraph and admonition block with multiple elements", func() {
			actualContent := `[CAUTION]                      
this is an admonition paragraph.
								
								
[NOTE]                         
.Title2                        
====                           
This is an admonition block
								
with another paragraph    
====      `
			expectedResult := `<div class="admonitionblock caution">
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
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("quote blocks", func() {

		It("single-line quote with author and title ", func() {
			actualContent := `[quote, john doe, quote title]
____
some *quote* content
____`
			expectedResult := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some <strong>quote</strong> content</p>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multi-line quote with author and title", func() {
			actualContent := `[quote, john doe, quote title]
____
- some 
- quote 
- content
____`
			expectedResult := `<div class="quoteblock">
<blockquote>
<div class="ulist">
<ul>
<li>
<p>some</p>
</li>
<li>
<p>quote</p>
</li>
<li>
<p>content</p>
</li>
</ul>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multi-line quote with author only and nested listing", func() {
			actualContent := `[quote, john doe]
____
* some
----
* quote 
----
* content
____`
			expectedResult := `<div class="quoteblock">
<blockquote>
<div class="ulist">
<ul>
<li>
<p>some</p>
</li>
</ul>
</div>
<div class="listingblock">
<div class="content">
<pre>* quote</pre>
</div>
</div>
<div class="ulist">
<ul>
<li>
<p>content</p>
</li>
</ul>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("single-line quote with title only", func() {
			actualContent := `[quote, , quote title]
____
some quote content
____`
			expectedResult := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some quote content</p>
</div>
</blockquote>
<div class="attribution">
&#8212; quote title
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multi-line quote without author and title", func() {
			actualContent := `[quote]
____
lines 
	and tabs 
are preserved, but not trailing spaces   
____`

			expectedResult := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>lines
	and tabs
are preserved, but not trailing spaces</p>
</div>
</blockquote>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("empty quote without author and title", func() {
			actualContent := `[quote]
____
____`
			// asciidoctor will include an emtpy line in the `blockquote` element, I'm not sure why.
			expectedResult := `<div class="quoteblock">
<blockquote>
</blockquote>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)

		})
	})

	Context("verse blocks", func() {

		It("single-line verse with author and title ", func() {
			actualContent := `[verse, john doe, verse title]
____
some *verse* content
____`
			expectedResult := `<div class="verseblock">
<pre class="content">some <strong>verse</strong> content</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multi-line verse with author and title", func() {
			actualContent := `[verse, john doe, verse title]
____
- some 
- verse 
- content
____`
			expectedResult := `<div class="verseblock">
<pre class="content">- some
- verse
- content</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("single-line verse with author only", func() {
			actualContent := `[verse, john doe]
____
some verse content
____`
			expectedResult := `<div class="verseblock">
<pre class="content">some verse content</pre>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("single-line verse with title only", func() {
			actualContent := `[verse, , verse title]
____
some verse content
____`
			expectedResult := `<div class="verseblock">
<pre class="content">some verse content</pre>
<div class="attribution">
&#8212; verse title
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multi-line verse without author and title", func() {
			actualContent := `[verse]
____
lines 
	and tabs 
are preserved
____`

			expectedResult := `<div class="verseblock">
<pre class="content">lines
	and tabs
are preserved</pre>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("empty verse without author and title", func() {
			actualContent := `[verse]
____
____`
			expectedResult := `<div class="verseblock">
<pre class="content"></pre>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)

		})
	})

})
