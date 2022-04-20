package go_log_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	log "storyly/pkg/log"
	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log level tests:", func() {
	Describe("When log levels are set for multiple loggers externally", func() {
		var (
			logger1, logger2, logger3 *log.Logger
			once                      sync.Once
		)
		BeforeEach(func() {
			once.Do(func() {
				log.SetLogLevels(map[string]string{"logger1": "debug", "logger2": "warn"})

				logger1 = log.NewLoggerByName("logger1", nil)
				logger2 = log.NewLoggerByName("logger2", nil)
				logger3 = log.NewLoggerByName("logger3", nil)
			})
		})

		It("should use the correct log levels for all the set loggers", func() {
			Expect(log.GetLogLevel(logger1.GetName())).To(Equal(logrus.DebugLevel))
			Expect(log.GetLogLevel(logger2.GetName())).To(Equal(logrus.WarnLevel))
		})

		It("should use the default log level for all the unset loggers", func() {
			Expect(log.GetLogLevel(logger3.GetName())).To(Equal(logrus.InfoLevel))
		})
	})

	Describe("When log level is set to Debug", func() {

		Describe("And logger is called to log in Debug level without any additional fields", func() {

			It("should log fields", func() {
				logAndAssert("debug",
					nil,
					func(l *log.Logger) {
						l.Debug("post not found")
					},
					func(fields logrus.Fields) {

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("debug"))
					})
			})

			It("should log with default fields ", func() {

				logAndAssertWithDefaultFields("debug",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.Debug("post not found")
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("debug"))
					})
			})

			It("should log formatted string ", func() {
				postId := 103

				logAndAssertWithDefaultFields("debug",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.Debug("post not found with postId: %d", postId)
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["msg"]).To(Equal("post not found with postId: 103"))
						Expect(fields["level"]).To(Equal("debug"))
					})
			})
		})

		Describe("And logger is called to log in Debug level with additional fields", func() {

			It("should log both default fields and given fields", func() {

				logAndAssertWithDefaultFields("debug",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.DebugWithFields(map[string]interface{}{"calling.method": "TestFunction"}, "post not found")
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["calling_method"]).To(Equal("TestFunction"))

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("debug"))
					})
			})

			It("should log formatted string with given fields", func() {
				postId := 103

				logAndAssert("debug",
					nil,
					func(l *log.Logger) {
						l.DebugWithFields(map[string]interface{}{"calling.method": "TestFunction"}, "post not found with postId: %d", postId)
					},
					func(fields logrus.Fields) {
						Expect(fields["calling_method"]).To(Equal("TestFunction"))
						Expect(fields["msg"]).To(Equal("post not found with postId: 103"))
						Expect(fields["level"]).To(Equal("debug"))
					})
			})
		})
	})

	Describe("When log level is set to Info", func() {

		Describe("And logger is called to log in Info level without any additional fields", func() {
			It("should log fields", func() {

				logAndAssert("info", nil, func(l *log.Logger) {
					l.Info("post not found")
				}, func(fields logrus.Fields) {
					Expect(fields["msg"]).To(Equal("post not found"))
					Expect(fields["level"]).To(Equal("info"))
				})

			})

			It("should log with default fields", func() {

				logAndAssertWithDefaultFields("info",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)}, nil,
					func(l *log.Logger) {
						l.Info("post not found")
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("info"))
					})

			})

			It("should log formatted string ", func() {
				postId := 103

				logAndAssertWithDefaultFields("info",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.Info("post not found with postId: %d", postId)
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["msg"]).To(Equal("post not found with postId: 103"))
						Expect(fields["level"]).To(Equal("info"))
					})
			})

			It("should filter out debug logs", func() {

				logAndAssertNoLog("info", nil, func(l *log.Logger) {
					l.Debug("post not found")
				})
			})
		})

		Describe("And logger is called to log in Info level with additional fields", func() {
			It("should log both default fields and given fields", func() {

				logAndAssertWithDefaultFields("info",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.InfoWithFields(map[string]interface{}{"calling.method": "TestFunction"}, "post not found")
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["calling_method"]).To(Equal("TestFunction"))

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("info"))
					})
			})

			It("should log formatted string with given fields", func() {
				postId := 103

				logAndAssert("info",
					nil,
					func(l *log.Logger) {
						l.InfoWithFields(map[string]interface{}{"calling.method": "TestFunction"}, "post not found with postId: %d", postId)
					},
					func(fields logrus.Fields) {
						Expect(fields["calling_method"]).To(Equal("TestFunction"))
						Expect(fields["msg"]).To(Equal("post not found with postId: 103"))
						Expect(fields["level"]).To(Equal("info"))
					})
			})
		})
	})

	Describe("When log level is set to Warn", func() {

		Describe("And logger is called to log in Warn level without any additional fields", func() {
			It("should log fields", func() {

				logAndAssert("warn", nil, func(l *log.Logger) {
					l.Warn("post not found")
				}, func(fields logrus.Fields) {
					Expect(fields["msg"]).To(Equal("post not found"))
					Expect(fields["level"]).To(Equal("warning"))
				})

			})

			It("should log with default fields", func() {

				logAndAssertWithDefaultFields("warn",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.Warn("post not found")
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("warning"))
					})

			})

			It("should log formatted string ", func() {
				postId := 103

				logAndAssertWithDefaultFields("warn",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.Warn("post not found with postId: %d", postId)
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["msg"]).To(Equal("post not found with postId: 103"))
						Expect(fields["level"]).To(Equal("warning"))
					})
			})

			It("should filter out debug logs", func() {
				logAndAssertNoLog("warn", nil, func(l *log.Logger) {
					l.Debug("post not found")
				})

			})

			It("should filter out info logs", func() {
				logAndAssertNoLog("warn", nil, func(l *log.Logger) {
					l.Info("post not found")
				})
			})
		})

		Describe("And logger is called to log in Warn level with additional fields", func() {
			It("should log both default fields and given fields", func() {

				logAndAssertWithDefaultFields("debug",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.WarnWithFields(map[string]interface{}{"calling.method": "TestFunction"}, "post not found")
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["calling_method"]).To(Equal("TestFunction"))

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("warning"))
					})
			})

			It("should log formatted string with given fields", func() {
				postId := 103

				logAndAssert("debug",
					nil,
					func(l *log.Logger) {
						l.WarnWithFields(map[string]interface{}{"calling.method": "TestFunction"}, "post not found with postId: %d", postId)
					},
					func(fields logrus.Fields) {
						Expect(fields["calling_method"]).To(Equal("TestFunction"))
						Expect(fields["msg"]).To(Equal("post not found with postId: 103"))
						Expect(fields["level"]).To(Equal("warning"))
					})
			})
		})

	})

	Describe("When log level is set to Error", func() {

		Describe("And logger is called to log in Error level without any additional fields", func() {
			It("should log fields", func() {

				logAndAssert("error", nil, func(l *log.Logger) {
					l.Error("post not found")
				}, func(fields logrus.Fields) {
					Expect(fields["msg"]).To(Equal("post not found"))
					Expect(fields["level"]).To(Equal("error"))
				})

			})

			It("should log with default fields", func() {

				logAndAssertWithDefaultFields("error",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.Error("post not found")
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("error"))

					})

			})

			It("should log formatted string ", func() {
				postId := 103

				logAndAssertWithDefaultFields("debug",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.Error("post not found with postId: %d", postId)
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["msg"]).To(Equal("post not found with postId: 103"))
						Expect(fields["level"]).To(Equal("error"))
					})
			})

			It("should filter out debug logs", func() {
				logAndAssertNoLog("error", nil, func(l *log.Logger) {
					l.Debug("post not found")
				})
			})

			It("should filter out info logs", func() {
				logAndAssertNoLog("error", nil, func(l *log.Logger) {
					l.Info("post not found")
				})
			})

			It("should filter out warn logs", func() {
				logAndAssertNoLog("error", nil, func(l *log.Logger) {
					l.Warn("post not found")
				})
			})
		})

		Describe("And logger is called to log in Error level with additional fields", func() {
			It("should log both default fields and given fields", func() {

				logAndAssertWithDefaultFields("debug",
					map[string]interface{}{"defaultKey1": "defaultValue", "defaultKey2": float64(789)},
					nil,
					func(l *log.Logger) {
						l.ErrorWithFields(map[string]interface{}{"calling.method": "TestFunction"}, "post not found")
					},
					func(fields logrus.Fields) {
						Expect(fields["defaultKey1"]).To(Equal("defaultValue"))
						Expect(fields["defaultKey2"]).To(Equal(float64(789)))

						Expect(fields["calling_method"]).To(Equal("TestFunction"))

						Expect(fields["msg"]).To(Equal("post not found"))
						Expect(fields["level"]).To(Equal("error"))
					})
			})

			It("should log formatted string with given fields", func() {
				postId := 103

				logAndAssert("debug",
					nil,
					func(l *log.Logger) {
						l.ErrorWithFields(map[string]interface{}{"calling.method": "TestFunction"}, "post not found with postId: %d", postId)
					},
					func(fields logrus.Fields) {
						Expect(fields["calling_method"]).To(Equal("TestFunction"))
						Expect(fields["msg"]).To(Equal("post not found with postId: 103"))
						Expect(fields["level"]).To(Equal("error"))
					})
			})
		})
	})
})

