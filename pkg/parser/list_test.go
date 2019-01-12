package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("lists", func() {

	Context("labeled lists", func() {

		It("labeled list with a term and description on 2 lines", func() {
			actualContent := `Item1::
Item 1 description
on 2 lines`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Item1",
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "Item 1 description"},
									},
									{
										types.StringElement{Content: "on 2 lines"},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("labeled list with a single term and no description", func() {
			actualContent := `Item1::`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Term:       "Item1",
						Level:      1,
						Elements:   []interface{}{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("labeled list with a horizontal layout attribute", func() {
			actualContent := `[horizontal]
Item1:: foo`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{
					"layout": "horizontal",
				},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Item1",
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("labeled list with a single term and a blank line", func() {
			actualContent := `Item1::
			`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Item1",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("labeled list with multiple sibling items", func() {
			actualContent := `Item 1::
Item 1 description
Item 2:: 
Item 2 description
Item 3:: 
Item 3 description`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Item 1",
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "Item 1 description"},
									},
								},
							},
						},
					},
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Item 2",
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "Item 2 description"},
									},
								},
							},
						},
					},
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Item 3",
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "Item 3 description"},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("labeled list with multiple nested items", func() {
			actualContent := `Item 1::
Item 1 description
Item 2:::
Item 2 description
Item 3::::
Item 3 description`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Item 1",
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "Item 1 description"},
									},
								},
							},
							types.LabeledList{
								Attributes: types.ElementAttributes{},
								Items: []types.LabeledListItem{
									{
										Attributes: types.ElementAttributes{},
										Level:      2,
										Term:       "Item 2",
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "Item 2 description"},
													},
												},
											},
											types.LabeledList{
												Attributes: types.ElementAttributes{},
												Items: []types.LabeledListItem{
													{
														Attributes: types.ElementAttributes{},
														Level:      3,
														Term:       "Item 3",
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "Item 3 description"},
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

		It("labeled list with nested list", func() {
			actualContent := `Empty item:: 
* foo
* bar
Item with description:: something simple`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Empty item",
						Elements: []interface{}{
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Attributes:  types.ElementAttributes{},
										Level:       1,
										BulletStyle: types.OneAsterisk,
										CheckStyle:  types.NoCheck,
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
									{
										Attributes:  types.ElementAttributes{},
										Level:       1,
										BulletStyle: types.OneAsterisk,
										CheckStyle:  types.NoCheck,
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
					},
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Item with description",
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "something simple"},
									},
								},
							},
						},
					},
				},
			}

			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("labeled list with a single item and paragraph", func() {
			actualContent := `Item 1::
foo
bar

a normal paragraph.`
			expectedResult := types.Document{
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LabeledList{
						Attributes: types.ElementAttributes{},
						Items: []types.LabeledListItem{
							{
								Attributes: types.ElementAttributes{},
								Level:      1,
								Term:       "Item 1",
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "foo"},
											},
											{
												types.StringElement{Content: "bar"},
											},
										},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a normal paragraph."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("labeled list with item continuation", func() {
			actualContent := `Item 1::
+
----
a fenced block
----
Item 2:: something simple
+
----
another fenced block
----`
			expectedResult := types.Document{
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LabeledList{
						Attributes: types.ElementAttributes{},
						Items: []types.LabeledListItem{
							{
								Attributes: types.ElementAttributes{},
								Level:      1,
								Term:       "Item 1",
								Elements: []interface{}{
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
															Content: "a fenced block",
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Attributes: types.ElementAttributes{},
								Level:      1,
								Term:       "Item 2",
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "something simple"},
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
															Content: "another fenced block",
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

			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("labeled list without item continuation", func() {
			actualContent := `Item 1::
----
a fenced block
----
Item 2:: something simple
----
another fenced block
----`
			expectedResult := types.Document{
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LabeledList{
						Attributes: types.ElementAttributes{},
						Items: []types.LabeledListItem{
							{
								Attributes: types.ElementAttributes{},
								Level:      1,
								Term:       "Item 1",
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
											Content: "a fenced block",
										},
									},
								},
							},
						},
					},
					types.LabeledList{
						Attributes: types.ElementAttributes{},
						Items: []types.LabeledListItem{
							{
								Attributes: types.ElementAttributes{},
								Level:      1,
								Term:       "Item 2",
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "something simple"},
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
											Content: "another fenced block",
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

		It("labeled list with nested unordered list", func() {
			actualContent := `Labeled item::
- unordered item`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Labeled item",
						Elements: []interface{}{
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Attributes:  types.ElementAttributes{},
										Level:       1,
										BulletStyle: types.Dash,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "unordered item"},
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

	Context("unordered list", func() {
		Context("ordered list item alone", func() {

			// same single item in the list for each test in this context
			elements := []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "item"},
						},
					},
				},
			}
			It("ordered list item with implicit numbering style", func() {
				actualContent := `.. item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.LowerAlpha,
							Attributes:     map[string]interface{}{},
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list item with arabic numbering style", func() {
				actualContent := `1. item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list item with lower alpha numbering style", func() {
				actualContent := `b. item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.LowerAlpha,
							Attributes:     map[string]interface{}{},
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list item with upper alpha numbering style", func() {
				actualContent := `B. item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{

							Level:          1,
							Position:       1,
							NumberingStyle: types.UpperAlpha,
							Attributes:     map[string]interface{}{},
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list item with lower roman numbering style", func() {
				actualContent := `i) item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.LowerRoman,
							Attributes:     map[string]interface{}{},
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list item with upper roman numbering style", func() {
				actualContent := `I) item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{

							Level:          1,
							Position:       1,
							NumberingStyle: types.UpperRoman,
							Attributes:     map[string]interface{}{},
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list item with explicit numbering style", func() {
				actualContent := `[lowerroman]
. item
. item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{
						"lowerroman": nil,
					},
					Items: []types.OrderedListItem{
						{
							Attributes:     types.ElementAttributes{},
							Level:          1,
							Position:       1,
							NumberingStyle: types.LowerRoman,
							Elements:       elements,
						},
						{
							Attributes:     types.ElementAttributes{},
							Level:          1,
							Position:       2,
							NumberingStyle: types.LowerRoman,
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list item with explicit start only", func() {
				actualContent := `[start=5]
	. item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{
						"start": "5",
					},
					Items: []types.OrderedListItem{
						{

							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     types.ElementAttributes{},
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list item with explicit quoted numbering and start", func() {
				actualContent := `["lowerroman", start="5"]
	. item`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{
						"lowerroman": nil,
						"start":      "5",
					},
					Items: []types.OrderedListItem{
						{

							Level:          1,
							Position:       1,
							NumberingStyle: types.LowerRoman,
							Attributes:     types.ElementAttributes{},
							Elements:       elements,
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

		})

		Context("items without numbers", func() {

			It("ordered list with simple unnumbered items", func() {
				actualContent := `. a
	. b`

				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "a"},
										},
									},
								},
							},
						},
						{
							Level:          1,
							Position:       2,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "b"},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list with unnumbered items", func() {
				actualContent := `. item 1
	. item 2`

				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "item 1"},
										},
									},
								},
							},
						},
						{
							Level:          1,
							Position:       2,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
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

			It("ordered list with custom numbering on child items with tabs ", func() {
				// note: the [upperroman] attribute must be at the beginning of the line
				actualContent := `. item 1
				.. item 1.1
[upperroman]
				... item 1.1.1
				... item 1.1.2
				.. item 1.2
				. item 2
				.. item 2.1`

				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "item 1"},
										},
									},
								},
								types.OrderedList{
									Attributes: types.ElementAttributes{},
									Items: []types.OrderedListItem{
										{
											Level:          2,
											Position:       1,
											NumberingStyle: types.LowerAlpha,
											Attributes:     map[string]interface{}{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
														{
															types.StringElement{Content: "item 1.1"},
														},
													},
												},
												types.OrderedList{
													Attributes: types.ElementAttributes{},
													Items: []types.OrderedListItem{
														{
															Level:          3,
															Position:       1,
															NumberingStyle: types.UpperRoman,
															Attributes: types.ElementAttributes{
																"upperroman": nil,
															},
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
														{
															Level:          3,
															Position:       2,
															NumberingStyle: types.UpperRoman,
															Attributes:     map[string]interface{}{},
															Elements: []interface{}{
																types.Paragraph{
																	Attributes: types.ElementAttributes{},
																	Lines: []types.InlineElements{
																		{
																			types.StringElement{Content: "item 1.1.2"},
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
											Level:          2,
											Position:       2,
											NumberingStyle: types.LowerAlpha,
											Attributes:     map[string]interface{}{},
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
							Level:          1,
							Position:       2,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "item 2"},
										},
									},
								},
								types.OrderedList{
									Attributes: types.ElementAttributes{},
									Items: []types.OrderedListItem{
										{
											Level:          2,
											Position:       1,
											NumberingStyle: types.LowerAlpha,
											Attributes:     map[string]interface{}{},
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

			It("ordered list with all default styles and blank lines", func() {
				actualContent := `. level 1
	
	.. level 2
	
	
	... level 3
	
	
	
	.... level 4
	..... level 5.
	
	
	`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "level 1"},
										},
									},
								},
								types.OrderedList{
									Attributes: types.ElementAttributes{},
									Items: []types.OrderedListItem{
										{
											Level:          2,
											Position:       1,
											NumberingStyle: types.LowerAlpha,
											Attributes:     map[string]interface{}{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
														{
															types.StringElement{Content: "level 2"},
														},
													},
												},
												types.OrderedList{
													Attributes: types.ElementAttributes{},
													Items: []types.OrderedListItem{
														{
															Level:          3,
															Position:       1,
															NumberingStyle: types.LowerRoman,
															Attributes:     map[string]interface{}{},
															Elements: []interface{}{
																types.Paragraph{
																	Attributes: types.ElementAttributes{},
																	Lines: []types.InlineElements{
																		{
																			types.StringElement{Content: "level 3"},
																		},
																	},
																},
																types.OrderedList{
																	Attributes: types.ElementAttributes{},
																	Items: []types.OrderedListItem{
																		{
																			Level:          4,
																			Position:       1,
																			NumberingStyle: types.UpperAlpha,
																			Attributes:     map[string]interface{}{},
																			Elements: []interface{}{
																				types.Paragraph{
																					Attributes: types.ElementAttributes{},
																					Lines: []types.InlineElements{
																						{
																							types.StringElement{Content: "level 4"},
																						},
																					},
																				},
																				types.OrderedList{
																					Attributes: types.ElementAttributes{},
																					Items: []types.OrderedListItem{
																						{
																							Level:          5,
																							Position:       1,
																							NumberingStyle: types.UpperRoman,
																							Attributes:     map[string]interface{}{},
																							Elements: []interface{}{
																								types.Paragraph{
																									Attributes: types.ElementAttributes{},
																									Lines: []types.InlineElements{
																										{
																											types.StringElement{Content: "level 5."},
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

		Context("numbered items", func() {

			It("ordered list with simple numbered items", func() {
				actualContent := `1. a
	2. b`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "a"},
										},
									},
								},
							},
						},
						{
							Level:          1,
							Position:       2,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "b"},
										},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("ordered list with numbered items", func() {
				actualContent := `1. item 1
	a. item 1.a
	2. item 2
	b. item 2.a`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "item 1"},
										},
									},
								},
								types.OrderedList{
									Attributes: types.ElementAttributes{},
									Items: []types.OrderedListItem{
										{
											Level:          2,
											Position:       1,
											NumberingStyle: types.LowerAlpha,
											Attributes:     map[string]interface{}{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
														{
															types.StringElement{Content: "item 1.a"},
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
							Level:          1,
							Position:       2,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "item 2"},
										},
									},
								},
								types.OrderedList{
									Attributes: types.ElementAttributes{},
									Items: []types.OrderedListItem{
										{
											Level:          2,
											Position:       1,
											NumberingStyle: types.LowerAlpha,
											Attributes:     map[string]interface{}{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
														{
															types.StringElement{Content: "item 2.a"},
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

		Context("mixed lists", func() {

			It("ordered list with nested unordered lists", func() {
				actualContent := `. Item 1
	* Item A
	* Item B
	. Item 2
	* Item C
	* Item D`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Attributes:     map[string]interface{}{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "Item 1"},
										},
									},
								},
								types.UnorderedList{
									Attributes: types.ElementAttributes{},
									Items: []types.UnorderedListItem{
										{
											Attributes:  types.ElementAttributes{},
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
														{
															types.StringElement{Content: "Item A"},
														},
													},
												},
											},
										},
										{
											Attributes:  types.ElementAttributes{},
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
														{
															types.StringElement{Content: "Item B"},
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
							Attributes:     types.ElementAttributes{},
							Level:          1,
							Position:       2,
							NumberingStyle: types.Arabic,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "Item 2"},
										},
									},
								},
								types.UnorderedList{
									Attributes: types.ElementAttributes{},
									Items: []types.UnorderedListItem{
										{
											Attributes:  types.ElementAttributes{},
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
														{
															types.StringElement{Content: "Item C"},
														},
													},
												},
											},
										},
										{
											Attributes:  types.ElementAttributes{},
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
														{
															types.StringElement{Content: "Item D"},
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

		Context("invalid ordered list item prefix", func() {

			It("should not match", func() {
				actualContent := `foo. content`
				verifyError(GinkgoT(), actualContent, parser.Entrypoint("List")) // here we expect this will not be a valid list
			})
		})
	})

	Context("unordered lists", func() {

		Context("valid content", func() {

			It("unordered list with a basic single item", func() {
				actualContent := `* a list item`
				expectedResult := types.UnorderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.UnorderedListItem{
						{
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
															Attributes:  types.ElementAttributes{},
															Level:       3,
															BulletStyle: types.ThreeAsterisks,
															CheckStyle:  types.NoCheck,
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
															Attributes:  types.ElementAttributes{},
															Level:       3,
															BulletStyle: types.ThreeAsterisks,
															CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
															Attributes:  types.ElementAttributes{},
															Level:       3,
															BulletStyle: types.ThreeAsterisks,
															CheckStyle:  types.NoCheck,
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
															Attributes:  types.ElementAttributes{},
															Level:       3,
															BulletStyle: types.ThreeAsterisks,
															CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.Dash,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.Dash,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.Dash,
							CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.Dash,
							CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
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
															Attributes:  types.ElementAttributes{},
															Level:       3,
															BulletStyle: types.TwoAsterisks,
															CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
									Attributes:  types.ElementAttributes{},
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
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
									Attributes:  types.ElementAttributes{},
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
															Attributes:  types.ElementAttributes{},
															Level:       3,
															BulletStyle: types.ThreeAsterisks,
															CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
															Attributes:  types.ElementAttributes{},
															Level:       3,
															BulletStyle: types.ThreeAsterisks,
															CheckStyle:  types.NoCheck,
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
											Attributes:  types.ElementAttributes{},
											Level:       2,
											BulletStyle: types.TwoAsterisks,
											CheckStyle:  types.NoCheck,
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
							Attributes:  types.ElementAttributes{},
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
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
									Attributes:  types.ElementAttributes{},
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
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
									Attributes:  types.ElementAttributes{},
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
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
									Attributes:  types.ElementAttributes{},
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
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
									Attributes:  types.ElementAttributes{},
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
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

	Context("complex cases", func() {

		It("complex case 1 - mixed lists", func() {
			actualContent := `- unordered 1
	1. ordered 1.1
		a. ordered 1.1.a
		b. ordered 1.1.b
		c. ordered 1.1.c
	2. ordered 1.2
		i)  ordered 1.2.i
		ii) ordered 1.2.ii
	3. ordered 1.3
	4. ordered 1.4
	- unordered 2
	* unordered 2.1`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "unordered 1"},
									},
								},
							},
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.1"},
													},
												},
											},
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       1,
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.a"},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       2,
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.b"},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       3,
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.c"},
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
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       2,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.2"},
													},
												},
											},
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       1,
														NumberingStyle: types.LowerRoman,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.2.i"},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       2,
														NumberingStyle: types.LowerRoman,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.2.ii"},
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
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       3,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.3"},
													},
												},
											},
										},
									},
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       4,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.4"},
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
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "unordered 2"},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.OneAsterisk,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "unordered 2.1"},
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

		It("complex case 2 - mixed lists", func() {
			actualContent := `- unordered 1
1. ordered 1.1
a. ordered 1.1.a
b. ordered 1.1.b
c. ordered 1.1.c
2. ordered 1.2
i)  ordered 1.2.i
ii) ordered 1.2.ii
3. ordered 1.3
4. ordered 1.4
- unordered 2
* unordered 2.1
** unordered 2.1.1
	with some
	extra lines.
** unordered 2.1.2
* unordered 2.2
- unordered 3
. ordered 3.1
. ordered 3.2
[upperroman]
.. ordered 3.2.I
.. ordered 3.2.II
. ordered 3.3`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "unordered 1"},
									},
								},
							},
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.1"},
													},
												},
											},
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       1,
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.a"},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       2,
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.b"},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       3,
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.c"},
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
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       2,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.2"},
													},
												},
											},
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       1,
														NumberingStyle: types.LowerRoman,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.2.i"},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       2,
														NumberingStyle: types.LowerRoman,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.2.ii"},
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
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       3,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.3"},
													},
												},
											},
										},
									},
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       4,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.4"},
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
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "unordered 2"},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.OneAsterisk,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "unordered 2.1"},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Attributes:  types.ElementAttributes{},
														Level:       3,
														BulletStyle: types.TwoAsterisks,
														CheckStyle:  types.NoCheck,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "unordered 2.1.1"},
																	},
																	{
																		types.StringElement{Content: "\twith some"},
																	},
																	{
																		types.StringElement{Content: "\textra lines."},
																	},
																},
															},
														},
													},
													{
														Attributes:  types.ElementAttributes{},
														Level:       3,
														BulletStyle: types.TwoAsterisks,
														CheckStyle:  types.NoCheck,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "unordered 2.1.2"},
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
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.OneAsterisk,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "unordered 2.2"},
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
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "unordered 3"},
									},
								},
							},
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 3.1"},
													},
												},
											},
										},
									},
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       2,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 3.2"},
													},
												},
											},
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Level:          2,
														Position:       1,
														NumberingStyle: types.UpperRoman,
														Attributes: types.ElementAttributes{
															"upperroman": nil,
														},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 3.2.I"},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          2,
														Position:       2,
														NumberingStyle: types.UpperRoman,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 3.2.II"},
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
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       3,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 3.3"},
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

		It("complex case 3 - unordered list with continuation", func() {
			actualContent := `.Unordered, complex
* level 1
** level 2
*** level 3
This is a new line inside an unordered list using {plus} symbol.
We can even force content to start on a separate line... +
Amazing, isn't it?
**** level 4
+
The {plus} symbol is on a new line.

***** level 5
`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Unordered, complex",
				},
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 1",
										},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.NoCheck,
										Attributes:  types.ElementAttributes{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "level 2",
														},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														CheckStyle:  types.NoCheck,
														Attributes:  types.ElementAttributes{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "level 3",
																		},
																	},
																	{
																		types.StringElement{
																			Content: "This is a new line inside an unordered list using ",
																		},
																		types.DocumentAttributeSubstitution{
																			Name: "plus",
																		},
																		types.StringElement{
																			Content: " symbol.",
																		},
																	},
																	{
																		types.StringElement{
																			Content: "We can even force content to start on a separate line...",
																		},
																		types.LineBreak{},
																	},
																	{
																		types.StringElement{
																			Content: "Amazing, isn't it?",
																		},
																	},
																},
															},
															types.UnorderedList{
																Attributes: types.ElementAttributes{},
																Items: []types.UnorderedListItem{
																	{
																		Level:       4,
																		BulletStyle: types.FourAsterisks,
																		CheckStyle:  types.NoCheck,
																		Attributes:  types.ElementAttributes{},
																		Elements: []interface{}{
																			types.Paragraph{
																				Attributes: types.ElementAttributes{},
																				Lines: []types.InlineElements{
																					{
																						types.StringElement{
																							Content: "level 4",
																						},
																					},
																				},
																			},
																			// the `+` continuation produces the second paragrap below
																			types.Paragraph{
																				Attributes: types.ElementAttributes{},
																				Lines: []types.InlineElements{
																					{
																						types.StringElement{
																							Content: "The ",
																						},
																						types.DocumentAttributeSubstitution{
																							Name: "plus",
																						},
																						types.StringElement{
																							Content: " symbol is on a new line.",
																						},
																					},
																				},
																			},

																			types.UnorderedList{
																				Attributes: types.ElementAttributes{},
																				Items: []types.UnorderedListItem{
																					{
																						Level:       5,
																						BulletStyle: types.FiveAsterisks,
																						CheckStyle:  types.NoCheck,
																						Attributes:  types.ElementAttributes{},
																						Elements: []interface{}{
																							types.Paragraph{
																								Attributes: types.ElementAttributes{},
																								Lines: []types.InlineElements{
																									{
																										types.StringElement{
																											Content: "level 5",
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

		It("complex case 4 - mixed lists", func() {
			actualContent := `.Mixed
Operating Systems::
  . Fedora
    * Desktop`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Mixed",
				},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Operating Systems",
						Elements: []interface{}{
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										Position:       1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "Fedora",
														},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Attributes:  types.ElementAttributes{},
														Level:       1,
														BulletStyle: types.OneAsterisk,
														CheckStyle:  types.NoCheck,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "Desktop",
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
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("complex case 5 - mixed lists", func() {
			actualContent := `.Mixed
Operating Systems::
  Linux:::
    . Fedora
      * Desktop
    . Ubuntu
      * Desktop
      * Server
  BSD:::
    . FreeBSD
    . NetBSD

Cloud Providers::
  PaaS:::
    . OpenShift
    . CloudBees
  IaaS:::
    . Amazon EC2
    . Rackspace
`
			expectedResult := types.LabeledList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Mixed",
				},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Operating Systems",
						Elements: []interface{}{
							types.LabeledList{
								Attributes: types.ElementAttributes{},
								Items: []types.LabeledListItem{
									{
										Attributes: types.ElementAttributes{},
										Level:      2,
										Term:       "Linux",
										Elements: []interface{}{
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          1,
														Position:       1,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "Fedora",
																		},
																	},
																},
															},
															types.UnorderedList{
																Attributes: types.ElementAttributes{},
																Items: []types.UnorderedListItem{
																	{
																		Attributes:  types.ElementAttributes{},
																		Level:       1,
																		BulletStyle: types.OneAsterisk,
																		CheckStyle:  types.NoCheck,
																		Elements: []interface{}{
																			types.Paragraph{
																				Attributes: types.ElementAttributes{},
																				Lines: []types.InlineElements{
																					{
																						types.StringElement{
																							Content: "Desktop",
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
														Attributes:     types.ElementAttributes{},
														Level:          1,
														Position:       2,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "Ubuntu",
																		},
																	},
																},
															},
															types.UnorderedList{
																Attributes: types.ElementAttributes{},
																Items: []types.UnorderedListItem{
																	{
																		Attributes:  types.ElementAttributes{},
																		Level:       1,
																		BulletStyle: types.OneAsterisk,
																		CheckStyle:  types.NoCheck,
																		Elements: []interface{}{
																			types.Paragraph{
																				Attributes: types.ElementAttributes{},
																				Lines: []types.InlineElements{
																					{
																						types.StringElement{
																							Content: "Desktop",
																						},
																					},
																				},
																			},
																		},
																	},
																	{
																		Attributes:  types.ElementAttributes{},
																		Level:       1,
																		BulletStyle: types.OneAsterisk,
																		CheckStyle:  types.NoCheck,
																		Elements: []interface{}{
																			types.Paragraph{
																				Attributes: types.ElementAttributes{},
																				Lines: []types.InlineElements{
																					{
																						types.StringElement{
																							Content: "Server",
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
									{
										Attributes: types.ElementAttributes{},
										Level:      2,
										Term:       "BSD",
										Elements: []interface{}{
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          1,
														Position:       1,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "FreeBSD",
																		},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          1,
														Position:       2,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "NetBSD",
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
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "Cloud Providers",
						Elements: []interface{}{
							types.LabeledList{
								Attributes: types.ElementAttributes{},
								Items: []types.LabeledListItem{
									{
										Attributes: types.ElementAttributes{},
										Level:      2,
										Term:       "PaaS",
										Elements: []interface{}{
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          1,
														Position:       1,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "OpenShift",
																		},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          1,
														Position:       2,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "CloudBees",
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
										Attributes: types.ElementAttributes{},
										Level:      2,
										Term:       "IaaS",
										Elements: []interface{}{
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          1,
														Position:       1,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "Amazon EC2",
																		},
																	},
																},
															},
														},
													},
													{
														Attributes:     types.ElementAttributes{},
														Level:          1,
														Position:       2,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "Rackspace",
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
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("checklists", func() {

		It("checklist with title and dashes", func() {
			actualContent := `.Checklist
- [*] checked
- [x] also checked
- [ ] not checked
-     normal list item`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Checklist",
				},
				Items: []types.UnorderedListItem{
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.Checked,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{
									types.AttrCheckStyle: types.Checked,
								},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "checked",
										},
									},
								},
							},
						},
					},
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.Checked,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{
									types.AttrCheckStyle: types.Checked,
								},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "also checked",
										},
									},
								},
							},
						},
					},
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.Unchecked,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{
									types.AttrCheckStyle: types.Unchecked,
								},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "not checked",
										},
									},
								},
							},
						},
					},
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "normal list item",
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

		It("parent checklist with title and nested checklist", func() {
			actualContent := `.Checklist
* [ ] parent not checked
** [*] checked
** [x] also checked
** [ ] not checked
*     normal list item`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Checklist",
				},
				Items: []types.UnorderedListItem{
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.Unchecked,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{
									types.AttrCheckStyle: types.Unchecked,
								},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "parent not checked",
										},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.Checked,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{
													types.AttrCheckStyle: types.Checked,
												},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "checked",
														},
													},
												},
											},
										},
									},
									{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.Checked,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{
													types.AttrCheckStyle: types.Checked,
												},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "also checked",
														},
													},
												},
											},
										},
									},
									{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.Unchecked,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{
													types.AttrCheckStyle: types.Unchecked,
												},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "not checked",
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
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "normal list item",
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

		It("parent checklist with title and nested normal list", func() {
			actualContent := `.Checklist
* [ ] parent not checked
** a normal list item
** another normal list item
*     normal list item`
			expectedResult := types.UnorderedList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Checklist",
				},
				Items: []types.UnorderedListItem{
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.Unchecked,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{
									types.AttrCheckStyle: types.Unchecked,
								},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "parent not checked",
										},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "a normal list item",
														},
													},
												},
											},
										},
									},
									{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "another normal list item",
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
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "normal list item",
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

	Context("distinct list blocks", func() {

		It("same list without attributes", func() {
			actualContent := `[lowerroman, start=5]
. Five
.. a
. Six`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{ // a single ordered list
					types.OrderedList{
						Attributes: types.ElementAttributes{
							"lowerroman": nil,
							"start":      "5",
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								Position:       1,
								NumberingStyle: types.LowerRoman,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "Five",
												},
											},
										},
									},
									types.OrderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.OrderedListItem{
											{
												Attributes:     types.ElementAttributes{},
												Level:          2,
												Position:       1,
												NumberingStyle: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{
																	Content: "a",
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
								Attributes:     types.ElementAttributes{},
								Level:          1,
								Position:       2,
								NumberingStyle: types.LowerRoman,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "Six",
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})

		It("same list with custom number style on sublist", func() {
			actualContent := `[lowerroman, start=5]
. Five
[upperalpha]
.. a
.. b
. Six`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{ // a single ordered list
					types.OrderedList{
						Attributes: types.ElementAttributes{
							"lowerroman": nil,
							"start":      "5",
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								Position:       1,
								NumberingStyle: types.LowerRoman,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "Five",
												},
											},
										},
									},
									types.OrderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.OrderedListItem{
											{
												Attributes: types.ElementAttributes{
													"upperalpha": nil,
												},
												Level:          2,
												Position:       1,
												NumberingStyle: types.UpperAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{
																	Content: "a",
																},
															},
														},
													},
												},
											},
											{
												Attributes:     types.ElementAttributes{},
												Level:          2,
												Position:       2,
												NumberingStyle: types.UpperAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{
																	Content: "b",
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
								Attributes:     types.ElementAttributes{},
								Level:          1,
								Position:       2,
								NumberingStyle: types.LowerRoman,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "Six",
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})

		It("distinct lists - case 1", func() {
			actualContent := `[lowerroman, start=5]
. Five

[loweralpha]
.. a
. Six`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{ // a single ordered list
					types.OrderedList{
						Attributes: types.ElementAttributes{
							"lowerroman": nil,
							"start":      "5",
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								Position:       1,
								NumberingStyle: types.LowerRoman,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "Five",
												},
											},
										},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.OrderedList{
						Attributes: types.ElementAttributes{
							"loweralpha": nil,
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								Position:       1,
								NumberingStyle: types.LowerAlpha,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "a",
												},
											},
										},
									},
									types.OrderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.OrderedListItem{
											{
												Attributes:     types.ElementAttributes{},
												Level:          2,
												Position:       1,
												NumberingStyle: types.Arabic,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{
																	Content: "Six",
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})

		It("distinct lists - case 2", func() {

			actualContent := `.Checklist
- [*] checked
-     normal list item

.Ordered, basic
. Step 1
. Step 2`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Checklist",
						},
						Items: []types.UnorderedListItem{
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.Checked,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{
											types.AttrCheckStyle: types.Checked,
										},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "checked",
												},
											},
										},
									},
								},
							},
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "normal list item",
												},
											},
										},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Ordered, basic",
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								Position:       1,
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "Step 1",
												},
											},
										},
									},
								},
							},
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								Position:       2,
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "Step 2",
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})
	})

})
