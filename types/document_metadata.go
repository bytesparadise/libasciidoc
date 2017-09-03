package types

// DocumentMetadata the document metadata
type DocumentMetadata map[string]string

const (
	title string = "title"
)

// GetTitle retrieves the document title in its metadata, or returns nil if the title was not specified
func (m DocumentMetadata) GetTitle() *string {
	if t, ok := m[title]; ok {
		return &t
	}
	return nil
}

// SetTitle sets the title in the document metadata
func (m DocumentMetadata) SetTitle(t string) {
	m[title] = t
}
