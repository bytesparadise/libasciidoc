package main_test

import (
	"bytes"

	main "github.com/bytesparadise/libasciidoc/cmd/libasciidoc"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/require"
)

var _ = Describe("version cmd", func() {

	It("ok", func() {
		// given
		versionCmd := main.NewVersionCmd()
		buf := new(bytes.Buffer)
		versionCmd.SetOutput(buf)
		versionCmd.SetArgs([]string{})
		// when
		err := versionCmd.Execute()
		// then
		require.NoError(GinkgoT(), err)
		require.NotEmpty(GinkgoT(), buf.String())
	})

})
