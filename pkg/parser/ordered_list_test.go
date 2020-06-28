package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("ordered lists - draft", func() {

	Context("ordered list item alone", func() {

		// same single item in the list for each test in this context
		elements := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "item"},
					},
				},
			},
		}

		It("ordered list item with implicit numbering style", func() {
			source := `.. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:    2,
						Style:    types.LowerAlpha,
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list item with arabic numbering style", func() {
			source := `1. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:    1,
						Style:    types.Arabic,
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list item with lower alpha numbering style", func() {
			source := `b. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:    1,
						Style:    types.LowerAlpha,
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list item with upper alpha numbering style", func() {
			source := `B. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:    1,
						Style:    types.UpperAlpha,
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list item with lower roman numbering style", func() {
			source := `i) item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:    1,
						Style:    types.LowerRoman,
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list item with upper roman numbering style", func() {
			source := `I) item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:    1,
						Style:    types.UpperRoman,
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list item with explicit numbering style", func() {
			source := `[lowerroman]
. item
. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Attributes: types.Attributes{
							types.AttrStyle: types.LowerRoman,
						},
						Level:    1,
						Style:    types.Arabic,
						Elements: elements,
					},
					types.OrderedListItem{
						Level:    1,
						Style:    types.Arabic,
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list item with explicit start only", func() {
			source := `[start=5]
. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Attributes: types.Attributes{
							"start": "5",
						},
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list item with explicit quoted numbering and start", func() {
			source := `["lowerroman", start="5"]
. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Attributes: types.Attributes{
							types.AttrStyle: types.LowerRoman,
							"start":         "5",
						},
						Elements: elements,
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("max level of ordered items - case 1", func() {
			source := `.Ordered, max nesting
. level 1
.. level 2
... level 3
.... level 4
..... level 5
. level 1`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Attributes: types.Attributes{
							types.AttrTitle: "Ordered, max nesting",
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 1",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 2,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 2",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 3,
						Style: types.LowerRoman,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 3",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 4,
						Style: types.UpperAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 4",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 5,
						Style: types.UpperRoman,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 5",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("max level of ordered items - case 2", func() {
			source := `.Ordered, max nesting
. level 1
.. level 2
... level 3
.... level 4
..... level 5
.. level 2b`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Attributes: types.Attributes{
							types.AttrTitle: "Ordered, max nesting",
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 1",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 2,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 2",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 3,
						Style: types.LowerRoman,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 3",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 4,
						Style: types.UpperAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 4",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 5,
						Style: types.UpperRoman,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 5",
										},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 2,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "level 2b",
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

	Context("items without numbers", func() {

		It("ordered list with simple unnumbered items", func() {
			source := `. a
. b`

			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "a"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "b"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list with unnumbered items", func() {
			source := `. item 1
. item 2`

			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list with custom numbering on child items with tabs ", func() {
			// note: the [upperroman] attribute must be at the beginning of the line
			source := `. item 1
			.. item 1.1
[upperroman]
			... item 1.1.1
			... item 1.1.2
			.. item 1.2
			. item 2
			.. item 2.1`

			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 2,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 3,
						Style: types.LowerRoman,
						Attributes: types.Attributes{
							types.AttrStyle: types.UpperRoman,
						},
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.1.1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 3,
						Style: types.LowerRoman,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.1.2"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 2,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.2"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 2,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 2.1"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list with all default styles and blank lines", func() {
			source := `. level 1

.. level 2


... level 3



.... level 4
..... level 5.


`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "level 1"},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.OrderedListItem{
						Level: 2,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "level 2"},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.BlankLine{},
					types.OrderedListItem{
						Level: 3,
						Style: types.LowerRoman,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "level 3"},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.BlankLine{},
					types.BlankLine{},
					types.OrderedListItem{
						Level: 4,
						Style: types.UpperAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "level 4"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 5,
						Style: types.UpperRoman,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "level 5."},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.BlankLine{},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})

	Context("numbered items", func() {

		It("ordered list with simple numbered items", func() {
			source := `1. a
2. b`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "a"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "b"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list with numbered items", func() {
			source := `1. item 1
a. item 1.a
2. item 2
b. item 2.a`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.a"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.LowerAlpha,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 2.a"},
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

	Context("list item continuation", func() {

		It("ordered list with item continuation - case 1", func() {
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
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
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
					types.ContinuedListItemElement{
						Offset: 0,
						Element: types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "a delimited block",
								},
							},
						},
					},
					types.ContinuedListItemElement{
						Offset: 0,
						Element: types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "another delimited block",
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
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
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("ordered list with item continuation - case 2", func() {
			source := `. {blank}
+
----
print("one")
----
. {blank}
+
----
print("one")
----`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.AttributeSubstitution{
											Name: "blank",
										},
									},
								},
							},
						},
					},
					types.ContinuedListItemElement{
						Offset: 0,
						Element: types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "print(\"one\")",
								},
							},
						},
					},
					types.OrderedListItem{
						Level: 1,
						Style: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.AttributeSubstitution{
											Name: "blank",
										},
									},
								},
							},
						},
					},
					types.ContinuedListItemElement{
						Offset: 0,
						Element: types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "print(\"one\")",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})
})

var _ = Describe("ordered lists - document", func() {

	Context("ordered list item alone", func() {

		// same single item in the list for each test in this context
		elements := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "item"},
					},
				},
			},
		}

		It("ordered list item with implicit numbering style", func() {
			source := `.. item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level:    1,
								Style:    types.LowerAlpha,
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list item with arabic numbering style", func() {
			source := `1. item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level:    1,
								Style:    types.Arabic,
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list item with lower alpha numbering style", func() {
			source := `b. item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level:    1,
								Style:    types.LowerAlpha,
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list item with upper alpha numbering style", func() {
			source := `B. item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{

								Level:    1,
								Style:    types.UpperAlpha,
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list item with lower roman numbering style", func() {
			source := `i) item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level:    1,
								Style:    types.LowerRoman,
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list item with upper roman numbering style", func() {
			source := `I) item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level:    1,
								Style:    types.UpperRoman,
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list item with predefined attribute", func() {
			source := `. {amp}`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "&amp;"},
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

		It("ordered list item with explicit numbering style", func() {
			source := `[lowerroman]
