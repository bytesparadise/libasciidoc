package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("List Items", func() {

	Context("Unordered ", func() {
		Context("Valid content", func() {
			It("1 list with a single item", func() {
				actualContent := "* a list item"
				expectedDocument := &types.List{
					Items: []*types.ListItem{
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a list item"},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("List"))
			})
			It("1 list with an ID and a single item", func() {
				actualContent := "[#listID]\n" +
					"* a list item"
				expectedDocument := &types.List{
					ID: &types.ElementID{Value: "listID"},
					Items: []*types.ListItem{
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a list item"},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("List"))
			})

			It("1 list with 2 items with stars", func() {
				actualContent := "* a first item\n" +
					"* a second item with *bold content*"
				expectedDocument := &types.List{
					Items: []*types.ListItem{
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a first item"},
										},
									},
								},
							},
						},
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a second item with "},
											&types.QuotedText{Kind: types.Bold,
												Elements: []types.InlineElement{
													&types.StringElement{Content: "bold content"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("List"))
			})
			It("1 list with 2 items with carets", func() {
				actualContent := "- a first item\n" +
					"- a second item with *bold content*"
				expectedDocument := &types.List{
					Items: []*types.ListItem{
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a first item"},
										},
									},
								},
							},
						},
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a second item with "},
											&types.QuotedText{Kind: types.Bold,
												Elements: []types.InlineElement{
													&types.StringElement{Content: "bold content"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("List"))
			})
			It("1 list with 2 items with empty line in-between", func() {
				// fist line after list item is swallowed
				actualContent := "* a first item\n" +
					"\n" +
					"* a second item with *bold content*"
				expectedDocument := &types.List{
					Items: []*types.ListItem{
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a first item"},
										},
									},
								},
							},
						},
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a second item with "},
											&types.QuotedText{Kind: types.Bold,
												Elements: []types.InlineElement{
													&types.StringElement{Content: "bold content"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("List"))
			})
			It("1 list with 2 items on multiple lines", func() {
				actualContent := "* item 1\n" +
					"  on 2 lines.\n" +
					"* item 2\n" +
					"on 2 lines, too."
				expectedDocument := &types.List{
					Items: []*types.ListItem{
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "item 1"},
										},
									},
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "  on 2 lines."},
										},
									},
								},
							},
						},
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "item 2"},
										},
									},
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "on 2 lines, too."},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("List"))
			})
			It("2 lists with 2 empty lines in-between", func() {
				// the first blank lines after the first list is swallowed (for the list item)
				actualContent := "* an item in the first list\n" +
					"\n" +
					"\n" +
					"* an item in the second list"
				expectedDocument := &types.Document{
					Attributes:        map[string]interface{}{},
					ElementReferences: map[string]interface{}{},
					Elements: []types.DocElement{
						&types.List{
							Items: []*types.ListItem{
								&types.ListItem{
									Level: 1,
									Content: &types.ListItemContent{
										Lines: []*types.InlineContent{
											&types.InlineContent{
												Elements: []types.InlineElement{
													&types.StringElement{Content: "an item in the first list"},
												},
											},
										},
									},
								},
							},
						},
						&types.List{
							Items: []*types.ListItem{
								&types.ListItem{
									Level: 1,
									Content: &types.ListItemContent{
										Lines: []*types.InlineContent{
											&types.InlineContent{
												Elements: []types.InlineElement{
													&types.StringElement{Content: "an item in the second list"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent) // parse the whole document to get 2 lists
			})
		})

		Context("List of multiple levels", func() {
			It("a list with items on 3 levels", func() {
				actualContent := "* item 1\n" +
					"** item 1.1\n" +
					"** item 1.2\n" +
					"*** item 1.2.1\n" +
					"** item 1.3\n" +
					"** item 1.4\n" +
					"* item 2\n" +
					"** item 2.1\n"
				expectedDocument := &types.List{
					Items: []*types.ListItem{
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "item 1"},
										},
									},
								},
							},
							Children: &types.List{
								Items: []*types.ListItem{
									&types.ListItem{
										Level: 2,
										Content: &types.ListItemContent{
											Lines: []*types.InlineContent{
												&types.InlineContent{
													Elements: []types.InlineElement{
														&types.StringElement{Content: "item 1.1"},
													},
												},
											},
										},
									},
									&types.ListItem{
										Level: 2,
										Content: &types.ListItemContent{
											Lines: []*types.InlineContent{
												&types.InlineContent{
													Elements: []types.InlineElement{
														&types.StringElement{Content: "item 1.2"},
													},
												},
											},
										},
										Children: &types.List{
											Items: []*types.ListItem{
												&types.ListItem{
													Level: 3,
													Content: &types.ListItemContent{
														Lines: []*types.InlineContent{
															&types.InlineContent{
																Elements: []types.InlineElement{
																	&types.StringElement{Content: "item 1.2.1"},
																},
															},
														},
													},
												},
											},
										},
									},
									&types.ListItem{
										Level: 2,
										Content: &types.ListItemContent{
											Lines: []*types.InlineContent{
												&types.InlineContent{
													Elements: []types.InlineElement{
														&types.StringElement{Content: "item 1.3"},
													},
												},
											},
										},
									},
									&types.ListItem{
										Level: 2,
										Content: &types.ListItemContent{
											Lines: []*types.InlineContent{
												&types.InlineContent{
													Elements: []types.InlineElement{
														&types.StringElement{Content: "item 1.4"},
													},
												},
											},
										},
									},
								},
							},
						},
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "item 2"},
										},
									},
								},
							},
							Children: &types.List{
								Items: []*types.ListItem{
									&types.ListItem{
										Level: 2,
										Content: &types.ListItemContent{
											Lines: []*types.InlineContent{
												&types.InlineContent{
													Elements: []types.InlineElement{
														&types.StringElement{Content: "item 2.1"},
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
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("List"))
			})

		})

		Context("Invalid content", func() {
			It("a list with items on 2 levels - bad numbering", func() {
				actualContent := "* item 1\n" +
					"*** item 1.1\n" +
					"*** item 1.1.1\n" +
					"** item 1.2\n" +
					"* item 2"
				expectedDocument := &types.List{
					Items: []*types.ListItem{
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "item 1"},
										},
									},
								},
							},
							Children: &types.List{
								Items: []*types.ListItem{
									&types.ListItem{
										Level: 2,
										Content: &types.ListItemContent{
											Lines: []*types.InlineContent{
												&types.InlineContent{
													Elements: []types.InlineElement{
														&types.StringElement{Content: "item 1.1"},
													},
												},
											},
										},
										Children: &types.List{
											Items: []*types.ListItem{
												&types.ListItem{
													Level: 3,
													Content: &types.ListItemContent{
														Lines: []*types.InlineContent{
															&types.InlineContent{
																Elements: []types.InlineElement{
																	&types.StringElement{Content: "item 1.1.1"},
																},
															},
														},
													},
												},
											},
										},
									},
									&types.ListItem{
										Level: 2,
										Content: &types.ListItemContent{
											Lines: []*types.InlineContent{
												&types.InlineContent{
													Elements: []types.InlineElement{
														&types.StringElement{Content: "item 1.2"},
													},
												},
											},
										},
									},
								},
							},
						},
						&types.ListItem{
							Level: 1,
							Content: &types.ListItemContent{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "item 2"},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("List"))
			})

			It("invalid list item", func() {
				actualContent := "*an invalid list item"
				expectedDocument := &types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.StringElement{Content: "*an invalid list item"},
							},
						},
					},
				}
				verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("Paragraph"))
			})
		})
	})
})
