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

// ConfigureLogger configures the logger to write to a `Readable`.
// Also returns a func that can be used to reset the logger at the
// end of the test.
func ConfigureLogger() (io.Reader, func()) {
	fmtr := log.StandardLogger().Formatter
	t := tee{
		buf: bytes.NewBuffer(nil),
		out: os.Stdout,
	}
	log.SetOutput(t)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: true,
	})
	return t, func() {
		log.SetOutput(os.Stdout)
		log.SetFormatter(fmtr)
	}
}

type tee struct {
	buf io.ReadWriter
	out io.Writer
}

func (t tee) Write(p []byte) (n int, err error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		n, err = t.out.Write(p)
		if err != nil {
			return n, err
		}
	}
	return t.buf.Write(p)
}

func (t tee) Read(p []byte) (n int, err error) {
	return t.buf.Read(p)
}

// ---------------------------
// ContainMessageWithLevel
// ---------------------------

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

func (m *containMessageMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected console to contain message '%s' with level '%v'", m.msg, m.level)
}

func (m *containMessageMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected console not to contain message '%s' with level '%v'", m.msg, m.level)
}

// ---------------------------
// ContainAnyMessageWithLevels
// ---------------------------

// ContainAnyMessageWithLevels a custom Matcher to verify that no message with the any of the given levels was logged
func ContainAnyMessageWithLevels(level log.Level, otherLevels ...log.Level) types.GomegaMatcher {
	return &containAnyMessageMatcher{
		levels: append([]log.Level{level}, otherLevels...),
	}
}

type containAnyMessageMatcher struct {
	levels []log.Level
}

func (m *containAnyMessageMatcher) Match(actual interface{}) (success bool, err error) {
	console, ok := actual.(io.Reader)
	if !ok {
		return false, errors.Errorf("ContainAnyMessageWithLevels matcher expects an io.Reader (actual: %T)", actual)
	}
	scanner := bufio.NewScanner(console)
	for scanner.Scan() {
		out := make(map[string]interface{})
		err := json.Unmarshal(scanner.Bytes(), &out)
		if err != nil {
			return false, errors.Wrapf(err, "failed to decode console line")
		}
		if level, ok := out["level"].(string); ok {
			for _, l := range m.levels {
				if l.String() == level {
					return true, nil
				}
			}
		}
	}
	// no match found
	return false, nil
}

func (m *containAnyMessageMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected console to contain a message at level '%v'", m.levels)
}

func (m *containAnyMessageMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected console not to contain a message at level '%v'", m.levels)
}
