package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Parsing Headings", func() {

	It("heading only", func() {
		actualContent := "= a heading"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Heading{
					Level: 1,
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a heading"},
						},
					},
					ID: &types.ElementID{
						Value: "_a_heading",
					},
				},
			}}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("heading invalid1", func() {
		actualContent := "=a heading"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "=a heading"},
							},
						},
					},
				},
			}}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("heading invalid2", func() {
		actualContent := " = a heading"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: " = a heading"},
							},
						},
					},
				},
			}}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("section2", func() {
		actualContent := `== section 1`
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Heading{
					Level: 2,
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "section 1"},
						},
					},
					ID: &types.ElementID{
						Value: "_section_1",
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("heading with section2", func() {
		actualContent := "= a heading\n" +
			"\n" +
			"== section 1"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Heading{
					Level: 1,
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a heading"},
						},
					},
					ID: &types.ElementID{
						Value: "_a_heading",
					},
				},
				&types.BlankLine{},
				&types.Heading{
					Level: 2,
					Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "section 1"},
						},
					},
					ID: &types.ElementID{
						Value: "_section_1",
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("heading with invalid section2", func() {
		actualContent := "= a heading\n" +
			"\n" +
			" == section 1"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Heading{
					Level: 1, Content: &types.InlineContent{
						Elements: []types.DocElement{
							&types.StringElement{Content: "a heading"},
						},
					},
					ID: &types.ElementID{
						Value: "_a_heading",
					},
				},
				&types.BlankLine{},
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: " == section 1"},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})
})
