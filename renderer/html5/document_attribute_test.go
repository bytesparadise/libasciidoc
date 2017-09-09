package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering With Attributes", func() {
	It("some attributes then a paragraph", func() {
		content := `:toc:
:date: 2017-01-01
:author: Xavier
a paragraph`

		expected := `<div class="paragraph">
<p>a paragraph</p>
</div>`
		verify(GinkgoT(), expected, content)
	})
})

var _ = Describe("Rendering With Attributes", func() {
	It("a paragraph then some attributes", func() {
		content := `a paragraph

:toc:
:date: 2017-01-01
:author: Xavier`

		expected := `<div class="paragraph">
<p>a paragraph</p>
</div>`
		verify(GinkgoT(), expected, content)
	})
})
