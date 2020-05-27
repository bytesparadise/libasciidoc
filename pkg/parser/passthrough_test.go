package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("passthroughs - draft", func() {

	Context("tripleplus inline passthrough", func() {

		It("tripleplus inline passthrough with words", func() {
			source := `+++hello, world+++`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("tripleplus empty passthrough ", func() {
			source := `++++++`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.InlinePassthrough{
							Kind:     types.TriplePlusPassthrough,
							Elements: []interface{}{},
						},
					},
				},
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("tripleplus inline passthrough with spaces", func() {
			source := `+++ *hello*, world +++`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("tripleplus inline passthrough with only spaces", func() {
			source := `+++ +++`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("tripleplus inline passthrough with line breaks", func() {
			source := "+++\nhello,\nworld\n+++"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("tripleplus inline passthrough in paragraph", func() {
			source := `The text +++<u>underline & me</u>+++ is underlined.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("tripleplus inline passthrough with embedded image", func() {
			source := `+++image:foo.png[]+++`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

	})

	Context("singleplus passthrough", func() {

		It("singleplus passthrough with words", func() {
			source := `+hello, world+`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("singleplus empty passthrough", func() {
			source := `++`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "++",
						},
					},
				},
			}
			result, err := ParseDocumentBlock(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(expected))
		})

		It("singleplus passthrough with embedded image", func() {
			source := `+image:foo.png[]+`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("invalid singleplus passthrough with spaces - case 1", func() {
			source := `+*hello*, world +` // invalid: space before last `+`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("invalid singleplus passthrough with spaces - case 2", func() {
			source := `+ *hello*, world+` // invalid: space after first `+`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("invalid singleplus passthrough with spaces - case 3", func() {
			source := `+ *hello*, world +` // invalid: spaces within
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
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
			}
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

		It("invalid singleplus passthrough with line break", func() {
			source := "+hello,\nworld+"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
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
			Expect(ParseDocumentBlock(source)).To(Equal(expected))
		})

	})

	Context("passthrough macro", func() {

		Context("passthrough base macro", func() {

			It("passthrough macro with single word", func() {
				source := `pass:[hello]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.InlinePassthrough{
								Kind: types.PassthroughMacro,
								Elements: []interface{}{
									types.StringElement{
										Content: "hello",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("passthrough macro with words", func() {
				source := `pass:[hello, world]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.InlinePassthrough{
								Kind: types.PassthroughMacro,
								Elements: []interface{}{
									types.StringElement{
										Content: "hello, world",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("empty passthrough macro", func() {
				source := `pass:[]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.InlinePassthrough{
								Kind:     types.PassthroughMacro,
								Elements: []interface{}{},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("passthrough macro with spaces", func() {
				source := `pass:[ *hello*, world ]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.InlinePassthrough{
								Kind: types.PassthroughMacro,
								Elements: []interface{}{
									types.StringElement{
										Content: " *hello*, world ",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("passthrough macro with line break", func() {
				source := "pass:[hello,\nworld]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.InlinePassthrough{
								Kind: types.PassthroughMacro,
								Elements: []interface{}{
									types.StringElement{
										Content: "hello,\nworld",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})
		})

		Context("passthrough macro with Quoted Text", func() {

			It("passthrough macro with single quoted word", func() {
				source := `pass:q[*hello*]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.InlinePassthrough{
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
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("passthrough macro with quoted word in sentence", func() {
				source := `pass:q[ a *hello*, world ]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.InlinePassthrough{
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
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})
		})
	})

})
