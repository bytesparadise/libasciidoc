package types

// DocumentAttributes the document attributes
type DocumentAttributes map[string]interface{}

const (
	title string = "title"
)

// GetTitle retrieves the document title in its metadata, or returns nil if the title was not specified
func (m DocumentAttributes) GetTitle() *string {
	if t, ok := m[title]; ok {
		title := t.(string)
		return &title
	}
	return nil
}

// SetTitle sets the title in the document attributes
func (m DocumentAttributes) SetTitle(t string) {
	m[title] = t
}

// AddAll adds all given attributes
func (m DocumentAttributes) AddAll(attributes map[string]interface{}) {
	for name, value := range attributes {
		// TODO: raise a warning if there was already a name/value
		m[name] = value
	}
}

// Add adds the given attribute
func (m DocumentAttributes) Add(a DocumentAttributeDeclaration) {
	// TODO: raise a warning if there was already a name/value
	m[a.Name] = a.Value
}

// Reset resets the given attribute
func (m DocumentAttributes) Reset(a DocumentAttributeReset) {
	delete(m, a.Name)
}

// Get gets the given value for the given attribute, or nil if none was found
func (m DocumentAttributes) Get(a DocumentAttributeSubstitution) interface{} {
	// TODO: raise a warning if there was no entry found
	if value, ok := m[a.Name]; ok {
		return &value
	}
	return nil
}
