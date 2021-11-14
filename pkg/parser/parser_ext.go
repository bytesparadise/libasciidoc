package parser

import (
	"errors"
	"fmt"
	"sort"
	"strings"

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
