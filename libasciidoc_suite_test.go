package libasciidoc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	_ "github.com/bytesparadise/libasciidoc/log"
)

func TestLibasciidoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Libasciidoc Suite")
}
