package types

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// DocumentAttributes the document attributes
type DocumentAttributes map[string]interface{}

const (
	title     string = AttrTitle
	toclevels string = "toclevels"
)

// GetTOCLevels returns the value of the `toclevels` attribute if it was specified,
// or `2` as the default value
func (a DocumentAttributes) GetTOCLevels() (*int, error) {
	if levels, exists := a[toclevels]; exists {
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

// Has returns the true if an entry with the given key exists
func (a DocumentAttributes) Has(key string) bool {
	_, ok := a[key]
	return ok
}

// HasAuthors returns `true` if the document has one or more authors, `false` otherwise.
func (a DocumentAttributes) HasAuthors() bool {
	_, exists := a["author"]
	return exists
}

// HasTitle returns `true` if the document has a title, ie, a section with level = 0
func (a DocumentAttributes) HasTitle() bool {
	_, found := a[title]
	return found
}

// GetTitle retrieves the document title in its metadata, or returns nil if the title was not specified
func (a DocumentAttributes) GetTitle() (SectionTitle, error) {
	if t, found := a[title]; found {
		if t, ok := t.(SectionTitle); ok {
			return t, nil
		}
		return SectionTitle{}, errors.Errorf("document title type is not valid: %T", t)
	}
	return SectionTitle{}, nil
}

// AddAll adds the given attributes
func (a DocumentAttributes) AddAll(attrs DocumentAttributes) DocumentAttributes {
	for k, v := range attrs {
		a.Add(k, v)
	}
	return a
}

// Add adds the given attribute if its value is non-nil
// TODO: raise a warning if there was already a name/value
func (a DocumentAttributes) Add(key string, value interface{}) DocumentAttributes {
	// do not add nil or empty values
	if value == nil {
		return a
	}
	v := reflect.ValueOf(value)
	k := v.Kind()
	// if the argument is a pointer, then retrieve the value it points to
	if k == reflect.Ptr {
		if v.Elem().IsValid() {
			a[key] = v.Elem().Interface()
		}
	} else {
		a[key] = value
	}
	return a
}

// AddNonEmpty adds the given attribute if its value is non-nil and non-empty
// TODO: raise a warning if there was already a name/value
func (a DocumentAttributes) AddNonEmpty(key string, value interface{}) {
	// do not add nil or empty values
	if value == "" {
		return
	}
	a.Add(key, value)
}

// AddDeclaration adds the given attribute
// TODO: raise a warning if there was already a name/value
func (a DocumentAttributes) AddDeclaration(attr DocumentAttributeDeclaration) {
	a.Add(attr.Name, attr.Value)
}

// Reset resets the given attribute
func (a DocumentAttributes) Reset(attr DocumentAttributeReset) {
	delete(a, attr.Name)
}

// GetAsString gets the string value for the given key (+ `true`),
// or empty string (+ `false`) if none was found
func (a DocumentAttributes) GetAsString(key string) (string, bool) {
	// TODO: raise a warning if there was no entry found
	if value, found := a[key]; found {
		if value, ok := value.(string); ok {
			return value, true
		}
	}
	return "", false
}
