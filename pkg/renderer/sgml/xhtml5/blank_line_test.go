package xhtml5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("blank lines", func() {

	It("blank line between 2 paragraphs", func() {
		source := `first paragraph 

second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("blank line with spaces and tabs between 2 paragraphs", func() {
		source := `first paragraph
		  
second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("blank lines (tabs) at end of document", func() {
		source := `first paragraph
		
		
		`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("blank lines (spaces) at end of document", func() {
		source := `first paragraph
		
		
        `
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("hard break", func() {
		source := `first line +
second line`
		expected := `<div class="paragraph">
<p>first line<br/>
second line</p>
</div>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("thematic break", func() {
		source := "- - -"
		expected := "<hr/>\n"
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})
})
