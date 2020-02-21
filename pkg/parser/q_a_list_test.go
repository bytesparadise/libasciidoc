package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("q and a lists", func() {

	It("q and a with title", func() {
		source := `.Q&A
[qanda]
What is libsciidoc?::
	An implementation of the AsciiDoc processor in Golang.
What is the answer to the Ultimate Question?:: 42`

		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{
						types.AttrTitle: "Q&A",
						types.AttrQandA: nil,
					},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "What is libsciidoc?",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "An implementation of the AsciiDoc processor in Golang.",
											},
										},
									},
								},
							},
						},
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "What is the answer to the Ultimate Question?",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "42",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDocument(source)).To(Equal(expected))
	})
})
