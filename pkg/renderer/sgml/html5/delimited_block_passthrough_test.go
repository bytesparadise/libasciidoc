package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("passthrough blocks", func() {

	Context("delimited blocks", func() {

		It("with title", func() {
			source := `.a title
++++
_foo_

*bar*
++++`
			expected := `_foo_

*bar*
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with special characters", func() {
			source := `++++
<input>

<input>
++++`
			expected := `<input>

<input>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("paragraph blocks", func() {

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
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
