package parser

// import (
// 	. "github.com/onsi/ginkgo" //nolint golint
// 	. "github.com/onsi/gomega" //nolint golint
// )

// var _ = Describe("footnotes", func() {

// 	It("should replace footnotes with footnote references", func() {

// 		source := []interface{}{
// 			types.Paragraph{
// 				Lines: [][]interface{}{
// 					{
// 						&types.StringElement{
// 							Content: "A statement.",
// 						},
// 						types.Footnote{
// 							Elements: []interface{}{
// 								&types.StringElement{
// 									Content: "a regular footnote.",
// 								},
// 							},
// 						},
// 					},
// 					{
// 						&types.StringElement{
// 							Content: "A bold statement!",
// 						},
// 						types.Footnote{
// 							Ref: "disclaimer",
// 							Elements: []interface{}{
// 								&types.StringElement{
// 									Content: "Opinions are my own.",
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 			types.Paragraph{
// 				Lines: [][]interface{}{
// 					{
// 						&types.StringElement{
// 							Content: "Another outrageous statement.",
// 						},
// 						types.Footnote{
// 							Ref:      "disclaimer",
// 							Elements: []interface{}{},
// 						},
// 					},
// 				},
// 			},
// 		}
// 		expectedDraftDoc := []interface{}{
// 			types.Paragraph{
// 				Lines: [][]interface{}{
// 					{
// 						&types.StringElement{
// 							Content: "A statement.",
// 						},
// 						types.FootnoteReference{
// 							ID: 1,
// 						},
// 					},
// 					{
// 						&types.StringElement{
// 							Content: "A bold statement!",
// 						},
// 						types.FootnoteReference{
// 							ID:  2,
// 							Ref: "disclaimer",
// 						},
// 					},
// 				},
// 			},
// 			types.Paragraph{
// 				Lines: [][]interface{}{
// 					{
// 						&types.StringElement{
// 							Content: "Another outrageous statement.",
// 						},
// 						types.FootnoteReference{
// 							ID:        2,
// 							Ref:       "disclaimer",
// 							Duplicate: true, // this FootnoteReference targets an already-existing footnote
// 						},
// 					},
// 				},
// 			},
// 		}
// 		expectedFootnotes := []*types.Footnote{
// 			{
// 				ID: 1,
// 				Elements: []interface{}{
// 					&types.StringElement{
// 						Content: "a regular footnote.",
// 					},
// 				},
// 			},
// 			{
// 				ID:  2,
// 				Ref: "disclaimer",
// 				Elements: []interface{}{
// 					&types.StringElement{
// 						Content: "Opinions are my own.",
// 					},
// 				},
// 			},
// 		}
// 		actualDraftDoc, actualFootnotes := processFootnotes(source)
// 		Expect(actualDraftDoc).To(Equal(expectedDraftDoc))
// 		Expect(actualFootnotes).To(Equal(expectedFootnotes))
// 	})

// })
