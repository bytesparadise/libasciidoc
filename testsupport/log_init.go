package testsupport

import (
	"flag"
	"os"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.ErrorLevel)
	if debugMode() {
		log.SetLevel(log.DebugLevel)
		log.Info("Running test with logs in DEBUG level")
		log.SetFormatter(&log.TextFormatter{
			DisableQuote: true, // see https://github.com/sirupsen/logrus/issues/608#issuecomment-745137306
		})
	}

	// also, configuration for spew (when dumping structures to compare results)
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisablePointerMethods = true
	spew.Config.DisableUnexported = true
}

func debugMode() bool {
	debugMode := false
	flag.BoolVar(&debugMode, "debug", false, "when set, enables debug log messages")
	// if the `-debug` flag was passed and captured by the `flag.Parse`
	if debugMode {
		log.Info("`debug` flag found")
		return debugMode
	}
	// otherwise, check the OS args
	for _, arg := range os.Args {
		if arg == "-debug" {
			log.Info("`-debug` os env found")
			return true
		}
	}
	return false
}
