package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("labeled lists", func() {

	It("labeled list with a term and description on 2 lines", func() {
		actualContent := `Item1::
Item 1 description
on 2 lines`
		expectedResult := types.LabeledList{
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
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("labeled list with a single term and no description", func() {
		actualContent := `Item1::`
		expectedResult := types.LabeledList{
			Attributes: types.ElementAttributes{},
			Items: []types.LabeledListItem{
				{
					Attributes: types.ElementAttributes{},
					Term:       "Item1",
					Level:      1,
					Elements:   []interface{}{},
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("labeled list with a horizontal layout attribute", func() {
		actualContent := `[horizontal]
Item1:: foo`
		expectedResult := types.LabeledList{
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
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("labeled list with a single term and a blank line", func() {
		actualContent := `Item1::
			`
		expectedResult := types.LabeledList{
			Attributes: types.ElementAttributes{},
			Items: []types.LabeledListItem{
				{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Term:       "Item1",
					Elements:   []interface{}{},
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("labeled list with multiple sibling items", func() {
		actualContent := `Item 1::
Item 1 description
Item 2:: 
Item 2 description
Item 3:: 
Item 3 description`
		expectedResult := types.LabeledList{
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
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("labeled list with multiple nested items", func() {
		actualContent := `Item 1::
Item 1 description
Item 2:::
Item 2 description
Item 3::::
Item 3 description`
		expectedResult := types.LabeledList{
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
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("labeled list with nested list", func() {
		actualContent := `Empty item:: 
* foo
* bar
Item with description:: something simple`
		expectedResult := types.LabeledList{
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
		}

		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("labeled list with a single item and paragraph", func() {
		actualContent := `Item 1::
foo
bar

a normal paragraph.`
		expectedResult := types.Document{
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
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("labeled list with item continuation", func() {
		actualContent := `Item 1::
+
----
a fenced block
----
Item 2:: something simple
+
----
another fenced block
----`
		expectedResult := types.Document{
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

		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("labeled list without item continuation", func() {
		actualContent := `Item 1::
----
a fenced block
----
Item 2:: something simple
----
another fenced block
----`
		expectedResult := types.Document{
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
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("labeled list with nested unordered list", func() {
		actualContent := `Labeled item::
- unordered item`
		expectedResult := types.LabeledList{
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
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("labeled list with title", func() {
		actualContent := `.Labeled, single-line
first term:: definition of the first term
second term:: definition of the second term`
		expectedResult := types.LabeledList{
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
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})
})
