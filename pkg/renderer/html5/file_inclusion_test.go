package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("file inclusions", func() {

	It("include adoc file with leveloffset attribute", func() {
		actualContent := `= Master Document

preamble

include::includes/chapter-a.adoc[leveloffset=+1]`
		expectedResult := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>preamble</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_chapter_a">Chapter A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content</p>
</div>
</div>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("include non adoc file", func() {
		actualContent := `= Master Document

preamble

include::includes/hello_world.go[]`
		expectedResult := `<div class="paragraph">
<p>preamble</p>
</div>
<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>import &#34;fmt&#34;</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</p>
</div>`
		verify(GinkgoT(), expectedResult, actualContent)
	})

	Context("file inclusion in delimited blocks", func() {

		Context("adoc file inclusion in delimited blocks", func() {

			It("should include adoc file within listing block", func() {
				actualContent := `= Master Document

preamble

----
include::includes/chapter-a.adoc[]
----`
				expectedResult := `<div class="paragraph">
<p>preamble</p>
</div>
<div class="listingblock">
<div class="content">
<pre>= Chapter A

content</pre>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include adoc file within fenced block", func() {
				actualContent := "```\n" +
					"include::includes/chapter-a.adoc[]\n" +
					"```"
				expectedResult := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>= Chapter A

content</code></pre>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include adoc file within example block", func() {
				actualContent := `====
include::includes/chapter-a.adoc[]
====`
				expectedResult := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>= Chapter A</p>
</div>
<div class="paragraph">
<p>content</p>
</div>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include adoc file within quote block", func() {
				actualContent := `____
include::includes/chapter-a.adoc[]
____`
				expectedResult := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>= Chapter A</p>
</div>
<div class="paragraph">
<p>content</p>
</div>
</blockquote>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include adoc file within verse block", func() {
				actualContent := `[verse]
____
include::includes/chapter-a.adoc[]
____`
				expectedResult := `<div class="verseblock">
<pre class="content">= Chapter A

content</pre>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include adoc file within sidebar block", func() {
				actualContent := `****
include::includes/chapter-a.adoc[]
****`
				expectedResult := `<div class="sidebarblock">
<div class="content">
<div class="paragraph">
<p>= Chapter A</p>
</div>
<div class="paragraph">
<p>content</p>
</div>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include adoc file within passthrough block", func() {
				Skip("missing support for passthrough blocks")
				actualContent := `++++
include::includes/chapter-a.adoc[]
++++`
				expectedResult := ``
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})

		Context("other file inclusion in delimited blocks", func() {

			It("should include go file within listing block", func() {
				actualContent := `= Master Document

preamble

----
include::includes/hello_world.go[]
----`
				expectedResult := `<div class="paragraph">
<p>preamble</p>
</div>
<div class="listingblock">
<div class="content">
<pre>package includes

import "fmt"

func helloworld() {
	fmt.Println("hello, world!")
}</pre>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include go file within fenced block", func() {
				actualContent := "```\n" +
					"include::includes/hello_world.go[]\n" +
					"```"
				expectedResult := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>package includes

import "fmt"

func helloworld() {
	fmt.Println("hello, world!")
}</code></pre>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include go file within example block", func() {
				actualContent := `====
include::includes/hello_world.go[]
====`
				expectedResult := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>import &#34;fmt&#34;</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</p>
</div>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include go file within quote block", func() {
				actualContent := `____
include::includes/hello_world.go[]
____`
				expectedResult := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>import &#34;fmt&#34;</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</p>
</div>
</blockquote>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include go file within verse block", func() {
				actualContent := `[verse]
____
include::includes/hello_world.go[]
____`
				expectedResult := `<div class="verseblock">
<pre class="content">package includes

import &#34;fmt&#34;

func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</pre>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include go file within sidebar block", func() {
				actualContent := `****
include::includes/hello_world.go[]
****`
				expectedResult := `<div class="sidebarblock">
<div class="content">
<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>import &#34;fmt&#34;</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</p>
</div>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include go file within passthrough block", func() {
				Skip("missing support for passthrough blocks")
				actualContent := `++++
include::includes/hello_world.go[]
++++`
				expectedResult := ``
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})
	})
})
