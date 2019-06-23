package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("file inclusions", func() {

	It("include adoc file with leveloffset attribute", func() {
		source := `= Master Document

preamble

include::includes/chapter-a.adoc[leveloffset=+1]`
		expected := `<div id="preamble">
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
		verify(expected, source)
	})

	It("include non adoc file", func() {
		source := `= Master Document

preamble

include::includes/hello_world.go[]`
		expected := `<div class="paragraph">
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
		verify(expected, source)
	})

	It("include 2 files", func() {
		source := `= Master Document

preamble

include::includes/grandchild-include.adoc[]

include::includes/hello_world.go[]`
		expected := `<div class="paragraph">
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
		verify(expected, source)
	})

	It("include file and append following elements in included section", func() {
		source := `a first paragraph

include::includes/chapter-a.adoc[leveloffset=+1]

a second paragraph

a third paragraph`
		expected := `<div class="paragraph">
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
		verify(expected, source)
	})

	Context("file inclusion in delimited blocks", func() {

		Context("adoc file inclusion in delimited blocks", func() {

			It("should include adoc file within listing block", func() {
				source := `= Master Document

preamble

----
include::includes/chapter-a.adoc[]
----`
				expected := `<div class="paragraph">
<p>preamble</p>
</div>
<div class="listingblock">
<div class="content">
<pre>= Chapter A

content</pre>
</div>
</div>`
				verify(expected, source)
			})

			It("should include adoc file within fenced block", func() {
				source := "```\n" +
					"include::includes/chapter-a.adoc[]\n" +
					"```"
				expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>= Chapter A

content</code></pre>
</div>
</div>`
				verify(expected, source)
			})

			It("should include adoc file within example block", func() {
				source := `====
include::includes/chapter-a.adoc[]
====`
				expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>= Chapter A</p>
</div>
<div class="paragraph">
<p>content</p>
</div>
</div>
</div>`
				verify(expected, source)
			})

			It("should include adoc file within quote block", func() {
				source := `____
include::includes/chapter-a.adoc[]
____`
				expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>= Chapter A</p>
</div>
<div class="paragraph">
<p>content</p>
</div>
</blockquote>
</div>`
				verify(expected, source)
			})

			It("should include adoc file within verse block", func() {
				source := `[verse]
____
include::includes/chapter-a.adoc[]
____`
				expected := `<div class="verseblock">
<pre class="content">= Chapter A

content</pre>
</div>`
				verify(expected, source)
			})

			It("should include adoc file within sidebar block", func() {
				source := `****
include::includes/chapter-a.adoc[]
****`
				expected := `<div class="sidebarblock">
<div class="content">
<div class="paragraph">
<p>= Chapter A</p>
</div>
<div class="paragraph">
<p>content</p>
</div>
</div>
</div>`
				verify(expected, source)
			})

			It("should include adoc file within passthrough block", func() {
				Skip("missing support for passthrough blocks")
				source := `++++
include::includes/chapter-a.adoc[]
++++`
				expected := ``
				verify(expected, source)
			})
		})

		Context("other file inclusion in delimited blocks", func() {

			It("should include go file within listing block", func() {
				source := `= Master Document

preamble

----
include::includes/hello_world.go[]
----`
				expected := `<div class="paragraph">
<p>preamble</p>
</div>
<div class="listingblock">
<div class="content">
<pre>package includes

import &#34;fmt&#34;

func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</pre>
</div>
</div>`
				verify(expected, source)
			})

			It("should include go file within fenced block", func() {
				source := "```\n" +
					"include::includes/hello_world.go[]\n" +
					"```"
				expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>package includes

import "fmt"

func helloworld() {
	fmt.Println("hello, world!")
}</code></pre>
</div>
</div>`
				verify(expected, source)
			})

			It("should include go file within example block", func() {
				source := `====
include::includes/hello_world.go[]
====`
				expected := `<div class="exampleblock">
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
				verify(expected, source)
			})

			It("should include go file within quote block", func() {
				source := `____
include::includes/hello_world.go[]
____`
				expected := `<div class="quoteblock">
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
				verify(expected, source)
			})

			It("should include go file within verse block", func() {
				source := `[verse]
____
include::includes/hello_world.go[]
____`
				expected := `<div class="verseblock">
<pre class="content">package includes

import &#34;fmt&#34;

func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</pre>
</div>`
				verify(expected, source)
			})

			It("should include go file within sidebar block", func() {
				source := `****
include::includes/hello_world.go[]
****`
				expected := `<div class="sidebarblock">
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
				verify(expected, source)
			})
		})
	})

	Context("file inclusions with line range", func() {

		Context("file inclusions as paragraph with line range", func() {

			It("should include single line as paragraph", func() {
				source := `include::includes/hello_world.go[lines=1]`
				expected := `<div class="paragraph">
<p>package includes</p>
</div>`
				verify(expected, source)
			})

			It("should include multiple lines as paragraph", func() {
				source := `include::includes/hello_world.go[lines=5..7]`
				expected := `<div class="paragraph">
<p>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</p>
</div>`
				verify(expected, source)
			})

			It("should include multiple ranges as paragraph", func() {
				source := `include::includes/hello_world.go[lines=1..2;5..7]`
				expected := `<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</p>
</div>`
				verify(expected, source)
			})
		})

		Context("file inclusions in listing blocks with line range", func() {

			It("should include single line in listing block", func() {
				source := `----
include::includes/hello_world.go[lines=1]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>package includes</pre>
</div>
</div>`
				verify(expected, source)
			})

			It("should include multiple lines in listing block", func() {
				source := `----
include::includes/hello_world.go[lines=5..7]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</pre>
</div>
</div>`
				verify(expected, source)
			})

			It("should include multiple ranges in listing block", func() {
				source := `----
include::includes/hello_world.go[lines=1..2;5..7]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>package includes

func helloworld() {
	fmt.Println(&#34;hello, world!&#34;)
}</pre>
</div>
</div>`
				verify(expected, source)
			})
		})
	})

	Context("recursive file inclusions", func() {

		It("should include child and grandchild content in paragraphs", func() {
			source := `include::includes/parent-include.adoc[]`
			expected := `<div class="paragraph">
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
			verify(expected, source)
		})

		It("should include child and grandchild content in listing block", func() {
			source := `----
include::includes/parent-include.adoc[]
----`
			expected := `<div class="listingblock">
<div class="content">
<pre>first line of parent

first line of child

first line of grandchild

last line of grandchild

last line of child

last line of parent</pre>
</div>
</div>`
			verify(expected, source)
		})
	})

	Context("inclusion with attribute in path", func() {

		It("should resolve path with attribute in standalone block", func() {
			source := `:includedir: ./includes
			
include::{includedir}/grandchild-include.adoc[]`
			expected := `<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>`
			verify(expected, source)
		})

		It("should resolve path with attribute in delimited block", func() {
			source := `:includedir: ./includes

----
include::{includedir}/grandchild-include.adoc[]
----`
			expected := `<div class="listingblock">
<div class="content">
<pre>first line of grandchild

last line of grandchild</pre>
</div>
</div>`
			verify(expected, source)
		})
	})

	Context("missing file to include", func() {

		Context("in standalone block", func() {

			It("should replace with string element if file is missing", func() {

				source := `include::includes/unknown.adoc[leveloffset=+1]`
				expected := `<div class="paragraph">
<p>Unresolved directive in test.adoc - include::includes/unknown.adoc[leveloffset=&#43;1]</p>
</div>`
				// TODO: also verify that an error was reported in the console.
				verify(expected, source)
			})

			It("should replace with string element if file with attribute in path is not resolved", func() {

				source := `include::{includedir}/unknown.adoc[leveloffset=+1]`
				expected := `<div class="paragraph">
<p>Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=&#43;1]</p>
</div>`
				// TODO: also verify that an error was reported in the console.
				verify(expected, source)
			})
		})

		Context("in listing block", func() {

			It("should replace with string element if file is missing", func() {
				source := `----
include::includes/unknown.adoc[leveloffset=+1]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>Unresolved directive in test.adoc - include::includes/unknown.adoc[leveloffset=+1]</pre>
</div>
</div>`
				// TODO: also verify that an error was reported in the console.
				verify(expected, source)
			})

			It("should replace with string element if file with attribute in path is not resolved", func() {
				source := `----
include::{includedir}/unknown.adoc[leveloffset=+1]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]</pre>
</div>
</div>`
				// TODO: also verify that an error was reported in the console.
				verify(expected, source)
			})
		})
	})
})
