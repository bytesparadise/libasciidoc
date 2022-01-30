package libasciidoc_test

import (
	"testing"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golint
)

func TestLibasciidoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Libasciidoc Suite")
}
