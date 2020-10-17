package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("delimited block custom substitutions", func() {

	// testing custom substitutions on example blocks only, as
	// other verbatim blocks (fenced, literal, source, passthrough)
	// share the same implementation

	// also, see https://asciidoctor.org/docs/user-manual/#incremental-substitutions
	// "When you set the subs attribute on a block, you automatically remove all of its default substitutions.
	// For example, if you set subs on a literal block, and assign it a value of attributes,
	// only attributes are substituted."

	source := `:github-url: https://github.com
				
[subs="$SUBS"]
====
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* a list item
====

<1> a callout
`

	It("should apply the default substitution", func() {
		s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]\n", "")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "1",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
								{
									types.StringElement{
										Content: "and ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "more text",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " on the",
									},
									types.LineBreak{},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'normal' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "normal")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "normal",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " ",
									},
									types.SpecialCharacter{ // callout is not detected with the `normal` susbtitution
										Name: "<",
									},
									types.StringElement{
										Content: "1",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
								{
									types.StringElement{
										Content: "and ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "more text",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " on the",
									},
									types.LineBreak{},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'quotes' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "quotes")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "quotes",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to {github-url}[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'macros' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "macros")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "macros",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'attributes' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "attributes")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "attributes",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to https://github.com[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'attributes,macros' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "attributes,macros",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'specialchars' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "specialchars")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "specialchars",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "1",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
								{
									types.StringElement{
										Content: "and ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "more text",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'replacements' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "replacements")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "replacements",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'post_replacements' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "post_replacements")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "post_replacements",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the",
									},
									types.LineBreak{},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'quotes,macros' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "quotes,macros",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to {github-url}[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'macros,quotes' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "macros,quotes",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to {github-url}[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})

	It("should apply the 'none' substitution", func() {
		s := strings.ReplaceAll(source, "$SUBS", "none")
		expected := types.DraftDocument{
			Attributes: types.Attributes{
				"github-url": "https://github.com",
			},
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "github-url",
					Value: "https://github.com",
				},
				types.BlankLine{},
				types.ExampleBlock{
					Attributes: types.Attributes{
						types.AttrSubstitutions: "none",
					},
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
							},
						},
						types.BlankLine{},
						types.UnorderedListItem{
							Level:       1,
							BulletStyle: types.OneAsterisk,
							CheckStyle:  types.NoCheck,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a list item",
											},
										},
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.CalloutListItem{
					Ref: 1,
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a callout",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
	})
})
