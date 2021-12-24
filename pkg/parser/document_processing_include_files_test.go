package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
	log "github.com/sirupsen/logrus"
)

var _ = Describe("file inclusions", func() {

	Context("in final documents", func() {

		It("should include adoc file without leveloffset from local file", func() {
			logs, reset := ConfigureLogger(log.WarnLevel)
			defer reset()
			source := "include::../../test/includes/chapter-a.adoc[]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "content",
							},
						},
					},
				},
			}
			result, err := ParseDocument(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDocument(expected))
			// verify no error/warning in logs
			Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
		})

		It("should include adoc file with leveloffset", func() {
			logs, reset := ConfigureLogger(log.WarnLevel)
			defer reset()
			source := "include::../../test/includes/chapter-a.adoc[leveloffset=+1]"
			title := []interface{}{
				&types.StringElement{
					Content: "Chapter A",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_chapter_a",
						},
						Level: 1, // offset by +1
						Title: title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_chapter_a": title,
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
			// verify no error/warning in logs
			Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
		})

		It("should include file with attribute in path", func() {
			source := `:includedir: ../../test/includes

include::{includedir}/chapter-a.adoc[]`
			title := []interface{}{
				&types.StringElement{
					Content: "Chapter A",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "includedir",
						Value: "../../test/includes",
					},
					&types.DocumentHeader{
						Title: title,
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "content",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should not further process with non-asciidoc files", func() {
			source := `:includedir: ../../test/includes

include::{includedir}/include.foo[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "includedir",
						Value: "../../test/includes",
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "some strong content",
									},
								},
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `include::hello_world.go.txt[]`,
							},
						},
					},
				},
			}
			// Expect(ParseDocument(source, WithFilename("foo.bar"))).To(MatchDocumentFragments(expected)) // parent doc may not need to be a '.adoc'
			result, err := ParseDocument(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDocument(expected)) // parent doc may not need to be a '.adoc'
		})

		It("should include grandchild content without offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[]`
			title := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_grandchild_title",
						},
						Level: 1,
						Title: title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of grandchild",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "last line of grandchild",
									},
								},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_grandchild_title": title,
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include grandchild content with relative offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[leveloffset=+1]`
			title := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_grandchild_title",
						},
						Level: 2,
						Title: title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of grandchild",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "last line of grandchild",
									},
								},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_grandchild_title": title,
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include grandchild content with absolute offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[leveloffset=0]`
			title := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_grandchild_title",
						},
						Level: 0,
						Title: title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of grandchild",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "last line of grandchild",
									},
								},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_grandchild_title": title,
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include child and grandchild content with relative level offset", func() {
			source := `include::../../test/includes/parent-include-relative-offset.adoc[leveloffset=+1]`
			parentTitle := []interface{}{
				&types.StringElement{
					Content: "parent title",
				},
			}
			childSection1Title := []interface{}{
				&types.StringElement{
					Content: "child section 1",
				},
			}
			childSection2Title := []interface{}{
				&types.StringElement{
					Content: "child section 2",
				},
			}
			grandchildTitle := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_parent_title",
						},
						Level: 1, // here the level is changed from `0` to `1` since `root` doc has a `leveloffset=+1` during its inclusion
						Title: parentTitle,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of parent",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "child preamble",
									},
								},
							},
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_child_section_1",
								},
								Level: 3, // here the level is changed from `1` to `3` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
								Title: childSection1Title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "first line of child",
											},
										},
									},
									&types.Section{
										Attributes: types.Attributes{
											types.AttrID: "_grandchild_title",
										},
										Level: 4, // here the level is changed from `1` to `4` since both `root`, `parent` and `child` docs have a `leveloffset=+1` during their inclusion
										Title: grandchildTitle,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "first line of grandchild",
													},
												},
											},
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of grandchild",
													},
												},
											},
										},
									},
									&types.Section{
										Attributes: types.Attributes{
											types.AttrID: "_child_section_2",
										},
										Level: 4, // here the level is changed from `2` to `4` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
										Title: childSection2Title,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of child",
													},
												},
											},
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
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
				ElementReferences: types.ElementReferences{
					"_parent_title":     parentTitle,
					"_child_section_1":  childSection1Title,
					"_child_section_2":  childSection2Title,
					"_grandchild_title": grandchildTitle,
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include child and grandchild content with relative then absolute level offset", func() {
			source := `include::../../test/includes/parent-include-absolute-offset.adoc[leveloffset=+1]`
			parentTitle := []interface{}{
				&types.StringElement{
					Content: "parent title",
				},
			}
			childSection1Title := []interface{}{
				&types.StringElement{
					Content: "child section 1",
				},
			}
			childSection2Title := []interface{}{
				&types.StringElement{
					Content: "child section 2",
				},
			}
			grandchildTitle := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_parent_title",
						},
						Level: 1, // here the level is changed from `0` to `1` since `root` doc has a `leveloffset=+1` during its inclusion
						Title: parentTitle,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of parent",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "child preamble",
									},
								},
							},
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_child_section_1",
								},
								Level: 3, // here level is forced to "absolute 3"
								Title: childSection1Title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "first line of child",
											},
										},
									},
									&types.Section{
										Attributes: types.Attributes{
											types.AttrID: "_grandchild_title",
										},
										Level: 4, // here the level is set to `4` because it the first section was moved from `1` to `3` so we use the same offset here
										Title: grandchildTitle,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "first line of grandchild",
													},
												},
											},
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of grandchild",
													},
												},
											},
										},
									},
									&types.Section{
										Attributes: types.Attributes{
											types.AttrID: "_child_section_2",
										},
										Level: 4, // here the level is changed from `2` to `4` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
										Title: childSection2Title,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of child",
													},
												},
											},
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
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
				ElementReferences: types.ElementReferences{
					"_parent_title":     parentTitle,
					"_child_section_1":  childSection1Title,
					"_child_section_2":  childSection2Title,
					"_grandchild_title": grandchildTitle,
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include adoc with attributes file within main content", func() {
			source := `include::../../test/includes/attributes.adoc[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "author",
						Value: "Xavier",
					},
					&types.AttributeDeclaration{
						Name:  "leveloffset",
						Value: "+1",
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "some content",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include adoc with attributes file within listing block", func() {
			source := `----
include::../../test/includes/attributes.adoc[]
----`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Listing,
						Elements: []interface{}{
							&types.StringElement{
								Content: `:author: Xavier
:leveloffset: +1

some content`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include adoc file within fenced block", func() {
			source := "```\n" +
				"include::../../test/includes/parent-include.adoc[]\n" +
				"```\n" +
				"<1> a callout"
			// include the doc without parsing the elements (besides the file inclusions)
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Fenced,
						Elements: []interface{}{
							&types.StringElement{
								Content: `:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent `,
							},
							&types.Callout{
								Ref: 1,
							},
							&types.StringElement{
								Content: "\n\n:leveloffset!:",
							},
						},
					},
					&types.List{
						Kind: types.CalloutListKind,
						Elements: []types.ListElement{
							&types.CalloutListElement{
								Ref: 1,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "a callout",
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

		It("should include adoc file within quote block", func() {
			source := "____\n" +
				"include::../../test/includes/parent-include.adoc[]\n" +
				"____"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Quote,
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "leveloffset",
								Value: string("+1"),
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "= parent title",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of parent",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "= child title",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of child",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "== grandchild title",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of grandchild",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "last line of grandchild",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "last line of child",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "last line of parent ",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "1",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
								},
							},
							&types.AttributeReset{
								Name: "leveloffset",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		Context("with line ranges", func() {

			Context("unquoted", func() {

				It("with single unquoted line", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with multiple unquoted lines", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1..3]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "content",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with multiple unquoted ranges (becoming authors)", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]` // paragraph becomes the author since the in-between blank line is stripped out
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "content",
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

				It("with invalid unquoted range - case 1", func() {
					logs, reset := ConfigureLogger(log.WarnLevel)
					defer reset()
					source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
					expected := &types.Document{}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
					Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 0, 64, "Unresolved directive in test.adoc - include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]"))
				})

				It("with invalid unquoted range - case 2", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			Context("quoted", func() {

				It("with single line", func() {
					logs, reset := ConfigureLogger(log.WarnLevel)
					defer reset()
					source := `include::../../test/includes/chapter-a.adoc[lines="1"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
					// verify no error/warning in logs
					Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
				})

				It("with multiple lines", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines="1..2"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with multiple ranges with colons (becoming authors)", func() {
					// here, the `content` paragraph gets attached to the header and becomes the author
					source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "content",
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

				It("with multiple ranges with semicolons (becoming authors)", func() {
					// here, the `content` paragraph gets attached to the header and becomes the author
					source := `include::../../test/includes/chapter-a.adoc[lines="1;3..4;6..-1"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "content",
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

				It("with invalid range - case 1", func() {
					logs, reset := ConfigureLogger(log.WarnLevel)
					defer reset()
					source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
					_, err := ParseDocument(source)
					Expect(err).NotTo(HaveOccurred()) // parsing does not fail, but an error is logged for the fragment with the `include` macro
					Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 0, 66, "Unresolved directive in test.adoc - include::../../test/includes/chapter-a.adoc[lines=\"1,3..4,6..foo\"]"))
				})

				It("with ignored tags", func() {
					// include using a line range a file having tags
					source := `include::../../test/includes/tag-include.adoc[lines=3]`
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("with tag ranges", func() {

			It("with single tag", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::../../test/includes/tag-include.adoc[tag=section]`
				title := []interface{}{
					&types.StringElement{
						Content: "Section 1",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": title,
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				// verify no error/warning in logs
				Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("with surrounding tag", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::../../test/includes/tag-include.adoc[tag=doc]`
				title := []interface{}{
					&types.StringElement{
						Content: "Section 1",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": title,
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				// verify no error/warning in logs
				Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("with unclosed tag", func() {
				// setup logger to write in a buffer so we can check the output
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::../../test/includes/tag-include-unclosed.adoc[tag=unclosed]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "content",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "end",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				// verify error in logs
				Expect(logs).To(ContainJSONLog(log.WarnLevel, "detected unclosed tag 'unclosed' starting at line 6 of include file: ../../test/includes/tag-include-unclosed.adoc"))
			})

			It("with unknown tag", func() {
				// given
				// setup logger to write in a buffer so we can check the output
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::../../test/includes/tag-include.adoc[tag=unknown]`
				// when/then
				_, err := ParseDocument(source)
				// verify error in logs
				Expect(err).NotTo(HaveOccurred()) // parsing does not fail, but an error is logged for the fragment with the `include` macro
				Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 0, 58, "Unresolved directive in test.adoc - include::../../test/includes/tag-include.adoc[tag=unknown]: tag 'unknown' not found in file to include"))
			})

			It("with no tag", func() {
				source := `include::../../test/includes/tag-include.adoc[]`
				title := []interface{}{
					&types.StringElement{
						Content: "Section 1",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "content",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "end",
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": title,
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("permutations", func() {

				It("all lines", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**]` // includes all content except lines with tags
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "content",
											},
										},
									},
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "end",
											},
										},
									},
								},
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("all tagged regions", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=*]` // includes all regions
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "content",
											},
										},
									},
								},
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("all the lines outside and inside of tagged regions", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;*]` // includes all regions
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "content",
											},
										},
									},
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "end",
											},
										},
									},
								},
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("regions tagged doc, but not nested regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=doc;!content]` // includes all `doc` but `content`
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("all tagged regions, but excludes any regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=*;!content]` // includes all but `content`
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("all tagged regions, but excludes any regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;!content]` // includes all lines but `content`
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "end",
											},
										},
									},
								},
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("**;!* — selects only the regions of the document outside of tags", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;!*]` // excludes all regions
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "end",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("with missing file to include", func() {

			It("should fail if directory does not exist in standalone block", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
				expected := &types.Document{}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 0, 47, "Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]"))
			})

			It("should fail if file is missing in standalone block", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
				expected := &types.Document{}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 0, 47, "Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]"))
			})

			It("should fail if file with attribute in path is not resolved in standalone block", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::{includedir}/unknown.adoc[leveloffset=+1]`
				expected := &types.Document{}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 0, 50, "Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]"))
			})

			It("should fail if file is missing in delimited block", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `----
include::../../test/includes/unknown.adoc[leveloffset=+1]
----`
				expected := &types.Document{}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 0, 67, "Unresolved directive in test.adoc - include::../../test/includes/unknown.adoc[leveloffset=+1]"))
			})

			It("should fail if file with attribute in path is not resolved in delimited block", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `----
include::{includedir}/unknown.adoc[leveloffset=+1]
----`
				expected := &types.Document{}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 0, 60, "Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]"))
			})
		})

		Context("with inclusion with attribute in path", func() {

			It("should resolve path with attribute in standalone block from local file", func() {
				source := `:includedir: ../../test/includes
			
include::{includedir}/grandchild-include.adoc[]`
				title := []interface{}{
					&types.StringElement{
						Content: "grandchild title",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  "includedir",
							Value: "../../test/includes",
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_grandchild_title",
							},
							Level: 1,
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "first line of grandchild",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "last line of grandchild",
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_grandchild_title": title,
					},
				}
				// Expect(ParseDocument(source, WithFilename("foo.adoc"))).To(MatchDocument(expected))
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should resolve path with attribute in delimited block", func() {
				source := `:includedir: ../../test/includes

----
include::{includedir}/grandchild-include.adoc[]
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  "includedir",
							Value: "../../test/includes",
						},
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "== grandchild title\n\nfirst line of grandchild\n\nlast line of grandchild",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("inclusion of non-asciidoc file", func() {

			It("include go file without any range in listing block", func() {

				source := `----
include::../../test/includes/hello_world.go.txt[] 
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: `package includes

import "fmt"

func helloworld() {
	fmt.Println("hello, world!")
}`,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("include go file with a simple range in listing block", func() {

				source := `----
include::../../test/includes/hello_world.go.txt[lines=1] 
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: `package includes`,
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
