package types

import (
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("convert to inline elements", func() {

	It("inline content without trailing spaces", func() {
		source := []interface{}{
			StringElement{Content: "hello"},
			StringElement{Content: "world"},
		}
		expected := []interface{}{
			StringElement{Content: "helloworld"},
		}
		// when
		result := Merge(source...)
		// then
		Expect(result).To(Equal(expected))
	})
	It("inline content with trailing spaces", func() {
		source := []interface{}{
			StringElement{Content: "hello, "},
			StringElement{Content: "world   "},
		}
		expected := []interface{}{
			StringElement{Content: "hello, world   "},
		}
		// when
		result := Merge(source...)
		// then
		Expect(result).To(Equal(expected))
	})
})
