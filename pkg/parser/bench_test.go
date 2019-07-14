package parser_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
)

const (
	doc1line = `=== foo1
bar1`
	doc10lines = `=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar`
)

func BenchmarkParser(b *testing.B) {
	usecases := []struct {
		name    string
		content []byte
	}{
		{
			name:    "1 line",
			content: []byte(doc1line),
		},
		{
			name:    "10 lines",
			content: []byte(doc10lines),
		},
		{
			name:    "vert.x doc",
			content: load(b, "../../test/bench/vertx-examples.adoc"),
		},
	}
	for _, usecase := range usecases {
		name := usecase.name
		content := usecase.content
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, err := parser.Parse(name, content)
				if err != nil {
					b.Error(err)
				}
			}
		})
	}

}

func load(b *testing.B, filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		b.Error(err)
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			b.Error(err2)
		}
	}()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		b.Error(err)
	}
	return content
}
