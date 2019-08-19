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
	doctitle := types.InlineElements{
		types.StringElement{Content: "A Title"},
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
		Attributes: types.ElementAttributes{
			types.AttrID:       "section_1",
			types.AttrCustomID: false,
		},
		Title: types.InlineElements{
			types.StringElement{Content: "section 1"},
		},
		Elements: []interface{}{},
	}
	tableOfContents := types.TableOfContentsMacro{}

	It("table of contents with default placement and no header with content", func() {
		source := types.Document{
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
		expected := types.Document{
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
		verifyTableOfContents(expected, source)
	})

	It("table of contents with default placement and a header with content", func() {
		source := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "a_title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						toc,
						preamble,
						section,
					},
				},
			},
		}
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "a_title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						tableOfContents,
						toc,
						preamble,
						section,
					},
				},
			},
		}
		verifyTableOfContents(expected, source)
	})

	It("table of contents with default placement and a header without content", func() {
		source := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "a_title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						toc,
						preamble,
					},
				},
			},
		}
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "a_title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						tableOfContents,
						toc,
						preamble,
					},
				},
			},
		}
		verifyTableOfContents(expected, source)
	})

	It("table of contents with preamble placement and no header with content", func() {
		source := types.Document{
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
		expected := types.Document{
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
		verifyTableOfContents(expected, source)
	})

	It("table of contents with preamble placement and header with content", func() {
		source := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "a_title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						preambletoc,
						preamble,
						section,
					},
				},
			},
		}
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "a_title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						preambletoc,
						preamble,
						tableOfContents,
						section,
					},
				},
			},
		}
		verifyTableOfContents(expected, source)
	})

	It("table of contents with preamble placement and header without content", func() {
		source := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "a_title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						preambletoc,
						preamble,
					},
				},
			},
		}
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "a_title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						preambletoc,
						preamble,
						tableOfContents,
					},
				},
			},
		}
		verifyTableOfContents(expected, source)
	})

})

func verifyTableOfContents(expectedContent, source types.Document) {
	ctx := renderer.Wrap(context.Background(), source)
	renderer.IncludeTableOfContents(ctx)
	assert.Equal(GinkgoT(), expectedContent, ctx.Document)
}
