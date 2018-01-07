package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Table of Contents", func() {

	// reusable elements
	doctitleAttribute := &types.SectionTitle{
		Content: &types.InlineContent{
			Elements: []types.InlineElement{
				&types.StringElement{Content: "A Title"},
			},
		},
		ID: &types.ElementID{
			Value: "_a_title",
		},
	}
	tableOfContents := &types.TableOfContentsMacro{}
	preamble := &types.Preamble{
		Elements: []types.DocElement{
			&types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "A short preamble"},
						},
					},
				},
			},
		},
	}
	section := &types.Section{
		Level: 1,
		SectionTitle: types.SectionTitle{
			Content: &types.InlineContent{
				Elements: []types.InlineElement{
					&types.StringElement{Content: "section 1"},
				},
			},
			ID: &types.ElementID{
				Value: "_section_1",
			},
		},
		Elements: []types.DocElement{},
	}

	It("TOC with default placement", func() {

		actualContent := `= A Title
:toc:

A short preamble

== section 1`

		expectedDocument := &types.Document{
			Attributes: map[string]interface{}{
				"doctitle": doctitleAttribute,
				"toc":      "",
			},
			ElementReferences: map[string]interface{}{
				"_section_1": &types.SectionTitle{
					Content: &types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "section 1"},
						},
					},
					ID: &types.ElementID{
						Value: "_section_1",
					},
				},
			},
			Elements: []types.DocElement{
				tableOfContents,
				preamble,
				section,
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("TOC with preamble placement", func() {
		actualContent := `= A Title
:toc: preamble

A short preamble

== section 1`

		expectedDocument := &types.Document{
			Attributes: map[string]interface{}{
				"doctitle": doctitleAttribute,
				"toc":      "preamble",
			},
			ElementReferences: map[string]interface{}{
				"_section_1": &types.SectionTitle{
					Content: &types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "section 1"},
						},
					},
					ID: &types.ElementID{
						Value: "_section_1",
					},
				},
			},
			Elements: []types.DocElement{
				preamble,
				tableOfContents,
				section,
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	// Context("TOC macro", func() {

	// })
})
