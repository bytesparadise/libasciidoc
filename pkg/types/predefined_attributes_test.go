package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega" //nolint golint
)

var _ = DescribeTable("predefined attributes",
	func(code, rendered string) {
		Expect(types.Predefined[code]).To(Equal(rendered))
	},
	Entry("sp", "sp", " "),
	Entry("blank", "blank", ""),
	Entry("empty", "empty", ""),
	Entry("nbsp", "nbsp", "\u00a0"),
	Entry("zwsp", "zwsp", "\u200b"),
	Entry("wj", "wj", "\u2060"),
	Entry("apos", "apos", "&#39;"),
	Entry("quot", "quot", "&#34;"),
	Entry("lsquo", "lsquo", "\u2018"),
	Entry("rsquo", "rsquo", "\u2019"),
	Entry("ldquo", "ldquo", "\u201c"),
	Entry("rdquo", "rdquo", "\u201d"),
	Entry("deg", "deg", "\u00b0"),
	Entry("plus", "plus", "&#43;"),
	Entry("brvbar", "brvbar", "\u00a6"),
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
