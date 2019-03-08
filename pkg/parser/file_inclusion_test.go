package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("file inclusions", func() {

	It("include adoc file with leveloffset attribute", func() {
		actualContent := "include::includes/chapter-a.adoc[leveloffset=+1]"
		expectedResult := types.FileInclusion{
			Attributes: types.ElementAttributes{
				types.AttrLevelOffset: "+1",
			},
			Path: "includes/chapter-a.adoc",
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})
})
