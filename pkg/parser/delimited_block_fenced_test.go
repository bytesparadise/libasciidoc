package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("fenced blocks", func() {

	Context("in final documents", func() {

		It("with single line", func() {
			content := "some fenced code"
			source := "```\n" + content + "\n" + "```"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: content,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with no line", func() {
			source := "```\n```"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with multiple lines alone", func() {
			source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: "some fenced code\nwith an empty line\n\nin the middle",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with multiple lines then a paragraph", func() {
			source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: "some fenced code\nwith an empty line\n\nin the middle",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "then a normal paragraph."},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("after a paragraph", func() {
			content := "some fenced code"
			source := "a paragraph.\n\n```\n" + content + "\n" + "```\n"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "a paragraph.",
							},
						},
					},
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: content,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with unclosed delimiter", func() {
			source := "```\nEnd of file here"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: "End of file here",
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
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: "a https://example.com\nand more text on the\nnext lines",
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
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: "a https://example.com[]\nand more text on the\nnext lines",
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
				"* fenced \n" +
				"* content \n```"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: "* some\n* fenced\n* content", // suffix spaces are trimmed
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		Context("with variable delimiter length", func() {

			It("with 5 chars only", func() {
				source := "`````\n" +
					"some *fenced* content\n" +
					"`````"
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								&types.StringElement{
									Content: "some *fenced* content",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with 5 chars with nested with 4 chars", func() {
				source := "`````\n" +
					"````\n" +
					"some *fenced* content\n" +
					"````\n" +
					"`````"
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								&types.StringElement{
									Content: "````\nsome *fenced* content\n````",
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
