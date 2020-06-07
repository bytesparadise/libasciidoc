package html5_test

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
	log "github.com/sirupsen/logrus"
)

var _ = Describe("file inclusions", func() {

	It("should include adoc file without leveloffset from local file", func() {
		console, reset := ConfigureLogger()
		defer reset()
		lastUpdated := time.Now()
		source := "include::../../../../test/includes/grandchild-include.adoc[]"
		expected := `<div class="sect1">
<h2 id="_grandchild_title">grandchild title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
</div>
</div>`
		Expect(RenderHTML(source, configuration.WithFilename("test.adoc"), configuration.WithLastUpdated(lastUpdated))).To(Equal(expected))
		Expect(DocumentMetadata(source, lastUpdated)).To(Equal(types.Metadata{
			Title:       "",
			LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
			TableOfContents: types.TableOfContents{
				Sections: []types.ToCSection{
					{
						ID:       "_grandchild_title",
						Level:    1,
						Title:    "grandchild title",
						Children: []types.ToCSection{},
					},
				},
			},
		}))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file without leveloffset from relative file", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../../../../test/includes/grandchild-include.adoc[]"
		expected := `<div class="sect1">
<h2 id="_grandchild_title">grandchild title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
</div>
</div>`
		Expect(RenderHTML(source, configuration.WithFilename("tmp/foo.adoc"))).To(Equal(expected))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include grandchild content with relative offset", func() {
		source := `include::../../../../test/includes/grandchild-include.adoc[leveloffset=+1]`
		expected := `<div class="sect2">
<h3 id="_grandchild_title">grandchild title</h3>
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
</div>`
		Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(Equal(expected))
	})

	It("should include grandchild content with absolute offset", func() {
		source := `include::../../../../test/includes/grandchild-include.adoc[leveloffset=1]`
		expected := `<div class="sect1">
<h2 id="_grandchild_title">grandchild title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
</div>
</div>`
		Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(Equal(expected))
	})

	It("should include child and grandchild content with relative level offset", func() {
		source := `include::../../../../test/includes/parent-include-relative-offset.adoc[leveloffset=+1]`
		expected := `<div class="sect1">
<h2 id="_parent_title">parent title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of parent</p>
</div>
<div class="paragraph">
<p>child preamble</p>
</div>
<div class="sect3">
<h4 id="_child_section_1">child section 1</h4>
<div class="paragraph">
<p>first line of child</p>
</div>
<div class="sect4">
<h5 id="_grandchild_title">grandchild title</h5>
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
</div>
<div class="sect4">
<h5 id="_child_section_2">child section 2</h5>
<div class="paragraph">
<p>last line of child</p>
</div>
<div class="paragraph">
<p>last line of parent</p>
</div>
</div>
</div>
</div>
</div>`
		Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(Equal(expected))
	})

	It("should include child and grandchild content with relative then absolute level offset", func() {
		source := `include::../../../../test/includes/parent-include-absolute-offset.adoc[leveloffset=+1]`
		expected := `<div class="sect1">
<h2 id="_parent_title">parent title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of parent</p>
</div>
<div class="paragraph">
<p>child preamble</p>
</div>
<div class="sect3">
<h4 id="_child_section_1">child section 1</h4>
<div class="paragraph">
<p>first line of child</p>
</div>
<div class="sect4">
<h5 id="_grandchild_title">grandchild title</h5>
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
</div>
<div class="sect4">
<h5 id="_child_section_2">child section 2</h5>
<div class="paragraph">
<p>last line of child</p>
</div>
<div class="paragraph">
<p>last line of parent</p>
</div>
</div>
</div>
</div>
</div>`
		Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(Equal(expected))
	})

	It("include adoc file with leveloffset attribute", func() {
		source := `= Master Document

preamble

include::../../../../test/includes/chapter-a.adoc[leveloffset=+1]`
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("should not include section 0 by default", func() {
		source := `include::../../../../test/includes/chapter-a.adoc[]`
		expected := `<div class="paragraph">
<p>content</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("should not include section 0 when attribute found", func() {
		source := `:includedir: ../../../../test/includes

include::{includedir}/chapter-a.adoc[]`
		expected := `<div class="paragraph">
<p>content</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("include non adoc file", func() {
		source := `= Master Document

preamble

include::../../../../test/includes/hello_world.go.txt[]`
		expected := `<div class="paragraph">
<p>preamble</p>
</div>
<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>import "fmt"</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println("hello, world!")
}</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("should not further process with non-asciidoc files", func() {
		source := `:includedir: ../../../../test/includes

include::{includedir}/include.foo[]`
		expected := `<div class="paragraph">
<p><strong>some strong content</strong></p>
</div>
<div class="paragraph">
<p>include::hello_world.go.txt[]</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("include 2 files", func() {
		source := `= Master Document

preamble

include::../../../../test/includes/grandchild-include.adoc[]

include::../../../../test/includes/hello_world.go.txt[]`
		expected := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>preamble</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_grandchild_title">grandchild title</h2>
<div class="sectionbody">
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
<p>import "fmt"</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println("hello, world!")
}</p>
</div>
</div>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("include file and append following elements in included section", func() {
		source := `a first paragraph

include::../../../../test/includes/chapter-a.adoc[leveloffset=+1]

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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	Context("in delimited blocks", func() {

		Context("adoc file inclusion in delimited blocks", func() {

			It("should include adoc file within listing block", func() {
				source := `= Master Document

preamble

----
include::../../../../test/includes/chapter-a.adoc[]
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
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include adoc file within fenced block", func() {
				source := "```\n" +
					"include::../../../../test/includes/chapter-a.adoc[]\n" +
					"```"
				expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>= Chapter A

content</code></pre>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include adoc file within example block", func() {
				source := `====
include::../../../../test/includes/chapter-a.adoc[]
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
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include adoc file within quote block", func() {
				source := `____
include::../../../../test/includes/chapter-a.adoc[]
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
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include adoc file within verse block", func() {
				source := `[verse]
____
include::../../../../test/includes/chapter-a.adoc[]
____`
				expected := `<div class="verseblock">
<pre class="content">= Chapter A

content</pre>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include adoc file within sidebar block", func() {
				source := `****
include::../../../../test/includes/chapter-a.adoc[]
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
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include adoc file within passthrough block", func() {
				Skip("missing support for passthrough blocks")
				source := `++++
include::../../../../test/includes/chapter-a.adoc[]
++++`
				expected := ``
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})

		Context("other file inclusion in delimited blocks", func() {

			It("should include go file within listing block", func() {
				source := `= Master Document

preamble

----
include::../../../../test/includes/hello_world.go.txt[]
----`
				expected := `<div class="paragraph">
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
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include go file within fenced block", func() {
				source := "```\n" +
					"include::../../../../test/includes/hello_world.go.txt[]\n" +
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
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include go file within example block", func() {
				source := `====
include::../../../../test/includes/hello_world.go.txt[]
====`
				expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>import "fmt"</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println("hello, world!")
}</p>
</div>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include go file within quote block", func() {
				source := `____
include::../../../../test/includes/hello_world.go.txt[]
____`
				expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>import "fmt"</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println("hello, world!")
}</p>
</div>
</blockquote>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include go file within verse block", func() {
				source := `[verse]
____
include::../../../../test/includes/hello_world.go.txt[]
____`
				expected := `<div class="verseblock">
<pre class="content">package includes

import "fmt"

func helloworld() {
	fmt.Println("hello, world!")
}</pre>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include go file within sidebar block", func() {
				source := `****
include::../../../../test/includes/hello_world.go.txt[]
****`
				expected := `<div class="sidebarblock">
<div class="content">
<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>import "fmt"</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println("hello, world!")
}</p>
</div>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("file inclusions with line range", func() {

		Context("file inclusions as paragraph with line range", func() {

			It("should include single line as paragraph", func() {
				source := `include::../../../../test/includes/hello_world.go.txt[lines=1]`
				expected := `<div class="paragraph">
<p>package includes</p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include multiple lines as paragraph", func() {
				source := `include::../../../../test/includes/hello_world.go.txt[lines=5..7]`
				expected := `<div class="paragraph">
<p>func helloworld() {
	fmt.Println("hello, world!")
}</p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include multiple ranges as paragraph", func() {
				source := `include::../../../../test/includes/hello_world.go.txt[lines=1..2;5..7]`
				expected := `<div class="paragraph">
<p>package includes</p>
</div>
<div class="paragraph">
<p>func helloworld() {
	fmt.Println("hello, world!")
}</p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})

		Context("file inclusions in listing blocks with line range", func() {

			It("should include single line in listing block", func() {
				source := `----
include::../../../../test/includes/hello_world.go.txt[lines=1]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>package includes</pre>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include multiple lines in listing block", func() {
				source := `----
include::../../../../test/includes/hello_world.go.txt[lines=5..7]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>func helloworld() {
	fmt.Println("hello, world!")
}</pre>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should include multiple ranges in listing block", func() {
				source := `----
include::../../../../test/includes/hello_world.go.txt[lines=1..2;5..7]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>package includes

func helloworld() {
	fmt.Println("hello, world!")
}</pre>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("file inclusions with tag ranges", func() {

		It("file inclusion with single tag", func() {
			source := `include::../../../../test/includes/tag-include.adoc[tag=section]`
			expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("file inclusion with surrounding tag", func() {
			source := `include::../../../../test/includes/tag-include.adoc[tag=doc]`
			expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("file inclusion with unclosed tag", func() {
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../../../test/includes/tag-include-unclosed.adoc[tag=unclosed]`
			expected := `<div class="paragraph">
<p>content</p>
</div>
<div class="paragraph">
<p>end</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
			// verify error in logs
			Expect(console).To(
				ContainMessageWithLevel(
					log.ErrorLevel,
					"detected unclosed tag 'unclosed' starting at line 6 of include file: ../../../../test/includes/tag-include-unclosed.adoc",
				))
		})

		It("file inclusion with unknown tag", func() {
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../../../test/includes/tag-include.adoc[tag=unknown]`
			expected := ``
			Expect(RenderHTML(source)).To(MatchHTML(expected))
			// verify error in logs
			Expect(console).To(
				ContainMessageWithLevel(
					log.ErrorLevel,
					"tag 'unknown' not found in include file: ../../../../test/includes/tag-include.adoc",
				))
		})

		It("file inclusion with no tag", func() {
			source := `include::../../../../test/includes/tag-include.adoc[]`
			expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content</p>
</div>
<div class="paragraph">
<p>end</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		Context("permutations", func() {

			It("all lines", func() {
				source := `include::../../../../test/includes/tag-include.adoc[tag=**]` // includes all content except lines with tags
				expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content</p>
</div>
<div class="paragraph">
<p>end</p>
</div>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("all tagged regions", func() {
				source := `include::../../../../test/includes/tag-include.adoc[tag=*]` // includes all sections
				expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content</p>
</div>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("all the lines outside and inside of tagged regions", func() {
				source := `include::../../../../test/includes/tag-include.adoc[tag=**;*]` // includes all sections
				expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content</p>
</div>
<div class="paragraph">
<p>end</p>
</div>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("regions tagged doc, but not nested regions tagged content", func() {
				source := `include::../../../../test/includes/tag-include.adoc[tag=doc;!content]` // includes all sections
				expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("all tagged regions, but excludes any regions tagged content", func() {
				source := `include::../../../../test/includes/tag-include.adoc[tag=*;!content]` // includes all sections
				expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("all tagged regions, but excludes any regions tagged content", func() {
				source := `include::../../../../test/includes/tag-include.adoc[tag=**;!content]` // includes all sections
				expected := `<div class="sect1">
<h2 id="_section_1">Section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>end</p>
</div>
</div>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("**;!* — selects only the regions of the document outside of tags", func() {
				source := `include::../../../../test/includes/tag-include.adoc[tag=**;!*]` // includes all sections
				expected := `<div class="paragraph">
<p>end</p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("recursive file inclusions", func() {

		It("should include child and grandchild content in paragraphs", func() {
			source := `include::../../../../test/includes/parent-include.adoc[leveloffset=+1]`
			expected := `<div class="sect1">
<h2 id="_parent_title">parent title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of parent</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_child_title">child title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of child</p>
</div>
<div class="sect2">
<h3 id="_grandchild_title">grandchild title</h3>
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
</div>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("should include child and grandchild content in listing block", func() {
			source := `----
include::../../../../test/includes/parent-include.adoc[leveloffset=+1]
----`
			expected := `<div class="listingblock">
<div class="content">
<pre>= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("inclusion with attribute in path", func() {

		It("should resolve path with attribute in standalone block", func() {
			source := `:includedir: ../../../../test/includes
			
include::{includedir}/grandchild-include.adoc[]`
			expected := `<div class="sect1">
<h2 id="_grandchild_title">grandchild title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("should resolve path with attribute in delimited block", func() {
			source := `:includedir: ../../../../test/includes

----
include::{includedir}/grandchild-include.adoc[]
----`
			expected := `<div class="listingblock">
<div class="content">
<pre>== grandchild title

first line of grandchild

last line of grandchild</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("missing file to include", func() {

		Context("in standalone block", func() {

			It("should replace with error message if file is missing", func() {
				// setup logger to write in a buffer so we can check the output
				console, reset := ConfigureLogger()
				defer reset()

				source := `include::../../../test/includes/unknown.adoc[leveloffset=+1]`
				expected := `<div class="paragraph">
<p>Unresolved directive in test.adoc - include::../../../test/includes/unknown.adoc[leveloffset=+1]</p>
</div>`
				Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(Equal(expected))
				// verify error in logs
				Expect(console).To(ContainMessageWithLevel(log.ErrorLevel, "failed to include content of '../../../test/includes/unknown.adoc' in 'test.adoc'"))
			})

			It("should replace with error message if file with attribute in path is not resolved", func() {
				// setup logger to write in a buffer so we can check the output
				console, reset := ConfigureLogger()
				defer reset()

				source := `include::{includedir}/unknown.adoc[leveloffset=+1]`
				expected := `<div class="paragraph">
<p>Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]</p>
</div>`
				Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(Equal(expected))
				// verify error in logs
				Expect(console).To(ContainMessageWithLevel(log.ErrorLevel, "failed to include content of '{includedir}/unknown.adoc' in 'test.adoc'"))
			})
		})

		Context("in listing block", func() {

			It("should replace with error message if file is missing", func() {
				// setup logger to write in a buffer so we can check the output
				console, reset := ConfigureLogger()
				defer reset()

				source := `----
include::../../../test/includes/unknown.adoc[leveloffset=+1]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>Unresolved directive in test.adoc - include::../../../test/includes/unknown.adoc[leveloffset=+1]</pre>
</div>
</div>`
				Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(MatchHTML(expected))
				// verify error in logs
				Expect(console).To(ContainMessageWithLevel(log.ErrorLevel, "failed to include content of '../../../test/includes/unknown.adoc' in 'test.adoc'"))
			})

			It("should replace with error message if file with attribute in path is not resolved", func() {
				// setup logger to write in a buffer so we can check the output
				console, reset := ConfigureLogger()
				defer reset()

				source := `----
include::{includedir}/unknown.adoc[leveloffset=+1]
----`
				expected := `<div class="listingblock">
<div class="content">
<pre>Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]</pre>
</div>
</div>`
				Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(Equal(expected))
				// verify error in logs
				Expect(console).To(ContainMessageWithLevel(log.ErrorLevel, "failed to include content of '{includedir}/unknown.adoc' in 'test.adoc'"))
			})
		})
	})
})
