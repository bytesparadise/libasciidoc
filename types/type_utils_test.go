package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		verify(t, "hello", "hello")
	})
	t.Run("héllo with an accent", func(t *testing.T) {
		verify(t, "_héllo_1_2_and_3_spaces", " héllo 1.2   and 3 Spaces")
	})
	t.Run("a an accent and a swedish character", func(t *testing.T) {
		verify(t, `a_à`, `A à ⌘`)
	})
	t.Run("AŁA", func(t *testing.T) {
		verify(t, `ała_0_1`, `AŁA 0.1 ?`)
	})
	t.Run("it's  2 spaces, here !", func(t *testing.T) {
		verify(t, `it_s_2_spaces_here`, `it's  2 spaces, here !`)
	})
}

func verify(t *testing.T, expected, input string) {
	t.Logf("Processing '%s'", input)
	result, err := ReplaceNonAlphanumerics("_")(input)
	require.Nil(t, err)
	t.Logf("Normalized result: '%s'", string(result))
	assert.Equal(t, expected, string(result))
}
