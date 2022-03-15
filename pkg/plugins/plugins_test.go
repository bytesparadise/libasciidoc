package plugins_test

import (
	"fmt"

//	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/pkg/plugins"
//	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("plugins", func() {

	Context("loading from golang plugin", func() {

		It("loads PreRender from a golang plugin", func() {
      result, err := plugins.LoadPlugins([]string{"../../test/plugins/plugin.so"})
			Expect(err).NotTo(HaveOccurred())
      fmt.Println(result)
		})
	})
})
