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

	Context("messages", func() {

		It("failure message", func() {
			// given
			matcher := testsupport.EqualPreflightDocumentWithoutPreprocessing("foo")
			// when
			msg := matcher.FailureMessage("bar")
			// then
			Expect(msg).To(Equal(fmt.Sprintf("expected preflight documents to match:\n\texpected: '%v'\n\tactual'%v'", "foo", "bar")))
		})

		It("negated failure message", func() {
			// given
			matcher := testsupport.EqualPreflightDocumentWithoutPreprocessing("foo")
			// when
			msg := matcher.NegatedFailureMessage("bar")
			// then
			Expect(msg).To(Equal(fmt.Sprintf("expected preflight documents not to match:\n\texpected: '%v'\n\tactual'%v'", "foo", "bar")))

		})
	})

	Context("failures", func() {

		It("should return error when invalid type is input", func() {
			// given
			matcher := testsupport.EqualPreflightDocumentWithoutPreprocessing("")
			// when
			_, err := matcher.Match(1) // not a string
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("EqualDocumentBlock matcher expects a string (actual: int)"))
		})
	})

})