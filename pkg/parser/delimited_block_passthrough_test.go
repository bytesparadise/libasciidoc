package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golint
)

var _ = Describe("passthrough blocks", func() {

	Context("in raw documents", func() {

		Context("paragraph with attribute", func() {

			It("2-line paragraph followed by another paragraph", func() {
				source := `[pass]
_foo_
*bar*

another paragraph`
				expected := []types.DocumentFragment{
					{
						Position: types.Position{
							Start: 0,
							End:   19,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrStyle: "pass",
								},
								Elements: []interface{}{
									types.RawLine("_foo_\n"),
									types.RawLine("*bar*"),
								},
							},
						},
					},
					{
						Position: types.Position{
							Start: 19,
							End:   20,
						},
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Position: types.Position{
							Start: 20,
							End:   37,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("another paragraph"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
			})
		})

	})

	Context("in final documents", func() {

		Context("as delimited blocks", func() {

			It("with title", func() {
				source := `.a title
++++
_foo_

*bar*
++++`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Passthrough,
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "_foo_\n\n*bar*",
								},
							},
						},
					},
				}
				result, err := ParseDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDocument(expected))
			})

			It("with special characters", func() {
				source := `++++
<input>

<input>
++++`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Passthrough,
							Elements: []interface{}{
								&types.StringElement{
									Content: "<input>\n\n<input>",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with inline link", func() {
				source := `++++
http://example.com[]
++++`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Passthrough,
							Elements: []interface{}{
								&types.StringElement{
									Content: "http://example.com[]",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with inline pass", func() {
				source := `++++
pass:[foo]
++++`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Passthrough,
							Elements: []interface{}{
								&types.StringElement{
									Content: "pass:[foo]",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with variable delimiter length", func() {

				It("with 5 chars", func() {
					source := `+++++
some *passthrough* content
+++++`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DelimitedBlock{
								Kind: types.Passthrough,
								Elements: []interface{}{
									&types.StringElement{
										Content: "some *passthrough* content",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with 5 chars with nested with 4 chars", func() {
					source := `+++++
++++
some *passthrough* content
++++
+++++`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DelimitedBlock{
								Kind: types.Passthrough,
								Elements: []interface{}{
									&types.StringElement{
										Content: "++++\nsome *passthrough* content\n++++",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("paragraph with attribute", func() {

			It("2-line paragraph followed by another paragraph", func() {
				source := `[pass]
_foo_
*bar*

another paragraph`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Passthrough,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "_foo_\n*bar*",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "another paragraph",
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
