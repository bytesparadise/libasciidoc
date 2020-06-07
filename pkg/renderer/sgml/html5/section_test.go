package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("sections", func() {

	Context("sections only", func() {

		It("header section", func() {
			source := "=   a title  "
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := ``
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("section level 1 alone", func() {
			source := "== a title with *bold* content"
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div class="sect1">
<h2 id="_a_title_with_bold_content">a title with <strong>bold</strong> content</h2>
<div class="sectionbody">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("section level 2 alone", func() {
			source := "=== a title"
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div class="sect2">
<h3 id="_a_title">a title</h3>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("section level 1 with just bold content", func() {
			source := `==  *2 spaces and bold content*`
			expected := `<div class="sect1">
<h2 id="_2_spaces_and_bold_content"><strong>2 spaces and bold content</strong></h2>
<div class="sectionbody">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("section level 2 with nested bold content", func() {
			source := `=== a section title, with *bold content*`
			expected := `<div class="sect2">
<h3 id="_a_section_title_with_bold_content">a section title, with <strong>bold content</strong></h3>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("section level 1 with custom ID", func() {
			source := `
:idprefix: ignored_
			
[#custom_id]
== a section title, with *bold content*`
			expected := `<div class="sect1">
<h2 id="custom_id">a section title, with <strong>bold content</strong></h2>
<div class="sectionbody">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("section level 1 with custom prefix id", func() {
			source := `
:idprefix: id_

== a section title`
			expected := `<div class="sect1">
<h2 id="id_a_section_title">a section title</h2>
<div class="sectionbody">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("sections with same title", func() {
			source := `== section 1

== section 1`
			expected := `<div class="sect1">
<h2 id="_section_1">section 1</h2>
<div class="sectionbody">
</div>
</div>
<div class="sect1">
<h2 id="_section_1_2">section 1</h2>
<div class="sectionbody">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("section with elements", func() {

		It("section level 1 with 2 paragraphs", func() {
			source := `== a title
		
and a first paragraph

and a second paragraph`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div class="sect1">
<h2 id="_a_title">a title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>and a first paragraph</p>
</div>
<div class="paragraph">
<p>and a second paragraph</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("section with just a paragraph", func() {
			source := `= a title
		
a paragraph`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div class="paragraph">
<p>a paragraph</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("header with preamble then section level 1", func() {
			source := `= a title
		
a preamble

splitted in 2 paragraphs

== section 1

with some text`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>a preamble</p>
</div>
<div class="paragraph">
<p>splitted in 2 paragraphs</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_1">section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some text</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("header with preamble then 2 sections level 1", func() {
			source := `= a title
		
a preamble

splitted in 2 paragraphs

== section 1

with some text

== section 2

with some text, too`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>a preamble</p>
</div>
<div class="paragraph">
<p>splitted in 2 paragraphs</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_1">section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some text</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_2">section 2</h2>
<div class="sectionbody">
<div class="paragraph">
<p>with some text, too</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("section with listing block and subsection", func() {
			source := `==== Third level heading

[#id-for-listing-block]
.Listing block title
----
Content in a listing block is subject to verbatim substitutions.
Listing block content is commonly used to preserve code input.
----

===== Fourth level heading
foo`

			expected := `<div class="sect3">
<h4 id="_third_level_heading">Third level heading</h4>
<div id="id-for-listing-block" class="listingblock">
<div class="title">Listing block title</div>
<div class="content">
<pre>Content in a listing block is subject to verbatim substitutions.
Listing block content is commonly used to preserve code input.</pre>
</div>
</div>
<div class="sect4">
<h5 id="_fourth_level_heading">Fourth level heading</h5>
<div class="paragraph">
<p>foo</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("preambles", func() {

		It("should include preamble wrapper", func() {
			source := `= Title

preamble 
here

== section 1

content here`
			expected := `<div id="preamble">
<div class="sectionbody">
<div class="paragraph">
<p>preamble
here</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_1">section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content here</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("should not include preamble wrapper", func() {
			source := `preamble 
here

== section 1

content here
`
			expected := `<div class="paragraph">
<p>preamble
here</p>
</div>
<div class="sect1">
<h2 id="_section_1">section 1</h2>
<div class="sectionbody">
<div class="paragraph">
<p>content here</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
