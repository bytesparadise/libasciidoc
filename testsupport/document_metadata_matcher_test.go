package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document metadata assertions", func() {

	expected := types.DocumentAttributes{
		"foo": "bar",
	}

	It("should match", func() {
		// given
		actual := types.Document{
			Attributes: types.DocumentAttributes{
				"foo": "bar",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Attributes: types.ElementAttributes{},
					Title:      types.InlineElements{},
					Elements:   []interface{}{},
				},
			},
		}
		matcher := testsupport.HaveMetadata(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := types.Document{}
		matcher := testsupport.HaveMetadata(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		// also verify the messages
		obtained := types.DocumentAttributes{}
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected document metadata to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected document metadata not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.HaveMetadata(expected)
		// when
		result, err := matcher.Match("foo")
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("HaveMetadata matcher expects a Document (actual: string)"))
		Expect(result).To(BeFalse())
	})

})
