package types_test

import (
	. "github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Document Attributes", func() {

	It("normal value", func() {
		// given
		attributes := DocumentAttributes{}
		// when
		attributes.Add("foo", "bar")
		// then
		assert.Equal(GinkgoT(), "bar", attributes["foo"])
	})

	It("pointer to value", func() {
		// given
		attributes := DocumentAttributes{}
		// when
		bar := "bar"
		attributes.Add("foo", &bar)
		// then
		assert.Equal(GinkgoT(), "bar", attributes["foo"])
	})

	It("nil value", func() {
		// given
		attributes := DocumentAttributes{}
		// when
		attributes.Add("foo", nil)
		// then
		_, found := attributes["foo"]
		assert.False(GinkgoT(), found)
	})
})
