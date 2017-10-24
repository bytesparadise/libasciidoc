package types

import "reflect"

// DocumentAttributes the document attributes
type DocumentAttributes map[string]interface{}

const (
	title string = "doctitle"
)

// HasAuthors returns `true` if the document has one or more authors, `false` otherwise.
func (m DocumentAttributes) HasAuthors() bool {
	_, author := m["author"]
	return author
}

// GetTitle retrieves the document title in its metadata, or returns nil if the title was not specified
func (m DocumentAttributes) GetTitle() *string {
	if t, ok := m[title]; ok {
		title := t.(string)
		return &title
	}
	return nil
}

// Add adds the given attribute
// TODO: raise a warning if there was already a name/value
func (m DocumentAttributes) Add(key string, value interface{}) {
	// do not add nil values
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

// AddAttribute adds the given attribute
// TODO: raise a warning if there was already a name/value
func (m DocumentAttributes) AddAttribute(attr *DocumentAttributeDeclaration) {
	// do not add nil values
	if attr == nil {
		return
	}
	m.Add(attr.Name, attr.Value)
}

// Reset resets the given attribute
func (m DocumentAttributes) Reset(a DocumentAttributeReset) {
	delete(m, a.Name)
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
