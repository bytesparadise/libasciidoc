package test

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// initializes the level for the logger, using the optional '-debug' flag to activate the logs in 'debug' level.
// Other tests must import this 'test' package even if unused, using:
// import _ "github.com/bytesparadise/libasciidoc/test"
func init() {
	debugMode := false
	flag.BoolVar(&debugMode, "debug", false, "when set, enables debug log messages")
	// flag.Parse()
	fmt.Printf("Args: %v\n", flag.Args())
	if debugMode {
		log.SetLevel(log.DebugLevel)
		log.Warn("Running test with logs in debug-level")
	}
	log.SetFormatter(&log.TextFormatter{FullTimestamp: false})
}
