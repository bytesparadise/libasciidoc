package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("file inclusions", func() {

	It("should include adoc file with leveloffset attribute", func() {
		actualContent := "include::includes/chapter-a.adoc[leveloffset=+1]"
		expectedResult := types.FileInclusion{
			Attributes: types.ElementAttributes{
				types.AttrLevelOffset: "+1",
			},
			Path: "includes/chapter-a.adoc",
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	Context("file inclusion in delimited blocks", func() {

		It("should include adoc file within fenced block", func() {
			actualContent := "```\n" +
				"include::includes/chapter-a.adoc[]\n" +
				"```"
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Fenced,
				Elements: []interface{}{
					types.FileInclusion{
						Attributes: types.ElementAttributes{},
						Path:       "includes/chapter-a.adoc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within listing block", func() {
			actualContent := `----
include::includes/chapter-a.adoc[]
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Listing,
				Elements: []interface{}{
					types.FileInclusion{
						Attributes: types.ElementAttributes{},
						Path:       "includes/chapter-a.adoc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within example block", func() {
			actualContent := `====
include::includes/chapter-a.adoc[]
====`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Example,
				Elements: []interface{}{
					types.FileInclusion{
						Attributes: types.ElementAttributes{},
						Path:       "includes/chapter-a.adoc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within quote block", func() {
			actualContent := `____
include::includes/chapter-a.adoc[]
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Quote,
				Elements: []interface{}{
					types.FileInclusion{
						Attributes: types.ElementAttributes{},
						Path:       "includes/chapter-a.adoc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within verse block", func() {
			actualContent := `[verse]
____
include::includes/chapter-a.adoc[]
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Verse,
				Elements: []interface{}{
					types.FileInclusion{
						Attributes: types.ElementAttributes{},
						Path:       "includes/chapter-a.adoc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within sidebar block", func() {
			actualContent := `****
include::includes/chapter-a.adoc[]
****`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Sidebar,
				Elements: []interface{}{
					types.FileInclusion{
						Attributes: types.ElementAttributes{},
						Path:       "includes/chapter-a.adoc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within passthrough block", func() {
			Skip("missing support for passthrough blocks")
			actualContent := `++++
include::includes/chapter-a.adoc[]
++++`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				// Kind:       types.Passthrough,
				Elements: []interface{}{
					types.FileInclusion{
						Attributes: types.ElementAttributes{},
						Path:       "includes/chapter-a.adoc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})
})
