package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("parse document block", func() {

	expected := types.Paragraph{
		Lines: [][]interface{}{
			{
				types.StringElement{ // `ParseDocumentBlock` uses the `AsciidocRawDocument` grammar rule
					Content: "hello, world!",
				},
			},
		},
	}

	It("should match", func() {
		// given
		actual := "hello, world!"
		// when
		result, err := testsupport.ParseDocumentBlock(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expected))
	})

	It("should not match", func() {
		// given
		actual := "foo"
		// when
		result, err := testsupport.ParseDocumentBlock(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).NotTo(Equal(expected))
	})

})
