package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golintt
)

var _ = Describe("buttons", func() {

	Context("in final documents", func() {

		It("when experimental is enabled", func() {
			source := `:experimental:
 
Click on btn:[OK].`
			expected := `<div class="paragraph">
<p>Click on <b class="button">OK</b>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("when experimental is not enabled", func() {
			source := `Click on btn:[OK].`
			expected := `<div class="paragraph">
<p>Click on btn:[OK].</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
