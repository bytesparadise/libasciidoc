package types

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// DocumentAttributes the document attributes
type DocumentAttributes map[string]interface{}

const (
	title     string = "doctitle"
	toclevels string = "toclevels"
)

// GetTOCLevels returns the value of the `toclevels` attribute if it was specified,
// or `2` as the default value
func (m DocumentAttributes) GetTOCLevels() (*int, error) {
	if levels, exists := m["toclevels"]; exists {
		if levels, ok := levels.(int); ok {
			return &levels, nil
		}
		if _, ok := levels.(string); ok {
			levels, err := strconv.Atoi(levels.(string))
			if err != nil {
				return nil, errors.Wrapf(err, "the value of the 'toclevels' attribute is not an integer: %T", levels)
			}
			return &levels, nil
		}
		return nil, errors.Errorf("the value of the 'toclevels' attribute is not an integer: %T", levels)
	}
	// return default value if the "toclevels" doc attribute was not specified
	defaultLevels := 2
	return &defaultLevels, nil
}

// HasAuthors returns `true` if the document has one or more authors, `false` otherwise.
func (m DocumentAttributes) HasAuthors() bool {
	_, exists := m["author"]
	return exists
}

// GetTitle retrieves the document title in its metadata, or returns nil if the title was not specified
func (m DocumentAttributes) GetTitle() (SectionTitle, error) {
	if t, found := m[title]; found {
		if t, ok := t.(SectionTitle); ok {
			return t, nil
		}
		return SectionTitle{}, errors.Errorf("document title type is not valid: %T", t)
	}
	return SectionTitle{}, nil
}

// Add adds the given attribute if its value is non-nil
// TODO: raise a warning if there was already a name/value
func (m DocumentAttributes) Add(key string, value interface{}) {
	// do not add nil or empty values
	if value == nil {
		return
	}
	v := reflect.ValueOf(value)
	k := v.Kind()
	// if the argument is a pointer, then retrive the value it points to
	if k == reflect.Ptr {
		if v.Elem().IsValid() {
			m[key] = v.Elem().Interface()
		}
	} else {
		m[key] = value
	}
}

// AddNonEmpty adds the given attribute if its value is non-nil and non-empty
// TODO: raise a warning if there was already a name/value
func (m DocumentAttributes) AddNonEmpty(key string, value interface{}) {
	// do not add nil or empty values
	if value == "" {
		return
	}
	m.Add(key, value)
}

// AddAttribute adds the given attribute
// TODO: raise a warning if there was already a name/value
func (m DocumentAttributes) AddAttribute(attr DocumentAttributeDeclaration) {
	// do not add nil values
	// if attr == nil {
	// 	return
	// }
	m.Add(attr.Name, attr.Value)
}

// Reset resets the given attribute
func (m DocumentAttributes) Reset(attr DocumentAttributeReset) {
	delete(m, attr.Name)
}

// GetAsString gets the string value for the given key, or nil if none was found
func (m DocumentAttributes) GetAsString(key string) *string {
	// TODO: raise a warning if there was no entry found
	if value, found := m[key]; found {
		strValue := value.(string)
		return &strValue
	}
	return nil
}
