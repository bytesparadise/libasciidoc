package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Parsing Block Images", func() {

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
		compare(GinkgoT(), expectedDocument, actualContent)
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
		compare(GinkgoT(), expectedDocument, actualContent)
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
					Title: &types.ElementTitle{Content: "A title to foobar"},
					Link:  &types.ElementLink{Path: "http://foo.bar"},
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})

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
		compare(GinkgoT(), expectedDocument, actualContent)
	})
})
