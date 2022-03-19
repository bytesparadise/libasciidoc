package test_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2" //nolint:golint
	. "github.com/onsi/gomega"    //nolint:golint

	_ "github.com/bytesparadise/libasciidoc/testsupport"
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}
