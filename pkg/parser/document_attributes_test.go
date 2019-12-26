package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document attributes", func() {

	Context("valid Document Header", func() {

		It("header alone", func() {
			source := `= The Dangerous and Thrilling Documentation Chronicles
			
This journey begins on a bleary Monday morning.`

			title := []interface{}{
				types.StringElement{Content: "The Dangerous and Thrilling Documentation Chronicles"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"the_dangerous_and_thrilling_documentation_chronicles": title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrID:       "the_dangerous_and_thrilling_documentation_chronicles",
							types.AttrCustomID: false,
						},
						Title: title,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "This journey begins on a bleary Monday morning."},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		Context("document authors", func() {

			Context("single author", func() {

				It("all author data", func() {
					source := `= title
Kismet  Rainbow Chameleon  <kismet@asciidoctor.org>`
					title := []interface{}{
						types.StringElement{
							Content: "title",
						},
					}
					expected := types.Document{
						Attributes: types.DocumentAttributes{},
						ElementReferences: types.ElementReferences{
							"title": title,
						},
						Footnotes:          types.Footnotes{},
						FootnoteReferences: types.FootnoteReferences{},
						Elements: []interface{}{
							types.Section{
								Level: 0,
								Attributes: types.ElementAttributes{
									types.AttrID:       "title",
									types.AttrCustomID: false,
									types.AttrAuthors: []types.DocumentAuthor{
										{
											FullName: "Kismet  Rainbow Chameleon  ",
											Email:    "kismet@asciidoctor.org",
										},
									},
								},
								Title:    title,
								Elements: []interface{}{},
							},
						},
					}
					Expect(source).To(BecomeDocument(expected))
				})

				It("lastname with underscores", func() {
					source := `= title
Lazarus het_Draeke <lazarus@asciidoctor.org>`
					title := []interface{}{
						types.StringElement{
							Content: "title",
						},
					}
					expected := types.Document{
						Attributes: types.DocumentAttributes{},
						ElementReferences: types.ElementReferences{
							"title": title,
						},
						Footnotes:          types.Footnotes{},
						FootnoteReferences: types.FootnoteReferences{},
						Elements: []interface{}{
							types.Section{
								Level: 0,
								Attributes: types.ElementAttributes{
									types.AttrID:       "title",
									types.AttrCustomID: false,
									types.AttrAuthors: []types.DocumentAuthor{
										{
											FullName: "Lazarus het_Draeke ",
											Email:    "lazarus@asciidoctor.org",
										},
									},
								},
								Title:    title,
								Elements: []interface{}{},
							},
						},
					}
					Expect(source).To(BecomeDocument(expected))
				})

				It("firstname and lastname only", func() {
					source := `= title
Kismet Chameleon`
					title := []interface{}{
						types.StringElement{
							Content: "title",
						},
					}
					expected := types.Document{
						Attributes: types.DocumentAttributes{},
						ElementReferences: types.ElementReferences{
							"title": title,
						},
						Footnotes:          types.Footnotes{},
						FootnoteReferences: types.FootnoteReferences{},
						Elements: []interface{}{
							types.Section{
								Level: 0,
								Attributes: types.ElementAttributes{
									types.AttrID:       "title",
									types.AttrCustomID: false,
									types.AttrAuthors: []types.DocumentAuthor{
										{
											FullName: "Kismet Chameleon",
											Email:    "",
										},
									},
								},
								Title:    title,
								Elements: []interface{}{},
							},
						},
					}
					Expect(source).To(BecomeDocument(expected))
				})

				It("firstname only", func() {
					source := `= title
Chameleon`
					title := []interface{}{
						types.StringElement{
							Content: "title",
						},
					}
					expected := types.Document{
						Attributes: types.DocumentAttributes{},
						ElementReferences: types.ElementReferences{
							"title": title,
						},
						Footnotes:          types.Footnotes{},
						FootnoteReferences: types.FootnoteReferences{},
						Elements: []interface{}{
							types.Section{
								Level: 0,
								Attributes: types.ElementAttributes{
									types.AttrID:       "title",
									types.AttrCustomID: false,
									types.AttrAuthors: []types.DocumentAuthor{
										{
											FullName: "Chameleon",
											Email:    "",
										},
									},
								},
								Title:    title,
								Elements: []interface{}{},
							},
						},
					}
					Expect(source).To(BecomeDocument(expected))
				})

				It("alternate author input", func() {
					source := `= title
:author: Kismet Rainbow Chameleon` // `:"email":` is processed as a regular attribute
					title := []interface{}{
						types.StringElement{
							Content: "title",
						},
					}
					expected := types.Document{
						Attributes: types.DocumentAttributes{},
						ElementReferences: types.ElementReferences{
							"title": title,
						},
						Footnotes:          types.Footnotes{},
						FootnoteReferences: types.FootnoteReferences{},
						Elements: []interface{}{
							types.Section{
								Level: 0,
								Attributes: types.ElementAttributes{
									types.AttrID:       "title",
									types.AttrCustomID: false,
									types.AttrAuthors: []types.DocumentAuthor{
										{
											FullName: "Kismet Rainbow Chameleon",
											Email:    "",
										},
									},
								},
								Title:    title,
								Elements: []interface{}{},
							},
						},
					}
					Expect(source).To(BecomeDocument(expected))
				})
			})

			Context("multiple authors", func() {

				It("2 authors only", func() {
					source := `= title
Kismet  Rainbow Chameleon  <kismet@asciidoctor.org>; Lazarus het_Draeke <lazarus@asciidoctor.org>`
					title := []interface{}{
						types.StringElement{
							Content: "title",
						},
					}
					expected := types.Document{
						Attributes: types.DocumentAttributes{},
						ElementReferences: types.ElementReferences{
							"title": title,
						},
						Footnotes:          types.Footnotes{},
						FootnoteReferences: types.FootnoteReferences{},
						Elements: []interface{}{
							types.Section{
								Level: 0,
								Attributes: types.ElementAttributes{
									types.AttrID:       "title",
									types.AttrCustomID: false,
									types.AttrAuthors: []types.DocumentAuthor{
										{
											FullName: "Kismet  Rainbow Chameleon  ",
											Email:    "kismet@asciidoctor.org",
										},
										{
											FullName: "Lazarus het_Draeke ",
											Email:    "lazarus@asciidoctor.org",
										},
									},
								},
								Title:    title,
								Elements: []interface{}{},
							},
						},
					}
					Expect(source).To(BecomeDocument(expected))
				})
			})
		})

		Context("document revision", func() {

			It("full document revision", func() {
				source := `= title
				john doe
				v1.0, June 19, 2017: First incarnation`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "June 19, 2017",
									Revremark: "First incarnation",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision with revnumber and revdate only", func() {
				source := `= title
				john doe
				v1.0, June 19, 2017`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "June 19, 2017",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision with revnumber and revdate - with colon separator", func() {
				source := `= title
				john doe
				1.0, June 19, 2017:`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "June 19, 2017",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})
			It("revision with revnumber only - comma suffix", func() {
				source := `= title
				john doe
				1.0,`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision with revdate as number - spaces and no prefix no suffix", func() {
				source := `= title
				john doe
				1.0`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "",
									Revdate:   "1.0",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision with revdate as alphanum - spaces and no prefix no suffix", func() {
				source := `= title
				john doe
				1.0a`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "",
									Revdate:   "1.0a",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision with revnumber only", func() {
				source := `= title
				john doe
				v1.0:`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision with spaces and capital revnumber ", func() {
				source := `= title
				john doe
				V1.0:`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision only - with comma separator", func() {
				source := `= title
				john doe
				v1.0,`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision with revnumber plus comma and colon separators", func() {
				source := `= title
				john doe
				v1.0,:`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("revision with revnumber plus colon separator", func() {
				source := `= title
john doe
v1.0:`
				title := []interface{}{
					types.StringElement{
						Content: "title",
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"title": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "john doe",
										Email:    "",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "",
									Revremark: "",
								},
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

		})

		Context("document header Attributes", func() {

			It("valid attribute names", func() {
				source := `:a:
:author: Xavier
:_author: Xavier
:Author: Xavier
:0Author: Xavier
:Auth0r: Xavier`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "a"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "_author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "Author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "0Author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "Auth0r", Value: "Xavier"},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("attributes and paragraph without blank line in-between", func() {
				source := `:toc:
:date:  2017-01-01
:author: Xavier
:hardbreaks:
a paragraph`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: types.DocumentAttrHardBreaks},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("contiguous attributes and paragraph with blank line in-between", func() {
				source := `:toc:
:date: 2017-01-01
:author: Xavier

a paragraph`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("splitted attributes and paragraph with blank line in-between", func() {
				source := `:toc:
:date: 2017-01-01

:author: Xavier

:hardbreaks:

a paragraph`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "hardbreaks"},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("no header and attributes in body", func() {
				source := `a paragraph
	
:toc:
:date: 2017-01-01
:author: Xavier`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph"},
								},
							},
						},
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})
		})

		Context("document attribute substitutions", func() {

			It("paragraph with attribute substitution", func() {
				source := `:author: Xavier
			
a paragraph written by {author}.`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph written by Xavier."},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("paragraph with attribute resets", func() {
				source := `:author: Xavier
							
:!author1:
:author2!:
a paragraph written by {author}.`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.DocumentAttributeReset{Name: "author1"},
						types.DocumentAttributeReset{Name: "author2"},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph written by Xavier."},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("header with 2 authors, revision and attributes", func() {
				source := `= The Dangerous and Thrilling Documentation Chronicles
Kismet Rainbow Chameleon <kismet@asciidoctor.org>; Lazarus het_Draeke <lazarus@asciidoctor.org>
v1.0, June 19, 2017: First incarnation
:toc:
:keywords: documentation, team, obstacles, journey, victory

This journey begins on a bleary Monday morning.`
				title := []interface{}{
					types.StringElement{Content: "The Dangerous and Thrilling Documentation Chronicles"},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{},
					ElementReferences: types.ElementReferences{
						"the_dangerous_and_thrilling_documentation_chronicles": title,
					},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "the_dangerous_and_thrilling_documentation_chronicles",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "Kismet Rainbow Chameleon ",
										Email:    "kismet@asciidoctor.org",
									},
									{
										FullName: "Lazarus het_Draeke ",
										Email:    "lazarus@asciidoctor.org",
									},
								},
								types.AttrRevision: types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "June 19, 2017",
									Revremark: "First incarnation",
								},
							},
							Title: title,
							Elements: []interface{}{
								types.DocumentAttributeDeclaration{
									Name:  "toc",
									Value: "",
								},
								types.DocumentAttributeDeclaration{
									Name:  "keywords",
									Value: "documentation, team, obstacles, journey, victory",
								},
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "This journey begins on a bleary Monday morning."},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("paragraph with attribute substitution from front-matter", func() {
				source := `---
author: Xavier
---
			
a paragraph written by {author}.`
				expected := types.Document{
					Attributes: types.DocumentAttributes{
						"author": "Xavier",
					},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph written by Xavier."},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})
		})

	})

	Context("invalid document attributes", func() {

		It("paragraph without blank line before attribute declarations", func() {
			source := `a paragraph
:toc:
:date: 2017-01-01
:author: Xavier`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
							{
								types.StringElement{Content: ":toc:"},
							},
							{
								types.StringElement{Content: ":date: 2017-01-01"},
							},
							{
								types.StringElement{Content: ":author: Xavier"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})

		It("invalid attribute names", func() {
			source := `:@date: 2017-01-01
:{author}: Xavier`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: ":@date: 2017-01-01"},
							},
							{
								types.StringElement{Content: ":{author}: Xavier"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})
	})
})
