package testsupport_test

import (
	"testing"

	_ "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

func TestTestsupport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testsupport Suite")
}
