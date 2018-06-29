package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Blank lines", func() {
	It("blank line between 2 paragraphs", func() {
		actualDocument := `first paragraph
 
second paragraph`
		expectedResult := types.Document{
			Attributes:        map[string]interface{}{},
			ElementReferences: map[string]interface{}{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "first paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "second paragraph"},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualDocument)
	})
	It("blank line with spaces and tabs between 2 paragraphs and after second paragraph", func() {
		actualDocument := `first paragraph
		 

		
second paragraph
`
		expectedResult := types.Document{
			Attributes:        map[string]interface{}{},
			ElementReferences: map[string]interface{}{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "first paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.BlankLine{},
				types.BlankLine{},
				types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "second paragraph"},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualDocument)
	})

})
