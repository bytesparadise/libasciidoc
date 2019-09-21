package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("non-alphanumeric replacement assertions", func() {

	expected := "foo_bar"
	actual := types.InlineElements{
		types.StringElement{
			Content: "foo@bar",
		},
	}

	It("should match", func() {
		// given
		matcher := testsupport.EqualWithoutNonAlphanumeric(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		matcher := testsupport.EqualWithoutNonAlphanumeric(expected)
		// when
		result, err := matcher.Match(types.InlineElements{
			types.StringElement{
				Content: "foobar",
			},
		})
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
	})

	Context("messages", func() {

		It("failure", func() {
			// given
			matcher := testsupport.EqualWithoutNonAlphanumeric(expected)
			_, err := matcher.Match(actual)
			Expect(err).ToNot(HaveOccurred())
			// when
			msg := matcher.FailureMessage(actual)
			// then
			Expect(msg).To(Equal(fmt.Sprintf("expected non alphanumeric values to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, expected)))
		})

		It("negated failure message", func() {
			// given
			matcher := testsupport.EqualWithoutNonAlphanumeric(expected)
			_, err := matcher.Match(actual)
			Expect(err).ToNot(HaveOccurred())
			// when
			msg := matcher.NegatedFailureMessage(actual)
			// then
			Expect(msg).To(Equal(fmt.Sprintf("expected non alphanumeric values not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, expected)))

		})
	})

	Context("failures", func() {

		It("should return error when invalid type is input", func() {
			// given
			matcher := testsupport.EqualWithoutNonAlphanumeric("")
			_, err := matcher.Match(actual)
			Expect(err).ToNot(HaveOccurred())
			// when
			result, err := matcher.Match(1) // not a string
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("EqualWithoutNonAlphanumeric matcher expects an InlineElements (actual: int)"))
			Expect(result).To(BeFalse())
		})
	})
})
