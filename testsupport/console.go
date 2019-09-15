package testsupport

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

// VerifyConsoleOutput verifies that the readable 'console' contains the exact given msg
func VerifyConsoleOutput(console io.Reader, expectedLevel log.Level, expectedMsg string) { // TODO also add a param for the log level associated with the msg
	// fmt.Printf("console: %s\n", console.String())
	// read each line of the console as a separate document
	scanner := bufio.NewScanner(console)
	for scanner.Scan() {
		out := make(map[string]interface{})
		err := json.Unmarshal(scanner.Bytes(), &out)
		Expect(err).ShouldNot(HaveOccurred())
		GinkgoT().Logf("out: %v", out)
		if level, ok := out["level"].(string); !ok || level != expectedLevel.String() {
			continue
		}
		if msg, ok := out["msg"].(string); !ok || msg != expectedMsg {
			continue
		}
		// match found
		return
	}
	c, err := ioutil.ReadAll(console)
	Expect(err).ShouldNot(HaveOccurred())
	GinkgoT().Logf("console: %v", string(c))
	// no match found
	Fail("no match found in the console")
}

// ConfigureLogger configures the logger to write to a `Readable`.
// Also returns a func that can be used to reset the logger at the
// end of the test.
func ConfigureLogger() (io.Reader, func()) {
	fmtr := log.StandardLogger().Formatter
	// level := log.StandardLogger().Level
	buf := bytes.NewBuffer(nil)
	// log.SetLevel(log.WarnLevel)
	log.SetOutput(buf)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: true,
	})
	return buf, func() {
		log.SetOutput(os.Stdout)
		log.SetFormatter(fmtr)
		// log.SetLevel(level)
	}
}
