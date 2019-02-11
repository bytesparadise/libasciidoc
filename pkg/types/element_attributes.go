package types

import (
	"fmt"
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
	// AttrTitle the key to retrieve the title in the element attributes
	AttrTitle string = "title"
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
	// AttrSource the "source" attribute for a source block or a source paragraph (this is a placeholder, ie, it does not expect any value for this attribute)
	AttrSource string = "source"
	// AttrLanguage the associated "language" attribute for a source block or a source paragraph
	AttrLanguage string = "language"
	// AttrCheckStyle the attribute to mark the first element of an unordered list itemd as a checked or not
	AttrCheckStyle string = "checkstyle"
	// AttrStart the "start" attribute in an ordered list
	AttrStart string = "start"
	// AttrQandA the "qanda" attribute for Q&A labeled lists
	AttrQandA string = "qanda"
)

// ElementWithAttributes an element on which attributes can be added/set
type ElementWithAttributes interface {
	AddAttributes(attributes ElementAttributes)
}

// NewElementID initializes a new attribute map with a single entry for the ID using the given value
func NewElementID(id string) (ElementAttributes, error) {
	log.Debugf("initializing a new ElementID with ID=%s", id)
	return ElementAttributes{AttrID: id}, nil
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
	log.Debugf("initializing a new ElementRole with content=%s", role)
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
		log.Debugf("processing attribute element of type %T", a)
		if a, ok := a.(ElementAttributes); ok {
			for k, v := range a {
				log.Debugf("adding attribute %v='%v'", k, v)
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
	k := apply(key,
		// remove surrounding quotes
		func(s string) string {
			return strings.Trim(s, "\"")
		},
		strings.TrimSpace)
	if value, ok := value.(string); ok {
		v := apply(value,
			// remove surrounding quotes
			func(s string) string {
				return strings.Trim(s, "\"")
			},
			strings.TrimSpace)
		result[k] = v
	} else {
		result[k] = nil
	}
	// log.Debugf("initialized a new ElementAttributes: %v", result)
	return result, nil
}

// NewQuoteAttributes initializes a new map of attributes for a verse paragraph
func NewQuoteAttributes(kind, author, title string) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 3)
	switch kind {
	case "verse":
		result[AttrKind] = Verse
	default:
		result[AttrKind] = Quote
	}
	result[AttrQuoteAuthor] = strings.TrimSpace(author)
	result[AttrQuoteTitle] = strings.TrimSpace(title)
	log.Debugf("initialized new %s attributes: %v", kind, result)
	return result, nil
}

// NewLiteralAttribute initializes a new attribute map with a single entry for the literal kind of block
func NewLiteralAttribute() (ElementAttributes, error) {
	log.Debug("initializing a new Literal attribute")
	return ElementAttributes{AttrKind: Literal}, nil
}

// NewSourceAttributes initializes a new attribute map with two entries, one for the kind of element ("source") and another optional one for the language of the source code
func NewSourceAttributes(language string) (ElementAttributes, error) {
	log.Debugf("initializing a new source attribute (language='%s')", language)
	return ElementAttributes{
		AttrKind:     Source,
		AttrLanguage: strings.TrimSpace(language),
	}, nil
}

// WithAttributes set the attributes on the given elements if its type is supported, otherwise returns an error
func WithAttributes(element interface{}, attributes []interface{}) (interface{}, error) {
	attrbs := NewElementAttributes(attributes)
	if element, ok := element.(ElementWithAttributes); ok {
		if len(attributes) > 0 {
			log.Debugf("setting %d attribute(s) on element of type %T", len(attributes), element)
		}
		element.AddAttributes(attrbs)
		return element, nil
	}
	// special case for DelimitedBlock where we need a pointer receiver to modify the `Kind` field of the struct.
	if element, ok := element.(DelimitedBlock); ok {
		block := &element
		block.AddAttributes(attrbs)
		return element, nil
	}
	// special case for any ListItem where we need a pointer receiver to modify the `Kind` field of the struct.
	if element, ok := element.(OrderedListItem); ok {
		item := &element
		item.AddAttributes(attrbs)
		return element, nil
	}
	if element, ok := element.(UnorderedListItem); ok {
		item := &element
		item.AddAttributes(attrbs)
		return element, nil
	}
	if element, ok := element.(LabeledListItem); ok {
		item := &element
		item.AddAttributes(attrbs)
		return element, nil
	}

	log.Debugf("cannot set attribute(s) %[2]v on element of type %[1]T : %[1]v", element, attributes)
	return nil, errors.Errorf("cannot set attributes on element of type '%T'", element)
}

// ElementAttributes is a map[string]interface{} with some utility methods
type ElementAttributes map[string]interface{}

// Has returns the true if an entry with the given key exists
func (a ElementAttributes) Has(key string) bool {
	_, ok := a[key]
	return ok
}

// GetAsString returns the value of the key as a string, or empty string if the key did not exist
func (a ElementAttributes) GetAsString(key string) string {
	if v, ok := a[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

// AddAll adds all the given attributes to the current ones
func (a ElementAttributes) AddAll(attributes ElementAttributes) {
	for k, v := range attributes {
		a[k] = v
	}
}

// NewElementAttributes retrieves the ElementID, ElementTitle and ElementInlineLink from the given slice of attributes
func NewElementAttributes(attributes []interface{}, extras ...ElementAttributes) ElementAttributes {
	attrbs := ElementAttributes{}
	for _, attrb := range attributes {
		log.Debugf("processing attribute %[1]v (%[1]T)", attrb)
		switch attrb := attrb.(type) {
		case []interface{}:
			// nested case, because of the grammar syntax,
			// eg: `attributes:(ElementAttribute* LiteralAttribute ElementAttribute*)`
			// which is used to ensure that a `LiteralAttribute` element is set amongst the attributes
			r := NewElementAttributes(attrb)
			for k, v := range r {
				attrbs[k] = v
			}
		case ElementAttributes:
			// TODO: warn if attribute already exists and is overridden
			for k, v := range attrb {
				attrbs[k] = v
			}
		case map[string]interface{}:
			// TODO: warn if attribute already exists and is overridden
			for k, v := range attrb {
				attrbs[k] = v
			}
		case nil:
			// ignore
		default:
			log.Warnf("unexpected attributes of type: %T", attrb)
		}
	}
	for _, extra := range extras {
		for k, v := range extra {
			// no warning on override here
			attrbs[k] = v
		}
	}
	return attrbs
}
