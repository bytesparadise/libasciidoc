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
	// AttrDocType the "description" attribute
	AttrDescription = "description"
	// AttrSyntaxHighlighter the attribute to define the syntax highlighter on code source blocks
	AttrSyntaxHighlighter = "source-highlighter"
	// AttrChromaClassPrefix the class prefix used by Chroma when rendering source code (default: `tok-`)
	AttrChromaClassPrefix = "chroma-class-prefix"
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
	// AttrNumbered the `numbered` attribute to trigger section numbering at renderding time
	AttrNumbered = "numbered"
	// AttrSectionNumbers the `sectnums` attribute to trigger section numbering at renderding time (an alias for `numbered`)
	AttrSectionNumbering = "sectnums"
	// AttrTableOfContents the `toc` attribute at document level
	AttrTableOfContents = "toc"
	// AttrTableOfContentsLevels the document attribute which specifies the number of levels to display in the ToC
	AttrTableOfContentsLevels = "toclevels"
	// AttrTableOfContentsTitle the document attribute which specifies the title of the table of contents
	AttrTableOfContentsTitle = "toc-title"
	// AttrNoHeader attribute to disable the rendering of document footer
	AttrNoHeader = "noheader"
	// AttrNoFooter attribute to disable the rendering of document footer
	AttrNoFooter = "nofooter"
	// AttrCustomID the key to retrieve the flag that indicates if the element ID is custom or generated
	// AttrCustomID = "@customID"
	// AttrTitle the key to retrieve the title
	AttrTitle = "title"
	// AttrAuthors the key to the authors declared after the section level 0 (at the beginning of the doc)
	AttrAuthors = "authors"
	// AttrAuthor the key to the author's full name declared as a standalone attribute
	AttrAuthor = "author"
	// AttrAuthor the key to the author's email address declared as a standalone attribute
	AttrEmail = "email"
	// AttrRevision the key to the revision declared after the section level 0 (at the beginning of the doc)
	// or as a standalone attribute
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
	// AttrCheckStyle the attribute to mark the first element of an unordered list item as a checked or not
	AttrCheckStyle = "checkstyle"
	// AttrInteractive the attribute to mark the first element of an unordered list item as n interactive checkbox or not
	// (paired with `AttrCheckStyle`)
	AttrInteractive = "interactive"
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
	// AttrStyle paragraph, block or list style
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
	// AttrAutoWidth the `autowidth` attribute on a table
	AttrAutoWidth = "autowidth"
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
	// AttrImagesDir the `imagesdir` attribute
	AttrImagesDir = "imagesdir"
	// AttrXRefLabel the label of a cross reference
	AttrXRefLabel = "xrefLabel"
	// AttrExperimental a flag to enable experiment macros (for UI)
	AttrExperimental = "experimental"
	// AttrButtonLabel the label of a button
	AttrButtonLabel = "label"
	// AttrHardBreaks the attribute to set on a paragraph to render with hard breaks on each line
	AttrHardBreaks = "hardbreaks"
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
	positionalIndex := 0
	for _, attr := range attributes {
		switch attr := attr.(type) {
		case *PositionalAttribute:
			positionalIndex++
			attr.Index = positionalIndex
			result.Set(attr.Key(), attr.Value)
		case *Attribute:
			result.Set(attr.Key, attr.Value)
		default:
			return nil, fmt.Errorf("unexpected type of attribute: '%[1]T' (%[1]v)", attr)
		}
	}
	return result, nil
}

func (a Attributes) Clone() Attributes {
	result := Attributes{}
	for k, v := range a {
		result[k] = v
	}
	return result
}
func MergeAttributes(attributes ...interface{}) (Attributes, error) {
	if len(attributes) == 0 {
		return nil, nil
	}
	result := Attributes{}
	for _, attr := range attributes {
		switch attr := attr.(type) {
		case *Attribute:
			result.Set(attr.Key, attr.Value)
		case Attributes: // when an there were multiple attributes, eg: `[quote,author,title]`
			result.AddAll(attr)
		default:
			return nil, fmt.Errorf("unexpected type of attribute: '%[1]T' (%[1]v)", attr)
		}
	}
	return result, nil
}

// NilIfEmpty returns `nil` if this `attributes` is empty
func (a Attributes) AddAll(others Attributes) Attributes {
	if others == nil {
		return a
	}
	if a == nil {
		a = Attributes{}
	}
	for k, v := range others {
		a.Set(k, v)
	}
	return a
}

