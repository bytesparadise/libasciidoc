package main_test

import (
	"bytes"
	"io/ioutil"

	main "github.com/bytesparadise/libasciidoc/cmd/libasciidoc"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("root cmd", func() {

	It("render with STDOUT output", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-o", "-", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(buf.String()).ToNot(BeEmpty())
	})

	It("render with file output", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
		content, err := ioutil.ReadFile("test/test.html")
		Expect(err).ToNot(HaveOccurred())
		Expect(content).ToNot(BeEmpty())
	})

	It("fail to parse bad log level", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"--log", "debug1", "-s", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		GinkgoT().Logf("command output: %v", buf.String())
		Expect(err).To(HaveOccurred())
	})

	It("render without header/footer", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "-o", "-", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(buf.String()).ToNot(BeEmpty())
		Expect(buf.String()).ToNot(ContainSubstring(`<div id="footer">`))
	})

	It("render with attribute set", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "-o", "-", "-afoo1=bar1", "-afoo2=bar2", "test/doc_with_attributes.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(buf.String()).ToNot(BeEmpty())
		Expect(buf.String()).To(Equal(`<div class="paragraph">
<p>bar1 and bar2</p>
</div>
`))
	})

	It("render with attribute reset", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "-o", "-", "-afoo1=bar1", "-a!foo2", "test/doc_with_attributes.adoc"})
		// when
		err := root.Execute()
		// then
		GinkgoT().Logf("out: %v", buf.String())
		Expect(err).ToNot(HaveOccurred())
		Expect(buf.String()).ToNot(BeEmpty())
		// console output also includes a warning message
		Expect(buf.String()).To(Equal(`level=warning msg="unable to find attribute 'foo2'"
<div class="paragraph">
<p>bar1 and {foo2}</p>
</div>
`))
	})

	It("render multiple files", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "test/admonition.adoc", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
	})

	It("when rendering multiple files, return last error", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "test/doesnotexist.adoc", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).To(HaveOccurred())
	})

	It("show help when executed with no arg", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
	})

})
