package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega" // nolint:golint
)

var _ = DescribeTable("'FileLocation' pattern",
	func(filename string, expected interface{}) {
		reader := strings.NewReader(filename)
		actual, err := parser.ParseReader(filename, reader, parser.Entrypoint("FileLocation"))
		Expect(err).ToNot(HaveOccurred())
		// GinkgoT().Log("actual result: %s", spew.Sdump(actual))
		// GinkgoT().Log("expected result: %s", spew.Sdump(expected))
		Expect(actual).To(Equal(expected))
	},
	Entry("'chapter-a.adoc'", "chapter-a.adoc", &types.Location{
		Path: "chapter-a.adoc",
	}),
	Entry("'chapter_a.adoc'", "chapter_a.adoc", &types.Location{
		Path: "chapter_a.adoc",
	}),
	Entry("'../../test/includes/chapter_a.adoc'", "../../test/includes/chapter_a.adoc", &types.Location{
		Path: "../../test/includes/chapter_a.adoc",
	}),
	Entry("'{includedir}/chapter-{foo}.adoc'", "{includedir}/chapter-{foo}.adoc", &types.Location{
		Path: "{includedir}/chapter-{foo}.adoc", // attribute substitutions are treared as part of the string element
	}),
	Entry("'{scheme}://{path}'", "{scheme}://{path}", &types.Location{
		Path: "{scheme}://{path}",
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
