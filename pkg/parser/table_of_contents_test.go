package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("table of contents", func() {

	// reusable elements
	doctitleAttribute := types.SectionTitle{
		Attributes: map[string]interface{}{
			types.AttrID: "_a_title",
		},
		Content: types.InlineElements{
			types.StringElement{Content: "A Title"},
		},
	}
	tableOfContents := types.TableOfContentsMacro{}
	preamble := types.Preamble{
		Elements: []interface{}{
			types.BlankLine{},
			types.Paragraph{
				Attributes: map[string]interface{}{},
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
			Attributes: map[string]interface{}{
				types.AttrID: "_section_1",
			},
			Content: types.InlineElements{
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
			Attributes: map[string]interface{}{
				"doctitle": doctitleAttribute,
				"toc":      "",
			},
			ElementReferences: map[string]interface{}{
				"_section_1": types.SectionTitle{
					Attributes: map[string]interface{}{
						types.AttrID: "_section_1",
					},
					Content: types.InlineElements{
						types.StringElement{Content: "section 1"},
					},
				},
			},
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
			Attributes: map[string]interface{}{
				"doctitle": doctitleAttribute,
				"toc":      "preamble",
			},
			ElementReferences: map[string]interface{}{
				"_section_1": types.SectionTitle{
					Attributes: map[string]interface{}{
						types.AttrID: "_section_1",
					},
					Content: types.InlineElements{
						types.StringElement{Content: "section 1"},
					},
				},
			},
			Elements: []interface{}{
				preamble,
				tableOfContents,
				section,
			},
		}
		verify(GinkgoT(), expectedResult, actualContent)
	})

})
