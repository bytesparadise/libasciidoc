package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golint
)

var _ = Describe("unordered lists", func() {

	Context("in final documents", func() {

		Context("valid content", func() {

			It("with a basic single item", func() {
				source := `* a list element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a list element"},
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

			It("with ID, title, role and a single item", func() {
				source := `.mytitle
[#listID]
[.myrole]
* a list element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Attributes: types.Attributes{
								types.AttrID:    "listID",
								types.AttrTitle: "mytitle",
								types.AttrRoles: types.Roles{"myrole"},
							},
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a list element"},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"listID": "mytitle",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with a title and a single item", func() {
				source := `.a title
	* a list element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a list element"},
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

			It("with 2 items with stars", func() {
				source := `* a first item
					* a second item with *bold content*`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a first item"},
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
												&types.StringElement{Content: "a second item with "},
												&types.QuotedText{
													Kind: types.SingleQuoteBold,
													Elements: []interface{}{
														&types.StringElement{Content: "bold content"},
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

			It("with multiple levels", func() {
				source := `.Unordered list title
		* list element 1
		** nested list element A
		*** nested nested list element A.1
		*** nested nested list element A.2
		** nested list element B
		*** nested nested list element B.1
		*** nested nested list element B.2
		* list element 2`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Unordered list title",
							},
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "list element 1"},
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
																&types.StringElement{Content: "nested list element A"},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "nested nested list element A.1"},
																			},
																		},
																	},
																},
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "nested nested list element A.2"},
																			},
																		},
																	},
																},
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
																&types.StringElement{Content: "nested list element B"},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "nested nested list element B.1"},
																			},
																		},
																	},
																},
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "nested nested list element B.2"},
																			},
																		},
																	},
																},
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
												&types.StringElement{Content: "list element 2"},
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

			It("with 2 items with carets", func() {
				source := "- a first item\n" +
					"- a second item with *bold content*"
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
												&types.StringElement{Content: "a first item"},
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
												&types.StringElement{Content: "a second item with "},
												&types.QuotedText{
													Kind: types.SingleQuoteBold,
													Elements: []interface{}{
														&types.StringElement{Content: "bold content"},
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

			It("with items with mixed styles", func() {
				source := `- a parent item
					* a child item
					- another parent item
					* another child item
					** with a sub child item`
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
												&types.StringElement{Content: "a parent item"},
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
																&types.StringElement{Content: "a child item"},
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
												&types.StringElement{Content: "another parent item"},
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
																&types.StringElement{Content: "another child item"},
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
																				&types.StringElement{Content: "with a sub child item"},
																			},
																		},
																	},
																},
															},
														},
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

			It("with 2 items with empty line in-between", func() {
				// fist line after list element is swallowed
				source := "* a first item\n" +
					"\n" +
					"* a second item with *bold content*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a first item"},
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
												&types.StringElement{Content: "a second item with "},
												&types.QuotedText{
													Kind: types.SingleQuoteBold,
													Elements: []interface{}{
														&types.StringElement{Content: "bold content"},
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
			It("with 2 items on multiple lines", func() {
				source := `* item 1
  on 2 lines.
* item 2
on 2 lines, too.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "item 1\non 2 lines."}, // leading spaces are trimmed
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
												&types.StringElement{Content: "item 2\non 2 lines, too."},
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
			It("unordered lists with 2 empty lines in-between", func() {
				// the first blank lines after the first list is swallowed (for the list element)
				source := "* an item in the first list\n" +
					"\n" +
					"\n" +
					"* an item in the second list"
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "an item in the first list"},
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
												&types.StringElement{Content: "an item in the second list"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected)) // parse the whole document to get 2 lists
			})

			It("with items on 3 levels", func() {
				source := `* item 1
	** item 1.1
	** item 1.2
	*** item 1.2.1
	** item 1.3
	** item 1.4
	* item 2
	** item 2.1`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "item 1"},
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
																&types.StringElement{Content: "item 1.1"},
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
																&types.StringElement{Content: "item 1.2"},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "item 1.2.1"},
																			},
																		},
																	},
																},
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
																&types.StringElement{Content: "item 1.3"},
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
																&types.StringElement{Content: "item 1.4"},
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
												&types.StringElement{Content: "item 2"},
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
																&types.StringElement{Content: "item 2.1"},
															},
														},
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

			It("max level of unordered items - case 1", func() {
				source := `.Unordered, max nesting
* level 1
** level 2
*** level 3
**** level 4
***** level 5
* level 1`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Unordered, max nesting",
							},
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "level 1"},
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
																&types.StringElement{Content: "level 2"},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "level 3"},
																			},
																		},
																		&types.List{
																			Kind: types.UnorderedListKind,
																			Elements: []types.ListElement{
																				&types.UnorderedListElement{
																					BulletStyle: types.FourAsterisks,
																					CheckStyle:  types.NoCheck,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{Content: "level 4"},
																							},
																						},
																						&types.List{
																							Kind: types.UnorderedListKind,
																							Elements: []types.ListElement{
																								&types.UnorderedListElement{
																									BulletStyle: types.FiveAsterisks,
																									CheckStyle:  types.NoCheck,
																									Elements: []interface{}{
																										&types.Paragraph{
																											Elements: []interface{}{
																												&types.StringElement{Content: "level 5"},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
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
												&types.StringElement{Content: "level 1"},
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

			It("max level of unordered items - case 2", func() {
				source := `.Unordered, max nesting
* level 1
** level 2
*** level 3
**** level 4
***** level 5
** level 2`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Unordered, max nesting",
							},
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "level 1"},
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
																&types.StringElement{Content: "level 2"},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "level 3"},
																			},
																		},
																		&types.List{
																			Kind: types.UnorderedListKind,
																			Elements: []types.ListElement{
																				&types.UnorderedListElement{
																					BulletStyle: types.FourAsterisks,
																					CheckStyle:  types.NoCheck,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{Content: "level 4"},
																							},
																						},
																						&types.List{
																							Kind: types.UnorderedListKind,
																							Elements: []types.ListElement{
																								&types.UnorderedListElement{
																									BulletStyle: types.FiveAsterisks,
																									CheckStyle:  types.NoCheck,
																									Elements: []interface{}{
																										&types.Paragraph{
																											Elements: []interface{}{
																												&types.StringElement{Content: "level 5"},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
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
																&types.StringElement{Content: "level 2"},
															},
														},
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

			It("unordered list element with predefined attribute", func() {
				source := `* {amp}`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.PredefinedAttribute{Name: "amp"},
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

		Context("invalid content", func() {
			It("with items on 3 levels and bad numbering", func() {
				source := `* item 1
					*** item 1.1
					*** item 1.1.1
					** item 1.2
					* item 2`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "item 1"},
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
																&types.StringElement{Content: "item 1.1"},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "item 1.1.1"},
																			},
																		},
																	},
																},
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
																&types.StringElement{Content: "item 1.2"},
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
												&types.StringElement{Content: "item 2"},
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

			It("invalid list element", func() {
				source := "*an invalid list element"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "*an invalid list element"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("list element continuation", func() {

			It("case 1", func() {
				source := `* foo
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
				expected := &types.Document{
					Elements: []interface{}{
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
													Content: "foo",
												},
											},
										},
										&types.DelimitedBlock{
											Kind: types.Listing,
											Elements: []interface{}{
												&types.StringElement{
													Content: "a delimited block",
												},
											},
										},
										&types.DelimitedBlock{
											Kind: types.Listing,
											Elements: []interface{}{
												&types.StringElement{
													Content: "another delimited block",
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
													Content: "bar",
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

			It("case 2", func() {
				source := `.Unordered, complex
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Unordered, complex",
							},
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "level 1"},
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
																&types.StringElement{
																	Content: "level 2",
																},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
																					Content: "level 3\nThis is a new line inside an unordered list using ",
																				},
																				&types.PredefinedAttribute{
																					Name: "plus",
																				},
																				&types.StringElement{
																					Content: " symbol.\nWe can even force content to start on a separate line\u2026\u200b",
																				},
																				&types.LineBreak{},
																				&types.StringElement{
																					Content: "\nAmazing, isn\u2019t it?",
																				},
																			},
																		},
																		&types.List{
																			Kind: types.UnorderedListKind,
																			Elements: []types.ListElement{
																				&types.UnorderedListElement{
																					BulletStyle: types.FourAsterisks,
																					CheckStyle:  types.NoCheck,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{
																									Content: "level 4",
																								},
																							},
																						},
																						// the `+` continuation produces the second paragrap below
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{
																									Content: "The ",
																								},
																								&types.PredefinedAttribute{
																									Name: "plus",
																								},
																								&types.StringElement{
																									Content: " symbol is on a new line.",
																								},
																							},
																						},
																						&types.List{
																							Kind: types.UnorderedListKind,
																							Elements: []types.ListElement{
																								&types.UnorderedListElement{
																									BulletStyle: types.FiveAsterisks,
																									CheckStyle:  types.NoCheck,
																									Elements: []interface{}{
																										&types.Paragraph{
																											Elements: []interface{}{
																												&types.StringElement{
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
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("case 3", func() {
				source := `- here
+
_there_`
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
													Content: "here",
												},
											},
										},
										&types.Paragraph{
											Elements: []interface{}{
												&types.QuotedText{
													Kind: types.SingleQuoteItalic,
													Elements: []interface{}{
														&types.StringElement{
															Content: "there",
														},
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

			It("attach to grandparent item", func() {
				source := `* grandparent list element
** parent list element
*** child list element


+
paragraph attached to grandparent list element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.UnorderedListKind,
							Elements: []types.ListElement{
								&types.UnorderedListElement{
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "grandparent list element"},
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
																&types.StringElement{Content: "parent list element"},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "child list element"},
																			},
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
												&types.StringElement{Content: "paragraph attached to grandparent list element"},
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

			It("attach to parent item", func() {
				source := `* grandparent list element
** parent list element
*** child list element

+
paragraph attached to parent list element`
				expected := &types.Document{
					Elements: []interface{}{
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
													Content: "grandparent list element",
												},
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
																&types.StringElement{
																	Content: "parent list element",
																},
															},
														},
														&types.List{
															Kind: types.UnorderedListKind,
															Elements: []types.ListElement{
																&types.UnorderedListElement{
																	BulletStyle: types.ThreeAsterisks,
																	CheckStyle:  types.NoCheck,
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{
																					Content: "child list element",
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
																	Content: "paragraph attached to parent list element",
																},
															},
														},
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

		It("without item continuation", func() {
			source := `* foo

----
a delimited block
----

* bar

----
another delimited block
----`
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.UnorderedListKind,
						Elements: []types.ListElement{
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{Content: "foo"},
										},
									},
								},
							},
						},
					},
					&types.DelimitedBlock{
						Kind: types.Listing,
						Elements: []interface{}{
							&types.StringElement{
								Content: "a delimited block",
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
											&types.StringElement{Content: "bar"},
										},
									},
								},
							},
						},
					},
					&types.DelimitedBlock{
						Kind: types.Listing,
						Elements: []interface{}{
							&types.StringElement{
								Content: "another delimited block",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

	})
})
