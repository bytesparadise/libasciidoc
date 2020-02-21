package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document table of contents assertions", func() {

	preamble := types.Preamble{
		Elements: []interface{}{
			types.BlankLine{},
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "A short preamble"},
					},
				},
			},
			types.BlankLine{},
		},
	}
	section := types.Section{
		Level:      1,
		Attributes: types.ElementAttributes{},
		Title: []interface{}{
			types.StringElement{Content: "section 1"},
		},
		Elements: []interface{}{},
	}
	tableOfContents := types.TableOfContentsPlaceHolder{}

	expected := types.Document{
		Attributes: types.DocumentAttributes{
			types.AttrTableOfContents: "",
		},
		ElementReferences:  types.ElementReferences{},
		Footnotes:          types.Footnotes{},
		FootnoteReferences: types.FootnoteReferences{},
		Elements: []interface{}{
			tableOfContents,
			preamble,
			section,
		},
	}

	It("should match", func() {
		// given
		actual := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				preamble,
				section,
			},
		}
		// when
		result := testsupport.IncludeTableOfContentsPlaceHolder(actual)
		// then
		Expect(result).To(Equal(expected))
	})

	It("should not match", func() {
		// given
		actual := types.Document{}
		// when
		result := testsupport.IncludeTableOfContentsPlaceHolder(actual)
		// then
		Expect(result).NotTo(Equal(expected))
	})

})
