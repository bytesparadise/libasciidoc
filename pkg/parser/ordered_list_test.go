package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ordered lists - draft", func() {

	Context("ordered list item alone", func() {

		// same single item in the list for each test in this context
		elements := []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
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
						Level:          2,
						NumberingStyle: types.LowerAlpha,
						Attributes:     map[string]interface{}{},
						Elements:       elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list item with arabic numbering style", func() {
			source := `1. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements:       elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list item with lower alpha numbering style", func() {
			source := `b. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.LowerAlpha,
						Attributes:     map[string]interface{}{},
						Elements:       elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list item with upper alpha numbering style", func() {
			source := `B. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.UpperAlpha,
						Attributes:     map[string]interface{}{},
						Elements:       elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list item with lower roman numbering style", func() {
			source := `i) item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.LowerRoman,
						Attributes:     map[string]interface{}{},
						Elements:       elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list item with upper roman numbering style", func() {
			source := `I) item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.UpperRoman,
						Attributes:     map[string]interface{}{},
						Elements:       elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list item with explicit numbering style", func() {
			source := `[lowerroman]
. item
. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Attributes: types.ElementAttributes{
							"lowerroman": nil,
						},
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements:       elements,
					},
					types.OrderedListItem{
						Attributes:     types.ElementAttributes{},
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements:       elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list item with explicit start only", func() {
			source := `[start=5]
. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes: types.ElementAttributes{
							"start": "5",
						},
						Elements: elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list item with explicit quoted numbering and start", func() {
			source := `["lowerroman", start="5"]
. item`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes: types.ElementAttributes{
							"lowerroman": nil,
							"start":      "5",
						},
						Elements: elements,
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
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
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Ordered, max nesting",
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          2,
						NumberingStyle: types.LowerAlpha,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          3,
						NumberingStyle: types.LowerRoman,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          4,
						NumberingStyle: types.UpperAlpha,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          5,
						NumberingStyle: types.UpperRoman,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
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
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Ordered, max nesting",
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          2,
						NumberingStyle: types.LowerAlpha,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          3,
						NumberingStyle: types.LowerRoman,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          4,
						NumberingStyle: types.UpperAlpha,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          5,
						NumberingStyle: types.UpperRoman,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          2,
						NumberingStyle: types.LowerAlpha,
						Attributes:     types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})

	Context("items without numbers", func() {

		It("ordered list with simple unnumbered items", func() {
			source := `. a
. b`

			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "a"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list with unnumbered items", func() {
			source := `. item 1
. item 2`

			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
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
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          2,
						NumberingStyle: types.LowerAlpha,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          3,
						NumberingStyle: types.LowerRoman,
						Attributes: types.ElementAttributes{
							"upperroman": nil,
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.1.1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          3,
						NumberingStyle: types.LowerRoman,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.1.2"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          2,
						NumberingStyle: types.LowerAlpha,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.2"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          2,
						NumberingStyle: types.LowerAlpha,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
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
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          2,
						NumberingStyle: types.LowerAlpha,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          3,
						NumberingStyle: types.LowerRoman,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
						Level:          4,
						NumberingStyle: types.UpperAlpha,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "level 4"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          5,
						NumberingStyle: types.UpperRoman,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})

	Context("numbered items", func() {

		It("ordered list with simple numbered items", func() {
			source := `1. a
2. b`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "a"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("ordered list with numbered items", func() {
			source := `1. item 1
a. item 1.a
2. item 2
b. item 2.a`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.LowerAlpha,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 1.a"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.Arabic,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "item 2"},
									},
								},
							},
						},
					},
					types.OrderedListItem{
						Level:          1,
						NumberingStyle: types.LowerAlpha,
						Attributes:     map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
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
						Attributes:     types.ElementAttributes{},
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
							Attributes: types.ElementAttributes{},
							Kind:       types.Listing,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a delimited block",
											},
										},
									},
								},
							},
						},
					},
					types.ContinuedListItemElement{
						Offset: 0,
						Element: types.DelimitedBlock{
							Attributes: types.ElementAttributes{},
							Kind:       types.Listing,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "another delimited block",
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
								Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDraftDocument(expected))
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
						Attributes:     types.ElementAttributes{},
						Level:          1,
						NumberingStyle: types.Arabic,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.DocumentAttributeSubstitution{
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
							Kind:       types.Listing,
							Attributes: types.ElementAttributes{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "print(\"one\")"},
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
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.DocumentAttributeSubstitution{
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
							Kind:       types.Listing,
							Attributes: types.ElementAttributes{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "print(\"one\")"},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})
})

var _ = Describe("ordered lists - document", func() {

	Context("ordered list item alone", func() {

		// same single item in the list for each test in this context
		elements := []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
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
								NumberingStyle: types.LowerAlpha,
								Attributes:     map[string]interface{}{},
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with arabic numbering style", func() {
			source := `1. item`
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
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with lower alpha numbering style", func() {
			source := `b. item`
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
								NumberingStyle: types.LowerAlpha,
								Attributes:     map[string]interface{}{},
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with upper alpha numbering style", func() {
			source := `B. item`
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
								NumberingStyle: types.UpperAlpha,
								Attributes:     map[string]interface{}{},
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with lower roman numbering style", func() {
			source := `i) item`
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
								NumberingStyle: types.LowerRoman,
								Attributes:     map[string]interface{}{},
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with upper roman numbering style", func() {
			source := `I) item`
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
								NumberingStyle: types.UpperRoman,
								Attributes:     map[string]interface{}{},
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with predefined attribute", func() {
			source := `. {amp}`
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with explicit numbering style", func() {
			source := `[lowerroman]
