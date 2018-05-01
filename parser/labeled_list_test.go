package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("labeled lists", func() {

	It("labeled list with a term and description on 2 lines", func() {
		actualContent := `Item1::
Item 1 description
on 2 lines`
		expectedResult := types.LabeledList{
			Attributes: map[string]interface{}{},
			Items: []types.LabeledListItem{
				{
					Term: "Item1",
					Elements: []types.DocElement{
						types.ListParagraph{
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "Item 1 description"},
									},
								},
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "on 2 lines"},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
	})

	It("labeled list with a single term and no description", func() {
		actualContent := `Item1::`
		expectedResult := types.LabeledList{
			Attributes: map[string]interface{}{},
			Items: []types.LabeledListItem{
				{
					Term: "Item1",
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
	})

	It("labeled list with a horizontal layout attribute", func() {
		actualContent := `[horizontal]
Item1:: foo`
		expectedResult := types.LabeledList{
			Attributes: map[string]interface{}{
				"layout": "horizontal",
			},
			Items: []types.LabeledListItem{
				{
					Term: "Item1",
					Elements: []types.DocElement{
						types.ListParagraph{
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "foo"},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
	})

	It("labeled list with a single term and a blank line", func() {
		actualContent := `Item1::
			`
		expectedResult := types.LabeledList{
			Attributes: map[string]interface{}{},
			Items: []types.LabeledListItem{
				{
					Term: "Item1",
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
	})

	It("labeled list with multiple items", func() {
		actualContent := `Item 1::
Item 1 description
Item 2:: 
Item 2 description
Item 3:: 
Item 3 description`
		expectedResult := types.LabeledList{
			Attributes: map[string]interface{}{},
			Items: []types.LabeledListItem{
				{
					Term: "Item 1",
					Elements: []types.DocElement{
						types.ListParagraph{
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "Item 1 description"},
									},
								},
							},
						},
					},
				},
				{
					Term: "Item 2",
					Elements: []types.DocElement{
						types.ListParagraph{
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "Item 2 description"},
									},
								},
							},
						},
					},
				},
				{
					Term: "Item 3",
					Elements: []types.DocElement{
						types.ListParagraph{
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "Item 3 description"},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
	})

	It("labeled list with nested list", func() {
		actualContent := `Empty item:: 
* foo
* bar
Item with description:: something simple`
		expectedResult := types.LabeledList{
			Attributes: map[string]interface{}{},
			Items: []types.LabeledListItem{
				{
					Term: "Empty item",
					Elements: []types.DocElement{
						types.UnorderedList{
							Attributes: map[string]interface{}{},
							Items: []types.UnorderedListItem{
								{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									Elements: []types.DocElement{
										types.ListParagraph{
											Lines: []types.InlineContent{
												{
													Elements: []types.InlineElement{
														types.StringElement{Content: "foo"},
													},
												},
											},
										},
									},
								},
								{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									Elements: []types.DocElement{
										types.ListParagraph{
											Lines: []types.InlineContent{
												{
													Elements: []types.InlineElement{
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
				},
				{
					Term: "Item with description",
					Elements: []types.DocElement{
						types.ListParagraph{
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "something simple"},
									},
								},
							},
						},
					},
				},
			},
		}

		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
	})

	It("labeled list with a single item and paragraph", func() {
		actualContent := `Item 1::
foo
bar

a normal paragraph.`
		expectedResult := types.Document{
			Attributes:        map[string]interface{}{},
			ElementReferences: map[string]interface{}{},
			Elements: []types.DocElement{
				types.LabeledList{
					Attributes: map[string]interface{}{},
					Items: []types.LabeledListItem{
						{
							Term: "Item 1",
							Elements: []types.DocElement{
								types.ListParagraph{
									Lines: []types.InlineContent{
										{
											Elements: []types.InlineElement{
												types.StringElement{Content: "foo"},
											},
										},
										{
											Elements: []types.InlineElement{
												types.StringElement{Content: "bar"},
											},
										},
									},
								},
							},
						},
					},
				},
				types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineContent{
						{
							Elements: []types.InlineElement{
								types.StringElement{Content: "a normal paragraph."},
							},
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
			Attributes:        map[string]interface{}{},
			ElementReferences: map[string]interface{}{},
			Elements: []types.DocElement{
				types.LabeledList{
					Attributes: map[string]interface{}{},
					Items: []types.LabeledListItem{
						{
							Term: "Item 1",
							Elements: []types.DocElement{
								types.DelimitedBlock{
									Kind:       types.ListingBlock,
									Attributes: map[string]interface{}{},
									Elements: []types.DocElement{
										types.Paragraph{
											Attributes: map[string]interface{}{},
											Lines: []types.InlineContent{
												{
													Elements: []types.InlineElement{
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
						},
						{
							Term: "Item 2",
							Elements: []types.DocElement{
								types.ListParagraph{
									Lines: []types.InlineContent{
										{
											Elements: []types.InlineElement{
												types.StringElement{Content: "something simple"},
											},
										},
									},
								},
								types.DelimitedBlock{
									Kind:       types.ListingBlock,
									Attributes: map[string]interface{}{},
									Elements: []types.DocElement{
										types.Paragraph{
											Attributes: map[string]interface{}{},
											Lines: []types.InlineContent{
												{
													Elements: []types.InlineElement{
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
			Attributes:        map[string]interface{}{},
			ElementReferences: map[string]interface{}{},
			Elements: []types.DocElement{
				types.LabeledList{
					Attributes: map[string]interface{}{},
					Items: []types.LabeledListItem{
						{
							Term: "Item 1",
						},
					},
				},
				types.DelimitedBlock{
					Kind:       types.ListingBlock,
					Attributes: map[string]interface{}{},
					Elements: []types.DocElement{
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{
											Content: "a fenced block",
										},
									},
								},
							},
						},
					},
				},
				types.LabeledList{
					Attributes: map[string]interface{}{},
					Items: []types.LabeledListItem{
						{
							Term: "Item 2",
							Elements: []types.DocElement{
								types.ListParagraph{
									Lines: []types.InlineContent{
										{
											Elements: []types.InlineElement{
												types.StringElement{Content: "something simple"},
											},
										},
									},
								},
							},
						},
					},
				},
				types.DelimitedBlock{
					Kind:       types.ListingBlock,
					Attributes: map[string]interface{}{},
					Elements: []types.DocElement{
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
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

		verify(GinkgoT(), expectedResult, actualContent)
	})

})
