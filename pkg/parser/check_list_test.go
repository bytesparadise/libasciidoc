package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golint
)

var _ = Describe("checked lists", func() {

	Context("in final documents", func() {

		It("with title and dashes", func() {
			source := `.Checklist
- [*] checked
- [x] also checked
- [ ] not checked
-     normal list item`
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.UnorderedListKind,
						Attributes: types.Attributes{
							types.AttrTitle: "Checklist",
						},
						Elements: []types.ListElement{
							&types.UnorderedListElement{
								BulletStyle: types.Dash,
								CheckStyle:  types.Checked,
								Elements: []interface{}{
									&types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Checked,
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "checked",
											},
										},
									},
								},
							},
							&types.UnorderedListElement{
								BulletStyle: types.Dash,
								CheckStyle:  types.Checked,
								Elements: []interface{}{
									&types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Checked,
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "also checked",
											},
										},
									},
								},
							},
							&types.UnorderedListElement{
								BulletStyle: types.Dash,
								CheckStyle:  types.Unchecked,
								Elements: []interface{}{
									&types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Unchecked,
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "not checked",
											},
										},
									},
								},
							},
							&types.UnorderedListElement{
								BulletStyle: types.Dash,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "normal list item",
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

		It("with interactive checkboxes", func() {
			source := `[%interactive]
	* [*] checked
	* [x] also checked
	* [ ] not checked
	*     normal list item`
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.UnorderedListKind,
						Attributes: types.Attributes{
							types.AttrOptions: types.Options{
								types.AttrInteractive,
							},
						},
						Elements: []types.ListElement{
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.CheckedInteractive,
								Elements: []interface{}{
									&types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.CheckedInteractive,
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "checked",
											},
										},
									},
								},
							},
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.CheckedInteractive,
								Elements: []interface{}{
									&types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.CheckedInteractive,
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "also checked",
											},
										},
									},
								},
							},
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.UncheckedInteractive,
								Elements: []interface{}{
									&types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.UncheckedInteractive,
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "not checked",
											},
										},
									},
								},
							},
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "normal list item",
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

		It("with title and nested checklist", func() {
			source := `.Checklist
* [ ] parent not checked
** [*] checked
** [x] also checked
** [ ] not checked
*     normal list item`
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.UnorderedListKind,
						Attributes: types.Attributes{
							types.AttrTitle: "Checklist",
						},
						Elements: []types.ListElement{
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.Unchecked,
								Elements: []interface{}{
									&types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Unchecked,
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "parent not checked",
											},
										},
									},
									&types.List{
										Kind: types.UnorderedListKind,
										Elements: []types.ListElement{
											&types.UnorderedListElement{
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.Checked,
												Elements: []interface{}{
													&types.Paragraph{
														Attributes: types.Attributes{
															types.AttrCheckStyle: types.Checked,
														},
														Elements: []interface{}{
															&types.StringElement{
																Content: "checked",
															},
														},
													},
												},
											},
											&types.UnorderedListElement{
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.Checked,
												Elements: []interface{}{
													&types.Paragraph{
														Attributes: types.Attributes{
															types.AttrCheckStyle: types.Checked,
														},
														Elements: []interface{}{
															&types.StringElement{
																Content: "also checked",
															},
														},
													},
												},
											},
											&types.UnorderedListElement{
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.Unchecked,
												Elements: []interface{}{
													&types.Paragraph{
														Attributes: types.Attributes{
															types.AttrCheckStyle: types.Unchecked,
														},
														Elements: []interface{}{
															&types.StringElement{
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
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "normal list item",
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

		It("with title and nested normal list", func() {
			source := `.Checklist
* [ ] parent not checked
** a normal list item
** another normal list item
*     normal list item`
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.UnorderedListKind,
						Attributes: types.Attributes{
							types.AttrTitle: "Checklist",
						},
						Elements: []types.ListElement{
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.Unchecked,
								Elements: []interface{}{
									&types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Unchecked,
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "parent not checked",
											},
										},
									},
									&types.List{
										Kind: types.UnorderedListKind,
										Elements: []types.ListElement{
											&types.UnorderedListElement{
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													&types.Paragraph{
														Elements: []interface{}{
															&types.StringElement{
																Content: "a normal list item",
															},
														},
													},
												},
											},
											&types.UnorderedListElement{
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													&types.Paragraph{
														Elements: []interface{}{
															&types.StringElement{
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
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "normal list item",
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
