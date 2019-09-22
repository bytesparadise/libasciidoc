package testsupport_test

import (
	"fmt"
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
		// also verify the messages

	})

	It("should not find expected level/message with wrong msg", func() {
		// given
		matcher := testsupport.ContainMessageWithLevel(log.ErrorLevel, "foo") // unknown message
		// when
		result, err := matcher.Match(strings.NewReader(console))
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		// also verify the messages
		Expect(matcher.FailureMessage(strings.NewReader(console))).To(Equal(fmt.Sprintf("expected console to contain message '%s' with level '%v'", "foo", log.ErrorLevel)))
		Expect(matcher.NegatedFailureMessage(strings.NewReader(console))).To(Equal(fmt.Sprintf("expected console not to contain message '%s' with level '%v'", "foo", log.ErrorLevel)))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.ContainMessageWithLevel(log.ErrorLevel, "foo") // unknown message
		// when
		result, err := matcher.Match(1) // not a reader
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("ContainMessageWithLevel matcher expects an io.Reader (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
