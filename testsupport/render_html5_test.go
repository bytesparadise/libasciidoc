package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
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
</div>`
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(expected))
	})

})

var _ = Describe("html5 body renderer", func() {

	It("should match when title exists", func() {
		// given
		actual := `= hello, world!`
		expected := `hello, world!`
		// when
		result, err := testsupport.RenderHTML5Title(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expected))
	})

	It("should match title does not exist", func() {
		// given
		actual := `foo` // no title in this doc
		// when
		result, err := testsupport.RenderHTML5Title(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(""))
	})

})
