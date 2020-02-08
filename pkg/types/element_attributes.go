package types

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ------------------------------------------
// Elements attributes
// ------------------------------------------

const (
	// AttrID the key to retrieve the ID in the element attributes
	AttrID string = "id"
	// AttrCustomID the key to retrieve the flag that indicates if the element ID is custom or generated
	AttrCustomID string = "customID"
	// AttrIDPrefix the key to retrieve the ID Prefix in the element attributes
	AttrIDPrefix string = "idprefix"
	// DefaultIDPrefix the default ID Prefix in the element attributes
	DefaultIDPrefix string = "_"
	// AttrTitle the key to retrieve the title in the element attributes
	AttrTitle string = "title"
	// AttrAuthors the key to the authors declared after the section level 0 (at the beginning of the doc)
	AttrAuthors string = "authors"
	// AttrRevision the key to the revision declared after the section level 0 (at the beginning of the doc)
	AttrRevision string = "revision"
	// AttrTableOfContents the `toc` attribute at document level
	AttrTableOfContents string = "toc"
	// AttrTableOfContentsLevels the document attribute which specifies the number of levels to display in the ToC
	AttrTableOfContentsLevels string = "toclevels"
	// AttrRole the key to retrieve the role in the element attributes
	AttrRole string = "role"
	// AttrInlineLink the key to retrieve the link in the element attributes
	AttrInlineLink string = "link"
	// AttrAdmonitionKind the key to retrieve the kind of admonition in the element attributes, if a "masquerade" is used
	AttrAdmonitionKind string = "admonitionKind"
	// AttrQuoteAuthor attribute for the author of a verse
	AttrQuoteAuthor string = "quoteAuthor"
	// AttrQuoteTitle attribute for the title of a verse
	AttrQuoteTitle string = "quoteTitle"
	// AttrSource the `source` attribute for a source block or a source paragraph (this is a placeholder, ie, it does not expect any value for this attribute)
	AttrSource string = "source"
	// AttrLanguage the `language` attribute for a source block or a source paragraph
	AttrLanguage string = "language"
	// AttrLineNums the `linenums` attribute for a source block or a source paragraph
	AttrLineNums string = "linenums"
	// AttrCheckStyle the attribute to mark the first element of an unordered list itemd as a checked or not
	AttrCheckStyle string = "checkstyle"
	// AttrStart the `start` attribute in an ordered list
	AttrStart string = "start"
	// AttrNumberingStyle the numbering style of items in a list
	AttrNumberingStyle string = "numberingStyle"
	// AttrQandA the `qanda` attribute for Q&A labeled lists
	AttrQandA string = "qanda"
	// AttrLevelOffset the `leveloffset` attribute used in file inclusions
	AttrLevelOffset string = "leveloffset"
	// AttrLineRanges the `lines` attribute used in file inclusions
	AttrLineRanges string = "lines"
	// AttrTagRanges the `tag`/`tags` attribute used in file inclusions
	AttrTagRanges string = "tags"
	// AttrLastUpdated the "last updated" data in the document, i.e., the output/generation time
	AttrLastUpdated string = "LastUpdated"
	// AttrImageAlt the image `alt` attribute
	AttrImageAlt string = "alt"
	// AttrImageWidth the image `width` attribute
	AttrImageWidth string = "width"
	// AttrImageHeight the image `height` attribute
	AttrImageHeight string = "height"
	// AttrImageTitle the image `title` attribute
	AttrImageTitle string = "title"
)

// ElementWithAttributes an element on which attributes can be added/set
type ElementWithAttributes interface {
	AddAttributes(attributes ElementAttributes)
}

// NewElementID initializes a new attribute map with a single entry for the ID using the given value
func NewElementID(id string) (ElementAttributes, error) {
	// log.Debugf("initializing a new ElementID with ID=%s", id)
	return ElementAttributes{
		AttrID:       id,
		AttrCustomID: true,
	}, nil
}

// NewInlineElementID initializes a new attribute map with a single entry for the ID using the given value
func NewInlineElementID(id string) (ElementAttributes, error) {
	log.Debugf("initializing a new inline ElementID with ID=%s", id)
	return ElementAttributes{AttrID: id}, nil
}

// NewElementTitle initializes a new attribute map with a single entry for the title using the given value
func NewElementTitle(title string) (ElementAttributes, error) {
	log.Debugf("initializing a new ElementTitle with content=%s", title)
	return ElementAttributes{
		AttrTitle: title,
	}, nil
}

// NewElementRole initializes a new attribute map with a single entry for the title using the given value
func NewElementRole(role string) (ElementAttributes, error) {
	// log.Debugf("initializing a new ElementRole with content=%s", role)
	return ElementAttributes{
		AttrRole: role,
	}, nil
}

// NewAdmonitionAttribute initializes a new attribute map with a single entry for the admonition kind using the given value
func NewAdmonitionAttribute(k AdmonitionKind) (ElementAttributes, error) {
	return ElementAttributes{AttrAdmonitionKind: k}, nil
}

