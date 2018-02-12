package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Delimited Blocks", func() {

	Context("Fenced blocks", func() {

		It("delimited fenced block with single line", func() {
			content := "some fenced code"
			actualContent := "```\n" + content + "\n```"
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited fenced block with no line", func() {
			content := ""
			actualContent := "```\n" + content + "```"
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited fenced block with multiple lines", func() {
			content := "some fenced code\nwith an empty line\n\nin the middle"
			actualContent := "```\n" + content + "\n```"
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited fenced block with multiple lines then a paragraph", func() {
			content := "some fenced code\nwith an empty line\n\nin the middle"
			actualContent := "```\n" + content + "\n```\nthen a normal paragraph."
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.FencedBlock,
						Content: content,
					},
					&types.Paragraph{
						Lines: []*types.InlineContent{
							{
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

		It("delimited fenced block after a paragraph", func() {
			content := "some fenced code"
			actualContent := "a paragraph.\n```\n" + content + "\n```\n"
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							{
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

	Context("Listing blocks", func() {

		It("delimited listing block with single line", func() {
			actualContent := `----
some listing code
----`
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.ListingBlock,
						Content: "some listing code",
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited listing block with no line", func() {
			content := ""
			actualContent := "----\n" + content + "----"
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.ListingBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited listing block with multiple lines", func() {
			content := "some listing code\nwith an empty line\n\nin the middle"
			actualContent := "----\n" + content + "\n----"
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.ListingBlock,
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("delimited listing block with multiple lines then a paragraph", func() {
			content := "some listing code\nwith an empty line\n\nin the middle"
			actualContent := "----\n" + content + "\n----\nthen a normal paragraph."
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.DelimitedBlock{
						Kind:    types.ListingBlock,
						Content: content,
					},
					&types.Paragraph{
						Lines: []*types.InlineContent{
							{
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

		It("delimited listing block after a paragraph", func() {
			actualContent := `a paragraph.
			
----
some listing code
----`
			expectedDocument := &types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph."},
								},
							},
						},
					},
					&types.DelimitedBlock{
						Kind:    types.ListingBlock,
						Content: "some listing code",
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

	})

})
