package html5_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"

	_ "github.com/bytesparadise/libasciidoc/testsupport"
)

func TestHtml5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html5 Suite")
}
