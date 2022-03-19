package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golintt
)

var _ = Describe("inline menus", func() {

	Context("in final documents", func() {

		It("with main path", func() {
			source := `:experimental:
 
Select menu:File[].`
			expected := `<div class="paragraph">
<p>Select <b class="menuref">File</b>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with single sub path", func() {
			source := `:experimental:
 
Select menu:File[Save].`
			expected := `<div class="paragraph">
<p>Select <span class="menuseq"><b class="menu">File</b>&#160;<b class="caret">&#8250;</b> <b class="menuitem">Save</b></span>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with multiple sub paths", func() {
			source := `:experimental:
 
Select menu:File[Zoom > Reset].`
			expected := `<div class="paragraph">
<p>Select <span class="menuseq"><b class="menu">File</b>&#160;<b class="caret">&#8250;</b> <b class="submenu">Zoom</b>&#160;<b class="caret">&#8250;</b> <b class="menuitem">Reset</b></span>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("when experimental is not enabled", func() {
			source := `Select menu:File[Zoom > Reset].`
			expected := `<div class="paragraph">
<p>Select menu:File[Zoom &gt; Reset].</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
