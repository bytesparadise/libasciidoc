package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("passthroughs", func() {

	Context("triplePlus Passthrough", func() {

		It("tripleplus passthrough with words", func() {
			actualContent := `+++hello, world+++`
			expectedResult := types.Passthrough{
				Kind: types.TriplePlusPassthrough,
				Elements: []interface{}{
					types.StringElement{
						Content: "hello, world",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
		})

		It("tripleplus empty passthrough ", func() {
			actualContent := `++++++`
			expectedResult := types.Passthrough{
				Kind:     types.TriplePlusPassthrough,
				Elements: []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
		})

		It("tripleplus passthrough with spaces", func() {
			actualContent := `+++ *hello*, world +++`
			expectedResult := types.Passthrough{
				Kind: types.TriplePlusPassthrough,
				Elements: []interface{}{
					types.StringElement{
						Content: " *hello*, world ",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
		})

		It("tripleplus passthrough with only spaces", func() {
			actualContent := `+++ +++`
			expectedResult := types.Passthrough{
				Kind: types.TriplePlusPassthrough,
				Elements: []interface{}{
					types.StringElement{
						Content: " ",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
		})

		It("tripleplus passthrough with line break", func() {
			actualContent := "+++hello,\nworld+++"
			expectedResult := types.Passthrough{
				Kind: types.TriplePlusPassthrough,
				Elements: []interface{}{
					types.StringElement{
						Content: "hello,\nworld",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
		})
	})

	Context("singlePlus Passthrough", func() {

		It("singleplus passthrough with words", func() {
			actualContent := `+hello, world+`
			expectedResult := types.Passthrough{
				Kind: types.SinglePlusPassthrough,
				Elements: []interface{}{
					types.StringElement{
						Content: "hello, world",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
		})

		It("singleplus empty passthrough", func() {
			actualContent := `++`
			expectedResult := types.Passthrough{
				Kind:     types.SinglePlusPassthrough,
				Elements: []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
		})

		It("singleplus passthrough with spaces", func() {
			actualContent := `+ *hello*, world +`
			expectedResult := types.Passthrough{
				Kind: types.SinglePlusPassthrough,
				Elements: []interface{}{
					types.StringElement{
						Content: " *hello*, world ",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
		})

		It("singleplus passthrough with line break", func() {
			actualContent := "+hello,\nworld+"
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})
	})

	Context("passthrough Macro", func() {

		Context("passthrough Base Macro", func() {

			It("passthrough macro with single word", func() {
				actualContent := `pass:[hello]`
				expectedResult := types.Passthrough{
					Kind: types.PassthroughMacro,
					Elements: []interface{}{
						types.StringElement{
							Content: "hello",
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
			})

			It("passthrough macro with words", func() {
				actualContent := `pass:[hello, world]`
				expectedResult := types.Passthrough{
					Kind: types.PassthroughMacro,
					Elements: []interface{}{
						types.StringElement{
							Content: "hello, world",
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
			})

			It("empty passthrough macro", func() {
				actualContent := `pass:[]`
				expectedResult := types.Passthrough{
					Kind:     types.PassthroughMacro,
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
			})

			It("passthrough macro with spaces", func() {
				actualContent := `pass:[ *hello*, world ]`
				expectedResult := types.Passthrough{
					Kind: types.PassthroughMacro,
					Elements: []interface{}{
						types.StringElement{
							Content: " *hello*, world ",
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
			})

			It("passthrough macro with line break", func() {
				actualContent := "pass:[hello,\nworld]"
				expectedResult := types.Passthrough{
					Kind: types.PassthroughMacro,
					Elements: []interface{}{
						types.StringElement{
							Content: "hello,\nworld",
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
			})
		})

		Context("passthrough Macro with Quoted Text", func() {

			It("passthrough macro with single quoted word", func() {
				actualContent := `pass:q[*hello*]`
				expectedResult := types.Passthrough{
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
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
			})

			It("passthrough macro with quoted word in sentence", func() {
				actualContent := `pass:q[ a *hello*, world ]`
				expectedResult := types.Passthrough{
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
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Passthrough"))
			})
		})
	})

})
