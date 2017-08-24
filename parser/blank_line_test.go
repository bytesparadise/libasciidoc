package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Blank lines", func() {
	It("blank line between 2 paragraphs", func() {
		actualDocument := `first paragraph

second paragraph`
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "first paragraph"},
							},
						},
					},
				},
				// &types.BlankLine{},
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "second paragraph"},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualDocument)
	})
	It("blank line with spaces and tabs between 2 paragraphs and after second paragraph", func() {
		actualDocument := `first paragraph
		 

		
second paragraph
`
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "first paragraph"},
							},
						},
					},
				},
				// &types.BlankLine{},
				// &types.BlankLine{},
				// &types.BlankLine{},
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "second paragraph"},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualDocument)
	})

})
