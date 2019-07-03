package parser_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
)

func BenchmarkParser1(b *testing.B) {
	source := `=== foo1
bar1`

	for n := 0; n < b.N; n++ {
		err := parseReader(source)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParser10(b *testing.B) {
	// given
	source := `=== foo
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
	for n := 0; n < b.N; n++ {
		err := parseReader(source)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkParserFile(b *testing.B) {
	f, err := os.Open("../../test/bench/vertx-examples.adoc")
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

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := parser.Parse("vert.x samples", content)
		if err != nil {
			b.Error(err)
		}
	}
}

func parseReader(content string) error {
	reader := strings.NewReader(content)
	// stats := parser.Stats{}
	// opts := []parser.Option{parser.AllowInvalidUTF8(false), parser.Statistics(&stats, "no match")}
	opts := []parser.Option{parser.AllowInvalidUTF8(false)}
	// if os.Getenv("DEBUG") == "true" {
	// 	opts = append(opts, parser.Debug(true))
	// }
	_, err := parser.ParseDocument("", reader, opts...) //, Debug(true))
	return err
}