. item
. item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrNumberingStyle: "lowerroman", // will be used during rendering
						},
						Items: []types.OrderedListItem{
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								NumberingStyle: types.Arabic, // will be overridden during rendering
								Elements:       elements,
							},
							{
								Attributes:     types.ElementAttributes{},
								Level:          1,
								NumberingStyle: types.Arabic, // will be overridden during rendering
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with explicit start only", func() {
			source := `[start=5]
. item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.ElementAttributes{
							"start": "5",
						},
						Items: []types.OrderedListItem{
							{
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     types.ElementAttributes{},
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list item with explicit quoted numbering and start", func() {
			source := `["lowerroman", start="5"]
. item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrNumberingStyle: "lowerroman", // will be used during rendering
							"start":                  "5",
						},
						Items: []types.OrderedListItem{
							{
								Level:          1,
								NumberingStyle: types.Arabic, // will be overridden during rendering
								Attributes:     types.ElementAttributes{},
								Elements:       elements,
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
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
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Ordered, max nesting",
						},
						Items: []types.OrderedListItem{
							{
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
												NumberingStyle: types.LowerAlpha,
												Attributes:     types.ElementAttributes{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
																NumberingStyle: types.LowerRoman,
																Attributes:     types.ElementAttributes{},
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																				NumberingStyle: types.UpperAlpha,
																				Attributes:     types.ElementAttributes{},
																				Elements: []interface{}{
																					types.Paragraph{
																						Attributes: types.ElementAttributes{},
																						Lines: [][]interface{}{
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
																								NumberingStyle: types.UpperRoman,
																								Attributes:     types.ElementAttributes{},
																								Elements: []interface{}{
																									types.Paragraph{
																										Attributes: types.ElementAttributes{},
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
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDocument(expected))
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
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.OrderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Ordered, max nesting",
						},
						Items: []types.OrderedListItem{
							{
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
												NumberingStyle: types.LowerAlpha,
												Attributes:     types.ElementAttributes{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
																NumberingStyle: types.LowerRoman,
																Attributes:     types.ElementAttributes{},
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																				NumberingStyle: types.UpperAlpha,
																				Attributes:     types.ElementAttributes{},
																				Elements: []interface{}{
																					types.Paragraph{
																						Attributes: types.ElementAttributes{},
																						Lines: [][]interface{}{
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
																								NumberingStyle: types.UpperRoman,
																								Attributes:     types.ElementAttributes{},
																								Elements: []interface{}{
																									types.Paragraph{
																										Attributes: types.ElementAttributes{},
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
												Level:          2,
												NumberingStyle: types.LowerAlpha,
												Attributes:     types.ElementAttributes{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDocument(expected))
		})
	})

	Context("items without numbers", func() {

		It("ordered list with simple unnumbered items", func() {
			source := `. a
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
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     map[string]interface{}{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a"},
											},
										},
									},
								},
							},
							{
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     map[string]interface{}{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list with unnumbered items", func() {
			source := `. item 1
. item 2`

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
												types.StringElement{Content: "item 1"},
											},
										},
									},
								},
							},
							{
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     map[string]interface{}{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDocument(expected))
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
												types.StringElement{Content: "item 1"},
											},
										},
									},
									types.OrderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.OrderedListItem{
											{
												Level:          2,
												NumberingStyle: types.LowerAlpha,
												Attributes:     map[string]interface{}{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
															{
																types.StringElement{Content: "item 1.1"},
															},
														},
													},
													types.OrderedList{
														Attributes: types.ElementAttributes{
															types.AttrNumberingStyle: "upperroman",
														},
														Items: []types.OrderedListItem{
															{
																Level:          3,
																NumberingStyle: types.LowerRoman, // will be overridden during rendering
																Attributes:     types.ElementAttributes{},
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
																			{
																				types.StringElement{Content: "item 1.1.1"},
																			},
																		},
																	},
																},
															},
															{
																Level:          3,
																NumberingStyle: types.LowerRoman, // will be overridden during rendering
																Attributes:     map[string]interface{}{},
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
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
												Level:          2,
												NumberingStyle: types.LowerAlpha,
												Attributes:     map[string]interface{}{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
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
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     map[string]interface{}{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
												NumberingStyle: types.LowerAlpha,
												Attributes:     map[string]interface{}{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list with all default styles and blank lines", func() {
			source := `. level 1

.. level 2


... level 3



.... level 4
..... level 5.


`
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
												types.StringElement{Content: "level 1"},
											},
										},
									},
									types.OrderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.OrderedListItem{
											{
												Level:          2,
												NumberingStyle: types.LowerAlpha,
												Attributes:     map[string]interface{}{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: [][]interface{}{
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
																NumberingStyle: types.LowerRoman,
																Attributes:     map[string]interface{}{},
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: [][]interface{}{
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
																				NumberingStyle: types.UpperAlpha,
																				Attributes:     map[string]interface{}{},
																				Elements: []interface{}{
																					types.Paragraph{
																						Attributes: types.ElementAttributes{},
																						Lines: [][]interface{}{
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
																								NumberingStyle: types.UpperRoman,
																								Attributes:     map[string]interface{}{},
																								Elements: []interface{}{
																									types.Paragraph{
																										Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDocument(expected))
		})
	})

	Context("numbered items", func() {

		It("ordered list with simple numbered items", func() {
			source := `1. a
2. b`
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
												types.StringElement{Content: "a"},
											},
										},
									},
								},
							},
							{
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     map[string]interface{}{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("ordered list with numbered items", func() {
			source := `1. item 1
a. item 1.a
2. item 2
b. item 2.a`
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
												types.StringElement{Content: "item 1"},
											},
										},
									},
									types.OrderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.OrderedListItem{
											{
												Level:          2,
												NumberingStyle: types.LowerAlpha,
												Attributes:     map[string]interface{}{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
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
								Level:          1,
								NumberingStyle: types.Arabic,
								Attributes:     map[string]interface{}{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
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
												NumberingStyle: types.LowerAlpha,
												Attributes:     map[string]interface{}{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
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
			Expect(source).To(BecomeDocument(expected))
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
												types.StringElement{Content: "foo"},
											},
										},
									},
									types.DelimitedBlock{
										Attributes: types.ElementAttributes{},
										Kind:       types.Listing,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{
															Content: "a delimited block",
														},
													},
												},
											},
										},
									},
									types.DelimitedBlock{
										Attributes: types.ElementAttributes{},
										Kind:       types.Listing,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{
															Content: "another delimited block",
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
			Expect(source).To(BecomeDocument(expected))
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
										Lines:      [][]interface{}{},
									},
									types.DelimitedBlock{
										Kind:       types.Listing,
										Attributes: types.ElementAttributes{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "print(\"one\")"},
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
										Attributes: types.ElementAttributes{},
										Lines:      [][]interface{}{},
									},
									types.DelimitedBlock{
										Kind:       types.Listing,
										Attributes: types.ElementAttributes{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: [][]interface{}{
													{
														types.StringElement{Content: "print(\"one\")"},
													},
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
