package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("attributes", func() {

	// We test inline image attributes first.
	Context("at inline level", func() {

		It("block image with empty alt", func() {
			source := `image::foo.png[]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   16,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image with empty alt and extra whitespace", func() {
			source := `image::foo.png[ ]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   17,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image with empty positional parameters", func() {
			source := `image::foo.png[ , , ]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   21,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image with empty first parameter, non-empty width", func() {
			source := `image::foo.png[ , 200, ]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   24,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrWidth: "200",
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image with double quoted alt", func() {
			source := `image::foo.png["Quoted, Here"]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   30,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `Quoted, Here`,
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image with double quoted alt and escaped double quotes", func() {
			source := `image::foo.png["The Foo\"Bar\" here"]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   37,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `The Foo"Bar" here`,
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image with single quoted alt and escaped single quotes", func() {
			source := `image::foo.png['The Foo\'Bar\' here']`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   37,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `The Foo'Bar' here`,
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image with double quoted alt and standalone backslash", func() {
			source := `image::foo.png["The Foo\Bar here"]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   34,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `The Foo\Bar here`,
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image with single quoted alt and standalone backslash", func() {
			source := `image::foo.png['The Foo\Bar here']`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   34,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `The Foo\Bar here`,
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image alt and named pair", func() {
			source := `image::foo.png["Quoted, Here", height=100]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   42,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `Quoted, Here`,
								types.AttrHeight:   "100",
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			result, err := ParseDocumentFragments(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image alt, width, height, and named pair", func() {
			source := `image::foo.png["Quoted, Here", 1, 2, height=100]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   48,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `Quoted, Here`,
								types.AttrHeight:   "100", // named attribute one wins
								types.AttrWidth:    "1",
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image alt, width, height, and named pair (spacing)", func() {
			source := `image::foo.png["Quoted, Here", 1, 2, height=100, test1=123 ,test2 = second test ]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   81,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `Quoted, Here`,
								types.AttrHeight:   "100", // last one wins
								types.AttrWidth:    "1",
								"test1":            "123",
								"test2":            "second test", // shows trailing pad removed
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("block image alt, width, height, and named pair embedded quote", func() {
			source := "image::foo.png[\"Quoted, Here\", 1, 2, height=100, test1=123 ,test2 = second \"test\" ]"
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   83,
					},
					Elements: []interface{}{
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt: `Quoted, Here`,
								types.AttrHeight:   "100", // last one wins
								types.AttrWidth:    "1",
								"test1":            "123",
								"test2":            `second "test"`, // shows trailing pad removed
							},
							Location: &types.Location{
								Path: "foo.png",
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})
	})

	Context("at block level", func() {

		It("section with id attached", func() {
			source := `[id='custom']
[role="cookie"]
== Section A
`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   43,
					},
					Elements: []interface{}{
						&types.Section{
							Level: 1,
							Attributes: types.Attributes{
								types.AttrID:    "custom",
								types.AttrRoles: types.Roles{"cookie"},
							},
							Title: []interface{}{
								types.RawLine("Section A"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("section with attributes detached", func() {
			source := `[id='custom']
		
[role="cookie"]

== Section A
`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   47,
					},
					Elements: []interface{}{
						&types.Section{
							Level: 1,
							Attributes: types.Attributes{
								types.AttrID:    "custom",
								types.AttrRoles: types.Roles{"cookie"},
							},
							Title: []interface{}{
								types.RawLine("Section A"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})
	})

	Context("invalid syntax", func() {

		// no space should be allowed at the beginning of inline attributes,
		// (to be consistent with block attributes)

		It("block image with double quoted alt extra whitespace", func() {
			source := `image::foo.png[ "This \Backslash  2Spaced End Space " ]`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   55,
					},
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine(`image::foo.png[ "This \Backslash  2Spaced End Space " ]`),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
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
	Entry(`[attr1=cookie,attr2='pasta']`, `[attr1=cookie,attr2='pasta']`,
		types.Attributes{
			`attr1`: `cookie`,
			`attr2`: `pasta`,
		},
	),
	Entry(`[attr1=cookie,attr2=pasta]`, `[attr1=cookie,attr2="pasta"]`,
		types.Attributes{
			`attr1`: `cookie`,
			`attr2`: `pasta`,
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
			types.AttrRoles: types.Roles{`a_role`},
		},
	),
	Entry(`[.a_role.another_role]`, `[.a_role.another_role]`,
		types.Attributes{
			types.AttrRoles: types.Roles{`a_role`, `another_role`},
		},
	),
	Entry(`[source.a_role,go]`, `[source.a_role,go]`,
		types.Attributes{
			types.AttrPositional1: `source`,
			types.AttrPositional2: `go`,
			types.AttrRoles:       types.Roles{`a_role`},
		},
	),
	Entry(`[source,go.not_a_role]`, `[source,go.not_a_role]`,
		types.Attributes{
			types.AttrPositional1: `source`,
			types.AttrPositional2: `go.not_a_role`,
		},
	),

	// options (explicit)
	Entry(`[options=header]`, `[options=header]`,
		types.Attributes{
			types.AttrOptions: types.Options{"header"},
		},
	),
	Entry(`[options="header,footer"]`, `[options="header,footer"]`,
		types.Attributes{
			types.AttrOptions: types.Options{"header", "footer"},
		},
	),

	// option shorthand
	Entry(`[%hardbreaks]`, `[%hardbreaks]`,
		types.Attributes{
			types.AttrOptions: types.Options{"hardbreaks"},
		},
	),
	Entry(`[%header]`, `[%header]`,
		types.Attributes{
			types.AttrOptions: types.Options{"header"},
		},
	),
	Entry(`[%footer]`, `[%footer]`,
		types.Attributes{
			types.AttrOptions: types.Options{"footer"},
		},
	),
	Entry(`[%header%footer]`, `[%header%footer]`,
		types.Attributes{
			types.AttrOptions: types.Options{"header", "footer"},
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
			types.AttrRoles: types.Roles{`a_role`},
		},
	),
	Entry(`[#an_id.role_1%option_1.role_2]`, `[#an_id.role_1%option_1.role_2]`,
		types.Attributes{
			types.AttrID:      `an_id`,
			types.AttrRoles:   types.Roles{`role_1`, `role_2`},
			types.AttrOptions: types.Options{`option_1`},
		},
	),
	Entry(`[#an_id.role_1%option_1.role_2%option_2]`, `[#an_id.role_1%option_1.role_2%option_2]`,
		types.Attributes{
			types.AttrID:      `an_id`,
			types.AttrRoles:   types.Roles{`role_1`, `role_2`},
			types.AttrOptions: types.Options{`option_1`, `option_2`},
		},
	),
	Entry(`[#an_id,role="a role"]`, `[#an_id,role="a role"]`,
		types.Attributes{
			types.AttrID:    `an_id`,
			types.AttrRoles: types.Roles{`a role`},
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
			types.AttrPositional1: "*cookie*", // with the `InlineAttributes` rule, values are returned as-is by the parser
			types.AttrPositional2: "_chocolate_",
		},
	),
	// single-quoted positional attributes with quoted text value
	Entry(`[*cookie*,_chocolate_]`, `[*cookie*,_chocolate_]`,
		types.Attributes{
			types.AttrPositional1: "*cookie*", // with the `InlineAttributes` rule, values are returned as-is by the parser
			types.AttrPositional2: "_chocolate_",
		},
	),
	// double-quoted positional attributes with quoted text value
	Entry(`["*cookie*","_chocolate_"]`, `["*cookie*","_chocolate_"]`,
		types.Attributes{
			types.AttrPositional1: "*cookie*", // with the `InlineAttributes` rule, values are returned as-is by the parser
			types.AttrPositional2: "_chocolate_",
		},
	),
)

var _ = DescribeTable("invalid block attributes",

	func(source string) {
		// given
		// do not show parse errors in the logs for this test
		_, reset := ConfigureLogger(log.FatalLevel)
		defer reset()
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
