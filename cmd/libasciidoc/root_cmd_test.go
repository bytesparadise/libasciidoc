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

	It("render with STDOUT output", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-o", "-", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		require.NoError(GinkgoT(), err)
		require.NotEmpty(GinkgoT(), buf)
	})

	It("render with file output", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"test/test.adoc"})
		// when
		err := root.Execute()
		// then
		require.NoError(GinkgoT(), err)
		content, err := ioutil.ReadFile("test/test.html")
		require.NoError(GinkgoT(), err)
		require.NotEmpty(GinkgoT(), content)
	})

	It("fail to parse bad log level", func() {
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

	It("render without header/footer", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "-o", "-", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		require.NoError(GinkgoT(), err)
		require.NotEmpty(GinkgoT(), buf)
		Expect(buf.String()).ToNot(ContainSubstring(`<div id="footer">`))
	})

	It("process stdin", func() {
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
		root.SetArgs([]string{})
		// when
		err = root.Execute()
		//then
		GinkgoT().Logf("command output: %v", buf.String())
		Expect(buf.String()).To(ContainSubstring(content))
		require.NoError(GinkgoT(), err)
		require.NotEmpty(GinkgoT(), buf)
	})

	It("render multiple files", func() {
		// given
		root := main.NewRootCmd()
		root.SetArgs([]string{"-s", "test/admonition.adoc", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		require.NoError(GinkgoT(), err)
	})

	It("when rendering multiple files, return last error", func() {
		// given
		root := main.NewRootCmd()
		root.SetArgs([]string{"-s", "test/doesnotexist.adoc", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		require.Error(GinkgoT(), err)
	})
})

func TestRootCommand(t *testing.T) {

}
