package log_factory

import (
	log "storyly/pkg/log"
)

type Logger interface {
	Info(msg string, args ...interface{})
	InfoWithFields(fields map[string]interface{}, msg string, args ...interface{})
	Error(msg string, args ...interface{})
	ErrorWithFields(fields map[string]interface{}, msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	WarnWithFields(fields map[string]interface{}, msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	DebugWithFields(fields map[string]interface{}, msg string, args ...interface{})
}

type logger struct {
	logger *log.Logger
}

func (l logger) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

func (l logger) InfoWithFields(fields map[string]interface{}, msg string, args ...interface{}) {
	l.logger.InfoWithFields(fields, msg, args...)
}

func (l logger) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}

func (l logger) ErrorWithFields(fields map[string]interface{}, msg string, args ...interface{}) {
	l.logger.ErrorWithFields(fields, msg, args...)
}

func (l logger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

func (l logger) WarnWithFields(fields map[string]interface{}, msg string, args ...interface{}) {
	l.logger.WarnWithFields(fields, msg, args...)
}

func (l logger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}

func (l logger) DebugWithFields(fields map[string]interface{}, msg string, args ...interface{}) {
	l.logger.DebugWithFields(fields, msg, args...)
}
