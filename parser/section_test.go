package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Sections", func() {

	Context("Valid document", func() {

		It("header only", func() {
			actualContent := "= a header"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{
					"doctitle": "a header",
				},
				Elements: []types.DocElement{},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("header and paragraph", func() {
			actualContent := `= a header

and a paragraph`
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{
					"doctitle": "a header",
				},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "and a paragraph"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 1 alone", func() {
			actualContent := `== section 1`
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "section 1"},
								},
							},
							ID: &types.ElementID{
								Value: "_section_1",
							},
						},
						Elements: []types.DocElement{},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 1 with quoted text", func() {
			actualContent := `==  *2 spaces and bold content*`
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.QuotedText{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "2 spaces and bold content"},
										},
									},
								},
							},
							ID: &types.ElementID{
								Value: "__strong_2_spaces_and_bold_content_strong",
							},
						},
						Elements: []types.DocElement{},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 0 with nested section level 1", func() {
			actualContent := `= a header

== section 1`
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{
					"doctitle": "a header",
				}, Elements: []types.DocElement{
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "section 1"},
								},
							},
							ID: &types.ElementID{
								Value: "_section_1",
							},
						},
						Elements: []types.DocElement{},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 0 with preamble and section level 1", func() {
			actualContent := `= a header

a short preamble

== section 1`
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{
					"doctitle": "a header",
				}, Elements: []types.DocElement{
					&types.Preamble{
						Elements: []types.DocElement{
							&types.Paragraph{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a short preamble"},
										},
									},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "section 1"},
								},
							},
							ID: &types.ElementID{
								Value: "_section_1",
							},
						},
						Elements: []types.DocElement{},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 0 with nested section level 2", func() {
			actualContent := "= a header\n" +
				"\n" +
				"=== section 2"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{
					"doctitle": "a header",
				},
				Elements: []types.DocElement{
					&types.Section{
						Level: 2,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "section 2"},
								},
							},
							ID: &types.ElementID{
								Value: "_section_2",
							},
						},
						Elements: []types.DocElement{},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 1 with immediate paragraph", func() {
			actualContent := `== a title
and a paragraph`
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a title"},
								},
							},
							ID: &types.ElementID{
								Value: "_a_title",
							},
						},
						Elements: []types.DocElement{
							&types.Paragraph{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "and a paragraph"},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
		It("section level 1 with a paragraph separated by empty line", func() {
			actualContent := "== a title\n\nand a paragraph"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a title"},
								},
							},
							ID: &types.ElementID{
								Value: "_a_title",
							},
						},
						Elements: []types.DocElement{
							&types.Paragraph{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "and a paragraph"},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 1 with a paragraph separated by non-empty line", func() {
			actualContent := "== a title\n    \nand a paragraph"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a title"},
								},
							},
							ID: &types.ElementID{
								Value: "_a_title",
							},
						},
						Elements: []types.DocElement{
							&types.Paragraph{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "and a paragraph"},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section levels 1, 2, 3, 2", func() {
			actualContent := `= a title

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{
					"doctitle": "a title",
				},
				Elements: []types.DocElement{
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "Section A"},
								},
							},
							ID: &types.ElementID{
								Value: "_section_a",
							},
						},
						Elements: []types.DocElement{
							&types.Paragraph{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a paragraph"},
										},
									},
								},
							},
							// &types.BlankLine{},
							&types.Section{
								Level: 2,
								SectionTitle: types.SectionTitle{
									Content: &types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "Section A.a"},
										},
									},
									ID: &types.ElementID{
										Value: "_section_a_a",
									},
								},
								Elements: []types.DocElement{
									&types.Paragraph{
										Lines: []*types.InlineContent{
											&types.InlineContent{
												Elements: []types.InlineElement{
													&types.StringElement{Content: "a paragraph"},
												},
											},
										},
									},
									// &types.BlankLine{},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						SectionTitle: types.SectionTitle{
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "Section B"},
								},
							},
							ID: &types.ElementID{
								Value: "_section_b",
							},
						},
						Elements: []types.DocElement{
							&types.Paragraph{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a paragraph"},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})

	Context("Invalid document", func() {
		It("header invalid - missing space", func() {
			actualContent := "=a header"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "=a header"},
								},
							},
						},
					},
				}}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("header invalid - header space", func() {
			actualContent := " = a header"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{},
				Elements: []types.DocElement{
					&types.LiteralBlock{
						Content: " = a header",
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("header with invalid section1", func() {
			actualContent := "= a header\n" +
				"\n" +
				" == section 1"
			expectedDocument := &types.Document{
				Attributes: map[string]interface{}{
					"doctitle": "a header",
				},
				Elements: []types.DocElement{
					&types.LiteralBlock{
						Content: " == section 1",
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

	})
})
