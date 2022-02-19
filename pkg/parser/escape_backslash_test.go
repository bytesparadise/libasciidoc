package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("escapes with backslashes", func() {

	Context("in final documents", func() {

		It("should escape attribute reference", func() {
			source := `:id: cookie

In /items/\{id}, the id attribute is not replaced.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "id",
								Value: "cookie",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `In /items/{id}, the id attribute is not replaced.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse attribute reference", func() {
			source := `:id: cookie

In /items/{id}, the id attribute is replaced.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "id",
								Value: "cookie",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `In /items/cookie, the id attribute is replaced.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape single quoted bold text", func() {
			source := `\*Content* is not displayed as bold text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `*Content* is not displayed as bold text.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse single quoted bold text", func() {
			source := `*Content* is displayed as bold text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "Content",
									},
								},
							},
							&types.StringElement{
								Content: ` is displayed as bold text.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape single quoted italic text", func() {
			source := `\_Content_ is not displayed as italic text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `_Content_ is not displayed as italic text.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse single quoted italic text", func() {
			source := `_Content_ is displayed as italic text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteItalic,
								Elements: []interface{}{
									&types.StringElement{
										Content: "Content",
									},
								},
							},
							&types.StringElement{
								Content: ` is displayed as italic text.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape single bold text but not italic content", func() {
			source := `\*_Content_* is not displayed as bold text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "*",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteItalic,
								Elements: []interface{}{
									&types.StringElement{
										Content: "Content",
									},
								},
							},
							&types.StringElement{
								Content: `* is not displayed as bold text.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		// TODO: test all combinations of embedded quoted text, with and without attributes
		It("should escape single bold text and italic content", func() {
			source := `\*\_Stars_* is not displayed as bold text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `*_Stars_* is not displayed as bold text.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse single bold text with italic content", func() {
			source := `*_Content_* is displayed as bold and italic text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.QuotedText{
										Kind: types.SingleQuoteItalic,
										Elements: []interface{}{
											&types.StringElement{
												Content: "Content",
											},
										},
									},
								},
							},
							&types.StringElement{
								Content: ` is displayed as bold and italic text.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape single quoted monospace text", func() {
			source := "\\`Content` is not displayed as monospace text"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "`Content` is not displayed as monospace text",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse single quoted monospace text", func() {
			source := "`Content` is displayed as monospace text"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteMonospace,
								Elements: []interface{}{
									&types.StringElement{
										Content: "Content",
									},
								},
							},
							&types.StringElement{
								Content: " is displayed as monospace text",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape attributes on bold text", func() {
			Skip("needs clarification...")
			source := `\[.role]*bold* is displayed as bold text, but without attributes.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "[.role]",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "bold",
									},
								},
							},
							&types.StringElement{
								Content: " is displayed as bold text, but without attributes.",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape attributes and bold text", func() {
			source := `[.role]\*bold* is not displayed as bold text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "[.role]*bold* is not displayed as bold text.",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse attributes and bold text", func() {
			source := `[.role]*bold* is displayed as bold text.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Attributes: types.Attributes{
									types.AttrRoles: types.Roles{"role"},
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "bold",
									},
								},
							},
							&types.StringElement{
								Content: " is displayed as bold text.",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape anchor", func() {
			source := `\[[Word]] is not interpreted as an anchor.
The double brackets around it are preserved.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `[[Word]] is not interpreted as an anchor.
The double brackets around it are preserved.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse anchor", func() {
			source := `[[Word]] is interpreted as an anchor.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.InlineLink{
								Attributes: types.Attributes{
									types.AttrID: "Word",
								},
							},
							&types.StringElement{
								Content: ` is interpreted as an anchor.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape relative link without attributes", func() {
			source := `The URL \link:cookie.adoc[] is not converted into an active link.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `The URL link:cookie.adoc[] is not converted into an active link.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse relative link without attributes", func() {
			source := `The URL link:cookie.adoc[] is converted into an active link.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `The URL `,
							},
							&types.InlineLink{
								Location: &types.Location{
									Path: "cookie.adoc",
								},
							},
							&types.StringElement{
								Content: ` is converted into an active link.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape relative link with attributes", func() {
			source := `The URL \link:cookie.adoc[yummy] is not converted into an active link.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `The URL link:cookie.adoc[yummy] is not converted into an active link.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse relative link with attributes", func() {
			source := `The URL link:cookie.adoc[yummy] is converted into an active link.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `The URL `,
							},
							&types.InlineLink{
								Attributes: types.Attributes{
									types.AttrInlineLinkText: "yummy",
								},
								Location: &types.Location{
									Path: "cookie.adoc",
								},
							},
							&types.StringElement{
								Content: ` is converted into an active link.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape external link without attributes", func() {
			source := `The URL \https://example.org is not converted into an active link.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `The URL https://example.org is not converted into an active link.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse external link without attributes", func() {
			source := `The URL https://example.org is converted into an active link.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `The URL `,
							},
							&types.InlineLink{
								Location: &types.Location{
									Scheme: "https://",
									Path:   "example.org",
								},
							},
							&types.StringElement{
								Content: ` is converted into an active link.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape external link with attributes", func() {
			source := `The URL \https://example.org[example] is not converted into an active link.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `The URL https://example.org[example] is not converted into an active link.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse external link with attributes", func() {
			source := `The URL https://example.org[example] is converted into an active link.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `The URL `,
							},
							&types.InlineLink{
								Attributes: types.Attributes{
									types.AttrInlineLinkText: "example",
								},
								Location: &types.Location{
									Scheme: "https://",
									Path:   "example.org",
								},
							},
							&types.StringElement{
								Content: ` is converted into an active link.`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape apostrophe", func() {
			source := "Here\\`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Here`'",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse apostrophe", func() {
			source := "Here`'"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Here\u2019",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape copyright symbol", func() {
			source := `Copyright \(C)`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Copyright (C)",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse copyright symbol", func() {
			source := `Copyright (C)`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Copyright \u00a9",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape trademark symbol", func() {
			source := `Trademark \(TM)`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Trademark (TM)",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse trademark symbol", func() {
			source := `Trademark (TM)`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Trademark \u2122",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape registered symbol", func() {
			source := `Registered \(R)`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Registered (R)",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse registered symbol", func() {
			source := `Registered (R)`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Registered \u00ae",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape elipsis symbol", func() {
			source := `Elipsis\...`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Elipsis...",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse registered symbol", func() {
			source := `Elipsis...`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Elipsis\u2026\u200b",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should escape typographic quote", func() {
			source := `Here\'s`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Here's",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse typographic quote", func() {
			source := `Here's`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Here\u2019s",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
