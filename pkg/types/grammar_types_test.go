package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
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
					{
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
									{
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
													{
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
													{
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
									{
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

var _ = Describe("line ranges", func() {

	Context("single range", func() {
		ranges := newLineRanges(
			types.LineRange{Start: 2, End: 4},
		)

		It("should not match line 1", func() {
			Expect(ranges.Match(1)).Should(BeFalse())
		})

		It("should match line 2", func() {
			Expect(ranges.Match(2)).Should(BeTrue())
		})

		It("should not match line 5", func() {
			Expect(ranges.Match(1)).Should(BeFalse())
		})
	})

	Context("multiple ranges", func() {

		ranges := newLineRanges(
			types.LineRange{Start: 1, End: 1},
			types.LineRange{Start: 3, End: 4},
			types.LineRange{Start: 6, End: -1},
		)

		It("should match line 1", func() {
			Expect(ranges.Match(1)).Should(BeTrue())
		})

		It("should not match line 2", func() {
			Expect(ranges.Match(2)).Should(BeFalse())
		})

		It("should match line 6", func() {
			Expect(ranges.Match(6)).Should(BeTrue())
		})

		It("should match line 100", func() {
			Expect(ranges.Match(100)).Should(BeTrue())
		})
	})

})

func newLineRanges(values ...interface{}) types.LineRanges {
	return types.NewLineRanges(values...)
}

var _ = Describe("raw section title offset", func() {

	It("should apply relative positive offset", func() {
		actual := types.RawSectionTitlePrefix{
			Level:  []byte("=="),
			Spaces: []byte(" "),
		}
		expected := "=== "
		verifyLevelOffset(expected, actual, "+1")
	})
})

func verifyLevelOffset(expectation string, actual types.RawSectionTitlePrefix, levelOffset string) {
	result, err := actual.Bytes(levelOffset)
	require.NoError(GinkgoT(), err)
	assert.EqualValues(GinkgoT(), expectation, result)
}

// var _ = Describe("file inclusions", func() {

// 	DescribeTable("check asciidoc file",
// 		func(path string, expectation bool) {
// 			f := types.FileInclusion{
// 				Path: path,
// 			}
// 			Expect(f.IsAsciidoc()).Should(Equal(expectation))
// 		},
// 		Entry("foo.adoc", "foo.adoc", true),
// 		Entry("foo.asc", "foo.asc", true),
// 		Entry("foo.ad", "foo.ad", true),
// 		Entry("foo.asciidoc", "foo.asciidoc", true),
// 		Entry("foo.txt", "foo.txt", true),
// 		Entry("foo.csv", "foo.csv", false),
// 		Entry("foo.go", "foo.go", false),
// 	)
// })

var _ = Describe("Location resolution", func() {

	attrs := map[string]string{
		"includedir": "includes",
		"foo":        "bar",
	}
	DescribeTable("resolve URL",
		func(location types.Location, expectation string) {
			f := types.FileInclusion{
				Location: location,
			}
			Expect(f.Location.Resolve(attrs)).Should(Equal(expectation))
		},
		Entry("includes/file.ext", types.Location{
			types.StringElement{Content: "includes/file.ext"},
		}, "includes/file.ext"),
		Entry("./{includedir}/file.ext", types.Location{
			types.StringElement{Content: "./"},
			types.DocumentAttributeSubstitution{Name: "includedir"},
			types.StringElement{Content: "/file.ext"},
		}, "./includes/file.ext"),
		Entry("./{unknown}/file.ext", types.Location{
			types.StringElement{Content: "./"},
			types.DocumentAttributeSubstitution{Name: "unknown"},
			types.StringElement{Content: "/file.ext"},
		}, "./{unknown}/file.ext"),
	)
})
