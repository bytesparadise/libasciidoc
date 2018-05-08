package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("delimited Blocks", func() {

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
	})

	Context("listing blocks", func() {

		It("listing block with multiple lines", func() {
			actualContent := `----
some source code

here
----`
			expectedResult := `<div class="listingblock">
			<div class="content">
			<pre class="highlight"><code>some source code
			
			here</code></pre>
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
<div class="paragraph">
<p>This is an admonition block</p>
</div>
<div class="paragraph">
<p>with another paragraph    </p>
</div>
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
