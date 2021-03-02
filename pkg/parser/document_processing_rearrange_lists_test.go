package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("rearrange lists", func() {

	It("mixed lists - complex case 1", func() {
		// - unordered 1
		// 1. ordered 1.1
		// 	a. ordered 1.1.a
		// 	b. ordered 1.1.b
		// 	c. ordered 1.1.c
		// 2. ordered 1.2
		// 	i)  ordered 1.2.i
		// 	ii) ordered 1.2.ii
		// 3. ordered 1.3
		// 4. ordered 1.4
		// - unordered 2
		// * unordered 2.1
		actual := []interface{}{
			types.UnorderedListItem{
				Level:       1,
				BulletStyle: types.Dash,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "unordered 1"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 1,
				Style: types.Arabic,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.1"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 2,
				Style: types.LowerAlpha,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.1.a"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 2,
				Style: types.LowerAlpha,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.1.b"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 2,
				Style: types.LowerAlpha,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.1.c"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 1,
				Style: types.Arabic,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.2"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 2,
				Style: types.LowerRoman,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.2.i"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 2,
				Style: types.LowerRoman,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.2.ii"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 1,
				Style: types.Arabic,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.3"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Level: 1,
				Style: types.Arabic,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.4"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       1,
				BulletStyle: types.Dash,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "unordered 2"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       2,
				BulletStyle: types.OneAsterisk,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "unordered 2.1"},
							},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.UnorderedList{
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "unordered 1"},
									},
								},
							},
							types.OrderedList{
								Items: []types.OrderedListItem{
									{
										Level: 1,
										Style: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "ordered 1.1"},
													},
												},
											},
											types.OrderedList{
												Items: []types.OrderedListItem{
													{
														Level: 2,
														Style: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
																	{
																		types.StringElement{Content: "ordered 1.1.a"},
																	},
																},
															},
														},
													},
													{
														Level: 2,
														Style: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
																	{
																		types.StringElement{Content: "ordered 1.1.b"},
																	},
																},
															},
														},
													},
													{
														Level: 2,
														Style: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
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
										Level: 1,
										Style: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "ordered 1.2"},
													},
												},
											},
											types.OrderedList{
												Items: []types.OrderedListItem{
													{
														Level: 2,
														Style: types.LowerRoman,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
																	{
																		types.StringElement{Content: "ordered 1.2.i"},
																	},
																},
															},
														},
													},
													{
														Level: 2,
														Style: types.LowerRoman,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
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
										Level: 1,
										Style: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "ordered 1.3"},
													},
												},
											},
										},
									},
									{
										Level: 1,
										Style: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
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
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "unordered 2"},
									},
								},
							},
							types.UnorderedList{
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.OneAsterisk,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
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
			},
		}
		Expect(rearrangeListItems(actual, false)).To(Equal(expected))
	})

	It("labeled list with rich terms", func() {
		actual := []interface{}{
			types.LabeledListItem{
				Level: 1,
				Term: []interface{}{
					types.StringElement{
						Content: "`foo` term",
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "description 1"},
							},
						},
					},
				},
			},
			types.LabeledListItem{
				Level: 2,
				Term: []interface{}{
					types.StringElement{
						Content: "`bar` term",
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "description 2"},
							},
						},
					},
				},
			},
			types.LabeledListItem{
				Level: 2,
				Term: []interface{}{
					types.StringElement{
						Content: "icon:caution[]",
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "description 3"},
							},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.LabeledList{
				Items: []types.LabeledListItem{
					{
						Level: 1,
						Term: []interface{}{
							types.QuotedText{
								Kind: types.SingleQuoteMonospace,
								Elements: []interface{}{
									types.StringElement{
										Content: "foo",
									},
								},
							},
							types.StringElement{
								Content: " term",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "description 1"},
									},
								},
							},
							types.LabeledList{
								Items: []types.LabeledListItem{
									{
										Level: 2,
										Term: []interface{}{
											types.QuotedText{
												Kind: types.SingleQuoteMonospace,
												Elements: []interface{}{
													types.StringElement{
														Content: "bar",
													},
												},
											},
											types.StringElement{
												Content: " term",
											},
										},
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "description 2"},
													},
												},
											},
										},
									},
									{
										Level: 2,
										Term: []interface{}{
											types.Icon{
												Class: "caution",
											},
										},
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "description 3"},
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
		Expect(rearrangeListItems(actual, false)).To(Equal(expected))
	})

	It("callout list items and a block afterwards", func() {
		actual := []interface{}{
			types.CalloutListItem{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "description 1"},
							},
						},
					},
				},
			},
			types.CalloutListItem{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "description 2"},
							},
						},
					},
				},
			},
			types.ExampleBlock{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "foo",
								},
							},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.CalloutList{
				Items: []types.CalloutListItem{
					{
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "description 1"},
									},
								},
							},
						},
					},
					{
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "description 2"},
									},
								},
							},
						},
					},
				},
			},
			types.ExampleBlock{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "foo",
								},
							},
						},
					},
				},
			},
		}
		Expect(rearrangeListItems(actual, false)).To(Equal(expected))
	})

	It("unordered list items and continued list item attached to grandparent", func() {
		// * grandparent list item
		// ** parent list item
		// *** child list item
		//
		//
		// +
		// paragraph attached to parent list item

		actual := []interface{}{
			types.UnorderedListItem{
				Level:       1,
				BulletStyle: types.OneAsterisk,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "grandparent list item"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       2,
				BulletStyle: types.TwoAsterisks,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "parent list item"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       3,
				BulletStyle: types.ThreeAsterisks,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "child list item"},
							},
						},
					},
				},
			},
			types.BlankLine{},
			types.BlankLine{},
			types.ContinuedListItemElement{
				Element: types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "paragraph attached to grandparent list item"},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.UnorderedList{
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "grandparent list item"},
									},
								},
							},
							types.UnorderedList{
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "parent list item"},
													},
												},
											},
											types.UnorderedList{
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														CheckStyle:  types.NoCheck,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
																	{
																		types.StringElement{Content: "child list item"},
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
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "paragraph attached to grandparent list item"},
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(rearrangeListItems(actual, false)).To(Equal(expected))
	})

	It("unordered list items and continued list item attached to parent", func() {
		// * grandparent list item
		// ** parent list item
		// *** child list item
		//
		// +
		// paragraph attached to parent list item

		actual := []interface{}{
			types.UnorderedListItem{
				Level:       1,
				BulletStyle: types.OneAsterisk,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "grandparent list item"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       2,
				BulletStyle: types.TwoAsterisks,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "parent list item"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       3,
				BulletStyle: types.ThreeAsterisks,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "child list item"},
							},
						},
					},
				},
			},
			types.BlankLine{},
			types.ContinuedListItemElement{
				Element: types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "paragraph attached to parent list item"},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.UnorderedList{
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "grandparent list item"},
									},
								},
							},
							types.UnorderedList{
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "parent list item"},
													},
												},
											},
											types.UnorderedList{
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														CheckStyle:  types.NoCheck,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
																	{
																		types.StringElement{Content: "child list item"},
																	},
																},
															},
														},
													},
												},
											},
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "paragraph attached to parent list item"},
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
		Expect(rearrangeListItems(actual, false)).To(Equal(expected))
	})

	It("unordered list items and continued list item attached to parent", func() {
		// * grandparent list item
		// ** parent list item
		// *** child list item
		//
		// +
		// paragraph attached to parent list item

		actual := []interface{}{
			types.UnorderedListItem{
				Level:       1,
				BulletStyle: types.OneAsterisk,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "grandparent list item"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       2,
				BulletStyle: types.TwoAsterisks,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "parent list item"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       3,
				BulletStyle: types.ThreeAsterisks,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "child list item"},
							},
						},
					},
				},
			},
			types.BlankLine{},
			types.ContinuedListItemElement{
				Element: types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "paragraph attached to parent list item"},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.UnorderedList{
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "grandparent list item"},
									},
								},
							},
							types.UnorderedList{
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "parent list item"},
													},
												},
											},
											types.UnorderedList{
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														CheckStyle:  types.NoCheck,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
																	{
																		types.StringElement{Content: "child list item"},
																	},
																},
															},
														},
													},
												},
											},
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "paragraph attached to parent list item"},
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
		Expect(rearrangeListItems(actual, false)).To(Equal(expected))
	})

	It("unordered list items and continued list item attached to parent and grandparent", func() {
		// * grandparent list item
		// ** parent list item
		// *** child list item
		//
		// +
		// paragraph attached to parent list item
		//
		// +
		// paragraph attached to grandparent list item

		actual := []interface{}{
			types.UnorderedListItem{
				Level:       1,
				BulletStyle: types.OneAsterisk,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "grandparent list item"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       2,
				BulletStyle: types.TwoAsterisks,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "parent list item"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Level:       3,
				BulletStyle: types.ThreeAsterisks,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "child list item"},
							},
						},
					},
				},
			},
			types.BlankLine{},
			types.ContinuedListItemElement{
				Element: types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "paragraph attached to parent list item"},
						},
					},
				},
			},
			types.BlankLine{},
			types.ContinuedListItemElement{
				Element: types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "paragraph attached to grandparent list item"},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.UnorderedList{
				Items: []types.UnorderedListItem{
					{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "grandparent list item"},
									},
								},
							},
							types.UnorderedList{
								Items: []types.UnorderedListItem{
									{
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "parent list item"},
													},
												},
											},
											types.UnorderedList{
												Items: []types.UnorderedListItem{
													{
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														CheckStyle:  types.NoCheck,
														Elements: []interface{}{
															types.Paragraph{
																Lines: [][]interface{}{
																	{
																		types.StringElement{Content: "child list item"},
																	},
																},
															},
														},
													},
												},
											},
											types.Paragraph{
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "paragraph attached to parent list item"},
													},
												},
											},
										},
									},
								},
							},
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "paragraph attached to grandparent list item"},
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(rearrangeListItems(actual, false)).To(Equal(expected))
	})

})
