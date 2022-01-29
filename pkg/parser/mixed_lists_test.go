package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("mixed lists", func() {

	Context("in final documents", func() {

		Context("valid mixed lists", func() {

			It("ordered list with nested unordered lists", func() {
				source := `. Item 1
* Item A
* Item B
. Item 2
* Item C
* Item D`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "Item 1"},
											},
										},
										&types.List{
											Kind: types.UnorderedListKind,
											Elements: []types.ListElement{
												&types.UnorderedListElement{
													BulletStyle: types.OneAsterisk,
													CheckStyle:  types.NoCheck,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "Item A"},
															},
														},
													},
												},
												&types.UnorderedListElement{
													BulletStyle: types.OneAsterisk,
													CheckStyle:  types.NoCheck,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "Item B"},
															},
														},
													},
												},
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "Item 2"},
											},
										},
										&types.List{
											Kind: types.UnorderedListKind,
											Elements: []types.ListElement{
												&types.UnorderedListElement{
													BulletStyle: types.OneAsterisk,
													CheckStyle:  types.NoCheck,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "Item C"},
															},
														},
													},
												},
												&types.UnorderedListElement{
													BulletStyle: types.OneAsterisk,
													CheckStyle:  types.NoCheck,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "Item D"},
															},
														},
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unordered list item and order list item with roman numbering", func() {
				source := `- unordered list item
 II) ordered list item`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "unordered list item",
												},
											},
										},
										&types.List{
											Kind: types.OrderedListKind,
											Elements: []types.ListElement{
												&types.OrderedListElement{
													Style: types.UpperRoman,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{
																	Content: "ordered list item",
																},
															},
														},
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "unordered 1"},
											},
										},
										&types.List{
											Kind: types.OrderedListKind,
											Elements: []types.ListElement{
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 1.1"},
															},
														},
														&types.List{
															Kind: types.OrderedListKind,
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.LowerAlpha,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.1.a"},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.LowerAlpha,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.1.b"},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.LowerAlpha,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.1.c"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 1.2"},
															},
														},
														&types.List{
															Kind: types.OrderedListKind,
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.LowerRoman,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.2.i"},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.LowerRoman,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.2.ii"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 1.3"},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 1.4"},
															},
														},
													},
												},
											},
										},
									},
								},
								&types.UnorderedListElement{
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "unordered 2"},
											},
										},
										&types.List{
											Kind: types.UnorderedListKind,
											Elements: []types.ListElement{
												&types.UnorderedListElement{
													BulletStyle: types.OneAsterisk,
													CheckStyle:  types.NoCheck,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "unordered 2.1"},
															},
														},
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "unordered 1"},
											},
										},
										&types.List{
											Kind: types.OrderedListKind,
											Elements: []types.ListElement{
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 1.1"},
															},
														},
														&types.List{
															Kind: types.OrderedListKind,
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.LowerAlpha,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.1.a"},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.LowerAlpha,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.1.b"},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.LowerAlpha,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.1.c"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 1.2"},
															},
														},
														&types.List{
															Kind: types.OrderedListKind,
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.LowerRoman,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.2.i"},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.LowerRoman,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 1.2.ii"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 1.3"},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 1.4"},
															},
														},
													},
												},
											},
										},
									},
								},
								&types.UnorderedListElement{
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "unordered 2"},
											},
										},
										&types.List{
											Kind: types.UnorderedListKind,
											Elements: []types.ListElement{
												&types.UnorderedListElement{
													BulletStyle: types.OneAsterisk,
													CheckStyle:  types.NoCheck,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "unordered 2.1"},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.TwoAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "unordered 2.1.1\nwith some\nextra lines."}, // leading tabs are trimmed
																			},
																		},
																	},
																},
																&types.UnorderedListElement{
																	BulletStyle: types.TwoAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "unordered 2.1.2"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												&types.UnorderedListElement{
													BulletStyle: types.OneAsterisk,
													CheckStyle:  types.NoCheck,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "unordered 2.2"},
															},
														},
													},
												},
											},
										},
									},
								},
								&types.UnorderedListElement{
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "unordered 3"},
											},
										},
										&types.List{
											Kind: types.OrderedListKind,
											Elements: []types.ListElement{
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 3.1"},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 3.2"},
															},
														},
														&types.List{
															Kind: types.OrderedListKind,
															Attributes: types.Attributes{
																types.AttrStyle: "upperroman",
															},
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.LowerAlpha,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 3.2.I"},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.LowerAlpha,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "ordered 3.2.II"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "ordered 3.3"},
															},
														},
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("complex case 4 - mixed lists", func() {
				source := `.Mixed
Operating Systems::
  . Fedora
    * Desktop`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.LabeledListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Mixed",
							},
							Elements: []types.ListElement{
								&types.LabeledListElement{
									Style: "::",
									Term: []interface{}{
										&types.StringElement{
											Content: "Operating Systems",
										},
									},
									Elements: []interface{}{
										&types.List{
											Kind: types.OrderedListKind,
											Elements: []types.ListElement{
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{
																	Content: "Fedora",
																},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.OneAsterisk,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
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
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.LabeledListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Mixed",
							},
							Elements: []types.ListElement{
								&types.LabeledListElement{
									Style: "::",
									Term: []interface{}{
										&types.StringElement{
											Content: "Operating Systems",
										},
									},
									Elements: []interface{}{
										&types.List{
											Kind: types.LabeledListKind,
											Elements: []types.ListElement{
												&types.LabeledListElement{
													Style: ":::",
													Term: []interface{}{
														&types.StringElement{
															Content: "Linux",
														},
													},
													Elements: []interface{}{
														&types.List{
															Kind: types.OrderedListKind,
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.Arabic,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
																					Content: "Fedora",
																				},
																			},
																		},
																		&types.List{
																			Kind: types.UnorderedListKind,
																			Elements: []types.ListElement{
																				&types.UnorderedListElement{
																					BulletStyle: types.OneAsterisk,
																					CheckStyle:  types.NoCheck,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{
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
																&types.OrderedListElement{
																	Style: types.Arabic,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
																					Content: "Ubuntu",
																				},
																			},
																		},
																		&types.List{
																			Kind: types.UnorderedListKind,
																			Elements: []types.ListElement{
																				&types.UnorderedListElement{
																					BulletStyle: types.OneAsterisk,
																					CheckStyle:  types.NoCheck,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{
																									Content: "Desktop",
																								},
																							},
																						},
																					},
																				},
																				&types.UnorderedListElement{
																					BulletStyle: types.OneAsterisk,
																					CheckStyle:  types.NoCheck,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{
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
												&types.LabeledListElement{
													Style: ":::",
													Term: []interface{}{
														&types.StringElement{
															Content: "BSD",
														},
													},
													Elements: []interface{}{
														&types.List{
															Kind: types.OrderedListKind,
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.Arabic,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
																					Content: "FreeBSD",
																				},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.Arabic,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
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
								&types.LabeledListElement{
									Style: "::",
									Term: []interface{}{
										&types.StringElement{
											Content: "Cloud Providers",
										},
									},
									Elements: []interface{}{
										&types.List{
											Kind: types.LabeledListKind,
											Elements: []types.ListElement{
												&types.LabeledListElement{
													Style: ":::",
													Term: []interface{}{
														&types.StringElement{
															Content: "PaaS",
														},
													},
													Elements: []interface{}{
														&types.List{
															Kind: types.OrderedListKind,
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.Arabic,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
																					Content: "OpenShift",
																				},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.Arabic,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
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
												&types.LabeledListElement{
													Style: ":::",
													Term: []interface{}{
														&types.StringElement{
															Content: "IaaS",
														},
													},
													Elements: []interface{}{
														&types.List{
															Kind: types.OrderedListKind,
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.Arabic,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
																					Content: "Amazon EC2",
																				},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.Arabic,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
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
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a paragraph",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("distinct list blocks", func() {

			It("same list without attributes", func() {
				source := `[lowerroman, start=5]
	. Five
	.. a
	. Six`
				expected := &types.Document{
					Elements: []interface{}{ // a single ordered list
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								types.AttrStyle: "lowerroman",
								types.AttrStart: "5",
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic, // will be overridden during rendering
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "Five",
												},
											},
										},
										&types.List{
											Kind: types.OrderedListKind,
											Elements: []types.ListElement{
												&types.OrderedListElement{
													Style: types.LowerAlpha,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{
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
								&types.OrderedListElement{
									Style: types.Arabic, // will be overridden during rendering
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "Six",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("same list with custom number style on sublist", func() {
				// need to be aligned on first column of file
				source := `[lowerroman, start=5]
. Five
[upperalpha]
.. a
.. b
. Six`
				expected := &types.Document{
					Elements: []interface{}{ // a single ordered list
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								types.AttrStyle: "lowerroman",
								types.AttrStart: "5",
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic, // will be overridden during rendering
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "Five",
												},
											},
										},
										&types.List{
											Kind: types.OrderedListKind,
											Attributes: types.Attributes{
												types.AttrStyle: "upperalpha",
											},
											Elements: []types.ListElement{
												&types.OrderedListElement{
													Style: types.LowerAlpha, // will be overridden during rendering
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{
																	Content: "a",
																},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.LowerAlpha, // will be overridden during rendering
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{
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
								&types.OrderedListElement{
									Style: types.Arabic, // will be overridden during rendering
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "Six",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("distinct lists with blankline and item attribute - case 1", func() {
				source := `[lowerroman, start=5]
. Five

[upperalpha]
.. a
. Six`
				expected := &types.Document{
					Elements: []interface{}{ // 2 distinct ordered lists because of `[upperalpha]`
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								types.AttrStyle: "lowerroman",
								types.AttrStart: "5",
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "Five",
												},
											},
										},
									},
								},
							},
						},
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								types.AttrStyle: "upperalpha",
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.LowerAlpha,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a",
												},
											},
										},
										&types.List{
											Kind: types.OrderedListKind,
											Elements: []types.ListElement{
												&types.OrderedListElement{
													Style: types.Arabic,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{
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
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("distinct lists with blankline and item attribute - case 2", func() {

				source := `.Checklist
- [*] checked
-     normal list item

.Ordered, basic
. Step 1
. Step 2`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Checklist",
							},
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.Dash,
									CheckStyle:  types.Checked,
									Elements: []interface{}{
										&types.Paragraph{
											Attributes: types.Attributes{
												types.AttrCheckStyle: types.Checked,
											},
											Elements: []interface{}{
												&types.StringElement{
													Content: "checked",
												},
											},
										},
									},
								},
								&types.UnorderedListElement{
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "normal list item",
												},
											},
										},
									},
								},
							},
						},
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Ordered, basic",
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "Step 1",
												},
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "Step 2",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("same list with single comment line inside", func() {
				source := `. a
// -
. b`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a",
												},
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "b",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("same list with multiple comment lines inside", func() {
				source := `. a
// -
// -
// -
. b`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a",
												},
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "b",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("distinct lists separated by single comment line", func() {
				source := `. a
	
// -
. b`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a",
												},
											},
										},
									},
								},
							},
						},
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "b",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("distinct lists separated by multiple comment lines", func() {
				source := `. a
	
// -
// -
// -
. b`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a",
												},
											},
										},
									},
								},
							},
						},
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "b",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
