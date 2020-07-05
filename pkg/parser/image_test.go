package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("images", func() {

	Context("block images", func() {

		Context("inline elements", func() {

			It("with empty alt", func() {
				source := "image::images/foo.png[]"
				expected := types.ImageBlock{
					Location: types.Location{
						Path: []interface{}{
							types.StringElement{Content: "images/foo.png"},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("with empty alt and trailing spaces", func() {
				source := "image::images/foo.png[]  \t\t  "
				expected := types.ImageBlock{
					Location: types.Location{
						Path: []interface{}{
							types.StringElement{Content: "images/foo.png"},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("with alt", func() {
				source := `image::images/foo.png[the foo.png image]`
				expected := types.ImageBlock{
					Attributes: types.Attributes{
						types.AttrImageAlt: "the foo.png image",
					},
					Location: types.Location{
						Path: []interface{}{
							types.StringElement{Content: "images/foo.png"},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("with dimensions and id link title meta", func() {
				source := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image, 600, 400]`
				expected := types.ImageBlock{
					Attributes: types.Attributes{
						types.AttrID:          "img-foobar",
						types.AttrCustomID:    true,
						types.AttrTitle:       "A title to foobar",
						types.AttrInlineLink:  "http://foo.bar",
						types.AttrImageAlt:    "the foo.png image",
						types.AttrWidth:       "600",
						types.AttrImageHeight: "400",
					},
					Location: types.Location{
						Path: []interface{}{
							types.StringElement{Content: "images/foo.png"},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})
		})

		Context("final document", func() {

			It("with empty alt", func() {
				source := "image::images/foo.png[]"
				expected := types.Document{
					Elements: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with implicit imagesdir document attribute", func() {
				source := `
:imagesdir: ./path/to/images

image::foo.png[]`
				expected := types.Document{
					Attributes: types.Attributes{
						"imagesdir": "./path/to/images",
					},
					Elements: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "./path/to/images/foo.png"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with document attribute in URL", func() {
				source := `
:dir: ./path/to/images

image::{dir}/foo.png[]`
				expected := types.Document{
					Attributes: types.Attributes{
						"dir": "./path/to/images",
					},
					Elements: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "./path/to/images/foo.png"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with implicit imagesdir", func() {
				source := `
:imagesdir: ./path/to/images

image::foo.png[]`
				expected := types.Document{
					Attributes: types.Attributes{
						"imagesdir": "./path/to/images",
					},
					Elements: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "./path/to/images/foo.png"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with explicit duplicate imagesdir document attribute", func() {
				source := `
:imagesdir: ./path/to/images

image::{imagesdir}/foo.png[]`
				expected := types.Document{
					Attributes: types.Attributes{
						"imagesdir": "./path/to/images",
					},
					Elements: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "./path/to/images/./path/to/images/foo.png"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("2 block images", func() {
				source := `image::images/foo.png[]
image::images/bar.png[]`
				expected := types.Document{
					Elements: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "images/foo.png"},
								},
							},
						},
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: "bar",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "images/bar.png"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("errors", func() {

			It("appending inline content", func() {
				source := "a paragraph\nimage::images/foo.png[]"
				expected := types.Paragraph{
					Lines: []interface{}{
						types.RawLine{Content: "a paragraph"},
						types.RawLine{Content: "image::images/foo.png[]"},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("paragraph with block image with alt and dimensions", func() {
				source := "a foo image::foo.png[foo image, 600, 400] bar"
				expected := []interface{}{
					types.StringElement{Content: "a foo image::foo.png[foo image, 600, 400] bar"},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})
		})
	})

	Context("inline images", func() {

		Context("inline elements", func() {

			It("inline image with empty alt only", func() {
				source := "image:images/foo.png[]"
				expected := []interface{}{
					types.InlineImage{
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image with empty alt and trailing spaces", func() {
				source := "image:images/foo.png[]  \t\t  "
				expected := []interface{}{
					types.InlineImage{
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
					types.StringElement{
						Content: "  \t\t  ",
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image surrounded with test", func() {
				source := "a foo image:images/foo.png[] bar..."
				expected := []interface{}{
					types.StringElement{
						Content: "a foo ",
					},
					types.InlineImage{
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
					types.StringElement{
						Content: " bar\u2026\u200b",
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image with alt alone", func() {
				source := "image:images/foo.png[the foo.png image]"
				expected := []interface{}{
					types.InlineImage{
						Attributes: types.Attributes{
							types.AttrImageAlt: "the foo.png image",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image with alt and width", func() {
				source := "image:images/foo.png[the foo.png image, 600]"
				expected := []interface{}{
					types.InlineImage{
						Attributes: types.Attributes{
							types.AttrImageAlt: "the foo.png image",
							types.AttrWidth:    "600",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image with alt, width and height", func() {
				source := "image:images/foo.png[the foo.png image, 600, 400]"
				expected := []interface{}{
					types.InlineImage{
						Attributes: types.Attributes{
							types.AttrImageAlt:    "the foo.png image",
							types.AttrWidth:       "600",
							types.AttrImageHeight: "400",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image with alt, but empty width and height", func() {
				source := "image:images/foo.png[the foo.png image, , ]"
				expected := []interface{}{
					types.InlineImage{
						Attributes: types.Attributes{
							types.AttrImageAlt: "the foo.png image",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image with single other attribute only", func() {
				source := "image:images/foo.png[id=myid]"
				expected := []interface{}{
					types.InlineImage{
						Attributes: types.Attributes{
							types.AttrID:       "myid",
							types.AttrCustomID: true,
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image with multiple other attributes only", func() {
				source := "image:images/foo.png[id=myid, title= mytitle, role = myrole ]"
				expected := []interface{}{
					types.InlineImage{
						Attributes: types.Attributes{
							types.AttrID:       "myid",
							types.AttrCustomID: true,
							types.AttrTitle:    "mytitle",
							types.AttrRole:     "myrole",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image with alt, width, height and other attributes", func() {
				source := "image:images/foo.png[ foo, 600, 400, id=myid, title=mytitle, role=myrole ]"
				expected := []interface{}{
					types.InlineImage{
						Attributes: types.Attributes{
							types.AttrImageAlt:    "foo",
							types.AttrWidth:       "600",
							types.AttrImageHeight: "400",
							types.AttrID:          "myid",
							types.AttrCustomID:    true,
							types.AttrTitle:       "mytitle",
							types.AttrRole:        "myrole",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image in a paragraph with space after colon", func() {
				source := "this is an image: image:images/foo.png[]"
				expected := []interface{}{
					types.StringElement{
						Content: "this is an image: ",
					},
					types.InlineImage{
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

			It("inline image in a paragraph without space separator", func() {
				source := "this is an inline.image:images/foo.png[]"
				expected := []interface{}{
					types.StringElement{
						Content: "this is an inline.",
					},
					types.InlineImage{
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				}
				Expect(ParseInlineElements(source)).To(MatchInlineElements(expected))
			})

		})

		Context("final document", func() {

			It("inline image with empty alt only", func() {
				source := "image:images/foo.png[]"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.InlineImage{
										Attributes: types.Attributes{
											types.AttrImageAlt: "foo",
										},
										Location: types.Location{
											Path: []interface{}{
												types.StringElement{Content: "images/foo.png"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline image with document attribute in URL", func() {
				source := `
:dir: ./path/to/images

an image:{dir}/foo.png[].`
				expected := types.Document{
					Attributes: types.Attributes{
						"dir": "./path/to/images",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "an "},
									types.InlineImage{
										Attributes: types.Attributes{
											types.AttrImageAlt: "foo",
										},
										Location: types.Location{
											Path: []interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline image with implicit imagesdir document attribute", func() {
				source := `
:imagesdir: ./path/to/images

an image:foo.png[].`
				expected := types.Document{
					Attributes: types.Attributes{
						"imagesdir": "./path/to/images",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "an "},
									types.InlineImage{
										Attributes: types.Attributes{
											types.AttrImageAlt: "foo",
										},
										Location: types.Location{
											Path: []interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline image with explicit duplicate imagesdir document attribute", func() {
				source := `
:imagesdir: ./path/to/images

an image:{imagesdir}/foo.png[].`
				expected := types.Document{
					Attributes: types.Attributes{
						"imagesdir": "./path/to/images",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "an "},
									types.InlineImage{
										Attributes: types.Attributes{
											types.AttrImageAlt: "foo",
										},
										Location: types.Location{
											Path: []interface{}{
												types.StringElement{Content: "./path/to/images/./path/to/images/foo.png"},
											},
										},
									},
									types.StringElement{Content: "."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with document attribute in URL", func() {
				source := `:path: ./path/to/images

image::{path}/foo.png[]`
				expected := types.Document{
					Attributes: types.Attributes{
						"path": "./path/to/images",
					},
					Elements: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: "foo",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "./path/to/images/foo.png"}, // resolved
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("errors", func() {

			It("inline image appending inline content", func() {
				source := "a paragraph\nimage::images/foo.png[]"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "a paragraph"},
								},
								[]interface{}{
									types.StringElement{Content: "image::images/foo.png[]"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
