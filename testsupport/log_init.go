package testsupport

import (
	"flag"
	"os"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func init() {
	lvl := parseLogLevel()
	log.SetLevel(lvl)
	log.Warnf("Running test with logs in '%s' level", lvl.String())
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true, // see https://github.com/sirupsen/logrus/issues/608#issuecomment-745137306
	})

	// also, configuration for spew (when dumping structures to compare results)
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisablePointerMethods = true
	spew.Config.DisableUnexported = true
}

func parseLogLevel() log.Level {
	var logLevel string
	// needed to let ginkgo parse the flag, otherwise, `ginkgo -- --loglevel=...` will fail with `flag provided but not defined: -loglevel`
	flag.StringVar(&logLevel, "loglevel", "error", "log level to set [debug|info|warn|error|fatal|panic]")
	// parse with a custom flagset in which all other flags (ginkgo's) are ignored
	f := pflag.NewFlagSet("passthroughs", pflag.ContinueOnError)
	f.ParseErrorsWhitelist.UnknownFlags = true
	f.StringVarP(&logLevel, "loglevel", "l", "error", "log level to set [debug|info|warn|error|fatal|panic]")
	if err := f.Parse(os.Args[1:]); err != nil {
		panic(err)
	}
	lvl, err := log.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	return lvl
}
