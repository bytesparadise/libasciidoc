package renderer

import (
	"time"

	"github.com/pkg/errors"
)

//Options the options when rendering a document
type Options map[string]interface{}

const (
	//LastUpdated the key to specify the last update of the document to render.
	// Can be a string or a time, which will be formatted using the 2006/01/02 15:04:05 MST` pattern
	LastUpdated string = "LastUpdated"
	//IncludeHeaderFooter a bool value to indicate if the header and footer should be rendered
	IncludeHeaderFooter string = "IncludeHeaderFooter"
)

// LastUpdated returns the value of the 'LastUpdated' Option if it was present,
// otherwise it returns the current time using the `2006/01/02 15:04:05 MST` format
func (o Options) LastUpdated() (*string, error) {
	if lastUpdated, ok := o[LastUpdated]; ok {
		switch lastUpdated := lastUpdated.(type) {
		case string:
			return &lastUpdated, nil
		case time.Time:
			result := lastUpdated.Format("2006/01/02 15:04:05 MST")
			return &result, nil
		default:
			return nil, errors.Errorf("`LastUpdated` option is not in a valid format: %T", lastUpdated)
		}
	}
	result := time.Now().Format("2006/01/02 15:04:05 MST")
	return &result, nil
}

// IncludeHeaderFooter returns the value of the 'LastUpdated' Option if it was present,
// otherwise it returns `false``
func (o Options) IncludeHeaderFooter() (*bool, error) {
	if includeHeaderFooter, ok := o[IncludeHeaderFooter]; ok {
		switch includeHeaderFooter := includeHeaderFooter.(type) {
		case bool:
			return &includeHeaderFooter, nil
		default:
			return nil, errors.Errorf("`IncludeHeaderFooter` option is not in a valid format: %T", includeHeaderFooter)
		}
	}
	result := false
	return &result, nil
}
