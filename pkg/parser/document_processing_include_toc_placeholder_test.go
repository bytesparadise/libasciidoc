package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

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
	tocPlaceHolder := types.TableOfContentsPlaceHolder{}

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
				tocPlaceHolder,
				preamble,
				section,
			},
		}
		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
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
						tocPlaceHolder,
						preamble,
						section,
					},
				},
			},
		}
		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
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
						tocPlaceHolder,
						preamble,
					},
				},
			},
		}
		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
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
				tocPlaceHolder,
				section,
			},
		}
		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
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
						tocPlaceHolder,
						section,
					},
				},
			},
		}
		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
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
						tocPlaceHolder,
					},
				},
			},
		}
		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
	})

})
