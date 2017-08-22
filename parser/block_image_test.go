package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Parsing Block Images", func() {
	Context("Correct behaviour", func() {

		It("block image with empty alt", func() {
			actualContent := "image::images/foo.png[]"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.BlockImage{
						Macro: types.BlockImageMacro{
							Path: "images/foo.png",
							Alt:  "foo",
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("block image with empty alt and trailing spaces", func() {
			actualContent := "image::images/foo.png[]  \t\t  "
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.BlockImage{
						Macro: types.BlockImageMacro{
							Path: "images/foo.png",
							Alt:  "foo",
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("block image with line return", func() {
			// line return here is not considered as a blank line
			actualContent := `image::images/foo.png[]
`
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.BlockImage{
						Macro: types.BlockImageMacro{
							Path: "images/foo.png",
							Alt:  "foo",
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("block image with 1 empty blank line", func() {
			// here, there's a real blank line with some spaces
			actualContent := `image::images/foo.png[]
  `
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.BlockImage{
						Macro: types.BlockImageMacro{
							Path: "images/foo.png",
							Alt:  "foo",
						},
					},
					// &types.BlankLine{},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("block image with 2 blank lines with spaces and tabs", func() {
			actualContent := `image::images/foo.png[]
 
			`
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.BlockImage{
						Macro: types.BlockImageMacro{
							Path: "images/foo.png",
							Alt:  "foo",
						},
					},
					// &types.BlankLine{},
					// &types.BlankLine{},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("block image with alt", func() {
			actualContent := "image::images/foo.png[the foo.png image]"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.BlockImage{
						Macro: types.BlockImageMacro{
							Path: "images/foo.png",
							Alt:  "the foo.png image",
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("block image with dimensions and i d link title meta", func() {
			actualContent := "[#img-foobar]\n" +
				".A title to foobar\n" +
				"[link=http://foo.bar]\n" +
				"image::images/foo.png[the foo.png image, 600, 400]"
			width := "600"
			height := "400"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.BlockImage{
						Macro: types.BlockImageMacro{
							Path:   "images/foo.png",
							Alt:    "the foo.png image",
							Width:  &width,
							Height: &height,
						},
						ID:    &types.ElementID{Value: "img-foobar"},
						Title: &types.ElementTitle{Value: "A title to foobar"},
						Link:  &types.ElementLink{Path: "http://foo.bar"},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})
	Context("Errors", func() {
		It("block image appending inline content", func() {
			actualContent := "a paragraph\nimage::images/foo.png[]"
			expectedDocument := &types.Document{
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.DocElement{
									&types.StringElement{Content: "a paragraph"},
								},
							},
							&types.InlineContent{
								Elements: []types.DocElement{
									&types.StringElement{Content: "image::images/foo.png[]"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})
})
