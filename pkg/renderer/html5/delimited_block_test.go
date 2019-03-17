package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("delimited blocks", func() {

	Context("fenced blocks", func() {

		It("fenced block with multiple lines", func() {
			actualContent := "```\nsome source code\n\nhere\n\n\n\n```"
			expectedResult := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>some source code

here</code></pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("fenced block with id and title", func() {
			actualContent := "[#id-for-fences]\n.fenced block title\n```\nsome source code\n\nhere\n\n\n\n```"
			expectedResult := `<div id="id-for-fences" class="listingblock">
<div class="title">fenced block title</div>
<div class="content">
<pre class="highlight"><code>some source code

here</code></pre>
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

		It("listing block with ID and title", func() {
			actualContent := `[#id-for-listing-block]
.listing block title
----
some source code
----`
			expectedResult := `<div id="id-for-listing-block" class="listingblock">
<div class="title">listing block title</div>
<div class="content">
<pre>some source code</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})

	Context("source blocks", func() {

		It("with source attribute only", func() {
			actualContent := `[source]
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expectedResult := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>require 'sinatra'

get '/hi' do
  "Hello World!"
end</code></pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("with title, source and languages attributes", func() {
			actualContent := `[source,ruby]
.Source block title
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expectedResult := `<div class="listingblock">
<div class="title">Source block title</div>
<div class="content">
<pre class="highlight"><code class="language-ruby" data-lang="ruby">require 'sinatra'

get '/hi' do
  "Hello World!"
end</code></pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("with id, title, source and languages attributes", func() {
			actualContent := `[#id-for-source-block]
[source,ruby]
.app.rb
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expectedResult := `<div id="id-for-source-block" class="listingblock">
<div class="title">app.rb</div>
<div class="content">
<pre class="highlight"><code class="language-ruby" data-lang="ruby">require 'sinatra'

get '/hi' do
  "Hello World!"
end</code></pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("example blocks", func() {

		It("example block with multiple elements - case 1", func() {
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

		It("example block with multiple elements - case 2", func() {
			actualContent := `====
*bold content*

and more content
====`
			expectedResult := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p><strong>bold content</strong></p>
</div>
<div class="paragraph">
<p>and more content</p>
</div>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("example block with multiple elements - case 3", func() {
			actualContent := `====
*bold content*

and "more" content
====`
			expectedResult := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p><strong>bold content</strong></p>
</div>
<div class="paragraph">
<p>and &#34;more&#34; content</p>
</div>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("example block with ID and title", func() {
			actualContent := `[#id-for-example-block]
.example block title
====
foo
====`
			expectedResult := `<div id="id-for-example-block" class="exampleblock">
<div class="title">Example 1. example block title</div>
<div class="content">
<div class="paragraph">
<p>foo</p>
</div>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("admonition blocks", func() {

		It("admonition block with multiple elements alone", func() {
			actualContent := `[NOTE]
====
some listing code
with *bold content*

* and a list item
====`
			expectedResult := `<div class="admonitionblock note">
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

		It("admonition block with ID and title", func() {
			actualContent := `[NOTE]
[#id-for-admonition-block]
.title for admonition block
====
some listing code
with *bold content*

* and a list item
====`
			expectedResult := `<div id="id-for-admonition-block" class="admonitionblock note">
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
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
		It("admonition block with ID, title and icon", func() {
			actualContent := `:icons: font
			
[NOTE]
[#id-for-admonition-block]
.title for admonition block
====
some listing code
with *bold content*

* and a list item
====`
			expectedResult := `<div id="id-for-admonition-block" class="admonitionblock note">
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

		It("admonition paragraph with an icon", func() {
			actualContent := `:icons: font

TIP: an admonition text on
2 lines.`
			expectedResult := `<div class="admonitionblock tip">
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
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("admonition paragraph with ID, title and icon", func() {
			actualContent := `:icons: font

[#id-for-admonition-block]
.title for the admonition block
TIP: an admonition text on 1 line.`
			expectedResult := `<div id="id-for-admonition-block" class="admonitionblock tip">
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

		It("single-line quote with author and title, and ID and title ", func() {
			actualContent := `[#id-for-quote-block]
[quote, john doe, quote title]
.title for quote block
____
some *quote* content
____`
			expectedResult := `<div id="id-for-quote-block" class="quoteblock">
<div class="title">title for quote block</div>
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
			// asciidoctor will include an empty line in the `blockquote` element, I'm not sure why.
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

		It("single-line verse with author, id and title ", func() {
			actualContent := `[verse, john doe, verse title]
[#id-for-verse-block]
.title for verse block
____
some *verse* content
____`
			expectedResult := `<div id="id-for-verse-block" class="verseblock">
<div class="title">title for verse block</div>
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

	Context("sidebar blocks", func() {

		It("sidebar block with paragraph", func() {
			actualContent := `****
some *verse* content
****`
			expectedResult := `<div class="sidebarblock">
<div class="content">
<div class="paragraph">
<p>some <strong>verse</strong> content</p>
</div>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("sidebar block with id, title, paragraph and sourcecode block", func() {
			actualContent := `[#id-for-sidebar]
.title for sidebar
****
some *verse* content
----
foo
bar
----
****`
			expectedResult := `<div id="id-for-sidebar" class="sidebarblock">
<div class="content">
<div class="title">title for sidebar</div>
<div class="paragraph">
<p>some <strong>verse</strong> content</p>
</div>
<div class="listingblock">
<div class="content">
<pre>foo
bar</pre>
</div>
</div>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
