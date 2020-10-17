package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("checked lists", func() {

	Context("final documents", func() {

		It("checklist with title and dashes", func() {
			source := `.Checklist
- [*] checked
- [x] also checked
- [ ] not checked
-     normal list item`
			expected := types.Document{
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.Attributes{
							types.AttrTitle: "Checklist",
						},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.Checked,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Checked,
										},
										Lines: [][]interface{}{
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
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.Checked,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Checked,
										},
										Lines: [][]interface{}{
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
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.Unchecked,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Unchecked,
										},
										Lines: [][]interface{}{
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
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
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
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("parent checklist with title and nested checklist", func() {
			source := `.Checklist
* [ ] parent not checked
** [*] checked
** [x] also checked
** [ ] not checked
*     normal list item`
			expected := types.Document{
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.Attributes{
							types.AttrTitle: "Checklist",
						},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.Unchecked,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Unchecked,
										},
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "parent not checked",
												},
											},
										},
									},
									types.UnorderedList{
										Items: []types.UnorderedListItem{
											{
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.Checked,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.Attributes{
															types.AttrCheckStyle: types.Checked,
														},
														Lines: [][]interface{}{
															{types.StringElement{
																Content: "checked",
															},
															},
														},
													},
												},
											},
											{
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.Checked,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.Attributes{
															types.AttrCheckStyle: types.Checked,
														},
														Lines: [][]interface{}{
															{types.StringElement{
																Content: "also checked",
															},
															},
														},
													},
												},
											},
											{
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.Unchecked,
												Elements: []interface{}{
													types.Paragraph{
														Attributes: types.Attributes{
															types.AttrCheckStyle: types.Unchecked,
														},
														Lines: [][]interface{}{
															{types.StringElement{
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
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
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
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("parent checklist with title and nested normal list", func() {
			source := `.Checklist
* [ ] parent not checked
** a normal list item
** another normal list item
*     normal list item`
			expected := types.Document{
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.Attributes{
							types.AttrTitle: "Checklist",
						},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.Unchecked,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.Attributes{
											types.AttrCheckStyle: types.Unchecked,
										},
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "parent not checked",
												},
											},
										},
									},
									types.UnorderedList{
										Items: []types.UnorderedListItem{
											{
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
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
												Level:       2,
												BulletStyle: types.TwoAsterisks,
												CheckStyle:  types.NoCheck,
												Elements: []interface{}{
													types.Paragraph{
														Lines: [][]interface{}{
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
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
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
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
