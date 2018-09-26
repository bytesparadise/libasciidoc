package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("unordered lists", func() {

	Context("valid content", func() {

		It("unordered list with a basic single item", func() {
			actualContent := `* a list item`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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

		It("unordered list with ID, title, role and a single item", func() {
			actualContent := `.mytitle
[#listID]
[.myrole]
* a list item`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{
					types.AttrID:    "listID",
					types.AttrTitle: "mytitle",
					types.AttrRole:  "myrole",
				},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
		It("unordered list with a title and a single item", func() {
			actualContent := `.a title
* a list item`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "a title",
				},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			actualContent := `* a first item
				* a second item with *bold content*`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a second item with "},
										types.QuotedText{
											Attributes: types.ElementAttributes{
												types.AttrKind: types.Bold,
											},
											Elements: types.InlineElements{
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

		It("unordered list based on article.adoc (with heading spaces)", func() {
			actualContent := `.Unordered list title
	* list item 1
	** nested list item A
	*** nested nested list item A.1
	*** nested nested list item A.2
	** nested list item B
	*** nested nested list item B.1
	*** nested nested list item B.2
	* list item 2`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Unordered list title",
				},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "list item 1"},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "nested list item A"},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "nested nested list item A.1"},
																	},
																},
															},
														},
													},
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "nested nested list item A.2"},
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
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "nested list item B"},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "nested nested list item B.1"},
																	},
																},
															},
														},
													},
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "nested nested list item B.2"},
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
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "list item 2"},
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
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a second item with "},
										types.QuotedText{
											Attributes: types.ElementAttributes{
												types.AttrKind: types.Bold,
											},
											Elements: types.InlineElements{
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
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a parent item"},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.OneAsterisk,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
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
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "another parent item"},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.OneAsterisk,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "another child item"},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.TwoAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a second item with "},
										types.QuotedText{
											Attributes: types.ElementAttributes{
												types.AttrKind: types.Bold,
											},
											Elements: types.InlineElements{
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
			actualContent := `* item 1
  on 2 lines.
* item 2
on 2 lines, too.`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
								Attributes: types.ElementAttributes{},
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
		It("unordered lists with 2 empty lines in-between", func() {
			// the first blank lines after the first list is swallowed (for the list item)
			actualContent := "* an item in the first list\n" +
				"\n" +
				"\n" +
				"* an item in the second list"
			expectedResult := types.Document{
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "an item in the first list"},
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
										Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
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
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.2"},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
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
												Attributes: types.ElementAttributes{},
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
												Attributes: types.ElementAttributes{},
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
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.1"},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
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
												Attributes: types.ElementAttributes{},
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
								Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{},
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
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "foo"},
											},
										},
									},
									types.DelimitedBlock{
										Attributes: types.ElementAttributes{
											types.AttrKind: types.Listing,
										},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
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
										Attributes: types.ElementAttributes{
											types.AttrKind: types.Listing,
										},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
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
										Attributes: types.ElementAttributes{},
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
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
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
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Listing,
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
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
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Listing,
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
