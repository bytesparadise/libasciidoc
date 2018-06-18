package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("images", func() {

	Context("block Images", func() {

		Context("correct behaviour", func() {

			It("block image with empty alt", func() {
				actualContent := "image::images/foo.png[]"
				expectedResult := types.BlockImage{
					Attributes: map[string]interface{}{},
					Macro: types.ImageMacro{
						Path: "images/foo.png",
						Alt:  "foo",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with empty alt and trailing spaces", func() {
				actualContent := "image::images/foo.png[]  \t\t  "
				expectedResult := types.BlockImage{
					Attributes: map[string]interface{}{},
					Macro: types.ImageMacro{
						Path: "images/foo.png",
						Alt:  "foo",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with line return", func() {
				// line return here is not considered as a blank line
				actualContent := `image::images/foo.png[]
`
				expectedResult := types.BlockImage{
					Attributes: map[string]interface{}{},
					Macro: types.ImageMacro{
						Path: "images/foo.png",
						Alt:  "foo",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with 1 empty blank line", func() {
				// here, there's a real blank line with some spaces
				actualContent := `image::images/foo.png[]
  `
				expectedResult := types.BlockImage{
					Attributes: map[string]interface{}{},
					Macro: types.ImageMacro{
						Path: "images/foo.png",
						Alt:  "foo",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with 2 blank lines with spaces and tabs", func() {
				actualContent := `image::images/foo.png[]
			`
				expectedResult := types.BlockImage{
					Attributes: map[string]interface{}{},
					Macro: types.ImageMacro{
						Path: "images/foo.png",
						Alt:  "foo",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with alt", func() {
				actualContent := "image::images/foo.png[the foo.png image]"
				expectedResult := types.BlockImage{
					Attributes: map[string]interface{}{},
					Macro: types.ImageMacro{
						Path: "images/foo.png",
						Alt:  "the foo.png image",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with dimensions and id link title meta", func() {
				actualContent := "[#img-foobar]\n" +
					".A title to foobar\n" +
					"[link=http://foo.bar]\n" +
					"image::images/foo.png[the foo.png image, 600, 400]"
				width := "600"
				height := "400"
				expectedResult := types.BlockImage{
					Attributes: map[string]interface{}{
						types.AttrID:    "img-foobar",
						types.AttrTitle: "A title to foobar",
						"link":          "http://foo.bar",
					},
					Macro: types.ImageMacro{
						Path:   "images/foo.png",
						Alt:    "the foo.png image",
						Width:  &width,
						Height: &height,
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})

		Context("errors", func() {

			Context("parsing the paragraph only", func() {

				It("block image appending inline content", func() {
					actualContent := "a paragraph\nimage::images/foo.png[]"
					expectedResult := types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a paragraph"},
							},
							{
								types.StringElement{Content: "image::images/foo.png[]"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
				})
			})

			Context("parsing the whole document", func() {

				It("paragraph with block image with alt and dimensions", func() {
					actualContent := "a foo image::foo.png[foo image, 600, 400] bar"
					expectedResult := types.Document{
						Attributes:        map[string]interface{}{},
						ElementReferences: map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a foo image::foo.png[foo image, 600, 400] bar"},
									},
								},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent)
				})
			})
		})
	})

	Context("inline Images", func() {

		Context("correct behaviour", func() {

			It("inline image with empty alt", func() {
				actualContent := "image:images/foo.png[]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Macro: types.ImageMacro{
							Path: "images/foo.png",
							Alt:  "foo",
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("inline image with empty alt and trailing spaces", func() {
				actualContent := "image:images/foo.png[]  \t\t  "
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Macro: types.ImageMacro{
									Path: "images/foo.png",
									Alt:  "foo",
								},
							},
							types.StringElement{
								Content: "  \t\t  ",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("inline image surrounded with test", func() {
				actualContent := "a foo image:images/foo.png[] bar..."
				expectedResult := types.InlineElements{
					types.StringElement{
						Content: "a foo ",
					},
					types.InlineImage{
						Macro: types.ImageMacro{
							Path: "images/foo.png",
							Alt:  "foo",
						},
					},
					types.StringElement{
						Content: " bar...",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("inline image with alt", func() {
				actualContent := "image:images/foo.png[the foo.png image]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Macro: types.ImageMacro{
							Path: "images/foo.png",
							Alt:  "the foo.png image",
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})
		})
		Context("errors", func() {
			It("inline image appending inline content", func() {
				actualContent := "a paragraph\nimage::images/foo.png[]"
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a paragraph"},
						},
						{
							types.StringElement{Content: "image::images/foo.png[]"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})
		})
	})
})
