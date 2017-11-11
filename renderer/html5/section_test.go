package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Sections", func() {

	Context("Sections only", func() {

		It("header section", func() {
			actualContent := "= a title"
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := ``
			verify(GinkgoT(), expected, actualContent)
		})

		It("section level 1 alone", func() {
			actualContent := "== a title"
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div class="sect1">
<h2 id="_a_title">a title</h2>
<div class="sectionbody">
</div>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("section level 2 alone", func() {
			actualContent := "=== a title"
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div class="sect2">
<h3 id="_a_title">a title</h3>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("section level 1 with just bold content", func() {
			actualContent := `==  *2 spaces and bold content*`
			expected := `<div class="sect1">
<h2 id="__strong_2_spaces_and_bold_content_strong"><strong>2 spaces and bold content</strong></h2>
<div class="sectionbody">
</div>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("section level 2 with nested bold content", func() {
			actualContent := `=== a section title, with *bold content*`
			expected := `<div class="sect2">
<h3 id="_a_section_title_with_strong_bold_content_strong">a section title, with <strong>bold content</strong></h3>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("section level 1 with custom ID", func() {
			actualContent := `[#custom_id]
== a section title, with *bold content*`
			expected := `<div class="sect1">
<h2 id="custom_id">a section title, with <strong>bold content</strong></h2>
<div class="sectionbody">
</div>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})

	Context("Section with elements", func() {

		It("section level 1 with 2 paragraphs", func() {
			actualContent := `== a title
		
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
			verify(GinkgoT(), expected, actualContent)
		})

		It("section with just a paragraph", func() {
			actualContent := `= a title
		
a paragraph`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div class="paragraph">
<p>a paragraph</p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})

		It("header with preamble then section level 1", func() {
			actualContent := `= a title
		
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
			verify(GinkgoT(), expected, actualContent)
		})

		It("header with preamble then 2 sections level 1", func() {
			actualContent := `= a title
		
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
			verify(GinkgoT(), expected, actualContent)
		})

	})
})
