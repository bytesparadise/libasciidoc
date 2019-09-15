package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("unordered lists - preflight", func() {

	Context("valid content", func() {

		It("unordered list with a basic single item", func() {
			source := `* a list item`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "a list item"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("unordered list with ID, title, role and a single item", func() {
			source := `.mytitle
[#listID]
[.myrole]
* a list item`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.UnorderedListItem{
						Attributes: types.ElementAttributes{
							types.AttrTitle:    "mytitle",
							types.AttrID:       "listID",
							types.AttrCustomID: true,
							types.AttrRole:     "myrole",
						},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a list item"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})
		It("unordered list with a title and a single item", func() {
			source := `.a title
	* a list item`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.UnorderedListItem{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "a title",
						},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a list item"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("unordered list with 2 items with stars", func() {
			source := `* a first item
					* a second item with *bold content*`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "a first item"},
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
										types.StringElement{Content: "a second item with "},
										types.QuotedText{
											Kind: types.Bold,
											Elements: types.InlineElements{
												types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("unordered list based on article.adoc (with heading spaces)", func() {
			source := `.Unordered list title
		* list item 1
		** nested list item A
		*** nested nested list item A.1
		*** nested nested list item A.2
		** nested list item B
		*** nested nested list item B.1
		*** nested nested list item B.2
		* list item 2`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.UnorderedListItem{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Unordered list title",
						},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "list item 1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "nested list item A"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "nested nested list item A.1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "nested nested list item A.2"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "nested list item B"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "nested nested list item B.1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "nested nested list item B.2"},
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
										types.StringElement{Content: "list item 2"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("unordered list with 2 items with carets", func() {
			source := "- a first item\n" +
				"- a second item with *bold content*"
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "a first item"},
									},
								},
							},
						},
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
										types.StringElement{Content: "a second item with "},
										types.QuotedText{
											Kind: types.Bold,
											Elements: types.InlineElements{
												types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("unordered list with items with mixed styles", func() {
			source := `- a parent item
					* a child item
					- another parent item
					* another child item
					** with a sub child item`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "a parent item"},
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
										types.StringElement{Content: "a child item"},
									},
								},
							},
						},
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
										types.StringElement{Content: "another parent item"},
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
										types.StringElement{Content: "another child item"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "with a sub child item"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("unordered list with 2 items with empty line in-between", func() {
			// fist line after list item is swallowed
			source := "* a first item\n" +
				"\n" +
				"* a second item with *bold content*"
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "a first item"},
									},
								},
							},
						},
					},
					types.BlankLine{},
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
										types.StringElement{Content: "a second item with "},
										types.QuotedText{
											Kind: types.Bold,
											Elements: types.InlineElements{
												types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})
		It("unordered list with 2 items on multiple lines", func() {
			source := `* item 1
  on 2 lines.
* item 2
on 2 lines, too.`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "item 1"},
									},
									{
										types.StringElement{Content: "  on 2 lines."},
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
										types.StringElement{Content: "item 2"},
									},
									{
										types.StringElement{Content: "on 2 lines, too."},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})
		It("unordered lists with 2 empty lines in-between", func() {
			source := `* an item in the first list
			

* an item in the second list`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "an item in the first list"},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.BlankLine{},
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
										types.StringElement{Content: "an item in the second list"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected)) // parse the whole document to get 2 lists
		})

		It("unordered list with items on 3 levels", func() {
			source := `* item 1
	** item 1.1
	** item 1.2
	*** item 1.2.1
	** item 1.3
	** item 1.4
	* item 2
	** item 2.1`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "item 1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1.1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1.2"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1.2.1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1.3"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1.4"},
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
										types.StringElement{Content: "item 2"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 2.1"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("max level of unordered items - case 1", func() {
			source := `.Unordered, max nesting
* level 1
** level 2
*** level 3
**** level 4
***** level 5
* level 1`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.UnorderedListItem{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Unordered, max nesting",
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 1",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 2",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 3",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       4,
						BulletStyle: types.FourAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 4",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       5,
						BulletStyle: types.FiveAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 5",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
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
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("max level of unordered items - case 2", func() {
			source := `.Unordered, max nesting
* level 1
** level 2
*** level 3
**** level 4
***** level 5
** level 2`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.UnorderedListItem{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Unordered, max nesting",
						},
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 1",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 2",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 3",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{Level: 4,
						BulletStyle: types.FourAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 4",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       5,
						BulletStyle: types.FiveAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 5",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 2",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})
	})

	Context("invalid content", func() {
		It("unordered list with items on 2 levels - bad numbering", func() {
			source := `* item 1
					*** item 1.1
					*** item 1.1.1
					** item 1.2
					* item 2`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "item 1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1.1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1.1.1"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "item 1.2"},
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
										types.StringElement{Content: "item 2"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("invalid list item", func() {
			source := "*an invalid list item"
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "*an invalid list item"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})
	})

	Context("list item continuation", func() {

		It("unordered list with item continuation - case 1", func() {
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
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
									Lines: []types.InlineElements{
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
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("unordered list with item continuation - case 2", func() {
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
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.UnorderedListItem{
						Level:       1,
						BulletStyle: types.OneAsterisk,
						CheckStyle:  types.NoCheck,
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Unordered, complex",
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 1",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 2",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 3",
										},
									},
									{
										types.StringElement{
											Content: "This is a new line inside an unordered list using ",
										},
										types.DocumentAttributeSubstitution{
											Name: "plus",
										},
										types.StringElement{
											Content: " symbol.",
										},
									},
									{
										types.StringElement{
											Content: "We can even force content to start on a separate line...",
										},
										types.LineBreak{},
									},
									{
										types.StringElement{
											Content: "Amazing, isn't it?",
										},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Level:       4,
						BulletStyle: types.FourAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "level 4",
										},
									},
								},
							},
						},
					},
					// the `+` continuation produces the second paragraph below
					types.ContinuedListItemElement{
						Offset: 0,
						Element: types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "The ",
									},
									types.DocumentAttributeSubstitution{
										Name: "plus",
									},
									types.StringElement{
										Content: " symbol is on a new line.",
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.UnorderedListItem{
						Level:       5,
						BulletStyle: types.FiveAsterisks,
						CheckStyle:  types.NoCheck,
						Attributes:  types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
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
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("unordered list without item continuation", func() {
			source := `* foo
----
a delimited block
----
* bar
----
another delimited block
----`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "a delimited block",
										},
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
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
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
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})
	})

	Context("attach to ancestor", func() {

		It("attach to grandparent item", func() {
			source := `* grand parent list item
** parent list item
*** child list item


+
paragraph attached to grand parent list item`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "grand parent list item"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "parent list item"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "child list item"},
									},
								},
							},
						},
					},
					types.ContinuedListItemElement{
						Offset: -2,
						Element: types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "paragraph attached to grand parent list item"},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("attach to parent item", func() {
			source := `* grandparent list item
** parent list item
*** child list item

+
paragraph attached to parent list item`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
										types.StringElement{Content: "grandparent list item"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       2,
						BulletStyle: types.TwoAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "parent list item"},
									},
								},
							},
						},
					},
					types.UnorderedListItem{
						Attributes:  types.ElementAttributes{},
						Level:       3,
						BulletStyle: types.ThreeAsterisks,
						CheckStyle:  types.NoCheck,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "child list item"},
									},
								},
							},
						},
					},
					types.ContinuedListItemElement{
						Offset: -1,
						Element: types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "paragraph attached to parent list item"},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualPreflightDocument(expected))
		})
	})
})

