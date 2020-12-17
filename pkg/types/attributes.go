package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// ------------------------------------------
// Attributes
// ------------------------------------------

const (
	// AttrDocType the "doctype" attribute
	AttrDocType = "doctype"
	// AttrSyntaxHighlighter the attribute to define the syntax highlighter on code source blocks
	AttrSyntaxHighlighter = "source-highlighter"
	// AttrID the key to retrieve the ID
	AttrID = "id"
	// AttrIDPrefix the key to retrieve the ID Prefix
	AttrIDPrefix = "idprefix"
	// DefaultIDPrefix the default ID Prefix
	DefaultIDPrefix = "_"
	// AttrIDSeparator the key to retrieve the ID Separator
	AttrIDSeparator = "idseparator"
	// DefaultIDSeparator the default ID Separator
	DefaultIDSeparator = "_"
	// AttrTableOfContents the `toc` attribute at document level
	AttrTableOfContents = "toc"
	// AttrTableOfContentsLevels the document attribute which specifies the number of levels to display in the ToC
	AttrTableOfContentsLevels = "toclevels"
	// AttrNoHeader attribute to disable the rendering of document footer
	AttrNoHeader = "noheader"
	// AttrNoFooter attribute to disable the rendering of document footer
	AttrNoFooter = "nofooter"
	// AttrCustomID the key to retrieve the flag that indicates if the element ID is custom or generated
	AttrCustomID = "@customID"
	// AttrTitle the key to retrieve the title
	AttrTitle = "title"
	// AttrAuthors the key to the authors declared after the section level 0 (at the beginning of the doc)
	AttrAuthors = "authors"
	// AttrRevision the key to the revision declared after the section level 0 (at the beginning of the doc)
	AttrRevision = "revision"
	// AttrRole the key for a single role attribute
	AttrRole = "role"
	// AttrRoles the key to retrieve the roles attribute
	AttrRoles = "roles"
	// AttrOption the key for a single option attribute
	AttrOption = "option"
	// AttrOptions the key to retrieve the options attribute
	AttrOptions = "options"
	// AttrOpts alias for AttrOptions
	AttrOpts = "opts"
	// AttrInlineLink the key to retrieve the link
	AttrInlineLink = "link"
	// AttrQuoteAuthor attribute for the author of a verse
	AttrQuoteAuthor = "quoteAuthor"
	// AttrQuoteTitle attribute for the title of a verse
	AttrQuoteTitle = "quoteTitle"
	// AttrSource the `source` attribute for a source block or a source paragraph (this is a placeholder, ie, it does not expect any value for this attribute)
	AttrSource = "source"
	// AttrLanguage the `language` attribute for a source block or a source paragraph
	AttrLanguage = "language"
	// AttrLineNums the `linenums` attribute for a source block or a source paragraph
	AttrLineNums = "linenums"
	// AttrCheckStyle the attribute to mark the first element of an unordered list itemd as a checked or not
	AttrCheckStyle = "checkstyle"
	// AttrStart the `start` attribute in an ordered list
	AttrStart = "start"
	// AttrLevelOffset the `leveloffset` attribute used in file inclusions
	AttrLevelOffset = "leveloffset"
	// AttrLineRanges the `lines` attribute used in file inclusions
	AttrLineRanges = "lines"
	// AttrTagRanges the `tag`/`tags` attribute used in file inclusions
	AttrTagRanges = "tags"
	// AttrLastUpdated the "last updated" data in the document, i.e., the output/generation time
	AttrLastUpdated = "LastUpdated"
	// AttrImageAlt the image `alt` attribute
	AttrImageAlt = "alt"
	// AttrHeight the image `height` attribute
	AttrHeight = "height"
	// AttrImageWindow the `window` attribute, which becomes the target for the link
	AttrImageWindow = "window"
	// AttrImageAlign is for image alignment
	AttrImageAlign = "align"
	// AttrIconSize the icon `size`, and can be one of 1x, 2x, 3x, 4x, 5x, lg, fw
	AttrIconSize = "size"
	// AttrIconRotate the icon `rotate` attribute, and can be one of 90, 180, or 270
	AttrIconRotate = "rotate"
	// AttrIconFlip the icon `flip` attribute, and if set can be "horizontal" or "vertical"
	AttrIconFlip = "flip"
	// AttrUnicode local libasciidoc attribute to encode output as UTF-8 instead of ASCII.
	AttrUnicode = "unicode"
	// AttrCaption is the caption for block images, tables, and so forth
	AttrCaption = "caption"
	// AttrStyle block or list style
	AttrStyle = "style"
	// AttrInlineLinkText the text attribute (first positional) of links
	AttrInlineLinkText = "text"
	// AttrInlineLinkTarget the 'window' attribute
	AttrInlineLinkTarget = "window"
	// AttrWidth the `width` attribute used ior images, tables, and so forth
	AttrWidth = "width"
	// AttrFrame the frame used mostly for tables (all, topbot, sides, none)
	AttrFrame = "frame"
	// AttrGrid the grid (none, all, cols, rows) in tables
	AttrGrid = "grid"
	// AttrStripes controls table row background (even, odd, all, none, hover)
	AttrStripes = "stripes"
	// AttrFloat is for image or table float (text flows around)
	AttrFloat = "float"
	// AttrCols the table columns attribute
	AttrCols = "cols"
	// AttrPositionalPrefix positional parameter prefix (DEPRECATED - use `AttrPositionalIndex`)
	AttrPositionalPrefix = "@"
	// AttrPositionalIndex positional parameter index
	AttrPositionalIndex = "@positional-"
	// AttrPositional1 positional parameter 1
	AttrPositional1 = "@positional-1"
	// AttrPositional2 positional parameter 2
	AttrPositional2 = "@positional-2"
	// AttrPositional3 positional parameter 3
	AttrPositional3 = "@positional-3"
	// AttrVersionLabel labels the version number in the document
	AttrVersionLabel = "version-label"
	// AttrExampleCaption is the example caption
	AttrExampleCaption = "example-caption"
	// AttrFigureCaption is the figure (image) caption
	AttrFigureCaption = "figure-caption"
	// AttrTableCaption is the table caption
	AttrTableCaption = "table-caption"
	// AttrCautionCaption is the CAUTION caption
	AttrCautionCaption = "caution-caption"
	// AttrImportantCaption is the IMPORTANT caption
	AttrImportantCaption = "important-caption"
	// AttrNoteCaption is the NOTE caption
	AttrNoteCaption = "note-caption"
	// AttrTipCaption is the TIP caption
	AttrTipCaption = "tip-caption"
	// AttrWarningCaption is the TIP caption
	AttrWarningCaption = "warning-caption"
	// AttrSubstitutions the "subs" attribute to configure substitutions on delimited blocks and paragraphs
	AttrSubstitutions = "subs"
)

