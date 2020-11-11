package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("attributes", func() {

	// We test inline image attributes first.
	Context("inline attributes", func() {

		It("block image with empty alt", func() {
			source := "image::foo.png[]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image with empty alt and extra whitespace", func() {
			source := "image::foo.png[ ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image with empty positional parameters", func() {
			source := "image::foo.png[ , , ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image with empty first parameter, non-empty width", func() {
			source := "image::foo.png[ , 200, ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrWidth: "200",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image with simple double quoted alt", func() {
			source := "image::foo.png[\"Quoted, Here\"]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: `Quoted, Here`,
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image with double quoted alt and embedded quotes", func() {
			source := `image::foo.png[  "The Ascii\"Doctor\" Is In" ]`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: `The Ascii"Doctor" Is In`,
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image with double quoted alt extra whitespace", func() {
			source := `image::foo.png[ "This \Backslash  2Spaced End Space " ]`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: `This \Backslash  2Spaced End Space `, // trailing space is retained
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image with single quoted alt and embedded quotes", func() {
			source := "image::foo.png[  'It\\'s It!' ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: `It's It!`,
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image with single quoted alt extra whitespace", func() {
			source := "image::foo.png[ 'This \\Backslash  2Spaced End Space ' ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: `This \Backslash  2Spaced End Space `, // trailing space within quotes is retained
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			result, err := ParseDraftDocument(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDraftDocument(expected))
		})

		It("block image alt and named pair", func() {
			source := `image::foo.png["Quoted, Here", height=100]`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt:    `Quoted, Here`,
							types.AttrImageHeight: "100",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			result, err := ParseDraftDocument(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDraftDocument(expected))
		})

		It("block image alt, width, height, and named pair", func() {
			source := "image::foo.png[\"Quoted, Here\", 1, 2, height=100]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt:    `Quoted, Here`,
							types.AttrImageHeight: "100", // last one wins
							types.AttrWidth:       "1",
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image alt, width, height, and named pair (spacing)", func() {
			source := "image::foo.png[\"Quoted, Here\", 1, 2, height=100, test1=123 ,test2 = second test ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt:    `Quoted, Here`,
							types.AttrImageHeight: "100", // last one wins
							types.AttrWidth:       "1",
							"test1":               "123",
							"test2":               "second test", // shows trailing pad removed
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("block image alt, width, height, and named pair embedded quote", func() {
			source := "image::foo.png[\"Quoted, Here\", 1, 2, height=100, test1=123 ,test2 = second \"test\" ]"
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt:    `Quoted, Here`,
							types.AttrImageHeight: "100", // last one wins
							types.AttrWidth:       "1",
							"test1":               "123",
							"test2":               `second "test"`, // shows trailing pad removed
						},
						Location: types.Location{
							Path: []interface{}{
								types.StringElement{Content: "foo.png"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})

	Context("recursive attributes", func() {

		It("should substitute an attribute in another attribute", func() {
			source := `:def: foo
:abc: {def}bar

{abc}`
			expected := types.DraftDocument{
				Attributes: types.Attributes{
					"def": "foo",
					"abc": "foobar", // resolved
				},
				Elements: []interface{}{
					types.AttributeDeclaration{
						Name:  "def",
						Value: "foo",
					},
					types.AttributeDeclaration{
						Name:  "abc",
						Value: "foobar", // resolved
					},
					types.BlankLine{},
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "foobar",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("should not substitute an attribute in another attribute when not defined", func() {
			source := `:abc: {def}bar

{abc}`
			expected := types.DraftDocument{
				Attributes: types.Attributes{
					"abc": "{def}bar", // unresolved
				},
				Elements: []interface{}{
					types.AttributeDeclaration{
						Name:  "abc",
						Value: "{def}bar", // unresolved
					},
					types.BlankLine{},
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "{def}bar",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})
})
