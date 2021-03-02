package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("line ranges", func() {

	Context("single range", func() {
		// given
		ranges := types.NewLineRanges(
			types.LineRange{StartLine: 2, EndLine: 4},
		)

		DescribeTable("match line range",
			func(line int, expectation bool) {
				Expect(ranges.Match(line)).To(Equal(expectation))
			},
			Entry("should not match line 1", 1, false),
			Entry("should match line 2", 2, true),
			Entry("should match line 3", 3, true),
			Entry("should match line 4", 4, true),
			Entry("should not match line 5", 5, false),
		)
	})

	Context("multiple ranges", func() {

		ranges := types.NewLineRanges([]interface{}{
			types.LineRange{StartLine: 1, EndLine: 1},
			types.LineRange{StartLine: 3, EndLine: 4},
			types.LineRange{StartLine: 6, EndLine: -1},
		})

		DescribeTable("match line range",
			func(line int, expectation bool) {
				Expect(ranges.Match(line)).To(Equal(expectation))
			},
			Entry("should match line 1", 1, true),
			Entry("should not match line 2", 2, false),
			Entry("should match line 3", 3, true),
			Entry("should match line 4", 4, true),
			Entry("should match line 6", 6, true),
			Entry("should match line 100", 100, true),
		)
	})

})

