package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("images", func() {

	Context("block images", func() {

		Context("correct behaviour", func() {

			It("block image with empty alt", func() {
				source := "image::images/foo.png[]"
				expected := types.ImageBlock{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt: "foo",
					},
					Path: "images/foo.png",
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("block image with empty alt and trailing spaces", func() {
				source := "image::images/foo.png[]  \t\t  "
				expected := types.ImageBlock{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt: "foo",
					},
					Path: "images/foo.png",
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("block image with line return", func() {
				// line return here is not considered as a blank line
				source := `image::images/foo.png[]
`
				expected := types.ImageBlock{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt: "foo",
					},
					Path: "images/foo.png",
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("block image with 1 empty blank line", func() {
				// here, there's a real blank line with some spaces
				source := `image::images/foo.png[]
  `
				expected := types.ImageBlock{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt: "foo",
					},
					Path: "images/foo.png",
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("block image with 2 blank lines with spaces and tabs", func() {
				source := `image::images/foo.png[]
			`
				expected := types.ImageBlock{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt: "foo",
					},
					Path: "images/foo.png",
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("block image with alt", func() {
				source := `image::images/foo.png[the foo.png image]`
				expected := types.ImageBlock{
					Attributes: types.ElementAttributes{
						types.AttrImageAlt: "the foo.png image",
					},
					Path: "images/foo.png",
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("block image with dimensions and id link title meta", func() {
				source := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image, 600, 400]`
				expected := types.ImageBlock{
					Attributes: types.ElementAttributes{
						types.AttrID:          "img-foobar",
						types.AttrCustomID:    true,
						types.AttrTitle:       "A title to foobar",
						types.AttrInlineLink:  "http://foo.bar",
						types.AttrImageAlt:    "the foo.png image",
						types.AttrImageWidth:  "600",
						types.AttrImageHeight: "400",
					},
					Path: "images/foo.png",
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("2 block images", func() {
				source := `image::app.png[]
image::appa.png[]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt: "app",
							},
							Path: "app.png",
						},
						types.ImageBlock{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt: "appa",
							},
							Path: "appa.png",
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})
		})

		Context("errors", func() {

			Context("parsing the paragraph only", func() {

				It("block image appending inline content", func() {
					source := "a paragraph\nimage::images/foo.png[]"
					expected := types.Paragraph{
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
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})

			Context("parsing the whole document", func() {

				It("paragraph with block image with alt and dimensions", func() {
					source := "a foo image::foo.png[foo image, 600, 400] bar"
					expected := types.PreflightDocument{
						Blocks: []interface{}{
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
					Expect(source).To(BecomePreflightDocument(expected))
				})
			})
		})
	})

	Context("inline images", func() {

		Context("correct behaviour", func() {

			It("inline image with empty alt only", func() {
				source := "image:images/foo.png[]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "foo",
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image with empty alt and trailing spaces", func() {
				source := "image:images/foo.png[]  \t\t  "
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "foo",
								},
								Path: "images/foo.png",
							},
							types.StringElement{
								Content: "  \t\t  ",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image surrounded with test", func() {
				source := "a foo image:images/foo.png[] bar..."
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "a foo ",
							},
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "foo",
								},
								Path: "images/foo.png",
							},
							types.StringElement{
								Content: " bar...",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image with alt alone", func() {
				source := "image:images/foo.png[the foo.png image]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "the foo.png image",
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image with alt and width", func() {
				source := "image:images/foo.png[the foo.png image, 600]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt:   "the foo.png image",
									types.AttrImageWidth: "600",
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image with alt, width and height", func() {
				source := "image:images/foo.png[the foo.png image, 600, 400]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt:    "the foo.png image",
									types.AttrImageWidth:  "600",
									types.AttrImageHeight: "400",
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image with alt, but empty width and height", func() {
				source := "image:images/foo.png[the foo.png image, , ]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "the foo.png image",
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image with single other attribute only", func() {
				source := "image:images/foo.png[id=myid]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "foo", // based on filename
									types.AttrID:       "myid",
									types.AttrCustomID: true,
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image with multiple other attributes only", func() {
				source := "image:images/foo.png[id=myid, title= mytitle, role = myrole ]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "foo", // based on filename
									types.AttrID:       "myid",
									types.AttrCustomID: true,
									types.AttrTitle:    "mytitle",
									types.AttrRole:     "myrole",
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image with alt, width, height and other attributes", func() {
				source := "image:images/foo.png[ foo, 600, 400, id=myid, title=mytitle, role=myrole ]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt:    "foo",
									types.AttrImageWidth:  "600",
									types.AttrImageHeight: "400",
									types.AttrID:          "myid",
									types.AttrCustomID:    true,
									types.AttrTitle:       "mytitle",
									types.AttrRole:        "myrole",
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image in a paragraph with space after colon", func() {
				source := "this is an image: image:images/foo.png[]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "this is an image: ",
							},
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "foo",
								},
								Path: "images/foo.png",
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("inline image in a paragraph without space keyword", func() {
				source := "this is an inline.image:images/foo.png[]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "this is an inline.",
							},
							types.InlineImage{
								Attributes: types.ElementAttributes{
									types.AttrImageAlt: "foo",
								},
								Path: "images/foo.png",
							},
						},
					},
				}

				Expect(source).To(EqualDocumentBlock(expected))
			})
		})
		Context("errors", func() {
			It("inline image appending inline content", func() {
				source := "a paragraph\nimage::images/foo.png[]"
				expected := types.Paragraph{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})
		})
	})
})
