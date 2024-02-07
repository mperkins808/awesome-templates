package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func CleanupPID(pidFile string) {
	if pidFile != "" {
		err := os.Remove(pidFile)
		if err != nil {
			log.Errorf("could not delete pid file: %v", err)
		}
	}
}
