package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/marcio-garcia/go-microservices/api/config"
	"github.com/sirupsen/logrus"
)

// Log -
var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}

	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
	}
}

// Debug logs the text in the parameter for the debug log level
func Debug(message string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(message)
}

// Info logs the text in the parameter for the info log level
func Info(message string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(message)
}

// Error logs the text in the parameter for the debug log level
func Error(message string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	msg := fmt.Sprintf("%s - ERROR - %v", message, err)
	Log.WithFields(parseFields(tags...)).Error(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))

	for _, tag := range tags {
		elements := strings.Split(tag, ":")
		key := strings.TrimSpace(elements[0])
		value := strings.TrimSpace(elements[1])
		result[key] = value
	}

	return result
}
