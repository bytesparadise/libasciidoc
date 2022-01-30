package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("thenatic breaks", func() {

	It("simple", func() {
		source := `before

'''

after`
		expected := `<div class="paragraph">
<p>before</p>
</div>
<hr>
<div class="paragraph">
<p>after</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})
