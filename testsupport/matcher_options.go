package testsupport

type filenameMatcher interface {
	setFilename(string)
}

// FilenameOption an option to set the name of the file being treated by the matcher
type FilenameOption func(m filenameMatcher)

// WithFilename configures the filename, which can be absolute or relative
func WithFilename(filename string) FilenameOption {
	return func(m filenameMatcher) {
		m.setFilename(filename)
	}
}

// BecomeDraftDocumentOption an option to configure the BecomeDraftDocument matcher
type BecomeDraftDocumentOption func(c *drafDocumentParserConfig)

// WithoutPreprocessing disables document preprocessing
func WithoutPreprocessing() BecomeDraftDocumentOption {
	return func(c *drafDocumentParserConfig) {
		c.preprocessing = false
	}
}
