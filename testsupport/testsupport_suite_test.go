package testsupport_test

import (
	"testing"

	_ "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestsupport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testsupport Suite")
}
