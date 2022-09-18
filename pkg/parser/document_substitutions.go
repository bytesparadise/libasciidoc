package parser

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

const (
	// AttributeRefs the "attribute_refs" substitution
	AttributeRefs string = "attributes" // TODO: no need to export
	// Callouts the "callouts" substitution
	Callouts string = "callouts"
	// InlinePassthroughs the "inline_passthrough" substitution
	InlinePassthroughs string = "inline_passthrough" //nolint:gosec
	// Macros the "macros" substitution
	Macros string = "macros"
	// None the "none" substitution
	None string = "none"
	// PostReplacements the "post_replacements" substitution
	PostReplacements string = "post_replacements"
	// Quotes the "quotes" substitution
	Quotes string = "quotes"
	// Replacements the "replacements" substitution
	Replacements string = "replacements"
	// SpecialCharacters the "specialchars" substitution
	SpecialCharacters string = "specialchars"
)

func normalSubstitutions() *substitutions {
	return &substitutions{
		sequence: []string{
			InlinePassthroughs,
			AttributeRefs,
			SpecialCharacters,
			Quotes,
			Replacements,
			Macros,
			PostReplacements,
		},
	}
}

func headerSubstitutions() *substitutions {
	return &substitutions{
		sequence: []string{
			InlinePassthroughs,
			AttributeRefs,
			SpecialCharacters,
			Quotes,
			Macros,
			Replacements,
		},
	}
}

func attributeSubstitutions() *substitutions {
	return &substitutions{
		sequence: []string{
			InlinePassthroughs,
			AttributeRefs,
			SpecialCharacters,
			Quotes,
			// Macros,
			Replacements,
		},
	}
}

func noneSubstitutions() *substitutions {
	return &substitutions{}
}

func verbatimSubstitutions() *substitutions {
	return &substitutions{
		sequence: []string{
			Callouts,
			SpecialCharacters,
		},
	}
}

type substitutions struct {
	sequence []string
}

func newSubstitutions(b types.WithElements) (*substitutions, error) {
	// look-up the optional `subs` attribute in the element
	attrSub, found := b.GetAttributes().GetAsString(types.AttrSubstitutions)
	if !found {
		return defaultSubstitutions(b), nil
	}
	subs := strings.Split(attrSub, ",")
	var result *substitutions
	// when dealing with incremental substitutions, use default sub as a baseline and append or prepend the incremental subs
	if allIncremental(subs) {
		result = defaultSubstitutions(b)
	} else {
		result = &substitutions{
			sequence: make([]string, 0, len(subs)),
		}
	}
	for _, sub := range subs {
		// log.Debugf("checking subs '%s'", sub)
		switch {
		case strings.HasSuffix(sub, "+"): // prepend
			if err := result.prepend(strings.TrimSuffix(sub, "+")); err != nil {
				return nil, err
			}
		case strings.HasPrefix(sub, "+"): // append
			if err := result.append(strings.TrimPrefix(sub, "+")); err != nil {
				return nil, err
			}
		case strings.HasPrefix(sub, "-"): // remove from all substitutions
			if err := result.remove(strings.TrimPrefix(sub, "-")); err != nil {
				return nil, err
			}
		default:
			if err := result.append(sub); err != nil {
				return nil, err
			}
		}
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("substitutions to apply on block of type '%T': %s", b, spew.Sdump(result))
	}
	return result, nil
}

func defaultSubstitutions(b types.WithElements) *substitutions {
	// log.Debugf("looking-up default substitution for block of type '%T'", b)
	switch b := b.(type) {
	case *types.DelimitedBlock:
		switch b.Kind {
		case types.Example, types.Quote, types.Verse, types.Sidebar, types.MarkdownQuote, types.Open:
			return normalSubstitutions()
		case types.Comment, types.Passthrough:
			return noneSubstitutions()
		default: // includes `types.Listing`, `types.Fenced`, `types.Literal`
			return verbatimSubstitutions()
		}
	case *types.Paragraph:
		// if listing paragraph:
		switch b.GetAttributes().GetAsStringWithDefault(types.AttrStyle, "") {
		case types.Listing:
			return verbatimSubstitutions()
		case types.Passthrough:
			return noneSubstitutions()
		default:
			return normalSubstitutions()
		}
	default:
		return normalSubstitutions()
	}
}

// checks if all the given subs are incremental (ie, prefixed with `+|-` or suffixed with `-`)
func allIncremental(subs []string) bool {
	for _, sub := range subs {
		if !(strings.HasPrefix(sub, "+") ||
			strings.HasPrefix(sub, "-") ||
			strings.HasSuffix(sub, "+")) {
			return false
		}
	}
	return true
}

func (s *substitutions) toString() string {
	return strings.Join(s.sequence, ",")
}

// split the actual substitutions in 2 parts, the first one containing
// all substitutions, the second part all substitutions except `inline_passthrough` and `attributes`
// (or nil if there were no other substitutions)
func (s *substitutions) split() (*substitutions, *substitutions) {
	phase1 := &substitutions{
		sequence: s.sequence, // all by default (in case not split needed)
	}
	var phase2 *substitutions
	for i, sub := range s.sequence {
		if sub == AttributeRefs && i < len(s.sequence)-1 {
			phase2 = &substitutions{
				sequence: s.sequence[i+1:],
			}
		}
	}
	return phase1, phase2
}

func (s *substitutions) contains(expected string) bool {
	for _, sub := range s.sequence {
		if sub == expected {
			return true
		}
	}
	return false
}

func (s *substitutions) append(v string) error {
	other, err := substitutionsFor(v)
	if err != nil {
		return err
	}
	s.sequence = append(s.sequence, other.sequence...)
	return nil
}

func (s *substitutions) prepend(v string) error {
	other, err := substitutionsFor(v)
	if err != nil {
		return err
	}
	s.sequence = append(other.sequence, s.sequence...)
	return nil
}

func (s *substitutions) remove(v string) error {
	other, err := substitutionsFor(v)
	if err != nil {
		return err
	}

	for _, toRemove := range other.sequence {
		sequence := make([]string, 0, len(s.sequence))
		for j := range s.sequence {
			if s.sequence[j] != toRemove {
				sequence = append(sequence, s.sequence[j])
			}
		}
		s.sequence = sequence
	}
	return nil
}

func substitutionsFor(s string) (*substitutions, error) {
	switch s {
	case "normal":
		return normalSubstitutions(), nil
	case "none":
		return noneSubstitutions(), nil
	case "verbatim":
		return verbatimSubstitutions(), nil
	case "attributes", "macros", "quotes", "replacements", "post_replacements", "callouts", "specialchars":
		return &substitutions{
			sequence: []string{s},
		}, nil
	default:
		// TODO: return `none` instead of `err` and log an error with the fragment position (use logger with fields?)
		return nil, fmt.Errorf("unsupported substitution: '%v'", s)
	}
}
