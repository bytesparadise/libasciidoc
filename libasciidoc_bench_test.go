package libasciidoc_test

import (
	"testing"

	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/stretchr/testify/require"
)

func BenchmarkLibasciidoc(b *testing.B) {
	filename := "./test/bench/mocking.adoc"
	for i := 0; i < b.N; i++ {
		_, err := testsupport.RenderHTML5Document(filename)
		require.NoError(b, err)
	}
}
