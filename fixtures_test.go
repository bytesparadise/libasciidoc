// +build fixtures

package libasciidoc_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc"
	. "github.com/onsi/ginkgo"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const adocExt = ".adoc"

var _ = Describe("fixtures", func() {
	matches, err := filepath.Glob("test/fixtures/*" + adocExt)
	require.NoError(GinkgoT(), err)

	for _, input := range matches {
		in := input
		It("render HTML "+in, func() {
			logrus.SetLevel(logrus.FatalLevel)
			verifyHTMLFixture(GinkgoT(), in)
		})
	}
})

func verifyHTMLFixture(t GinkgoTInterface, sourcePath string) {
	w := bytes.NewBuffer(nil)
	_, err := libasciidoc.ConvertFileToHTML(context.Background(), sourcePath, w)
	require.NoError(t, err)
	result := w.String()

	goldPath := sourcePath[:len(sourcePath)-len(adocExt)] + ".html"
	b, err := ioutil.ReadFile(goldPath)
	require.NoError(t, err)
	expect := string(b)

	assert.Equal(t, expect, result, showDiff(expect, result))
}

func showDiff(expect, actual string) string {
	dmp := diffmatchpatch.New()
	// prepend normal color code to clear ginkgo colorization.
	return "\x1b[0m" + dmp.DiffPrettyText(dmp.DiffMain(actual, expect, true))
}
