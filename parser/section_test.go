package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Sections", func() {

	Context("Valid document", func() {

		It("section with heading only", func() {
			actualContent := "= a heading"
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{
					"title": "a heading",
				},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 1,
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a heading"},
								},
							},
							ID: &types.ElementID{
								Value: "_a_heading",
							},
						},
						Elements: []types.DocElement{},
					},
				}}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 2", func() {
			actualContent := `== section 2`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 2,
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

		It("section level 2 with quoted text", func() {
			actualContent := `==  *2 spaces and bold content*`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 2,
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

		It("section level 1 with nested section level 2", func() {
			actualContent := "= a heading\n" +
				"\n" +
				"== section 2"
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{
					"title": "a heading",
				}, Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 1,
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a heading"},
								},
							},
							ID: &types.ElementID{
								Value: "_a_heading",
							},
						},
						Elements: []types.DocElement{
							&types.Section{
								Heading: types.Heading{
									Level: 2,
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
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 1 with nested section level 3", func() {
			actualContent := "= a heading\n" +
				"\n" +
				"=== section 3"
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{
					"title": "a heading",
				},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 1,
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a heading"},
								},
							},
							ID: &types.ElementID{
								Value: "_a_heading",
							},
						},
						Elements: []types.DocElement{
							&types.Section{
								Heading: types.Heading{
									Level: 3,
									Content: &types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "section 3"},
										},
									},
									ID: &types.ElementID{
										Value: "_section_3",
									},
								},
								Elements: []types.DocElement{},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("section level 2 with immediate paragraph", func() {
			actualContent := `== a title
and a paragraph`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 2,
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
		It("section level 2 with a paragraph separated by empty line", func() {
			actualContent := "== a title\n\nand a paragraph"
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 2,
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

		It("section level 2 with a paragraph separated by non-empty line", func() {
			actualContent := "== a title\n    \nand a paragraph"
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 2,
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
				Attributes: &types.DocumentAttributes{
					"title": "a title",
				},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 1,
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
							&types.Section{
								Heading: types.Heading{
									Level: 2,
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
										Heading: types.Heading{
											Level: 3,
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
								Heading: types.Heading{
									Level: 2,
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
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})

	Context("Invalid document", func() {
		It("heading invalid - missing space", func() {
			actualContent := "=a heading"
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "=a heading"},
								},
							},
						},
					},
				}}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("heading invalid - heading space", func() {
			actualContent := " = a heading"
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: " = a heading"},
								},
							},
						},
					},
				}}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("heading with invalid section2", func() {
			actualContent := "= a heading\n" +
				"\n" +
				" == section 2"
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{
					"title": "a heading",
				},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 1, Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a heading"},
								},
							},
							ID: &types.ElementID{
								Value: "_a_heading",
							},
						},
						Elements: []types.DocElement{
							&types.Paragraph{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: " == section 2"},
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
})
