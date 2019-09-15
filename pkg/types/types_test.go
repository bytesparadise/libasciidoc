package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("line ranges", func() {

	Context("single range", func() {
		ranges := newLineRanges(
			types.LineRange{Start: 2, End: 4},
		)

		It("should not match line 1", func() {
			Expect(ranges.Match(1)).To(BeFalse())
		})

		It("should match line 2", func() {
			Expect(ranges.Match(2)).To(BeTrue())
		})

		It("should not match line 5", func() {
			Expect(ranges.Match(1)).To(BeFalse())
		})
	})

	Context("multiple ranges", func() {

		ranges := newLineRanges(
			types.LineRange{Start: 1, End: 1},
			types.LineRange{Start: 3, End: 4},
			types.LineRange{Start: 6, End: -1},
		)

		It("should match line 1", func() {
			Expect(ranges.Match(1)).To(BeTrue())
		})

		It("should not match line 2", func() {
			Expect(ranges.Match(2)).To(BeFalse())
		})

		It("should match line 6", func() {
			Expect(ranges.Match(6)).To(BeTrue())
		})

		It("should match line 100", func() {
			Expect(ranges.Match(100)).To(BeTrue())
		})
	})

})

func newLineRanges(values ...interface{}) types.LineRanges {
	return types.NewLineRanges(values...)
}

// var _ = Describe("raw section title offset", func() {

// 	It("should apply relative positive offset", func() {
// 		actual := types.RawSectionTitlePrefix{
// 			Level:  []byte("=="),
// 			Spaces: []byte(" "),
// 		}
// 		expected := "=== "
// 		verifyLevelOffset(expected, actual, "+1")
// 	})
// })

// func verifyLevelOffset(expectation string, actual types.RawSectionTitlePrefix, levelOffset string) {
// 	result, err := actual.Bytes(levelOffset)
// 	require.NoError(GinkgoT(), err)
// 	assert.EqualValues(GinkgoT(), expectation, result)
// }

var _ = Describe("file inclusions", func() {

	DescribeTable("check asciidoc file",
		func(path string, expectation bool) {
			Expect(types.IsAsciidoc(path)).To(Equal(expectation))
		},
		Entry("foo.adoc", "foo.adoc", true),
		Entry("foo.asc", "foo.asc", true),
		Entry("foo.ad", "foo.ad", true),
		Entry("foo.asciidoc", "foo.asciidoc", true),
		Entry("foo.txt", "foo.txt", true),
		Entry("foo.csv", "foo.csv", false),
		Entry("foo.go", "foo.go", false),
	)
})

var _ = Describe("Location resolution", func() {

	attrs := types.DocumentAttributes{
		"includedir": "includes",
		"foo":        "bar",
	}
	DescribeTable("resolve URL",
		func(location types.Location, expectation string) {
			f := types.FileInclusion{
				Location: location,
			}
			Expect(f.Location.Resolve(attrs)).To(Equal(expectation))
		},
		Entry("includes/file.ext", types.Location{
			types.StringElement{Content: "includes/file.ext"},
		}, "includes/file.ext"),
		Entry("./{includedir}/file.ext", types.Location{
			types.StringElement{Content: "./"},
			types.DocumentAttributeSubstitution{Name: "includedir"},
			types.StringElement{Content: "/file.ext"},
		}, "./includes/file.ext"),
		Entry("./{unknown}/file.ext", types.Location{
			types.StringElement{Content: "./"},
			types.DocumentAttributeSubstitution{Name: "unknown"},
			types.StringElement{Content: "/file.ext"},
		}, "./{unknown}/file.ext"),
	)
})
