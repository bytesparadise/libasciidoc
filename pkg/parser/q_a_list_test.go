package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("q and a lists", func() {

	It("q and a with title", func() {
		actualContent := `.Q&A
[qanda]
What is libsciidoc?::
	An implementation of the AsciiDoc processor in Golang.
What is the answer to the Ultimate Question?:: 42`

		expectedResult := types.LabeledList{
			Attributes: types.ElementAttributes{
				types.AttrTitle: "Q&A",
				types.AttrQandA: nil,
			},
			Items: []types.LabeledListItem{
				{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "What is libsciidoc?",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Term:       "What is the answer to the Ultimate Question?",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
		}
		verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})
})
