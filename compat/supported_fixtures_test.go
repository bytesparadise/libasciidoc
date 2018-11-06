package compat_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

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
	// compare actual vs reference
	GinkgoT().Logf("actual:\n%v\nexpected:%v", []byte(actual), []byte(expected))
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
