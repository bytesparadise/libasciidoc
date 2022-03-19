package sgml_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSgml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sgml Suite")
}