func logAndAssert(loggingLevel string, createLogger func() *log.Logger, loggingAction func(*log.Logger), assert func(fields logrus.Fields)) {
	buffer := new(bytes.Buffer)

	var fields logrus.Fields

	var l *log.Logger

	if createLogger == nil {
		l = log.NewLoggerByName("NewLogger", nil)
	} else {
		l = createLogger()
	}

	log.SetLogLevel(l.GetName(), loggingLevel)

	l.SetFormatter(new(logrus.JSONFormatter))
	l.SetOut(buffer)

	loggingAction(l)

	if buffer.Bytes() != nil {
		err := json.Unmarshal(buffer.Bytes(), &fields)

		if err != nil {
			fmt.Println(err)
		}

		Expect(err).ToNot(HaveOccurred())
	}

	assert(fields)
}

func logAndAssertWithDefaultFields(loggingLevel string, defaultFields map[string]interface{}, createLogger func(map[string]interface{}) *log.Logger, loggingAction func(*log.Logger), assert func(fields logrus.Fields)) {
	buffer := new(bytes.Buffer)

	var fields logrus.Fields

	var l *log.Logger

	if createLogger == nil {
		l = log.NewLoggerByName("NewLoggerWithDefaultFields", defaultFields)
	} else {
		l = createLogger(defaultFields)
	}

	log.SetLogLevel(l.GetName(), loggingLevel)

	l.SetFormatter(new(logrus.JSONFormatter))
	l.SetOut(buffer)

	loggingAction(l)

	err := json.Unmarshal(buffer.Bytes(), &fields)
	if err != nil {
		fmt.Println(err)
	}

	Expect(err).ToNot(HaveOccurred())

	assert(fields)
}

func logAndAssertNoLog(loggingLevel string, createLogger func() *log.Logger, loggingAction func(*log.Logger)) {
	buffer := new(bytes.Buffer)

	var l *log.Logger

	if createLogger == nil {
		l = log.NewLoggerByName("NewLogger", nil)
	} else {
		l = createLogger()
	}

	log.SetLogLevel(l.GetName(), loggingLevel)

	l.SetFormatter(new(logrus.JSONFormatter))
	l.SetOut(buffer)

	loggingAction(l)

	Expect(buffer.Bytes()).To(HaveLen(0))
}
