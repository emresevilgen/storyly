package go_log_test

import (
	"reflect"
	log "storyly/pkg/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type dummyLogger struct{}

var _ = Describe("Logger tests:", func() {

	Describe("when getting loggers by type", func() {

		It("should name the logger with package name and type name", func() {
			logger := log.NewLoggerByType(reflect.TypeOf(dummyLogger{}), nil)

			Expect(logger.GetName()).To(Equal("go_log_test.dummyLogger"))
		})
	})

	Describe("when getting loggers by same name twice", func() {
		var (
			loggerByName     = "debugLoggerByName"
			logger1, logger2 *log.Logger
		)

		BeforeAll(func() {
			logger1 = log.NewLoggerByName(loggerByName, nil)
			logger2 = log.NewLoggerByName(loggerByName, nil)
		})

		It("should return a New instance for each call", func() {
			Expect(&logger1).NotTo(BeIdenticalTo(&logger2))
		})

		It("should have the same logger name for both loggers", func() {
			Expect(logger1.GetName()).To(Equal(logger2.GetName()))
		})
	})

	Describe("when getting loggers by same type", func() {
		var logger1, logger2 *log.Logger

		BeforeAll(func() {
			logger1 = log.NewLoggerByType(reflect.TypeOf(dummyLogger{}), nil)
			logger2 = log.NewLoggerByType(reflect.TypeOf(dummyLogger{}), nil)
		})

		It("should return a New instance for each call", func() {
			Expect(&logger1).NotTo(BeIdenticalTo(&logger2))
		})

		It("should have the same logger name for both loggers", func() {
			Expect(logger1.GetName()).To(Equal(logger2.GetName()))
		})
	})
})
