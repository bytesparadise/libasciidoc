package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("mixed lists", func() {

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
			// need to be aligned on first column of file
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

		It("distinct lists with blankline and item attribute - case 1", func() {
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

		It("distinct lists with blankline and item attribute - case 2", func() {

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

		It("same list with single comment line inside", func() {
			actualContent := `. a
	// -
	. b`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
													Content: "a",
												},
											},
										},
									},
									types.SingleLineComment{
										Content: " -",
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})

		It("same list with multiple comment lines inside", func() {
			actualContent := `. a
	// -
	// -
	// -
	. b`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
													Content: "a",
												},
											},
										},
									},
									types.SingleLineComment{
										Content: " -",
									},
									types.SingleLineComment{
										Content: " -",
									},
									types.SingleLineComment{
										Content: " -",
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})

		It("distinct lists separated by single comment line", func() {
			actualContent := `. a
	
	// -
	. b`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
													Content: "a",
												},
											},
										},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.SingleLineComment{
						Content: " -",
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})

		It("distinct lists separated by multiple comment lines", func() {
			actualContent := `. a
	
// -
// -
// -
. b`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
													Content: "a",
												},
											},
										},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.SingleLineComment{
						Content: " -",
					},
					types.SingleLineComment{
						Content: " -",
					},
					types.SingleLineComment{
						Content: " -",
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})
	})
})