var _ = Describe("tag ranges", func() {

	DescribeTable("single range",
		func(line int, c types.CurrentRanges, expectation bool) {
			// given
			ranges := types.NewTagRanges(types.TagRange{
				Name:     "foo",
				Included: true,
			})
			// when
			match := ranges.Match(line, c)
			// then
			Expect(match).To(Equal(expectation))
		},
		Entry("should match within expected tag range", 2, types.CurrentRanges{
			"foo": &types.CurrentTagRange{
				StartLine: 1,
				EndLine:   -1, // range must be "open"
			},
		}, true),
		Entry("should not match outside expected tag range", 4, types.CurrentRanges{
			"foo": &types.CurrentTagRange{
				StartLine: 1,
				EndLine:   3,
			},
		}, false),
		Entry("should not match within unexpected tag range", 20, types.CurrentRanges{
			"bar": &types.CurrentTagRange{
				StartLine: 10,
				EndLine:   30,
			},
		}, false),
		Entry("should not match outside unexpected tag range", 40, types.CurrentRanges{
			"bar": &types.CurrentTagRange{
				StartLine: 10,
				EndLine:   30,
			},
		}, false),
	)

	DescribeTable("multiple ranges",
		func(line int, c types.CurrentRanges, expectation bool) {
			// given
			ranges := types.NewTagRanges([]interface{}{
				types.TagRange{
					Name:     "foo",
					Included: true,
				}, types.TagRange{
					Name:     "bar",
					Included: true,
				},
			})
			// when
			match := ranges.Match(line, c)
			// then
			Expect(match).To(Equal(expectation))
		},
		Entry("should match within first expected tag range", 2, types.CurrentRanges{
			"foo": &types.CurrentTagRange{
				StartLine: 1,
				EndLine:   -1, // range must be "open"
			},
		}, true),
		Entry("should match within second expected tag ranges", 5, types.CurrentRanges{
			"foo": &types.CurrentTagRange{
				StartLine: 1,
				EndLine:   3, // range must be "open"
			},
			"bar": &types.CurrentTagRange{
				StartLine: 4,
				EndLine:   -1, // range must be "open"
			},
		}, true),
		Entry("should not match outside expected tag range", 15, types.CurrentRanges{
			"foo": &types.CurrentTagRange{
				StartLine: 1,
				EndLine:   3,
			},
			"bar": &types.CurrentTagRange{
				StartLine: 10,
				EndLine:   20,
			},
		}, false),
		Entry("should not match within unexpected tag range", 25, types.CurrentRanges{
			"foo": &types.CurrentTagRange{
				StartLine: 1,
				EndLine:   3,
			},
			"baz": &types.CurrentTagRange{
				StartLine: 10,
				EndLine:   30,
			},
		}, false),
		Entry("should not match outside unexpected tag range", 40, types.CurrentRanges{
			"foo": &types.CurrentTagRange{
				StartLine: 1,
				EndLine:   3,
			},
			"baz": &types.CurrentTagRange{
				StartLine: 10,
				EndLine:   30,
			},
		}, false),
	)

	Context("permutations", func() {

		DescribeTable("** - all lines", // except lines containing a tag directive
			func(line int, c types.CurrentRanges, expectation bool) {
				// given
				ranges := types.NewTagRanges(types.TagRange{
					Name:     "**",
					Included: true,
				})
				// when
				match := ranges.Match(line, c)
				// then
				Expect(match).To(Equal(expectation))
			},
			Entry("should match within any tag ranges", 15, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   3, // range must be "open"
				},
				"bar": &types.CurrentTagRange{
					StartLine: 10,
					EndLine:   -1, // range must be "open"
				},
			}, true),
			Entry("should match outside any tag range", 25, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   3,
				},
				"bar": &types.CurrentTagRange{
					StartLine: 10,
					EndLine:   20,
				},
			}, true),
		)

		DescribeTable("* - all tagged regions", // except lines containing a tag directive
			func(line int, c types.CurrentRanges, expectation bool) {
				// given
				ranges := types.NewTagRanges(types.TagRange{
					Name:     "*",
					Included: true,
				})

				// when
				match := ranges.Match(line, c)
				// then
				Expect(match).To(Equal(expectation))
			},
			Entry("should match within any tag ranges", 15, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   3, // range must be "open"
				},
				"bar": &types.CurrentTagRange{
					StartLine: 10,
					EndLine:   -1, // range must be "open"
				},
			}, true),
			Entry("should not match outside any tag range", 25, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   3,
				},
				"bar": &types.CurrentTagRange{
					StartLine: 10,
					EndLine:   20,
				},
			}, false),
		)

		DescribeTable("**;* - all the lines outside and inside of tagged regions", // except lines containing a tag directive
			func(line int, c types.CurrentRanges, expectation bool) {
				// given
				ranges := types.NewTagRanges([]interface{}{
					types.TagRange{
						Name:     "**",
						Included: true,
					}, types.TagRange{
						Name:     "*",
						Included: true,
					},
				})
				// when
				match := ranges.Match(line, c)
				// then
				Expect(match).To(Equal(expectation))
			},
			Entry("should match within any tag ranges", 15, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   3, // range must be "open"
				},
				"bar": &types.CurrentTagRange{
					StartLine: 10,
					EndLine:   -1, // range must be "open"
				},
			}, true),
			Entry("should match outside any tag range", 25, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   3,
				},
				"bar": &types.CurrentTagRange{
					StartLine: 10,
					EndLine:   20,
				},
			}, true),
		)

		DescribeTable("foo;!bar - regions tagged foo, but not nested regions tagged bar",
			func(line int, c types.CurrentRanges, expectation bool) {
				// given
				ranges := types.NewTagRanges([]interface{}{types.TagRange{
					Name:     "foo",
					Included: true,
				}, types.TagRange{
					Name:     "bar",
					Included: false,
				},
				})
				// when
				match := ranges.Match(line, c)
				// then
				Expect(match).To(Equal(expectation))
			},
			Entry("should match within expected tag range", 2, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   -1, // range must be "open"
				},
				// "bar" is not be here yet, since we're still processing lines before its "start" tag
			}, true),
			Entry("should match within expected tag range", 16, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   -1, // range must be "open"
				},
				"bar": &types.CurrentTagRange{
					StartLine: 10,
					EndLine:   15,
				},
			}, true),
			Entry("should not match within excluded tag range", 12, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   -1, // range must be "open"
				},
				"bar": &types.CurrentTagRange{ // this range is excluded, and since we're on line 12, we can't include it
					StartLine: 10,
					EndLine:   -1,
				},
			}, false),
		)

		DescribeTable("*;!foo — all tagged regions, but excludes any regions tagged foo",
			func(line int, c types.CurrentRanges, expectation bool) {
				// given
				ranges := types.NewTagRanges([]interface{}{
					types.TagRange{
						Name:     "*",
						Included: true,
					}, types.TagRange{
						Name:     "foo",
						Included: false,
					},
				})
				// when
				match := ranges.Match(line, c)
				// then
				Expect(match).To(Equal(expectation))
			},
			Entry("should not match before any tag range", 1, types.CurrentRanges{}, false),
			Entry("should not match within foo tag range", 2, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   -1, // range must be "open"
				},
			}, false),
			Entry("should match in another range", 20, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   10, // range must be "open"
				},
				"bar": &types.CurrentTagRange{
					StartLine: 15,
					EndLine:   -1, // range must be "open"
				},
			}, true),
			Entry("should match in a range but outside foo tag range", 20, types.CurrentRanges{
				"bar": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   -1, // range must be "open"
				},
				"foo": &types.CurrentTagRange{
					StartLine: 3,
					EndLine:   10, // range is closed/passed
				},
			}, true),
			Entry("should not match after all tag ranges", 30, types.CurrentRanges{
				"bar": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   25, // range must be "open"
				},
				"foo": &types.CurrentTagRange{
					StartLine: 3,
					EndLine:   10, // range is closed/passed
				},
			}, false),
		)

		DescribeTable("**;!foo — selects all the lines of the document except for regions tagged foo",
			func(line int, c types.CurrentRanges, expectation bool) {
				// given
				ranges := types.NewTagRanges([]interface{}{
					types.TagRange{
						Name:     "**",
						Included: true,
					}, types.TagRange{
						Name:     "foo",
						Included: false,
					},
				})
				// when
				match := ranges.Match(line, c)
				// then
				Expect(match).To(Equal(expectation))
			},
			Entry("should match before any tag range", 1, types.CurrentRanges{}, true),
			Entry("should not match within foo tag range", 2, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   -1, // range must be "open"
				},
			}, false),
			Entry("should match in another range", 20, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   10, // range must be "open"
				},
				"bar": &types.CurrentTagRange{
					StartLine: 15,
					EndLine:   -1, // range must be "open"
				},
			}, true),
			Entry("should match in a range but outside foo tag range", 20, types.CurrentRanges{
				"bar": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   -1, // range must be "open"
				},
				"foo": &types.CurrentTagRange{
					StartLine: 3,
					EndLine:   10, // range is closed/passed
				},
			}, true),
			Entry("should match after all tag ranges", 30, types.CurrentRanges{
				"bar": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   25, // range must be "open"
				},
				"foo": &types.CurrentTagRange{
					StartLine: 3,
					EndLine:   10, // range is closed/passed
				},
			}, true),
		)

		DescribeTable("**;!* — selects only the regions of the document outside of tags (i.e., non-tagged regions).",
			func(line int, c types.CurrentRanges, expectation bool) {
				// given
				ranges := types.NewTagRanges([]interface{}{
					types.TagRange{
						Name:     "**",
						Included: true,
					}, types.TagRange{
						Name:     "*",
						Included: false,
					},
				})
				// when
				match := ranges.Match(line, c)
				// then
				Expect(match).To(Equal(expectation))
			},
			Entry("should match before any tag range", 1, types.CurrentRanges{}, true),
			Entry("should not match within foo tag range", 2, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   -1, // range must be "open"
				},
			}, false),
			Entry("should not match in another range", 20, types.CurrentRanges{
				"foo": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   10, // range must be "open"
				},
				"bar": &types.CurrentTagRange{
					StartLine: 15,
					EndLine:   -1, // range must be "open"
				},
			}, false),
			Entry("should match after all tag ranges", 30, types.CurrentRanges{
				"bar": &types.CurrentTagRange{
					StartLine: 1,
					EndLine:   25, // range must be "open"
				},
				"foo": &types.CurrentTagRange{
					StartLine: 3,
					EndLine:   10, // range is closed/passed
				},
			}, true),
		)
	})

	It("invalid tage ranges", func() {
		// when
		ranges := types.NewTagRanges([]interface{}{"foo", "bar"})
		// then
		Expect(ranges).To(BeEmpty())
	})

})