// Attribute is a key/value pair wrapper
type Attribute struct {
	Key   string
	Value interface{}
}

// Attributes the element attributes
// a map[string]interface{} with some utility methods
type Attributes map[string]interface{}

// NewAttributes retrieves the ElementID, ElementTitle and ElementInlineLink from the given slice of attributes
func NewAttributes(attributes ...interface{}) (Attributes, error) {
	if len(attributes) == 0 {
		return nil, nil
	}
	result := Attributes{}
	for _, attr := range attributes {
		switch attr := attr.(type) {
		case Attribute:
			result.Set(attr.Key, attr.Value)
		case Attributes: // when an there were multiple attributes, eg: `[quote,author,title]`
			result.SetAll(attr)
		case nil:
			// ignore
		default:
			return nil, fmt.Errorf("unexpected type of attribute: '%[1]T' (%[1]v)", attr)
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("initialized attributes:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, nil
}

func toAttributes(attrs interface{}) Attributes {
	if attrs, ok := attrs.(Attributes); ok {
		return attrs
	}
	return nil
}

func toAttributesWithMapping(attrs interface{}, mapping map[string]string) Attributes {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("processing attributes with mapping on")
		spew.Fdump(log.StandardLogger().Out, attrs)
	}
	if attrs, ok := attrs.(Attributes); ok {
		for source, target := range mapping {
			if v, exists := attrs[source]; exists {
				if v != nil {
					// (a bit hack-ish) make sure that `roles` is an `[]interface{}` if it came from a positional (1) attribute
					if source == AttrPositional1 && target == AttrRoles {
						v = []interface{}{v}
					}

					// do not override if already exists
					if _, exists := attrs[target]; !exists {
						// if key == value, replace value with `true`, so we don't have something
						// like `linenums=linenums`
						if target == v {
							attrs[target] = true
						} else {
							attrs[target] = v
						}
					}
				}
				delete(attrs, source)
			}
		}
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debug("processed attributes with mapping:")
			spew.Fdump(log.StandardLogger().Out, attrs)
		}
		if len(attrs) == 0 {
			return nil
		}
		return attrs
	}

	return nil
}

// NewAttributeGroup initializes an AttributeGroup
// We always pass in an array of Attributes.  Special handling for certain attributes:
//
// AttrID - these are strings, the first occurrence of this wins, and we set AttrCustomID
// AttrRoles - these are strings, we append them into an array
// AttrOptions - comma separated list, we split and put into a map
func NewAttributeGroup(attributes ...interface{}) (Attributes, error) {
	if len(attributes) == 0 {
		return nil, nil
	}
	result := Attributes{}
	for _, item := range attributes {
		if item == nil {
			continue
		}
		var attrs Attributes
		var err error
		switch item := item.(type) {
		case Attribute:
			attrs = Attributes{
				item.Key: item.Value,
			}
		case Attributes:
			attrs = item
		case []interface{}:
			attrs, err = NewAttributeGroup(item...)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unexpected type of attributes: %T", item)
		}
		for k, v := range attrs {
			switch k {
			case AttrID:
				// The first ID set wins.
				result[AttrID] = v
			case AttrRoles:
				result.append(AttrRoles, v)
			case AttrOptions, AttrOpts: // TODO handle the split in the `NewOptionAttribute` func
				if !result.Has(AttrOptions) {
					result[AttrOptions] = map[string]bool{}
				}
				m := result[AttrOptions].(map[string]bool)
				switch v := v.(type) {
				case string:
					for _, o := range strings.Split(v, ",") {
						m[o] = true
					}
				case map[string]bool:
					for o := range v {
						m[o] = v[o]
					}
				}
			default:
				result[k] = v
			}
		}
	}
	if len(result) == 0 {
		return nil, nil // don't retain groups with no attribute
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("new attribute group:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, nil
}

// HasAttributeWithValue checks that there is an entry for the given key/value pair
func HasAttributeWithValue(attributes interface{}, key string, value interface{}) bool {
	switch a := attributes.(type) {
	case Attribute:
		if a.Key == key && a.Value == value {
			return true
		}
	case Attributes:
		if v, ok := a[key]; ok && v == value {
			return true
		}
	case []interface{}:
		for _, attr := range a {
			if HasAttributeWithValue(attr, key, value) {
				return true
			}
		}
	}
	// make sure we return false here in case there was no match, so we can print the log msg
	log.Debugf("no attribute '%s:%v' found in %v", key, value, attributes)
	return false
}

// HasNotAttribute checks that there is no entry for the given key
func HasNotAttribute(attributes interface{}, key string) bool {
	switch a := attributes.(type) {
	case Attribute:
		if a.Key == key {
			return false
		}
	case Attributes:
		if _, ok := a[key]; ok {
			return false
		}
	case []interface{}:
		for _, attr := range a {
			if !HasNotAttribute(attr, key) {
				return false
			}
		}
	}
	return true
}

// NewPositionalAttribute returns a new attribute who key is the position in the group
func NewPositionalAttribute(index int, value interface{}) (Attribute, error) {
	result := Attribute{
		Key:   AttrPositionalIndex + strconv.Itoa(index),
		Value: value,
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("new positional attribute:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, nil
}

// NewOptionAttribute sets a boolean option.
func NewOptionAttribute(options interface{}) (Attribute, error) {
	return Attribute{
		Key:   AttrOption,
		Value: Reduce(options),
	}, nil
}

// NewNamedAttribute a named (or positional) element
func NewNamedAttribute(key string, value interface{}) (Attribute, error) {
	value = Reduce(value)
	if key == AttrOpts { // Handle the alias
		key = AttrOptions
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new named attribute '%s':", key)
		spew.Fdump(log.StandardLogger().Out, value)
	}
	return Attribute{
		Key:   key,
		Value: value,
	}, nil
}

// NewInlineIDAttribute initializes a new attribute map with a single entry for the ID using the given value
func NewInlineIDAttribute(id string) (Attribute, error) {
	log.Debugf("initializing a new inline ElementID with ID=%s", id)
	return Attribute{
		Key:   AttrID,
		Value: id,
	}, nil
}

// NewTitleAttribute initializes a new attribute map with a single entry for the title using the given value
func NewTitleAttribute(title interface{}) (Attribute, error) {
	log.Debugf("initializing a new Title attribute with content=%v", title)
	return Attribute{
		Key:   AttrTitle,
		Value: title,
	}, nil
}

// NewRoleAttribute initializes a new attribute map with a single entry for the title using the given value
func NewRoleAttribute(role interface{}) (Attribute, error) {
	role = Reduce(role)
	return Attribute{
		Key:   AttrRole,
		Value: role,
	}, nil
}

// NewIDAttribute initializes a new attribute map with a single entry for the ID using the given value
func NewIDAttribute(id interface{}) (Attribute, error) {
	log.Debugf("initializing a new ID attribute with ID='%v'", id)
	return Attribute{
		Key:   AttrID,
		Value: id,
	}, nil
}

// NewStyleAttribute initializes a new attribute map with a single entry for the style
// TODO: remove the `extras` parameter
func NewStyleAttribute(style interface{}, extras ...interface{}) (Attributes, error) {
	result := Attributes{
		AttrStyle: style,
	}
	extraAttrs, _ := NewExtraAttributes(extras...)
	result.SetAll(extraAttrs)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("new style/roles/options attributes:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, nil
}

// NewExtraAttributes return all the "extra" role and option attributes in an `Attributes` object
func NewExtraAttributes(extras ...interface{}) (Attributes, error) {
	roles := make([]interface{}, 0, len(extras))
	options := make([]interface{}, 0, len(extras))
	for _, extra := range extras {
		if a, ok := extra.(Attribute); ok {
			switch a.Key {
			case AttrRole:
				roles = append(roles, a.Value)
			case AttrOption:
				options = append(options, a.Value)
			}
		}
	}
	result := Attributes{}
	if len(roles) > 0 {
		result[AttrRoles] = roles
	}
	if len(options) > 0 {
		result[AttrOptions] = options
	}
	return result, nil
}

// NewAdmonitionAttribute initializes a new attribute map with a single entry for the admonition kind using the given value
func NewAdmonitionAttribute(kind string) (Attribute, error) {
	return Attribute{
		Key:   AttrStyle,
		Value: kind,
	}, nil
}

// NewQuoteAttributes initializes a new map of attributes for a verse paragraph
func NewQuoteAttributes(kind string, author, title interface{}) (Attributes, error) {
	result := make(map[string]interface{}, 3)
	switch kind {
	case "verse":
		result[AttrStyle] = Verse
	default:
		result[AttrStyle] = Quote
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

// NewLiteralBlockAttribute initializes a new attribute map with a single entry for the `literal` kind of block
func NewLiteralBlockAttribute() (Attribute, error) {
	return Attribute{
		Key:   AttrStyle,
		Value: Literal,
	}, nil
}

// NewPassthroughBlockAttribute initializes a new attribute map with a single entry for the `passthrough` kind of block
func NewPassthroughBlockAttribute() (Attribute, error) {
	return Attribute{
		Key:   AttrStyle,
		Value: Passthrough,
	}, nil
}

// NewExampleBlockAttribute initializes a new attribute map with a single entry for the `example`` kind of block
func NewExampleBlockAttribute() (Attribute, error) {
	return Attribute{
		Key:   AttrStyle,
		Value: Example,
	}, nil
}

// NewListingBlockAttribute initializes a new attribute map with a single entry for the `listing`` kind of block
func NewListingBlockAttribute() (Attribute, error) {
	return Attribute{
		Key:   AttrStyle,
		Value: Listing,
	}, nil
}

// NewSourceAttributes initializes a new attribute map with two entries, one for the kind of element ("source") and another optional one for the language of the source code
func NewSourceAttributes(language interface{}, option interface{}, others ...interface{}) (Attributes, error) {
	result := Attributes{
		AttrStyle: Source,
	}
	if language := Reduce(language); language != nil {
		result[AttrLanguage] = language
	}
	if option, ok := option.(string); ok {
		result[AttrSourceBlockOption] = strings.TrimSpace(option)
	}
	for _, other := range others {
		result = result.SetAll(other)
	}
	return result, nil
}

// Set adds the given attribute to the current ones
// with some `key` replacements/grouping (Role->Roles and Option->Options)
func (a Attributes) Set(key string, value interface{}) Attributes {
	log.Debugf("setting attribute %s=%v", key, value)
	if a == nil {
		a = Attributes{}
	}
	switch key {
	case AttrRole:
		if roles, ok := a[AttrRoles].([]interface{}); ok {
			log.Debugf("appending role to existin one(s): %v", value)
			a[AttrRoles] = append(roles, value)
		} else {
			log.Debugf("setting first role: %v (%T)", value, a[AttrRoles])
			a[AttrRoles] = []interface{}{value}
		}
	case AttrRoles:
		if r, ok := value.([]interface{}); ok { // value should be an []interface{}
			if roles, ok := a[AttrRoles].([]interface{}); ok {
				log.Debugf("appending role to existin one(s): %v", value)
				a[AttrRoles] = append(roles, r...)
			} else {
				log.Debugf("setting first role: %v (%T)", value, a[AttrRoles])
				a[AttrRoles] = r
			}
		}
	case AttrOption:
		if options, ok := a[AttrOptions].([]interface{}); ok {
			a[AttrOptions] = append(options, value)
		} else {
			a[AttrOptions] = []interface{}{value}
		}
	default:
		a[key] = value
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		spew.Fdump(log.StandardLogger().Out, a)

	}
	return a
}

// SetAll adds the given attributes to the current ones
func (a Attributes) SetAll(attr interface{}) Attributes {
	log.Debug("setting attributes")
	switch attr := attr.(type) {
	case Attribute:
		if a == nil {
			a = Attributes{}
		}
		a[attr.Key] = attr.Value
	case Attributes:
		if len(attr) == 0 {
			return a
		}
		if a == nil {
			a = Attributes{}
		}
		for k, v := range attr {
			a.Set(k, v)
		}
	case []interface{}:
		for i := range attr {
			a = a.SetAll(attr[i])
		}
	}
	// will return nil if it was nil before and nothing was added
	return a
}

// Has returns the true if an entry with the given key exists
func (a Attributes) Has(key string) bool {
	_, ok := a[key]
	return ok
}

// HasOption returns true if the option is set.
func (a Attributes) HasOption(key string) bool {
	// in block attributes: search key in the `Options`
	if opts, ok := a[AttrOptions].([]interface{}); ok {
		for _, opt := range opts {
			if opt == key {
				return true
			}
		}
	}
	// in document attributes: direct lookup
	if a.Has(key) {
		return true
	}
	return false
}

// sets the value as a singular string value if it did not exist yet,
// or move the existing value in a slice and append the new one
func (a Attributes) append(key string, value interface{}) {
	v, found := a[key]
	if !found {
		a[key] = value
		return
	}
	switch v := v.(type) {
	case []interface{}:
		switch value := value.(type) {
		case []interface{}:
			a[key] = append(v, value...)
		default:
			a[key] = append(v, value) // just append the new value into the slice
		}
	default:
		switch value := value.(type) {
		case []interface{}:
			a[key] = append([]interface{}{v}, value...)
		default:
			a[key] = []interface{}{v, value} // move existing value in a slice, along with the new one
		}
	}
}

// GetAsString gets the string value for the given key (+ `true`),
// or empty string (+ `false`) if none was found
func (a Attributes) GetAsString(key string) (string, bool) {
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
func (a Attributes) GetAsStringWithDefault(key, defaultValue string) string {
	if value, found := a[key]; found {
		if value == nil {
			return "" // nil present means attribute was reset
		}
		if value, ok := value.(string); ok {
			return value
		} else if v, ok := a[key]; ok {
			return fmt.Sprintf("%v", v)
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
