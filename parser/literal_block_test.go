package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Literal Blocks", func() {

	Context("Literal blocks with spaces indentation", func() {

		It("literal block from 1-line paragraph with single space", func() {
			actualContent := ` some literal content`
			expectedResult := types.LiteralBlock{
				Content: " some literal content",
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("literal block from paragraph with single space on first line", func() {
			actualContent := ` some literal content
on 2 lines.`
			expectedResult := types.LiteralBlock{
				Content: " some literal content\non 2 lines.",
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("mixing literal block and paragraph ", func() {
			actualContent := `   some literal content

a normal paragraph.`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.LiteralBlock{
						Content: "   some literal content",
					},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a normal paragraph."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("Literal blocks with block delimiter", func() {

		It("literal block from 1-line paragraph with delimiter", func() {
			actualContent := `....
some literal content
....
a normal paragraph.`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.LiteralBlock{
						Content: "some literal content",
					},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a normal paragraph."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})

	Context("Literal blocks with attribute", func() {

		It("literal block from 1-line paragraph with attribute", func() {
			actualContent := `[literal]   
some literal content

a normal paragraph.`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.LiteralBlock{
						Content: "some literal content",
					},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a normal paragraph."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

})
