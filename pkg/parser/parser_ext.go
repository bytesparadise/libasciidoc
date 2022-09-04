package parser

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

func parseContent(content []byte, opts ...Option) ([]interface{}, error) {
	p := newParser("", content, opts...)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("parsing content from '%s' entrypoint", p.entrypoint)
	}
	result, err := p.parse(g)
	if err != nil {
		return nil, err
	}
	if result, ok := result.([]interface{}); ok {
		return result, nil
	}
	return nil, fmt.Errorf("unexpected type of result after parsing content: '%T'", result)
}

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

const suffixTrackingKey = "element_suffix_tracking"

func (c *current) trackElement(element interface{}) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("---tracking element of type '%T' at pos=%s text='%s'", element, c.pos.String(), string(c.text))
	}
	c.globalStore[suffixTrackingKey] = rune(c.text[len(c.text)-1])
}

func (c *current) isPrecededBySpace() bool {
	if r, ok := c.globalStore[suffixTrackingKey].(rune); ok {
		log.Debugf("---checking if preceded by space (tracked='%s')", string(r))
		result := unicode.IsSpace(r) || unicode.IsControl(r)
		return result
	}
	log.Debugf("---is not preceded by space (no previous character)")
	return false
}

// verifies that previous last character of previous match is neither a
// letter, number, underscore, colon, semicolon, or closing curly bracket
func (c *current) isSingleQuotedTextAllowed() bool {
	if r, ok := c.globalStore[suffixTrackingKey].(rune); ok {
		log.Debugf("---checking if single quoted text is allowed (tracked='%s')", string(r))
		// r := rune(d[c.pos.offset])
		result := !unicode.IsLetter(r) &&
			!unicode.IsNumber(r) &&
			// r != '_' &&
			r != ',' &&
			r != ';' &&
			r != '}'
		return result
	}
	log.Debugf("---single quoted text is allowed (no previous character)")
	return true
}

func (c *current) isPrecededByAlphanum() bool {
	if r, ok := c.globalStore[suffixTrackingKey].(rune); ok {
		log.Debugf("---checking if preceded by alphanum before (tracked='%s')", string(r))
		result := unicode.IsLetter(r) || unicode.IsNumber(r)
		return result
	}
	log.Debugf("---is not preceded by alphanum (no previous character)")
	return false
}

// verifies that the content does not end with a space
func validateSingleQuoteElements(elements []interface{}) (bool, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("checking that there is no space at the end of:\n%s", spew.Sdump(elements))
	}
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

const usermacrosKey = "user_macros"

func (c *current) hasUserMacro(name string) bool {
	if macros, ok := c.globalStore[usermacrosKey].(map[string]configuration.MacroTemplate); ok {
		_, found := macros[name]
		// log.Debugf("user macro '%s' registered: %t", name, found)
		return found
	}
	// log.Debugf("no user macro registered")
	return false
}

// enabledSubstitutionsKey the key in which enabled substitutions are stored in the parser's GlobalStore
const enabledSubstitutionsKey string = "enabled_substitutions"

func (c *current) withSubstitutions(subs *substitutions) {
	c.state[enabledSubstitutionsKey] = subs
}

func (c *current) lookupCurrentSubstitutions() (*substitutions, bool) {
	if s, found := c.state[enabledSubstitutionsKey].(*substitutions); found {
		return s, true
	}
	s, found := c.globalStore[enabledSubstitutionsKey].(*substitutions)
	return s, found
}

func (c *current) isSubstitutionEnabled(k string) bool {
	subs, found := c.lookupCurrentSubstitutions()
	if !found {
		log.Debugf("substitutions not set in globalStore: assuming '%s' not enabled", k)
		return false // TODO: should return `true`, at least for `attributes`?
	}
	for _, s := range subs.sequence {
		if s == k {
			// log.Debugf("'%s' is enabled", k)
			return true
		}
	}
	// log.Debugf("'%s' is not enabled", k)
	return false
}
