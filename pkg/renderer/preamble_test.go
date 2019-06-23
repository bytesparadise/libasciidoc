package renderer_test

import (
	"context"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("preambles", func() {

	sectionATitle := types.InlineElements{
		&types.StringElement{Content: "Section A"},
	}

	sectionBTitle := types.InlineElements{
		&types.StringElement{Content: "Section B"},
	}

	It("doc without sections", func() {
		source := &types.Document{
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
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{Content: "another short paragraph"},
						},
					},
				},
			},
		}
		expectedContent := &types.Document{
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
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{Content: "another short paragraph"},
						},
					},
				},
			},
		}
		verifyPreamble(expectedContent, source)
	})

	It("doc with 1-paragraph preamble", func() {
		source := &types.Document{
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
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				&types.Section{
					Level: 1,
					Title: sectionATitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_a",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				&types.Section{
					Level: 1,
					Title: sectionBTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_b",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{},
				},
			},
		}
		expectedContent := &types.Document{
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
				&types.Preamble{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									&types.StringElement{Content: "a short paragraph"},
								},
							},
						},
						types.BlankLine{},
					},
				},
				&types.Section{
					Level: 1,
					Title: sectionATitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_a",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				&types.Section{
					Level: 1,
					Title: sectionBTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_b",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{},
				},
			},
		}
		verifyPreamble(expectedContent, source)
	})

	It("doc with 2-paragraph preamble", func() {
		source := &types.Document{
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
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{Content: "another short paragraph"},
						},
					},
				},
				types.BlankLine{},
				&types.Section{
					Level: 1,
					Title: sectionATitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_a",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				&types.Section{
					Level: 1,
					Title: sectionBTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_b",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{},
				},
			},
		}
		expectedContent := &types.Document{
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
				&types.Preamble{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									&types.StringElement{Content: "a short paragraph"},
								},
							},
						},
						types.BlankLine{},
						&types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									&types.StringElement{Content: "another short paragraph"},
								},
							},
						},
						types.BlankLine{},
					},
				},
				&types.Section{
					Level: 1,
					Title: sectionATitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_a",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				&types.Section{
					Level: 1,
					Title: sectionBTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_b",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{},
				},
			},
		}
		verifyPreamble(expectedContent, source)
	})

})

func verifyPreamble(expectedContent, source *types.Document) {
	ctx := renderer.Wrap(context.Background(), source)
	renderer.IncludePreamble(ctx)
	assert.Equal(GinkgoT(), expectedContent, ctx.Document)
}
