package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("passthroughs", func() {

	Context("triplePlus Passthrough", func() {

		It("tripleplus passthrough with words", func() {
			actualContent := `+++hello, world+++`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.Passthrough{
							Kind: types.TriplePlusPassthrough,
							Elements: types.InlineElements{
								types.StringElement{
									Content: "hello, world",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("tripleplus empty passthrough ", func() {
			actualContent := `++++++`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.Passthrough{
							Kind:     types.TriplePlusPassthrough,
							Elements: types.InlineElements{},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("tripleplus passthrough with spaces", func() {
			actualContent := `+++ *hello*, world +++`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.Passthrough{
							Kind: types.TriplePlusPassthrough,
							Elements: types.InlineElements{
								types.StringElement{
									Content: " *hello*, world ",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("tripleplus passthrough with only spaces", func() {
			actualContent := `+++ +++`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.Passthrough{
							Kind: types.TriplePlusPassthrough,
							Elements: types.InlineElements{
								types.StringElement{
									Content: " ",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("tripleplus passthrough with line breaks", func() {
			actualContent := "+++\nhello,\nworld\n+++"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.Passthrough{
							Kind: types.TriplePlusPassthrough,
							Elements: types.InlineElements{
								types.StringElement{
									Content: "\nhello,\nworld\n",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("tripleplus passthrough in paragraph", func() {
			actualContent := `The text +++<u>underline & me</u>+++ is underlined.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "The text "},
						types.Passthrough{
							Kind: types.TriplePlusPassthrough,
							Elements: types.InlineElements{
								types.StringElement{
									Content: "<u>underline & me</u>",
								},
							},
						},
						types.StringElement{Content: " is underlined."},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("tripleplus passthrough with embedded image", func() {
			actualContent := `+++image:foo.png[]+++`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.Passthrough{
							Kind: types.TriplePlusPassthrough,
							Elements: types.InlineElements{
								types.StringElement{
									Content: "image:foo.png[]",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

	})

	Context("singleplus passthrough", func() {

		It("singleplus passthrough with words", func() {
			actualContent := `+hello, world+`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.Passthrough{
							Kind: types.SinglePlusPassthrough,
							Elements: types.InlineElements{
								types.StringElement{
									Content: "hello, world",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("singleplus empty passthrough", func() {
			actualContent := `++`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "++",
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("singleplus passthrough with embedded image", func() {
			actualContent := `+image:foo.png[]+`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.Passthrough{
							Kind: types.SinglePlusPassthrough,
							Elements: types.InlineElements{
								types.StringElement{
									Content: "image:foo.png[]",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("invalid singleplus passthrough with spaces - case 1", func() {
			actualContent := `+*hello*, world +` // invalid: space before last `+`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "+",
						},
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("invalid singleplus passthrough with spaces - case 2", func() {
			actualContent := `+ *hello*, world+` // invalid: space after first `+`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "+ ",
						},
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("invalid singleplus passthrough with spaces - case 3", func() {
			actualContent := `+ *hello*, world +` // invalid: spaces within
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "+ ",
						},
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("invalid singleplus passthrough with line break", func() {
			actualContent := "+hello,\nworld+"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "+hello,",
						},
					},
					{
						types.StringElement{
							Content: "world+",
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

	})

	Context("passthrough macro", func() {

		Context("passthrough base macro", func() {

			It("passthrough macro with single word", func() {
				actualContent := `pass:[hello]`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.Passthrough{
								Kind: types.PassthroughMacro,
								Elements: types.InlineElements{
									types.StringElement{
										Content: "hello",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("passthrough macro with words", func() {
				actualContent := `pass:[hello, world]`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.Passthrough{
								Kind: types.PassthroughMacro,
								Elements: types.InlineElements{
									types.StringElement{
										Content: "hello, world",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("empty passthrough macro", func() {
				actualContent := `pass:[]`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.Passthrough{
								Kind:     types.PassthroughMacro,
								Elements: types.InlineElements{},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("passthrough macro with spaces", func() {
				actualContent := `pass:[ *hello*, world ]`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.Passthrough{
								Kind: types.PassthroughMacro,
								Elements: types.InlineElements{
									types.StringElement{
										Content: " *hello*, world ",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("passthrough macro with line break", func() {
				actualContent := "pass:[hello,\nworld]"
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.Passthrough{
								Kind: types.PassthroughMacro,
								Elements: types.InlineElements{
									types.StringElement{
										Content: "hello,\nworld",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})

		Context("passthrough macro with Quoted Text", func() {

			It("passthrough macro with single quoted word", func() {
				actualContent := `pass:q[*hello*]`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.Passthrough{
								Kind: types.PassthroughMacro,
								Elements: types.InlineElements{
									types.QuotedText{
										Kind: types.Bold,
										Elements: types.InlineElements{
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
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("passthrough macro with quoted word in sentence", func() {
				actualContent := `pass:q[ a *hello*, world ]`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.Passthrough{
								Kind: types.PassthroughMacro,
								Elements: types.InlineElements{
									types.StringElement{
										Content: " a ",
									},
									types.QuotedText{
										Kind: types.Bold,
										Elements: types.InlineElements{
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
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})
	})

})
