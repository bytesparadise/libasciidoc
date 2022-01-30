package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("apply substitutions", func() {

	var c chan types.DocumentFragment
	done := make(<-chan interface{})
	var ctx *parser.ParseContext

	BeforeEach(func() {
		c = make(chan types.DocumentFragment, 1)
		ctx = parser.NewParseContext(configuration.NewConfiguration(
			configuration.WithAttributes(map[string]interface{}{
				"role1":   "role_1",
				"role2":   "my_role_2",
				"title":   "Title",
				"option1": "option_1",
				"option2": "option_2",
				"cookie":  "yummy",
			}),
		))
	})

	AfterEach(func() {
		close(c)
	})

	It("should process substitutions in document header", func() {
		// given
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.DocumentHeader{
					Attributes: types.Attributes{
						types.AttrTitle: []interface{}{
							&types.AttributeReference{
								Name: "title",
							},
						},
					},
					Title: []interface{}{
						types.RawLine("Document [.{role1}]*{title}*"),
					},
					Elements: []interface{}{
						&types.DocumentAuthors{ // TODO: support attribute references in document authors
							{
								DocumentAuthorFullName: &types.DocumentAuthorFullName{
									FirstName: "Xavier",
								},
								Email: "xavier@example.com",
							},
						},
						&types.DocumentRevision{ // TODO: support attribute references in document revision
							Revnumber: "1.0",
						},
					},
				},
			},
		}
		// when
		result := parser.ApplySubstitutions(ctx, done, c)
		// then
		Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
			Elements: []interface{}{
				&types.DocumentHeader{
					Attributes: types.Attributes{
						types.AttrTitle: "Title",
					},
					Title: []interface{}{
						&types.StringElement{
							Content: "Document ",
						},
						&types.QuotedText{
							Kind: types.SingleQuoteBold,
							Attributes: types.Attributes{
								types.AttrRoles: types.Roles{"role_1"},
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "Title",
								},
							},
						},
					},
					Elements: []interface{}{
						&types.DocumentAuthors{ // TODO: support attribute references in document authors
							{
								DocumentAuthorFullName: &types.DocumentAuthorFullName{
									FirstName: "Xavier",
								},
								Email: "xavier@example.com",
							},
						},
						&types.DocumentRevision{ // TODO: support attribute references in document revision
							Revnumber: "1.0",
						},
					},
				},
			},
		}))
	})

	It("should process substitutions in section title", func() {
		// given
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.Section{
					Level: 1,
					Attributes: types.Attributes{
						types.AttrTitle: []interface{}{
							&types.AttributeReference{
								Name: "title",
							},
						},
					},
					Title: []interface{}{
						types.RawLine("Section [.{role1}]*{title}*"),
					},
				},
			},
		}
		// when
		result := parser.ApplySubstitutions(ctx, done, c)
		// then
		Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
			Elements: []interface{}{
				&types.Section{
					Level: 1,
					Attributes: types.Attributes{
						types.AttrTitle: "Title",
					},
					Title: []interface{}{
						&types.StringElement{
							Content: "Section ",
						},
						&types.QuotedText{
							Kind: types.SingleQuoteBold,
							Attributes: types.Attributes{
								types.AttrRoles: types.Roles{"role_1"},
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "Title",
								},
							},
						},
					},
				},
			},
		}))
	})

	It("should process substitutions in paragraph lines", func() {
		// given
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.Paragraph{
					Attributes: types.Attributes{
						types.AttrTitle: "The Title",
					},
					Elements: []interface{}{
						types.RawLine("a paragraph called\n"),
						types.RawLine("the {title}."),
					},
				},
			},
		}
		// when
		result := parser.ApplySubstitutions(ctx, done, c)
		// then
		Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
			Elements: []interface{}{
				&types.Paragraph{
					Attributes: types.Attributes{
						types.AttrTitle: "The Title",
					},
					Elements: []interface{}{
						&types.StringElement{
							Content: "a paragraph called\nthe Title.",
						},
					},
				},
			},
		}))
	})

	It("should process substitutions in paragraph attributes", func() {
		// given
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.Paragraph{
					Attributes: types.Attributes{
						types.AttrTitle: []interface{}{
							&types.StringElement{
								Content: "The ",
							},
							&types.AttributeReference{
								Name: "title",
							},
						},
						types.AttrRoles: types.Roles{
							[]interface{}{
								&types.StringElement{
									Content: "my_",
								},
								&types.AttributeReference{
									Name: "role1",
								},
							},
							[]interface{}{
								&types.AttributeReference{
									Name: "role2",
								},
							},
						},
						types.AttrOptions: types.Options{
							[]interface{}{
								&types.StringElement{
									Content: "an_",
								},
								&types.AttributeReference{
									Name: "option1",
								},
							},
							[]interface{}{
								&types.AttributeReference{
									Name: "option2",
								},
							},
						},
					},
					Elements: []interface{}{
						types.RawLine("a line"),
					},
				},
			},
		}
		// when
		result := parser.ApplySubstitutions(ctx, done, c)
		// then
		Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
			Elements: []interface{}{
				&types.Paragraph{
					Attributes: types.Attributes{
						types.AttrTitle: "The Title",
						types.AttrRoles: types.Roles{
							"my_role_1",
							"my_role_2",
						},
						types.AttrOptions: types.Options{
							"an_option_1",
							"option_2",
						},
					},
					Elements: []interface{}{
						&types.StringElement{
							Content: "a line",
						},
					},
				},
			},
		}))
	})

	It("should process substitutions in quoted text", func() {
		// given
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.Paragraph{
					Elements: []interface{}{
						types.RawLine("quoted text [.{role1}]*here*."),
					},
				},
			},
		}
		// when
		result := parser.ApplySubstitutions(ctx, done, c)
		// then
		Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
			Elements: []interface{}{
				&types.Paragraph{
					Elements: []interface{}{
						&types.StringElement{
							Content: "quoted text ",
						},
						&types.QuotedText{
							Kind: types.SingleQuoteBold,
							Attributes: types.Attributes{
								types.AttrRoles: types.Roles{
									"role_1",
								},
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "here",
								},
							},
						},
						&types.StringElement{
							Content: ".",
						},
					},
				},
			},
		}))
	})

	It("should process substitutions in inline image", func() {
		// given
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.Paragraph{
					Elements: []interface{}{
						types.RawLine("image:{cookie}.png[[.{role1}]_yummy!_]"),
					},
				},
			},
		}
		// when
		result := parser.ApplySubstitutions(ctx, done, c)
		// then
		Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
			Elements: []interface{}{
				&types.Paragraph{
					Elements: []interface{}{
						&types.InlineImage{
							Attributes: types.Attributes{
								types.AttrImageAlt: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteItalic,
										Attributes: types.Attributes{
											types.AttrRoles: types.Roles{"role_1"},
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "yummy!",
											},
										},
									},
								},
							},
							Location: &types.Location{
								Path: "yummy.png",
							},
						},
					},
				},
			},
		}))
	})

	It("should process substitutions in block image", func() {
		// given
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.ImageBlock{
					Attributes: types.Attributes{
						types.AttrRoles: types.Roles{
							[]interface{}{
								&types.StringElement{
									Content: "my_",
								},
								&types.AttributeReference{
									Name: "role1",
								},
							},
						},
					},
					Location: &types.Location{ // TODO: support substitutions in scheme?
						Path: []interface{}{
							&types.StringElement{
								Content: "path/to/",
							},
							&types.AttributeReference{
								Name: "cookie",
							},
							&types.StringElement{
								Content: ".png",
							},
						},
					},
				},
			},
		}
		// when
		result := parser.ApplySubstitutions(ctx, done, c)
		// then
		Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
			Elements: []interface{}{
				&types.ImageBlock{
					Attributes: types.Attributes{
						types.AttrRoles: types.Roles{"my_role_1"},
					},
					Location: &types.Location{ // TODO: support substitutions in scheme?
						Path: "path/to/yummy.png",
					},
				},
			},
		}))
	})

	It("should process substitutions in table block", func() {
		// given
		c <- types.DocumentFragment{
			Elements: []interface{}{
				&types.Table{
					Attributes: types.Attributes{
						types.AttrTitle: []interface{}{
							&types.AttributeReference{
								Name: "title",
							},
						},
					},
					Header: &types.TableRow{
						Cells: []*types.TableCell{
							{
								Elements: []interface{}{
									types.RawLine("[.{role1}]_yummy header!_"),
								},
							},
						},
					},
					Rows: []*types.TableRow{
						{
							Cells: []*types.TableCell{
								{
									Elements: []interface{}{
										types.RawLine("image:{cookie}.png[[.{role1}]_yummy row!_]"),
									},
								},
							},
						},
					},
					Footer: &types.TableRow{
						Cells: []*types.TableCell{
							{
								Elements: []interface{}{
									types.RawLine("[.{role1}]_yummy footer!_"),
								},
							},
						},
					},
				},
			},
		}
		// when
		result := parser.ApplySubstitutions(ctx, done, c)
		// then
		Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
			Elements: []interface{}{
				&types.Table{
					Attributes: types.Attributes{
						types.AttrTitle: "Title",
					},
					Header: &types.TableRow{
						Cells: []*types.TableCell{
							{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteItalic,
										Attributes: types.Attributes{
											types.AttrRoles: types.Roles{"role_1"},
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "yummy header!",
											},
										},
									},
								},
							},
						},
					},
					Rows: []*types.TableRow{
						{
							Cells: []*types.TableCell{
								{
									Elements: []interface{}{
										&types.InlineImage{
											Attributes: types.Attributes{
												types.AttrImageAlt: []interface{}{
													&types.QuotedText{
														Kind: types.SingleQuoteItalic,
														Attributes: types.Attributes{
															types.AttrRoles: types.Roles{"role_1"},
														},
														Elements: []interface{}{
															&types.StringElement{
																Content: "yummy row!",
															},
														},
													},
												},
											},
											Location: &types.Location{
												Path: "yummy.png",
											},
										},
									},
								},
							},
						},
					},
					Footer: &types.TableRow{
						Cells: []*types.TableCell{
							{
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteItalic,
										Attributes: types.Attributes{
											types.AttrRoles: types.Roles{"role_1"},
										},
										Elements: []interface{}{
											&types.StringElement{
												Content: "yummy footer!",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}))
	})

	// It("should process substitutions in listing block", func() {
	// 	Fail("not implemented yet")
	// })

	// It("should process substitutions in example block", func() {
	// 	Fail("not implemented yet")
	// })

})
