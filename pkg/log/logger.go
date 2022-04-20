package go_log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"path"
	"runtime"
)

type Logger struct {
	name          string
	logrusLogger  *logrus.Logger
	defaultFields map[string]interface{}
}

func (l *Logger) GetName() string {
	return l.name
}

func (l *Logger) SetOut(o io.Writer) {
	l.logrusLogger.Out = o
}

func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	l.logrusLogger.Formatter = formatter
}

func (l *Logger) GetFormatter() logrus.Formatter {
	return l.logrusLogger.Formatter
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if !isDebugEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Debug(format)
}

func (l *Logger) DebugWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if !isDebugEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Debug(format)
}

func (l *Logger) Info(format string, args ...interface{}) {
	if !isInfoEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Info(format)
}

func (l *Logger) InfoWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if !isInfoEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Info(format)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	if !isWarnEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Warn(format)
}

func (l *Logger) WarnWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if !isWarnEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}
	mapFieldsToEntry(entry, fields)
	entry.Warn(format)
}

func (l *Logger) Error(format string, args ...interface{}) {
	if !isErrorEnabled(l.name) {
		return
	}
	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	entry.Error(format)
}

func (l *Logger) ErrorWithFields(fields map[string]interface{}, format string, args ...interface{}) {
	if !isErrorEnabled(l.name) {
		return
	}

	entry := l.newLogEntry()

	if args != nil {
		format = fmt.Sprintf(format, args...)
	}

	mapFieldsToEntry(entry, fields)

	entry.Error(format)
}

func (l *Logger) IsDebugEnabled() bool {
	return isDebugEnabled(l.name)
}

func (l *Logger) IsInfoEnabled() bool {
	return isInfoEnabled(l.name)
}

func (l *Logger) IsWarnEnabled() bool {
	return isWarnEnabled(l.name)
}

func (l *Logger) IsErrorEnabled() bool {
	return isErrorEnabled(l.name)
}

func (l *Logger) newLogEntry() *logrus.Entry {
	defaultFields := logrus.Fields(l.defaultFields)

	entry := l.logrusLogger.WithFields(defaultFields)

	entry.Data["logger"] = l.name
	entry.Data["callStack"] = l.getCallStack()

	return entry
}

func (l *Logger) getCallStack() map[string]interface{} {
	pc, file, line, _ := runtime.Caller(3)
	_, fileName := path.Split(file)
	fullPath := runtime.FuncForPC(pc).Name()

	return map[string]interface{}{
		"fullPath": fullPath,
		"fileName": fileName,
		"line":     line,
	}
}
