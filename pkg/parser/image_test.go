package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("images", func() {

	Context("block images", func() {

		Context("correct behaviour", func() {

			It("block image with empty alt", func() {
				actualContent := "image::images/foo.png[]"
				expectedResult := types.BlockImage{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt:    "foo",
						types.AttrImageWidth:  "",
						types.AttrImageHeight: "",
					},
					Path: "images/foo.png",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with empty alt and trailing spaces", func() {
				actualContent := "image::images/foo.png[]  \t\t  "
				expectedResult := types.BlockImage{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt:    "foo",
						types.AttrImageWidth:  "",
						types.AttrImageHeight: "",
					},
					Path: "images/foo.png",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with line return", func() {
				// line return here is not considered as a blank line
				actualContent := `image::images/foo.png[]
`
				expectedResult := types.BlockImage{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt:    "foo",
						types.AttrImageWidth:  "",
						types.AttrImageHeight: "",
					},
					Path: "images/foo.png",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with 1 empty blank line", func() {
				// here, there's a real blank line with some spaces
				actualContent := `image::images/foo.png[]
  `
				expectedResult := types.BlockImage{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt:    "foo",
						types.AttrImageWidth:  "",
						types.AttrImageHeight: "",
					},
					Path: "images/foo.png",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with 2 blank lines with spaces and tabs", func() {
				actualContent := `image::images/foo.png[]
			`
				expectedResult := types.BlockImage{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt:    "foo",
						types.AttrImageWidth:  "",
						types.AttrImageHeight: "",
					},
					Path: "images/foo.png",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with alt", func() {
				actualContent := `image::images/foo.png[the foo.png image]`
				expectedResult := types.BlockImage{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt:    "the foo.png image",
						types.AttrImageWidth:  "",
						types.AttrImageHeight: "",
					},
					Path: "images/foo.png",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("block image with dimensions and id link title meta", func() {
				actualContent := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image, 600, 400]`
				expectedResult := types.BlockImage{
					Attributes: types.ElementAttributes{
						types.AttrID:          "img-foobar",
						types.AttrTitle:       "A title to foobar",
						types.AttrInlineLink:  "http://foo.bar",
						types.AttrImageAlt:    "the foo.png image",
						types.AttrImageWidth:  "600",
						types.AttrImageHeight: "400",
					},
					Path: "images/foo.png",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("2 block images", func() {
				actualContent := `image::app.png[]
image::appa.png[]`
				expectedResult := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.BlockImage{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt:    "app",
								types.AttrImageWidth:  "",
								types.AttrImageHeight: "",
							},
							Path: "app.png",
						},
						types.BlockImage{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt:    "appa",
								types.AttrImageWidth:  "",
								types.AttrImageHeight: "",
							},
							Path: "appa.png",
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})

		Context("errors", func() {

			Context("parsing the paragraph only", func() {

				It("block image appending inline content", func() {
					actualContent := "a paragraph\nimage::images/foo.png[]"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
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
						Attributes:         map[string]interface{}{},
						ElementReferences:  map[string]interface{}{},
						Footnotes:          types.Footnotes{},
						FootnoteReferences: types.FootnoteReferences{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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

	Context("inline images", func() {

		Context("correct behaviour", func() {

			It("inline image with empty alt", func() {
				actualContent := "image:images/foo.png[]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Attributes: types.ElementAttributes{
							types.AttrImageAlt:    "foo",
							types.AttrImageWidth:  "",
							types.AttrImageHeight: "",
						},
						Path: "images/foo.png",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("inline image with empty alt and trailing spaces", func() {
				actualContent := "image:images/foo.png[]  \t\t  "
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt:    "foo",
									types.AttrImageWidth:  "",
									types.AttrImageHeight: "",
								},
								Path: "images/foo.png",
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
						Attributes: types.ElementAttributes{
							types.AttrImageAlt:    "foo",
							types.AttrImageWidth:  "",
							types.AttrImageHeight: "",
						},
						Path: "images/foo.png",
					},
					types.StringElement{
						Content: " bar...",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("inline image with alt alone", func() {
				actualContent := "image:images/foo.png[the foo.png image]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Attributes: types.ElementAttributes{
							types.AttrImageAlt:    "the foo.png image",
							types.AttrImageWidth:  "",
							types.AttrImageHeight: "",
						},
						Path: "images/foo.png",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("inline image with alt and width", func() {
				actualContent := "image:images/foo.png[the foo.png image, 600]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Attributes: types.ElementAttributes{
							types.AttrImageAlt:    "the foo.png image",
							types.AttrImageWidth:  "600",
							types.AttrImageHeight: "",
						},
						Path: "images/foo.png",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("inline image with alt, width and height", func() {
				actualContent := "image:images/foo.png[the foo.png image, 600, 400]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Attributes: types.ElementAttributes{
							types.AttrImageAlt:    "the foo.png image",
							types.AttrImageWidth:  "600",
							types.AttrImageHeight: "400",
						},
						Path: "images/foo.png",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("inline image with single other attribute only", func() {
				actualContent := "image:images/foo.png[id=myid]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Attributes: types.ElementAttributes{
							types.AttrImageAlt:    "foo", // based on filename
							types.AttrImageWidth:  "",
							types.AttrImageHeight: "",
							types.AttrID:          "myid",
						},
						Path: "images/foo.png",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"), parser.Debug(false))
			})

			It("inline image with multiple other attributes only", func() {
				actualContent := "image:images/foo.png[id=myid, title= mytitle, role = myrole ]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Attributes: types.ElementAttributes{
							types.AttrImageAlt:    "foo", // based on filename
							types.AttrImageWidth:  "",
							types.AttrImageHeight: "",
							types.AttrID:          "myid",
							types.AttrTitle:       "mytitle",
							types.AttrRole:        "myrole",
						},
						Path: "images/foo.png",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("inline image with alt, width, height and other attributes", func() {
				actualContent := "image:images/foo.png[ foo, 600, 400, id=myid, title=mytitle, role=myrole ]"
				expectedResult := types.InlineElements{
					types.InlineImage{
						Attributes: types.ElementAttributes{
							types.AttrImageAlt:    "foo",
							types.AttrImageWidth:  "600",
							types.AttrImageHeight: "400",
							types.AttrID:          "myid",
							types.AttrTitle:       "mytitle",
							types.AttrRole:        "myrole",
						},
						Path: "images/foo.png",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})
		})
		Context("errors", func() {
			It("inline image appending inline content", func() {
				actualContent := "a paragraph\nimage::images/foo.png[]"
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
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
