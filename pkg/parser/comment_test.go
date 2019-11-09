package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("comments - draft", func() {

	Context("single line comments", func() {

		It("single line comment alone", func() {
			source := `// A single-line comment.`
			expected := types.SingleLineComment{
				Content: " A single-line comment.",
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single line comment with prefixing spaces alone", func() {
			source := `  // A single-line comment.`
			expected := types.SingleLineComment{
				Content: " A single-line comment.",
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single line comment with prefixing tabs alone", func() {
			source := "\t\t// A single-line comment."
			expected := types.SingleLineComment{
				Content: " A single-line comment.",
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single line comment at end of line", func() {
			source := `foo // A single-line comment.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "foo // A single-line comment."},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single line comment within a paragraph", func() {
			source := `a first line
// A single-line comment.
another line`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a first line"},
					},
					{
						types.SingleLineComment{Content: " A single-line comment."},
					},
					{
						types.StringElement{Content: "another line"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single line comment within a paragraph with tab", func() {
			source := `a first line
	// A single-line comment.
another line`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a first line"},
					},
					{
						types.SingleLineComment{Content: " A single-line comment."},
					},
					{
						types.StringElement{Content: "another line"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})
	})

	Context("comment blocks", func() {

		It("comment block alone", func() {
			source := `//// 
a *comment* block
with multiple lines
////`
			expected := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Comment,
				Elements: []interface{}{
					types.StringElement{
						Content: "a *comment* block",
					},
					types.StringElement{
						Content: "with multiple lines",
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("comment block with paragraphs around", func() {
			source := `a first paragraph
//// 
a *comment* block
with multiple lines
////
a second paragraph`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a first paragraph"},
							},
						},
					},
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Comment,
						Elements: []interface{}{
							types.StringElement{
								Content: "a *comment* block",
							},
							types.StringElement{
								Content: "with multiple lines",
							},
						},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a second paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})

})
