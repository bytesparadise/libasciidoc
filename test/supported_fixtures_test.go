package test_test

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega" // nolint:golintt
	log "github.com/sirupsen/logrus"
)

var autocrlf string

func init() {
	cmd := exec.Command("git", "config", "--get", "core.autocrlf")
	// Attach buffer to command
	cmdOutput := bytes.NewBuffer(nil)
	cmd.Stdout = cmdOutput
	// execute command
	err := cmd.Run()
	if err != nil {
		log.Errorf("failed to check the `git config --get autocrlf': %v", err)
		return
	}
	autocrlf = strings.Trim(cmdOutput.String(), "\r\n")
	log.Warnf("git autocrlf='%s'", autocrlf)

}

var _ = Describe("fixtures", func() {

	// verifies that all files in the `supported` subfolder match their sibling golden file
	DescribeTable("supported", compare, entries("compat/*.adoc")...)
})

func compare(filename string) {
	actual := &strings.Builder{}
	_, err := libasciidoc.ConvertFile(actual, configuration.NewConfiguration(
		configuration.WithFilename(filename),
		configuration.WithBackEnd("html5"),
		configuration.WithAttribute("libasciidoc-version", "0.7.0"),
	))
	Expect(err).NotTo(HaveOccurred())
	// retrieve the reference document
	path := strings.TrimSuffix(filename, ".adoc") + ".html"
	content, err := ioutil.ReadFile(path)
	Expect(err).NotTo(HaveOccurred())
	expected := string(content)
	// if tests are executed on windows platform and git 'autocrlf' is set to 'true',
	// then we need to remove the `\r` characters that were added in the 'expected'
	// source at the time of the checkout
	if runtime.GOOS == "windows" && autocrlf == "true" {
		expected = strings.Replace(expected, "\r", "", -1)
	}
	// compare actual vs reference
	Expect(actual.String()).To(MatchHTML(expected))
}

func entries(pattern string) []TableEntry {
	files, _ := filepath.Glob(pattern)
	result := make([]TableEntry, 0, len(files))
	for _, file := range files {
		if file == "compat/include.adoc" {
			// skip
			continue
		}
		result = append(result, Entry(file, file))
	}
	return result
}
