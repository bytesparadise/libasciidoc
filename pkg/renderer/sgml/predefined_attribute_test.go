package sgml

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega" // nolint:golint
)

var _ = DescribeTable("predefined attributes",
	func(code, rendered string) {
		Expect(predefinedAttribute(code)).To(Equal(rendered))
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
	Entry("amp", "amp", "&"),
	Entry("lt", "lt", "<"),
	Entry("gt", "gt", ">"),
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
