package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func NewLogger(name string) (*logrus.Logger, error) {
	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter)                     // default
	log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Level = logrus.TraceLevel

	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		return nil, err
	}

	return log, nil
}
