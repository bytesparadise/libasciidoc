package parser

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// extra methods for the parser

func (p *parser) setup(g *grammar) (err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}
	p.read() // advance to first rune

	return p.errs.err()
}

// TODO: return immediately if end of doc was reached? (would simplify the grammar, avoiding to check for !EOF before parsing a new element)
func (p *parser) next() (val interface{}, err error) {
	if p.pt.offset == len(p.data) {
		log.Debugf("reached end of document")
		return nil, nil
	}
	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}
	startRule, ok := p.rules[p.entrypoint]
	if !ok {
		log.Errorf("invalid entrypoint: '%s'", p.entrypoint)
		p.addErr(errInvalidEntrypoint)
		return nil, p.errs.err()
	}

	val, ok = p.parseRule(startRule)
	if !ok {
		if len(*p.errs) == 0 {
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			maxFailExpectedMap := make(map[string]struct{}, len(p.maxFailExpected))
			for _, v := range p.maxFailExpected {
				maxFailExpectedMap[v] = struct{}{}
			}
			expected := make([]string, 0, len(maxFailExpectedMap))
			eof := false
			if _, ok := maxFailExpectedMap["!."]; ok {
				delete(maxFailExpectedMap, "!.")
				eof = true
			}
			for k := range maxFailExpectedMap {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}

		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

const suffixTrackingKey = "space_suffix_tracking"
const spaceSuffix = "space_suffix"
const alphanumSuffix = "alphanum_suffix"

func (c *current) trackSuffix(element interface{}) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("tracking space at the end of:\n%s", spew.Sdump(element))
	// }
	switch e := element.(type) {
	case string:
		doTrackSuffix(c, e)
	case *types.StringElement:
		doTrackSuffix(c, e.Content)
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("space suffix detected: %t", c.globalStore[spaceSuffixTrackingKey])
	// }
}

func doTrackSuffix(c *current, content string) {
	r := []rune(content)
	suffix := r[len(r)-1]
	switch {
	case suffix == ' ': // strict space, not `\n`, `\r`, etc.
		c.globalStore[suffixTrackingKey] = spaceSuffix
	case unicode.IsLetter(suffix) || unicode.IsNumber(suffix):
		c.globalStore[suffixTrackingKey] = alphanumSuffix
	default:
		delete(c.globalStore, suffixTrackingKey)
	}
}

func (c *current) isPreceededBySpace() bool {
	k, found := c.globalStore[suffixTrackingKey]
	return found && k == spaceSuffix
}

func (c *current) isPreceededByAlphanum() bool {
	k, found := c.globalStore[suffixTrackingKey]
	return found && k == alphanumSuffix
}

// verifies that the content does not end with a space
func validateSingleQuoteElements(elements []interface{}) (bool, error) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("checking that there is no space at the end of:\n%s", spew.Sdump(elements))
	// }
	if len(elements) == 0 {
		return true, nil
	}
	switch s := elements[len(elements)-1].(type) {
	case *types.StringElement:
		return !strings.HasSuffix(s.Content, " "), nil
	case string:
		return !strings.HasSuffix(s, " "), nil
	default:
		return true, nil
	}
}

const rawSectionEnabledKey = "raw_section_enabled"

// sectionEnabled parser option to enable detection of (raw) section during preparsing
func sectionEnabled() Option {
	return GlobalStore(rawSectionEnabledKey, true)
}

// state info to determine if parsing is happening within a delimited block (any kind),
// in which case some grammar rules need to be disabled
func (c *current) isSectionEnabled() bool {
	enabled, found := c.globalStore[rawSectionEnabledKey].(bool)
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("raw sections enabled: %t", found && enabled)
	// }
	return found && enabled
}

const withinDelimitedBlockKey = "within_delimited_block"

func withinDelimitedBlock(v bool) Option {
	return GlobalStore(withinDelimitedBlockKey, v)
}

// state info to determine if parsing is happening within a delimited block (any kind),
// in which case some grammar rules need to be disabled
func (c *current) isWithinDelimitedBlock() bool {
	w, found := c.globalStore[withinDelimitedBlockKey].(bool)
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("checking if within delimited block: %t/%t", found, w)
	// }
	return found && w
}

const blockDelimiterLengthKey = "block_delimiter_length"

// state info to indicate the length of the current block delimiter
func (c *current) setBlockDelimiterLength(length int) (bool, error) {
	c.globalStore[blockDelimiterLengthKey] = length
	return true, nil
}

// check if the length of the current block delimiters match
func (c *current) matchBlockDelimiterLength(length int) (bool, error) {
	return c.globalStore[blockDelimiterLengthKey] == length, nil
}

type blockDelimiterTracker struct {
	stack []blockDelimiter
}

type blockDelimiter struct {
	kind   types.BlockDelimiterKind
	length int
}

func newBlockDelimiterTracker() *blockDelimiterTracker {
	return &blockDelimiterTracker{
		stack: []blockDelimiter{},
	}
}

func (t *blockDelimiterTracker) push(kind types.BlockDelimiterKind, length int) {
	switch {
	case len(t.stack) > 0 && t.stack[len(t.stack)-1].kind == kind && t.stack[len(t.stack)-1].length == length:
		// trim
		t.stack = t.stack[:len(t.stack)-1]
	default:
		// append
		t.stack = append(t.stack, blockDelimiter{
			kind:   kind,
			length: length,
		})
	}
}

func (t *blockDelimiterTracker) withinDelimitedBlock() bool {
	return len(t.stack) > 0
}

const usermacrosKey = "user_macros"

func (c storeDict) hasUserMacro(name string) bool {
	if macros, exists := c[usermacrosKey].(map[string]configuration.MacroTemplate); exists {
		_, exists := macros[name]
		return exists
	}
	// log.Debugf("no user macro registered")
	return false
}
