module github.com/bytesparadise/libasciidoc

go 1.17

require (
	github.com/alecthomas/chroma/v2 v2.3.0
	github.com/davecgh/go-spew v1.1.1
	github.com/felixge/fgtrace v0.1.0
	github.com/google/go-cmp v0.5.9
	github.com/mna/pigeon v1.1.0
	github.com/onsi/ginkgo/v2 v2.7.0
	github.com/onsi/gomega v1.24.2
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.6.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/DataDog/gostackparse v0.5.0 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/google/pprof v0.0.0-20210407192527-94a9f03dee38 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// include support for disabling unexported fields
// TODO: still needed?
replace github.com/davecgh/go-spew => github.com/flw-cn/go-spew v1.1.2-0.20200624141737-10fccbfd0b23
