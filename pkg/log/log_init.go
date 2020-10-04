package log

import (
	log "github.com/sirupsen/logrus"
)

// Setup configures the logger
func Setup(level log.Level) {
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		EnvironmentOverrideColors: true,
		DisableLevelTruncation:    true,
		DisableTimestamp:          true,
	})
}
