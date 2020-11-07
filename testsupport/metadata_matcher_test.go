package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("html5 body renderer", func() {

	It("should match when title exists", func() {
		// given
		actual := `= hello, world!`
		expected := `hello, world!`
		// when
		result, err := testsupport.MetadataTitle(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expected))
	})

	It("should match when title does not exist", func() {
		// given
		actual := `foo` // no title in this doc
		// when
		result, err := testsupport.MetadataTitle(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(""))
	})
})
