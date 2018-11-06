package compat_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCompat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Compat Suite")
}
