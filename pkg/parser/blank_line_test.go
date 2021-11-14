package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("blank lines", func() {

	Context("in raw documents", func() {

		It("blank line between 2 paragraphs", func() {
			source := `first paragraph
 
second paragraph`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("first paragraph"),
							},
						},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("second paragraph"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("blank line with spaces and tabs between 2 paragraphs and after second paragraph", func() {
			source := `first paragraph
		 

		
second paragraph
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("first paragraph"),
							},
						},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("second paragraph"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("blank line with attributes", func() {
			source := `.ignored
 
`
			expected := []types.DocumentFragment{
				{
					Elements: []interface{}{
						// types.Attributes{
						// 	types.AttrTitle: "ignored",
						// },
						&types.BlankLine{},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})
	})

	Context("in final documents", func() {

		It("blank line between 2 paragraphs", func() {
			source := `first paragraph
 
second paragraph`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "first paragraph"},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "second paragraph"},
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
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "first paragraph"},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "second paragraph"},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