var _ = DescribeTable("raw document attributes",
	func(d types.RawDocument, expectation types.Attributes) {
		Expect(d.Attributes()).To(Equal(expectation))
	},

	Entry("should use attribute declarations at top of document only",
		types.RawDocument{
			Elements: []interface{}{
				types.AttributeDeclaration{
					Name:  "foo1",
					Value: "bar1",
				},
				types.AttributeDeclaration{
					Name:  "foo2",
					Value: "bar2",
				},
				types.BlankLine{},
				types.AttributeDeclaration{
					Name:  "foo3",
					Value: "bar3",
				},
			},
		},
		types.Attributes{
			"foo1": "bar1",
			"foo2": "bar2",
		},
	),
	Entry("should use attribute declarations right after section 0 only",
		types.RawDocument{
			Elements: []interface{}{
				types.Section{
					Level: 0,
				},
				types.AttributeDeclaration{
					Name:  "foo1",
					Value: "bar1",
				},
				types.AttributeDeclaration{
					Name:  "foo2",
					Value: "bar2",
				},
				types.BlankLine{},
				types.AttributeDeclaration{
					Name:  "foo3",
					Value: "bar3",
				},
			},
		},
		types.Attributes{
			"foo1": "bar1",
			"foo2": "bar2",
		},
	),
	Entry("should include attributes of section 0 only",
		types.RawDocument{
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Attributes: types.Attributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "foo",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
						},
						"other": "unused",
					},
				},
				types.BlankLine{},
				types.Section{
					Level: 1,
					Attributes: types.Attributes{
						"foo1": "bar1",
						"foo2": "bar2",
					},
				},
				types.AttributeDeclaration{
					Name:  "foo3",
					Value: "bar3",
				},
			},
		},
		types.Attributes{
			types.AttrAuthors: []types.DocumentAuthor{
				{
					FullName: "foo",
				},
			},
			// "authors" attribute gets "expanded" when being moved to the document attributes level
			"author":         "foo",
			"firstname":      "foo",
			"authorinitials": "f",
			// "revision" attribute is also "expanded"
			types.AttrRevision: types.DocumentRevision{
				Revnumber: "1.0",
			},
			"revnumber": "1.0",
		},
	),
	Entry("should ignore attribute declarations elsewhere",
		types.RawDocument{
			Elements: []interface{}{
				types.Section{
					Level: 1,
				},
				types.AttributeDeclaration{
					Name:  "foo1",
					Value: "bar1",
				},
				types.AttributeDeclaration{
					Name:  "foo2",
					Value: "bar2",
				},
				types.BlankLine{},
				types.AttributeDeclaration{
					Name:  "foo3",
					Value: "bar3",
				},
			},
		},
		types.Attributes{},
	),
)

