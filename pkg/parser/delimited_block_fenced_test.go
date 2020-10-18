package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("fenced blocks", func() {

	Context("draft documents", func() {

		Context("delimited blocks", func() {

			It("with single line", func() {
				content := "some fenced code"
				source := "```\n" + content + "\n" + "```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: content,
									},
								},
							},
						},
					},
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
			})

			It("with special characters line", func() {
				source := "```\n<some fenced code>\n" + "```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "some fenced code",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
							},
						},
					},
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
			})

			It("with no line", func() {
				source := "```\n```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some fenced code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some fenced code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{}, // empty line
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "then a normal paragraph.",
									},
								},
							},
						},
					},
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
			})

			It("after a paragraph", func() {
				content := "some fenced code"
				source := "a paragraph.\n\n```\n" + content + "\n" + "```\n"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.BlankLine{},
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: content,
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := "```\nEnd of file here"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with external link inside - without attributes", func() {
				source := "```" + "\n" +
					"a https://example.com\n" +
					"and more text on the\n" +
					"next lines\n" +
					"```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a https://example.com",
									},
								},
								{
									types.StringElement{
										Content: "and more text on the",
									},
								},
								{
									types.StringElement{
										Content: "next lines",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with external link inside - with attributes", func() {
				source := "```" + "\n" +
					"a https://example.com[]" + "\n" +
					"and more text on the" + "\n" +
					"next lines" + "\n" +
					"```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a https://example.com[]",
									},
								},
								{
									types.StringElement{
										Content: "and more text on the",
									},
								},
								{
									types.StringElement{
										Content: "next lines",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unrendered list", func() {
				source := "```\n" +
					"* some \n" +
					"* listing \n" +
					"* content \n```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "* some ",
									},
								},
								{
									types.StringElement{
										Content: "* listing ",
									},
								},
								{
									types.StringElement{
										Content: "* content ",
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

	Context("final documents", func() {

		Context("delimited blocks", func() {

			It("with single line", func() {
				content := "some fenced code"
				source := "```\n" + content + "\n" + "```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: content,
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with no line", func() {
				source := "```\n```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some fenced code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some fenced code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "then a normal paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("after a paragraph", func() {
				content := "some fenced code"
				source := "a paragraph.\n\n```\n" + content + "\n" + "```\n"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: content,
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := "```\nEnd of file here"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with external link inside - without attributes", func() {
				source := "```\n" +
					"a https://example.com\n" +
					"and more text on the\n" +
					"next lines\n" +
					"```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a https://example.com",
									},
								},
								{
									types.StringElement{
										Content: "and more text on the",
									},
								},
								{
									types.StringElement{
										Content: "next lines",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with external link inside - with attributes", func() {
				source := "```" + "\n" +
					"a https://example.com[]" + "\n" +
					"and more text on the" + "\n" +
					"next lines" + "\n" +
					"```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a https://example.com[]",
									},
								},
								{
									types.StringElement{
										Content: "and more text on the",
									},
								},
								{
									types.StringElement{
										Content: "next lines",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unrendered list", func() {
				source := "```\n" +
					"* some \n" +
					"* listing \n" +
					"* content \n```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "* some ",
									},
								},
								{
									types.StringElement{
										Content: "* listing ",
									},
								},
								{
									types.StringElement{
										Content: "* content ",
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
})
