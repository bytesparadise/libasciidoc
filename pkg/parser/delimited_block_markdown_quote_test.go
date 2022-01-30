package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("markdown-style quote blocks", func() {

	Context("in final documents", func() {

		It("with single marker without author", func() {
			source := `> some text
on *multiple lines*`

			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.MarkdownQuote,
						Elements: []interface{}{
							&types.StringElement{
								Content: "some text\non ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "multiple lines",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with marker on each line without author", func() {
			source := `> some text
> on *multiple lines*`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.MarkdownQuote,
						Elements: []interface{}{
							&types.StringElement{
								Content: "some text\non ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "multiple lines",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with marker on each line with author only", func() {
			source := `> some text
> on *multiple lines*
> -- John Doe`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.MarkdownQuote,
						Attributes: types.Attributes{
							types.AttrQuoteAuthor: "John Doe",
						},
						Elements: []interface{}{
							&types.StringElement{
								Content: "some text\non ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "multiple lines",
									},
								},
							},
							&types.StringElement{
								Content: "\n",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with marker on each line with author and title", func() {
			source := `.title
> some text
> on *multiple lines*
> -- John Doe`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.MarkdownQuote,
						Attributes: types.Attributes{
							types.AttrTitle:       "title",
							types.AttrQuoteAuthor: "John Doe",
						},
						Elements: []interface{}{
							&types.StringElement{
								Content: "some text\non ",
							},
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "multiple lines",
									},
								},
							},
							&types.StringElement{
								Content: "\n",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with with author only", func() {
			source := `> -- John Doe`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.MarkdownQuote,
						Attributes: types.Attributes{
							types.AttrQuoteAuthor: "John Doe",
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
