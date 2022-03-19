package xhtml5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("passthrough blocks", func() {

	Context("as delimited blocks", func() {

		It("with title", func() {
			source := `.a title
++++
_foo_

*bar*
++++`
			expected := `_foo_

*bar*
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

		It("with special characters", func() {
			source := `++++
<input>

<input>
++++`
			expected := `<input>

<input>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("open block", func() {

		It("2-line paragraph followed by another paragraph", func() {
			source := `[pass]
_foo_
*bar*

another paragraph`
			expected := `_foo_
*bar*
<div class="paragraph">
<p>another paragraph</p>
</div>
`
			Expect(RenderXHTML(source)).To(MatchHTML(expected))
		})
	})
})
