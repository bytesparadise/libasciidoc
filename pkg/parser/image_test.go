package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("block images", func() {

	Context("draft documents", func() {

		It("with empty alt", func() {
			source := "image::images/foo.png[]"
			expected := types.DraftDocument{
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with empty alt and trailing spaces", func() {
			source := "image::images/foo.png[]  \t\t  "
			expected := types.DraftDocument{
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with alt", func() {
			source := `image::images/foo.png[the foo.png image]`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: "the foo.png image",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with dimensions and id link title meta", func() {
			source := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image, 600, 400]`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
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
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
		It("with roles", func() {
			source := `[.role1.role2]
image::images/foo.png[the foo.png image, 600, 400]`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt:    "the foo.png image",
							types.AttrWidth:       "600",
							types.AttrImageHeight: "400",
							types.AttrRole: []interface{}{
								types.ElementRole{"role1"},
								types.ElementRole{"role2"},
							},
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "images/foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})

	Context("final documents", func() {

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
								types.StringElement{Content: "./path/to/images"},
								types.StringElement{Content: "foo.png"},
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
								types.StringElement{Content: "./path/to/images"},
								types.StringElement{Content: "foo.png"},
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
								types.StringElement{Content: "./path/to/images"},
								types.StringElement{Content: "./path/to/images/foo.png"},
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
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "a paragraph",
								},
							},
							{
								types.StringElement{
									Content: "image::images/foo.png[]",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("paragraph with block image with alt and dimensions", func() {
			source := "a foo image::foo.png[foo image, 600, 400] bar"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a foo image::foo.png[foo image, 600, 400] bar"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})
})

var _ = Describe("inline images", func() {

	Context("draft documents", func() {

		It("with empty alt only", func() {
			source := "image:images/foo.png[]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with empty alt and trailing spaces", func() {
			source := "image:images/foo.png[]  \t\t  "
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
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
								types.StringElement{
									Content: "  \t\t  ",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("inline image surrounded with test", func() {
			source := "a foo image:images/foo.png[] bar..."
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "a foo ",
								},
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
								types.StringElement{
									Content: " bar\u2026\u200b",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with alt alone", func() {
			source := "image:images/foo.png[the foo.png image]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
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
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with alt and width", func() {
			source := "image:images/foo.png[the foo.png image, 600]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
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
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with alt, width and height", func() {
			source := "image:images/foo.png[the foo.png image, 600, 400]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
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
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with alt, but empty width and height", func() {
			source := "image:images/foo.png[the foo.png image, , ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
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
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with single other attribute only", func() {
			source := "image:images/foo.png[id=myid]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.InlineImage{
									Attributes: types.Attributes{
										types.AttrImageAlt: "foo",
										types.AttrID:       "myid",
										types.AttrCustomID: true,
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with multiple other attributes only", func() {
			source := "image:images/foo.png[id=myid, title= mytitle, role = myrole ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.InlineImage{
									Attributes: types.Attributes{
										types.AttrImageAlt: "foo",
										types.AttrID:       "myid",
										types.AttrCustomID: true,
										types.AttrTitle:    "mytitle",
										types.AttrRole:     types.ElementRole{"myrole"},
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("with alt, width, height and other attributes", func() {
			source := "image:images/foo.png[ foo, 600, 400, id=myid, title=mytitle, role=myrole ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.InlineImage{
									Attributes: types.Attributes{
										types.AttrImageAlt:    "foo",
										types.AttrWidth:       "600",
										types.AttrImageHeight: "400",
										types.AttrID:          "myid",
										types.AttrCustomID:    true,
										types.AttrTitle:       "mytitle",
										types.AttrRole:        types.ElementRole{"myrole"},
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("inline image in a paragraph with space after colon", func() {
			source := "this is an image: image:images/foo.png[]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "this is an image: ",
								},
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("inline image in a paragraph without space separator", func() {
			source := "this is an inline.image:images/foo.png[]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "this is an inline.",
								},
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
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

	})

	Context("final documents", func() {

		It("with empty alt only", func() {
			source := "image:images/foo.png[]"
			expected := types.Document{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
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

		It("with document attribute in URL", func() {
			source := `
:dir: ./path/to/images

an image:{dir}/foo.png[].`
			expected := types.Document{
				Attributes: types.Attributes{
					"dir": "./path/to/images",
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
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

		It("with implicit imagesdir document attribute", func() {
			source := `
:imagesdir: ./path/to/images

an image:foo.png[].`
			expected := types.Document{
				Attributes: types.Attributes{
					"imagesdir": "./path/to/images",
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "an "},
								types.InlineImage{
									Attributes: types.Attributes{
										types.AttrImageAlt: "foo",
									},
									Location: types.Location{
										Path: []interface{}{
											types.StringElement{Content: "./path/to/images"},
											types.StringElement{Content: "foo.png"},
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

		It("with explicit duplicate imagesdir document attribute", func() {
			source := `
:imagesdir: ./path/to/images

an image:{imagesdir}/foo.png[].`
			expected := types.Document{
				Attributes: types.Attributes{
					"imagesdir": "./path/to/images",
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "an "},
								types.InlineImage{
									Attributes: types.Attributes{
										types.AttrImageAlt: "foo",
									},
									Location: types.Location{
										Path: []interface{}{
											types.StringElement{Content: "./path/to/images"},
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
						Lines: [][]interface{}{
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
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
