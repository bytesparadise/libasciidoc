package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("blank lines", func() {

	Context("draft documents", func() {

		It("blank line between 2 paragraphs", func() {
			source := `first paragraph
 
second paragraph`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "first paragraph",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "second paragraph",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("blank line with spaces and tabs between 2 paragraphs and after second paragraph", func() {
			source := `first paragraph
		 

		
second paragraph
`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "first paragraph",
								},
							},
						},
					},
					types.BlankLine{},
					types.BlankLine{},
					types.BlankLine{},
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "second paragraph",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("blank line with attributes", func() {
			source := `.ignored
 
`
			expected := types.DraftDocument{
				Elements: []interface{}{
					types.BlankLine{},
				},
			}
			result, err := ParseDraftDocument(source) // , parser.Debug(true))
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDraftDocument(expected))
		})
	})

	Context("final documents", func() {

		It("blank line between 2 paragraphs", func() {
			source := `first paragraph
 
second paragraph`
			expected := types.Document{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "first paragraph"},
							},
						},
					},
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "second paragraph"},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
		It("blank line with spaces and tabs between 2 paragraphs and after second paragraph", func() {
			source := `first paragraph
		 

		
second paragraph
`
			expected := types.Document{
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "first paragraph"},
							},
						},
					},
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "second paragraph"},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
