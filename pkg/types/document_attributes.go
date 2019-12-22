package types

import (
	"reflect"
)

// DocumentAttributes the document attributes
type DocumentAttributes map[string]interface{}

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

// AddAll adds the given attributes
func (a DocumentAttributes) AddAll(attrs map[string]interface{}) DocumentAttributes {
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

// GetAsStringWithDefault gets the string value for the given key,
// or returns the given default value
func (a DocumentAttributes) GetAsStringWithDefault(key, defaultValue string) string {
	// TODO: raise a warning if there was no entry found
	if value, found := a[key]; found {
		if value, ok := value.(string); ok {
			return value
		}
	}
	return defaultValue
}
