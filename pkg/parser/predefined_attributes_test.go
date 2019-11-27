package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("document with attributes", func() {

	DescribeTable("predefined attributes",
		func(code, rendered string) {
			Expect(parser.Predefined[code]).To(Equal(rendered))
		},
		Entry("sp", "sp", " "),
		Entry("blank", "blank", ""),
		Entry("empty", "empty", ""),
		Entry("nbsp", "nbsp", "&#160;"),
		Entry("zwsp", "zwsp", "&#8203;"),
		Entry("wj", "wj", "&#8288;"),
		Entry("apos", "apos", "&#39;"),
		Entry("quot", "quot", "&#34;"),
		Entry("lsquo", "lsquo", "&#8216;"),
		Entry("rsquo", "rsquo", "&#8217;"),
		Entry("ldquo", "ldquo", "&#8220;"),
		Entry("rdquo", "rdquo", "&#8221;"),
		Entry("deg", "deg", "&#176;"),
		Entry("plus", "plus", "&#43;"),
		Entry("brvbar", "brvbar", "&#166;"),
		Entry("vbar", "vbar", "|"),
		Entry("amp", "amp", "&amp;"),
		Entry("lt", "lt", "&lt;"),
		Entry("gt", "gt", "&gt;"),
		Entry("startsb", "startsb", "["),
		Entry("endsb", "endsb", "]"),
		Entry("caret", "caret", "^"),
		Entry("asterisk", "asterisk", "*"),
		Entry("tilde", "tilde", "~"),
		Entry("backslash", "backslash", `\`),
		Entry("backtick", "backtick", "`"),
		Entry("two-colons", "two-colons", "::"),
		Entry("two-semicolons", "two-semicolons", ";"),
		Entry("cpp", "cpp", "C++"),
	)
})
