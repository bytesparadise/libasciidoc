package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("quote blocks", func() {

	Context("draft documents", func() {

		Context("delimited blocks", func() {

			It("single-line quote block with author and title", func() {
				source := `[quote, john doe, quote title]
____
some *quote* content
____`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "quote",
													},
												},
											},
											types.StringElement{
												Content: " content",
											},
										},
									},
								},
							},
						},
					},
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
			})

			It("multi-line quote with author only", func() {
				source := `[quote, john doe,   ]
____
- some 
- quote 
- content 
____
`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Quote,
								types.AttrQuoteAuthor: "john doe",
							},
							Elements: []interface{}{
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "some ",
													},
												},
											},
										},
									},
								},
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "quote ",
													},
												},
											},
										},
									},
								},
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "content ",
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-line quote with title only", func() {
				source := `[quote, ,quote title]
____
some quote content 
____
`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle:      types.Quote,
								types.AttrQuoteTitle: "quote title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "some quote content ",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line quote with rendered list and without author and title", func() {
				source := `[quote]
____
* some


* quote 


* content
____`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "some",
													},
												},
											},
										},
									},
								},
								types.BlankLine{},
								types.BlankLine{},
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "quote ",
													},
												},
											},
										},
									},
								},
								types.BlankLine{},
								types.BlankLine{},
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
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
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("empty quote without author and title", func() {
				source := `[quote]
____
____`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("unclosed quote without author and title", func() {
				source := `[quote]
____
foo
`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
	})

	Context("final documents", func() {

		Context("delimited blocks", func() {

			It("single-line quote block with author and title", func() {
				source := `[quote, john doe, quote title]
____
some *quote* content
____`
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "quote",
													},
												},
											},
											types.StringElement{
												Content: " content",
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

			It("multi-line quote with author only", func() {
				source := `[quote, john doe,   ]
____
- some 
- quote 
- content 
____
`
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Quote,
								types.AttrQuoteAuthor: "john doe",
							},
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.Dash,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
														{
															types.StringElement{
																Content: "some ",
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
																Content: "quote ",
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
																Content: "content ",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-line quote with title only", func() {
				source := `[quote, ,quote title]
____
some quote content 
____
`
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle:      types.Quote,
								types.AttrQuoteTitle: "quote title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "some quote content ",
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

			It("with single line starting with a dot", func() {
				source := `[quote]
____
.foo
____`
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{},
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
* quote 
----
* content
____`
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
														{
															types.StringElement{
																Content: "some",
															},
														},
													},
												},
											},
										},
									},
								},
								types.ListingBlock{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "* quote ",
											},
										},
									},
								},
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
														{
															types.StringElement{
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
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
														{
															types.StringElement{
																Content: "some",
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
																Content: "quote ",
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
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("empty quote without author and title", func() {
				source := `[quote]
____
____`
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{},
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
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrStyle: types.Quote,
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "foo",
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
