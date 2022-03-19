package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("quoted strings", func() {

	Context("in final documents", func() {

		It("simple single quoted string", func() {
			source := "'`curly was single`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "curly was single",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("interior spaces with single quoted string", func() {
			source := "'` curly was single `'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: " curly was single ",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("interior ending space with single quoted string", func() {
			source := "'`curly was single `'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "curly was single ",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("interior leading space with single quoted string", func() {
			source := "'` curly was single`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: " curly was single",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("apostrophe in single bold word", func() {
			source := "this *cookie`'*"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "this ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "cookie",
									},
									&types.Symbol{
										Name: "`'",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("apostrophes in single bold", func() {
			source := "this *mother's brothers' sisters`'*"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "this ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "mothe",
									},
									&types.Symbol{
										Prefix: "r",
										Name:   "'",
									},
									&types.StringElement{
										Content: "s brothers' sisters",
									},
									&types.Symbol{
										Name: "`'",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("apostrophes in double bold", func() {
			source := "this **mother's brothers' sisters`'**"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "this ",
							},
							&types.QuotedText{
								Kind: types.DoubleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "mothe",
									},
									&types.Symbol{
										Prefix: "r",
										Name:   "'",
									},
									&types.StringElement{
										Content: "s brothers' sisters",
									},
									&types.Symbol{
										Name: "`'",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("bold in single quoted string", func() {
			source := "'`curly *was* single`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "curly ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "was",
									},
								},
							},
							&types.StringElement{
								Content: " single",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("italics in single quoted string", func() {
			source := "'`curly _was_ single`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "curly ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteItalic,
								Elements: []interface{}{
									&types.StringElement{
										Content: "was",
									},
								},
							},
							&types.StringElement{
								Content: " single",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("span in single quoted string", func() {
			source := "'`curly [.strikeout]#was#_is_ single`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "curly ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteMarked,
								Attributes: types.Attributes{
									types.AttrRoles: types.Roles{"strikeout"},
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "was",
									},
								},
							},
							&types.QuotedText{
								Kind: types.SingleQuoteItalic,
								Elements: []interface{}{
									&types.StringElement{
										Content: "is",
									},
								},
							},
							&types.StringElement{
								Content: " single",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly in monospace string", func() {
			source := "'`curly `is` single`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "curly ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteMonospace,
								Elements: []interface{}{
									&types.StringElement{
										Content: "is",
									},
								},
							},
							&types.StringElement{
								Content: " single",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly as monospace string", func() {
			source := "'``curly``'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteMonospace,
								Elements: []interface{}{
									&types.StringElement{
										Content: "curly",
									},
								},
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly with nested double curly", func() {
			source := "'`single\"`double`\"`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "single",
							},
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "double",
							},
							&types.Symbol{
								Name: "`\"",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly in monospace string", func() {
			source := "`'`curly`'`"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteMonospace,
								Elements: []interface{}{
									&types.Symbol{
										Name: "'`",
									},
									&types.StringElement{
										Content: "curly",
									},
									&types.Symbol{
										Name: "`'",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly in italics", func() {
			source := "_'`curly`'_"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteItalic,
								Elements: []interface{}{
									&types.Symbol{
										Name: "'`",
									},
									&types.StringElement{
										Content: "curly",
									},
									&types.Symbol{
										Name: "`'",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly in bold", func() {
			source := "*'`curly`'*"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.Symbol{
										Name: "'`",
									},
									&types.StringElement{
										Content: "curly",
									},
									&types.Symbol{
										Name: "`'",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly in link", func() {
			source := "https://www.example.com/a['`example`']"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineLink{
								Location: &types.Location{
									Scheme: "https://",
									Path:   "www.example.com/a",
								},
								Attributes: types.Attributes{
									types.AttrInlineLinkText: []interface{}{
										&types.Symbol{
											Name: "'`",
										},
										&types.StringElement{
											Content: "example",
										},
										&types.Symbol{
											Name: "`'",
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

		It("single curly in single quoted link attribute", func() {
			source := "https://www.example.com/a['an '`example`'']"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineLink{
								Location: &types.Location{
									Scheme: "https://",
									Path:   "www.example.com/a",
								},
								Attributes: types.Attributes{
									types.AttrInlineLinkText: []interface{}{
										&types.StringElement{
											Content: "an ",
										},
										&types.Symbol{
											Name: "'`",
										},
										&types.StringElement{
											Content: "example",
										},
										&types.Symbol{
											Name: "`'",
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

		It("single curly in double quoted link attribute", func() {
			source := "https://www.example.com/a[\"an '`example`'\"]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineLink{
								Location: &types.Location{
									Scheme: "https://",
									Path:   "www.example.com/a",
								},
								Attributes: types.Attributes{
									types.AttrInlineLinkText: []interface{}{
										&types.StringElement{
											Content: "an ",
										},
										&types.Symbol{
											Name: "'`",
										},
										&types.StringElement{
											Content: "example",
										},
										&types.Symbol{
											Name: "`'",
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

		It("image in curly", func() {
			source := "'`a image:foo.png[]`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "a ",
							},
							&types.InlineImage{
								Location: &types.Location{
									Path: "foo.png",
								},
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("icon in curly", func() {
			source := "'`a icon:note[]`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "a ",
							},
							&types.Icon{
								Class: "note",
							},
							&types.Symbol{
								Name: "`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("simple double quoted string", func() {
			source := "\"`curly was single`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "curly was single",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("interior spaces with double quoted string", func() {
			source := "\"` curly was single `\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: " curly was single ",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("interior ending space with double quoted string", func() {
			source := "\"`curly was single `\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "curly was single ",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("interior leading space with double quoted string", func() {
			source := "\"` curly was single`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: " curly was single",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("bold in double quoted string", func() {
			source := "\"`curly *was* single`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "curly ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "was",
									},
								},
							},
							&types.StringElement{
								Content: " single",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("italics in double quoted string", func() {
			source := "\"`curly _was_ single`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "curly ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteItalic,
								Elements: []interface{}{
									&types.StringElement{
										Content: "was",
									},
								},
							},
							&types.StringElement{
								Content: " single",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("span in double quoted string", func() {
			source := "\"`curly [.strikeout]#was#_is_ single`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "curly ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteMarked,
								Attributes: types.Attributes{
									types.AttrRoles: types.Roles{"strikeout"},
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "was",
									},
								},
							},
							&types.QuotedText{
								Kind: types.SingleQuoteItalic,
								Elements: []interface{}{
									&types.StringElement{
										Content: "is",
									},
								},
							},
							&types.StringElement{
								Content: " single",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("double curly in monospace string", func() {
			source := "\"`curly `is` single`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "curly ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteMonospace,
								Elements: []interface{}{
									&types.StringElement{
										Content: "is",
									},
								},
							},
							&types.StringElement{
								Content: " single",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("double curly as monospace string", func() {
			source := "\"``curly``\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteMonospace,
								Elements: []interface{}{
									&types.StringElement{Content: "curly"},
								},
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("double curly with nested single curly", func() {
			source := "\"`double'`single`'`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "double",
							},
							&types.Symbol{
								Name: "'`",
							},
							&types.StringElement{
								Content: "single",
							},
							&types.Symbol{
								Name: "`'",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("double curly in monospace string", func() {
			source := "`\"`curly`\"`"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteMonospace,
								Elements: []interface{}{
									&types.Symbol{
										Name: "\"`",
									},
									&types.StringElement{
										Content: "curly",
									},
									&types.Symbol{
										Name: "`\"",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("double curly in italics", func() {
			source := "_\"`curly`\"_"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteItalic,
								Elements: []interface{}{
									&types.Symbol{
										Name: "\"`",
									},
									&types.StringElement{
										Content: "curly",
									},
									&types.Symbol{
										Name: "`\"",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("double curly in bold", func() {
			source := "*\"`curly`\"*"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.Symbol{
										Name: "\"`",
									},
									&types.StringElement{
										Content: "curly",
									},
									&types.Symbol{
										Name: "`\"",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		// In a link, the quotes are ambiguous, and we default to assuming they are for enclosing
		// the link text.  Nest them explicitly if this is needed.
		It("double curly in link attribute", func() {
			source := "https://www.example.com/a[\"`example`\"]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineLink{
								Location: &types.Location{
									Scheme: "https://",
									Path:   "www.example.com/a",
								},
								Attributes: types.Attributes{
									types.AttrInlineLinkText: []interface{}{
										&types.Symbol{
											Name: "\"`",
										},
										&types.StringElement{
											Content: "example",
										},
										&types.Symbol{
											Name: "`\"",
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

		// This is the unambiguous form.
		It("double curly in double quoted link attribute", func() {
			source := "https://www.example.com/a['\"`example`\"']"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineLink{
								Location: &types.Location{
									Scheme: "https://",
									Path:   "www.example.com/a",
								},
								Attributes: types.Attributes{
									types.AttrInlineLinkText: []interface{}{
										&types.Symbol{
											Name: "\"`",
										},
										&types.StringElement{
											Content: "example",
										},
										&types.Symbol{
											Name: "`\"",
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

		// This is the unambiguous form.
		It("double curly in double quoted link attribute", func() {
			source := "https://www.example.com/a[\"\"`example`\"\"]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineLink{
								Location: &types.Location{
									Scheme: "https://",
									Path:   "www.example.com/a",
								},
								Attributes: types.Attributes{
									types.AttrInlineLinkText: []interface{}{
										&types.Symbol{
											Name: "\"`",
										},
										&types.StringElement{
											Content: "example",
										},
										&types.Symbol{
											Name: "`\"",
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

		It("image in double curly", func() {
			source := "\"`a image:foo.png[]`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "a ",
							},
							&types.InlineImage{
								Location: &types.Location{
									Path: "foo.png",
								},
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("icon in double curly", func() {
			source := "\"`a icon:note[]`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.Symbol{
								Name: "\"`",
							},
							&types.StringElement{
								Content: "a ",
							},
							&types.Icon{
								Class: "note",
							},
							&types.Symbol{
								Name: "`\"",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly in title", func() {
			source := "== a '`curly`' episode"
			title := []interface{}{
				&types.StringElement{
					Content: "a ",
				},
				&types.Symbol{
					Name: "'`",
				},
				&types.StringElement{
					Content: "curly",
				},
				&types.Symbol{
					Name: "`'",
				},
				&types.StringElement{
					Content: " episode",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_a_curly_episode",
						},
						Title: title,
					},
				},
				ElementReferences: types.ElementReferences{
					"_a_curly_episode": title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_a_curly_episode",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("curly in list element", func() {
			source := "* a '`curly`' episode"
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.UnorderedListKind,
						Elements: []types.ListElement{
							&types.UnorderedListElement{
								CheckStyle:  types.NoCheck,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "a ",
											},
											&types.Symbol{
												Name: "'`",
											},
											&types.StringElement{
												Content: "curly",
											},
											&types.Symbol{
												Name: "`'",
											},
											&types.StringElement{
												Content: " episode",
											},
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

		It("curly in labeled list", func() {
			source := "'`term`':: something '`quoted`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.LabeledListKind,
						Elements: []types.ListElement{
							&types.LabeledListElement{
								Style: types.DoubleColons,
								Term: []interface{}{
									&types.Symbol{
										Name: "'`",
									},
									&types.StringElement{
										Content: "term",
									},
									&types.Symbol{
										Name: "`'",
									},
								},
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "something ",
											},
											&types.Symbol{
												Name: "'`",
											},
											&types.StringElement{
												Content: "quoted",
											},
											&types.Symbol{
												Name: "`'",
											},
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

		It("double curly in title", func() {
			source := "== a \"`curly`\" episode"
			title := []interface{}{
				&types.StringElement{
					Content: "a ",
				},
				&types.Symbol{
					Name: "\"`",
				},
				&types.StringElement{
					Content: "curly",
				},
				&types.Symbol{
					Name: "`\"",
				},
				&types.StringElement{
					Content: " episode",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_a_curly_episode",
						},
						Title: title,
					},
				},
				ElementReferences: types.ElementReferences{
					"_a_curly_episode": title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_a_curly_episode",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("double curly in labeled list", func() {
			source := "\"`term`\":: something \"`quoted`\""
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.LabeledListKind,
						Elements: []types.ListElement{
							&types.LabeledListElement{
								Style: types.DoubleColons,
								Term: []interface{}{
									&types.Symbol{
										Name: "\"`",
									},
									&types.StringElement{
										Content: "term",
									},
									&types.Symbol{
										Name: "`\"",
									},
								},
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "something ",
											},
											&types.Symbol{
												Name: "\"`",
											},
											&types.StringElement{
												Content: "quoted",
											},
											&types.Symbol{
												Name: "`\"",
											},
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

		It("double in list element", func() {
			source := "* a \"`curly`\" episode"
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.UnorderedListKind,
						Elements: []types.ListElement{
							&types.UnorderedListElement{
								CheckStyle:  types.NoCheck,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "a ",
											},
											&types.Symbol{
												Name: "\"`",
											},
											&types.StringElement{
												Content: "curly",
											},
											&types.Symbol{
												Name: "`\"",
											},
											&types.StringElement{
												Content: " episode",
											},
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

		Context("prevented substitution", func() {

			It("apostrophes and escaped quoted strings", func() {
				source := "here`'s an escaped slash \\'`/\\`' and here`'s an escaped backslash \\'`\\\\`'."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "here",
								},
								&types.Symbol{
									Name: "`'",
								},
								&types.StringElement{
									Content: "s an escaped slash '`/`' and here",
								},
								&types.Symbol{
									Name: "`'",
								},
								&types.StringElement{
									Content: "s an escaped backslash '`\\`'.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unescaped single quoted string alone", func() {
				source := "'`hello`'"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.Symbol{
									Name: "'`",
								},
								&types.StringElement{
									Content: `hello`,
								},
								&types.Symbol{
									Name: "`'",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unescaped single quoted strings in a sentence", func() {
				source := "'`quoted content`' and \\'`escaped content\\`'."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.Symbol{
									Name: "'`",
								},
								&types.StringElement{
									Content: "quoted content",
								},
								&types.Symbol{
									Name: "`'",
								},
								&types.StringElement{
									Content: " and '`escaped content`'.",
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
