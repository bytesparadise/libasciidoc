package libasciidoc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLibasciidoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Libasciidoc Suite")
}
