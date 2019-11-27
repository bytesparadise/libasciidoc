package types

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ReplaceNonAlphanumerics replace all non alpha numeric characters with the given `replacement`
func ReplaceNonAlphanumerics(source InlineElements, replacement string) (string, error) {
	v := newReplaceNonAlphanumericsVisitor(replacement)
	err := source.AcceptVisitor(v)
	if err != nil {
		return "", err
	}
	return v.normalizedContent(), nil
}

//ReplaceNonAlphanumericsVisitor a visitor that builds a string representation of the visited elements,
// in which all non-alphanumeric characters have been replaced with a "_"
type ReplaceNonAlphanumericsVisitor struct {
	buf         bytes.Buffer
	replacement string
}

var _ Visitor = &ReplaceNonAlphanumericsVisitor{}

// newReplaceNonAlphanumericsVisitor returns a new ReplaceNonAlphanumericsVisitor
func newReplaceNonAlphanumericsVisitor(replacement string) *ReplaceNonAlphanumericsVisitor {
	buf := bytes.NewBuffer(nil)
	return &ReplaceNonAlphanumericsVisitor{
		buf:         *buf,
		replacement: replacement,
	}
}

// Visit method called when an element is visited
func (v *ReplaceNonAlphanumericsVisitor) Visit(element Visitable) error {
	switch element := element.(type) {
	case StringElement:
		return v.process(element.Content)
	case InlineLink:
		return v.process(element.Location.String())
	default:
		// other types are ignored
		return nil
	}
}

func (v *ReplaceNonAlphanumericsVisitor) process(content string) error {
	if v.buf.Len() > 0 {
		v.buf.WriteString("_")
	}
	normalized, err := v.normalize(content)
	if err != nil {
		return errors.Wrapf(err, "error while normalizing String Element")
	}
	v.buf.WriteString(normalized)
	return nil
}

// normalize returns the normalized content
func (v *ReplaceNonAlphanumericsVisitor) normalize(source string) (string, error) {
	log.Debugf("normalizing '%s'", source)
	buf := bytes.NewBuffer(nil)
	lastCharIsSpace := false
	for _, r := range strings.TrimLeft(source, " ") { // ignore header spaces
		if unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Number, r) {
			_, err := buf.WriteString(strings.ToLower(string(r)))
			if err != nil {
				return "", errors.Wrapf(err, "unable to normalize value")
			}
			lastCharIsSpace = false
		} else if !lastCharIsSpace && (unicode.Is(unicode.Space, r) || unicode.Is(unicode.Punct, r)) {
			_, err := buf.WriteString(v.replacement)
			if err != nil {
				return "", errors.Wrapf(err, "unable to normalize value")
			}
			lastCharIsSpace = true
		}
	}
	return strings.TrimSuffix(buf.String(), v.replacement), nil
}

// NormalizedContent returns the normalized content
func (v *ReplaceNonAlphanumericsVisitor) normalizedContent() string {
	return v.buf.String()
}
