package renderer_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("table of contents", func() {

	// reusable elements
	doctitle := []interface{}{
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
				Lines: [][]interface{}{
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
		Title: []interface{}{
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
		Expect(source).To(HaveTableOfContents(expected))
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
		Expect(source).To(HaveTableOfContents(expected))
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
		Expect(source).To(HaveTableOfContents(expected))
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
		Expect(source).To(HaveTableOfContents(expected))
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
		Expect(source).To(HaveTableOfContents(expected))
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
		Expect(source).To(HaveTableOfContents(expected))
	})

})
