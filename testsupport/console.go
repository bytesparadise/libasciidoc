package testsupport

import (
	"bytes"
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

// VerifyConsoleOutput verifies that the readable 'console' contains the exact given msg
func VerifyConsoleOutput(console Readable, msg string) { // TODO also add a param for the log level associated with the msg
	GinkgoT().Logf(console.String())
	out := make(map[string]interface{})
	// TODO: wrap the content of 'console` in a JSON document to avoid parsing error if there are multiple lines
	err := json.Unmarshal(console.Bytes(), &out)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(out["level"]).Should(Equal("error"))
	Expect(out["msg"]).Should(Equal(msg))
}

// ConfigureLogger configures the logger to write to a `Readable`.
// Also returns a func that can be used to reset the logger at the
// end of the test.
func ConfigureLogger() (Readable, func()) {
	fmtr := log.StandardLogger().Formatter
	level := log.StandardLogger().Level
	buf := bytes.NewBuffer(nil)
	log.SetLevel(log.WarnLevel)
	log.SetOutput(buf)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: true,
	})
	return buf, func() {
		log.SetOutput(os.Stdout)
		log.SetFormatter(fmtr)
		log.SetLevel(level)
	}
}

// Readable an interface for types which can return strings or arrays of bytes
type Readable interface {
	Bytes() []byte
	String() string
}
