package html5_test

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
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("blank line with spaces and tabs between 2 paragraphs", func() {
		source := `first paragraph
		  
second paragraph`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>
<div class="paragraph">
<p>second paragraph</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("blank lines (tabs) at end of document", func() {
		source := `first paragraph
		
		
		`
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("blank lines (spaces) at end of document", func() {
		source := `first paragraph
		
		
        `
		expected := `<div class="paragraph">
<p>first paragraph</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})
