package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("process file inclusions", func() {

	It("should include adoc file with leveloffset attribute", func() {
		actualContent := "include::includes/chapter-a.adoc[leveloffset=+1]"
		expectedResult := types.Section{
			Level: 1,
			Title: types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "chapter_a",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{
						Content: "Chapter A",
					},
				},
			},
			Attributes: types.ElementAttributes{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "content",
							},
						},
					},
				},
			},
		}
		verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "= Chapter A",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within listing block", func() {
			actualContent := `----
include::includes/chapter-a.adoc[]
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Listing,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "= Chapter A",
								},
							},
							{},
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within example block", func() {
			actualContent := `====
include::includes/chapter-a.adoc[]
====`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Example,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "= Chapter A",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within quote block", func() {
			actualContent := `____
include::includes/chapter-a.adoc[]
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Quote,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "= Chapter A",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "= Chapter A",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should include adoc file within sidebar block", func() {
			actualContent := `****
include::includes/chapter-a.adoc[]
****`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Sidebar,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "= Chapter A",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "= Chapter A",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("file inclusion with line ranges", func() {

		Context("file inclusion with unquoted line ranges", func() {

			It("file inclusion with single unquoted line", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1]`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with multiple unquoted lines", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1..2]`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with multiple unquoted ranges", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1;3..4;6..-1]`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with invalid unquoted range - case 1", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with invalid unquoted range - case 2", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1;3..4;6..-1]` // using commas instead of semi-colons
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})

		Context("file inclusion with quoted line ranges", func() {

			It("file inclusion with single quoted line", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1"]`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with multiple quoted lines", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1..2"]`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with multiple quoted ranges", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with invalid quoted range - case 1", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with invalid quoted range - case 2", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})
	})

	Context("missing file to include", func() {

		It("should replace with string element if directory does not exist in standalone block", func() {
			actualContent := `include::{unknown}/unknown.adoc[leveloffset=+1]`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]",
						},
					},
				},
			}
			// TODO: also verify that an error was reported in the console.
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should replace with string element if file is missing in standalone block", func() {
			actualContent := `include::includes/unknown.adoc[leveloffset=+1]`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "Unresolved directive in test.adoc - include::includes/unknown.adoc[leveloffset=+1]",
						},
					},
				},
			}
			// TODO: also verify that an error was reported in the console.
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("should replace with string element if file is missing in delimited block", func() {
			actualContent := `----
include::includes/unknown.adoc[leveloffset=+1]
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Listing,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "Unresolved directive in test.adoc - include::includes/unknown.adoc[leveloffset=+1]",
								},
							},
						},
					},
				},
			}
			// TODO: also verify that an error was reported in the console.
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})
})

var _ = Describe("ignore file inclusions", func() {

	It("should include adoc file with leveloffset attribute", func() {
		actualContent := "include::includes/chapter-a.adoc[leveloffset=+1]"
		expectedResult := types.FileInclusion{
			Attributes: types.ElementAttributes{
				types.AttrLevelOffset: "+1",
			},
			Path:    "includes/chapter-a.adoc",
			RawText: `include::includes/chapter-a.adoc[leveloffset=+1]`,
		}
		verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
						RawText:    `include::includes/chapter-a.adoc[]`,
					},
				},
			}
			verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
						RawText:    `include::includes/chapter-a.adoc[]`,
					},
				},
			}
			verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
						RawText:    `include::includes/chapter-a.adoc[]`,
					},
				},
			}
			verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
						RawText:    `include::includes/chapter-a.adoc[]`,
					},
				},
			}
			verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
						RawText:    `include::includes/chapter-a.adoc[]`,
					},
				},
			}
			verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
						RawText:    `include::includes/chapter-a.adoc[]`,
					},
				},
			}
			verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
						RawText:    `include::includes/chapter-a.adoc[]`,
					},
				},
			}
			verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("file inclusion with line ranges", func() {

		Context("file inclusion with unquoted line ranges", func() {

			It("file inclusion with single unquoted line", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1]`
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: types.LineRanges{
							{Start: 1, End: 1},
						},
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines=1]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with multiple unquoted lines", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1..2]`
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: types.LineRanges{
							{Start: 1, End: 2},
						},
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines=1..2]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with multiple unquoted ranges", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1;3..4;6..-1]`
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: types.LineRanges{
							{Start: 1, End: 1},
							{Start: 3, End: 4},
							{Start: 6, End: -1},
						},
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines=1;3..4;6..-1]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with invalid unquoted range - case 1", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: `1;3..4;6..foo`,
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines=1;3..4;6..foo]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with invalid unquoted range - case 2", func() {
				actualContent := `include::includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: types.LineRanges{
							{Start: 1, End: 1},
						},
						"3..4":  nil,
						"6..-1": nil,
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines=1,3..4,6..-1]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})

		Context("file inclusion with quoted line ranges", func() {

			It("file inclusion with single quoted line", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1"]`
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: types.LineRanges{
							{Start: 1, End: 1},
						},
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines="1"]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with multiple quoted lines", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1..2"]`
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: types.LineRanges{
							{Start: 1, End: 2},
						},
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines="1..2"]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with multiple quoted ranges", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: types.LineRanges{
							{Start: 1, End: 1},
							{Start: 3, End: 4},
							{Start: 6, End: -1},
						},
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines="1,3..4,6..-1"]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with invalid quoted range - case 1", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: `"1`, // viewed as a string
						"3..4":               nil,
						"6..foo":             nil,
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines="1,3..4,6..foo"]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("file inclusion with invalid quoted range - case 2", func() {
				actualContent := `include::includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
				expectedResult := types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLineRanges: `"1;3..4;6..10"`,
					},
					Path:    "includes/chapter-a.adoc",
					RawText: `include::includes/chapter-a.adoc[lines="1;3..4;6..10"]`,
				}
				verifyWithoutPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})
	})
})
