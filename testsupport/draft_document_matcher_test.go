package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("draft document assertions", func() {

	Context("with preprocessing", func() {

		expected := types.DraftDocument{
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
			matcher := testsupport.BecomeDraftDocument(expected)
			actual := "hello, world!"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.BecomeDraftDocument(expected)
			actual := "foo"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also check the error messages
			obtained := types.DraftDocument{
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
			Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected draft documents to match:\n%s", compare(obtained, expected))))
			Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected draft documents not to match:\n%s", compare(obtained, expected))))
		})

	})

	Context("without preprocessing", func() {

		expected := types.DraftDocument{
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
			matcher := testsupport.BecomeDraftDocument(expected, testsupport.WithoutPreprocessing())
			// when
			result, err := matcher.Match("hello, world!")
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.BecomeDraftDocument(expected, testsupport.WithoutPreprocessing())
			actual := "foo"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also check the error messages
			obtained := types.DraftDocument{
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
			Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected draft documents to match:\n%s", compare(obtained, expected))))
			Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected draft documents not to match:\n%s", compare(obtained, expected))))
		})
	})

})
