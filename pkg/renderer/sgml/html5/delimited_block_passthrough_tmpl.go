package html5

const (

	// the name here is weird because "pass" as a prefix triggers a false security warning
	passthroughBlock = "{{ .Content }}\n" // nolint (avoids a Gosec false positive because the const name starts with 'pass')
)
