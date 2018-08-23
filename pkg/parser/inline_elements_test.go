package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("inline elements", func() {

	It("bold text without paranthesis", func() {
		actualContent := "*some bold content*"
		expectedResult := types.InlineElements{
			types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some bold content"},
				},
			},
		}

		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
	})

	It("bold text within paranthesis", func() {
		actualContent := "(*some bold content*)"
		expectedResult := types.InlineElements{
			types.StringElement{Content: "("},
			types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some bold content"},
				},
			},
			types.StringElement{Content: ")"},
		}

		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
	})

	It("bold text within words", func() {
		actualContent := "some*bold*content"
		expectedResult := types.InlineElements{
			types.StringElement{Content: "some*bold*content"},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
	})

	It("invalid bold portion of text", func() {
		actualContent := "*foo*bar"
		expectedResult := types.InlineElements{
			types.StringElement{Content: "*foo*bar"},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
	})

	It("valid bold portion of text", func() {
		actualContent := "**foo**bar"
		expectedResult := types.InlineElements{
			types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "foo"},
				},
			},
			types.StringElement{Content: "bar"},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
	})
})
