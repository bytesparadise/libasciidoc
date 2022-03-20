// +build bench

package libasciidoc_test

import (
	"strings"
	"testing"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/stretchr/testify/require"
)

func BenchmarkRealDocumentProcessing(b *testing.B) {
	b.Run("demo.adoc", processDocument("./test/compat/demo.adoc"))
	b.Run("vertx-examples.adoc", processDocument("./test/bench/vertx-examples.adoc"))
	b.Run("mocking.adoc", processDocument("./test/bench/mocking.adoc"))
}

func processDocument(filename string) func(b *testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out := &strings.Builder{}
			_, err := libasciidoc.ConvertFile(out,
				configuration.NewConfiguration(
					configuration.WithFilename(filename),
					configuration.WithCSS([]string{"path/to/style.css"}),
					configuration.WithHeaderFooter(true)))
			require.NoError(b, err)
		}
	}
}
