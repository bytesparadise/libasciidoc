package types

// DocumentAttributes the document metadata
type DocumentAttributes map[string]interface{}

const (
	title string = "title"
)

// GetTitle retrieves the document title in its metadata, or returns nil if the title was not specified
func (m DocumentAttributes) GetTitle() *string {
	if t, ok := m[title]; ok {
		if t, ok := t.(string); ok {
			return &t
		}
	}
	return nil
}

// SetTitle sets the title in the document metadata
func (m DocumentAttributes) SetTitle(t string) {
	m[title] = t
}

// Add adds the given attribute
func (m DocumentAttributes) Add(a *DocumentAttribute) {
	// TODO: raise a warning if there was already a name/value
	m[a.Name] = a.Value
}
