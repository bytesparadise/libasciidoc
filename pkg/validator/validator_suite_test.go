package validator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validator Suite")
}
