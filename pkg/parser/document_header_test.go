package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("headers", func() {

	Context("in final documents", func() {

		Context("valid cases", func() {

			It("header alone", func() {
				source := `= Title
			
This journey continues.`

				Title := []interface{}{
					&types.StringElement{Content: "Title"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: Title,
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "This journey continues."},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header with attributes", func() {
				source := `[.role1#anchor.role2]
= Title
			
This journey continues.`

				Title := []interface{}{
					&types.StringElement{Content: "Title"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: Title,
							Attributes: types.Attributes{
								types.AttrRoles:    types.Roles{"role1", "role2"},
								types.AttrID:       "anchor",
								types.AttrCustomID: true,
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "This journey continues."},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header with attribute declarations and resets", func() {
				source := `= A Title
:author: Xavier
:version-label!:
`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{
									Content: "A Title",
								},
							},
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name: "authors",
									Value: types.DocumentAuthors{
										{
											DocumentAuthorFullName: &types.DocumentAuthorFullName{
												FirstName: "Xavier",
											},
										},
									},
								},
								&types.AttributeReset{
									Name: "version-label",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with authors", func() {

				Context("single author", func() {

					It("all author data with extra spaces", func() {
						source := `= Title
John  Foo    Doe  <johndoe@example.com>`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName:  "John",
														MiddleName: "Foo",
														LastName:   "Doe",
													},
													Email: "johndoe@example.com",
												},
											},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("lastname with underscores", func() {
						source := `= Title
Jane the_Doe <jane@example.com>`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName: "Jane",
														LastName:  "the Doe",
													},
													Email: "jane@example.com",
												},
											},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("with middlename and composed lastname", func() {
						source := `= Title
Jane Foo the Doe <jane@example.com>`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName:  "Jane",
														MiddleName: "Foo",
														LastName:   "the Doe",
													},
													Email: "jane@example.com",
												},
											},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("firstname and lastname only", func() {
						source := `= Title
John Doe`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName: "John",
														LastName:  "Doe",
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

					It("firstname only", func() {
						source := `= Title
Doe`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName: "Doe",
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

					It("alternate author input", func() {
						source := `= Title
:author: John Foo Doe` // `:"email":` is processed as a regular attribute
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName:  "John",
														MiddleName: "Foo",
														LastName:   "Doe",
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

				Context("multiple authors", func() {

					It("2 authors", func() {
						source := `= Title
John  Foo Doe  <johndoe@example.com>; Jane the_Doe <jane@example.com>`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName:  "John",
														MiddleName: "Foo",
														LastName:   "Doe",
													},
													Email: "johndoe@example.com",
												},
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName: "Jane",
														LastName:  "the Doe",
													},
													Email: "jane@example.com",
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

				Context("authors and comments", func() {

					It("authors commented out", func() {
						source := `= Title
// John  Foo Doe  <johndoe@example.com>; Jane the_Doe <jane@example.com>`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title:    Title,
									Elements: nil, // single comment is filtered out
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("authors after a single comment line", func() {
						source := `= Title
// a comment
John  Foo Doe  <johndoe@example.com>; Jane the_Doe <jane@example.com>`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName:  "John",
														MiddleName: "Foo",
														LastName:   "Doe",
													},
													Email: "johndoe@example.com",
												},
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName: "Jane",
														LastName:  "the Doe",
													},
													Email: "jane@example.com",
												},
											},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("authors after a comment block", func() {
						source := `= Title
//// 
a comment
////
John  Foo Doe  <johndoe@example.com>; Jane the_Doe <jane@example.com>`
						Title := []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						}
						expected := &types.Document{
							Elements: []interface{}{
								&types.DocumentHeader{
									Title: Title,
									Elements: []interface{}{
										&types.AttributeDeclaration{
											Name: types.AttrAuthors,
											Value: types.DocumentAuthors{
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName:  "John",
														MiddleName: "Foo",
														LastName:   "Doe",
													},
													Email: "johndoe@example.com",
												},
												{
													DocumentAuthorFullName: &types.DocumentAuthorFullName{
														FirstName: "Jane",
														LastName:  "the Doe",
													},
													Email: "jane@example.com",
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

				It("with author used in a paragraph", func() {
					source := `= Title
Xavier <xavier@example.com>

written by {author}`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "Xavier",
												},
												Email: "xavier@example.com",
											},
										},
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "written by Xavier", //
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			Context("with revisions", func() {

				It("full document revision without any comment", func() {
					source := `= Title
John Doe
v1.0, March 29, 2020: Updated revision`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
											Revdate:   "March 29, 2020",
											Revremark: "Updated revision",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("full document revision with various comments", func() {
					source := `= Title
// a single-line comment
John Doe
////
a comment block

with an empty line
////
v1.0, March 29, 2020: Updated revision
////
another comment block

with another empty line
////
`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
											Revdate:   "March 29, 2020",
											Revremark: "Updated revision",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("full document revision with a comment before author", func() {
					source := `= Title
// a comment
John Doe
v1.0, March 29, 2020: Updated revision`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
											Revdate:   "March 29, 2020",
											Revremark: "Updated revision",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("full document revision with a singleline comment before revision", func() {
					source := `= Title
John Doe
// a comment
v1.0, March 29, 2020: Updated revision`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
											Revdate:   "March 29, 2020",
											Revremark: "Updated revision",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with revnumber and revdate only", func() {
					source := `= Title
				John Doe
				v1.0, March 29, 2020`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
											Revdate:   "March 29, 2020",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with revnumber and revdate - with colon separator", func() {
					source := `= Title
				John Doe
				1.0, March 29, 2020:`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
											Revdate:   "March 29, 2020",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with revnumber only - comma suffix", func() {
					source := `= Title
				John Doe
				1.0,`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with revdate as number - spaces and no prefix no suffix", func() {
					source := `= Title
				John Doe
				1.0`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revdate: "1.0",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with revdate as alphanum - spaces and no prefix no suffix", func() {
					source := `= Title
				John Doe
				1.0a`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revdate: "1.0a",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with revnumber only", func() {
					source := `= Title
				John Doe
				v1.0:`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with spaces and capital revnumber ", func() {
					source := `= Title
				John Doe
				V1.0:`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision only - with comma separator", func() {
					source := `= Title
				John Doe
				v1.0,`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with revnumber plus comma and colon separators", func() {
					source := `= Title
				John Doe
				v1.0,:`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("revision with revnumber and empty revremark", func() {
					source := `= Title
John Doe
v1.0:`
					Title := []interface{}{
						&types.StringElement{
							Content: "Title",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: Title,
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "John",
													LastName:  "Doe",
												},
											},
										},
									},
									&types.AttributeDeclaration{
										Name: types.AttrRevision,
										Value: &types.DocumentRevision{
											Revnumber: "1.0",
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

		Context("document header attributes", func() {

			It("valid attribute names", func() {
				source := `:a:
:author: Xavier
:_author: Xavier`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name: "a",
								},
								&types.AttributeDeclaration{
									Name:  "author",
									Value: "Xavier",
								},
								&types.AttributeDeclaration{
									Name:  "_author",
									Value: "Xavier",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("attributes and paragraph without blank line in-between", func() {
				source := `:toc:
:date:  2017-01-01
:author: Xavier
:hardbreaks:
a paragraph`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name: "toc",
								},
								&types.AttributeDeclaration{
									Name:  "date",
									Value: "2017-01-01",
								},
								&types.AttributeDeclaration{
									Name:  "author",
									Value: "Xavier",
								},
								&types.AttributeDeclaration{
									Name: "hardbreaks",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a paragraph",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("contiguous attributes and paragraph with blank line in-between", func() {
				source := `:toc:
:date: 2017-01-01
:author: Xavier

a paragraph`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name: "toc",
								},
								&types.AttributeDeclaration{
									Name:  "date",
									Value: "2017-01-01",
								},
								&types.AttributeDeclaration{
									Name:  "author",
									Value: "Xavier",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("splitted attributes and paragraph with blank line in-between", func() {
				source := `:toc:
:date: 2017-01-01

:author: Xavier

:hardbreaks:

a paragraph`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name: "toc",
								},
								&types.AttributeDeclaration{
									Name:  "date",
									Value: "2017-01-01",
								},
								&types.AttributeDeclaration{
									Name:  "author",
									Value: "Xavier",
								},
								&types.AttributeDeclaration{
									Name: "hardbreaks",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("no header and attributes in body", func() {
				source := `a paragraph

:toc:
:date: 2017-01-01
:author: Xavier`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a paragraph"},
							},
						},
						&types.AttributeDeclaration{
							Name: "toc",
						},
						&types.AttributeDeclaration{
							Name:  "date",
							Value: "2017-01-01",
						},
						&types.AttributeDeclaration{
							Name:  "author",
							Value: "Xavier",
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with soft-wrapping", func() {

				It("alone without indentation", func() {
					source := `:description: a long \
description on \
multiple \
lines.`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "description",
										Value: "a long description on multiple lines.",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with other attributes without indentation", func() {
					source := `:hardbreaks:
:description: a long \
description on \
multiple \
lines.
:author: Xavier`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: "hardbreaks",
									},
									&types.AttributeDeclaration{
										Name:  "description",
										Value: "a long description on multiple lines.",
									},
									&types.AttributeDeclaration{
										Name:  "author",
										Value: "Xavier",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with other attributes and with variable indentation", func() {
					source := `:hardbreaks:
:description: a long \
    description on \
      multiple \
    lines.
:author: Xavier`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: "hardbreaks",
									},
									&types.AttributeDeclaration{
										Name:  "description",
										Value: "a long description on multiple lines.",
									},
									&types.AttributeDeclaration{
										Name:  "author",
										Value: "Xavier",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("invalid cases", func() {

			It("paragraph without blank line before attribute declarations", func() {
				source := `a paragraph
:toc:
:date: 2017-01-01
:author: Xavier`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: `a paragraph
:toc:
:date: 2017-01-01
:author: Xavier`,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid attribute names", func() {
				source := `:@date: 2017-01-01
:{author}: Xavier`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: ":@date: 2017-01-01\n:{author}: Xavier", // attribute susbtitution "failed"
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("with overrides", func() {

			It("custom icon attribute", func() {
				// given
				attrs := map[string]interface{}{
					"icons":              "font",
					"source-highlighter": "pygments",
				}
				source := `{icons}`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "font"},
							},
						},
					},
				}
				Expect(ParseDocument(source, configuration.WithAttributes(attrs))).To(MatchDocument(expected))
			})
		})
	})
})
