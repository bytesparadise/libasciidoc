package compat_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bytesparadise/libasciidoc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
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
	DescribeTable("supported", compare, entries("fixtures/supported/*.adoc")...)
})

func compare(file string) {
	// set logger to a minimal verbose level, then restore at its initial level afterwards
	level := log.GetLevel()
	log.SetLevel(log.WarnLevel)
	defer func() {
		log.SetLevel(level)
	}()
	actual, err := convert(file)
	Expect(err).ShouldNot(HaveOccurred())
	expected, err := getGoldenFile(file)
	Expect(err).ShouldNot(HaveOccurred())
	// if tests are executed on windows platform and git 'autocrlf' is set to 'true',
	// then we need to remove the `\r` characters that were added in the 'expected'
	// source at the time of the checkout
	if runtime.GOOS == "windows" && autocrlf == "true" {
		expected = strings.Replace(expected, "\r", "", -1)
	}
	// compare actual vs reference
	Expect(actual).To(Equal(expected))
}

const adocExt = ".adoc"

func entries(pattern string) []TableEntry {
	files, _ := filepath.Glob(pattern)
	result := make([]TableEntry, len(files))
	for i, file := range files {
		result[i] = Entry(file, file)
	}
	return result
}

func convert(sourcePath string) (string, error) {
	// generate the HTML output
	buff := bytes.NewBuffer(nil)
	_, err := libasciidoc.ConvertFileToHTML(context.Background(), sourcePath, buff)
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}

func getGoldenFile(sourcePath string) (string, error) {
	// retrieve the reference document
	goldPath := sourcePath[:len(sourcePath)-len(adocExt)] + ".html"
	content, err := ioutil.ReadFile(goldPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
