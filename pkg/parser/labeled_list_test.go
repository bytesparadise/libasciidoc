package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("labeled lists", func() {

	Context("draft documents", func() {

		It("with a term and description on 2 lines", func() {
			source := `Item1::
Item 1 description
on 2 lines`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "Item 1 description"},
									},
									{
										types.StringElement{Content: "on 2 lines"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with a single term and no description", func() {
			source := `Item1::`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Term: []interface{}{
							types.StringElement{
								Content: "Item1",
							},
						},
						Level:    1,
						Elements: []interface{}{},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with a quoted text in term and in description", func() {
			source := "`foo()`::\n" +
				`This function is _untyped_.`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Term: []interface{}{
							types.StringElement{
								Content: "`foo()`", // the term is a raw string in the DraftDocument
							},
						},
						Level: 1,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "This function is ",
										},
										types.QuotedText{
											Kind: types.Italic,
											Elements: []interface{}{
												types.StringElement{
													Content: "untyped",
												},
											},
										},
										types.StringElement{
											Content: ".",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with a horizontal layout attribute", func() {
			source := `[horizontal]
Item1:: foo`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Attributes: types.Attributes{
							"style": "horizontal",
						},
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "foo"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with a title attribute", func() {
			source := `[title="Fighters"]
Item1:: foo`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Attributes: types.Attributes{
							types.AttrTitle: "Fighters",
						},
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "foo"},
									},
								},
							},
						},
					},
				},
			}
			result, err := ParseDraftDocument(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDraftDocument(expected))
		})

		It("with a single term and a blank line", func() {
			source := `Item1::
			`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item1",
							},
						},
						Elements: []interface{}{},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with multiple sibling items", func() {
			source := `Item 1::
Item 1 description
Item 2:: 
Item 2 description
Item 3:: 
Item 3 description`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "Item 1 description"},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 2",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "Item 2 description"},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 3",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "Item 3 description"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with multiple nested items", func() {
			source := `Item 1::
Item 1 description
Item 2:::
Item 2 description
Item 3::::
Item 3 description`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "Item 1 description"},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 2,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 2",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "Item 2 description"},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 3,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 3",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "Item 3 description"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with nested unordered list - case 1", func() {
			source := `Empty item:: 
* foo
* bar
Item with description:: something simple`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Empty item",
							},
						},
						Elements: []interface{}{},
					},
					types.UnorderedListItem{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "foo"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "bar"},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item with description",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "something simple"},
									},
								},
							},
						},
					},
				},
			}

			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with a single item and paragraph", func() {
			source := `Item 1::
foo
bar

a normal paragraph.`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "foo"},
									},
									{
										types.StringElement{Content: "bar"},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a normal paragraph."},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with item continuation", func() {
			source := `Item 1::
+
----
a fenced block
----
Item 2:: something simple
+
----
another fenced block
----`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 1",
							},
						},
						Elements: []interface{}{},
					},
					// the `+` continuation produces the element below
					types.ContinuedListItemElement{
						Offset: 0,
						Element: types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a fenced block",
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 2",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "something simple",
										},
									},
								},
							},
						},
					},
					// the `+` continuation produces the second element below
					types.ContinuedListItemElement{
						Offset: 0,
						Element: types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "another fenced block",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("without item continuation", func() {
			source := `Item 1::
----
a fenced block
----
Item 2:: something simple
----
another fenced block
----`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 1",
							},
						},
						Elements: []interface{}{},
					},
					types.ListingBlock{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "a fenced block",
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Item 2",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "something simple"},
									},
								},
							},
						},
					},
					types.ListingBlock{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "another fenced block",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with nested unordered list - case 2", func() {
			source := `Labeled item::
- unordered item`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "Labeled item",
							},
						},
						Elements: []interface{}{},
					},
					types.UnorderedListItem{
						Level:       1,
						BulletStyle: types.Dash,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "unordered item"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with title", func() {
			source := `.Labeled, single-line
first term:: definition of the first term
second term:: definition of the second term`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Attributes: types.Attributes{
							types.AttrTitle: "Labeled, single-line",
						},
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "first term",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "definition of the first term",
										},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "second term",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "definition of the second term",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("max level of labeled items - case 1", func() {
			source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 1:: description 1`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Attributes: types.Attributes{
							types.AttrTitle: "Labeled, max nesting",
						},
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "level 1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "description 1",
										},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 2,
						Term: []interface{}{
							types.StringElement{
								Content: "level 2",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "description 2",
										},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 3,
						Term: []interface{}{
							types.StringElement{
								Content: "level 3",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "description 3",
										},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "level 1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "description 1",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("max level of labeled items - case 2", func() {
			source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 2::: description 2`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.LabeledListItem{
						Attributes: types.Attributes{
							types.AttrTitle: "Labeled, max nesting",
						},
						Level: 1,
						Term: []interface{}{
							types.StringElement{
								Content: "level 1",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "description 1",
										},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 2,
						Term: []interface{}{
							types.StringElement{
								Content: "level 2",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "description 2",
										},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 3,
						Term: []interface{}{
							types.StringElement{
								Content: "level 3",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "description 3",
										},
									},
								},
							},
						},
					},
					types.LabeledListItem{
						Level: 2,
						Term: []interface{}{
							types.StringElement{
								Content: "level 2",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "description 2",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})

	Context("final documents", func() {

		It("with a term and description on 2 lines", func() {
			source := `Item1::
Item 1 description
on 2 lines`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item1",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "Item 1 description"},
											},
											{
												types.StringElement{Content: "on 2 lines"},
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

		It("with a single term and no description", func() {
			source := `Item1::`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Term: []interface{}{
									types.StringElement{
										Content: "Item1",
									},
								},

								Level:    1,
								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with a quoted text in term and in description", func() {
			source := "`foo()`::\n" +
				`This function is _untyped_.`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Term: []interface{}{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{
												Content: "foo()",
											},
										},
									},
								},
								Level: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "This function is ",
												},
												types.QuotedText{
													Kind: types.Italic,
													Elements: []interface{}{
														types.StringElement{
															Content: "untyped",
														},
													},
												},
												types.StringElement{
													Content: ".",
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
		It("with a index term", func() {
			source := "((`foo`))::\n" +
				`This function is _untyped_.`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Term: []interface{}{
									types.IndexTerm{
										Term: []interface{}{
											types.QuotedText{
												Kind: types.Monospace,
												Elements: []interface{}{
													types.StringElement{
														Content: "foo",
													},
												},
											},
										},
									},
								},
								Level: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "This function is ",
												},
												types.QuotedText{
													Kind: types.Italic,
													Elements: []interface{}{
														types.StringElement{
															Content: "untyped",
														},
													},
												},
												types.StringElement{
													Content: ".",
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

		It("with a concealed index term in term", func() {
			source := "(((foo,bar)))::\n" +
				`This function is _untyped_.`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Term: []interface{}{
									types.ConcealedIndexTerm{
										Term1: "foo",
										Term2: "bar",
									},
								},
								Level: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "This function is ",
												},
												types.QuotedText{
													Kind: types.Italic,
													Elements: []interface{}{
														types.StringElement{
															Content: "untyped",
														},
													},
												},
												types.StringElement{
													Content: ".",
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

		It("with a horizontal layout attribute", func() {
			source := `[horizontal]
Item1:: foo`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Attributes: types.Attributes{
							"style": "horizontal",
						},
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item1",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "foo"},
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

		It("with a single term and a blank line", func() {
			source := `Item1::
			`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item1",
									},
								},

								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with multiple sibling items", func() {
			source := `Item 1::
Item 1 description
Item 2:: 
Item 2 description
Item 3:: 
Item 3 description`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 1",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "Item 1 description"},
											},
										},
									},
								},
							},
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 2",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "Item 2 description"},
											},
										},
									},
								},
							},
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 3",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "Item 3 description"},
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

		It("with multiple nested items", func() {
			source := `Item 1::
Item 1 description
Item 2:::
Item 2 description
Item 3::::
Item 3 description`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 1",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "Item 1 description"},
											},
										},
									},
									types.LabeledList{
										Items: []types.LabeledListItem{
											{
												Level: 2,
												Term: []interface{}{
													types.StringElement{
														Content: "Item 2",
													},
												},

												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{Content: "Item 2 description"},
															},
														},
													},
													types.LabeledList{
														Items: []types.LabeledListItem{
															{
																Level: 3,
																Term: []interface{}{
																	types.StringElement{
																		Content: "Item 3",
																	},
																},

																Elements: []interface{}{
																	types.Paragraph{
																		Lines: [][]interface{}{
																			{
																				types.StringElement{Content: "Item 3 description"},
																			},
																		},
																	},
																},
															},
														},
													},
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

		It("with nested unordered list - case 1", func() {
			source := `Empty item:: 
* foo
* bar
Item with description:: something simple`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Empty item",
									},
								},

								Elements: []interface{}{
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
																types.StringElement{Content: "foo"},
															},
														},
													},
												},
											},
											{
												Level:       1,
												BulletStyle: types.OneAsterisk,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{Content: "bar"},
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
								Term: []interface{}{
									types.StringElement{
										Content: "Item with description",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "something simple"},
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

		It("with a single item and paragraph", func() {
			source := `Item 1::
foo
bar

a normal paragraph.`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 1",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "foo"},
											},
											{
												types.StringElement{Content: "bar"},
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
								types.StringElement{Content: "a normal paragraph."},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with item continuation", func() {
			source := `Item 1::
+
----
a fenced block
----
Item 2:: something simple
+
----
another fenced block
----`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 1",
									},
								},

								Elements: []interface{}{
									types.ListingBlock{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a fenced block",
												},
											},
										},
									},
								},
							},
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 2",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "something simple"},
											},
										},
									},
									types.ListingBlock{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "another fenced block",
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

		It("without item continuation", func() {
			source := `Item 1::
----
a fenced block
----
Item 2:: something simple
----
another fenced block
----`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 1",
									},
								},

								Elements: []interface{}{},
							},
						},
					},
					types.ListingBlock{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "a fenced block",
								},
							},
						},
					},
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 2",
									},
								},

								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "something simple"},
											},
										},
									},
								},
							},
						},
					},
					types.ListingBlock{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "another fenced block",
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with nested unordered list - case 2", func() {
			source := `Labeled item::
- unordered item`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Labeled item",
									},
								},
								Elements: []interface{}{
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
																types.StringElement{Content: "unordered item"},
															},
														},
													},
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

		It("with title", func() {
			source := `.Labeled, single-line
first term:: definition of the first term
second term:: definition of the second term`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Attributes: types.Attributes{
							types.AttrTitle: "Labeled, single-line",
						},
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "first term",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "definition of the first term",
												},
											},
										},
									},
								},
							},
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "second term",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "definition of the second term",
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
			result, err := ParseDocument(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDocument(expected))
		})

		It("max level of labeled items - case 1", func() {
			source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 1:: description 1`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Attributes: types.Attributes{
							types.AttrTitle: "Labeled, max nesting",
						},
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "level 1",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "description 1",
												},
											},
										},
									},
									types.LabeledList{
										Items: []types.LabeledListItem{
											{
												Level: 2,
												Term: []interface{}{
													types.StringElement{
														Content: "level 2",
													},
												},
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{
																	Content: "description 2",
																},
															},
														},
													},
													types.LabeledList{
														Items: []types.LabeledListItem{
															{
																Level: 3,
																Term: []interface{}{
																	types.StringElement{
																		Content: "level 3",
																	},
																},
																Elements: []interface{}{
																	types.Paragraph{
																		Lines: [][]interface{}{
																			{
																				types.StringElement{
																					Content: "description 3",
																				},
																			},
																		},
																	},
																},
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
								Term: []interface{}{
									types.StringElement{
										Content: "level 1",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "description 1",
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

		It("max level of labeled items - case 2", func() {
			source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 2::: description 2`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Attributes: types.Attributes{
							types.AttrTitle: "Labeled, max nesting",
						},
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "level 1",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "description 1",
												},
											},
										},
									},
									types.LabeledList{
										Items: []types.LabeledListItem{
											{
												Level: 2,
												Term: []interface{}{
													types.StringElement{
														Content: "level 2",
													},
												},
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{
																	Content: "description 2",
																},
															},
														},
													},
													types.LabeledList{
														Items: []types.LabeledListItem{
															{
																Level: 3,
																Term: []interface{}{
																	types.StringElement{
																		Content: "level 3",
																	},
																},
																Elements: []interface{}{
																	types.Paragraph{
																		Lines: [][]interface{}{
																			{
																				types.StringElement{
																					Content: "description 3",
																				},
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
												Level: 2,
												Term: []interface{}{
													types.StringElement{
														Content: "level 2",
													},
												},
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{
																	Content: "description 2",
																},
															},
														},
													},
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

		It("item with predefined attribute", func() {
			source := `level 1:: {amp}`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "level 1",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.PredefinedAttribute{Name: "amp"},
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

		It("item with a colon the term", func() {
			source := `what: ever:: text`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "what: ever",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "text"},
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

		It("with multiple item continuations", func() {
			source := `Item 1::
content 1
+
NOTE: note

Item 2::
content 2
+
addition
+
IMPORTANT: important
+
TIP: tip`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 1",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "content 1"},
											},
										},
									},
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrStyle: types.Note,
										},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "note"},
											},
										},
									},
								},
							},
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "Item 2",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "content 2"},
											},
										},
									},
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "addition"},
											},
										},
									},
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrStyle: types.Important,
										},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "important"},
											},
										},
									},
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrStyle: types.Tip,
										},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "tip"},
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

		It("with list item continuations", func() {
			source := `item::
This is the first line of the first paragraph.
This is the second line of the first paragraph.
+
This is the first line of the continuation paragraph.
This is the second line of the continuation paragraph.
+
This is the next continuation paragraph.
+
TIP: We can embed admonitions too!
`
			expected := types.Document{
				Elements: []interface{}{
					types.LabeledList{
						Items: []types.LabeledListItem{
							{
								Level: 1,
								Term: []interface{}{
									types.StringElement{
										Content: "item",
									},
								},
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "This is the first line of the first paragraph.",
												},
											},
											{
												types.StringElement{
													Content: "This is the second line of the first paragraph.",
												},
											},
										},
									},
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "This is the first line of the continuation paragraph.",
												},
											},
											{
												types.StringElement{
													Content: "This is the second line of the continuation paragraph.",
												},
											},
										},
									},
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "This is the next continuation paragraph.",
												},
											},
										},
									},
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrStyle: types.Tip,
										},
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "We can embed admonitions too!",
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

})
