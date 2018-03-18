package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Table of Contents", func() {

	// reusable elements
	doctitleAttribute := types.SectionTitle{
		ID: types.ElementID{
			Value: "_a_title",
		},
		Content: types.InlineContent{
			Elements: []types.InlineElement{
				types.StringElement{Content: "A Title"},
			},
		},
	}
	tableOfContents := types.TableOfContentsMacro{}
	preamble := types.Preamble{
		Elements: []types.DocElement{
			types.Paragraph{
				Lines: []types.InlineContent{
					{
						Elements: []types.InlineElement{
							types.StringElement{Content: "A short preamble"},
						},
					},
				},
			},
		},
	}
	section := types.Section{
		Level: 1,
		Title: types.SectionTitle{
			ID: types.ElementID{
				Value: "_section_1",
			},
			Content: types.InlineContent{
				Elements: []types.InlineElement{
					types.StringElement{Content: "section 1"},
				},
			},
		},
		Elements: []types.DocElement{},
	}

	It("TOC with default placement", func() {

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
					ID: types.ElementID{
						Value: "_section_1",
					},
					Content: types.InlineContent{
						Elements: []types.InlineElement{
							types.StringElement{Content: "section 1"},
						},
					},
				},
			},
			Elements: []types.DocElement{
				tableOfContents,
				preamble,
				section,
			},
		}
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("TOC with preamble placement", func() {
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
					ID: types.ElementID{
						Value: "_section_1",
					},
					Content: types.InlineContent{
						Elements: []types.InlineElement{
							types.StringElement{Content: "section 1"},
						},
					},
				},
			},
			Elements: []types.DocElement{
				preamble,
				tableOfContents,
				section,
			},
		}
		verify(GinkgoT(), expectedResult, actualContent)
	})

	// Context("TOC macro", func() {

	// })
})
