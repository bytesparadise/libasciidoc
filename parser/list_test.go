package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Parsing Unordered List Items", func() {

	Context("Valid content", func() {
		It("1 list with a single item", func() {
			actualContent := "* a list item"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
												&types.StringElement{Content: "a list item"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			compare(GinkgoT(), expectedDocument, actualContent)
		})
		It("1 list with an ID and a single item", func() {
			actualContent := "[#listID]\n" +
				"* a list item"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						ID: &types.ElementID{Value: "listID"},
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
												&types.StringElement{Content: "a list item"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			compare(GinkgoT(), expectedDocument, actualContent)
		})

		It("1 list with 2 items with stars", func() {
			actualContent := "* a first item\n" +
				"* a second item with *bold content*"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
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
											Elements: []types.DocElement{
												&types.StringElement{Content: "a second item with "},
												&types.QuotedText{Kind: types.Bold,
													Elements: []types.DocElement{
														&types.StringElement{Content: "bold content"},
													},
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
			compare(GinkgoT(), expectedDocument, actualContent)
		})
		It("1 list with 2 items with carets", func() {
			actualContent := "- a first item\n" +
				"- a second item with *bold content*"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
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
											Elements: []types.DocElement{
												&types.StringElement{Content: "a second item with "},
												&types.QuotedText{Kind: types.Bold,
													Elements: []types.DocElement{
														&types.StringElement{Content: "bold content"},
													},
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
			compare(GinkgoT(), expectedDocument, actualContent)
		})
		It("1 list with 2 items with empty line in-between", func() {
			actualContent := "* a first item\n" +
				"\n" +
				"* a second item with *bold content*"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
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
											Elements: []types.DocElement{
												&types.StringElement{Content: "a second item with "},
												&types.QuotedText{Kind: types.Bold,
													Elements: []types.DocElement{
														&types.StringElement{Content: "bold content"},
													},
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
			compare(GinkgoT(), expectedDocument, actualContent)
		})
		It("1 list with 2 items on multiple lines", func() {
			actualContent := "* item 1\n" +
				"  on 2 lines.\n" +
				"* item 2\n" +
				"on 2 lines, too."
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
												&types.StringElement{Content: "item 1"},
											},
										},
										&types.InlineContent{
											Elements: []types.DocElement{
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
											Elements: []types.DocElement{
												&types.StringElement{Content: "item 2"},
											},
										},
										&types.InlineContent{
											Elements: []types.DocElement{
												&types.StringElement{Content: "on 2 lines, too."},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			compare(GinkgoT(), expectedDocument, actualContent)
		})
		It("2 lists with 2 empty line in-between", func() {
			actualContent := "* an item in the first list\n" +
				"\n" +
				"\n" +
				"* an item in the second list"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
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
											Elements: []types.DocElement{
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
			compare(GinkgoT(), expectedDocument, actualContent)
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
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
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

														Elements: []types.DocElement{
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
														Elements: []types.DocElement{
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
																	Elements: []types.DocElement{
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
														Elements: []types.DocElement{
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
														Elements: []types.DocElement{
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
											Elements: []types.DocElement{
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
														Elements: []types.DocElement{
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
					},
				},
			}
			compare(GinkgoT(), expectedDocument, actualContent)
		})

	})

	Context("Invalid content", func() {
		It("a list with items on 2 levels - bad numbering", func() {
			actualContent := "* item 1\n" +
				"*** item 1.1\n" +
				"*** item 1.1.1\n" +
				"** item 1.2\n" +
				"* item 2"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.List{
						Items: []*types.ListItem{
							&types.ListItem{
								Level: 1,
								Content: &types.ListItemContent{
									Lines: []*types.InlineContent{
										&types.InlineContent{
											Elements: []types.DocElement{
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
														Elements: []types.DocElement{
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
																	Elements: []types.DocElement{
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
														Elements: []types.DocElement{
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
											Elements: []types.DocElement{
												&types.StringElement{Content: "item 2"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			compare(GinkgoT(), expectedDocument, actualContent)
		})

		It("invalid list item", func() {
			actualContent := "*an invalid list item"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.DocElement{
									&types.StringElement{Content: "*an invalid list item"},
								},
							},
						},
					},
				},
			}
			compare(GinkgoT(), expectedDocument, actualContent)
		})
	})

})
