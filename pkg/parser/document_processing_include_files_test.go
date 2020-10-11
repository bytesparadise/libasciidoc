package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega" //nolint golint
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
	Entry("'chapter-a.adoc'", "chapter-a.adoc", types.Location{
		Path: []interface{}{
			types.StringElement{
				Content: "chapter-a.adoc",
			},
		},
	}),
	Entry("'chapter_a.adoc'", "chapter_a.adoc", types.Location{
		Path: []interface{}{
			types.StringElement{
				Content: "chapter_a.adoc",
			},
		},
	}),
	Entry("'../../test/includes/chapter_a.adoc'", "../../test/includes/chapter_a.adoc", types.Location{
		Path: []interface{}{
			types.StringElement{
				Content: "../../test/includes/chapter_a.adoc",
			},
		},
	}),
	Entry("'{includedir}/chapter-{foo}.adoc'", "{includedir}/chapter-{foo}.adoc", types.Location{
		Path: []interface{}{
			types.StringElement{
				Content: "{includedir}/chapter-{foo}.adoc", // attribute substitutions are treared as part of the string element
			},
		},
	}),
	Entry("'{scheme}://{path}'", "{scheme}://{path}", types.Location{
		Path: []interface{}{
			types.StringElement{ // attribute substitutions are treared as part of the string element
				Content: "{scheme}://{path}",
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

var _ = Describe("file inclusions", func() {

	Context("raw document", func() {

		Context("with file inclusions", func() {

			It("should include adoc file without leveloffset from local file", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := "include::../../test/includes/chapter-a.adoc[]"
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				result, err := ParseRawDocument(source, WithFilename("foo.adoc"))
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchRawDocument(expected))
				// verify no error/warning in logs
				Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("should include adoc file without leveloffset from relative file", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := "include::../../../test/includes/chapter-a.adoc[]"
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source, WithFilename("tmp/foo.adoc"))).To(MatchRawDocument(expected))
				// verify no error/warning in logs
				Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("should include adoc file with leveloffset", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := "include::../../test/includes/chapter-a.adoc[leveloffset=+1]"
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 1,
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
				// verify no error/warning in logs
				Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("should include section 0 by default", func() {
				source := "include::../../test/includes/chapter-a.adoc[]"
				// at this level (parsing), it is expected that the Section 0 is part of the Prefligh document
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
			})

			It("should not include section 0 when attribute found", func() {
				source := `:includedir: ../../test/includes

include::{includedir}/chapter-a.adoc[]`
				// at this level (parsing), it is expected that the Section 0 is part of the Prefligh document
				expected := types.RawDocument{
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "includedir",
							Value: "../../test/includes",
						},
						types.BlankLine{},
						types.Section{
							Level: 0,
							Title: []interface{}{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				result, err := ParseRawDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchRawDocument(expected))
			})

			It("should not further process with non-asciidoc files", func() {
				source := `:includedir: ../../test/includes

include::{includedir}/include.foo[]`
				expected := types.RawDocument{
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "includedir",
							Value: "../../test/includes",
						},
						types.BlankLine{},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "*some strong content*",
									},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source, WithFilename("foo.bar"))).To(MatchRawDocument(expected)) // parent doc may not need to be a '.adoc'
			})

			It("should include grandchild content without offset", func() {
				source := `include::../../test/includes/grandchild-include.adoc[]`
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 1,
							Title: []interface{}{
								types.StringElement{
									Content: "grandchild title",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source, WithFilename("test.adoc"))).To(MatchRawDocument(expected))
			})

			It("should include grandchild content with relative offset", func() {
				source := `include::../../test/includes/grandchild-include.adoc[leveloffset=+1]`
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 2,
							Title: []interface{}{
								types.StringElement{
									Content: "grandchild title",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source, WithFilename("test.adoc"))).To(MatchRawDocument(expected))
			})

			It("should include grandchild content with absolute offset", func() {
				source := `include::../../test/includes/grandchild-include.adoc[leveloffset=0]`
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Title: []interface{}{
								types.StringElement{
									Content: "grandchild title",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source, WithFilename("test.adoc"))).To(MatchRawDocument(expected))
			})

			It("should include child and grandchild content with relative level offset", func() {
				source := `include::../../test/includes/parent-include-relative-offset.adoc[leveloffset=+1]`
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 1, // here the level is changed from `0` to `1` since `root` doc has a `leveloffset=+1` during its inclusion
							Title: []interface{}{
								types.StringElement{
									Content: "parent title",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
							Level: 3, // here the level is changed from `1` to `3` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
							Title: []interface{}{
								types.StringElement{
									Content: "child section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
							Level: 4, // here the level is changed from `1` to `4` since both `root`, `parent` and `child` docs have a `leveloffset=+1` during their inclusion
							Title: []interface{}{
								types.StringElement{
									Content: "grandchild title",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
							Level: 4, // here the level is changed from `2` to `4` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
							Title: []interface{}{
								types.StringElement{
									Content: "child section 2",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source, WithFilename("test.adoc"))).To(MatchRawDocument(expected))
			})

			It("should include child and grandchild content with relative then absolute level offset", func() {
				source := `include::../../test/includes/parent-include-absolute-offset.adoc[leveloffset=+1]`
				expected := types.RawDocument{
					Elements: []interface{}{
						types.Section{
							Level: 1, // here the level is offset by `+1` as per root doc attribute in the `include` macro
							Title: []interface{}{
								types.StringElement{
									Content: "parent title",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
							Level: 3, // here level is forced to "absolute 3"
							Title: []interface{}{
								types.StringElement{
									Content: "child section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
							Level: 4, // here the level is set to `4` because it was its parent was offset by 3...
							Title: []interface{}{
								types.StringElement{
									Content: "grandchild title",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
							Level: 4, // here the level is set to `4` because it the first section was moved from `1` to `3` so we use the same offset here
							Title: []interface{}{
								types.StringElement{
									Content: "child section 2",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
						types.Paragraph{
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
				Expect(ParseRawDocument(source, WithFilename("test.adoc"))).To(MatchRawDocument(expected))
			})

			It("should include adoc file within fenced block", func() {
				source := "```\n" +
					"include::../../test/includes/parent-include.adoc[]\n" +
					"```"
				expected := types.RawDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "= parent title",
									},
								},
								{
									types.StringElement{
										Content: "",
									},
								},
								{
									types.StringElement{
										Content: "first line of parent",
									},
								},
								{
									types.StringElement{
										Content: "",
									},
								},
								{
									types.StringElement{
										Content: "= child title",
									},
								},
								{
									types.StringElement{
										Content: "",
									},
								},
								{
									types.StringElement{
										Content: "first line of child",
									},
								},
								{
									types.StringElement{
										Content: "",
									},
								},
								{
									types.StringElement{
										Content: "== grandchild title",
									},
								},
								{
									types.StringElement{
										Content: "",
									},
								},
								{
									types.StringElement{
										Content: "first line of grandchild",
									},
								},
								{
									types.StringElement{
										Content: "",
									},
								},
								{
									types.StringElement{
										Content: "last line of grandchild",
									},
								},
								{
									types.StringElement{
										Content: "",
									},
								},
								{
									types.StringElement{
										Content: "last line of child",
									},
								},
								{
									types.StringElement{
										Content: "",
									},
								},
								{
									types.StringElement{
										Content: "last line of parent",
									},
								},
							},
						},
					},
				}
				Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
			})

			It("should include adoc file within quote block", func() {
				source := "____\n" +
					"include::../../test/includes/parent-include.adoc[]\n" +
					"____"
				expected := types.RawDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Elements: []interface{}{
								types.Paragraph{
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
				result, err := ParseRawDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchRawDocument(expected))
			})

			Context("with line ranges", func() {

				Context("with unquoted line ranges", func() {

					It("file inclusion with single unquoted line", func() {
						source := `include::../../test/includes/chapter-a.adoc[lines=1]`
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
									Title: []interface{}{
										types.StringElement{
											Content: "Chapter A",
										},
									},
									Elements: []interface{}{},
								},
							},
						}
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("file inclusion with multiple unquoted lines", func() {
						source := `include::../../test/includes/chapter-a.adoc[lines=1..2]`
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("file inclusion with multiple unquoted ranges", func() {
						source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]` // paragraph becomes the author since the in-between blank line is stripped out
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
									Attributes: types.Attributes{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("file inclusion with invalid unquoted range - case 1", func() {
						source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
									Title: []interface{}{
										types.StringElement{
											Content: "Chapter A",
										},
									},
									Elements: []interface{}{},
								},
								types.BlankLine{},
								types.Paragraph{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("file inclusion with invalid unquoted range - case 2", func() {
						source := `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
									Title: []interface{}{
										types.StringElement{
											Content: "Chapter A",
										},
									},
									Elements: []interface{}{},
								},
							},
						}
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})
				})

				Context("with quoted line ranges", func() {

					It("file inclusion with single quoted line", func() {
						logs, reset := ConfigureLogger(log.WarnLevel)
						defer reset()
						source := `include::../../test/includes/chapter-a.adoc[lines="1"]`
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
									Title: []interface{}{
										types.StringElement{
											Content: "Chapter A",
										},
									},
									Elements: []interface{}{},
								},
							},
						}
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
						// verify no error/warning in logs
						Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
					})

					It("file inclusion with multiple quoted lines", func() {
						source := `include::../../test/includes/chapter-a.adoc[lines="1..2"]`
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("file inclusion with multiple quoted ranges", func() {
						// here, the `content` paragraph gets attached to the header and becomes the author
						source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
									Attributes: types.Attributes{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("file inclusion with invalid quoted range - case 1", func() {
						source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
									Title: []interface{}{
										types.StringElement{
											Content: "Chapter A",
										},
									},
									Elements: []interface{}{},
								},
								types.BlankLine{},
								types.Paragraph{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("file inclusion with invalid quoted range - case 2", func() {
						source := `include::../../test/includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 0,
									Title: []interface{}{
										types.StringElement{
											Content: "Chapter A",
										},
									},
									Elements: []interface{}{},
								},
								types.BlankLine{},
								types.Paragraph{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("file inclusion with ignored tags", func() {
						// include using a line range a file having tags
						source := `include::../../test/includes/tag-include.adoc[lines=3]`
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 1,
									Title: []interface{}{
										types.StringElement{
											Content: "Section 1",
										},
									},
									Elements: []interface{}{},
								},
							},
						}
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})
				})
			})

			Context("with tag ranges", func() {

				It("file inclusion with single tag", func() {
					logs, reset := ConfigureLogger(log.WarnLevel)
					defer reset()
					source := `include::../../test/includes/tag-include.adoc[tag=section]`
					expected := types.RawDocument{
						Elements: []interface{}{
							types.Section{
								Level: 1,
								Title: []interface{}{
									types.StringElement{
										Content: "Section 1",
									},
								},
								Elements: []interface{}{},
							},
						},
					}
					Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					// verify no error/warning in logs
					Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
				})

				It("file inclusion with surrounding tag", func() {
					logs, reset := ConfigureLogger(log.WarnLevel)
					defer reset()
					source := `include::../../test/includes/tag-include.adoc[tag=doc]`
					expected := types.RawDocument{
						Elements: []interface{}{
							types.Section{
								Level: 1,
								Title: []interface{}{
									types.StringElement{
										Content: "Section 1",
									},
								},
								Elements: []interface{}{},
							},
							types.BlankLine{},
							types.Paragraph{
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
					Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					// verify no error/warning in logs
					Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
				})

				It("file inclusion with unclosed tag", func() {
					// setup logger to write in a buffer so we can check the output
					logs, reset := ConfigureLogger(log.WarnLevel)
					defer reset()
					source := `include::../../test/includes/tag-include-unclosed.adoc[tag=unclosed]`
					expected := types.RawDocument{
						Elements: []interface{}{
							// leading blanklines are ignored
							types.Paragraph{
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
					Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					// verify error in logs
					Expect(logs).To(ContainMessageWithLevel(log.WarnLevel,
						"detected unclosed tag 'unclosed' starting at line 6 of include file: ../../test/includes/tag-include-unclosed.adoc",
					))
				})

				It("file inclusion with unknown tag", func() {
					// given
					source := `include::../../test/includes/tag-include.adoc[tag=unknown]`
					// when/then
					_, err := ParseRawDocument(source)
					// verify error in logs
					Expect(err).To(MatchError("tag 'unknown' not found in include file: ../../test/includes/tag-include.adoc"))
				})

				It("file inclusion with no tag", func() {
					source := `include::../../test/includes/tag-include.adoc[]`
					expected := types.RawDocument{
						Elements: []interface{}{
							types.Section{
								Level: 1,
								Title: []interface{}{
									types.StringElement{
										Content: "Section 1",
									},
								},
								Elements: []interface{}{},
							},
							types.BlankLine{},
							types.Paragraph{
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
					Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
				})

				Context("permutations", func() {

					It("all lines", func() {
						source := `include::../../test/includes/tag-include.adoc[tag=**]` // includes all content except lines with tags
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 1,
									Title: []interface{}{
										types.StringElement{
											Content: "Section 1",
										},
									},
									Elements: []interface{}{},
								},
								types.BlankLine{},
								types.Paragraph{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("all tagged regions", func() {
						source := `include::../../test/includes/tag-include.adoc[tag=*]` // includes all sections
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 1,
									Title: []interface{}{
										types.StringElement{
											Content: "Section 1",
										},
									},
									Elements: []interface{}{},
								},
								types.BlankLine{},
								types.Paragraph{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("all the lines outside and inside of tagged regions", func() {
						source := `include::../../test/includes/tag-include.adoc[tag=**;*]` // includes all sections
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 1,
									Title: []interface{}{
										types.StringElement{
											Content: "Section 1",
										},
									},
									Elements: []interface{}{},
								},
								types.BlankLine{},
								types.Paragraph{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("regions tagged doc, but not nested regions tagged content", func() {
						source := `include::../../test/includes/tag-include.adoc[tag=doc;!content]` // includes all sections
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 1,
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("all tagged regions, but excludes any regions tagged content", func() {
						source := `include::../../test/includes/tag-include.adoc[tag=*;!content]` // includes all sections
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 1,
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("all tagged regions, but excludes any regions tagged content", func() {
						source := `include::../../test/includes/tag-include.adoc[tag=**;!content]` // includes all sections
						expected := types.RawDocument{
							Elements: []interface{}{
								types.Section{
									Level: 1,
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})

					It("**;!* — selects only the regions of the document outside of tags", func() {
						source := `include::../../test/includes/tag-include.adoc[tag=**;!*]` // includes all sections
						expected := types.RawDocument{
							Elements: []interface{}{
								// leading blanklines are ignored
								types.Paragraph{
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
						Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
					})
				})
			})

			Context("missing file to include", func() {

				It("should fail if directory does not exist in standalone block", func() {
					source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
					_, err := ParseRawDocument(source)
					Expect(err).To(MatchError("Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]"))
				})

				It("should fail if file is missing in standalone block", func() {
					source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
					_, err := ParseRawDocument(source)
					Expect(err).To(MatchError("Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]"))
				})

				It("should fail if file with attribute in path is not resolved in standalone block", func() {
					source := `include::{includedir}/unknown.adoc[leveloffset=+1]`
					_, err := ParseRawDocument(source)
					Expect(err).To(MatchError("Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]"))
				})

				It("should fail if file is missing in delimited block", func() {
					source := `----
include::../../test/includes/unknown.adoc[leveloffset=+1]
----`
					_, err := ParseRawDocument(source)
					Expect(err).To(MatchError("Unresolved directive in test.adoc - include::../../test/includes/unknown.adoc[leveloffset=+1]"))
				})

				It("should fail if file with attribute in path is not resolved in delimited block", func() {
					// setup logger to write in a buffer so we can check the output
					source := `----
include::{includedir}/unknown.adoc[leveloffset=+1]
----`
					_, err := ParseRawDocument(source)
					Expect(err).To(MatchError("Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]"))
				})
			})

			Context("inclusion with attribute in path", func() {

				It("should resolve path with attribute in standalone block from local file", func() {
					source := `:includedir: ../../test/includes
			
include::{includedir}/grandchild-include.adoc[]`
					expected := types.RawDocument{
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "includedir",
								Value: "../../test/includes",
							},
							types.BlankLine{},
							types.Section{
								Level: 1,
								Title: []interface{}{
									types.StringElement{
										Content: "grandchild title",
									},
								},
								Elements: []interface{}{},
							},
							types.BlankLine{},
							types.Paragraph{
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
					Expect(ParseRawDocument(source, WithFilename("foo.adoc"))).To(MatchRawDocument(expected))
				})

				It("should resolve path with attribute in standalone block from relative file", func() {
					source := `:includedir: ../../../test/includes
			
include::{includedir}/grandchild-include.adoc[]`
					expected := types.RawDocument{
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "includedir",
								Value: "../../../test/includes",
							},
							types.BlankLine{},
							types.Section{
								Level: 1,
								Title: []interface{}{
									types.StringElement{
										Content: "grandchild title",
									},
								},
								Elements: []interface{}{},
							},
							types.BlankLine{},
							types.Paragraph{
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
					Expect(ParseRawDocument(source, WithFilename("tmp/foo.adoc"))).To(MatchRawDocument(expected))
				})

				It("should resolve path with attribute in delimited block", func() {
					source := `:includedir: ../../test/includes

----
include::{includedir}/grandchild-include.adoc[]
----`
					expected := types.RawDocument{
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "includedir",
								Value: "../../test/includes",
							},
							types.BlankLine{},
							types.ListingBlock{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "== grandchild title",
										},
									},
									{
										types.StringElement{},
									},
									{
										types.StringElement{
											Content: "first line of grandchild",
										},
									},
									{
										types.StringElement{},
									},
									{
										types.StringElement{
											Content: "last line of grandchild",
										},
									},
								},
							},
						},
					}
					Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
				})
			})

			Context("inclusion of non-asciidoc file", func() {

				It("include go file without any range", func() {

					source := `----
include::../../test/includes/hello_world.go.txt[] 
----`
					expected := types.RawDocument{
						Elements: []interface{}{
							types.ListingBlock{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: `package includes`,
										},
									},
									{
										// TODO: why do we need such an empty StringElement?
										types.StringElement{
											Content: ``,
										},
									},
									{
										types.StringElement{
											Content: `import "fmt"`,
										},
									},
									{
										types.StringElement{
											Content: ``,
										},
									},
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
					}
					Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
				})

				It("include go file with a simple range", func() {

					source := `----
include::../../test/includes/hello_world.go.txt[lines=1] 
----`
					expected := types.RawDocument{
						Elements: []interface{}{
							types.ListingBlock{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: `package includes`,
										},
									},
								},
							},
						},
					}
					Expect(ParseRawDocument(source)).To(MatchRawDocument(expected))
				})
			})
		})

	})

	Context("final document", func() {

		It("should include child and grandchild content with relative level offset", func() {
			source := `include::../../test/includes/parent-include-relative-offset.adoc[leveloffset=+1]`
			expected := types.Document{
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
				Elements: []interface{}{
					types.Section{
						Attributes: types.Attributes{
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
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "first line of parent",
										},
									},
								},
							},
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "child preamble",
										},
									},
								},
							},
							types.Section{
								Attributes: types.Attributes{
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
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "first line of child",
												},
											},
										},
									},
									types.Section{
										Attributes: types.Attributes{
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
												Lines: [][]interface{}{
													{
														types.StringElement{
															Content: "first line of grandchild",
														},
													},
												},
											},
											types.Paragraph{
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
										Attributes: types.Attributes{
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
												Lines: [][]interface{}{
													{
														types.StringElement{
															Content: "last line of child",
														},
													},
												},
											},
											types.Paragraph{
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
				Elements: []interface{}{
					types.Section{
						Attributes: types.Attributes{
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
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "first line of parent",
										},
									},
								},
							},
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "child preamble",
										},
									},
								},
							},
							types.Section{
								Attributes: types.Attributes{
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
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "first line of child",
												},
											},
										},
									},
									types.Section{
										Attributes: types.Attributes{
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
												Lines: [][]interface{}{
													{
														types.StringElement{
															Content: "first line of grandchild",
														},
													},
												},
											},
											types.Paragraph{
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
										Attributes: types.Attributes{
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
												Lines: [][]interface{}{
													{
														types.StringElement{
															Content: "last line of child",
														},
													},
												},
											},
											types.Paragraph{
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
})
