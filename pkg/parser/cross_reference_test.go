package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golint
)

var _ = Describe("cross references", func() {

	Context("in final documents", func() {

		Context("internal references", func() {

			It("with custom id alone", func() {
				source := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
				title := []interface{}{
					&types.StringElement{
						Content: "a title",
					},
				}
				expected := &types.Document{
					ElementReferences: types.ElementReferences{
						"thetitle": title,
					},
					Elements: []interface{}{
						&types.Section{
							Level: 1,
							Attributes: types.Attributes{
								types.AttrID:       "thetitle",
								types.AttrCustomID: true,
							},
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "with some content linked to ",
										},
										&types.InternalCrossReference{
											ID: "thetitle",
										},
										&types.StringElement{
											Content: "!",
										},
									},
								},
							},
						},
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "thetitle",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with custom id and label", func() {
				source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
				title := []interface{}{
					&types.StringElement{
						Content: "a title",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Level: 1,
							Attributes: types.Attributes{
								types.AttrID:       "thetitle",
								types.AttrCustomID: true,
							},
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "with some content linked to ",
										},
										&types.InternalCrossReference{
											ID:    "thetitle",
											Label: "a label to the title",
										},
										&types.StringElement{
											Content: "!",
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"thetitle": title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "thetitle",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to section defined later in the document", func() {
				source := `a reference to <<section>>
	
[#section]
== A section with a link to https://example.com

some content`
				title := []interface{}{
					&types.StringElement{
						Content: "A section with a link to ",
					},
					&types.InlineLink{
						Location: &types.Location{
							Scheme: "https://",
							Path:   "example.com",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a reference to ",
								},
								&types.InternalCrossReference{
									ID: "section",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID:       "section",
								types.AttrCustomID: true,
							},
							Level: 1,
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "some content",
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"section": title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "section",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to delimited block defined later in the document", func() {
				source := `a reference to <<block>>
	
[#block]
.The block
----
some content
----`

				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a reference to ",
								},
								&types.InternalCrossReference{
									ID: "block",
								},
							},
						},
						&types.DelimitedBlock{
							Kind: types.Listing,
							Attributes: types.Attributes{
								types.AttrID:    "block",
								types.AttrTitle: "The block",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "some content",
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"block": "The block",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to paragraph defined later in the document", func() {
				source := `a reference to <<a-paragraph>>
	
[#a-paragraph]
.another paragraph
some content`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a reference to ",
								},
								&types.InternalCrossReference{
									ID: "a-paragraph",
								},
							},
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrID:    "a-paragraph",
								types.AttrTitle: "another paragraph",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "some content",
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"a-paragraph": "another paragraph",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to table defined later in the document", func() {
				source := `a reference to <<table>>
	
[#table]
.The table
|===
| A | B
|===
`

				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a reference to ",
								},
								&types.InternalCrossReference{
									ID: "table",
								},
							},
						},
						&types.Table{
							Attributes: types.Attributes{
								types.AttrID:    "table",
								types.AttrTitle: "The table",
							},
							Rows: []*types.TableRow{
								{
									Cells: []*types.TableCell{
										{
											Elements: []interface{}{
												&types.StringElement{
													Content: "A ",
												},
											},
										},
										{
											Elements: []interface{}{
												&types.StringElement{
													Content: "B",
												},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"table": "The table",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to attached element in a list", func() {
				source := `a reference to <<table>>
	
. list element
+				
[#table]
.The table
|===
| A | B
|===
`

				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a reference to ",
								},
								&types.InternalCrossReference{
									ID: "table",
								},
							},
						},
						&types.List{
							Kind: types.OrderedListKind,
							Elements: []types.ListElement{
								&types.OrderedListElement{
									Style: types.Arabic,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "list element",
												},
											},
										},
										&types.Table{
											Attributes: types.Attributes{
												types.AttrID:    "table",
												types.AttrTitle: "The table",
											},
											Rows: []*types.TableRow{
												{
													Cells: []*types.TableCell{
														{
															Elements: []interface{}{
																&types.StringElement{
																	Content: "A ",
																},
															},
														},
														{
															Elements: []interface{}{
																&types.StringElement{
																	Content: "B",
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
					ElementReferences: types.ElementReferences{
						"table": "The table",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to term in labeled list", func() {
				source := `[[a_term]]term::
// a comment

Here's a reference to the definition of <<a_term>>.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.LabeledListKind,
							Elements: []types.ListElement{
								&types.LabeledListElement{
									Style: types.DoubleColons,
									Term: []interface{}{
										&types.InlineLink{
											Attributes: types.Attributes{
												types.AttrID: "a_term",
											},
										},
										&types.StringElement{
											Content: "term",
										},
									},
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "Here’s a reference to the definition of ", // note that the quote is transformed
								},
								&types.InternalCrossReference{
									ID: "a_term",
								},
								&types.StringElement{
									Content: ".",
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"a_term": []interface{}{
							&types.StringElement{ // the term content, excluding the inline anchor
								Content: "term",
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("external references", func() {

			It("to other doc with plain text location and rich label", func() {
				source := `some content linked to xref:another-doc.adoc[*another doc*]!`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "some content linked to ",
								},
								&types.ExternalCrossReference{
									Location: &types.Location{
										Path: "another-doc.adoc",
									},
									Attributes: types.Attributes{
										types.AttrXRefLabel: []interface{}{
											&types.QuotedText{
												Kind: types.SingleQuoteBold,
												Elements: []interface{}{
													&types.StringElement{
														Content: "another doc",
													},
												},
											},
										},
									},
								},
								&types.StringElement{
									Content: "!",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to other doc with document attribute in location", func() {
				source := `some content linked to xref:{foo}.adoc[another doc]!`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "some content linked to ",
								},
								&types.ExternalCrossReference{
									Location: &types.Location{
										Path: "{foo}.adoc", // attribute resolution failed
									},
									Attributes: types.Attributes{
										types.AttrXRefLabel: "another doc",
									},
								},
								&types.StringElement{
									Content: "!",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to other doc with document attribute in location and label with special chars", func() {
				source := `
:foo: another-doc.adoc

some content linked to xref:{foo}[another_doc()]!`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "foo",
									Value: "another-doc.adoc",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "some content linked to ",
								},
								&types.ExternalCrossReference{
									Location: &types.Location{
										Path: "another-doc.adoc",
									},
									Attributes: types.Attributes{
										types.AttrXRefLabel: "another_doc()",
									},
								},
								&types.StringElement{
									Content: "!",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
