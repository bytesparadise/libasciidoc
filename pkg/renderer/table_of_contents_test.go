package renderer_test

import (
	"context"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("table of contents", func() {

	// reusable elements
	doctitleAttribute := types.SectionTitle{
		Attributes: types.ElementAttributes{
			types.AttrID:       "a_title",
			types.AttrCustomID: false,
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
				types.AttrID:       "section_1",
				types.AttrCustomID: false,
			},
			Elements: types.InlineElements{
				types.StringElement{Content: "section 1"},
			},
		},
		Attributes: types.ElementAttributes{},
		Elements:   []interface{}{},
	}

	It("table of contents with default placement", func() {

		actualContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: doctitleAttribute,
				"toc":           "",
			},
			ElementReferences: types.ElementReferences{
				"section_1": types.SectionTitle{
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_1",
						types.AttrCustomID: false,
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
				section,
			},
		}

		expectedResult := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: doctitleAttribute,
				"toc":           "",
			},
			ElementReferences: types.ElementReferences{
				"section_1": types.SectionTitle{
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_1",
						types.AttrCustomID: false,
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
		verifyTableOfContents(expectedResult, actualContent)
	})

	It("table of contents with preamble placement", func() {
		actualContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: doctitleAttribute,
				"toc":           "preamble",
			},
			ElementReferences: types.ElementReferences{
				"section_1": types.SectionTitle{
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_1",
						types.AttrCustomID: false,
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
				section,
			},
		}

		expectedResult := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: doctitleAttribute,
				"toc":           "preamble",
			},
			ElementReferences: types.ElementReferences{
				"section_1": types.SectionTitle{
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_1",
						types.AttrCustomID: false,
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
		verifyTableOfContents(expectedResult, actualContent)
	})

})

func verifyTableOfContents(expectedContent, actualContent types.Document) {
	ctx := renderer.Wrap(context.Background(), actualContent)
	renderer.IncludeTableOfContents(ctx)
	assert.Equal(GinkgoT(), expectedContent, ctx.Document)
}
