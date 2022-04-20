package go_log

import (
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strings"
)

const lowestLogLevelPossible = logrus.DebugLevel

var logLevels = make(map[string]logrus.Level)

func SetLogLevels(logLevels map[string]string) {
	for loggerName, logLevel := range logLevels {
		SetLogLevel(loggerName, logLevel)
	}
}

func SetLogLevel(loggerName string, newLevel string) {
	newLevel = strings.ToLower(newLevel)

	if newLevel == "debug" {
		logLevels[loggerName] = logrus.DebugLevel
	} else if newLevel == "info" {
		logLevels[loggerName] = logrus.InfoLevel
	} else if newLevel == "warn" {
		logLevels[loggerName] = logrus.WarnLevel
	} else if newLevel == "error" {
		logLevels[loggerName] = logrus.ErrorLevel
	} else {
		logLevels[loggerName] = logrus.InfoLevel
	}
}

func GetLogLevel(loggerName string) logrus.Level {
	return logLevels[loggerName]
}

func NewLoggerByType(loggerType reflect.Type, defaultFields map[string]interface{}) *Logger {
	return NewLoggerByName(loggerType.String(), defaultFields)
}

func NewLoggerByName(loggerName string, defaultFields map[string]interface{}) *Logger {
	logger := new(Logger)

	logger.name = loggerName
	logger.defaultFields = getDefaultFields(defaultFields)

	logger.logrusLogger = logrus.StandardLogger()

	if suppressLogs() {
		logger.logrusLogger.Level = logrus.FatalLevel
	} else {
		logger.logrusLogger.Level = lowestLogLevelPossible
	}

	if isBuildMode() {
		logrus.SetFormatter(&PrettyJSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
	}

	addLogLevelIfMissing(loggerName, "info")

	return logger
}

func getDefaultFields(defaultFields map[string]interface{}) map[string]interface{} {
	fields := make(map[string]interface{})

	for field, value := range defaultFields {
		if !isSensitiveField(field) {
			fields[field] = value
		}
	}
	return fields
}

func mapFieldsToEntry(entry *logrus.Entry, fields map[string]interface{}) {
	for field, value := range fields {
		if !isSensitiveField(field) {
			escapedField := strings.Replace(field, ".", "_", -1)
			entry.Data[escapedField] = value
		}
	}
}

func isSensitiveField(field string) bool {
	return field == "Authorization"
}

func addLogLevelIfMissing(loggerName string, newLevel string) {
	if logLevels[loggerName] == logrus.PanicLevel {
		SetLogLevel(loggerName, newLevel)
	}
}

func isDebugEnabled(loggerName string) bool {
	return logLevels[loggerName] >= logrus.DebugLevel
}

func isInfoEnabled(loggerName string) bool {
	return logLevels[loggerName] >= logrus.InfoLevel
}

func isWarnEnabled(loggerName string) bool {
	return logLevels[loggerName] >= logrus.WarnLevel
}

func isErrorEnabled(loggerName string) bool {
	return logLevels[loggerName] >= logrus.ErrorLevel
}

func isBuildMode() bool {
	return os.Getenv("BUILD_MODE") == "1"
}

func suppressLogs() bool {
	return os.Getenv("SUPPRESS_LOGS") == "1"
}