var _ = Describe("unordered lists - document", func() {

	Context("valid content", func() {

		It("unordered list with a basic single item", func() {
			source := `* a list item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "a list item"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("unordered list with ID, title, role and a single item", func() {
			source := `.mytitle
[#listID]
[.myrole]
* a list item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{
							types.AttrID:       "listID",
							types.AttrCustomID: true,
							types.AttrTitle:    "mytitle",
							types.AttrRole:     "myrole",
						},
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
												types.StringElement{Content: "a list item"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})
		It("unordered list with a title and a single item", func() {
			source := `.a title
	* a list item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "a title",
						},
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
												types.StringElement{Content: "a list item"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("unordered list with 2 items with stars", func() {
			source := `* a first item
					* a second item with *bold content*`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "a first item"},
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
												types.StringElement{Content: "a second item with "},
												types.QuotedText{
													Kind: types.Bold,
													Elements: types.InlineElements{
														types.StringElement{Content: "bold content"},
													},
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
			Expect(source).To(EqualDocument(expected))
		})

		It("unordered list based on article.adoc (with heading spaces)", func() {
			source := `.Unordered list title
		* list item 1
		** nested list item A
		*** nested nested list item A.1
		*** nested nested list item A.2
		** nested list item B
		*** nested nested list item B.1
		*** nested nested list item B.2
		* list item 2`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Unordered list title",
						},
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
												types.StringElement{Content: "list item 1"},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "nested list item A"},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "nested nested list item A.1"},
																			},
																		},
																	},
																},
															},
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "nested nested list item A.2"},
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
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "nested list item B"},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "nested nested list item B.1"},
																			},
																		},
																	},
																},
															},
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "nested nested list item B.2"},
																			},
																		},
																	},
																},
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
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "list item 2"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("unordered list with 2 items with carets", func() {
			source := "- a first item\n" +
				"- a second item with *bold content*"
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "a first item"},
											},
										},
									},
								},
							},
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
												types.StringElement{Content: "a second item with "},
												types.QuotedText{
													Kind: types.Bold,
													Elements: types.InlineElements{
														types.StringElement{Content: "bold content"},
													},
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
			Expect(source).To(EqualDocument(expected))
		})

		It("unordered list with items with mixed styles", func() {
			source := `- a parent item
					* a child item
					- another parent item
					* another child item
					** with a sub child item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "a parent item"},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.OneAsterisk,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "a child item"},
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
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "another parent item"},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.OneAsterisk,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "another child item"},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.TwoAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "with a sub child item"},
																			},
																		},
																	},
																},
															},
														},
													},
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
			Expect(source).To(EqualDocument(expected))
		})

		It("unordered list with 2 items with empty line in-between", func() {
			// fist line after list item is swallowed
			source := "* a first item\n" +
				"\n" +
				"* a second item with *bold content*"
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "a first item"},
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
												types.StringElement{Content: "a second item with "},
												types.QuotedText{
													Kind: types.Bold,
													Elements: types.InlineElements{
														types.StringElement{Content: "bold content"},
													},
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
			Expect(source).To(EqualDocument(expected))
		})
		It("unordered list with 2 items on multiple lines", func() {
			source := `* item 1
  on 2 lines.
* item 2
on 2 lines, too.`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "item 1"},
											},
											{
												types.StringElement{Content: "  on 2 lines."},
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
												types.StringElement{Content: "item 2"},
											},
											{
												types.StringElement{Content: "on 2 lines, too."},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})
		It("unordered lists with 2 empty lines in-between", func() {
			// the first blank lines after the first list is swallowed (for the list item)
			source := "* an item in the first list\n" +
				"\n" +
				"\n" +
				"* an item in the second list"
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "an item in the first list"},
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
												types.StringElement{Content: "an item in the second list"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected)) // parse the whole document to get 2 lists
		})

		It("unordered list with items on 3 levels", func() {
			source := `* item 1
	** item 1.1
	** item 1.2
	*** item 1.2.1
	** item 1.3
	** item 1.4
	* item 2
	** item 2.1`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "item 1"},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "item 1.1"},
															},
														},
													},
												},
											},
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "item 1.2"},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "item 1.2.1"},
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
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "item 1.3"},
															},
														},
													},
												},
											},
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "item 1.4"},
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
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "item 2"},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
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
			Expect(source).To(EqualDocument(expected))
		})

		It("max level of unordered items - case 1", func() {
			source := `.Unordered, max nesting
* level 1
** level 2
*** level 3
**** level 4
***** level 5
* level 1`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Unordered, max nesting",
						},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Attributes:  types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "level 1",
												},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Attributes:  types.ElementAttributes{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{
																	Content: "level 2",
																},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Attributes:  types.ElementAttributes{},
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{
																					Content: "level 3",
																				},
																			},
																		},
																	},
																	types.UnorderedList{
																		Attributes: types.ElementAttributes{},
																		Items: []types.UnorderedListItem{
																			{
																				Level:       4,
																				BulletStyle: types.FourAsterisks,
																				CheckStyle:  types.NoCheck,
																				Attributes:  types.ElementAttributes{},
																				Elements: []interface{}{
																					types.Paragraph{
																						Attributes: types.ElementAttributes{},
																						Lines: []types.InlineElements{
																							{
																								types.StringElement{
																									Content: "level 4",
																								},
																							},
																						},
																					},
																					types.UnorderedList{
																						Attributes: types.ElementAttributes{},
																						Items: []types.UnorderedListItem{
																							{
																								Level:       5,
																								BulletStyle: types.FiveAsterisks,
																								CheckStyle:  types.NoCheck,
																								Attributes:  types.ElementAttributes{},
																								Elements: []interface{}{
																									types.Paragraph{
																										Attributes: types.ElementAttributes{},
																										Lines: []types.InlineElements{
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
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Attributes:  types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
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
			Expect(source).To(EqualDocument(expected))
		})

		It("max level of unordered items - case 2", func() {
			source := `.Unordered, max nesting
* level 1
** level 2
*** level 3
**** level 4
***** level 5
** level 2`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Unordered, max nesting",
						},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Attributes:  types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "level 1",
												},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Attributes:  types.ElementAttributes{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{
																	Content: "level 2",
																},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Attributes:  types.ElementAttributes{},
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{
																					Content: "level 3",
																				},
																			},
																		},
																	},
																	types.UnorderedList{
																		Attributes: types.ElementAttributes{},
																		Items: []types.UnorderedListItem{
																			{
																				Level:       4,
																				BulletStyle: types.FourAsterisks,
																				CheckStyle:  types.NoCheck,
																				Attributes:  types.ElementAttributes{},
																				Elements: []interface{}{
																					types.Paragraph{
																						Attributes: types.ElementAttributes{},
																						Lines: []types.InlineElements{
																							{
																								types.StringElement{
																									Content: "level 4",
																								},
																							},
																						},
																					},
																					types.UnorderedList{
																						Attributes: types.ElementAttributes{},
																						Items: []types.UnorderedListItem{
																							{
																								Level:       5,
																								BulletStyle: types.FiveAsterisks,
																								CheckStyle:  types.NoCheck,
																								Attributes:  types.ElementAttributes{},
																								Elements: []interface{}{
																									types.Paragraph{
																										Attributes: types.ElementAttributes{},
																										Lines: []types.InlineElements{
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
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Attributes:  types.ElementAttributes{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{
																	Content: "level 2",
																},
															},
														},
													},
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
			Expect(source).To(EqualDocument(expected))
		})
	})

	Context("invalid content", func() {
		It("unordered list with items on 2 levels - bad numbering", func() {
			source := `* item 1
					*** item 1.1
					*** item 1.1.1
					** item 1.2
					* item 2`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "item 1"},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "item 1.1"},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "item 1.1.1"},
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
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
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
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
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
			Expect(source).To(EqualDocument(expected))
		})

		It("invalid list item", func() {
			source := "*an invalid list item"
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "*an invalid list item"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})
	})

	Context("list item continuation", func() {

		It("unordered list with item continuation - case 1", func() {
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
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
									types.DelimitedBlock{
										Attributes: types.ElementAttributes{},
										Kind:       types.Listing,
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
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
												Lines: []types.InlineElements{
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
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("unordered list with item continuation - case 2", func() {
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
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "Unordered, complex",
						},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Attributes:  types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "level 1",
												},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Attributes:  types.ElementAttributes{},
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{
																	Content: "level 2",
																},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Attributes:  types.ElementAttributes{},
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{
																					Content: "level 3",
																				},
																			},
																			{
																				types.StringElement{
																					Content: "This is a new line inside an unordered list using ",
																				},
																				types.DocumentAttributeSubstitution{
																					Name: "plus",
																				},
																				types.StringElement{
																					Content: " symbol.",
																				},
																			},
																			{
																				types.StringElement{
																					Content: "We can even force content to start on a separate line...",
																				},
																				types.LineBreak{},
																			},
																			{
																				types.StringElement{
																					Content: "Amazing, isn't it?",
																				},
																			},
																		},
																	},
																	types.UnorderedList{
																		Attributes: types.ElementAttributes{},
																		Items: []types.UnorderedListItem{
																			{
																				Level:       4,
																				BulletStyle: types.FourAsterisks,
																				CheckStyle:  types.NoCheck,
																				Attributes:  types.ElementAttributes{},
																				Elements: []interface{}{
																					types.Paragraph{
																						Attributes: types.ElementAttributes{},
																						Lines: []types.InlineElements{
																							{
																								types.StringElement{
																									Content: "level 4",
																								},
																							},
																						},
																					},
																					// the `+` continuation produces the second paragrap below
																					types.Paragraph{
																						Attributes: types.ElementAttributes{},
																						Lines: []types.InlineElements{
																							{
																								types.StringElement{
																									Content: "The ",
																								},
																								types.DocumentAttributeSubstitution{
																									Name: "plus",
																								},
																								types.StringElement{
																									Content: " symbol is on a new line.",
																								},
																							},
																						},
																					},

																					types.UnorderedList{
																						Attributes: types.ElementAttributes{},
																						Items: []types.UnorderedListItem{
																							{
																								Level:       5,
																								BulletStyle: types.FiveAsterisks,
																								CheckStyle:  types.NoCheck,
																								Attributes:  types.ElementAttributes{},
																								Elements: []interface{}{
																									types.Paragraph{
																										Attributes: types.ElementAttributes{},
																										Lines: []types.InlineElements{
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
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("unordered list without item continuation", func() {
			source := `* foo
----
a delimited block
----
* bar
----
another delimited block
----`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
											Content: "a delimited block",
										},
									},
								},
							},
						},
					},
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
												types.StringElement{Content: "bar"},
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
											Content: "another delimited block",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})
	})

	Context("attach to ancestor", func() {

		It("attach to grandparent item", func() {
			source := `* grand parent list item
** parent list item
*** child list item


+
paragraph attached to grand parent list item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "grand parent list item"},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "parent list item"},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "child list item"},
																			},
																		},
																	},
																},
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
												types.StringElement{Content: "paragraph attached to grand parent list item"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("attach to parent item", func() {
			source := `* grandparent list item
** parent list item
*** child list item

+
paragraph attached to parent list item`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
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
												types.StringElement{Content: "grandparent list item"},
											},
										},
									},
									types.UnorderedList{
										Attributes: types.ElementAttributes{},
										Items: []types.UnorderedListItem{
											{
												Attributes:  types.ElementAttributes{},
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.ElementAttributes{},
														Lines: []types.InlineElements{
															{
																types.StringElement{Content: "parent list item"},
															},
														},
													},
													types.UnorderedList{
														Attributes: types.ElementAttributes{},
														Items: []types.UnorderedListItem{
															{
																Attributes:  types.ElementAttributes{},
																Level:       3,
																BulletStyle: types.ThreeAsterisks,
																CheckStyle:  types.NoCheck,
																Elements: []interface{}{
																	types.Paragraph{
																		Attributes: types.ElementAttributes{},
																		Lines: []types.InlineElements{
																			{
																				types.StringElement{Content: "child list item"},
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
																types.StringElement{Content: "paragraph attached to parent list item"},
															},
														},
													},
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
			Expect(source).To(EqualDocument(expected))
		})
	})
})
