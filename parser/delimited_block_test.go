package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Delimited Blocks", func() {

	Context("Source blocks", func() {

		It("delimited source block with single line", func() {
			content := "some source code"
			actualContent := "```\n" + content + "\n```"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited source block with no line", func() {
			content := ""
			actualContent := "```\n" + content + "```"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited source block with multiple lines", func() {
			content := "some source code\nwith an empty line\n\nin the middle"
			actualContent := "```\n" + content + "\n```"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited source block with multiple lines then a paragraph", func() {
			content := "some source code\nwith an empty line\n\nin the middle"
			actualContent := "```\n" + content + "\n```\nthen a normal paragraph."
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "then a normal paragraph."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited source block after a paragraph", func() {
			content := "some source code"
			actualContent := "a paragraph.\n```\n" + content + "\n```\n"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph."},
								},
							},
						},
					},
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

	})

})
