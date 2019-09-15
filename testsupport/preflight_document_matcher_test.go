package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("preflight document assertions", func() {

	Context("with preprocessing", func() {

		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
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
			matcher := testsupport.EqualPreflightDocument(expected)
			// when
			result, err := matcher.Match("hello, world!")
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.EqualPreflightDocument(expected)
			// when
			result, err := matcher.Match("meh")
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
		})
	})

	Context("without preprocessing", func() {

		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
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
			matcher := testsupport.EqualPreflightDocumentWithoutPreprocessing(expected)
			// when
			result, err := matcher.Match("hello, world!")
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.EqualPreflightDocumentWithoutPreprocessing(expected)
			// when
			result, err := matcher.Match("meh")
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
		})
	})

})