var _ = Describe("element id resolution", func() {

	Context("sections", func() {

		Context("default it", func() {

			It("simple title", func() {
				// given
				section := types.Section{
					Level:      0,
					Attributes: types.Attributes{},
					Title: []interface{}{
						types.StringElement{
							Content: "foo",
						},
					},
					Elements: []interface{}{},
				}
				// when
				section, err := section.ResolveID(types.AttributesWithOverrides{
					Content:   map[string]interface{}{},
					Overrides: map[string]string{},
				})
				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(section.Attributes[types.AttrID]).To(Equal("_foo"))
			})

			It("title with link", func() {
				// given
				section := types.Section{
					Level:      0,
					Attributes: types.Attributes{},
					Title: []interface{}{
						types.StringElement{
							Content: "a link to ",
						},
						types.InlineLink{
							Location: types.Location{
								Scheme: "https://",
								Path: []interface{}{
									types.StringElement{
										Content: "foo.com",
									},
								},
							},
						},
					},
					Elements: []interface{}{},
				}
				// when
				section, err := section.ResolveID(types.AttributesWithOverrides{
					Content:   map[string]interface{}{},
					Overrides: map[string]string{},
				})
				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(section.Attributes[types.AttrID]).To(Equal("_a_link_to_httpsfoo_com")) // TODO: should be `httpsfoo`
			})
		})

		Context("custom id prefix", func() {

			It("simple title", func() {
				// given
				section := types.Section{
					Level:      0,
					Attributes: types.Attributes{},
					Title: []interface{}{
						types.StringElement{
							Content: "foo",
						},
					},
					Elements: []interface{}{},
				}
				// when
				section, err := section.ResolveID(types.AttributesWithOverrides{
					Content: map[string]interface{}{
						types.AttrIDPrefix: "custom_",
					},
					Overrides: map[string]string{},
				})
				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(section.Attributes[types.AttrID]).To(Equal("custom_foo"))
			})

			It("title with link", func() {
				// given
				section := types.Section{
					Level:      0,
					Attributes: types.Attributes{},
					Title: []interface{}{
						types.StringElement{
							Content: "a link to ",
						},
						types.InlineLink{
							Location: types.Location{
								Scheme: "https://",
								Path: []interface{}{
									types.StringElement{
										Content: "foo.com",
									},
								},
							},
						},
					},
					Elements: []interface{}{},
				}
				// when
				section, err := section.ResolveID(types.AttributesWithOverrides{
					Content: map[string]interface{}{
						types.AttrIDPrefix: "custom_",
					},
					Overrides: map[string]string{},
				})
				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(section.Attributes[types.AttrID]).To(Equal("custom_a_link_to_httpsfoo_com")) // TODO: should be `httpsfoo`
			})
		})

		Context("custom id", func() {

			It("simple title", func() {
				// given
				section := types.Section{
					Level: 0,
					Attributes: types.Attributes{
						types.AttrID: "bar",
					},
					Title: []interface{}{
						types.StringElement{
							Content: "foo",
						},
					},
					Elements: []interface{}{},
				}
				// when
				section, err := section.ResolveID(types.AttributesWithOverrides{
					Content: map[string]interface{}{
						types.AttrIDPrefix: "custom_",
					},
					Overrides: map[string]string{},
				})
				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(section.Attributes[types.AttrID]).To(Equal("bar"))
			})

			It("title with link", func() {
				// given
				section := types.Section{
					Level: 0,
					Attributes: types.Attributes{
						types.AttrID: "bar",
					},
					Title: []interface{}{
						types.StringElement{
							Content: "a link to ",
						},
						types.InlineLink{
							Location: types.Location{
								Scheme: "https://",
								Path: []interface{}{
									types.StringElement{
										Content: "foo.com",
									},
								},
							},
						},
					},
					Elements: []interface{}{},
				}
				// when
				section, err := section.ResolveID(types.AttributesWithOverrides{
					Content: map[string]interface{}{
						types.AttrIDPrefix: "custom_",
					},
					Overrides: map[string]string{},
				})
				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(section.Attributes[types.AttrID]).To(Equal("bar"))
			})
		})
	})
})

