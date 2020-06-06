package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega" //nolint golint
)

var _ = DescribeTable("document attribute overrides",
	func(key string, expectedValue string, expectedFound bool) {
		// given
		attributes := types.AttributesWithOverrides{
			Content: map[string]interface{}{
				"normal":   "ok",
				"override": "ok, too",
			},
			Overrides: map[string]string{
				"foo":      "cheesecake",
				"!bar":     "",
				"baz":      "",
				"override": "overridden",
			},
		}
		// when
		value, found := attributes.GetAsString(key)
		// then
		Expect(found).To(Equal(expectedFound))
		Expect(value).To(Equal(expectedValue))
	},
	Entry("normal", "normal", "ok", true),
	Entry("override", "override", "overridden", true), // entry is overridden
	Entry("foo", "foo", "cheesecake", true),
	Entry("!bar", "bar", "", false), // entry is reset
	Entry("baz", "baz", "", true),   // entry exists but its value is empty
)

var _ = DescribeTable("document attribute overrides with default",
	func(key string, expectedValue string) {
		// given
		attributes := types.AttributesWithOverrides{
			Content: map[string]interface{}{
				"normal":   "ok",
				"override": "ok, too",
			},
			Overrides: map[string]string{
				"foo":      "cheesecake",
				"!bar":     "",
				"baz":      "",
				"override": "overridden",
			},
		}
		// when
		value := attributes.GetAsStringWithDefault(key, "default")
		// then
		Expect(value).To(Equal(expectedValue))
	},
	Entry("normal", "normal", "ok"),
	Entry("override", "override", "overridden"), // entry is overridden
	Entry("foo", "foo", "cheesecake"),
	Entry("!bar", "bar", "default"), // entry is reset, default is returned
	Entry("baz", "baz", ""),         // entry exists but its value is empty
)
