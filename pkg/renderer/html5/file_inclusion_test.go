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

	It("include 2 files", func() {
		actualContent := `= Master Document

preamble

include::includes/grandchild-include.adoc[]

include::includes/hello_world.go[]`
		expectedResult := `<div class="paragraph">
<p>preamble</p>
</div>
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
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

	It("include file and append following elements in included section", func() {
		actualContent := `a first paragraph

include::includes/chapter-a.adoc[leveloffset=+1]

a second paragraph

a third paragraph`
		expectedResult := `<div class="paragraph">
<p>a first paragraph</p>
</div>
<div class="sect1">
<h2 id="_chapter_a">Chapter A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content</p>
</div>
<div class="paragraph">
<p>a second paragraph</p>
</div>
<div class="paragraph">
<p>a third paragraph</p>
</div>
</div>
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
		})
	})

	Context("file inclusions with line range", func() {

		Context("file inclusions as paragraph with line range", func() {

			It("should include single line as paragraph", func() {
				actualContent := `include::includes/hello_world.go[lines=1]`
				expectedResult := `<div class="paragraph">
<p>package includes</p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include multiple lines as paragraph", func() {
				actualContent := `include::includes/hello_world.go[lines=5..7]`
				expectedResult := `<div class="paragraph">
<p>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include multiple ranges as paragraph", func() {
				actualContent := `include::includes/hello_world.go[lines=1..2;5..7]`
				expectedResult := `<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})

		Context("file inclusions in listing blocks with line range", func() {

			It("should include single line in listing block", func() {
				actualContent := `----
include::includes/hello_world.go[lines=1]
----`
				expectedResult := `<div class="listingblock">
<div class="content">
<pre>package includes</pre>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include multiple lines in listing block", func() {
				actualContent := `----
include::includes/hello_world.go[lines=5..7]
----`
				expectedResult := `<div class="listingblock">
<div class="content">
<pre>func helloworld() {
	fmt.Println("hello, world!")
}</pre>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("should include multiple ranges in listing block", func() {
				actualContent := `----
include::includes/hello_world.go[lines=1..2;5..7]
----`
				expectedResult := `<div class="listingblock">
<div class="content">
<pre>package includes

func helloworld() {
	fmt.Println("hello, world!")
}</pre>
</div>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})
	})

	Context("recursive file inclusions", func() {

		It("should include child and grandchild content in paragraphs", func() {
			actualContent := `include::includes/parent-include.adoc[]`
			expectedResult := `<div class="paragraph">
<p>first line of parent</p>
</div>
<div class="paragraph">
<p>first line of child</p>
</div>
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of child</p>
</div>
<div class="paragraph">
<p>last line of parent</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("should include child and grandchild content in listing block", func() {
			actualContent := `----
include::includes/parent-include.adoc[]
----`
			expectedResult := `<div class="listingblock">
<div class="content">
<pre>first line of parent

first line of child

first line of grandchild

last line of grandchild

last line of child

last line of parent</pre>
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
