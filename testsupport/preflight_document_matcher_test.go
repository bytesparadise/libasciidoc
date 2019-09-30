package testsupport_test

import (
	"fmt"

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
			matcher := testsupport.BecomePreflightDocument(expected)
			actual := "hello, world!"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.BecomePreflightDocument(expected)
			actual := "foo"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also check the error messages
			obtained := types.PreflightDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: actual,
								},
							},
						},
					},
				},
			}
			Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected preflight documents to match:\n%s", compare(obtained, expected))))
			Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected preflight documents not to match:\n%s", compare(obtained, expected))))
		})

		It("should return error when invalid type is input", func() {
			// given
			matcher := testsupport.BecomePreflightDocumentWithoutPreprocessing("")
			// when
			_, err := matcher.Match(1) // not a string
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("EqualDocumentBlock matcher expects a string (actual: int)"))
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
			matcher := testsupport.BecomePreflightDocumentWithoutPreprocessing(expected)
			// when
			result, err := matcher.Match("hello, world!")
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.BecomePreflightDocumentWithoutPreprocessing(expected)
			actual := "foo"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also check the error messages
			obtained := types.PreflightDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: actual,
								},
							},
						},
					},
				},
			}
			Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected preflight documents to match:\n%s", compare(obtained, expected))))
			Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected preflight documents not to match:\n%s", compare(obtained, expected))))
		})
	})

})
