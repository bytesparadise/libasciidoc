package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("include table of contents", func() {

	// reusable elements
	doctitle := []interface{}{
		types.StringElement{Content: "A Title"},
	}
	preamble := types.Preamble{
		Elements: []interface{}{
			types.BlankLine{},
			types.Paragraph{
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
		Attributes: types.Attributes{
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
			Attributes: types.Attributes{
				types.AttrTableOfContents: "",
			},
			Elements: []interface{}{
				preamble,
				section,
			},
		}
		expected := types.Document{
			Attributes: types.Attributes{
				types.AttrTableOfContents: "",
			},
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
			Attributes: types.Attributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences: types.ElementReferences{}, // can leave empty for this test
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.Attributes{
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
			Attributes: types.Attributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences: types.ElementReferences{}, // can leave empty for this test
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.Attributes{
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
			Attributes: types.Attributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences: types.ElementReferences{}, // can leave empty for this test
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.Attributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						preamble,
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.Attributes{
				types.AttrTableOfContents: "",
			},
			ElementReferences: types.ElementReferences{}, // can leave empty for this test
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.Attributes{
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
			Attributes: types.Attributes{
				types.AttrTableOfContents: "preamble",
			},
			Elements: []interface{}{
				preamble,
				section,
			},
		}
		expected := types.Document{
			Attributes: types.Attributes{
				types.AttrTableOfContents: "preamble",
			},
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
			Attributes: types.Attributes{
				types.AttrTableOfContents: "preamble",
			},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.Attributes{
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
			Attributes: types.Attributes{
				types.AttrTableOfContents: "preamble",
			},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.Attributes{
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
			Attributes: types.Attributes{
				types.AttrTableOfContents: "preamble",
			},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.Attributes{
						types.AttrID: "_a_title",
					},
					Elements: []interface{}{
						preamble,
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.Attributes{
				types.AttrTableOfContents: "preamble",
			},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: doctitle,
					Attributes: types.Attributes{
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
