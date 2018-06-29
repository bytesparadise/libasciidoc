package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("comments", func() {

	Context("single line comments", func() {

		It("single line comment alone", func() {
			actualDocument := `// A single-line comment.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.SingleLineComment{Content: " A single-line comment."},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualDocument, parser.Entrypoint("DocumentBlock"))
		})

		It("single line comment at end of line", func() {
			actualDocument := `foo // A single-line comment.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "foo // A single-line comment."},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualDocument, parser.Entrypoint("DocumentBlock"))
		})

		It("single line comment within a paragraph", func() {
			actualDocument := `a first line
// A single-line comment.
another line`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
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
			verify(GinkgoT(), expectedResult, actualDocument, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("comment blocks", func() {

		It("comment block alone", func() {
			actualDocument := `//// 
a *comment* block
with multiple lines
////`
			expectedResult := types.DelimitedBlock{
				Kind:       types.CommentBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.StringElement{
						Content: "a *comment* block",
					},
					types.StringElement{
						Content: "with multiple lines",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualDocument, parser.Entrypoint("DocumentBlock"))
		})

		It("comment block with paragraphs around", func() {
			actualDocument := `a first paragraph
//// 
a *comment* block
with multiple lines
////
a second paragraph`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a first paragraph"},
							},
						},
					},
					types.DelimitedBlock{
						Kind:       types.CommentBlock,
						Attributes: map[string]interface{}{},
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
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a second paragraph"},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualDocument)
		})
	})

})
