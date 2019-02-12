package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("lists", func() {

	Context("unordered list", func() {

		It("multi-level list", func() {
			// // given
			elements := []interface{}{
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.Dash,
					Elements: []interface{}{
						types.StringElement{
							Content: "item 1",
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       2,
					BulletStyle: types.OneAsterisk,
					Elements: []interface{}{
						types.StringElement{
							Content: "item 1.1",
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.Dash,
					Elements: []interface{}{
						types.StringElement{
							Content: "item 2",
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.UnorderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.Dash,
						Elements: []interface{}{
							types.StringElement{
								Content: "item 1",
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.OneAsterisk,
										Elements: []interface{}{
											types.StringElement{
												Content: "item 1.1",
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
						Elements: []interface{}{
							types.StringElement{
								Content: "item 2",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})

	})

	Context("labeled list", func() {
		It("labeled list with 3 items", func() {
			// // given
			elements := []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item 1",
					Elements: []interface{}{
						types.StringElement{
							Content: "item 1",
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item 2",
					Elements: []interface{}{
						types.StringElement{
							Content: "item 2",
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item 3",
					Elements: []interface{}{
						types.StringElement{
							Content: "item 3",
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "item 1",
						Elements: []interface{}{
							types.StringElement{
								Content: "item 1",
							},
						},
					},
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "item 2",
						Elements: []interface{}{
							types.StringElement{
								Content: "item 2",
							},
						},
					},
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "item 3",
						Elements: []interface{}{
							types.StringElement{
								Content: "item 3",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})
	})

	Context("mixed lists", func() {

		It("labeled list with unordered sublist", func() {
			// given
			elements := []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item A",
					Elements: []interface{}{
						types.StringElement{
							Content: "item A",
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.Dash,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.1",
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       2,
					BulletStyle: types.OneAsterisk,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.1.1",
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.Dash,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.2",
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item B",
					Elements: []interface{}{
						types.StringElement{
							Content: "item B",
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item C",
					Elements: []interface{}{
						types.StringElement{
							Content: "item C",
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "item A",
						Elements: []interface{}{
							types.StringElement{
								Content: "item A",
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									{
										Attributes:  types.ElementAttributes{},
										Level:       1,
										BulletStyle: types.Dash,
										Elements: []interface{}{
											types.StringElement{
												Content: "item A.1",
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Attributes:  types.ElementAttributes{},
														Level:       2,
														BulletStyle: types.OneAsterisk,
														Elements: []interface{}{
															types.StringElement{
																Content: "item A.1.1",
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
										Elements: []interface{}{
											types.StringElement{
												Content: "item A.2",
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
						Term:       "item B",
						Elements: []interface{}{
							types.StringElement{
								Content: "item B",
							},
						},
					},
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "item C",
						Elements: []interface{}{
							types.StringElement{
								Content: "item C",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})

		It("mixed lists - case 2", func() {
			// // given
			elements := []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item A",
					Elements: []interface{}{
						types.StringElement{
							Content: "item A",
						},
					},
				},
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          1,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.1",
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.1.1",
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "item A",
						Elements: []interface{}{
							types.StringElement{
								Content: "item A",
							},
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{

										Attributes:     types.ElementAttributes{},
										Level:          1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.StringElement{
												Content: "item A.1",
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Attributes:  types.ElementAttributes{},
														Level:       1,
														BulletStyle: types.OneAsterisk,
														Elements: []interface{}{
															types.StringElement{
																Content: "item A.1.1",
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
			verify(GinkgoT(), expectation, actual)
		})

		It("mixed lists - case 3", func() {
			// // given
			elements := []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item A",
					Elements: []interface{}{
						types.StringElement{
							Content: "item A",
						},
					},
				},
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          1,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.1",
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					Elements: []interface{}{
						types.StringElement{
							Content: "item A.1.1",
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item B",
					Elements: []interface{}{
						types.StringElement{
							Content: "item B",
						},
					},
				},
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          1,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.StringElement{
							Content: "item B.1",
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					Elements: []interface{}{
						types.StringElement{
							Content: "item B.1.1",
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "item A",
						Elements: []interface{}{
							types.StringElement{
								Content: "item A",
							},
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.StringElement{
												Content: "item A.1",
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Attributes:  types.ElementAttributes{},
														Level:       1,
														BulletStyle: types.OneAsterisk,
														Elements: []interface{}{
															types.StringElement{
																Content: "item A.1.1",
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
						Term:       "item B",
						Elements: []interface{}{
							types.StringElement{
								Content: "item B",
							},
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									{
										Attributes:     types.ElementAttributes{},
										Level:          1,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.StringElement{
												Content: "item B.1",
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													{
														Attributes:  types.ElementAttributes{},
														Level:       1,
														BulletStyle: types.OneAsterisk,
														Elements: []interface{}{
															types.StringElement{
																Content: "item B.1.1",
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
			verify(GinkgoT(), expectation, actual)
		})
	})

	Context("list continuations", func() {

		It("attach to parent ancestor of same kind", func() {
			elements := []interface{}{
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          1,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1",
									},
								},
							},
						},
					},
				},
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          2,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1",
									},
								},
							},
						},
					},
				},
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          3,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1.1",
									},
								},
							},
						},
						types.ContinuedListElement{
							Offset: -1,
							Element: types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1.1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.OrderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.OrderedListItem{
					{
						Attributes:     types.ElementAttributes{},
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1",
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
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.1",
														},
													},
												},
											},
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          3,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "item 1.1.1",
																		},
																	},
																},
															},
														},
													},
												},
											},
											types.Paragraph{
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.1 continuation",
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
			verify(GinkgoT(), expectation, actual)
		})

		It("attach to grand parent ancestor of OrderedListItem kind", func() {
			elements := []interface{}{
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          1,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1",
									},
								},
							},
						},
					},
				},
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          2,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1",
									},
								},
							},
						},
					},
				},
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          3,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1.1",
									},
								},
							},
						},
						types.ContinuedListElement{
							Offset: -2,
							Element: types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.OrderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.OrderedListItem{
					{
						Attributes:     types.ElementAttributes{},
						Position:       0,
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1",
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
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.1",
														},
													},
												},
											},
											types.OrderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.OrderedListItem{
													{
														Attributes:     types.ElementAttributes{},
														Level:          3,
														Position:       0,
														NumberingStyle: types.Arabic,
														Elements: []interface{}{
															types.Paragraph{
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "item 1.1.1",
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
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})
		It("attach to grand parent ancestor of LabeledListItem kind", func() {
			elements := []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "item 1",
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1",
									},
								},
							},
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      2,
					Term:       "item 1.1",
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1",
									},
								},
							},
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      3,
					Term:       "item 1.1.1",
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1.1",
									},
								},
							},
						},
						types.ContinuedListElement{
							Offset: -2,
							Element: types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.LabeledList{
				Attributes: types.ElementAttributes{},
				Items: []types.LabeledListItem{
					types.LabeledListItem{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Term:       "item 1",
						Elements: []interface{}{
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1",
										},
									},
								},
							},
							types.LabeledList{
								Attributes: types.ElementAttributes{},
								Items: []types.LabeledListItem{
									types.LabeledListItem{
										Attributes: types.ElementAttributes{},
										Level:      2,
										Term:       "item 1.1",
										Elements: []interface{}{
											types.Paragraph{
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.1",
														},
													},
												},
											},
											types.LabeledList{
												Attributes: types.ElementAttributes{},
												Items: []types.LabeledListItem{
													types.LabeledListItem{
														Attributes: types.ElementAttributes{},
														Level:      3,
														Term:       "item 1.1.1",
														Elements: []interface{}{
															types.Paragraph{
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "item 1.1.1",
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
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})

		It("attach to grand parent ancestor of UnorderedListItem kind", func() {
			elements := []interface{}{
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					CheckStyle:  types.NoCheck,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1",
									},
								},
							},
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.TwoAsterisks,
					CheckStyle:  types.NoCheck,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1",
									},
								},
							},
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.ThreeAsterisks,
					CheckStyle:  types.NoCheck,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1.1",
									},
								},
							},
						},
						types.ContinuedListElement{
							Offset: -2,
							Element: types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.UnorderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.UnorderedListItem{
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1",
										},
									},
								},
							},
							types.UnorderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.UnorderedListItem{
									types.UnorderedListItem{
										Attributes:  types.ElementAttributes{},
										Level:       2,
										BulletStyle: types.TwoAsterisks,
										CheckStyle:  types.NoCheck,
										Elements: []interface{}{
											types.Paragraph{
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.1",
														},
													},
												},
											},
											types.UnorderedList{
												Attributes: types.ElementAttributes{},
												Items: []types.UnorderedListItem{
													types.UnorderedListItem{
														Attributes:  types.ElementAttributes{},
														Level:       3,
														BulletStyle: types.ThreeAsterisks,
														CheckStyle:  types.NoCheck,
														Elements: []interface{}{
															types.Paragraph{
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "item 1.1.1",
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
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})

		It("attach to grand parent ancestor of different kind", func() {
			elements := []interface{}{
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          1,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1",
									},
								},
							},
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					CheckStyle:  types.NoCheck,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1",
									},
								},
							},
						},
					},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					CheckStyle:  types.NoCheck,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.2",
									},
								},
							},
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "term",
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.2.1",
									},
								},
							},
						},
						types.ContinuedListElement{
							Offset: -1,
							Element: types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1.2 continuation",
										},
									},
								},
							},
						},
						types.ContinuedListElement{
							Offset: -2,
							Element: types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
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
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 2",
									},
								},
							},
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.OrderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.OrderedListItem{
					{
						Attributes:     types.ElementAttributes{},
						Position:       0,
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1",
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
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.1",
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
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.2",
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
														Term:       "term",
														Elements: []interface{}{
															types.Paragraph{
																Lines: []types.InlineElements{
																	{
																		types.StringElement{
																			Content: "item 1.2.1",
																		},
																	},
																},
															},
														},
													},
												},
											},
											types.Paragraph{
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.2 continuation",
														},
													},
												},
											},
										},
									},
								},
							},
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
					{
						Attributes:     types.ElementAttributes{},
						Position:       1,
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 2",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})

		It("attach to ancestor with over offset", func() {
			elements := []interface{}{
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          1,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1",
									},
								},
							},
						},
					},
				},
				types.OrderedListItem{
					Attributes:     types.ElementAttributes{},
					Level:          2,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						types.Paragraph{
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "item 1.1",
									},
								},
							},
						},
						types.ContinuedListElement{
							Offset: -2,
							Element: types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			// when
			actual, err := types.NewList(elements)
			// then
			require.NoError(GinkgoT(), err)
			expectation := types.OrderedList{
				Attributes: types.ElementAttributes{},
				Items: []types.OrderedListItem{
					{
						Attributes:     types.ElementAttributes{},
						Position:       0,
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1",
										},
									},
								},
							},
							types.OrderedList{
								Attributes: types.ElementAttributes{},
								Items: []types.OrderedListItem{
									types.OrderedListItem{
										Attributes:     types.ElementAttributes{},
										Position:       0,
										Level:          2,
										NumberingStyle: types.Arabic,
										Elements: []interface{}{
											types.Paragraph{
												Lines: []types.InlineElements{
													{
														types.StringElement{
															Content: "item 1.1",
														},
													},
												},
											},
										},
									},
								},
							},
							types.Paragraph{
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "item 1 continuation",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectation, actual)
		})
	})
})

func verify(t GinkgoTInterface, expectation, actual interface{}) {
	t.Logf("actual document: `%s`", spew.Sdump(actual))
	t.Logf("expected document: `%s`", spew.Sdump(expectation))
	assert.EqualValues(t, expectation, actual)
}
