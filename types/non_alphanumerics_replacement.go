package types

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//NormalizationFunc a function that is used to normalize a string.
type NormalizationFunc func(string) ([]byte, error)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// NewReplaceNonAlphanumericsFunc replaces all non alphanumerical characters and remove (accents)
// in the given 'source' with the given 'replacement'.
func NewReplaceNonAlphanumericsFunc(replacement string) NormalizationFunc {
	return func(source string) ([]byte, error) {
		buf := bytes.NewBuffer(nil)
		lastCharIsSpace := false
		for _, r := range strings.TrimLeft(source, " ") { // ignore header spaces
			if unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Number, r) {
				_, err := buf.WriteString(strings.ToLower(string(r)))
				if err != nil {
					return nil, errors.Wrapf(err, "unable to normalize value")
				}
				lastCharIsSpace = false
			} else if !lastCharIsSpace && (unicode.Is(unicode.Space, r) || unicode.Is(unicode.Punct, r)) {
				_, err := buf.WriteString(replacement)
				if err != nil {
					return nil, errors.Wrapf(err, "unable to normalize value")
				}
				lastCharIsSpace = true
			}
		}
		result := strings.TrimSuffix(buf.String(), replacement)
		return []byte(result), nil
	}
}

// ReplaceNonAlphanumerics replace all non alpha numeric characters with the given `replacement`
func ReplaceNonAlphanumerics(source InlineContent, replacement string) (string, error) {
	v := NewReplaceNonAlphanumericsVisitor()
	err := source.Accept(&v)
	if err != nil {
		return "", err
	}
	return v.NormalizedContent(), nil
}

//ReplaceNonAlphanumericsVisitor a visitor that builds a string representation of the visited elements,
// in which all non-alphanumeric characters have been replaced with a "_"
type ReplaceNonAlphanumericsVisitor struct {
	buf       bytes.Buffer
	normalize NormalizationFunc
}

// NewReplaceNonAlphanumericsVisitor returns a new ReplaceNonAlphanumericsVisitor
func NewReplaceNonAlphanumericsVisitor() ReplaceNonAlphanumericsVisitor {
	buf := bytes.NewBuffer(nil)
	return ReplaceNonAlphanumericsVisitor{
		buf:       *buf,
		normalize: NewReplaceNonAlphanumericsFunc("_"),
	}
}

// Visit method called when an element is visited
func (v *ReplaceNonAlphanumericsVisitor) Visit(element Visitable) error {
	log.Debugf("visiting element of type '%T'", element)
	switch element := element.(type) {
	case InlineContent:
		// log.Debugf("Prefixing with '_' while processing '%T'", element)
		v.buf.WriteString("_")
	case StringElement:
		normalized, err := v.normalize(element.Content)
		if err != nil {
			return errors.Wrapf(err, "error while normalizing String Element")
		}
		v.buf.Write(normalized)
	default:
		// ignore
	}
	return nil
}

// BeforeVisit method called before visiting an element. Allows for performing "pre-actions"
func (v *ReplaceNonAlphanumericsVisitor) BeforeVisit(element Visitable) error {
	log.Debugf("Before visiting element of type '%T'...", element)
	switch element := element.(type) {
	case QuotedText:
		// log.Debugf("Before visiting quoted element...")
		switch element.Kind {
		case Bold:
			v.buf.WriteString("_strong_")
		case Italic:
			v.buf.WriteString("_italic_")
		case Monospace:
			v.buf.WriteString("_monospace_")
		default:
			return errors.Errorf("unsupported kind of quoted text: %d", element.Kind)
		}
	default:
		// ignore
	}
	return nil
}

// AfterVisit method called before visiting an element. Allows for performing "post-actions"
func (v *ReplaceNonAlphanumericsVisitor) AfterVisit(element Visitable) error {
	log.Debugf("After visiting element of type '%T'...", element)
	switch element := element.(type) {
	case QuotedText:
		switch element.Kind {
		case Bold:
			v.buf.WriteString("_strong")
		case Italic:
			v.buf.WriteString("_italic")
		case Monospace:
			v.buf.WriteString("_monospace")
		default:
			return errors.Errorf("unsupported kind of quoted text: %d", element.Kind)
		}
	default:
		// ignore
	}
	return nil
}

// NormalizedContent returns the normalized content
func (v *ReplaceNonAlphanumericsVisitor) NormalizedContent() string {
	return v.buf.String()
}
