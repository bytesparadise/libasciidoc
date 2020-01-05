package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("labeled lists - draft", func() {

	It("labeled list with a term and description on 2 lines", func() {
		source := `Item1::
Item 1 description
on 2 lines`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item1",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with a single term and no description", func() {
		source := `Item1::`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with a quoted text in term and in description", func() {
		source := "`foo()`::\n" +
			`This function is _untyped_.`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Term: []interface{}{
						types.StringElement{
							Content: "`foo()`", // the term is a raw string in the DraftDocument
						},
					},
					Level: 1,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with a horizontal layout attribute", func() {
		source := `[horizontal]
Item1:: foo`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{
						"layout": "horizontal",
					},
					Level: 1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item1",
						},
					},
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
			},
		}
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with a single term and a blank line", func() {
		source := `Item1::
			`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item1",
						},
					},
					Elements: []interface{}{},
				},
			},
		}
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with multiple sibling items", func() {
		source := `Item 1::
Item 1 description
Item 2:: 
Item 2 description
Item 3:: 
Item 3 description`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 1",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "Item 1 description"},
								},
							},
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 2",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "Item 2 description"},
								},
							},
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 3",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with multiple nested items", func() {
		source := `Item 1::
Item 1 description
Item 2:::
Item 2 description
Item 3::::
Item 3 description`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 1",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "Item 1 description"},
								},
							},
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      2,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 2",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "Item 2 description"},
								},
							},
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      3,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 3",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with nested unordered list - case 1", func() {
		source := `Empty item:: 
* foo
* bar
Item with description:: something simple`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Empty item",
						},
					},
					Elements: []interface{}{},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					CheckStyle:  types.NoCheck,
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
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					CheckStyle:  types.NoCheck,
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
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item with description",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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

		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with a single item and paragraph", func() {
		source := `Item 1::
foo
bar

a normal paragraph.`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 1",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a normal paragraph."},
						},
					},
				},
			},
		}
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with item continuation", func() {
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
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
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
					Element: types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 2",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "something simple"},
								},
							},
						},
					},
				},
				// the `+` continuation produces the second element below
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
											Content: "another fenced block",
										},
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

	It("labeled list without item continuation", func() {
		source := `Item 1::
----
a fenced block
----
Item 2:: something simple
----
another fenced block
----`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 1",
						},
					},
					Elements: []interface{}{},
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
										Content: "a fenced block",
									},
								},
							},
						},
					},
				},
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Item 2",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "something simple"},
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
										Content: "another fenced block",
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

	It("labeled list with nested unordered list - case 2", func() {
		source := `Labeled item::
- unordered item`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "Labeled item",
						},
					},
					Elements: []interface{}{},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.Dash,
					CheckStyle:  types.NoCheck,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("labeled list with title", func() {
		source := `.Labeled, single-line
first term:: definition of the first term
second term:: definition of the second term`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{
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
							Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term: []interface{}{
						types.StringElement{
							Content: "second term",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("max level of labeled items - case 1", func() {
		source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 1:: description 1`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{
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
							Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})

	It("max level of labeled items - case 2", func() {
		source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 2::: description 2`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{
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
							Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDraftDocument(expected))
	})
})

var _ = Describe("labeled lists - document", func() {

	It("labeled list with a term and description on 2 lines", func() {
		source := `Item1::
Item 1 description
on 2 lines`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item1",
								},
							},

							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with a single term and no description", func() {
		source := `Item1::`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with a quoted text in term and in description", func() {
		source := "`foo()`::\n" +
			`This function is _untyped_.`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
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
									Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with a horizontal layout attribute", func() {
		source := `[horizontal]
Item1:: foo`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{
						"layout": "horizontal",
					},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item1",
								},
							},

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
					},
				},
			},
		}
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with a single term and a blank line", func() {
		source := `Item1::
			`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with multiple sibling items", func() {
		source := `Item 1::
Item 1 description
Item 2:: 
Item 2 description
Item 3:: 
Item 3 description`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item 1",
								},
							},

							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "Item 1 description"},
										},
									},
								},
							},
						},
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item 2",
								},
							},

							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "Item 2 description"},
										},
									},
								},
							},
						},
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item 3",
								},
							},

							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with multiple nested items", func() {
		source := `Item 1::
Item 1 description
Item 2:::
Item 2 description
Item 3::::
Item 3 description`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item 1",
								},
							},

							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "Item 1 description"},
										},
									},
								},
								types.LabeledList{
									Attributes: types.ElementAttributes{},
									Items: []types.LabeledListItem{
										{
											Attributes: types.ElementAttributes{},
											Level:      2,
											Term: []interface{}{
												types.StringElement{
													Content: "Item 2",
												},
											},

											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: [][]interface{}{
														{
															types.StringElement{Content: "Item 2 description"},
														},
													},
												},
												types.LabeledList{
													Attributes: types.ElementAttributes{},
													Items: []types.LabeledListItem{
														{
															Attributes: types.ElementAttributes{},
															Level:      3,
															Term: []interface{}{
																types.StringElement{
																	Content: "Item 3",
																},
															},

															Elements: []interface{}{
																types.Paragraph{
																	Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with nested unordered list - case 1", func() {
		source := `Empty item:: 
* foo
* bar
Item with description:: something simple`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Empty item",
								},
							},

							Elements: []interface{}{
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
															types.StringElement{Content: "foo"},
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
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item with description",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
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

		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with a single item and paragraph", func() {
		source := `Item 1::
foo
bar

a normal paragraph.`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item 1",
								},
							},

							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
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
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a normal paragraph."},
						},
					},
				},
			},
		}
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with item continuation", func() {
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
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item 1",
								},
							},

							Elements: []interface{}{
								types.DelimitedBlock{
									Attributes: types.ElementAttributes{},
									Kind:       types.Listing,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
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
							},
						},
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item 2",
								},
							},

							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "something simple"},
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
				},
			},
		}

		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list without item continuation", func() {
		source := `Item 1::
----
a fenced block
----
Item 2:: something simple
----
another fenced block
----`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Item 1",
								},
							},

							Elements: []interface{}{},
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
										Content: "a fenced block",
									},
								},
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
							Term: []interface{}{
								types.StringElement{
									Content: "Item 2",
								},
							},

							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
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
				types.DelimitedBlock{
					Attributes: types.ElementAttributes{},
					Kind:       types.Listing,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
		}
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with nested unordered list - case 2", func() {
		source := `Labeled item::
- unordered item`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "Labeled item",
								},
							},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list with title", func() {
		source := `.Labeled, single-line
first term:: definition of the first term
second term:: definition of the second term`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{
						types.AttrTitle: "Labeled, single-line",
					},
					Items: []types.LabeledListItem{
						{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "first term",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
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
							Attributes: types.ElementAttributes{},
							Level:      1,
							Term: []interface{}{
								types.StringElement{
									Content: "second term",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("max level of labeled items - case 1", func() {
		source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 1:: description 1`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{
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
							Attributes: types.ElementAttributes{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "description 1",
											},
										},
									},
								},
								types.LabeledList{
									Attributes: types.ElementAttributes{},
									Items: []types.LabeledListItem{
										{
											Level: 2,
											Term: []interface{}{
												types.StringElement{
													Content: "level 2",
												},
											},
											Attributes: types.ElementAttributes{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: [][]interface{}{
														{
															types.StringElement{
																Content: "description 2",
															},
														},
													},
												},
												types.LabeledList{
													Attributes: types.ElementAttributes{},
													Items: []types.LabeledListItem{
														{
															Level: 3,
															Term: []interface{}{
																types.StringElement{
																	Content: "level 3",
																},
															},
															Attributes: types.ElementAttributes{},
															Elements: []interface{}{
																types.Paragraph{
																	Attributes: types.ElementAttributes{},
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
							Attributes: types.ElementAttributes{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("max level of labeled items - case 2", func() {
		source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 2::: description 2`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{
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
							Attributes: types.ElementAttributes{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "description 1",
											},
										},
									},
								},
								types.LabeledList{
									Attributes: types.ElementAttributes{},
									Items: []types.LabeledListItem{
										{
											Level: 2,
											Term: []interface{}{
												types.StringElement{
													Content: "level 2",
												},
											},
											Attributes: types.ElementAttributes{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: [][]interface{}{
														{
															types.StringElement{
																Content: "description 2",
															},
														},
													},
												},
												types.LabeledList{
													Attributes: types.ElementAttributes{},
													Items: []types.LabeledListItem{
														{
															Level: 3,
															Term: []interface{}{
																types.StringElement{
																	Content: "level 3",
																},
															},
															Attributes: types.ElementAttributes{},
															Elements: []interface{}{
																types.Paragraph{
																	Attributes: types.ElementAttributes{},
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
											Attributes: types.ElementAttributes{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
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
		Expect(source).To(BecomeDocument(expected))
	})

	It("labeled list item with predefined attribute", func() {
		source := `level 1:: {amp}`
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.LabeledList{
					Attributes: types.ElementAttributes{},
					Items: []types.LabeledListItem{
						{
							Level: 1,
							Term: []interface{}{
								types.StringElement{
									Content: "level 1",
								},
							},
							Attributes: types.ElementAttributes{},
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
})
