package testsupport_test

import (
	"context"
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
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
		ctx := renderer.Wrap(context.Background(), actual)
		renderer.IncludePreamble(ctx)
		obtained := ctx.Document
		GinkgoT().Logf(matcher.FailureMessage(actual))
		GinkgoT().Logf(fmt.Sprintf("expected documents to match:\n%s", compare(obtained, expected)))
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected documents to match:\n%s", compare(obtained, expected))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected documents not to match:\n%s", compare(obtained, expected))))
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
