package main_test

import (
	"bytes"
	"testing"

	main "github.com/bytesparadise/libasciidoc/cmd/libasciidoc"
	"github.com/stretchr/testify/require"
)

import . "github.com/onsi/ginkgo"

var _ = Describe("root cmd", func() {

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
})

func TestRootCommand(t *testing.T) {

}
