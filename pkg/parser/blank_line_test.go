package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("blank lines - preflight", func() {
	It("blank line between 2 paragraphs", func() {
		source := `first paragraph
 
second paragraph`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "first paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "second paragraph"},
						},
					},
				},
			},
		}
		Expect(source).To(BecomePreflightDocument(expected))
	})
	It("blank line with spaces and tabs between 2 paragraphs and after second paragraph", func() {
		source := `first paragraph
		 

		
second paragraph
`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "second paragraph"},
						},
					},
				},
			},
		}
		Expect(source).To(BecomePreflightDocument(expected))
	})

})
