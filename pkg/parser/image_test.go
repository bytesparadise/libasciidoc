package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("images", func() {

	Context("block images", func() {

		Context("draft document", func() {

			It("block image with empty alt", func() {
				source := "image::images/foo.png[]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("block image with empty alt and trailing spaces", func() {
				source := "image::images/foo.png[]  \t\t  "
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("block image with line return", func() {
				// line return here is not considered as a blank line
				source := `image::images/foo.png[]
`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("block image with 1 empty blank line", func() {
				// here, there's a real blank line with some spaces
				source := `image::images/foo.png[]
  `
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
						types.BlankLine{},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("block image with 2 blank lines with spaces and tabs", func() {
				source := `image::images/foo.png[]
			`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
						types.BlankLine{},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("block image with alt", func() {
				source := `image::images/foo.png[the foo.png image]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt: "the foo.png image",
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("block image with dimensions and id link title meta", func() {
				source := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image, 600, 400]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{
								types.AttrID:          "img-foobar",
								types.AttrCustomID:    true,
								types.AttrTitle:       "A title to foobar",
								types.AttrInlineLink:  "http://foo.bar",
								types.AttrImageAlt:    "the foo.png image",
								types.AttrImageWidth:  "600",
								types.AttrImageHeight: "400",
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("2 block images", func() {
				source := `image::images/foo.png[]
image::images/bar.png[]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
						types.ImageBlock{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/bar.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})
		})

		Context("final document", func() {

			It("block image with empty alt", func() {
				source := "image::images/foo.png[]"
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.ImageBlock{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("image block with document attribute in URL", func() {
				source := `:imagesdir: ./path/to/images

image::{imagesdir}/foo.png[]`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{
							Name:  "imagesdir",
							Value: "./path/to/images",
						},
						types.ImageBlock{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{Content: "./path/to/images/foo.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})
		})

		Context("errors", func() {

			Context("parsing the paragraph only", func() {

				It("block image appending inline content", func() {
					source := "a paragraph\nimage::images/foo.png[]"
					expected := types.DraftDocument{
						Blocks: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a paragraph"},
									},
									{
										types.StringElement{Content: "image::images/foo.png[]"},
									},
								},
							},
						},
					}
					Expect(source).To(BecomeDraftDocument(expected))
				})
			})

			Context("parsing the whole document", func() {

				It("paragraph with block image with alt and dimensions", func() {
					source := "a foo image::foo.png[foo image, 600, 400] bar"
					expected := types.DraftDocument{
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
					Expect(source).To(BecomeDraftDocument(expected))
				})
			})
		})
	})

	Context("inline images", func() {

		Context("draft document", func() {

			It("inline image with empty alt only", func() {
				source := "image:images/foo.png[]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image with empty alt and trailing spaces", func() {
				source := "image:images/foo.png[]  \t\t  "
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
									types.StringElement{
										Content: "  \t\t  ",
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image surrounded with test", func() {
				source := "a foo image:images/foo.png[] bar..."
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "a foo ",
									},
									types.InlineImage{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
									types.StringElement{
										Content: " bar...",
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image with alt alone", func() {
				source := "image:images/foo.png[the foo.png image]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{
											types.AttrImageAlt: "the foo.png image",
										},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image with alt and width", func() {
				source := "image:images/foo.png[the foo.png image, 600]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{
											types.AttrImageAlt:   "the foo.png image",
											types.AttrImageWidth: "600",
										},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image with alt, width and height", func() {
				source := "image:images/foo.png[the foo.png image, 600, 400]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{
											types.AttrImageAlt:    "the foo.png image",
											types.AttrImageWidth:  "600",
											types.AttrImageHeight: "400",
										},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image with alt, but empty width and height", func() {
				source := "image:images/foo.png[the foo.png image, , ]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{
											types.AttrImageAlt: "the foo.png image",
										},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image with single other attribute only", func() {
				source := "image:images/foo.png[id=myid]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{
											types.AttrID:       "myid",
											types.AttrCustomID: true,
										},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image with multiple other attributes only", func() {
				source := "image:images/foo.png[id=myid, title= mytitle, role = myrole ]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{
											types.AttrID:       "myid",
											types.AttrCustomID: true,
											types.AttrTitle:    "mytitle",
											types.AttrRole:     "myrole",
										},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image with alt, width, height and other attributes", func() {
				source := "image:images/foo.png[ foo, 600, 400, id=myid, title=mytitle, role=myrole ]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
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
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image in a paragraph with space after colon", func() {
				source := "this is an image: image:images/foo.png[]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "this is an image: ",
									},
									types.InlineImage{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("inline image in a paragraph without space separator", func() {
				source := "this is an inline.image:images/foo.png[]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "this is an inline.",
									},
									types.InlineImage{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})

			It("image block with document attribute in URL", func() {
				source := `:imagesdir: ./path/to/images

image::{imagesdir}/foo.png[]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DocumentAttributeDeclaration{
							Name:  "imagesdir",
							Value: "./path/to/images",
						},
						types.BlankLine{},
						types.ImageBlock{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.DocumentAttributeSubstitution{
										Name: "imagesdir",
									},
									types.StringElement{Content: "/foo.png"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})
		})

		Context("final document", func() {

			It("inline image with empty alt only", func() {
				source := "image:images/foo.png[]"
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.InlineImage{
										Attributes: types.ElementAttributes{
											types.AttrImageAlt: "foo",
										},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})

			It("inline image with document attribute in URL", func() {
				source := `:imagesdir: ./path/to/images

an image:{imagesdir}/foo.png[].`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{
							Name:  "imagesdir",
							Value: "./path/to/images",
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "an "},
									types.InlineImage{
										Attributes: types.ElementAttributes{
											types.AttrImageAlt: "foo",
										},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{Content: "./path/to/images/foo.png"},
											},
										},
									},
									types.StringElement{Content: "."},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDocument(expected))
			})
		})

		Context("errors", func() {
			It("inline image appending inline content", func() {
				source := "a paragraph\nimage::images/foo.png[]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a paragraph"},
								},
								{
									types.StringElement{Content: "image::images/foo.png[]"},
								},
							},
						},
					},
				}
				Expect(source).To(BecomeDraftDocument(expected))
			})
		})
	})
})
