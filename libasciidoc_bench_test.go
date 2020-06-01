package libasciidoc_test

import (
	"testing"

	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/stretchr/testify/require"
)

func BenchmarkLibasciidoc(b *testing.B) {
	filename := "./test/bench/mocking.adoc"
	_, err := testsupport.RenderHTML5Document(filename)
	require.NoError(b, err)
}
