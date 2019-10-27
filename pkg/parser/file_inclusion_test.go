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

var _ = Describe("file location", func() {

	DescribeTable("'FileLocation' pattern",
		func(filename string, expected interface{}) {
			reader := strings.NewReader(filename)
			actual, err := parser.ParseReader(filename, reader, parser.Entrypoint("FileLocation"))
			Expect(err).ToNot(HaveOccurred())
			GinkgoT().Log("actual result: %s", spew.Sdump(actual))
			GinkgoT().Log("expected result: %s", spew.Sdump(expected))
			Expect(actual).To(Equal(expected))
		},
		Entry("'chapter'", "chapter", types.Location{
			types.StringElement{
				Content: "chapter",
			},
		}),
		Entry("'chapter.adoc'", "chapter.adoc", types.Location{
			types.StringElement{
				Content: "chapter.adoc",
			},
		}),
		Entry("'chapter-a.adoc'", "chapter-a.adoc", types.Location{
			types.StringElement{
				Content: "chapter-a.adoc",
			},
		}),
		Entry("'chapter_a.adoc'", "chapter_a.adoc", types.Location{
			types.StringElement{
				Content: "chapter_a.adoc",
			},
		}),
		Entry("'../../test/includes/chapter_a.adoc'", "../../test/includes/chapter_a.adoc", types.Location{
			types.StringElement{
				Content: "../../test/includes/chapter_a.adoc",
			},
		}),
		Entry("'chapter-{foo}.adoc'", "chapter-{foo}.adoc", types.Location{
			types.StringElement{
				Content: "chapter-",
			},
			types.DocumentAttributeSubstitution{
				Name: "foo",
			},
			types.StringElement{
				Content: ".adoc",
			},
		}),
		Entry("'{includedir}/chapter-{foo}.adoc'", "{includedir}/chapter-{foo}.adoc", types.Location{
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
		}),
	)
})

