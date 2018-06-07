package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("unordered lists", func() {

	Context("valid content", func() {

		It("unordered list with a single item", func() {
			actualContent := "* a list item"
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a list item"},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("unordered list with an ID and a single item", func() {
			actualContent := "[#listID]\n" +
				"* a list item"
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{
					"ID": "listID",
				},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a list item"},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("unordered list with 2 items with stars", func() {
			actualContent := "* a first item\n" +
				"* a second item with *bold content*"
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a first item"},
									},
								},
							},
						},
					},
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a second item with "},
										types.QuotedText{Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("unordered list with 2 items with carets", func() {
			actualContent := "- a first item\n" +
				"- a second item with *bold content*"
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a first item"},
									},
								},
							},
						},
					},
					{
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a second item with "},
										types.QuotedText{Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("unordered list with items with mixed styles", func() {
			actualContent := `- a parent item
				* a child item
				- another parent item
				* another child item
				** with a sub child item`
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a parent item"},
									},
								},
							},
							types.UnorderedList{
								Attributes: map[string]interface{}{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.OneAsterisk,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "a child item"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "another parent item"},
									},
								},
							},
							types.UnorderedList{
								Attributes: map[string]interface{}{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.OneAsterisk,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "another child item"},
													},
												},
											},
											types.UnorderedList{
												Attributes: map[string]interface{}{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.TwoAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "with a sub child item"},
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
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("unordered list with 2 items with empty line in-between", func() {
			// fist line after list item is swallowed
			actualContent := "* a first item\n" +
				"\n" +
				"* a second item with *bold content*"
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a first item"},
									},
								},
							},
						},
					},
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a second item with "},
										types.QuotedText{Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
		It("unordered list with 2 items on multiple lines", func() {
			actualContent := "* item 1\n" +
				"  on 2 lines.\n" +
				"* item 2\n" +
				"on 2 lines, too."
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1"},
									},
									{
										types.StringElement{Content: "  on 2 lines."},
									},
								},
							},
						},
					},
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 2"},
									},
									{
										types.StringElement{Content: "on 2 lines, too."},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
		It("2 unordered lists with 2 empty lines in-between", func() {
			// the first blank lines after the first list is swallowed (for the list item)
			actualContent := "* an item in the first list\n" +
				"\n" +
				"\n" +
				"* an item in the second list"
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: map[string]interface{}{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: map[string]interface{}{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "an item in the first list"},
											},
										},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.UnorderedList{
						Attributes: map[string]interface{}{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: map[string]interface{}{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "an item in the second list"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent) // parse the whole document to get 2 lists
		})

		It("unordered list with items on 3 levels", func() {
			actualContent := `* item 1
** item 1.1
** item 1.2
*** item 1.2.1
** item 1.3
** item 1.4
* item 2
** item 2.1`
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
							types.UnorderedList{
								Attributes: map[string]interface{}{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.1"},
													},
												},
											},
										},
									},
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.2"},
													},
												},
											},
											types.UnorderedList{
												Attributes: map[string]interface{}{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "item 1.2.1"},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.3"},
													},
												},
											},
										},
									},
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.4"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
							types.UnorderedList{
								Attributes: map[string]interface{}{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 2.1"},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

	})

	Context("invalid content", func() {
		It("unordered list with items on 2 levels - bad numbering", func() {
			actualContent := `* item 1
				*** item 1.1
				*** item 1.1.1
				** item 1.2
				* item 2`
			expectedResult := types.UnorderedList{
				Attributes: map[string]interface{}{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
							types.UnorderedList{
								Attributes: map[string]interface{}{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.1"},
													},
												},
											},
											types.UnorderedList{
												Attributes: map[string]interface{}{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "item 1.1.1"},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.2"},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("invalid list item", func() {
			actualContent := "*an invalid list item"
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "*an invalid list item"},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})
	})

	Context("list item continuation", func() {

		It("unordered list with item continuation", func() {
			actualContent := `* foo
+
----
a delimited block
----
+
----
another delimited block
----
* bar
`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: map[string]interface{}{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: map[string]interface{}{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "foo"},
											},
										},
									},
									types.DelimitedBlock{
										Kind:       types.ListingBlock,
										Attributes: map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "a delimited block",
														},
													},
												},
											},
										},
									},
									types.DelimitedBlock{
										Kind:       types.ListingBlock,
										Attributes: map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "another delimited block",
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: map[string]interface{}{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "bar"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("unordered list without item continuation", func() {
			actualContent := `* foo
----
a delimited block
----
* bar
----
another delimited block
----`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: map[string]interface{}{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: map[string]interface{}{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "foo"},
											},
										},
									},
								},
							},
						},
					},
					types.DelimitedBlock{
						Kind:       types.ListingBlock,
						Attributes: map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "a delimited block",
										},
									},
								},
							},
						},
					},
					types.UnorderedList{
						Attributes: map[string]interface{}{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: map[string]interface{}{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "bar"},
											},
										},
									},
								},
							},
						},
					},
					types.DelimitedBlock{
						Kind:       types.ListingBlock,
						Attributes: map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "another delimited block",
										},
									},
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
