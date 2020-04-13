package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

var _ = DescribeTable("'FileLocation' pattern",
	func(filename string, expected interface{}) {
		reader := strings.NewReader(filename)
		actual, err := parser.ParseReader(filename, reader, parser.Entrypoint("FileLocation"))
		Expect(err).ToNot(HaveOccurred())
		GinkgoT().Log("actual result: %s", spew.Sdump(actual))
		GinkgoT().Log("expected result: %s", spew.Sdump(expected))
		Expect(actual).To(Equal(expected))
	},
	Entry("'chapter'", "chapter", types.Location{
		Elements: []interface{}{
			types.StringElement{
				Content: "chapter",
			},
		},
	}),
	Entry("'chapter.adoc'", "chapter.adoc", types.Location{
		Elements: []interface{}{
			types.StringElement{
				Content: "chapter.adoc",
			},
		},
	}),
	Entry("'chapter-a.adoc'", "chapter-a.adoc", types.Location{
		Elements: []interface{}{
			types.StringElement{
				Content: "chapter-a.adoc",
			},
		},
	}),
	Entry("'chapter_a.adoc'", "chapter_a.adoc", types.Location{
		Elements: []interface{}{
			types.StringElement{
				Content: "chapter_a.adoc",
			},
		},
	}),
	Entry("'../../test/includes/chapter_a.adoc'", "../../test/includes/chapter_a.adoc", types.Location{
		Elements: []interface{}{
			types.StringElement{
				Content: "../../test/includes/chapter_a.adoc",
			},
		},
	}),
	Entry("'chapter-{foo}.adoc'", "chapter-{foo}.adoc", types.Location{
		Elements: []interface{}{
			types.StringElement{
				Content: "chapter-",
			},
			types.DocumentAttributeSubstitution{
				Name: "foo",
			},
			types.StringElement{
				Content: ".adoc",
			},
		},
	}),
	Entry("'{includedir}/chapter-{foo}.adoc'", "{includedir}/chapter-{foo}.adoc", types.Location{
		Elements: []interface{}{
			types.DocumentAttributeSubstitution{
				Name: "includedir",
			},
			types.StringElement{
				Content: "/chapter-",
			},
			types.DocumentAttributeSubstitution{
				Name: "foo",
			},
			types.StringElement{
				Content: ".adoc",
			},
		},
	}),
)

var _ = DescribeTable("check asciidoc file",
	func(path string, expectation bool) {
		Expect(parser.IsAsciidoc(path)).To(Equal(expectation))
	},
	Entry("foo.adoc", "foo.adoc", true),
	Entry("foo.asc", "foo.asc", true),
	Entry("foo.ad", "foo.ad", true),
	Entry("foo.asciidoc", "foo.asciidoc", true),
	Entry("foo.txt", "foo.txt", true),
	Entry("foo.csv", "foo.csv", false),
	Entry("foo.go", "foo.go", false),
)

