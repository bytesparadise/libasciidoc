package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
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

		ranges := types.NewLineRanges(
			types.LineRange{StartLine: 1, EndLine: 1},
			types.LineRange{StartLine: 3, EndLine: 4},
			types.LineRange{StartLine: 6, EndLine: -1},
		)

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
			ranges, _ := types.NewTagRanges(types.TagRange{
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
			ranges, _ := types.NewTagRanges(types.TagRange{
				Name:     "foo",
				Included: true,
			}, types.TagRange{
				Name:     "bar",
				Included: true,
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
				ranges, _ := types.NewTagRanges(types.TagRange{
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
				ranges, _ := types.NewTagRanges(types.TagRange{
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
				ranges, _ := types.NewTagRanges(types.TagRange{
					Name:     "**",
					Included: true,
				}, types.TagRange{
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
				ranges, _ := types.NewTagRanges(types.TagRange{
					Name:     "foo",
					Included: true,
				}, types.TagRange{
					Name:     "bar",
					Included: false,
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
				ranges, _ := types.NewTagRanges(types.TagRange{
					Name:     "*",
					Included: true,
				}, types.TagRange{
					Name:     "foo",
					Included: false,
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
				ranges, _ := types.NewTagRanges(types.TagRange{
					Name:     "**",
					Included: true,
				}, types.TagRange{
					Name:     "foo",
					Included: false,
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
				ranges, _ := types.NewTagRanges(types.TagRange{
					Name:     "**",
					Included: true,
				}, types.TagRange{
					Name:     "*",
					Included: false,
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
		_, err := types.NewTagRanges("foo", "bar")
		// then
		Expect(err).To(HaveOccurred())
	})

})

var _ = Describe("location resolution", func() {

	attrs := map[string]string{
		"includedir": "includes",
		"foo":        "bar",
	}
	DescribeTable("resolve URL",
		func(actual types.Location, expectation types.Location) {
			actual.Resolve(attrs)
			Expect(actual).To(Equal(expectation))
		},
		Entry("includes/file.ext",
			types.Location{
				Elements: []interface{}{
					types.StringElement{
						Content: "includes/file.ext",
					},
				},
			},
			types.Location{
				Elements: []interface{}{
					types.StringElement{
						Content: "includes/file.ext",
					},
				},
			}),
		Entry("./{includedir}/file.ext",
			types.Location{
				Elements: []interface{}{
					types.StringElement{
						Content: "./",
					},
					types.DocumentAttributeSubstitution{
						Name: "includedir",
					},
					types.StringElement{
						Content: "/file.ext",
					},
				},
			},
			types.Location{
				Elements: []interface{}{
					types.StringElement{
						Content: "./",
					},
					types.DocumentAttributeSubstitution{
						Name: "includedir",
					},
					types.StringElement{
						Content: "/file.ext",
					},
				},
			},
		),
		Entry("./{unknown}/file.ext",
			types.Location{
				Elements: []interface{}{
					types.StringElement{
						Content: "./",
					},
					types.DocumentAttributeSubstitution{
						Name: "unknown",
					},
					types.StringElement{
						Content: "/file.ext",
					},
				},
			},
			types.Location{
				Elements: []interface{}{
					types.StringElement{
						Content: "./",
					},
					types.DocumentAttributeSubstitution{
						Name: "unknown",
					},
					types.StringElement{
						Content: "/file.ext",
					},
				},
			},
		),
	)
})
