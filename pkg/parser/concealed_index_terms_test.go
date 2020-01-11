package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("concealed index terms", func() {

	Context("draft document", func() {

		It("index term in existing paragraph line", func() {
			source := `a paragaph with an index term (((index, term, here))).`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "a paragaph with an index term ",
								},
								types.ConceleadIndexTerm{
									Term1: "index",
									Term2: "term",
									Term3: "here",
								},
								types.StringElement{
									Content: ".",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("index term in single paragraph line", func() {
			source := `(((index, term)))
a paragaph with an index term.`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.ConceleadIndexTerm{
									Term1: "index",
									Term2: "term",
								},
							},
							{
								types.StringElement{
									Content: "a paragaph with an index term.",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})

	Context("final document", func() {

		It("index term in existing paragraph line", func() {
			source := `a paragaph with an index term (((index, term, here))).`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "a paragaph with an index term ",
								},
								types.StringElement{
									Content: ".",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("index term in single paragraph line", func() {
			source := `(((index, term)))
a paragaph with an index term.`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "a paragaph with an index term.",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})
	})
})
