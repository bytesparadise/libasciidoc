package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document preamble assertions", func() {

	expected := types.Document{
		Attributes:         types.DocumentAttributes{},
		ElementReferences:  types.ElementReferences{},
		Footnotes:          types.Footnotes{},
		FootnoteReferences: types.FootnoteReferences{},
		Elements: []interface{}{
			types.Preamble{
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a short paragraph"},
							},
						},
					},
					types.BlankLine{},
				},
			},
			types.Section{
				Level: 1,
				Title: types.InlineElements{
					types.StringElement{Content: "Section A"},
				},
				Attributes: types.ElementAttributes{},
				Elements:   []interface{}{},
			},
			types.Section{
				Level: 1,
				Title: types.InlineElements{
					types.StringElement{Content: "Section B"},
				},
				Attributes: types.ElementAttributes{},
				Elements:   []interface{}{},
			},
		},
	}

	It("should match", func() {
		// given
		matcher := testsupport.HavePreamble(expected)
		actual := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Level: 1,
					Title: types.InlineElements{
						types.StringElement{Content: "Section A"},
					},
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				},
				types.Section{
					Level: 1,
					Title: types.InlineElements{
						types.StringElement{Content: "Section B"},
					},
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				},
			},
		}
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		matcher := testsupport.HavePreamble(expected)
		actual := types.Document{}
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		// also verify messages
		obtained := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements:           []interface{}{},
		}
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected document to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected document not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.HavePreamble(types.Document{})
		// when
		result, err := matcher.Match(1) // not a doc
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("HavePreamble matcher expects a Document (actual: int)"))
		Expect(result).To(BeFalse())
	})
})