var _ = Describe("footnote replacements", func() {

	Context("sections", func() {

		It("title with footnote without ref", func() {
			// given
			section := types.Section{
				Level:      0,
				Attributes: types.Attributes{},
				Title: []interface{}{
					types.StringElement{
						Content: "foo",
					},
					types.Footnote{
						Elements: []interface{}{
							types.StringElement{
								Content: "a regular footnote.",
							},
						},
					},
				},
				Elements: []interface{}{},
			}
			footnotes := types.NewFootnotes()
			// when
			section.SubstituteFootnotes(footnotes)
			// then
			Expect(section).To(Equal(types.Section{
				Level:      0,
				Attributes: types.Attributes{},
				Title: []interface{}{
					types.StringElement{
						Content: "foo",
					},
					types.FootnoteReference{
						ID: 1,
					},
				},
				Elements: []interface{}{},
			}))
			Expect(footnotes.Notes()).To(Equal([]types.Footnote{
				{
					ID: 1,
					Elements: []interface{}{
						types.StringElement{
							Content: "a regular footnote.",
						},
					},
				},
			}))
		})

		It("title with footnote with ref", func() {
			// given
			section := types.Section{
				Level:      0,
				Attributes: types.Attributes{},
				Title: []interface{}{
					types.StringElement{
						Content: "foo",
					},
					types.Footnote{
						Ref: "disclaimer",
						Elements: []interface{}{
							types.StringElement{
								Content: "a regular footnote.",
							},
						},
					},
				},
				Elements: []interface{}{},
			}
			footnotes := types.NewFootnotes()
			// when
			section.SubstituteFootnotes(footnotes)
			// then
			Expect(section).To(Equal(types.Section{
				Level:      0,
				Attributes: types.Attributes{},
				Title: []interface{}{
					types.StringElement{
						Content: "foo",
					},
					types.FootnoteReference{
						ID:  1,
						Ref: "disclaimer",
					},
				},
				Elements: []interface{}{},
			}))
			Expect(footnotes.Notes()).To(Equal([]types.Footnote{
				{
					ID:  1,
					Ref: "disclaimer",
					Elements: []interface{}{
						types.StringElement{
							Content: "a regular footnote.",
						},
					},
				},
			}))
		})
	})

	Context("paragraphs", func() {

		It("paragraph with multiple footnotes", func() {
			// given
			paragraph := types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "first line",
						},
						types.Footnote{
							Ref: "disclaimer",
							Elements: []interface{}{
								types.StringElement{
									Content: "a disclaimer.",
								},
							},
						},
					},
					{
						types.StringElement{
							Content: "second line",
						},
						types.Footnote{
							Elements: []interface{}{
								types.StringElement{
									Content: "a regular footnote.",
								},
							},
						},
					},
					{
						types.StringElement{
							Content: "third line",
						},
						types.Footnote{
							Ref:      "disclaimer",
							Elements: []interface{}{},
						},
					},
				},
			}
			footnotes := types.NewFootnotes()
			// when
			paragraph.SubstituteFootnotes(footnotes)
			// then
			Expect(paragraph).To(Equal(types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "first line",
						},
						types.FootnoteReference{
							ID:  1,
							Ref: "disclaimer",
						},
					},
					{
						types.StringElement{
							Content: "second line",
						},
						types.FootnoteReference{
							ID: 2,
						},
					},
					{
						types.StringElement{
							Content: "third line",
						},
						types.FootnoteReference{
							ID:        1,
							Ref:       "disclaimer",
							Duplicate: true,
						},
					},
				},
			}))
			Expect(footnotes.Notes()).To(Equal([]types.Footnote{
				{
					ID:  1,
					Ref: "disclaimer",
					Elements: []interface{}{
						types.StringElement{
							Content: "a disclaimer.",
						},
					},
				},
				{
					ID: 2,
					Elements: []interface{}{
						types.StringElement{
							Content: "a regular footnote.",
						},
					},
				},
			}))
		})

		It("paragraph with invalid footnote reference", func() {
			// given
			paragraph := types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "first line",
						},
						types.Footnote{
							Ref: "disclaimer",
							Elements: []interface{}{
								types.StringElement{
									Content: "a disclaimer.",
								},
							},
						},
					},
					{
						types.StringElement{
							Content: "second line",
						},
						types.Footnote{
							Elements: []interface{}{
								types.StringElement{
									Content: "a regular footnote.",
								},
							},
						},
					},
					{
						types.StringElement{
							Content: "third line",
						},
						types.Footnote{
							Ref:      "disclaimer_",
							Elements: []interface{}{},
						},
					},
				},
			}
			footnotes := types.NewFootnotes()
			// when
			paragraph.SubstituteFootnotes(footnotes)
			// then
			Expect(paragraph).To(Equal(types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "first line",
						},
						types.FootnoteReference{
							ID:  1,
							Ref: "disclaimer",
						},
					},
					{
						types.StringElement{
							Content: "second line",
						},
						types.FootnoteReference{
							ID: 2,
						},
					},
					{
						types.StringElement{
							Content: "third line",
						},
						types.FootnoteReference{
							ID:  types.InvalidFootnoteReference, // marks as an invalid reference
							Ref: "disclaimer_",
						},
					},
				},
			}))
			Expect(footnotes.Notes()).To(Equal([]types.Footnote{
				{
					ID:  1,
					Ref: "disclaimer",
					Elements: []interface{}{
						types.StringElement{
							Content: "a disclaimer.",
						},
					},
				},
				{
					ID: 2,
					Elements: []interface{}{
						types.StringElement{
							Content: "a regular footnote.",
						},
					},
				},
			}))
		})
	})
})

