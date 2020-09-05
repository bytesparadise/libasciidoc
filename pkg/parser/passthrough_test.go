package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("passthroughs", func() {

	Context("final document", func() {

		Context("tripleplus inline passthrough", func() {

			It("tripleplus inline passthrough with words", func() {
				source := `+++hello, world+++`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind: types.TriplePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: "hello, world",
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

			It("tripleplus empty passthrough ", func() {
				source := `++++++`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind:     types.TriplePlusPassthrough,
										Elements: []interface{}{},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("tripleplus inline passthrough with spaces and nested attribute substitution", func() {
				source := `:hello: HELLO
				
+++ {hello}, world +++` // attribute susbsitution must not occur
				expected := types.Document{
					Attributes: types.Attributes{
						"hello": "HELLO",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind: types.TriplePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: " {hello}, world ",
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

			It("tripleplus inline passthrough with spaces aned nested quoted text", func() {
				source := `+++ *hello*, world +++` // macro susbsitution must not occur
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind: types.TriplePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: " *hello*, world ",
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

			It("tripleplus inline passthrough with only spaces", func() {
				source := `+++ +++`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind: types.TriplePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: " ",
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

			It("tripleplus inline passthrough with line breaks", func() {
				source := "+++\nhello,\nworld\n+++"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind: types.TriplePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: "\nhello,\nworld\n",
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

			It("tripleplus inline passthrough in paragraph", func() {
				source := `The text +++<u>underline & me</u>+++ is underlined.`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "The text "},
									types.InlinePassthrough{
										Kind: types.TriplePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: "<u>underline & me</u>",
											},
										},
									},
									types.StringElement{Content: " is underlined."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("tripleplus inline passthrough with embedded image", func() {
				source := `+++image:foo.png[]+++`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind: types.TriplePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: "image:foo.png[]",
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

		Context("singleplus passthrough", func() {

			It("singleplus passthrough with words", func() {
				source := `+hello, world+`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind: types.SinglePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: "hello, world",
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

			It("singleplus empty passthrough", func() {
				source := `++`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "++",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("singleplus passthrough with embedded image", func() {
				source := `+image:foo.png[]+`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlinePassthrough{
										Kind: types.SinglePlusPassthrough,
										Elements: []interface{}{
											types.StringElement{
												Content: "image:foo.png[]",
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

			It("invalid singleplus passthrough with spaces - case 1", func() {
				source := `+*hello*, world +` // invalid: space before last `+`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "+",
									},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "hello",
											},
										},
									},
									types.StringElement{
										Content: ", world",
									},
									types.LineBreak{},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid singleplus passthrough with spaces - case 2", func() {
				source := `+ *hello*, world+` // invalid: space after first `+`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "+ ",
									},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "hello",
											},
										},
									},
									types.StringElement{
										Content: ", world+",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid singleplus passthrough with spaces - case 3", func() {
				source := `+ *hello*, world +` // invalid: spaces within
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "+ ",
									},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "hello",
											},
										},
									},
									types.StringElement{
										Content: ", world",
									},
									types.LineBreak{},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid singleplus passthrough with line break", func() {
				source := "+hello,\nworld+"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "+hello,",
									},
								},
								[]interface{}{
									types.StringElement{
										Content: "world+",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("invalid cases", func() {
				It("invalid singleplus passthrough in paragraph", func() {
					source := `The text + *hello*, world + is not passed through.`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{
											Content: "The text + ",
										},
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "hello",
												},
											},
										},
										types.StringElement{
											Content: ", world + is not passed through.",
										},
									},
								},
							},
						},
					}
					result, err := ParseDocument(source)
					Expect(err).NotTo(HaveOccurred())
					Expect(result).To(MatchDocument(expected))
				})
			})

		})

		Context("passthrough macro", func() {

			Context("passthrough base macro", func() {

				It("passthrough macro with single word", func() {
					source := `pass:[hello]`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: []interface{}{
									[]interface{}{types.InlinePassthrough{
										Kind: types.PassthroughMacro,
										Elements: []interface{}{
											types.StringElement{
												Content: "hello",
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

				It("passthrough macro with words", func() {
					source := `pass:[hello, world]`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: []interface{}{
									[]interface{}{types.InlinePassthrough{
										Kind: types.PassthroughMacro,
										Elements: []interface{}{
											types.StringElement{
												Content: "hello, world",
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

				It("empty passthrough macro", func() {
					source := `pass:[]`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: []interface{}{
									[]interface{}{types.InlinePassthrough{
										Kind:     types.PassthroughMacro,
										Elements: []interface{}{},
									},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("passthrough macro with spaces", func() {
					source := `pass:[ *hello*, world ]`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: []interface{}{
									[]interface{}{types.InlinePassthrough{
										Kind: types.PassthroughMacro,
										Elements: []interface{}{
											types.StringElement{
												Content: " *hello*, world ",
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

				It("passthrough macro with line break", func() {
					source := "pass:[hello,\nworld]"
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: []interface{}{
									[]interface{}{types.InlinePassthrough{
										Kind: types.PassthroughMacro,
										Elements: []interface{}{
											types.StringElement{
												Content: "hello,\nworld",
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

			Context("passthrough macro with Quoted Text", func() {

				It("passthrough macro with single quoted word", func() {
					source := `pass:q[*hello*]`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: []interface{}{
									[]interface{}{types.InlinePassthrough{
										Kind: types.PassthroughMacro,
										Elements: []interface{}{
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "hello",
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

				It("passthrough macro with quoted word in sentence", func() {
					source := `pass:q[ a *hello*, world ]`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: []interface{}{
									[]interface{}{types.InlinePassthrough{
										Kind: types.PassthroughMacro,
										Elements: []interface{}{
											types.StringElement{
												Content: " a ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "hello",
													},
												},
											},
											types.StringElement{
												Content: ", world ",
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
