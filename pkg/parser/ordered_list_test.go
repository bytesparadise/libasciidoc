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

		It("max level of ordered items - case 1", func() {
			actualContent := `.Ordered, max nesting
. level 1
.. level 2
... level 3
.... level 4
..... level 5
. level 1`
			expectedResult := types.OrderedList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Ordered, max nesting",
				},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     types.ElementAttributes{},
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
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{
										Level:          2,
										Position:       1,
										NumberingStyle: types.LowerAlpha,
										Attributes:     types.ElementAttributes{},
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
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Level:          3,
														Position:       1,
														NumberingStyle: types.LowerRoman,
														Attributes:     types.ElementAttributes{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "level 3",
																		},
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
																		Attributes:     types.ElementAttributes{},
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
																			types.OrderedList{
																				Attributes: types.ElementAttributes{},
																				Items: []types.OrderedListItem{
																					{
																						Level:          5,
																						Position:       1,
																						NumberingStyle: types.UpperRoman,
																						Attributes:     types.ElementAttributes{},
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
					{
						Level:          1,
						Position:       2,
						NumberingStyle: types.Arabic,
						Attributes:     types.ElementAttributes{},
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
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("max level of ordered items - case 2", func() {
			actualContent := `.Ordered, max nesting
. level 1
.. level 2
... level 3
.... level 4
..... level 5
.. level 2`
			expectedResult := types.OrderedList{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "Ordered, max nesting",
				},
				Items: []types.OrderedListItem{
					{
						Level:          1,
						Position:       1,
						NumberingStyle: types.Arabic,
						Attributes:     types.ElementAttributes{},
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
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{
										Level:          2,
										Position:       1,
										NumberingStyle: types.LowerAlpha,
										Attributes:     types.ElementAttributes{},
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
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Level:          3,
														Position:       1,
														NumberingStyle: types.LowerRoman,
														Attributes:     types.ElementAttributes{},
														Elements: []interface{}{
															types.Paragraph{
																Attributes: types.ElementAttributes{},
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "level 3",
																		},
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
																		Attributes:     types.ElementAttributes{},
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
																			types.OrderedList{
																				Attributes: types.ElementAttributes{},
																				Items: []types.OrderedListItem{
																					{
																						Level:          5,
																						Position:       1,
																						NumberingStyle: types.UpperRoman,
																						Attributes:     types.ElementAttributes{},
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
									{
										Level:          2,
										Position:       2,
										NumberingStyle: types.LowerAlpha,
										Attributes:     types.ElementAttributes{},
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
})
