package parser_test 

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("comments - preflight", func() {

	Context("single line comments", func() {

		It("single line comment alone", func() {
			doc := `// A single-line comment.`
			expected := &types.SingleLineComment{
				Content: " A single-line comment.",
			}
			verifyDocumentBlock(expected, doc)
		})

		It("single line comment with prefixing spaces alone", func() {
			doc := `  // A single-line comment.`
			expected := &types.SingleLineComment{
				Content: " A single-line comment.",
			}
			verifyDocumentBlock(expected, doc)
		})

		It("single line comment with prefixing tabs alone", func() {
			doc := "\t\t// A single-line comment."
			expected := &types.SingleLineComment{
				Content: " A single-line comment.",
			}
			verifyDocumentBlock(expected, doc)
		})

		It("single line comment at end of line", func() {
			doc := `foo // A single-line comment.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "foo // A single-line comment."},
					},
				},
			}
			verifyDocumentBlock(expected, doc)
		})

		It("single line comment within a paragraph", func() {
			doc := `a first line
// A single-line comment.
another line`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "a first line"},
					},
					{
						&types.SingleLineComment{Content: " A single-line comment."},
					},
					{
						&types.StringElement{Content: "another line"},
					},
				},
			}
			verifyDocumentBlock(expected, doc)
		})

		It("single line comment within a paragraph with tab", func() {
			doc := `a first line
	// A single-line comment.
another line`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "a first line"},
					},
					{
						&types.SingleLineComment{Content: " A single-line comment."},
					},
					{
						&types.StringElement{Content: "another line"},
					},
				},
			}
			verifyDocumentBlock(expected, doc)
		})
	})

	Context("comment blocks", func() {

		It("comment block alone", func() {
			doc := `//// 
a *comment* block
with multiple lines
////`
			expected := &types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Comment,
				Elements: []interface{}{
					&types.StringElement{
						Content: "a *comment* block",
					},
					&types.StringElement{
						Content: "with multiple lines",
					},
				},
			}
			verifyDocumentBlock(expected, doc)
		})

		It("comment block with paragraphs around", func() {
			doc := `a first paragraph
//// 
a *comment* block
with multiple lines
////
a second paragraph`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{Content: "a first paragraph"},
							},
						},
					},
					&types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Comment,
						Elements: []interface{}{
							&types.StringElement{
								Content: "a *comment* block",
							},
							&types.StringElement{
								Content: "with multiple lines",
							},
						},
					},
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{Content: "a second paragraph"},
							},
						},
					},
				},
			}
			verifyPreflight(expected, doc)
		})
	})

})
