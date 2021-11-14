package parser

// import (
// 	"github.com/bytesparadise/libasciidoc/pkg/types"

// 	. "github.com/onsi/ginkgo" //nolint golint
// 	. "github.com/onsi/gomega" //nolint golint
// )

// var _ = Describe("include table of contents", func() {

// 	// reusable elements
// 	doctitle := []interface{}{
// 		&types.StringElement{Content: "A Title"},
// 	}
// 	preamble := &types.Preamble{
// 		Elements: []interface{}{
// 			&types.BlankLine{},
// 			&types.Paragraph{
// 				Elements: []interface{}{
// 					&types.StringElement{Content: "A short preamble"},
// 				},
// 			},
// 			&types.BlankLine{},
// 		},
// 	}
// 	section := &types.Section{
// 		Level: 1,
// 		Attributes: types.Attributes{
// 			types.AttrID: "_section_1",
// 		},
// 		Title: []interface{}{
// 			&types.StringElement{Content: "section 1"},
// 		},
// 		Elements: []interface{}{},
// 	}

// 	It("table of contents with default placement and no header with content", func() {
// 		source := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: nil,
// 			},
// 			Elements: []interface{}{
// 				preamble,
// 				section,
// 			},
// 		}
// 		expected := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: nil,
// 			},
// 			Elements: []interface{}{
// 				&types.TableOfContents{},
// 				preamble,
// 				section,
// 			},
// 		}
// 		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
// 	})

// 	It("table of contents with default placement and a header with content", func() {
// 		source := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: nil,
// 			},
// 			ElementReferences: types.ElementReferences{}, // can leave empty for this test
// 			Elements: []interface{}{
// 				&types.Section{
// 					Level: 0,
// 					Title: doctitle,
// 					Attributes: types.Attributes{
// 						types.AttrID: "_a_title",
// 					},
// 					Elements: []interface{}{
// 						preamble,
// 						section,
// 					},
// 				},
// 			},
// 		}
// 		expected := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: nil,
// 			},
// 			ElementReferences: types.ElementReferences{}, // can leave empty for this test
// 			Elements: []interface{}{
// 				&types.Section{
// 					Level: 0,
// 					Title: doctitle,
// 					Attributes: types.Attributes{
// 						types.AttrID: "_a_title",
// 					},
// 					Elements: []interface{}{
// 						&types.TableOfContents{},
// 						preamble,
// 						section,
// 					},
// 				},
// 			},
// 		}
// 		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
// 	})

// 	It("table of contents with default placement and a header without content", func() {
// 		source := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: nil,
// 			},
// 			ElementReferences: types.ElementReferences{}, // can leave empty for this test
// 			Elements: []interface{}{
// 				&types.Section{
// 					Level: 0,
// 					Title: doctitle,
// 					Attributes: types.Attributes{
// 						types.AttrID: "_a_title",
// 					},
// 					Elements: []interface{}{
// 						preamble,
// 					},
// 				},
// 			},
// 		}
// 		expected := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: nil,
// 			},
// 			ElementReferences: types.ElementReferences{}, // can leave empty for this test
// 			Elements: []interface{}{
// 				&types.Section{
// 					Level: 0,
// 					Title: doctitle,
// 					Attributes: types.Attributes{
// 						types.AttrID: "_a_title",
// 					},
// 					Elements: []interface{}{
// 						&types.TableOfContents{},
// 						preamble,
// 					},
// 				},
// 			},
// 		}
// 		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
// 	})

// 	It("table of contents with preamble placement and no header with content", func() {
// 		source := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: "preamble",
// 			},
// 			Elements: []interface{}{
// 				preamble,
// 				section,
// 			},
// 		}
// 		expected := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: "preamble",
// 			},
// 			Elements: []interface{}{
// 				preamble,
// 				&types.TableOfContents{},
// 				section,
// 			},
// 		}
// 		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
// 	})

// 	It("table of contents with preamble placement and header with content", func() {
// 		source := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: "preamble",
// 			},
// 			Elements: []interface{}{
// 				&types.Section{
// 					Level: 0,
// 					Title: doctitle,
// 					Attributes: types.Attributes{
// 						types.AttrID: "_a_title",
// 					},
// 					Elements: []interface{}{
// 						preamble,
// 						section,
// 					},
// 				},
// 			},
// 		}
// 		expected := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: "preamble",
// 			},
// 			Elements: []interface{}{
// 				&types.Section{
// 					Level: 0,
// 					Title: doctitle,
// 					Attributes: types.Attributes{
// 						types.AttrID: "_a_title",
// 					},
// 					Elements: []interface{}{
// 						preamble,
// 						&types.TableOfContents{},
// 						section,
// 					},
// 				},
// 			},
// 		}
// 		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
// 	})

// 	It("table of contents with preamble placement and header without content", func() {
// 		source := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: "preamble",
// 			},
// 			Elements: []interface{}{
// 				&types.Section{
// 					Level: 0,
// 					Title: doctitle,
// 					Attributes: types.Attributes{
// 						types.AttrID: "_a_title",
// 					},
// 					Elements: []interface{}{
// 						preamble,
// 					},
// 				},
// 			},
// 		}
// 		expected := &types.Document{
// 			Attributes: types.Attributes{
// 				types.AttrTableOfContents: "preamble",
// 			},
// 			Elements: []interface{}{
// 				&types.Section{
// 					Level: 0,
// 					Title: doctitle,
// 					Attributes: types.Attributes{
// 						types.AttrID: "_a_title",
// 					},
// 					Elements: []interface{}{
// 						preamble,
// 						&types.TableOfContents{},
// 					},
// 				},
// 			},
// 		}
// 		Expect(includeTableOfContentsPlaceHolder(source)).To(Equal(expected))
// 	})

// })
