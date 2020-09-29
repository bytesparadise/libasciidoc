package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("special characters", func() {

	It("should parse in paragraph", func() {
		source := "<b>*</b> &apos; &amp;"
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Lines: []interface{}{
						[]interface{}{
							types.SpecialCharacter{
								Name: "<",
							},
							types.StringElement{
								Content: "b",
							},
							types.SpecialCharacter{
								Name: ">",
							},
							types.StringElement{
								Content: "*",
							},
							types.SpecialCharacter{
								Name: "<",
							},
							types.StringElement{
								Content: "/b",
							},
							types.SpecialCharacter{
								Name: ">",
							},
							types.StringElement{
								Content: " ",
							},
							types.SpecialCharacter{
								Name: "&",
							},
							types.StringElement{
								Content: "apos; ",
							},
							types.SpecialCharacter{
								Name: "&",
							},
							types.StringElement{
								Content: "amp;",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("should parse in delimited block", func() {
		source := "```" + "\n" +
			"<b>*</b> &apos; &amp;" + "\n" +
			"```"
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.DelimitedBlock{
					Kind: types.Fenced,
					Elements: []interface{}{
						[]interface{}{
							types.SpecialCharacter{
								Name: "<",
							},
							types.StringElement{
								Content: "b",
							},
							types.SpecialCharacter{
								Name: ">",
							},
							types.StringElement{
								Content: "*",
							},
							types.SpecialCharacter{
								Name: "<",
							},
							types.StringElement{
								Content: "/b",
							},
							types.SpecialCharacter{
								Name: ">",
							},
							types.StringElement{
								Content: " ",
							},
							types.SpecialCharacter{
								Name: "&",
							},
							types.StringElement{
								Content: "apos; ",
							},
							types.SpecialCharacter{
								Name: "&",
							},
							types.StringElement{
								Content: "amp;",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})
})
