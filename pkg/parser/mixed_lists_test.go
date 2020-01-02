package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("mixed lists - document", func() {

	Context("valid mixed lists", func() {

		It("ordered list with nested unordered lists", func() {
			source := `. Item 1
* Item A
* Item B
. Item 2
* Item C
* Item D`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.ElementAttributes{},
						Items: []types.OrderedListItem{
							{
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     map[string]interface{}{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
														Lines: [][]interface{}{
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
														Lines: [][]interface{}{
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
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
														Lines: [][]interface{}{
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
														Lines: [][]interface{}{
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
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})
	})

	Context("complex cases", func() {

		It("complex case 1 - mixed lists", func() {
			source := `- unordered 1
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
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("complex case 2 - mixed lists", func() {
			source := `- unordered 1
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
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
																		Lines: [][]interface{}{
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
																		Lines: [][]interface{}{
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
														Lines: [][]interface{}{
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
										Lines: [][]interface{}{
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
												NumberingStyle: types.Arabic,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
												NumberingStyle: types.Arabic,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
															{
																types.StringElement{Content: "ordered 3.2"},
															},
														},
													},
													types.OrderedList{
														Attributes: types.ElementAttributes{
															types.AttrNumberingStyle: "upperroman",
														},
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
																				types.StringElement{Content: "ordered 3.2.I"},
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
												NumberingStyle: types.Arabic,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("complex case 4 - mixed lists", func() {
			source := `.Mixed
Operating Systems::
  . Fedora
    * Desktop`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LabeledList{
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
												NumberingStyle: types.Arabic,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
																		Lines: [][]interface{}{
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
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("complex case 5 - mixed lists and a paragraph", func() {
			source := `.Mixed
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
	
a paragraph
`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LabeledList{
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
																NumberingStyle: types.Arabic,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																						Lines: [][]interface{}{
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
																NumberingStyle: types.Arabic,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																						Lines: [][]interface{}{
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
																						Lines: [][]interface{}{
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
																NumberingStyle: types.Arabic,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																NumberingStyle: types.Arabic,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																NumberingStyle: types.Arabic,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																NumberingStyle: types.Arabic,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																NumberingStyle: types.Arabic,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																NumberingStyle: types.Arabic,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
					},
					// types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "a paragraph",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})
	})

	Context("distinct list blocks", func() {

		It("same list without attributes", func() {
			source := `[lowerroman, start=5]
	. Five
	.. a
	. Six`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{ // a single ordered list
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrNumberingStyle: "lowerroman",
							types.AttrStart:          "5",
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								NumberingStyle: types.Arabic, // will be overridden during rendering
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
												NumberingStyle: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
								NumberingStyle: types.Arabic, // will be overridden during rendering
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("same list with custom number style on sublist", func() {
			// need to be aligned on first column of file
			source := `[lowerroman, start=5]
. Five
[upperalpha]
.. a
.. b
. Six`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{ // a single ordered list
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrNumberingStyle: "lowerroman",
							types.AttrStart:          "5",
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								NumberingStyle: types.Arabic, // will be overridden during rendering
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "Five",
												},
											},
										},
									},
									types.OrderedList{
										Attributes: types.ElementAttributes{
											types.AttrNumberingStyle: "upperalpha",
										},
										Items: []types.OrderedListItem{
											{
												Attributes:     types.ElementAttributes{},
												Level:          2,
												NumberingStyle: types.LowerAlpha, // will be overridden during rendering
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
												NumberingStyle: types.LowerAlpha, // will be overridden during rendering
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
								NumberingStyle: types.Arabic, // will be overridden during rendering
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("distinct lists with blankline and item attribute - case 1", func() {
			source := `[lowerroman, start=5]
. Five

[upperalpha]
.. a
. Six`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{ // a single ordered list
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrNumberingStyle: "lowerroman",
							types.AttrStart:          "5",
						},
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
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrNumberingStyle: "upperalpha",
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								NumberingStyle: types.LowerAlpha,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
												NumberingStyle: types.Arabic,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("distinct lists with blankline and item attribute - case 2", func() {

			source := `.Checklist
- [*] checked
-     normal list item

.Ordered, basic
. Step 1
. Step 2`
			expected := types.Document{
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
										Lines: [][]interface{}{
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
										Lines: [][]interface{}{
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
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Ordered, basic",
						},
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
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("same list with single comment line inside", func() {
			source := `. a
// -
. b`
			expected := types.Document{
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
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
								Level:          1,
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("same list with multiple comment lines inside", func() {
			source := `. a
// -
// -
// -
. b`
			expected := types.Document{
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
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
								Level:          1,
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("distinct lists separated by single comment line", func() {
			source := `. a
	
// -
. b`
			expected := types.Document{
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
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("distinct lists separated by multiple comment lines", func() {
			source := `. a
	
// -
// -
// -
. b`
			expected := types.Document{
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
								NumberingStyle: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})
	})
})