var _ = Describe("file inclusions - preflight with preprocessing", func() {

	It("should include adoc file without leveloffset from local file", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../test/includes/chapter-a.adoc[]"
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID:       "chapter_a",
						types.AttrCustomID: false,
					},
					Level: 0,
					Title: types.InlineElements{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
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
		Expect(source).To(BecomePreflightDocument(expected, WithFilename("foo.adoc")))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file without leveloffset from relative file", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../../test/includes/chapter-a.adoc[]"
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID:       "chapter_a",
						types.AttrCustomID: false,
					},
					Level: 0,
					Title: types.InlineElements{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
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
		Expect(source).To(BecomePreflightDocument(expected, WithFilename("tmp/foo.adoc")))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file with leveloffset", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../test/includes/chapter-a.adoc[leveloffset=+1]"
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID:       "chapter_a",
						types.AttrCustomID: false,
					},
					Level: 1,
					Title: types.InlineElements{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
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
		Expect(source).To(BecomePreflightDocument(expected))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include section 0 by default", func() {
		source := "include::../../test/includes/chapter-a.adoc[]"
		// at this level (parsing), it is expected that the Section 0 is part of the Prefligh document
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID:       "chapter_a",
						types.AttrCustomID: false,
					},
					Level: 0,
					Title: types.InlineElements{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
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
		Expect(source).To(BecomePreflightDocument(expected))
	})

	It("should not include section 0 when attribute exists", func() {
		source := `:includedir: ../../test/includes

include::{includedir}/chapter-a.adoc[]`
		// at this level (parsing), it is expected that the Section 0 is part of the Prefligh document
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.DocumentAttributeDeclaration{
					Name:  "includedir",
					Value: "../../test/includes",
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID:       "chapter_a",
						types.AttrCustomID: false,
					},
					Level: 0,
					Title: types.InlineElements{
						types.StringElement{
							Content: "Chapter A",
						},
					},
					Elements: []interface{}{},
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
		Expect(source).To(BecomePreflightDocument(expected))
	})

	Context("file inclusions in delimited blocks", func() {

		It("should include adoc file within fenced block", func() {
			source := "```\n" +
				"include::../../test/includes/chapter-a.adoc[]\n" +
				"```"
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
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
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("should include adoc file within listing block", func() {
			source := `----
include::../../test/includes/chapter-a.adoc[]
----`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
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
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("should include adoc file within example block", func() {
			source := `====
include::../../test/includes/chapter-a.adoc[]
====`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
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
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("should include adoc file within quote block", func() {
			source := `____
include::../../test/includes/chapter-a.adoc[]
____`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
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
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("should include adoc file within verse block", func() {
			source := `[verse]
____
include::../../test/includes/chapter-a.adoc[]
____`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Verse,
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
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("should include adoc file within sidebar block", func() {
			source := `****
include::../../test/includes/chapter-a.adoc[]
****`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
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
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("should include adoc file within passthrough block", func() {
			Skip("missing support for passthrough blocks")
			source := `++++
include::../../test/includes/chapter-a.adoc[]
++++`
			expected := types.DelimitedBlock{
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
			Expect(source).To(BecomePreflightDocument(expected))
		})
	})

	Context("file inclusions with line ranges", func() {

		Context("file inclusions with unquoted line ranges", func() {

			It("file inclusion with single unquoted line", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Level: 0,
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("file inclusion with multiple unquoted lines", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1..2]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("file inclusion with multiple unquoted ranges", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]` // paragraph becomes the author since the in-between blank line is stripped out
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "content",
									},
								},
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("file inclusion with invalid unquoted range - case 1", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
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
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("file inclusion with invalid unquoted range - case 2", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})
		})

		Context("file inclusions with quoted line ranges", func() {

			It("file inclusion with single quoted line", func() {
				console, reset := ConfigureLogger()
				defer reset()
				source := `include::../../test/includes/chapter-a.adoc[lines="1"]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
				// verify no error/warning in logs
				Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("file inclusion with multiple quoted lines", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1..2"]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("file inclusion with multiple quoted ranges", func() {
				// here, the `content` paragraph gets attached to the header and becomes the author
				source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
								types.AttrAuthors: []types.DocumentAuthor{
									{
										FullName: "content",
									},
								},
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("file inclusion with invalid quoted range - case 1", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
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
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("file inclusion with invalid quoted range - case 2", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID:       "chapter_a",
								types.AttrCustomID: false,
							},
							Title: types.InlineElements{
								types.StringElement{
									Content: "Chapter A",
								},
							},
							Elements: []interface{}{},
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
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("file inclusion with ignored tags", func() {
				// include using a line range a file having tags
				source := `include::../../test/includes/tag-include.adoc[lines=3]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID:       "section_1",
								types.AttrCustomID: false,
							},
							Level: 1,
							Title: types.InlineElements{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})
		})
	})

	Context("file inclusions with tag ranges", func() {

		It("file inclusion with single tag", func() {
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../test/includes/tag-include.adoc[tag=section]`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1",
							types.AttrCustomID: false,
						},
						Level: 1,
						Title: types.InlineElements{
							types.StringElement{
								Content: "Section 1",
							},
						},
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
			// verify no error/warning in logs
			Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
		})

		It("file inclusion with surrounding tag", func() {
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../test/includes/tag-include.adoc[tag=doc]`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrCustomID: false,
							types.AttrID:       "section_1",
						},
						Level: 1,
						Title: types.InlineElements{
							types.StringElement{
								Content: "Section 1",
							},
						},
						Elements: []interface{}{},
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
					types.BlankLine{},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
			// verify no error/warning in logs
			Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
		})

		It("file inclusion with unclosed tag", func() {
			// setup logger to write in a buffer so we can check the output
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::../../test/includes/tag-include-unclosed.adoc[tag=unclosed]`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
					types.BlankLine{},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "end",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
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
			expected := types.PreflightDocument{
				Blocks: []interface{}{},
			}
			// when/then
			Expect(source).To(BecomePreflightDocument(expected))
			// verify error in logs
			Expect(console).To(
				ContainMessageWithLevel(
					log.ErrorLevel,
					"tag 'unknown' not found in include file: ../../test/includes/tag-include.adoc",
				))
		})

		It("file inclusion with no tag", func() {
			source := `include::../../test/includes/tag-include.adoc[]`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1",
							types.AttrCustomID: false,
						},
						Level: 1,
						Title: types.InlineElements{
							types.StringElement{
								Content: "Section 1",
							},
						},
						Elements: []interface{}{},
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
					types.BlankLine{},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "end",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		Context("permutations", func() {

			It("all lines", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=**]` // includes all content except lines with tags
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrCustomID: false,
								types.AttrID:       "section_1",
							},
							Level: 1,
							Title: types.InlineElements{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
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
						types.BlankLine{},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "end",
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("all tagged regions", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=*]` // includes all sections
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID:       "section_1",
								types.AttrCustomID: false,
							},
							Level: 1,
							Title: types.InlineElements{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
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
						types.BlankLine{},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("all the lines outside and inside of tagged regions", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=**;*]` // includes all sections
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrCustomID: false,
								types.AttrID:       "section_1",
							},
							Level: 1,
							Title: types.InlineElements{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
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
						types.BlankLine{},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "end",
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("regions tagged doc, but not nested regions tagged content", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=doc;!content]` // includes all sections
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrCustomID: false,
								types.AttrID:       "section_1",
							},
							Level: 1,
							Title: types.InlineElements{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("all tagged regions, but excludes any regions tagged content", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=*;!content]` // includes all sections
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrCustomID: false,
								types.AttrID:       "section_1",
							},
							Level: 1,
							Title: types.InlineElements{
								types.StringElement{
									Content: "Section 1",
								},
							},
							Elements: []interface{}{},
						},
						types.BlankLine{},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("all tagged regions, but excludes any regions tagged content", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=**;!content]` // includes all sections
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrCustomID: false,
								types.AttrID:       "section_1",
							},
							Level: 1,
							Title: types.InlineElements{
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
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "end",
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})

			It("**;!* — selects only the regions of the document outside of tags", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=**;!*]` // includes all sections
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "end",
									},
								},
							},
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected))
			})
		})
	})

	Context("missing file to include", func() {

		It("should replace with string element if directory does not exist in standalone block", func() {
			// setup logger to write in a buffer so we can check the output
			console, reset := ConfigureLogger()
			defer reset()
			source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
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
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "Unresolved directive in test.adoc - include::../../test/includes/unknown.adoc[leveloffset=+1]",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
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
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
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
			Expect(source).To(BecomePreflightDocument(expected))
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
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DocumentAttributeDeclaration{
						Name:  "includedir",
						Value: "../../test/includes",
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
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
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "last line of grandchild",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithFilename("foo.adoc")))
		})

		It("should resolve path with attribute in standalone block from relative file", func() {
			source := `:includedir: ../../../test/includes
			
include::{includedir}/grandchild-include.adoc[]`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DocumentAttributeDeclaration{
						Name:  "includedir",
						Value: "../../../test/includes",
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
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
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "last line of grandchild",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithFilename("tmp/foo.adoc")))
		})

		It("should resolve path with attribute in delimited block", func() {
			source := `:includedir: ../../test/includes

----
include::{includedir}/grandchild-include.adoc[]
----`
			expected := types.PreflightDocument{
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
								Lines: []types.InlineElements{
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
								Lines: []types.InlineElements{
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
			Expect(source).To(BecomePreflightDocument(expected))
		})
	})

	Context("inclusion of non-asciidoc file", func() {

		It("include go file without any range", func() {

			source := `----
include::../../test/includes/hello_world.go.txt[] 
----`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Kind:       types.Listing,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
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
								Lines: []types.InlineElements{
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
								Lines: []types.InlineElements{
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
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("include go file with a simple range", func() {

			source := `----
include::../../test/includes/hello_world.go.txt[lines=1] 
----`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Kind:       types.Listing,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
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
			Expect(source).To(BecomePreflightDocument(expected))
		})
	})
})

var _ = Describe("file inclusions - preflight without preprocessing", func() {

	It("should include adoc file without leveloffset in local dir", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../test/includes/chapter-a.adoc[]"
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.FileInclusion{
					Attributes: types.ElementAttributes{},
					Location: types.Location{
						types.StringElement{
							Content: "../../test/includes/chapter-a.adoc",
						},
					},
					RawText: source,
				},
			},
		}
		Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing(), WithFilename("foo.adoc")))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file without leveloffset in relative dir", func() {
		console, reset := ConfigureLogger()
		defer reset()
		source := "include::../../../test/includes/chapter-a.adoc[]"
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.FileInclusion{
					Attributes: types.ElementAttributes{},
					Location: types.Location{
						types.StringElement{
							Content: "../../../test/includes/chapter-a.adoc",
						},
					},
					RawText: source,
				},
			},
		}
		Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing(), WithFilename("tmp/foo.adoc")))
		// verify no error/warning in logs
		Expect(console).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
	})

	It("should include adoc file with leveloffset attribute", func() {
		source := "include::../../test/includes/chapter-a.adoc[leveloffset=+1]"
		expected := types.PreflightDocument{
			Blocks: []interface{}{
				types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLevelOffset: "+1",
					},
					Location: types.Location{
						types.StringElement{
							Content: "../../test/includes/chapter-a.adoc",
						},
					},
					RawText: source,
				},
			},
		}
		Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
	})

	Context("file inclusions in delimited blocks", func() {

		It("should include adoc file within fenced block", func() {
			source := "```\n" +
				"include::../../test/includes/chapter-a.adoc[]\n" +
				"```"
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Fenced,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
		})

		It("should include adoc file within listing block", func() {
			source := `----
include::../../test/includes/chapter-a.adoc[]
----`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
		})

		It("should include adoc file within example block", func() {
			source := `====
include::../../test/includes/chapter-a.adoc[]
====`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Example,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
		})

		It("should include adoc file within quote block", func() {
			source := `____
include::../../test/includes/chapter-a.adoc[]
____`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Quote,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
		})

		It("should include adoc file within verse block", func() {
			source := `[verse]
____
include::../../test/includes/chapter-a.adoc[]
____`
			expected := types.PreflightDocument{
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
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
		})

		It("should include adoc file within sidebar block", func() {
			source := `****
include::../../test/includes/chapter-a.adoc[]
****`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Sidebar,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
		})

		It("should include adoc file within passthrough block", func() {
			Skip("missing support for passthrough blocks")
			source := `++++
include::../../test/includes/chapter-a.adoc[]
++++`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						// Kind:       types.Passthrough,
						Elements: []interface{}{
							types.FileInclusion{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									types.StringElement{
										Content: "../../test/includes/chapter-a.adoc",
									},
								},
								RawText: `include::../../test/includes/chapter-a.adoc[]`,
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
		})
	})

	Context("file inclusions with line ranges", func() {

		Context("file inclusions with unquoted line ranges", func() {

			It("file inclusion with single unquoted line", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 1},
								},
							},
							Location: types.Location{
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with multiple unquoted lines", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1..2]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 2},
								},
							},
							Location: types.Location{
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1..2]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with multiple unquoted ranges", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]`
				expected := types.PreflightDocument{
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
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with invalid unquoted range - case 1", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `1;3..4;6..foo`,
							},
							Location: types.Location{
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with invalid unquoted range - case 2", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
				expected := types.PreflightDocument{
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
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with invalid unquoted range - case 3", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines=foo]` // using commas instead of semi-colons
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: "foo",
							},
							Location: types.Location{
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines=foo]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})
		})

		Context("file inclusions with quoted line ranges", func() {

			It("file inclusion with single quoted line", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1"]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 1},
								},
							},
							Location: types.Location{
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines="1"]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with multiple quoted lines", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1..2"]`
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{StartLine: 1, EndLine: 2},
								},
							},
							Location: types.Location{
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines="1..2"]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with multiple quoted ranges", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
				expected := types.PreflightDocument{
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
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with invalid quoted range - case 1", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `"1`, // viewed as a string
								"3..4":               nil,
								"6..foo":             nil,
							},
							Location: types.Location{
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]`,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with invalid quoted range - case 2", func() {
				source := `include::../../test/includes/chapter-a.adoc[lines="1;3..4;6..10"]` // using semi-colons instead of commas
				expected := types.PreflightDocument{
					Blocks: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: `"1;3..4;6..10"`,
							},
							Location: types.Location{
								types.StringElement{
									Content: "../../test/includes/chapter-a.adoc",
								},
							},
							RawText: source,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})
		})

		Context("file inclusions with tag ranges", func() {

			It("file inclusion with single tag", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=section]`
				expected := types.PreflightDocument{
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
								types.StringElement{
									Content: "../../test/includes/tag-include.adoc",
								},
							},
							RawText: source,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

			It("file inclusion with multiple tags", func() {
				source := `include::../../test/includes/tag-include.adoc[tags=section;content]`
				expected := types.PreflightDocument{
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
								types.StringElement{
									Content: "../../test/includes/tag-include.adoc",
								},
							},
							RawText: source,
						},
					},
				}
				Expect(source).To(BecomePreflightDocument(expected, WithoutPreprocessing()))
			})

		})
	})
})
