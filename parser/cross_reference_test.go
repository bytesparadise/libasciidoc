package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("cross References", func() {

	Context("reference to section", func() {

		Context("valid reference", func() {

			It("xref with custom id", func() {
				actualContent := `a link to <<thetitle>>.`
				expectedResult := types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a link to "},
						types.CrossReference{ID: "thetitle"},
						types.StringElement{Content: "."},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
			})
		})
	})
})
