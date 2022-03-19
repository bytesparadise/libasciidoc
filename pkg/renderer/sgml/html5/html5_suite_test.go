package html5_test

import (
	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golintt

	"testing"

	_ "github.com/bytesparadise/libasciidoc/testsupport"
)

func TestHtml5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html5 Suite")
}
