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

	. "github.com/onsi/ginkgo/v2"
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

var _ = DescribeTable("compat", args("compat/*.adoc")...)

func args(pattern string) []interface{} {
	result := []interface{}{}
	result = append(result, compare)
	files, _ := filepath.Glob(pattern)
	for _, file := range files {
		if filepath.Base(file) == "include.adoc" {
			// skip
			continue
		}
		result = append(result, Entry(file, file))
	}
	return result

}

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
