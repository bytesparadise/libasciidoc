package main_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	main "github.com/bytesparadise/libasciidoc/cmd/libasciidoc"
	"github.com/stretchr/testify/require"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("root cmd", func() {
	RegisterFailHandler(Fail)

	It("ok", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		require.NoError(GinkgoT(), err)
		require.NotEmpty(GinkgoT(), buf)
	})

	It("missing source flag", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		// when
		err := root.Execute()
		// then
		GinkgoT().Logf("command output: %v", buf.String())
		require.Error(GinkgoT(), err)
	})

	It("should fail to parse bad log level", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"--log", "debug1", "-s", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		GinkgoT().Logf("command output: %v", buf.String())
		require.Error(GinkgoT(), err)
	})

	It("should process stdin", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		content := "some content"
		tmpfile, err := ioutil.TempFile("", "example")
		if err != nil {
			log.Fatal(err)
		}

		defer os.Remove(tmpfile.Name()) // clean up

		if _, err := tmpfile.Write([]byte(content)); err != nil {
			log.Fatal(err)
		}
		tmpfile.Seek(0, 0)
		oldstdin := os.Stdin
		os.Stdin = tmpfile
		defer func() { os.Stdin = oldstdin }()

		root.SetArgs([]string{"-s", "-"})
		// when
		err = root.Execute()

		//then
		GinkgoT().Logf("command output: %v", buf.String())
		Expect(buf.String()).To(ContainSubstring(content))
		require.NoError(GinkgoT(), err)
		require.NotEmpty(GinkgoT(), buf)
	})
})

func TestRootCommand(t *testing.T) {

}
