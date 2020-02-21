package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("inline elements", func() {

	It("bold text without parenthesis", func() {
		source := "*some bold content*"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: [][]interface{}{
				{
					types.QuotedText{
						Kind: types.Bold,
						Elements: []interface{}{
							types.StringElement{Content: "some bold content"},
						},
					},
				},
			},
		}
		Expect(ParseDocumentBlock(source)).To(Equal(expected))
	})

	It("bold text within parenthesis", func() {
		source := "(*some bold content*)"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: [][]interface{}{
				{
					types.StringElement{Content: "("},
					types.QuotedText{
						Kind: types.Bold,
						Elements: []interface{}{
							types.StringElement{Content: "some bold content"},
						},
					},
					types.StringElement{Content: ")"},
				},
			},
		}
		Expect(ParseDocumentBlock(source)).To(Equal(expected))
	})

	It("non-bold text within words", func() {
		source := "some*bold*content"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: [][]interface{}{
				{
					types.StringElement{Content: "some*bold*content"},
				},
			},
		}
		Expect(ParseDocumentBlock(source)).To(Equal(expected))
	})

	It("non-italic text within words", func() {
		source := "some_italic_content"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: [][]interface{}{
				{
					types.StringElement{Content: "some_italic_content"},
				},
			},
		}
		Expect(ParseDocumentBlock(source)).To(Equal(expected))
	})
	It("non-monospace text within words", func() {
		source := "some`monospace`content"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: [][]interface{}{
				{
					types.StringElement{Content: "some`monospace`content"},
				},
			},
		}
		Expect(ParseDocumentBlock(source)).To(Equal(expected))
	})

	It("invalid bold portion of text", func() {
		source := "*foo*bar"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: [][]interface{}{
				{
					types.StringElement{Content: "*foo*bar"},
				},
			},
		}
		Expect(ParseDocumentBlock(source)).To(Equal(expected))
	})

	It("valid bold portion of text", func() {
		source := "**foo**bar"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: [][]interface{}{
				{
					types.QuotedText{
						Kind: types.Bold,
						Elements: []interface{}{
							types.StringElement{Content: "foo"},
						},
					},
					types.StringElement{Content: "bar"},
				},
			},
		}
		Expect(ParseDocumentBlock(source)).To(Equal(expected))
	})

	It("latin characters", func() {
		source := "à bientôt"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: [][]interface{}{
				{
					types.StringElement{Content: "à bientôt"},
				},
			},
		}
		Expect(ParseDocumentBlock(source)).To(Equal(expected))
	})
})
