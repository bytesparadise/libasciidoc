package html5_test

import (
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint

	"testing"

	_ "github.com/bytesparadise/libasciidoc/testsupport"
)

func TestHtml5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html5 Suite")
}
