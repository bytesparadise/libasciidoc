package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering Paragraph", func() {
	It("some paragraph", func() {

		content := "*bold content* \n" +
			"with more content afterwards..."
		expected := `<div class="paragraph">
<p><strong>bold content</strong> with more content afterwards...</p>
</div>`
		verify(GinkgoT(), expected, content)
	})
})
