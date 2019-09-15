package testsupport

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ContainMessageWithLevel a custom Matcher to verify that a message with at a given level was logged
func ContainMessageWithLevel(level log.Level, msg string) types.GomegaMatcher {
	return &containMessageMatcher{
		level: level,
		msg:   msg,
	}
}

type containMessageMatcher struct {
	level log.Level
	msg   string
}

func (m *containMessageMatcher) Match(actual interface{}) (success bool, err error) {
	console, ok := actual.(io.Reader)
	if !ok {
		return false, errors.Errorf("ContainMessageWithLevel matcher expects an io.Reader (actual: %T)", actual)
	}
	scanner := bufio.NewScanner(console)
	for scanner.Scan() {
		out := make(map[string]interface{})
		err := json.Unmarshal(scanner.Bytes(), &out)
		if err != nil {
			return false, errors.Wrapf(err, "failed to decode console line")
		}
		if level, ok := out["level"].(string); !ok || level != m.level.String() {
			continue
		}
		if msg, ok := out["msg"].(string); !ok || msg != m.msg {
			continue
		}
		// match found
		return true, nil
	}
	// no match found
	return false, nil
}

func (m *containMessageMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected console to contain message '%s' with level '%v'", m.msg, m.level)
}

func (m *containMessageMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected console not to contain message '%s' with level '%v'", m.msg, m.level)
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
