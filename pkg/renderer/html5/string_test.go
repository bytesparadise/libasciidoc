package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("strings", func() {

	It("text with ellipsis", func() {
		source := `some text...`
		// top-level section is not rendered per-say,
		// but the section will be used to set the HTML page's <title> element
		expected := `<div class="paragraph">
<p>some text&#8230;&#8203;</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("text with copyright", func() {
		source := `Copyright (C)`
		// top-level section is not rendered per-say,
		// but the section will be used to set the HTML page's <title> element
		expected := `<div class="paragraph">
<p>Copyright &#169;</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})
