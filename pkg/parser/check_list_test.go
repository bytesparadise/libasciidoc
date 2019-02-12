package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("checked lists", func() {

	It("checklist with title and dashes", func() {
		actualContent := `.Checklist
- [*] checked
- [x] also checked
- [ ] not checked
-     normal list item`
		expectedResult := types.UnorderedList{
			Attributes: types.ElementAttributes{
				types.AttrTitle: "Checklist",
			},
			Items: []types.UnorderedListItem{
				{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.Dash,
					CheckStyle:  types.Checked,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{
								types.AttrCheckStyle: types.Checked,
							},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "checked",
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
					CheckStyle:  types.Checked,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{
								types.AttrCheckStyle: types.Checked,
							},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "also checked",
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
					CheckStyle:  types.Unchecked,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{
								types.AttrCheckStyle: types.Unchecked,
							},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "not checked",
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
									types.StringElement{
										Content: "normal list item",
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

	It("parent checklist with title and nested checklist", func() {
		actualContent := `.Checklist
* [ ] parent not checked
** [*] checked
** [x] also checked
** [ ] not checked
*     normal list item`
		expectedResult := types.UnorderedList{
			Attributes: types.ElementAttributes{
				types.AttrTitle: "Checklist",
			},
			Items: []types.UnorderedListItem{
				{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					CheckStyle:  types.Unchecked,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{
								types.AttrCheckStyle: types.Unchecked,
							},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "parent not checked",
									},
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
									CheckStyle:  types.Checked,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{
												types.AttrCheckStyle: types.Checked,
											},
											Lines: []types.InlineElements{
												{
													types.StringElement{
														Content: "checked",
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
									CheckStyle:  types.Checked,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{
												types.AttrCheckStyle: types.Checked,
											},
											Lines: []types.InlineElements{
												{
													types.StringElement{
														Content: "also checked",
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
									CheckStyle:  types.Unchecked,
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{
												types.AttrCheckStyle: types.Unchecked,
											},
											Lines: []types.InlineElements{
												{
													types.StringElement{
														Content: "not checked",
													},
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
									types.StringElement{
										Content: "normal list item",
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

	It("parent checklist with title and nested normal list", func() {
		actualContent := `.Checklist
* [ ] parent not checked
** a normal list item
** another normal list item
*     normal list item`
		expectedResult := types.UnorderedList{
			Attributes: types.ElementAttributes{
				types.AttrTitle: "Checklist",
			},
			Items: []types.UnorderedListItem{
				{
					Attributes:  types.ElementAttributes{},
					Level:       1,
					BulletStyle: types.OneAsterisk,
					CheckStyle:  types.Unchecked,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{
								types.AttrCheckStyle: types.Unchecked,
							},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "parent not checked",
									},
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
													types.StringElement{
														Content: "a normal list item",
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
													types.StringElement{
														Content: "another normal list item",
													},
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
									types.StringElement{
										Content: "normal list item",
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
