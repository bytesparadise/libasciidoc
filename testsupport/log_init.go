package testsupport

import (
	"flag"
	"os"

	logsupport "github.com/bytesparadise/libasciidoc/pkg/log"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func init() {
	logsupport.Setup(logrus.ErrorLevel)
	if debugMode() {
		log.SetLevel(log.DebugLevel)
		log.Info("Running test with logs in DEBUG level")
	}
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
