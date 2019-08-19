package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("labeled lists - preflight", func() {

	It("labeled list with a term and description on 2 lines", func() {
		source := `Item1::
Item 1 description
on 2 lines`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Item1",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
		verifyPreflight(expected, source)
	})

	It("labeled list with a single term and no description", func() {
		source := `Item1::`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Term:       "Item1",
					Level:      1,
					Elements:   []interface{}{},
				},
			},
		}
		verifyPreflight(expected, source)
	})

	It("labeled list with a horizontal layout attribute", func() {
		source := `[horizontal]
Item1:: foo`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{
						"layout": "horizontal",
					},
					Level: 1,
					Term:  "Item1",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "foo"},
								},
							},
						},
					},
				},
			},
		}
		verifyPreflight(expected, source)
	})

	It("labeled list with a single term and a blank line", func() {
		source := `Item1::
			`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Item1",
					Elements:   []interface{}{},
				},
			},
		}
		verifyPreflight(expected, source)
	})

	It("labeled list with multiple sibling items", func() {
		source := `Item 1::
Item 1 description
Item 2:: 
Item 2 description
Item 3:: 
Item 3 description`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Item 1",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Term:       "Item 2",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Term:       "Item 3",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "Item 3 description"},
								},
							},
						},
					},
				},
			},
		}
		verifyPreflight(expected, source)
	})

	It("labeled list with multiple nested items", func() {
		source := `Item 1::
Item 1 description
Item 2:::
Item 2 description
Item 3::::
Item 3 description`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Item 1",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Term:       "Item 2",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Term:       "Item 3",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "Item 3 description"},
								},
							},
						},
					},
				},
			},
		}
		verifyPreflight(expected, source)
	})

	It("labeled list with nested unordered list - case 1", func() {
		source := `Empty item:: 
* foo
* bar
Item with description:: something simple`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Empty item",
					Elements:   []interface{}{},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					CheckStyle:  types.NoCheck,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
							Lines: []types.InlineElements{
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
					Term:       "Item with description",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "something simple"},
								},
							},
						},
					},
				},
			},
		}

		verifyPreflight(expected, source)
	})

	It("labeled list with a single item and paragraph", func() {
		source := `Item 1::
foo
bar

a normal paragraph.`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Item 1",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a normal paragraph."},
						},
					},
				},
			},
		}
		verifyPreflight(expected, source)
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
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Item 1",
					Elements:   []interface{}{},
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
								Lines: []types.InlineElements{
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
					Term:       "Item 2",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
								Lines: []types.InlineElements{
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
		verifyPreflight(expected, source)
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
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Item 1",
					Elements:   []interface{}{},
				},
				types.DelimitedBlock{
					Attributes: types.ElementAttributes{},
					Kind:       types.Listing,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Term:       "Item 2",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
							Lines: []types.InlineElements{
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
		verifyPreflight(expected, source)
	})

	It("labeled list with nested unordered list - case 2", func() {
		source := `Labeled item::
- unordered item`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Labeled item",
					Elements:   []interface{}{},
				},
				types.UnorderedListItem{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.Dash,
					CheckStyle:  types.NoCheck,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "unordered item"},
								},
							},
						},
					},
				},
			},
		}
		verifyPreflight(expected, source)
	})

	It("labeled list with title", func() {
		source := `.Labeled, single-line
first term:: definition of the first term
second term:: definition of the second term`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{
						types.AttrTitle: "Labeled, single-line",
					},
					Level: 1,
					Term:  "first term",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Term:       "second term",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
		verifyPreflight(expected, source)
	})

	It("max level of labeled items - case 1", func() {
		source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 1:: description 1`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{
						types.AttrTitle: "Labeled, max nesting",
					},
					Level: 1,
					Term:  "level 1",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Level:      2,
					Term:       "level 2",
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Level:      3,
					Term:       "level 3",
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Level:      1,
					Term:       "level 1",
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
		verifyPreflight(expected, source)
	})

	It("max level of labeled items - case 2", func() {
		source := `.Labeled, max nesting
level 1:: description 1
level 2::: description 2
level 3:::: description 3
level 2::: description 2`
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.LabeledListItem{
					Attributes: types.ElementAttributes{
						types.AttrTitle: "Labeled, max nesting",
					},
					Level: 1,
					Term:  "level 1",
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Level:      2,
					Term:       "level 2",
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Level:      3,
					Term:       "level 3",
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
					Level:      2,
					Term:       "level 2",
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
		verifyPreflight(expected, source)
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
							Term:       "Item1",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
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
							Term:       "Item1",
							Level:      1,
							Elements:   []interface{}{},
						},
					},
				},
			},
		}
		verifyDocument(expected, source)
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
							Term:       "Item1",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
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
							Term:       "Item1",
							Elements:   []interface{}{},
						},
					},
				},
			},
		}
		verifyDocument(expected, source)
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
							Term:       "Item 1",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
							Term:       "Item 2",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
							Term:       "Item 3",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
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
							Term:       "Item 1",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
											Term:       "Item 2",
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
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
															Term:       "Item 3",
															Elements: []interface{}{
																types.Paragraph{
																	Attributes: types.ElementAttributes{},
																	Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
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
							Term:       "Empty item",
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
													Lines: []types.InlineElements{
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
													Lines: []types.InlineElements{
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
							Term:       "Item with description",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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

		verifyDocument(expected, source)
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
							Term:       "Item 1",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a normal paragraph."},
						},
					},
				},
			},
		}
		verifyDocument(expected, source)
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
							Term:       "Item 1",
							Elements: []interface{}{
								types.DelimitedBlock{
									Attributes: types.ElementAttributes{},
									Kind:       types.Listing,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: []types.InlineElements{
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
							Term:       "Item 2",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
											Lines: []types.InlineElements{
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

		verifyDocument(expected, source)
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
							Term:       "Item 1",
							Elements:   []interface{}{},
						},
					},
				},
				types.DelimitedBlock{
					Attributes: types.ElementAttributes{},
					Kind:       types.Listing,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
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
							Term:       "Item 2",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
							Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
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
							Term:       "Labeled item",
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
													Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
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
							Term:       "first term",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
							Term:       "second term",
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
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
							Level:      1,
							Term:       "level 1",
							Attributes: types.ElementAttributes{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
											Level:      2,
											Term:       "level 2",
											Attributes: types.ElementAttributes{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
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
															Level:      3,
															Term:       "level 3",
															Attributes: types.ElementAttributes{},
															Elements: []interface{}{
																types.Paragraph{
																	Attributes: types.ElementAttributes{},
																	Lines: []types.InlineElements{
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
							Level:      1,
							Term:       "level 1",
							Attributes: types.ElementAttributes{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
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
							Level:      1,
							Term:       "level 1",
							Attributes: types.ElementAttributes{},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
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
											Level:      2,
											Term:       "level 2",
											Attributes: types.ElementAttributes{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
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
															Level:      3,
															Term:       "level 3",
															Attributes: types.ElementAttributes{},
															Elements: []interface{}{
																types.Paragraph{
																	Attributes: types.ElementAttributes{},
																	Lines: []types.InlineElements{
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
											Level:      2,
											Term:       "level 2",
											Attributes: types.ElementAttributes{},
											Elements: []interface{}{
												types.Paragraph{
													Attributes: types.ElementAttributes{},
													Lines: []types.InlineElements{
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
		verifyDocument(expected, source)
	})
})
