package testsupport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ConfigureLogger configures the logger to write to a `Readable`.
// Also returns a func that can be used to reset the logger at the
// end of the test.

type TeeOption func(*ConsoleOutput)

var IncludeStdOut = func(t *ConsoleOutput) {
	t.out = os.Stdout
}

func ConfigureLogger(level log.Level, opts ...TeeOption) (*ConsoleOutput, func()) {
	t := NewConsoleOutput()
	for _, apply := range opts {
		apply(t)
	}
	if level == log.DebugLevel {
		t.out = os.Stdout // assume tee to stdout is needed too
	}
	if level == log.DebugLevel {
		t.out = os.Stdout // assume tee to stdout is needed too
	}
	log.SetOutput(t)
	fmtr := log.StandardLogger().Formatter
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: true,
	})
	oldLevel := log.GetLevel()
	log.SetLevel(level)

	return t, func() {
		log.SetOutput(os.Stdout)
		log.SetFormatter(fmtr)
		log.SetLevel(oldLevel)
	}
}

func NewConsoleOutput() *ConsoleOutput {
	return &ConsoleOutput{
		buf: &strings.Builder{},
		out: ioutil.Discard,
	}
}

type ConsoleOutput struct {
	buf *strings.Builder
	out io.Writer
}

func (t *ConsoleOutput) Write(p []byte) (n int, err error) {
	n, err = t.out.Write(p)
	if err != nil {
		return n, err
	}
	return t.buf.Write(p)
}

func (t ConsoleOutput) Content() io.Reader {
	return strings.NewReader(t.buf.String())
}

// ---------------------------
// ContainLogWithLevel
// ---------------------------

// ContainJSONLog a custom Matcher to verify that a message at a given level was logged
func ContainJSONLog(level log.Level, msg string) types.GomegaMatcher {
	return &containMessageMatcher{
		level:       level,
		msg:         msg,
		startOffset: float64(-1),
		endOffset:   float64(-1),
	}
}

// ContainJSONLogWithOffset a custom Matcher to verify that a message with offset position and at a given level was logged
func ContainJSONLogWithOffset(level log.Level, startOffset int, endOffset int, msg string) types.GomegaMatcher {
	return &containMessageMatcher{
		level:       level,
		msg:         msg,
		startOffset: float64(startOffset),
		endOffset:   float64(endOffset),
	}
}

type containMessageMatcher struct {
	level       log.Level
	msg         string
	startOffset float64
	endOffset   float64
}

type Console interface {
	Content() io.Reader
}

func (m *containMessageMatcher) Match(actual interface{}) (success bool, err error) {
	console, ok := actual.(Console)
	if !ok {
		return false, errors.Errorf("ContainJSONLog matcher expects an io.Reader (actual: %T)", actual)
	}
	scanner := bufio.NewScanner(console.Content())
scan:
	for scanner.Scan() {
		out := make(map[string]interface{})
		err := json.Unmarshal(scanner.Bytes(), &out)
		if err != nil {
			fmt.Printf("invalid content %s: %s\n", scanner.Text(), err.Error())
			return false, errors.Wrapf(err, "failed to decode console line")
		}
		if !strings.HasPrefix(out["msg"].(string), m.msg) ||
			out["level"] != m.level.String() ||
			(m.startOffset != -1 && out["start_offset"] != m.startOffset) ||
			(m.endOffset != -1 && out["end_offset"] != m.endOffset) {
			continue scan
		}
		// match found
		return true, nil
	}
	// no match found
	return false, nil
}

func (m *containMessageMatcher) FailureMessage(_ interface{}) (message string) {
	if m.startOffset != -1 || m.endOffset != -1 {
		return fmt.Sprintf(`expected console to contain log {"level": "%s", "start_offset":%d, "end_offset":%d, "msg":"%s"}`, m.level.String(), int(m.startOffset), int(m.endOffset), m.msg)
	}
	return fmt.Sprintf(`expected console to contain log {"level": "%s", "msg":"%s"}`, m.level.String(), m.msg)
}

func (m *containMessageMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	if m.startOffset != -1 || m.endOffset != -1 {
		return fmt.Sprintf(`expected console not to contain log {"level": "%s", "start_offset":%d, "end_offset":%d, "msg":"%s"}`, m.level.String(), int(m.startOffset), int(m.endOffset), m.msg)
	}
	return fmt.Sprintf(`expected console not to contain log {"level": "%s", "msg":"%s"}`, m.level.String(), m.msg)
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
	console, ok := actual.(Console)
	if !ok {
		return false, errors.Errorf("ContainAnyMessageWithLevels matcher expects an io.Reader (actual: %T)", actual)
	}
	scanner := bufio.NewScanner(console.Content())
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
