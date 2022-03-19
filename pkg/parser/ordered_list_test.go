package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golintt
)

var _ = Describe("ordered lists", func() {

	Context("in final documents", func() {

		// same single element in the list for each test in this context
		elements := []interface{}{
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{
						Content: "element",
					},
				},
			},
		}

		Context("ordered list element alone", func() {

			It("with implicit numbering style on a single line", func() {
				source := `. element on a single line`
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
													Content: "element on a single line",
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

			It("with implicit numbering style on multiple lines with tabs", func() {
				// leading and trailing spaces must be trimmed on each line
				source := `. element 
	on 
	multiple 
	lines
`
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
													Content: "element\non\nmultiple\nlines", // spaces are trimmed
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

			It("with implicit numbering style", func() {
				source := `. element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.Arabic,
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unnecessary level and numbering style adjustments", func() {
				source := `.. element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.LowerAlpha,
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with arabic numbering style", func() {
				source := `1. element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.Arabic,
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with lower alpha numbering style", func() {
				source := `b. element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.LowerAlpha,
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with upper alpha numbering style", func() {
				source := `B. element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.UpperAlpha,
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with lower roman numbering style", func() {
				source := `i) element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.LowerRoman,
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with upper roman numbering style", func() {
				source := `I) element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.UpperRoman,
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with predefined attribute", func() {
				source := `. {amp}`
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

			It("with explicit start only", func() {
				source := `[start=5]
. element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								"start": "5",
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.Arabic,
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with explicit quoted numbering and start", func() {
				source := `["lowerroman", start="5"]
. element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								types.AttrStyle: "lowerroman", // will be used during rendering
								"start":         "5",
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.Arabic, // will be overridden during rendering
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("elements without numbers", func() {

			It("with simple unnumbered elements", func() {
				source := `. a
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
												&types.StringElement{Content: "a"},
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "b"},
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

			It("with explicit numbering style", func() {
				source := `[lowerroman]
. element
. element`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								types.AttrStyle: "lowerroman", // will be used during rendering
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style:    types.Arabic, // will be overridden during rendering
									Elements: elements,
								},
								&types.OrderedListElement{
									Style:    types.Arabic, // will be overridden during rendering
									Elements: elements,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unnumbered elements", func() {
				source := `. element 1
. element 2`

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
												&types.StringElement{Content: "element 1"},
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "element 2"},
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

			It("with custom numbering on child elements with tabs ", func() {
				// note: the [upperroman] attribute must be at the beginning of the line
				source := `. element 1
			.. element 1.1
[upperroman]
			... element 1.1.1
			... element 1.1.2
			.. element 1.2
			. element 2
			.. element 2.1`

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
												&types.StringElement{Content: "element 1"},
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
																&types.StringElement{Content: "element 1.1"},
															},
														},
														&types.List{
															Kind: types.OrderedListKind,
															Attributes: types.Attributes{
																types.AttrStyle: "upperroman",
															},
															Elements: []types.ListElement{
																&types.OrderedListElement{
																	Style: types.LowerRoman, // will be overridden during rendering
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "element 1.1.1"},
																			},
																		},
																	},
																},
																&types.OrderedListElement{
																	Style: types.LowerRoman, // will be overridden during rendering
																	Elements: []interface{}{
																		&types.Paragraph{
																			Elements: []interface{}{
																				&types.StringElement{Content: "element 1.1.2"},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												&types.OrderedListElement{
													Style: types.LowerAlpha,
													Elements: []interface{}{
														&types.Paragraph{
															Elements: []interface{}{
																&types.StringElement{Content: "element 1.2"},
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
												&types.StringElement{Content: "element 2"},
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
																&types.StringElement{Content: "element 2.1"},
															},
														},
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

			It("with all default styles and blank lines", func() {
				source := `. level 1

.. level 2


... level 3



.... level 4
..... level 5.


`
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
												&types.StringElement{Content: "level 1"},
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
																&types.StringElement{Content: "level 2"},
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
																				&types.StringElement{Content: "level 3"},
																			},
																		},
																		&types.List{
																			Kind: types.OrderedListKind,
																			Elements: []types.ListElement{
																				&types.OrderedListElement{
																					Style: types.UpperAlpha,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{Content: "level 4"},
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
																												&types.StringElement{Content: "level 5."},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
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

			It("with all default styles and no blank line", func() {
				source := `. level 1
.. level 2
... level 3
.... level 4
..... level 5.
`
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
												&types.StringElement{Content: "level 1"},
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
																&types.StringElement{Content: "level 2"},
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
																				&types.StringElement{Content: "level 3"},
																			},
																		},
																		&types.List{
																			Kind: types.OrderedListKind,
																			Elements: []types.ListElement{
																				&types.OrderedListElement{
																					Style: types.UpperAlpha,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{Content: "level 4"},
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
																												&types.StringElement{Content: "level 5."},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
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

		Context("numbered elements", func() {

			It("with simple numbered elements", func() {
				source := `1. a
2. b`
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
												&types.StringElement{Content: "a"},
											},
										},
									},
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "b"},
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

			It("max level of ordered elements - case 1", func() {
				source := `.Ordered, max nesting
. level 1
.. level 2
... level 3
.... level 4
..... level 5
. level 1`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.OrderedListKind,
							Attributes: types.Attributes{
								types.AttrTitle: "Ordered, max nesting",
							},
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "level 1",
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
																	Content: "level 2",
																},
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
																				&types.StringElement{
																					Content: "level 3",
																				},
																			},
																		},
																		&types.List{
																			Kind: types.OrderedListKind,
																			Elements: []types.ListElement{
																				&types.OrderedListElement{
																					Style: types.UpperAlpha,
																					Elements: []interface{}{
																						&types.Paragraph{
																							Elements: []interface{}{
																								&types.StringElement{
																									Content: "level 4",
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
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "level 1",
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

			It("with numbered elements", func() {
				source := `1. element 1
a. element 1.a
2. element 2
b. element 2.a`
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
												&types.StringElement{Content: "element 1"},
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
																&types.StringElement{Content: "element 1.a"},
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
												&types.StringElement{Content: "element 2"},
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
																&types.StringElement{Content: "element 2.a"},
															},
														},
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

		Context("list element continuation", func() {

			It("case 1", func() {
				source := `. foo
+
----
a delimited block
----
+
----
another delimited block
----
. bar
`
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
												&types.StringElement{Content: "foo"},
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
								&types.OrderedListElement{
									Style: types.Arabic,
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
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("case 2", func() {
				source := `. {blank}
+
----
print("one")
----
. {blank}
+
----
print("two")
----`
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
												&types.PredefinedAttribute{Name: "blank"},
											},
										},
										&types.DelimitedBlock{
											Kind: types.Listing,
											Elements: []interface{}{
												&types.StringElement{
													Content: "print(\"one\")",
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
												&types.PredefinedAttribute{Name: "blank"},
											},
										},
										&types.DelimitedBlock{
											Kind: types.Listing,
											Elements: []interface{}{
												&types.StringElement{
													Content: "print(\"two\")",
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
				// continuation with "continued element" being a list element (ie, kinda invalid/empty continuation in the middle of a list)
				source := `. element 1
+
a paragraph
. element 2
+
. element 3
`
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
													Content: "element 1",
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
								},
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "element 2",
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
													Content: "element 3",
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

			It("case 4", func() {
				source := `. cookie
+
image::cookie.png[]
+
. chocolate
+
image::chocolate.png[]`
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
													Content: "cookie",
												},
											},
										},
										&types.ImageBlock{
											Location: &types.Location{
												Path: "cookie.png",
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
													Content: "chocolate",
												},
											},
										},
										&types.ImageBlock{
											Location: &types.Location{
												Path: "chocolate.png",
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
