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
func VerifyConsoleOutput(console Readable, msg string) {
	GinkgoT().Logf(console.String())
	out := make(map[string]interface{})
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

	buf := bytes.NewBuffer(nil)
	log.SetOutput(buf)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: true,
	})
	return buf, func() {
		log.SetOutput(os.Stdout)
		log.SetFormatter(fmtr)
	}
}

// Readable an interface for types which can return strings or arrays of bytes
type Readable interface {
	Bytes() []byte
	String() string
}
