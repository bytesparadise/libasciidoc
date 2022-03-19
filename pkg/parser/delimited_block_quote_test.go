package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("quote blocks", func() {

	Context("in final documents", func() {

		Context("as delimited blocks", func() {

			It("with single-line content and author and title attributes", func() {
				source := `[quote, john doe, quote title]
____
some *quote* content
____`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Attributes: types.Attributes{
								types.AttrStyle:       types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "some ",
										},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{
													Content: "quote",
												},
											},
										},
										&types.StringElement{
											Content: " content",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multi-line content and author attribute", func() {
				source := `[quote, john doe,   ]
____
- some 
- quote 
- content 
____
`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Attributes: types.Attributes{
								types.AttrStyle:       types.Quote,
								types.AttrQuoteAuthor: "john doe",
							},
							Elements: []interface{}{
								&types.List{
									Kind: types.UnorderedListKind,
									Elements: []types.ListElement{
										// suffix spaces are trimmed on each line
										&types.UnorderedListElement{
											BulletStyle: types.Dash,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												&types.Paragraph{
													Elements: []interface{}{
														&types.StringElement{
															Content: "some",
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
															Content: "quote",
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
															Content: "content",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with single-line content with title attribute", func() {
				source := `[quote, ,quote title]
____
some quote content 
____
`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Attributes: types.Attributes{
								types.AttrStyle:      types.Quote,
								types.AttrQuoteTitle: "quote title",
							},
							Elements: []interface{}{
								// suffix spaces are trimmed on each line
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "some quote content",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with single line starting with a dot", func() {
				source := `[quote]
____
.standalone
____`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
						},
					},
				}
				result, err := ParseDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDocument(expected))
			})

			It("with last line starting with a dot", func() {
				source := `
----
some content

.standalone
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "some content\n\n.standalone", // includes all lines
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line quote with rendered lists and block and without author and title", func() {
				source := `[quote]
____
* some
----
* listing 
----
* content
____`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{
								// suffix spaces are trimmed on each line
								&types.List{
									Kind: types.UnorderedListKind,
									Elements: []types.ListElement{
										&types.UnorderedListElement{
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												&types.Paragraph{
													Elements: []interface{}{
														&types.StringElement{
															Content: "some",
														},
													},
												},
											},
										},
									},
								},
								&types.DelimitedBlock{
									Kind: types.Listing,
									Elements: []interface{}{
										&types.StringElement{
											Content: "* listing",
										},
									},
								},
								&types.List{
									Kind: types.UnorderedListKind,
									Elements: []types.ListElement{
										&types.UnorderedListElement{
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												&types.Paragraph{
													Elements: []interface{}{
														&types.StringElement{
															Content: "content",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line quote with rendered list and without author and title", func() {
				source := `[quote]
____
* some


* quote 


* content
____`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{
								&types.List{
									Kind: types.UnorderedListKind,
									Elements: []types.ListElement{
										// suffix spaces are trimmed on each line
										&types.UnorderedListElement{
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												&types.Paragraph{
													Elements: []interface{}{
														&types.StringElement{
															Content: "some",
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
															Content: "quote",
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
															Content: "content",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("empty quote without author and title", func() {
				source := `[quote]
____
____`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unclosed quote without author and title", func() {
				source := `[quote]
____
foo
`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "foo",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with variable delimiter length", func() {

				It("with 5 chars", func() {
					source := `[quote]
_____
some *quote* content
_____`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DelimitedBlock{
								Kind: types.Quote,
								Attributes: types.Attributes{
									types.AttrStyle: types.Quote,
								},
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "some ",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteBold,
												Elements: []interface{}{
													&types.StringElement{
														Content: "quote",
													},
												},
											},
											&types.StringElement{
												Content: " content",
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with 5 chars with nested with 4 chars", func() {
					source := `[quote]
_____
[quote]
____
some *quote* content
____
_____`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DelimitedBlock{
								Kind: types.Quote,
								Attributes: types.Attributes{
									types.AttrStyle: types.Quote,
								},
								Elements: []interface{}{
									&types.DelimitedBlock{
										Kind: types.Quote,
										Attributes: types.Attributes{
											types.AttrStyle: types.Quote,
										},
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "some ",
													},
													&types.QuotedText{
														Kind: types.SingleQuoteBold,
														Elements: []interface{}{
															&types.StringElement{
																Content: "quote",
															},
														},
													},
													&types.StringElement{
														Content: " content",
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
	})
})
