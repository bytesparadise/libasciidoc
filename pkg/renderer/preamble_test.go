package renderer_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("preambles", func() {

	sectionATitle := types.SectionTitle{
		Attributes: types.ElementAttributes{
			types.AttrID:       "section_a",
			types.AttrCustomID: false,
		},
		Elements: types.InlineElements{
			types.StringElement{Content: "Section A"},
		},
	}

	sectionBTitle := types.SectionTitle{
		Attributes: types.ElementAttributes{
			types.AttrID:       "section_b",
			types.AttrCustomID: false,
		},
		Elements: types.InlineElements{
			types.StringElement{Content: "Section B"},
		},
	}

	It("doc without sections", func() {
		actualContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "another short paragraph"},
						},
					},
				},
			},
		}
		expectedContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "another short paragraph"},
						},
					},
				},
			},
		}
		verifyPreamble(expectedContent, actualContent)
	})

	It("doc with 1-paragraph preamble", func() {
		actualContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Level:    1,
					Title:    sectionATitle,
					Elements: []interface{}{},
				},
				types.Section{
					Level:    1,
					Title:    sectionBTitle,
					Elements: []interface{}{},
				},
			},
		}
		expectedContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Preamble{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a short paragraph"},
								},
							},
						},
						types.BlankLine{},
					},
				},
				types.Section{
					Level:    1,
					Title:    sectionATitle,
					Elements: []interface{}{},
				},
				types.Section{
					Level:    1,
					Title:    sectionBTitle,
					Elements: []interface{}{},
				},
			},
		}
		verifyPreamble(expectedContent, actualContent)
	})

	It("doc with 2-paragraph preamble", func() {
		actualContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "another short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Level:    1,
					Title:    sectionATitle,
					Elements: []interface{}{},
				},
				types.Section{
					Level:    1,
					Title:    sectionBTitle,
					Elements: []interface{}{},
				},
			},
		}
		expectedContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Preamble{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a short paragraph"},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "another short paragraph"},
								},
							},
						},
						types.BlankLine{},
					},
				},
				types.Section{
					Level:    1,
					Title:    sectionATitle,
					Elements: []interface{}{},
				},
				types.Section{
					Level:    1,
					Title:    sectionBTitle,
					Elements: []interface{}{},
				},
			},
		}
		verifyPreamble(expectedContent, actualContent)
	})

})

func verifyPreamble(expectedContent, actualContent types.Document) {
	result, err := renderer.IncludePreamble(actualContent)
	assert.NoError(GinkgoT(), err)
	assert.Equal(GinkgoT(), expectedContent, result)
}
