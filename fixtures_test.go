// +build fixtures

package libasciidoc_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc"
	. "github.com/onsi/ginkgo"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const adocExt = ".adoc"

var _ = Describe("fixtures", func() {
	matches, err := filepath.Glob("test/fixtures/*" + adocExt)
	require.NoError(GinkgoT(), err)

	for _, input := range matches {
		Context("["+input+"]", func() {
			in := input
			It("render HTML", func() {
				verifyHTMLFixture(GinkgoT(), in)
			})
		})
	}
})

func verifyHTMLFixture(t GinkgoTInterface, sourcePath string) {
	w := bytes.NewBuffer(nil)
	f, err := os.Open(sourcePath)
	require.NoError(t, err)

	_, err = libasciidoc.ConvertToHTML(context.Background(), f, w)
	require.NoError(t, err)
	result := w.String()

	goldPath := sourcePath[:len(sourcePath)-len(adocExt)] + ".html"
	b, err := ioutil.ReadFile(goldPath)
	require.NoError(t, err)
	expect := string(b)

	assert.Equal(t, expect, result, showDiff(goldPath, expect, result))
}

func showDiff(name, expect, actual string) string {
	d, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(expect),
		B:        difflib.SplitLines(actual),
		FromFile: "expect (" + name + ")",
		ToFile:   "actual",
		Context:  1,
	})
	return d
}
