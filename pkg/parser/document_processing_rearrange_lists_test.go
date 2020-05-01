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
				Attributes:  types.ElementAttributes{},
				Level:       1,
				BulletStyle: types.Dash,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "unordered 1"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          1,
				NumberingStyle: types.Arabic,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.1"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          2,
				NumberingStyle: types.LowerAlpha,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.1.a"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          2,
				NumberingStyle: types.LowerAlpha,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.1.b"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          2,
				NumberingStyle: types.LowerAlpha,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.1.c"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          1,
				NumberingStyle: types.Arabic,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.2"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          2,
				NumberingStyle: types.LowerRoman,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.2.i"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          2,
				NumberingStyle: types.LowerRoman,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.2.ii"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          1,
				NumberingStyle: types.Arabic,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.3"},
							},
						},
					},
				},
			},
			types.OrderedListItem{
				Attributes:     types.ElementAttributes{},
				Level:          1,
				NumberingStyle: types.Arabic,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "ordered 1.4"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Attributes:  types.ElementAttributes{},
				Level:       1,
				BulletStyle: types.Dash,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "unordered 2"},
							},
						},
					},
				},
			},
			types.UnorderedListItem{
				Attributes:  types.ElementAttributes{},
				Level:       2,
				BulletStyle: types.OneAsterisk,
				CheckStyle:  types.NoCheck,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
								Lines: [][]interface{}{
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
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
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
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: [][]interface{}{
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
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: [][]interface{}{
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
														NumberingStyle: types.LowerAlpha,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
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
										Attributes:     types.ElementAttributes{},
										Level:          1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
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
														NumberingStyle: types.LowerRoman,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: [][]interface{}{
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
														NumberingStyle: types.LowerRoman,
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
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
										Attributes:     types.ElementAttributes{},
										Level:          1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
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
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
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
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
		result, err := rearrangeListItems(actual, false)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(expected))
	})

	It("labeled list with rich terms", func() {
		actual := []interface{}{
			types.LabeledListItem{
				Attributes: types.ElementAttributes{},
				Level:      1,
				Term: []interface{}{
					types.StringElement{
						Content: "`foo` term",
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "description 1"},
							},
						},
					},
				},
			},
			types.LabeledListItem{
				Attributes: types.ElementAttributes{},
				Level:      2,
				Term: []interface{}{
					types.StringElement{
						Content: "`bar` term",
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "description 2"},
							},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term: []interface{}{
							types.QuotedText{
								Kind: types.Monospace,
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
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "description 1"},
									},
								},
							},
							types.LabeledList{
								Attributes: types.ElementAttributes{},
								Items: []types.LabeledListItem{
									{
										Attributes: types.ElementAttributes{},
										Level:      2,
										Term: []interface{}{
											types.QuotedText{
												Kind: types.Monospace,
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
												Attributes: types.ElementAttributes{},
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
						},
					},
				},
			},
		}
		result, err := rearrangeListItems(actual, false)
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(expected))
	})
})
