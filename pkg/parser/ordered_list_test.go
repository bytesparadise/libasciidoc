package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("ordered lists", func() {

	Context("ordered list item alone", func() {

		// same single item in the list for each test in this context
		elements := []interface{}{
			types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "item"},
					},
				},
			},
		}
		It("ordered list item with implicit numbering style", func() {
			actualContent := `.. item`
			expectedResult := types.OrderedListItem{
				Level:          2,
				Position:       1,
				NumberingStyle: types.LowerAlpha,
				Attributes:     map[string]interface{}{},
				Elements:       elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

		It("ordered list item with arabic numbering style", func() {
			actualContent := `1. item`
			expectedResult := types.OrderedListItem{
				Level:          1,
				Position:       1,
				NumberingStyle: types.Arabic,
				Attributes:     map[string]interface{}{},
				Elements:       elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

		It("ordered list item with lower alpha numbering style", func() {
			actualContent := `b. item`
			expectedResult := types.OrderedListItem{
				Level:          1,
				Position:       1,
				NumberingStyle: types.LowerAlpha,
				Attributes:     map[string]interface{}{},
				Elements:       elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

		It("ordered list item with upper alpha numbering style", func() {
			actualContent := `B. item`
			expectedResult := types.OrderedListItem{
				Level:          1,
				Position:       1,
				NumberingStyle: types.UpperAlpha,
				Attributes:     map[string]interface{}{},
				Elements:       elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

		It("ordered list item with lower roman numbering style", func() {
			actualContent := `i) item`
			expectedResult := types.OrderedListItem{
				Level:          1,
				Position:       1,
				NumberingStyle: types.LowerRoman,
				Attributes:     map[string]interface{}{},
				Elements:       elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

		It("ordered list item with upper roman numbering style", func() {
			actualContent := `I) item`
			expectedResult := types.OrderedListItem{
				Level:          1,
				Position:       1,
				NumberingStyle: types.UpperRoman,
				Attributes:     map[string]interface{}{},
				Elements:       elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

		It("ordered list item with explicit numbering type only", func() {
			actualContent := `[lowerroman]
. item`
			expectedResult := types.OrderedListItem{
				Level:          1,
				Position:       1,
				NumberingStyle: types.Arabic,
				Attributes: map[string]interface{}{
					"lowerroman": nil,
				},
				Elements: elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

		It("ordered list item with explicit start only", func() {
			actualContent := `[start=5]
. item`
			expectedResult := types.OrderedListItem{
				Level:          1,
				Position:       1,
				NumberingStyle: types.Arabic,
				Attributes: map[string]interface{}{
					"start": "5",
				},
				Elements: elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

		It("ordered list item with explicit quoted numbering and start", func() {
			actualContent := `["lowerroman", start="5"]
. item`
			expectedResult := types.OrderedListItem{
				Level:          1,
				Position:       1,
				NumberingStyle: types.Arabic,
				Attributes: map[string]interface{}{
					"lowerroman": nil,
					"start":      "5",
				},
				Elements: elements,
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("OrderedListItem"))
		})

	})

	Context("items without numbers", func() {

		It("ordered list with simple unnumbered items", func() {
			actualContent := `. a
. b`

			expectedResult := types.OrderedList{
				Attributes: map[string]interface{}{},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
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
								Attributes: map[string]interface{}{},
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
				Attributes: map[string]interface{}{},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
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

		It("ordered list with custom numbering on child items", func() {
			actualContent := `. item 1
.. item 1.1
[upperroman]
... item 1.1.1
... item 1.1.2
.. item 1.2
. item 2
.. item 2.1`

			expectedResult := types.OrderedList{
				Attributes: map[string]interface{}{},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
							types.OrderedList{
								Attributes: map[string]interface{}{},
								Items: []types.OrderedListItem{
									{
										Level:          2,
										Position:       1,
										NumberingStyle: types.LowerAlpha,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "item 1.1"},
													},
												},
											},
											types.OrderedList{
												Attributes: map[string]interface{}{},
												Items: []types.OrderedListItem{
													{
														Level:          3,
														Position:       1,
														NumberingStyle: types.UpperRoman,
														Attributes: map[string]interface{}{
															"upperroman": nil,
														},
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
													{
														Level:          3,
														Position:       2,
														NumberingStyle: types.UpperRoman,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
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
						Level:          1,
						Position:       2,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
							types.OrderedList{
								Attributes: map[string]interface{}{},
								Items: []types.OrderedListItem{
									{
										Level:          2,
										Position:       1,
										NumberingStyle: types.LowerAlpha,
										Attributes:     map[string]interface{}{},
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

		It("ordered list with all default styles", func() {
			actualContent := `. level 1
.. level 2
... level 3
.... level 4
..... level 5.`
			expectedResult := types.OrderedList{
				Attributes: map[string]interface{}{},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "level 1"},
									},
								},
							},
							types.OrderedList{
								Attributes: map[string]interface{}{},
								Items: []types.OrderedListItem{
									{
										Level:          2,
										Position:       1,
										NumberingStyle: types.LowerAlpha,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "level 2"},
													},
												},
											},
											types.OrderedList{
												Attributes: map[string]interface{}{},
												Items: []types.OrderedListItem{
													{
														Level:          3,
														Position:       1,
														NumberingStyle: types.LowerRoman,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "level 3"},
																	},
																},
															},
															types.OrderedList{
																Attributes: map[string]interface{}{},
																Items: []types.OrderedListItem{
																	{
																		Level:          4,
																		Position:       1,
																		NumberingStyle: types.UpperAlpha,
																		Attributes:     map[string]interface{}{},
																		Elements: []interface{}{
																			types.Paragraph{
																				Attributes: map[string]interface{}{},
																				Lines: []types.InlineElements{
																					{
																						types.StringElement{Content: "level 4"},
																					},
																				},
																			},
																			types.OrderedList{
																				Attributes: map[string]interface{}{},
																				Items: []types.OrderedListItem{
																					{
																						Level:          5,
																						Position:       1,
																						NumberingStyle: types.UpperRoman,
																						Attributes:     map[string]interface{}{},
																						Elements: []interface{}{
																							types.Paragraph{
																								Attributes: map[string]interface{}{},
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
				Attributes: map[string]interface{}{},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
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
								Attributes: map[string]interface{}{},
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
				Attributes: map[string]interface{}{},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
							types.OrderedList{
								Attributes: map[string]interface{}{},
								Items: []types.OrderedListItem{
									{
										Level:          2,
										Position:       1,
										NumberingStyle: types.LowerAlpha,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
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
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
							types.OrderedList{
								Attributes: map[string]interface{}{},
								Items: []types.OrderedListItem{
									{
										Level:          2,
										Position:       1,
										NumberingStyle: types.LowerAlpha,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
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
				Attributes: map[string]interface{}{},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "Item 1"},
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
														types.StringElement{Content: "Item A"},
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
						Level:          1,
						Position:       2,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "Item 2"},
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
														types.StringElement{Content: "Item C"},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("List"))
		})

		It("ordered list mixed with unordered list - complex case 1", func() {
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
										types.StringElement{Content: "unordered 1"},
									},
								},
							},
							types.OrderedList{
								Attributes: map[string]interface{}{},
								Items: []types.OrderedListItem{
									{
										Level:          1,
										Position:       1,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.1"},
													},
												},
											},
											types.OrderedList{
												Attributes: map[string]interface{}{},
												Items: []types.OrderedListItem{
													{
														Level:          2,
														Position:       1,
														NumberingStyle: types.LowerAlpha,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.a"},
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
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.b"},
																	},
																},
															},
														},
													},
													{
														Level:          2,
														Position:       3,
														NumberingStyle: types.LowerAlpha,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
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
										Level:          1,
										Position:       2,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.2"},
													},
												},
											},
											types.OrderedList{
												Attributes: map[string]interface{}{},
												Items: []types.OrderedListItem{
													{
														Level:          2,
														Position:       1,
														NumberingStyle: types.LowerRoman,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.2.i"},
																	},
																},
															},
														},
													},
													{
														Level:          2,
														Position:       2,
														NumberingStyle: types.LowerRoman,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
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
										Level:          1,
										Position:       3,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.3"},
													},
												},
											},
										},
									},
									{
										Level:          1,
										Position:       4,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
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
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "unordered 2"},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("List"))
		})

		It("ordered list mixed with unordered list - complex case 2", func() {
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
										types.StringElement{Content: "unordered 1"},
									},
								},
							},
							types.OrderedList{
								Attributes: map[string]interface{}{},
								Items: []types.OrderedListItem{
									{
										Level:          1,
										Position:       1,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.1"},
													},
												},
											},
											types.OrderedList{
												Attributes: map[string]interface{}{},
												Items: []types.OrderedListItem{
													{
														Level:          2,
														Position:       1,
														NumberingStyle: types.LowerAlpha,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.a"},
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
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.1.b"},
																	},
																},
															},
														},
													},
													{
														Level:          2,
														Position:       3,
														NumberingStyle: types.LowerAlpha,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
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
										Level:          1,
										Position:       2,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.2"},
													},
												},
											},
											types.OrderedList{
												Attributes: map[string]interface{}{},
												Items: []types.OrderedListItem{
													{
														Level:          2,
														Position:       1,
														NumberingStyle: types.LowerRoman,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 1.2.i"},
																	},
																},
															},
														},
													},
													{
														Level:          2,
														Position:       2,
														NumberingStyle: types.LowerRoman,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
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
										Level:          1,
										Position:       3,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 1.3"},
													},
												},
											},
										},
									},
									{
										Level:          1,
										Position:       4,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
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
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "unordered 2"},
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
														types.StringElement{Content: "unordered 2.1"},
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
														Level:       3,
														BulletStyle: types.TwoAsterisks,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
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
										Level:       2,
										BulletStyle: types.OneAsterisk,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
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
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "unordered 3"},
									},
								},
							},
							types.OrderedList{
								Attributes: map[string]interface{}{},
								Items: []types.OrderedListItem{
									{
										Level:          1,
										Position:       1,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 3.1"},
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
												Attributes: map[string]interface{}{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "ordered 3.2"},
													},
												},
											},
											types.OrderedList{
												Attributes: map[string]interface{}{},
												Items: []types.OrderedListItem{
													{
														Level:          2,
														Position:       1,
														NumberingStyle: types.UpperRoman,
														Attributes: map[string]interface{}{
															"upperroman": nil,
														},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{Content: "ordered 3.2.I"},
																	},
																},
															},
														},
													},
													{
														Level:          2,
														Position:       2,
														NumberingStyle: types.UpperRoman,
														Attributes:     map[string]interface{}{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: map[string]interface{}{},
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
										Level:          1,
										Position:       3,
										NumberingStyle: types.Arabic,
										Attributes:     map[string]interface{}{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: map[string]interface{}{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("List"))
		})
	})

})