var _ = DescribeTable("match for attribute with key and value",
	func(key string, value interface{}, expected bool) {
		// given
		attributes := []interface{}{
			types.Attribute{ // single attribute
				Key:   types.AttrStyle,
				Value: types.Quote,
			},
			types.Attributes{ // multiple attributes
				types.AttrStyle: types.Verse,
				types.AttrTitle: "verse title",
			},
		}
		// when
		result := types.HasAttributeWithValue(attributes, key, value)

		// then
		Expect(result).To((Equal(expected)))

	},
	Entry("match for block-kind: verse", types.AttrStyle, types.Verse, true),
	Entry("match for block-kind: quote", types.AttrStyle, types.Quote, true),
	Entry("no match for block-kind: quote", types.AttrID, "unknown", false),
)

var _ = DescribeTable("no match attribute with key",
	func(key string, expected bool) {
		// given
		attributes := []interface{}{
			types.Attribute{ // single attribute
				Key:   types.AttrStyle,
				Value: types.Quote,
			},
			types.Attributes{ // multiple attributes
				types.AttrStyle: types.Verse,
				types.AttrTitle: "verse title",
			},
		}
		// when
		result := types.HasNotAttribute(attributes, key)

		// then
		Expect(result).To((Equal(expected)))

	},
	Entry("match for block-kind: verse", types.AttrStyle, false),
	Entry("match for block-kind: quote", types.AttrStyle, false),
	Entry("no match for block-kind: quote", types.AttrID, true),
)

