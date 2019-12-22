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
			types.AttrID: "_section_1",
		},
		Title: []interface{}{
			types.StringElement{Content: "section 1"},
		},
		Elements: []interface{}{},
	}
	tableOfContents := types.TableOfContentsMacro{}

	It("table of contents with default placement and no header with content", func() {
		source := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				preamble,
				section,
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				tableOfContents,
				preamble,
				section,
			},
		}
		Expect(source).To(HaveTableOfContents(expected))
	})

	It("table of contents with default placement and a header with content", func() {
		source := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						preamble,
						section,
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						tableOfContents,
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
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						preamble,
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences:  types.ElementReferences{}, // can leave empty for this test
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						tableOfContents,
						preamble,
					},
				},
			},
		}
		Expect(source).To(HaveTableOfContents(expected))
	})

	It("table of contents with preamble placement and no header with content", func() {
		source := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "preamble",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				preamble,
				section,
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "preamble",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				preamble,
				tableOfContents,
				section,
			},
		}
		Expect(source).To(HaveTableOfContents(expected))
	})

	It("table of contents with preamble placement and header with content", func() {
		source := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "preamble",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						preamble,
						section,
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "preamble",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
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
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "preamble",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						preamble,
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTableOfContents: "preamble",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						preamble,
						tableOfContents,
					},
				},
			},
		}
		Expect(source).To(HaveTableOfContents(expected))
	})

})
