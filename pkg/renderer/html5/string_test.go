package html5_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("strings", func() {

	Context("ellipsis conversion", func() {

		It("text with ellipsis", func() {
			source := `some text...`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expected := `<div class="paragraph">
<p>some text&#8230;&#8203;</p>
</div>`
			verify(expected, source)
		})
	})
})
