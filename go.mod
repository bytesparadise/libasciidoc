module github.com/bytesparadise/libasciidoc

go 1.11

require (
	github.com/alecthomas/chroma v0.7.1
	github.com/davecgh/go-spew v1.1.1
	github.com/google/go-cmp v0.5.5
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mna/pigeon v1.1.0
	github.com/modocache/gover v0.0.0-20171022184752-b58185e213c5 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/pkg/errors v0.9.1
	github.com/sergi/go-diff v1.0.0
	github.com/sirupsen/logrus v1.7.0
	github.com/sozorogami/gover v0.0.0-20171022184752-b58185e213c5
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0
)

// include support for disabling unexported fields 
// TODO: still needed?
replace github.com/davecgh/go-spew => github.com/flw-cn/go-spew v1.1.2-0.20200624141737-10fccbfd0b23 
