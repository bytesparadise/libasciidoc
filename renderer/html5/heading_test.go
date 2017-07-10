package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("render headings", func() {
	It("heading level 1", func() {
		content := "= a title"
		expected := `<div id="header">
<h1>a title</h1>
</div>`
		verify(GinkgoT(), expected, content)
	})
	It("heading with just bold content", func() {
		content := `==  *2 spaces and bold content*`
		expected := `<h2 id="__strong_2_spaces_and_bold_content_strong"><strong>2 spaces and bold content</strong></h2>`
		verify(GinkgoT(), expected, content)
	})
	It("heading with nested bold content", func() {
		content := `== a section title, with *bold content*`
		expected := `<h2 id="_a_section_title_with_strong_bold_content_strong">a section title, with <strong>bold content</strong></h2>`
		verify(GinkgoT(), expected, content)
	})
	It("heading with custom ID", func() {
		content := `[#custom_id]
== a section title, with *bold content*`
		expected := `<h2 id="custom_id">a section title, with <strong>bold content</strong></h2>`
		verify(GinkgoT(), expected, content)
	})
})
