package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("html5 body renderer", func() {

	It("should match", func() {
		// given
		actual := "hello, world!"
		// when
		result, err := testsupport.RenderHTML(actual)
		// then
		expected := `<div class="paragraph">
<p>hello, world!</p>
</div>
`
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(expected))
	})
})
