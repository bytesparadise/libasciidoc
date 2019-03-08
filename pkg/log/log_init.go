package log

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

// initializes the level for the logger, using the optional '-debug' flag to activate the logs in 'debug' level.
// Other tests must import this 'test' package even if unused, using:
// import _ "github.com/bytesparadise/libasciidoc/pkg/log"
func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	if debugMode() {
		log.SetLevel(log.DebugLevel)
		log.Warn("Running test with logs in DEBUG level")
	}
	log.SetFormatter(&log.TextFormatter{FullTimestamp: false})
}

func debugMode() bool {
	debugMode := false
	flag.BoolVar(&debugMode, "debug", false, "when set, enables debug log messages")
	if !flag.Parsed() {
		flag.Parse()
	}
	// if the `-debug` flag was passed and captured by the `flag.Parse`
	if debugMode {
		// log.Info("`debug` flag found")
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
