package testsupport_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

var _ = Describe("console assertions", func() {

	console := `{"level":"debug","msg":"processing foo.adoc: ----\ninclude::../../test/includes/unknown.adoc[leveloffset=+1]\n----"}
{"level":"debug","msg":"initializing a new PreflightDocument with 1 block element(s)"}
{"level":"debug","msg":"parsing '../../test/includes/unknown.adoc'..."}
{"error":"open unknown.adoc: no such file or directory","level":"error","msg":"failed to include '../../test/includes/unknown.adoc'"}
{"level":"debug","msg":"restoring current working dir to: github.com/bytesparadise/libasciidoc/pkg/parser"}`

	It("should find expected level/message", func() {
		// given
		matcher := testsupport.ContainMessageWithLevel(log.ErrorLevel, "failed to include '../../test/includes/unknown.adoc'")
		// when
		result, err := matcher.Match(strings.NewReader(console))
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not find expected level/message with wrong level", func() {
		// given
		matcher := testsupport.ContainMessageWithLevel(log.WarnLevel, "failed to include '../../test/includes/unknown.adoc'") // wrong level
		// when
		result, err := matcher.Match(strings.NewReader(console))
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
	})

	It("should not find expected level/message with wrong msg", func() {
		// given
		matcher := testsupport.ContainMessageWithLevel(log.ErrorLevel, "foo") // unknown message
		// when
		result, err := matcher.Match(strings.NewReader(console))
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
	})

})
