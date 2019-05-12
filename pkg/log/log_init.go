package log

import (
	log "github.com/sirupsen/logrus"
)

// Setup configures the logger
func Setup() {
	customFormatter := new(log.TextFormatter)
	customFormatter.EnvironmentOverrideColors = true
	customFormatter.DisableLevelTruncation = true
	customFormatter.DisableTimestamp = true
	log.SetFormatter(customFormatter)
}
