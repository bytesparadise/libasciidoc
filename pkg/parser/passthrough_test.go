package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("passthroughs - draft", func() {

	Context("triplePlus Passthrough", func() {

		It("tripleplus passthrough with words", func() {
			source := `+++hello, world+++`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus empty passthrough ", func() {
			source := `++++++`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus passthrough with spaces", func() {
			source := `+++ *hello*, world +++`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus passthrough with only spaces", func() {
			source := `+++ +++`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus passthrough with line breaks", func() {
			source := "+++\nhello,\nworld\n+++"
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus passthrough in paragraph", func() {
			source := `The text +++<u>underline & me</u>+++ is underlined.`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus passthrough with embedded image", func() {
			source := `+++image:foo.png[]+++`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

	})

	Context("singleplus passthrough", func() {

		It("singleplus passthrough with words", func() {
			source := `+hello, world+`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("singleplus empty passthrough", func() {
			source := `++`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "++",
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("singleplus passthrough with embedded image", func() {
			source := `+image:foo.png[]+`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("invalid singleplus passthrough with spaces - case 1", func() {
			source := `+*hello*, world +` // invalid: space before last `+`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("invalid singleplus passthrough with spaces - case 2", func() {
			source := `+ *hello*, world+` // invalid: space after first `+`
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("invalid singleplus passthrough with spaces - case 3", func() {
			source := `+ *hello*, world +` // invalid: spaces within
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("invalid singleplus passthrough with line break", func() {
			source := "+hello,\nworld+"
			expected := types.Paragraph{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

	})

	Context("passthrough macro", func() {

		Context("passthrough base macro", func() {

			It("passthrough macro with single word", func() {
				source := `pass:[hello]`
				expected := types.Paragraph{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("passthrough macro with words", func() {
				source := `pass:[hello, world]`
				expected := types.Paragraph{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("empty passthrough macro", func() {
				source := `pass:[]`
				expected := types.Paragraph{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("passthrough macro with spaces", func() {
				source := `pass:[ *hello*, world ]`
				expected := types.Paragraph{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("passthrough macro with line break", func() {
				source := "pass:[hello,\nworld]"
				expected := types.Paragraph{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})
		})

		Context("passthrough macro with Quoted Text", func() {

			It("passthrough macro with single quoted word", func() {
				source := `pass:q[*hello*]`
				expected := types.Paragraph{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("passthrough macro with quoted word in sentence", func() {
				source := `pass:q[ a *hello*, world ]`
				expected := types.Paragraph{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})
		})
	})

})