var _ = DescribeTable("rawtest",
	func(element types.RawText, expected string) {
		Expect(element.RawText()).To(Equal(expected))
	},
	// quoted text
	Entry("single quote bold text",
		types.QuotedText{
			Kind: types.SingleQuoteBold,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"*content*"),
	Entry("double quote bold text",
		types.QuotedText{
			Kind: types.DoubleQuoteBold,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"**content**"),
	Entry("single quote italic text",
		types.QuotedText{
			Kind: types.SingleQuoteItalic,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"_content_"),
	Entry("double quote italic text",
		types.QuotedText{
			Kind: types.DoubleQuoteItalic,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"__content__"),
	Entry("single quote monospace text",
		types.QuotedText{
			Kind: types.SingleQuoteMonospace,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"`content`"),
	Entry("double quote monospace text",
		types.QuotedText{
			Kind: types.DoubleQuoteMonospace,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"``content``"),
	Entry("single quote marked text",
		types.QuotedText{
			Kind: types.SingleQuoteMarked,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"#content#"),
	Entry("double quote marked text",
		types.QuotedText{
			Kind: types.DoubleQuoteMarked,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"##content##"),
	Entry("single quote subscript text",
		types.QuotedText{
			Kind: types.SingleQuoteSubscript,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"~content~"),
	Entry("single quote superscript text",
		types.QuotedText{
			Kind: types.SingleQuoteSuperscript,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"^content^"),
	// quoted string
	Entry("single quoted string",
		types.QuotedString{
			Kind: types.SingleQuote,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"`'content'`"),
	Entry("double quoted string",
		types.QuotedString{
			Kind: types.DoubleQuote,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"`\"content\"`"),
	// inline passthrough
	Entry("singleplus inline passthrough",
		types.InlinePassthrough{
			Kind: types.SinglePlusPassthrough,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"+content+"),
	Entry("tripleplus inline passthrough",
		types.InlinePassthrough{
			Kind: types.TriplePlusPassthrough,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"+++content+++"),
	Entry("macro inline passthrough",
		types.InlinePassthrough{
			Kind: types.PassthroughMacro,
			Elements: []interface{}{
				types.StringElement{
					Content: "content",
				},
			},
		},
		"pass:[content]"),
	// special characters
	Entry("special character",
		types.SpecialCharacter{
			Name: "<",
		},
		"<"),
	// attribute substitution
	Entry("attribute substitution",
		types.AttributeSubstitution{
			Name: "cookie",
		},
		"{cookie}"),
	// mixins
	Entry("mixins",
		types.QuotedText{
			Kind: types.SingleQuoteBold,
			Elements: []interface{}{
				types.StringElement{
					Content: "some ",
				},
				types.QuotedString{
					Kind: types.DoubleQuote,
					Elements: []interface{}{
						types.StringElement{
							Content: "content",
						},
						types.SpecialCharacter{
							Name: "<",
						},
						types.SpecialCharacter{
							Name: ">",
						},
					},
				},
				types.StringElement{
					Content: " ",
				},
				types.AttributeSubstitution{
					Name: "here",
				},
			},
		},
		"*some `\"content<>\"` {here}*"),
)