func toAttributesWithMapping(attrs interface{}, mapping map[string]string) Attributes {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("processing attributes with mapping on\n%s", spew.Sdump(attrs))
	// }
	if attrs, ok := attrs.(Attributes); ok {
		for source, target := range mapping {
			if v, exists := attrs[source]; exists {
				if v != nil && v != "" { // nil and empty values are discarded (ie, not mapped to target key)
					// (a bit hack-ish) make sure that `roles` is an `[]interface{}` if it came from a positional (1) attribute
					if source == AttrPositional1 && target == AttrRoles {
						v = Roles{v}
					}

					// do not override if already exists
					if _, exists := attrs[target]; !exists {
						// if key == value, replace value with `true`,
						// so we don't have something such as `linenums=linenums`
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
		if len(attrs) == 0 {
			return nil
		}
		return attrs
	}
	return nil
}

// HasAttributeWithValue checks that there is an entry for the given key/value pair
func HasAttributeWithValue(attributes interface{}, key string, value interface{}) bool {
	switch a := attributes.(type) {
	case *Attribute:
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
	// log.Debugf("no attribute '%s:%v' found in %v", key, value, attributes)
	return false
}

// HasNotAttribute checks that there is no entry for the given key
func HasNotAttribute(attributes interface{}, key string) bool {
	switch a := attributes.(type) {
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

// PositionalAttribute an attribute whose key will be determined by its position,
// and which depends on the element it applies to.
type PositionalAttribute struct {
	Index int
	Value interface{}
}

// NewPositionalAttribute returns a new attribute who key is the position in the group
func NewPositionalAttribute(value interface{}) (*PositionalAttribute, error) {
	value = Reduce(value, strings.TrimSpace)
	// log.Debugf("new positional attribute: '%s'", value)
	return &PositionalAttribute{
		Value: value,
	}, nil
}

// Key returns the "temporary" key, based on the attribute index.
func (a *PositionalAttribute) Key() string {
	return AttrPositionalIndex + strconv.Itoa(a.Index)
}

type Options []interface{} // more explicit than `[]interface{}`, and to bypass the `Reduce` func that would merge all roles into a single string :/

// NewOptionAttribute sets a boolean option.
func NewOptionAttribute(option interface{}) (*Attribute, error) {
	option = Reduce(option, strings.TrimSpace)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("new option attribute: '%s'", spew.Sdump(option))
	}
	return &Attribute{
		Key:   AttrOption,
		Value: option,
	}, nil
}

// NewNamedAttribute a named (or positional) element
func NewNamedAttribute(key string, value interface{}) (*Attribute, error) {
	// value = Reduce(value, strings.TrimSpace)
	key = strings.TrimSpace(key)
	if key == AttrOpts { // Handle the alias
		key = AttrOptions
	}
	return &Attribute{
		Key:   key,
		Value: value,
	}, nil
}

// NewTitleAttribute initializes a new attribute map with a single entry for the title using the given value
func NewTitleAttribute(value interface{}) (*Attribute, error) {
	return NewNamedAttribute(AttrTitle, value)
}

// NewRoleAttribute initializes a new attribute map with a single entry for the title using the given value
func NewRoleAttribute(value interface{}) (*Attribute, error) {
	return NewNamedAttribute(AttrRole, value)
}

type Roles []interface{} // more explicit than `[]interface{}`, and to bypass the `Reduce` func that would merge all roles into a single string :/

// NewIDAttribute initializes a new attribute map with a single entry for the ID using the given value
func NewIDAttribute(value interface{}) (*Attribute, error) {
	return NewNamedAttribute(AttrID, value)
}

// Set adds the given attribute to the current ones
// with some `key` replacements/grouping (Role->Roles and Option->Options)
// returns the new `Attributes` if the current instance was `nil`
func (a Attributes) Set(key string, value interface{}) Attributes {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("setting attribute %s=%s", key, spew.Sdump(value))
	// }
	if a == nil {
		a = Attributes{}
	}
	switch key {
	// case AttrID:
	// 	a[key] = value
	// 	a[AttrCustomID] = true
	case AttrRole:
		if roles, ok := a[AttrRoles].(Roles); ok {
			a[AttrRoles] = append(roles, value)
		} else {
			a[AttrRoles] = Roles{value}
		}
	case AttrRoles:
		if r, ok := value.(Roles); ok {
			if roles, ok := a[AttrRoles].(Roles); ok {
				a[AttrRoles] = append(roles, r...)
			} else {
				a[AttrRoles] = Roles(r)
			}
		}
	case AttrOption: // move into `options`
		if options, ok := a[AttrOptions].(Options); ok {
			a[AttrOptions] = append(options, value)
		} else {
			a[AttrOptions] = Options{value}
		}
	case AttrOptions: // make sure the value is wrapped into a []interface{}
		switch v := value.(type) {
		case Options:
			a[AttrOptions] = v
		case string:
			values := strings.Split(v, ",")
			options := make(Options, len(values))
			for i, v := range values {
				options[i] = v
			}
			a[AttrOptions] = options
		default:
			a[AttrOptions] = Options{value}
		}
	default:
		a[key] = value
	}
	return a
}

func (a Attributes) Unset(key string) {
	delete(a, key)
}

// SetAll adds the given attributes to the current ones
func (a Attributes) SetAll(attr interface{}) Attributes {
	switch attr := attr.(type) {
	case *Attribute:
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
	if opts, ok := a[AttrOptions].(Options); ok {
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

// GetAsString gets the string value for the given key (+ `true`),
// or empty string (+ `false`) if none was found
func (a Attributes) GetAsString(k string) (string, bool) {
	if value, ok := a[k].(string); ok {
		return value, true
	}
	// log.Warnf("no entry for key '%s' in attributes, or value is not a string: '%T'", k, a[k])
	return "", false
}

// GetAsIntWithDefault gets the int value for the given key ,
// or default if none was found,
func (a Attributes) GetAsIntWithDefault(key string, defaultValue int) int {
	switch v := a[key].(type) {
	case int:
		return v
	case string:
		if result, err := strconv.Atoi(v); err == nil {
			return result
		}
	}
	return defaultValue
}

// GetAsBoolWithDefault gets the bool value for the given key ,
// or default if none was found,
func (a Attributes) GetAsBoolWithDefault(key string, defaultValue bool) bool {
	switch v := a[key].(type) {
	case bool:
		return v
	case string:
		if result, err := strconv.ParseBool(v); err == nil {
			return result
		}
	}
	return defaultValue
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
