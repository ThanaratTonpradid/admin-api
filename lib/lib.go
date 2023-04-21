package lib

import (
	defaultLog "log"
	"os"
	"strings"

	"github.com/dollarsignteam/go-logger"

	"mini-api/helper"
)

var log *logger.Logger

func init() {
	log = logger.NewLogger(logger.LoggerOptions{
		Name:       helper.GetPackageName(),
		HideCaller: true,
	})
}

type LogWriter struct{}

func (writer LogWriter) Write(bytes []byte) (int, error) {
	log.Info(strings.TrimSuffix(string(bytes), "\n"))
	return len(bytes), nil
}

func NewLibLogger() *defaultLog.Logger {
	logger := defaultLog.New(os.Stdout, "", 0)
	logger.SetOutput(new(LogWriter))
	return logger
}
