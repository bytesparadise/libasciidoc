package types

import (
	"github.com/stretchr/testify/require"

	_ "github.com/bytesparadise/libasciidoc/test"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("normalize string", func() {
	It("hello", func() {
		verify(GinkgoT(), "hello", "hello")
	})
	It("héllo with an accent", func() {
		verify(GinkgoT(), "_héllo_1_2_and_3_spaces", " héllo 1.2   and 3 Spaces")
	})
	It("a an accent and a swedish character", func() {
		verify(GinkgoT(), `a_à`, `A à ⌘`)
	})
	It("AŁA", func() {
		verify(GinkgoT(), `ała_0_1`, `AŁA 0.1 ?`)
	})
	It("it's  2 spaces, here !", func() {
		verify(GinkgoT(), `it_s_2_spaces_here`, `it's  2 spaces, here !`)
	})
})

func verify(t GinkgoTInterface, expected, input string) {
	t.Logf("Processing '%s'", input)
	result, err := ReplaceNonAlphanumerics("_")(input)
	require.Nil(t, err)
	t.Logf("Normalized result: '%s'", string(result))
	assert.Equal(t, expected, string(result))
}