// NewAttributeGroup initializes a group of attributes from the given generic attributes.
func NewAttributeGroup(attributes []interface{}) (ElementAttributes, error) {
	// log.Debugf("initializing a new AttributeGroup with %v", attributes)
	result := make(ElementAttributes)
	for _, a := range attributes {
		// log.Debugf("processing attribute element of type %T", a)
		if a, ok := a.(ElementAttributes); ok {
			for k, v := range a {
				// log.Debugf("adding attribute %v='%v'", k, v)
				result[k] = v
			}
		} else {
			return result, errors.Errorf("unable to process element of type '%[1]T': '%[1]s'", a)
		}
	}
	// log.Debugf("initialized a new AttributeGroup: %v", result)
	return result, nil
}

// NewGenericAttribute initializes a new ElementAttribute from the given key and optional value
func NewGenericAttribute(key string, value interface{}) (ElementAttributes, error) {
	result := make(map[string]interface{})
	k := Apply(key,
		// remove surrounding quotes
		func(s string) string {
			return strings.Trim(s, "\"")
		},
		strings.TrimSpace)
	result[k] = nil
	if value, ok := value.(string); ok {
		v := Apply(value,
			// remove surrounding quotes
			func(s string) string {
				return strings.Trim(s, "\"")
			},
			strings.TrimSpace)
		if len(v) > 0 {
			result[k] = v
		}
	}
	// log.Debugf("initialized a new ElementAttributes: %v", result)
	return result, nil
}

// NewQuoteAttributes initializes a new map of attributes for a verse paragraph
func NewQuoteAttributes(kind string, author, title interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 3)
	switch kind {
	case "verse":
		result[AttrKind] = Verse
	default:
		result[AttrKind] = Quote
	}
	if author, ok := author.(string); ok {
		author = strings.TrimSpace(author)
		if len(author) > 0 {
			result[AttrQuoteAuthor] = author
		}
	}
	if title, ok := title.(string); ok {
		title = strings.TrimSpace(title)
		if len(title) > 0 {
			result[AttrQuoteTitle] = title
		}
	}
	return result, nil
}

// NewLiteralAttribute initializes a new attribute map with a single entry for the literal kind of block
func NewLiteralAttribute() (ElementAttributes, error) {
	return ElementAttributes{AttrKind: Literal}, nil
}

// NewSourceAttributes initializes a new attribute map with two entries, one for the kind of element ("source") and another optional one for the language of the source code
func NewSourceAttributes(language interface{}, others ...interface{}) (ElementAttributes, error) {
	result := ElementAttributes{
		AttrKind: Source,
	}
	if language, ok := language.(string); ok {
		result[AttrLanguage] = strings.TrimSpace(language)
	}
	for _, other := range others {
		if other, ok := other.(ElementAttributes); ok {
			result.AddAll(other)
		}
	}
	return result, nil
}

// ElementAttributes is a map[string]interface{} with some utility methods
type ElementAttributes map[string]interface{}

// Has returns the true if an entry with the given key exists
func (a ElementAttributes) Has(key string) bool {
	_, ok := a[key]
	return ok
}

// NilSafeSet sets the key/value pair unless the value is nil or empty
func (a ElementAttributes) NilSafeSet(key string, value interface{}) {
	if value != nil && value != "" {
		a[key] = value
	}
}

// GetAsString returns the value of the key as a string, or empty string if the key did not exist
func (a ElementAttributes) GetAsString(key string) string {
	if v, ok := a[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

// GetAsBool returns the value of the key as a bool, or `false` if the key did not exist
// or if its value was not a bool
func (a ElementAttributes) GetAsBool(key string) bool {
	if v, ok := a[key]; ok {
		if v, ok := v.(bool); ok {
			return v
		}
	}
	return false
}

// AddAll adds all the given attributes to the current ones
func (a ElementAttributes) AddAll(attributes ElementAttributes) {
	if attributes == nil {
		return
	}
	for k, v := range attributes {
		a[k] = v
	}
}

// NewElementAttributes retrieves the ElementID, ElementTitle and ElementInlineLink from the given slice of attributes
func NewElementAttributes(attributes []interface{}) ElementAttributes {
	attrs := ElementAttributes{}
	for _, attr := range attributes {
		// log.Debugf("processing attribute %[1]v (%[1]T)", attr)
		switch attr := attr.(type) {
		case []interface{}:
			// nested case, because of the grammar syntax,
			// eg: `attributes:(ElementAttribute* LiteralAttribute ElementAttribute*)`
			// which is used to ensure that a `LiteralAttribute` element is set amongst the attributes
			r := NewElementAttributes(attr)
			for k, v := range r {
				attrs[k] = v
			}
		case ElementAttributes:
			// TODO: warn if attribute already exists and is overridden
			for k, v := range attr {
				attrs[k] = v
			}
		case map[string]interface{}:
			// TODO: warn if attribute already exists and is overridden
			for k, v := range attr {
				attrs[k] = v
			}
		case nil:
			// ignore
		default:
			log.Warnf("unexpected attributes of type: %T", attr)
		}
	}
	return attrs
}

// NewInlineAttributes returns a map of attributes
func NewInlineAttributes(attrs []interface{}) (ElementAttributes, error) {
	result := ElementAttributes{}
	for _, attr := range attrs {
		if attr, ok := attr.(ElementAttributes); ok {
			for k, v := range attr {
				result[k] = v
			}
		}
	}
	return result, nil
}

func resolveAlt(path Location) string {
	_, filename := filepath.Split(path.String())
	ext := filepath.Ext(filename)
	if ext != "" {
		return strings.TrimSuffix(filename, ext)
	}
	return filename
}