. item
. item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.Attributes{
							types.AttrStyle: "lowerroman", // will be used during rendering
						},
						Items: []types.OrderedListItem{
							{
								Level:    1,
								Style:    types.Arabic, // will be overridden during rendering
								Elements: elements,
							},
							{
								Level:    1,
								Style:    types.Arabic, // will be overridden during rendering
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list item with explicit start only", func() {
			source := `[start="5"]
. item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.Attributes{
							"start": "5",
						},
						Items: []types.OrderedListItem{
							{
								Level:    1,
								Style:    types.Arabic,
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list item with explicit quoted numbering and start", func() {
			source := `["lowerroman", start="5"]
. item`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.Attributes{
							types.AttrStyle: "lowerroman", // will be used during rendering
							"start":         "5",
						},
						Items: []types.OrderedListItem{
							{
								Level:    1,
								Style:    types.Arabic, // will be overridden during rendering
								Elements: elements,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("max level of ordered items - case 1", func() {
			source := `.Ordered, max nesting
. level 1
.. level 2
... level 3
.... level 4
..... level 5
. level 1`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.Attributes{
							types.AttrTitle: "Ordered, max nesting",
						},
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "level 1",
												},
											},
										},
									},
									types.OrderedList{
										Items: []types.OrderedListItem{
											{
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{
																	Content: "level 2",
																},
															},
														},
													},
													types.OrderedList{
														Items: []types.OrderedListItem{
															{
																Level: 3,
																Style: types.LowerRoman,
																Elements: []interface{}{
																	types.Paragraph{
																		Lines: [][]interface{}{
																			{
																				types.StringElement{
																					Content: "level 3",
																				},
																			},
																		},
																	},
																	types.OrderedList{
																		Items: []types.OrderedListItem{
																			{
																				Level: 4,
																				Style: types.UpperAlpha,
																				Elements: []interface{}{
																					types.Paragraph{
																						Lines: [][]interface{}{
																							{
																								types.StringElement{
																									Content: "level 4",
																								},
																							},
																						},
																					},
																					types.OrderedList{
																						Items: []types.OrderedListItem{
																							{
																								Level: 5,
																								Style: types.UpperRoman,
																								Elements: []interface{}{
																									types.Paragraph{
																										Lines: [][]interface{}{
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
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
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
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("max level of ordered items - case 2", func() {
			source := `.Ordered, max nesting
. level 1
.. level 2
... level 3
.... level 4
..... level 5
.. level 2b`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.Attributes{
							types.AttrTitle: "Ordered, max nesting",
						},
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "level 1",
												},
											},
										},
									},
									types.OrderedList{
										Items: []types.OrderedListItem{
											{
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{
																	Content: "level 2",
																},
															},
														},
													},
													types.OrderedList{
														Items: []types.OrderedListItem{
															{
																Level: 3,
																Style: types.LowerRoman,
																Elements: []interface{}{
																	types.Paragraph{
																		Lines: [][]interface{}{
																			{
																				types.StringElement{
																					Content: "level 3",
																				},
																			},
																		},
																	},
																	types.OrderedList{
																		Items: []types.OrderedListItem{
																			{
																				Level: 4,
																				Style: types.UpperAlpha,
																				Elements: []interface{}{
																					types.Paragraph{
																						Lines: [][]interface{}{
																							{
																								types.StringElement{
																									Content: "level 4",
																								},
																							},
																						},
																					},
																					types.OrderedList{
																						Items: []types.OrderedListItem{
																							{
																								Level: 5,
																								Style: types.UpperRoman,
																								Elements: []interface{}{
																									types.Paragraph{
																										Lines: [][]interface{}{
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
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{
																	Content: "level 2b",
																},
															},
														},
													},
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

	Context("items without numbers", func() {

		It("ordered list with simple unnumbered items", func() {
			source := `. a
. b`

			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a"},
											},
										},
									},
								},
							},
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "b"},
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

		It("ordered list with unnumbered items", func() {
			source := `. item 1
. item 2`

			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "item 1"},
											},
										},
									},
								},
							},
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "item 2"},
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

		It("ordered list with custom numbering on child items with tabs ", func() {
			// note: the [upperroman] attribute must be at the beginning of the line
			source := `. item 1
			.. item 1.1
[upperroman]
			... item 1.1.1
			... item 1.1.2
			.. item 1.2
			. item 2
			.. item 2.1`

			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "item 1"},
											},
										},
									},
									types.OrderedList{
										Items: []types.OrderedListItem{
											{
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{Content: "item 1.1"},
															},
														},
													},
													types.OrderedList{
														Attributes: types.Attributes{
															types.AttrStyle: "upperroman",
														},
														Items: []types.OrderedListItem{
															{
																Level: 3,
																Style: types.LowerRoman, // will be overridden during rendering
																Elements: []interface{}{
																	types.Paragraph{
																		Lines: [][]interface{}{
																			{
																				types.StringElement{Content: "item 1.1.1"},
																			},
																		},
																	},
																},
															},
															{
																Level: 3,
																Style: types.LowerRoman, // will be overridden during rendering
																Elements: []interface{}{
																	types.Paragraph{
																		Lines: [][]interface{}{
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
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
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
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "item 2"},
											},
										},
									},
									types.OrderedList{
										Items: []types.OrderedListItem{
											{
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
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
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list with all default styles and blank lines", func() {
			source := `. level 1

.. level 2


... level 3



.... level 4
..... level 5.


`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "level 1"},
											},
										},
									},
									types.OrderedList{
										Items: []types.OrderedListItem{
											{
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
															{
																types.StringElement{Content: "level 2"},
															},
														},
													},
													types.OrderedList{
														Items: []types.OrderedListItem{
															{
																Level: 3,
																Style: types.LowerRoman,
																Elements: []interface{}{
																	types.Paragraph{
																		Lines: [][]interface{}{
																			{
																				types.StringElement{Content: "level 3"},
																			},
																		},
																	},
																	types.OrderedList{
																		Items: []types.OrderedListItem{
																			{
																				Level: 4,
																				Style: types.UpperAlpha,
																				Elements: []interface{}{
																					types.Paragraph{
																						Lines: [][]interface{}{
																							{
																								types.StringElement{Content: "level 4"},
																							},
																						},
																					},
																					types.OrderedList{
																						Items: []types.OrderedListItem{
																							{
																								Level: 5,
																								Style: types.UpperRoman,
																								Elements: []interface{}{
																									types.Paragraph{
																										Lines: [][]interface{}{
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
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})

	Context("numbered items", func() {

		It("ordered list with simple numbered items", func() {
			source := `1. a
2. b`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a"},
											},
										},
									},
								},
							},
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "b"},
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

		It("ordered list with numbered items", func() {
			source := `1. item 1
a. item 1.a
2. item 2
b. item 2.a`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "item 1"},
											},
										},
									},
									types.OrderedList{
										Items: []types.OrderedListItem{
											{
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
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
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "item 2"},
											},
										},
									},
									types.OrderedList{
										Items: []types.OrderedListItem{
											{
												Level: 2,
												Style: types.LowerAlpha,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
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
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})

	Context("list item continuation", func() {

		It("ordered list with item continuation - case 1", func() {
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
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "foo"},
											},
										},
									},
									types.DelimitedBlock{
										Kind: types.Listing,
										Elements: []interface{}{
											types.VerbatimLine{
												Content: "a delimited block",
											},
										},
									},
									types.DelimitedBlock{
										Kind: types.Listing,
										Elements: []interface{}{
											types.VerbatimLine{
												Content: "another delimited block",
											},
										},
									},
								},
							},
							{
								Level: 1,
								Style: types.Arabic,
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
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("ordered list with item continuation - case 2", func() {
			source := `. {blank}
+
----
print("one")
----
. {blank}
+
----
print("one")
----`
			expected := types.Document{
				Elements: []interface{}{
					types.OrderedList{
						Items: []types.OrderedListItem{
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{},
									},
									types.DelimitedBlock{
										Kind: types.Listing,
										Elements: []interface{}{
											types.VerbatimLine{
												Content: "print(\"one\")",
											},
										},
									},
								},
							},
							{
								Level: 1,
								Style: types.Arabic,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{},
									},
									types.DelimitedBlock{
										Kind: types.Listing,
										Elements: []interface{}{
											types.VerbatimLine{
												Content: "print(\"one\")",
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