var _ = Describe("file inclusions - draft with preprocessing", func() {

	It("should include adoc file without leveloffset from local file", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../test/includes/chapter-a.adoc[]"
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      0,
					Title: []interface{}{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "content",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source, WithFilename("foo.adoc"))).To(Equal(expected))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file without leveloffset from relative file", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../../test/includes/chapter-a.adoc[]"
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      0,
					Title: []interface{}{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "content",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source, WithFilename("tmp/foo.adoc"))).To(Equal(expected))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file with leveloffset", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../test/includes/chapter-a.adoc[leveloffset=+1]"
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Title: []interface{}{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "content",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include section 0 by default", func() {
		source := "include::../../test/includes/chapter-a.adoc[]"
		// at this level (parsing), it is expected that the Section 0 is part of the Prefligh document
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      0,
					Title: []interface{}{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "content",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("should not include section 0 when attribute exists", func() {
		source := `:includedir: ../../test/includes

include::{includedir}/chapter-a.adoc[]`
		// at this level (parsing), it is expected that the Section 0 is part of the Prefligh document
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.DocumentAttributeDeclaration{
					Name:  "includedir",
					Value: "../../test/includes",
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      0,
					Title: []interface{}{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "content",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("should not further process with non-asciidoc files", func() {
		source := `:includedir: ../../test/includes

include::{includedir}/include.foo[]`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.DocumentAttributeDeclaration{
					Name:  "includedir",
					Value: "../../test/includes",
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.QuotedText{
								Kind: types.Bold,
								Elements: []interface{}{
									types.StringElement{
										Content: "some strong content",
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "include::hello_world.go.txt[]",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source, WithFilename("foo.bar"))).To(Equal(expected)) // parent doc may not need to be a '.adoc'
	})

	It("should include grandchild content without offset", func() {
		source := `include::../../test/includes/grandchild-include.adoc[]`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Title: []interface{}{
						types.StringElement{
							Content: "grandchild title",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of grandchild",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of grandchild",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source, WithFilename("test.adoc"))).To(Equal(expected))
	})

	It("should include grandchild content with relative offset", func() {
		source := `include::../../test/includes/grandchild-include.adoc[leveloffset=+1]`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      2,
					Title: []interface{}{
						types.StringElement{
							Content: "grandchild title",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of grandchild",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of grandchild",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source, WithFilename("test.adoc"))).To(Equal(expected))
	})

	It("should include grandchild content with absolute offset", func() {
		source := `include::../../test/includes/grandchild-include.adoc[leveloffset=0]`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      0,
					Title: []interface{}{
						types.StringElement{
							Content: "grandchild title",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of grandchild",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of grandchild",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source, WithFilename("test.adoc"))).To(Equal(expected))
	})

	It("should include child and grandchild content with relative level offset", func() {
		source := `include::../../test/includes/parent-include-relative-offset.adoc[leveloffset=+1]`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      1, // here the level is changed from `0` to `1` since `root` doc has a `leveloffset=+1` during its inclusion
					Title: []interface{}{
						types.StringElement{
							Content: "parent title",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of parent",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "child preamble",
							},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      3, // here the level is changed from `1` to `3` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
					Title: []interface{}{
						types.StringElement{
							Content: "child section 1",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of child",
							},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      4, // here the level is changed from `1` to `4` since both `root`, `parent` and `child` docs have a `leveloffset=+1` during their inclusion
					Title: []interface{}{
						types.StringElement{
							Content: "grandchild title",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of grandchild",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of grandchild",
							},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      4, // here the level is changed from `2` to `4` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
					Title: []interface{}{
						types.StringElement{
							Content: "child section 2",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of child",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of parent",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source, WithFilename("test.adoc"))).To(Equal(expected))
	})

	It("should include child and grandchild content with relative then absolute level offset", func() {
		source := `include::../../test/includes/parent-include-absolute-offset.adoc[leveloffset=+1]`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      1, // here the level is offset by `+1` as per root doc attribute in the `include` macro
					Title: []interface{}{
						types.StringElement{
							Content: "parent title",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of parent",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "child preamble",
							},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      3, // here level is forced to "absolute 3"
					Title: []interface{}{
						types.StringElement{
							Content: "child section 1",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of child",
							},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      4, // here the level is set to `4` because it was its parent was offset by 3...
					Title: []interface{}{
						types.StringElement{
							Content: "grandchild title",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "first line of grandchild",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of grandchild",
							},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      4, // here the level is set to `4` because it the first section was moved from `1` to `3` so we use the same offset here
					Title: []interface{}{
						types.StringElement{
							Content: "child section 2",
						},
					},
					Elements: []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of child",
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "last line of parent",
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source, WithFilename("test.adoc"))).To(Equal(expected))
	})

	Context("file inclusions in delimited blocks", func() {

		It("should include adoc file within fenced block", func() {
			source := "```\n" +
				"include::../../test/includes/parent-include.adoc[]\n" +
				"```"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Fenced,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "= parent title",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "first line of parent",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "= child title",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "first line of child",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "== grandchild title",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "first line of grandchild",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "last line of grandchild",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "last line of child",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "last line of parent",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("should include adoc file within listing block", func() {
			source := `----
include::../../test/includes/chapter-a.adoc[]
----`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("should include adoc file within example block", func() {
			source := `====
include::../../test/includes/chapter-a.adoc[]
====`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Example,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("should include adoc file within quote block", func() {
			source := `____
include::../../test/includes/chapter-a.adoc[]
____`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Quote,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("should include adoc file within verse block", func() {
			source := `[verse]
____
include::../../test/includes/chapter-a.adoc[]
____`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Verse,
						},
						Kind: types.Verse,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("should include adoc file within sidebar block", func() {
			source := `****
include::../../test/includes/chapter-a.adoc[]
****`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Sidebar,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("should include adoc file within passthrough block", func() {
			Skip("missing support for passthrough blocks")
			source := `++++
include::../../test/includes/chapter-a.adoc[]
++++`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						// Kind:       types.Passthrough,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})

	Context("file inclusions with line ranges", func() {

		Context("file inclusions with unquoted line ranges", func() {

			It("file inclusion with single unquoted line", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{},
							Level:      0,
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("file inclusion with multiple unquoted lines", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1..2]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level:      0,
							Attributes: types.ElementAttributes{},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("file inclusion with multiple unquoted ranges", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]` // paragraph becomes the author since the in-between blank line is stripped out
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "content",
									},
								},
							},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("file inclusion with invalid unquoted range - case 1", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level:      0,
							Attributes: types.ElementAttributes{},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("file inclusion with invalid unquoted range - case 2", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level:      0,
							Attributes: types.ElementAttributes{},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("file inclusions with quoted line ranges", func() {

			It("file inclusion with single quoted line", func() {
				console, reset := ConfigureLogger()
				defer reset()
				source := `include::../../test/includes/chapter-a.adoc[lines="1"]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level:      0,
							Attributes: types.ElementAttributes{},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
				// verify no error/warning in logs
				Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("file inclusion with multiple quoted lines", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1..2"]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level:      0,
							Attributes: types.ElementAttributes{},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("file inclusion with multiple quoted ranges", func() {
				// here, the `content` paragraph gets attached to the header and becomes the author
				source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "content",
									},
								},
							},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("file inclusion with invalid quoted range - case 1", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level:      0,
							Attributes: types.ElementAttributes{},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("file inclusion with invalid quoted range - case 2", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level:      0,
							Attributes: types.ElementAttributes{},
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("file inclusion with ignored tags", func() {
				// include using a line range a file having tags
				source := `include::../../test/includes/tag-include.adoc[lines=3]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Title: []interface{}{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
	})

	Context("file inclusions with tag ranges", func() {

		It("file inclusion with single tag", func() {
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../test/includes/tag-include.adoc[tag=section]`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title: []interface{}{
							types.StringElement{
								Content: "Section 1",
							},
						},
						Elements: []interface{}{},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			// verify no error/warning in logs
			Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
		})

		It("file inclusion with surrounding tag", func() {
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../test/includes/tag-include.adoc[tag=doc]`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title: []interface{}{
							types.StringElement{
								Content: "Section 1",
							},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
					types.BlankLine{},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			// verify no error/warning in logs
			Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
		})

		It("file inclusion with unclosed tag", func() {
			// setup logger to write in a buffer so we can check the output
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../test/includes/tag-include-unclosed.adoc[tag=unclosed]`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
					types.BlankLine{},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "end",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			// verify error in logs
			Expect(console).To(
				ContainMessageWithLevel(
					log.ErrorLevel,
					"detected unclosed tag 'unclosed' starting at line 6 of include file: ../../test/includes/tag-include-unclosed.adoc",
				))
		})

		It("file inclusion with unknown tag", func() {
			// given
			// setup logger to write in a buffer so we can check the output
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../test/includes/tag-include.adoc[tag=unknown]`
			expected := types.DraftDocument{
				Blocks: []interface{}{},
			}
			// when/then
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			// verify error in logs
			Expect(console).To(
				ContainMessageWithLevel(
					log.ErrorLevel,
					"tag 'unknown' not found in include file: ../../test/includes/tag-include.adoc",
				))
		})

		It("file inclusion with no tag", func() {
			source := `include::../../test/includes/tag-include.adoc[]`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title: []interface{}{
							types.StringElement{
								Content: "Section 1",
							},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "content",
								},
							},
						},
					},
					types.BlankLine{},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "end",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		Context("permutations", func() {

			It("all lines", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=**]` // includes all content except lines with tags
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Title: []interface{}{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
						types.BlankLine{},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "end",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("all tagged regions", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=*]` // includes all sections
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Title: []interface{}{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
						types.BlankLine{},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("all the lines outside and inside of tagged regions", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=**;*]` // includes all sections
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Title: []interface{}{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "content",
									},
								},
							},
						},
						types.BlankLine{},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "end",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("regions tagged doc, but not nested regions tagged content", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=doc;!content]` // includes all sections
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Title: []interface{}{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("all tagged regions, but excludes any regions tagged content", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=*;!content]` // includes all sections
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Title: []interface{}{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("all tagged regions, but excludes any regions tagged content", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=**;!content]` // includes all sections
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{},
							Level:      1,
							Title: []interface{}{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "end",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("**;!* — selects only the regions of the document outside of tags", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=**;!*]` // includes all sections
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "end",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
	})

	Context("missing file to include", func() {

		It("should replace with string element if directory does not exist in standalone block", func() {
			// setup logger to write in a buffer so we can check the output
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			// verify error in logs
			Expect(console).To(
				ContainMessageWithLevel(
					log.ErrorLevel,
					"failed to include '{unknown}/unknown.adoc'",
				))
		})

		It("should replace with string element if file is missing in standalone block", func() {
			// setup logger to write in a buffer so we can check the output
			console, reset := ConfigureLogger()
			defer reset()

			source := `include::../../test/includes/unknown.adoc[leveloffset=+1]`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "Unresolved directive in test.adoc - include::../../test/includes/unknown.adoc[leveloffset=+1]",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			// verify error in logs
			Expect(console).To(
				ContainMessageWithLevel(
					log.ErrorLevel,
					"failed to include '../../test/includes/unknown.adoc'",
				))
		})

		It("should replace with string element if file is missing in delimited block", func() {
			// setup logger to write in a buffer so we can check the output
			console, reset := ConfigureLogger()
			defer reset()

			source := `----
include::../../test/includes/unknown.adoc[leveloffset=+1]
----`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "Unresolved directive in test.adoc - include::../../test/includes/unknown.adoc[leveloffset=+1]",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			// verify error in logs
			Expect(console).To(
				ContainMessageWithLevel(
					log.ErrorLevel,
					"failed to include '../../test/includes/unknown.adoc'",
				))
		})
	})

	Context("inclusion with attribute in path", func() {

		It("should resolve path with attribute in standalone block from local file", func() {
			source := `:includedir: ../../test/includes
			
include::{includedir}/grandchild-include.adoc[]`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DocumentAttributeDeclaration{
						Name:  "includedir",
						Value: "../../test/includes",
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title: []interface{}{
							types.StringElement{
								Content: "grandchild title",
							},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "first line of grandchild",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "last line of grandchild",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source, WithFilename("foo.adoc"))).To(Equal(expected))
		})

		It("should resolve path with attribute in standalone block from relative file", func() {
			source := `:includedir: ../../../test/includes
			
include::{includedir}/grandchild-include.adoc[]`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DocumentAttributeDeclaration{
						Name:  "includedir",
						Value: "../../../test/includes",
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{},
						Level:      1,
						Title: []interface{}{
							types.StringElement{
								Content: "grandchild title",
							},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "first line of grandchild",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "last line of grandchild",
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source,
				WithFilename("tmp/foo.adoc"))).To(Equal(expected))
		})

		It("should resolve path with attribute in delimited block", func() {
			source := `:includedir: ../../test/includes

----
include::{includedir}/grandchild-include.adoc[]
----`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DocumentAttributeDeclaration{
						Name:  "includedir",
						Value: "../../test/includes",
					},
					types.BlankLine{},
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "== grandchild title",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "first line of grandchild",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "last line of grandchild",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})

	Context("inclusion of non-asciidoc file", func() {

		It("include go file without any range", func() {

			source := `----
include::../../test/includes/hello_world.go.txt[] 
----`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Kind:       types.Listing,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: `package includes`,
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: `import "fmt"`,
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: `func helloworld() {`,
										},
									},
									{
										types.StringElement{
											Content: `	fmt.Println("hello, world!")`,
										},
									},
									{
										types.StringElement{
											Content: `}`,
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("include go file with a simple range", func() {

			source := `----
include::../../test/includes/hello_world.go.txt[lines=1] 
----`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Kind:       types.Listing,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: `package includes`,
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})
})

var _ = Describe("file inclusions - final document", func() {

	It("should include child and grandchild content with relative level offset", func() {
		source := `include::../../test/includes/parent-include-relative-offset.adoc[leveloffset=+1]`
		expected := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"_parent_title": []interface{}{
					types.StringElement{
						Content: "parent title",
					},
				},
				"_child_section_1": []interface{}{
					types.StringElement{
						Content: "child section 1",
					},
				},
				"_child_section_2": []interface{}{
					types.StringElement{
						Content: "child section 2",
					},
				},
				"_grandchild_title": []interface{}{
					types.StringElement{
						Content: "grandchild title",
					},
				},
			},
			Footnotes: []types.Footnote{},
			Elements: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID: "_parent_title",
					},
					Level: 1, // here the level is changed from `0` to `1` since `root` doc has a `leveloffset=+1` during its inclusion
					Title: []interface{}{
						types.StringElement{
							Content: "parent title",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "first line of parent",
									},
								},
							},
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "child preamble",
									},
								},
							},
						},
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID: "_child_section_1",
							},
							Level: 3, // here the level is changed from `1` to `3` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
							Title: []interface{}{
								types.StringElement{
									Content: "child section 1",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "first line of child",
											},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_grandchild_title",
									},
									Level: 4, // here the level is changed from `1` to `4` since both `root`, `parent` and `child` docs have a `leveloffset=+1` during their inclusion
									Title: []interface{}{
										types.StringElement{
											Content: "grandchild title",
										},
									},
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "first line of grandchild",
													},
												},
											},
										},
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "last line of grandchild",
													},
												},
											},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_child_section_2",
									},
									Level: 4, // here the level is changed from `2` to `4` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
									Title: []interface{}{
										types.StringElement{
											Content: "child section 2",
										},
									},
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "last line of child",
													},
												},
											},
										},
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "last line of parent",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDocument(source)).To(MatchDocument(expected))
	})

	It("should include child and grandchild content with relative then absolute level offset", func() {
		source := `include::../../test/includes/parent-include-absolute-offset.adoc[leveloffset=+1]`
		expected := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"_parent_title": []interface{}{
					types.StringElement{
						Content: "parent title",
					},
				},
				"_child_section_1": []interface{}{
					types.StringElement{
						Content: "child section 1",
					},
				},
				"_child_section_2": []interface{}{
					types.StringElement{
						Content: "child section 2",
					},
				},
				"_grandchild_title": []interface{}{
					types.StringElement{
						Content: "grandchild title",
					},
				},
			},
			Footnotes: []types.Footnote{},
			Elements: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID: "_parent_title",
					},
					Level: 1, // here the level is offset by `+1` as per root doc attribute in the `include` macro
					Title: []interface{}{
						types.StringElement{
							Content: "parent title",
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "first line of parent",
									},
								},
							},
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "child preamble",
									},
								},
							},
						},
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID: "_child_section_1",
							},
							Level: 3, // here level is forced to "absolute 3"
							Title: []interface{}{
								types.StringElement{
									Content: "child section 1",
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "first line of child",
											},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_grandchild_title",
									},
									Level: 4, // here the level is set to `4` because it was its parent was offset by 3...
									Title: []interface{}{
										types.StringElement{
											Content: "grandchild title",
										},
									},
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "first line of grandchild",
													},
												},
											},
										},
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "last line of grandchild",
													},
												},
											},
										},
									},
								},
								types.Section{
									Attributes: types.ElementAttributes{
										types.AttrID: "_child_section_2",
									},
									Level: 4, // here the level is set to `4` because it the first section was moved from `1` to `3` so we use the same offset here
									Title: []interface{}{
										types.StringElement{
											Content: "child section 2",
										},
									},
									Elements: []interface{}{
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "last line of child",
													},
												},
											},
										},
										types.Paragraph{
											Attributes: types.ElementAttributes{},
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "last line of parent",
													},
												},
											},
										},
									},
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

