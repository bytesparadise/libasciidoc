package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("table of contents", func() {

	// reusable elements
	doctitleAttribute := types.SectionTitle{
		Attributes: types.ElementAttributes{
			types.AttrID: "_a_title",
		},
		Elements: types.InlineElements{
			types.StringElement{Content: "A Title"},
		},
	}
	tableOfContents := types.TableOfContentsMacro{}
	preamble := types.Preamble{
		Elements: []interface{}{
			types.BlankLine{},
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "A short preamble"},
					},
				},
			},
			types.BlankLine{},
		},
	}
	section := types.Section{
		Level: 1,
		Title: types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID: "_section_1",
			},
			Elements: types.InlineElements{
				types.StringElement{Content: "section 1"},
			},
		},
		Elements: []interface{}{},
	}

	It("toc with default placement", func() {

		actualContent := `= A Title
:toc:

A short preamble

== section 1`

		expectedResult := types.Document{
			Attributes: types.DocumentAttributes{
				"doctitle": doctitleAttribute,
				"toc":      "",
			},
			ElementReferences: types.ElementReferences{
				"_section_1": types.SectionTitle{
					Attributes: types.ElementAttributes{
						types.AttrID: "_section_1",
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "section 1"},
					},
				},
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				tableOfContents,
				preamble,
				section,
			},
		}
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("toc with preamble placement", func() {
		actualContent := `= A Title
:toc: preamble

A short preamble

== section 1`

		expectedResult := types.Document{
			Attributes: types.DocumentAttributes{
				"doctitle": doctitleAttribute,
				"toc":      "preamble",
			},
			ElementReferences: types.ElementReferences{
				"_section_1": types.SectionTitle{
					Attributes: types.ElementAttributes{
						types.AttrID: "_section_1",
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "section 1"},
					},
				},
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				preamble,
				tableOfContents,
				section,
			},
		}
		verify(GinkgoT(), expectedResult, actualContent)
	})

})
