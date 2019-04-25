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
	doctitle := types.SectionTitle{
		Attributes: types.ElementAttributes{
			types.AttrID:       "a_title",
			types.AttrCustomID: false,
		},
		Elements: types.InlineElements{
			types.StringElement{Content: "A Title"},
		},
	}
	toc := types.DocumentAttributeDeclaration{
		Name: "toc",
	}
	preambletoc := types.DocumentAttributeDeclaration{
		Name:  "toc",
		Value: "preamble",
	}
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
	tableOfContents := types.TableOfContentsMacro{}

	It("table of contents with default placement and no header with content", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				toc,
				preamble,
				section,
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				tableOfContents,
				toc,
				preamble,
				section,
			},
		}
		verifyTableOfContents(expectedResult, actualContent)
	})

	It("table of contents with default placement and a header with content", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Title:      doctitle,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						toc,
						preamble,
						section,
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Title:      doctitle,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						tableOfContents,
						toc,
						preamble,
						section,
					},
				},
			},
		}
		verifyTableOfContents(expectedResult, actualContent)
	})

	It("table of contents with default placement and a header without content", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Title:      doctitle,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						toc,
						preamble,
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Title:      doctitle,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						tableOfContents,
						toc,
						preamble,
					},
				},
			},
		}
		verifyTableOfContents(expectedResult, actualContent)
	})

	It("table of contents with preamble placement and no header with content", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				preambletoc,
				preamble,
				section,
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				preambletoc,
				preamble,
				tableOfContents,
				section,
			},
		}
		verifyTableOfContents(expectedResult, actualContent)
	})

	It("table of contents with preamble placement and header with content", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Title:      doctitle,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						preambletoc,
						preamble,
						section,
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Title:      doctitle,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						preambletoc,
						preamble,
						tableOfContents,
						section,
					},
				},
			},
		}
		verifyTableOfContents(expectedResult, actualContent)
	})

	It("table of contents with preamble placement and header without content", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Title:      doctitle,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						preambletoc,
						preamble,
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:      0,
					Title:      doctitle,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						preambletoc,
						preamble,
						tableOfContents,
					},
				},
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
