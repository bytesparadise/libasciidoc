package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("raw document assertions", func() {

	expected := types.RawDocument{
		Elements: []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "hello, world!",
						},
					},
				},
			},
		},
	}

	It("should match", func() {
		// given
		actual := "hello, world!"
		// when
		result, err := testsupport.ParseRawDocument(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expected))
	})

	It("should not match", func() {
		// given
		actual := "foo"
		// when
		result, err := testsupport.ParseRawDocument(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).NotTo(Equal(expected))
	})
})
