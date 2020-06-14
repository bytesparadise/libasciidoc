package types

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ------------------------------------------
// Attributes
// ------------------------------------------

const (
	// AttrDocType the "doctype" attribute
	AttrDocType string = "doctype"
	// AttrSyntaxHighlighter the attribute to define the syntax highlighter on code source blocks
	AttrSyntaxHighlighter string = "source-highlighter"
	// AttrIDPrefix the key to retrieve the ID Prefix
	AttrIDPrefix string = "idprefix"
	// DefaultIDPrefix the default ID Prefix
	DefaultIDPrefix string = "_"
	// AttrTableOfContents the `toc` attribute at document level
	AttrTableOfContents string = "toc"
	// AttrTableOfContentsLevels the document attribute which specifies the number of levels to display in the ToC
	AttrTableOfContentsLevels string = "toclevels"
	// AttrNoHeader attribute to disable the rendering of document footer
	AttrNoHeader string = "noheader"
	// AttrNoFooter attribute to disable the rendering of document footer
	AttrNoFooter string = "nofooter"
	// AttrID the key to retrieve the ID
	AttrID string = "id"
	// AttrCustomID the key to retrieve the flag that indicates if the element ID is custom or generated
	AttrCustomID string = "customID"
	// AttrTitle the key to retrieve the title
	AttrTitle string = "title"
	// AttrAuthors the key to the authors declared after the section level 0 (at the beginning of the doc)
	AttrAuthors string = "authors"
	// AttrRevision the key to the revision declared after the section level 0 (at the beginning of the doc)
	AttrRevision string = "revision"
	// AttrRole the key to retrieve the role
	AttrRole string = "role"
	// AttrInlineLink the key to retrieve the link
	AttrInlineLink string = "link"
	// AttrAdmonitionKind the key to retrieve the kind of admonition , if a "masquerade" is used
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

// NewElementID initializes a new attribute map with a single entry for the ID using the given value
func NewElementID(id string) (Attributes, error) {
	// log.Debugf("initializing a new ElementID with ID=%s", id)
	return Attributes{
		AttrID:       id,
		AttrCustomID: true,
	}, nil
}

// NewInlineElementID initializes a new attribute map with a single entry for the ID using the given value
func NewInlineElementID(id string) (Attributes, error) {
	log.Debugf("initializing a new inline ElementID with ID=%s", id)
	return Attributes{AttrID: id}, nil
}

// NewElementTitle initializes a new attribute map with a single entry for the title using the given value
func NewElementTitle(title string) (Attributes, error) {
	// log.Debugf("initializing a new ElementTitle with content=%s", title)
	return Attributes{
		AttrTitle: title,
	}, nil
}

// NewElementRole initializes a new attribute map with a single entry for the title using the given value
func NewElementRole(role string) (Attributes, error) {
	// log.Debugf("initializing a new ElementRole with content=%s", role)
	return Attributes{
		AttrRole: role,
	}, nil
}

// NewAdmonitionAttribute initializes a new attribute map with a single entry for the admonition kind using the given value
func NewAdmonitionAttribute(k AdmonitionKind) (Attributes, error) {
	return Attributes{AttrAdmonitionKind: k}, nil
}

// NewAttributeGroup initializes a group of attributes from the given generic attributes.
func NewAttributeGroup(attributes []interface{}) (Attributes, error) {
	// log.Debugf("initializing a new AttributeGroup with %v", attributes)
	result := make(Attributes)
	for _, a := range attributes {
		// log.Debugf("processing attribute element of type %T", a)
		if a, ok := a.(Attributes); ok {
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
func NewGenericAttribute(key string, value interface{}) (Attributes, error) {
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
	// log.Debugf("initialized a new Attributes: %v", result)
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
func NewLiteralAttribute() (Attributes, error) {
	return Attributes{AttrKind: Literal}, nil
}

// NewPassthroughBlockAttribute initializes a new attribute map with a single entry for the passthrough kind of block
func NewPassthroughBlockAttribute() (Attributes, error) {
	return Attributes{AttrKind: Passthrough}, nil
}

// NewSourceAttributes initializes a new attribute map with two entries, one for the kind of element ("source") and another optional one for the language of the source code
func NewSourceAttributes(language interface{}, others ...interface{}) (Attributes, error) {
	result := Attributes{
		AttrKind: Source,
	}
	if language, ok := language.(string); ok {
		result[AttrLanguage] = strings.TrimSpace(language)
	}
	for _, other := range others {
		result.Add(other)
	}
	return result, nil
}

// Attributes is a map[string]interface{} with some utility methods
type Attributes map[string]interface{}

// Set sets the key/value entry in the Attributes.
func (a Attributes) Set(key string, value interface{}) Attributes {
	if a == nil {
		a = Attributes{}
	}
	a[key] = value
	return a
}

// Add adds the given attributes to the current ones
func (a Attributes) Add(attrs interface{}) Attributes {
	if a == nil {
		a = Attributes{}
	}
	if attrs, ok := attrs.(Attributes); ok {
		for k, v := range attrs {
			a[k] = v
			if k == AttrID {
				a[AttrCustomID] = true
			}
		}
	}
	return a
}

// Has returns the true if an entry with the given key exists
func (a Attributes) Has(key string) bool {
	_, ok := a[key]
	return ok
}

// AppendString sets the value as a singular string value if it did not exist yet,
// or move the existing value in a slice of strings and append the new one
func (a Attributes) AppendString(key string, value string) {
	v, found := a[key]
	if !found {
		a[key] = value
		return
	}
	switch v := v.(type) {
	case string:
		a[key] = []string{v, value} // move existing value in a slice, along with the new one
	case []string:
		a[key] = append(v, value) // just append the new value into the slice
	}
}

// NilSafeSet sets the key/value pair unless the value is nil or empty
func (a Attributes) NilSafeSet(key string, value interface{}) {
	if value != nil && value != "" {
		a[key] = value
	}
}

// GetAsString gets the string value for the given key (+ `true`),
// or empty string (+ `false`) if none was found
// TODO: raise a warning if there was no entry found
func (a Attributes) GetAsString(key string) (string, bool) {
	// check in predefined attributes
	if value, found := Predefined[key]; found {
		return value, true
	}
	if value, found := a[key]; found {
		if value, ok := value.(string); ok {
			return value, true
		} else if v, ok := a[key]; ok {
			return fmt.Sprintf("%v", v), true
		}
	}
	return "", false
}

// GetAsStringWithDefault gets the string value for the given key,
// or returns the given default value
// TODO: raise a warning if there was no entry found
func (a Attributes) GetAsStringWithDefault(key, defaultValue string) string {
	// check in predefined attributes
	if value, found := Predefined[key]; found {
		return value
	}
	if value, found := a[key]; found {
		if value, ok := value.(string); ok {
			return value
		}
	}
	return defaultValue
}

// GetAsBool returns the value of the key as a bool, or `false` if the key did not exist
// or if its value was not a bool
func (a Attributes) GetAsBool(key string) bool {
	if v, ok := a[key]; ok {
		if v, ok := v.(bool); ok {
			return v
		}
	}
	return false
}

// AddNonEmpty adds the given attribute if its value is non-nil and non-empty
// TODO: raise a warning if there was already a name/value
func (a Attributes) AddNonEmpty(key string, value interface{}) {
	// do not add nil or empty values
	if value == "" {
		return
	}
	a[key] = value
}

// Positionals returns all positional attributes, ie, the values for the keys in the form of `positional-<int>`
func (a Attributes) Positionals() [][]interface{} {
	result := make([][]interface{}, 0, len(a))
	i := 0
	for {
		i++
		if arg, ok := a["positional-"+strconv.Itoa(i)].([]interface{}); ok {
			result = append(result, arg)
			continue
		}
		break
	}
	return result
}

// NewAttributes retrieves the ElementID, ElementTitle and ElementInlineLink from the given slice of attributes
func NewAttributes(attributes interface{}) (Attributes, error) {
	if attributes == nil {
		return nil, nil
	}
	switch attrs := attributes.(type) {
	case []interface{}:
		// nested case, because of the grammar syntax,
		// eg: `attributes:(ElementAttribute* LiteralAttribute ElementAttribute*)`
		// which is used to ensure that a `LiteralAttribute` element is set amongst the attributes
		if len(attrs) == 0 {
			return nil, nil
		}
		result := Attributes{}
		for _, a := range attrs {
			r, err := NewAttributes(a)
			if err != nil {
				return nil, err
			}
			for k, v := range r {
				result[k] = v
			}
		}
		return result, nil
	case Attributes:
		return attrs, nil
	case map[string]interface{}:
		if len(attrs) == 0 {
			return nil, nil
		}
		result := Attributes{}
		for k, v := range attrs {
			result[k] = v
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected type of attributes: '%T'", attrs)
	}
}

// NewQuotedTextAttributes retrieves the attributes for QuotedText elements.
// We always pass in an array of Attributes.  They may only const of
// AttrRole or AttrID elements.  We coalesce the AttrRole elements into a
// single array. We keep only the first AttrID element.
func NewQuotedTextAttributes(attributes interface{}) (Attributes, error) {
	if attributes == nil {
		return nil, nil
	}
	switch attrs := attributes.(type) {
	case Attributes:
		return attrs, nil
	case []interface{}:
		// this is never empty, because we always have at least a role.
		result := Attributes{}
		for _, a := range attrs {
			for k, v := range a.(Attributes) {
				switch k {
				case AttrID:
					// The first ID set wins.
					if !result.Has(AttrID) {
						result[AttrID] = v
						result[AttrCustomID] = true
					}
				case AttrRole:
					result.AppendString(AttrRole, v.(string)) // grammar only generates string values
				}
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected type of attributes: '%T'", attrs)

	}
}

func resolveAlt(path Location) string {
	_, filename := filepath.Split(path.String())
	ext := filepath.Ext(filename)
	if ext != "" {
		return strings.TrimSuffix(filename, ext)
	}
	return filename
}

// // Has returns the true if an entry with the given key exists
// func (a Attributes) Has(key string) bool {
// 	_, ok := a[key]
// 	return ok
// }

// // Add adds the given attributes
// func (a Attributes) Add(attrs map[string]interface{}) {
// 	for k, v := range attrs {
// 		a.Add(k, v)
// 	}
// }

// // Add adds the given attribute if its value is non-nil
// // TODO: raise a warning if there was already a name/value
// func (a Attributes) Add(key string, value interface{}) {
// 	a[key] = value
// }

// // AddNonEmpty adds the given attribute if its value is non-nil and non-empty
// // TODO: raise a warning if there was already a name/value
// func (a Attributes) AddNonEmpty(key string, value interface{}) {
// 	// do not add nil or empty values
// 	if value == "" {
// 		return
// 	}
// 	a.Add(key, value)
// }

// // AddDeclaration adds the given attribute
// // TODO: raise a warning if there was already a name/value
// func (a Attributes) AddDeclaration(attr AttributeDeclaration) {
// 	a.Add(attr.Name, attr.Value)
// }

// // Delete deletes the given attribute
// func (a Attributes) Delete(attr AttributeReset) {
// 	delete(a, attr.Name)
// }
