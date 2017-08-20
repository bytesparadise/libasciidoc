package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering sections", func() {
	Context("Headings only", func() {

		It("heading level 1", func() {
			content := "= a title"
			// top-level heading is not rendered per-say,
			// but the heading will be used to set the HTML page's <title> element
			expected := ``
			verify(GinkgoT(), expected, content)
		})

		It("heading level 2", func() {
			content := "== a title"
			// top-level heading is not rendered per-say,
			// but the heading will be used to set the HTML page's <title> element
			expected := `<div class="sect1">
<h2 id="_a_title">a title</h2>
<div class="sectionbody">
</div>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("heading level 3", func() {
			content := "=== a title"
			// top-level heading is not rendered per-say,
			// but the heading will be used to set the HTML page's <title> element
			expected := `<div class="sect2">
<h3 id="_a_title">a title</h3>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("heading level 2 with just bold content", func() {
			content := `==  *2 spaces and bold content*`
			expected := `<div class="sect1">
<h2 id="__strong_2_spaces_and_bold_content_strong"><strong>2 spaces and bold content</strong></h2>
<div class="sectionbody">
</div>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("heading level 2 with just bold content", func() {
			content := `==  *2 spaces and bold content*`
			expected := `<div class="sect1">
<h2 id="__strong_2_spaces_and_bold_content_strong"><strong>2 spaces and bold content</strong></h2>
<div class="sectionbody">
</div>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("heading level 3 with nested bold content", func() {
			content := `=== a section title, with *bold content*`
			expected := `<div class="sect2">
<h3 id="_a_section_title_with_strong_bold_content_strong">a section title, with <strong>bold content</strong></h3>
</div>`
			verify(GinkgoT(), expected, content)
		})

		It("heading level 2 with custom ID", func() {
			content := `[#custom_id]
== a section title, with *bold content*`
			expected := `<div class="sect1">
<h2 id="custom_id">a section title, with <strong>bold content</strong></h2>
<div class="sectionbody">
</div>
</div>`
			verify(GinkgoT(), expected, content)
		})
	})

	Context("Section with elements", func() {

		It("heading level 2 with 2 paragraphs", func() {
			content := `== a title
		
and a first paragraph

and a second paragraph`
			// top-level heading is not rendered per-say,
			// but the heading will be used to set the HTML page's <title> element
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
			verify(GinkgoT(), expected, content)
		})
	})
})
