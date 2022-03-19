package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golint
)

var _ = Describe("block images", func() {

	Context("in final documents", func() {

		It("with empty alt", func() {
			source := "image::images/cookie.png[]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ImageBlock{
						Location: &types.Location{
							Path: "images/cookie.png",
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with attribute substitution in alt text", func() {
			source := `:alt: the cookie.png image
			
image::images/cookie.png[{alt}]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "alt",
								Value: "the cookie.png image",
							},
						},
					},
					&types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: "the cookie.png image", // substituted
						},
						Location: &types.Location{
							Path: "images/cookie.png",
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with implicit imagesdir document attribute", func() {
			source := `
:imagesdir: ./path/to/images

image::cookie.png[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "imagesdir",
								Value: "./path/to/images",
							},
						},
					},
					&types.ImageBlock{
						Location: &types.Location{
							Path: "cookie.png", // `imagesdir` (image catalog) attribute not preprended at this point
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with document attribute in URL", func() {
			source := `
:dir: ./path/to/images

image::{dir}/cookie.png[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "dir",
								Value: "./path/to/images",
							},
						},
					},
					&types.ImageBlock{
						Location: &types.Location{
							Path: "./path/to/images/cookie.png",
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with implicit imagesdir", func() {
			source := `
:imagesdir: ./path/to/images

image::cookie.png[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "imagesdir",
								Value: "./path/to/images",
							},
						},
					},
					&types.ImageBlock{
						Location: &types.Location{
							Path: "cookie.png", // `imagesdir` (image catalog) attribute not preprended at this point
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with explicit duplicate imagesdir document attribute", func() {
			source := `
:imagesdir: ./path/to/images

image::{imagesdir}/cookie.png[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "imagesdir",
								Value: "./path/to/images",
							},
						},
					},
					&types.ImageBlock{
						Location: &types.Location{
							Path: "./path/to/images/cookie.png", // `imagesdir` (image catalog) attribute not preprended at this point
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("2 block images", func() {
			source := `image::images/cookie.png[]
image::images/pasta.png[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.ImageBlock{
						Location: &types.Location{
							Path: "images/cookie.png",
						},
					},
					&types.ImageBlock{
						Location: &types.Location{
							Path: "images/pasta.png",
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with role above", func() {
			source := `.mytitle
[#myid]
[.myrole]
image::cookie.png[cookie image, 600, 400]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrID:    "myid",
							types.AttrTitle: "mytitle",
							types.AttrRoles: types.Roles{
								"myrole",
							},
							types.AttrImageAlt: "cookie image",
							types.AttrWidth:    "600",
							types.AttrHeight:   "400",
						},
						Location: &types.Location{
							Path: "cookie.png",
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with link", func() {
			source := "image::cookie.png[cookie image, link=https://cookie.dev]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt:   "cookie image",
							types.AttrInlineLink: "https://cookie.dev",
						},
						Location: &types.Location{
							Path: "cookie.png",
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with special characters", func() {
			source := `image::http://example.com/foo.png?a=1&b=2[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.ImageBlock{
						Location: &types.Location{
							Scheme: "http://",
							Path:   "example.com/foo.png?a=1&b=2",
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		Context("errors", func() {

			It("appending inline content", func() {
				source := "a paragraph\nimage::images/cookie.png[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a paragraph\nimage::images/cookie.png[]",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("paragraph with block image with alt and dimensions", func() {
				source := "a cookie image::cookie.png[cookie image, 600, 400] image"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a cookie image::cookie.png[cookie image, 600, 400] image",
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

var _ = Describe("inline images", func() {

	Context("in final documents", func() {

		It("with empty alt only", func() {
			source := "image:images/cookie.png[]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineImage{
								Location: &types.Location{
									Path: "images/cookie.png",
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

an image:{dir}/cookie.png[].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "dir",
								Value: "./path/to/images",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "an ",
							},
							&types.InlineImage{
								Location: &types.Location{
									Path: "./path/to/images/cookie.png",
								},
							},
							&types.StringElement{
								Content: ".",
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

an image:cookie.png[].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "imagesdir",
								Value: "./path/to/images",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "an ",
							},
							&types.InlineImage{
								Location: &types.Location{
									Path: "cookie.png", // `imagesdir` (image catalog) attribute not preprended at this point
								},
							},
							&types.StringElement{
								Content: ".",
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

an image:{imagesdir}/cookie.png[].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "imagesdir",
								Value: "./path/to/images",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "an ",
							},
							&types.InlineImage{
								Location: &types.Location{
									Path: "./path/to/images/cookie.png", // `imagesdir` (image catalog) attribute not preprended at this point
								},
							},
							&types.StringElement{
								Content: ".",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with link", func() {
			source := "image:cookie.png[cookie image, link=https://cookie.dev]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineImage{
								Attributes: types.Attributes{
									types.AttrImageAlt:   "cookie image",
									types.AttrInlineLink: "https://cookie.dev",
								},
								Location: &types.Location{
									Path: "cookie.png",
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
