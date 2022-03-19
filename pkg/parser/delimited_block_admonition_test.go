package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("admonition blocks", func() {

	Context("in final documents", func() {

		Context("as delimited blocks", func() {

			It("example block as admonition", func() {
				source := `[NOTE]
====
cookie
====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
							Attributes: types.Attributes{
								types.AttrStyle: types.Note,
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "cookie",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
			It("example block as admonition with multiple lines", func() {
				source := `[NOTE]
====
multiple

paragraphs
====
`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
							Attributes: types.Attributes{
								types.AttrStyle: types.Note,
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "multiple",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "paragraphs",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("as paragraph", func() {

			It("basic admonition", func() {
				source := `[CAUTION]                      
this is an admonition paragraph.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Caution,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "this is an admonition paragraph.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("mixed", func() {

			It("admonition paragraph and admonition block with multiple elements", func() {
				source := `[CAUTION]                      
this is an admonition paragraph.
									
									
[NOTE]                         
.Title                     
====                           
This is an admonition block
								
with another paragraph    
====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Caution,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "this is an admonition paragraph.",
								},
							},
						},
						&types.DelimitedBlock{
							Kind: types.Example,
							Attributes: types.Attributes{
								types.AttrStyle: types.Note,
								types.AttrTitle: "Title",
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "This is an admonition block",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "with another paragraph",
										},
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
