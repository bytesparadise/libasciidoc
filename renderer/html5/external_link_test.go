package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Links", func() {
	Context("External links", func() {

		It("External link alone", func() {

			actualContent := "https://foo.com[the website]"
			expected := `<div class="paragraph">
<p><a href="https://foo.com">the website</a></p>
</div>`
			verify(GinkgoT(), expected, actualContent)
		})
	})
})
