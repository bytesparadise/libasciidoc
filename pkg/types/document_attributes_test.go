package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document attributes", func() {

	Context("custom attributes", func() {

		It("normal value", func() {
			// given
			attributes := types.DocumentAttributes{}
			// when
			attributes.Add("foo", "bar")
			// then
			Expect("bar").To(Equal(attributes["foo"]))
		})

		It("pointer to value", func() {
			// given
			attributes := types.DocumentAttributes{}
			// when
			bar := "bar"
			attributes.Add("foo", &bar)
			// then
			Expect("bar").To(Equal(attributes["foo"]))
		})

		It("nil value", func() {
			// given
			attributes := types.DocumentAttributes{}
			// when
			attributes.Add("foo", nil)
			// then
			_, found := attributes["foo"]
			Expect(found).To(BeFalse())
		})

	})
})
