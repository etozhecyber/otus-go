package utilities

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

//SetupLogger setting log level and logfile
func SetupLogger(config Config) error {
	switch strings.ToLower(config.LogLevel) {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}

	//setting log file
	logfile, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(logfile)
	return nil
}
