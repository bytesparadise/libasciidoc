package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"                  //nolint golint
	. "github.com/onsi/ginkgo/extensions/table" //nolint golint
	. "github.com/onsi/gomega"                  //nolint golint
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

		It("block image with double quoted alt", func() {
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
			source := `image::foo.png["The Foo\"Bar\" here"]`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: `The Foo"Bar" here`,
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

		It("block image alt and named pair", func() {
			source := `image::foo.png["Quoted, Here", height=100]`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.ImageBlock{
						Attributes: types.Attributes{
							types.AttrImageAlt: `Quoted, Here`,
							types.AttrHeight:   "100",
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
							types.AttrImageAlt: `Quoted, Here`,
							types.AttrHeight:   "100", // last one wins
							types.AttrWidth:    "1",
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
							types.AttrImageAlt: `Quoted, Here`,
							types.AttrHeight:   "100", // last one wins
							types.AttrWidth:    "1",
							"test1":            "123",
							"test2":            "second test", // shows trailing pad removed
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
							types.AttrImageAlt: `Quoted, Here`,
							types.AttrHeight:   "100", // last one wins
							types.AttrWidth:    "1",
							"test1":            "123",
							"test2":            `second "test"`, // shows trailing pad removed
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

	Context("attributes with substitutions", func() {

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

		Context("invalid syntax", func() {

			// no space should be allowed at the beginning of inline attributes,
			// (to be consistent with block attributes)

			It("block image with double quoted alt extra whitespace", func() {
				source := `image::foo.png[ "This \Backslash  2Spaced End Space " ]`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: `image::foo.png[ "This \Backslash  2Spaced End Space " ]`,
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
})

var _ = DescribeTable("valid block attributes",

	func(source string, expected types.Attributes) {
		// given
		log.Debugf("processing '%s'", source)
		content := strings.NewReader(source + "\n")
		// when parsing only (ie, no substitution applied)
		result, err := parser.ParseReader("", content, parser.Entrypoint("BlockAttributes"))
		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(expected))
	},

	// named attributes
	Entry(`[attr1=cookie]`, `[attr1=cookie]`,
		types.Attributes{
			`attr1`: `cookie`,
		},
	),
	Entry(`[attr1=cookie,foo2='bar2']`, `[attr1=cookie,foo2='bar2']`,
		types.Attributes{
			`attr1`: `cookie`,
			`foo2`:  `bar2`,
		},
	),
	Entry(`[attr1=cookie,foo2=bar2]`, `[attr1=cookie,foo2="bar2"]`,
		types.Attributes{
			`attr1`: `cookie`,
			`foo2`:  `bar2`,
		},
	),

	// positional attributes
	Entry(`[literal]`, `[literal]`,
		types.Attributes{
			types.AttrPositional1: `literal`,
		},
	),
	Entry(`[pass]`, `[pass]`,
		types.Attributes{
			types.AttrPositional1: `pass`,
		},
	),
	Entry(`[example]`, `[example]`,
		types.Attributes{
			types.AttrPositional1: `example`,
		},
	),
	Entry(`[listing]`, `[listing]`,
		types.Attributes{
			types.AttrPositional1: `listing`,
		},
	),
	Entry(`[NOTE]`, `[NOTE]`, // admonitions
		types.Attributes{
			types.AttrPositional1: `NOTE`,
		},
	),
	Entry(`[source,go]`, `[source,go]`,
		types.Attributes{
			types.AttrPositional1: `source`,
			types.AttrPositional2: `go`,
		},
	),
	Entry(`[source,go,foo=bar]`, `[source,go,foo=bar]`,
		types.Attributes{
			types.AttrPositional1: `source`,
			types.AttrPositional2: `go`,
			`foo`:                 `bar`,
		},
	),
	Entry(`[quote,an author,a title]`, `[quote,an author,a title]`,
		types.Attributes{
			types.AttrPositional1: `quote`,
			types.AttrPositional2: `an author`,
			types.AttrPositional3: `a title`,
		},
	),
	Entry(`[verse,an author,a title]`, `[verse,an author,a title]`,
		types.Attributes{
			types.AttrPositional1: `verse`,
			types.AttrPositional2: `an author`,
			types.AttrPositional3: `a title`,
		},
	),
	Entry(`[verse , an author , a title ]`, `[verse, an author , a title ]`, // with spaces around
		types.Attributes{
			types.AttrPositional1: `verse`,
			types.AttrPositional2: `an author`,
			types.AttrPositional3: `a title`,
		},
	),
	Entry(`[verse, ,a title]`, `[verse, ,a title]`, // with empty positional-2
		types.Attributes{
			types.AttrPositional1: `verse`,
			types.AttrPositional2: nil,
			types.AttrPositional3: `a title`,
		},
	),
	Entry(`[verse,,a title]`, `[verse,,a title]`, // with empty positional-2
		types.Attributes{
			types.AttrPositional1: `verse`,
			types.AttrPositional2: nil,
			types.AttrPositional3: `a title`,
		},
	),
	Entry(`[verse,an author,]`, `[verse,an author,]`, // with empty positional-3
		types.Attributes{
			types.AttrPositional1: `verse`,
			types.AttrPositional2: `an author`,
		},
	),
	Entry(`[verse,an author, ]`, `[verse,an author, ]`, // with empty positional-3
		types.Attributes{
			types.AttrPositional1: `verse`,
			types.AttrPositional2: `an author`,
		},
	),
	Entry(`[]`, `[]`, // with empty positional-1
		types.Attributes{},
	),
	Entry(`[ ]`, `[ ]`, // with empty positional-1
		types.Attributes{
			types.AttrPositional1: nil,
		},
	),
	Entry(`[,foo]`, `[,foo]`, // with empty positional-1
		types.Attributes{
			types.AttrPositional1: nil,
			types.AttrPositional2: "foo",
		},
	),
	Entry(`[ ,foo]`, `[ ,foo]`, // with empty positional-1
		types.Attributes{
			types.AttrPositional1: nil,
			types.AttrPositional2: "foo",
		},
	),
	Entry(`[,,]`, `[,,]`, // with empty positional-1
		types.Attributes{
			types.AttrPositional1: nil,
			types.AttrPositional2: nil,
			// types.AttrPositional3: nil,
		},
	),
	// quoted values
	Entry(`.a "title"`, ".a \"title\"",
		types.Attributes{
			types.AttrTitle: `a "title"`,
		},
	),

	// -------------------------
	// shorthand syntaxes
	// -------------------------

	// title shorthand
	Entry(`.a title`, `.a title`,
		types.Attributes{
			types.AttrTitle: `a title`,
		},
	),
	Entry(`.'a title'`, `.'a title'`,
		types.Attributes{
			types.AttrTitle: `'a title'`,
		},
	),
	Entry(`."a title"`, `."a title"`,
		types.Attributes{
			types.AttrTitle: `"a title"`,
		},
	),
	Entry(`.a title.not_a_role`, `.a title.not_a_role`,
		types.Attributes{
			types.AttrTitle: `a title.not_a_role`,
		},
	),

	// role shorthand
	Entry(`[.a_role]`, `[.a_role]`,
		types.Attributes{
			types.AttrRoles: []interface{}{`a_role`},
		},
	),
	Entry(`[.a_role.another_role]`, `[.a_role.another_role]`,
		types.Attributes{
			types.AttrRoles: []interface{}{`a_role`, `another_role`},
		},
	),
	Entry(`[source.a_role,go]`, `[source.a_role,go]`,
		types.Attributes{
			types.AttrPositional1: `source`,
			types.AttrPositional2: `go`,
			types.AttrRoles:       []interface{}{`a_role`},
		},
	),
	Entry(`[source,go.not_a_role]`, `[source,go.not_a_role]`,
		types.Attributes{
			types.AttrPositional1: `source`,
			types.AttrPositional2: `go.not_a_role`,
		},
	),
	Entry("multiple roles", "[.role1]\n[.role2]",
		types.Attributes{
			types.AttrRoles: []interface{}{`role1`, `role2`},
		},
	),

	// option shorthand
	Entry(`[%hardbreaks]`, `[%hardbreaks]`,
		types.Attributes{
			types.AttrOptions: []interface{}{"hardbreaks"},
		},
	),

	// id (alone)
	Entry(`[#an_id]`, `[#an_id]`,
		types.Attributes{
			types.AttrID: `an_id`,
		},
	),
	// id (with roles and options)
	Entry(`[#an_id.a_role]`, `[#an_id.a_role]`,
		types.Attributes{
			types.AttrID:    `an_id`,
			types.AttrRoles: []interface{}{`a_role`},
		},
	),
	Entry(`[#an_id.role_1%option_1.role_2]`, `[#an_id.role_1%option_1.role_2]`,
		types.Attributes{
			types.AttrID:      `an_id`,
			types.AttrRoles:   []interface{}{`role_1`, `role_2`},
			types.AttrOptions: []interface{}{`option_1`},
		},
	),
	Entry(`[#an_id.role_1%option_1.role_2%option_2]`, `[#an_id.role_1%option_1.role_2%option_2]`,
		types.Attributes{
			types.AttrID:      `an_id`,
			types.AttrRoles:   []interface{}{`role_1`, `role_2`},
			types.AttrOptions: []interface{}{`option_1`, `option_2`},
		},
	),
	Entry(`[#an_id,role="a role"]`, `[#an_id,role="a role"]`,
		types.Attributes{
			types.AttrID:    `an_id`,
			types.AttrRoles: []interface{}{`a role`},
		},
	),
	Entry(`[qanda#quiz]`, `[qanda#quiz]`,
		types.Attributes{
			types.AttrPositional1: "qanda",
			types.AttrID:          `quiz`,
		},
	),

	Entry(`[[here, an id]]`, `[[here, an id]]`,
		types.Attributes{
			types.AttrID: `here, an id`,
		},
	),
	Entry(`[[another id.not_a_role]]`, `[[another id.not_a_role]]`,
		types.Attributes{
			types.AttrID: `another id.not_a_role`,
		},
	),

	// TODO: attributes with substitutions
)

var _ = DescribeTable("valid inline attributes",

	func(source string, expected types.Attributes) {
		// given
		log.Debugf("processing '%s'", source)
		content := strings.NewReader(source + "\n")
		// when parsing only (ie, no substitution applied)
		result, err := parser.ParseReader("", content, parser.Entrypoint("InlineAttributes"))
		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(expected))
	},

	// ---------------------
	// named attributes
	// ---------------------
	// named attributes with plain text value
	Entry(`[attr1=cookie]`, `[attr1=cookie]`,
		types.Attributes{
			`attr1`: `cookie`,
		},
	),
	Entry(`[attr1=cookie,attr2=chocolate]`, `[attr1=cookie,attr2=chocolate]`,
		types.Attributes{
			`attr1`: `cookie`,
			"attr2": "chocolate",
		},
	),
	// named attributes with single quoted values
	Entry(`[attr1='cookie',attr2='chocolate']`, `[attr1='cookie',attr2='chocolate']`,
		types.Attributes{
			`attr1`: `cookie`,
			"attr2": "chocolate",
		},
	),
	// named attributes with double quoted values
	Entry(`[attr1="cookie",attr2="chocolate"]`, `[attr1="cookie",attr2="chocolate"]`,
		types.Attributes{
			`attr1`: `cookie`,
			"attr2": "chocolate",
		},
	),
	// ---------------------
	// positional attributes
	// ---------------------
	// unquoted positional attributes with plain text value
	Entry(`[cookie,chocolate]`, `[cookie,chocolate]`,
		types.Attributes{
			types.AttrPositional1: "cookie",
			types.AttrPositional2: "chocolate",
		},
	),
	// unquoted positional attributes with quoted text value
	Entry(`[*cookie*,_chocolate_]`, `[*cookie*,_chocolate_]`,
		types.Attributes{
			types.AttrPositional1: []interface{}{
				types.QuotedText{
					Kind: types.SingleQuoteBold,
					Elements: []interface{}{
						types.StringElement{
							Content: "cookie",
						},
					},
				},
			},
			types.AttrPositional2: []interface{}{
				types.QuotedText{
					Kind: types.SingleQuoteItalic,
					Elements: []interface{}{
						types.StringElement{
							Content: "chocolate",
						},
					},
				},
			},
		},
	),
	// single-quoted positional attributes with quoted text value
	Entry(`[*cookie*,_chocolate_]`, `[*cookie*,_chocolate_]`,
		types.Attributes{
			types.AttrPositional1: []interface{}{
				types.QuotedText{
					Kind: types.SingleQuoteBold,
					Elements: []interface{}{
						types.StringElement{
							Content: "cookie",
						},
					},
				},
			},
			types.AttrPositional2: []interface{}{
				types.QuotedText{
					Kind: types.SingleQuoteItalic,
					Elements: []interface{}{
						types.StringElement{
							Content: "chocolate",
						},
					},
				},
			},
		},
	),
	// double-quoted positional attributes with quoted text value
	Entry(`["*cookie*","_chocolate_"]`, `["*cookie*","_chocolate_"]`,
		types.Attributes{
			types.AttrPositional1: []interface{}{
				types.QuotedText{
					Kind: types.SingleQuoteBold,
					Elements: []interface{}{
						types.StringElement{
							Content: "cookie",
						},
					},
				},
			},
			types.AttrPositional2: []interface{}{
				types.QuotedText{
					Kind: types.SingleQuoteItalic,
					Elements: []interface{}{
						types.StringElement{
							Content: "chocolate",
						},
					},
				},
			},
		},
	),
)

var _ = DescribeTable("invalid block attributes",

	func(source string) {
		// given
		content := strings.NewReader(source + "\n")
		// when parsing only (ie, no substitution applied)
		_, err := parser.ParseReader("", content, parser.Entrypoint("BlockAttributes"))
		// then
		Expect(err).To(HaveOccurred())
	},

	// space after `[` is not allowed if more content exists
	Entry(`[ attr1=cookie]`, `[ attr1=cookie]`),

	Entry(`[ attr1=cookie ]`, `[ attr1=cookie ]`),
)