var _ = Describe("file inclusions - draft without preprocessing", func() {

	It("should include adoc file without leveloffset in local dir", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../test/includes/chapter-a.adoc[]"
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.FileInclusion{
					Attributes: types.ElementAttributes{},
					Location: types.Location{
						Elements: []interface{}{
							types.StringElement{
								Content: "../../test/includes/chapter-a.adoc",
							},
						},
					},
					RawText: source,
				},
			},
		}
		Expect(ParseDraftDocument(source,
			WithoutPreprocessing(), WithFilename("test.adoc"))).To(Equal(expected))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file without leveloffset in relative dir", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../../test/includes/chapter-a.adoc[]"
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.FileInclusion{
					Attributes: types.ElementAttributes{},
					Location: types.Location{
						Elements: []interface{}{
							types.StringElement{
								Content: "../../../test/includes/chapter-a.adoc",
							},
						},
					},
					RawText: source,
				},
			},
		}
		Expect(ParseDraftDocument(source,
			WithoutPreprocessing(), WithFilename("tmp/foo.adoc"))).To(Equal(expected))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file with leveloffset attribute", func() {
		source := "include::../../test/includes/chapter-a.adoc[leveloffset=+1]"
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLevelOffset: "+1",
					},
					Location: types.Location{
						Elements: []interface{}{
							types.StringElement{
								Content: "../../test/includes/chapter-a.adoc",
							},
						},
					},
					RawText: source,
				},
			},
		}
		Expect(ParseDraftDocument(source,
			WithoutPreprocessing())).To(Equal(expected))
	})

	Context("file inclusions in delimited blocks", func() {

		It("should include adoc file within fenced block", func() {
			source := "```\n" +
				"include::../../test/includes/chapter-a.adoc[]\n" +
				"```"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Fenced,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "../../test/includes/chapter-a.adoc",
										},
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source,
				WithoutPreprocessing())).To(Equal(expected))
		})

		It("should include adoc file within listing block", func() {
			source := `----
include::../../test/includes/chapter-a.adoc[]
----`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "../../test/includes/chapter-a.adoc",
										},
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source,
				WithoutPreprocessing())).To(Equal(expected))
		})

		It("should include adoc file within example block", func() {
			source := `====
include::../../test/includes/chapter-a.adoc[]
====`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Example,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "../../test/includes/chapter-a.adoc",
										},
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source,
				WithoutPreprocessing())).To(Equal(expected))
		})

		It("should include adoc file within quote block", func() {
			source := `____
include::../../test/includes/chapter-a.adoc[]
____`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Quote,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "../../test/includes/chapter-a.adoc",
										},
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source,
				WithoutPreprocessing())).To(Equal(expected))
		})

		It("should include adoc file within verse block", func() {
			source := `[verse]
____
include::../../test/includes/chapter-a.adoc[]
____`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Verse,
						},
						Kind: types.Verse,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "../../test/includes/chapter-a.adoc",
										},
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source,
				WithoutPreprocessing())).To(Equal(expected))
		})

		It("should include adoc file within sidebar block", func() {
			source := `****
include::../../test/includes/chapter-a.adoc[]
****`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Sidebar,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "../../test/includes/chapter-a.adoc",
										},
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source,
				WithoutPreprocessing())).To(Equal(expected))
		})

		It("should include adoc file within passthrough block", func() {
			Skip("missing support for passthrough blocks")
			source := `++++
include::../../test/includes/chapter-a.adoc[]
++++`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						// Kind:       types.Passthrough,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "../../test/includes/chapter-a.adoc",
										},
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source,
				WithoutPreprocessing())).To(Equal(expected))
		})
	})

	Context("file inclusions with line ranges", func() {

		Context("file inclusions with unquoted line ranges", func() {

			It("file inclusion with single unquoted line", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 1},
								},
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with multiple unquoted lines", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1..2]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 2},
								},
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1..2]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with multiple unquoted ranges", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 1},
									{StartLine: 3, EndLine: 4},
									{StartLine: 6, EndLine: -1},
								},
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with invalid unquoted range - case 1", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `1;3..4;6..foo`,
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with invalid unquoted range - case 2", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 1},
								},
								"3..4":  nil,
								"6..-1": nil,
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with invalid unquoted range - case 3", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=foo]` // using commas instead of semi-colons
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: "foo",
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=foo]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})
		})

		Context("file inclusions with quoted line ranges", func() {

			It("file inclusion with single quoted line", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1"]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 1},
								},
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines="1"]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with multiple quoted lines", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1..2"]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 2},
								},
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines="1..2"]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with multiple quoted ranges", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 1},
									{StartLine: 3, EndLine: 4},
									{StartLine: 6, EndLine: -1},
								},
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with invalid quoted range - case 1", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `"1`, // viewed as a string
								"3..4":               nil,
								"6..foo":             nil,
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]`,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with invalid quoted range - case 2", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `"1;3..4;6..10"`,
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
							},
							RawText: source,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})
		})

		Context("file inclusions with tag ranges", func() {

			It("file inclusion with single tag", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=section]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrTagRanges: types.TagRanges{
									{
										Name:     `section`,
										Included: true,
									},
								},
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/tag-include.adoc",
									},
								},
							},
							RawText: source,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})

			It("file inclusion with multiple tags", func() {
				source := `include::../../test/includes/tag-include.adoc[tags=section;content]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrTagRanges: types.TagRanges{
									{
										Name:     `section`,
										Included: true,
									},
									{
										Name:     "content",
										Included: true,
									},
								},
							},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "../../test/includes/tag-include.adoc",
									},
								},
							},
							RawText: source,
						},
					},
				}
				Expect(ParseDraftDocument(source,
					WithoutPreprocessing())).To(Equal(expected))
			})
		})
	})
})
