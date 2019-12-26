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

	It("should match", func() {
		// given
		actual := []interface{}{
			types.StringElement{
				Content: "foo@bar",
			},
		}
		matcher := testsupport.EqualWithoutNonAlphanumeric(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := []interface{}{
			types.StringElement{
				Content: "foobar",
			},
		}
		matcher := testsupport.EqualWithoutNonAlphanumeric(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		// also verify the messages
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected non-alphanumeric values to match:\n%s", compare("foobar", expected))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected non-alphanumeric values not to match:\n%s", compare("foobar", expected))))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.EqualWithoutNonAlphanumeric("")
		// when
		result, err := matcher.Match(1) // not a string
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("EqualWithoutNonAlphanumeric matcher expects an InlineElements (actual: int)"))
		Expect(result).To(BeFalse())
	})
})
