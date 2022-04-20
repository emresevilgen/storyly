package go_log_test

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	log "storyly/pkg/log"
	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formatters tests:", func() {

	Describe("when BUILD_MODE is 1", func() {
		var (
			logger              *log.Logger
			prettyJsonFormatter = reflect.TypeOf(&log.PrettyJSONFormatter{})
			actual              reflect.Type
		)

		BeforeAll(func() {
			os.Setenv("BUILD_MODE", "1")
			logger = log.NewLoggerByName("PrettyLogger", map[string]interface{}{"defaultKey": "defaultValue"})

			actual = reflect.TypeOf(logger.GetFormatter())

			os.Unsetenv("BUILD_MODE")
		})

		It("should use PrettyJSONFormatter", func() {
			Expect(actual).To(Equal(prettyJsonFormatter))
		})
	})

	Describe("when BUILD_MODE is not provided", func() {
		var (
			logger        *log.Logger
			jsonFormatter = reflect.TypeOf(&logrus.JSONFormatter{})
			actual        reflect.Type
		)

		BeforeAll(func() {
			os.Unsetenv("BUILD_MODE")
			logger = log.NewLoggerByName("Logger", map[string]interface{}{"defaultKey": "defaultValue"})

			actual = reflect.TypeOf(logger.GetFormatter())
		})

		It("should use default logrus JSON formatter", func() {
			Expect(actual).To(Equal(jsonFormatter))
		})
	})

	Describe("when using PrettyJSONFormatter", func() {
		var (
			buffer = new(bytes.Buffer)
			fields logrus.Fields
			err    error
		)

		BeforeAll(func() {
			logger := log.NewLoggerByName("PrettyLogger", map[string]interface{}{"defaultKey": "defaultValue"})
			logger.SetFormatter(&log.PrettyJSONFormatter{})
			logger.SetOut(buffer)

			logger.Error("A log message")

			err = json.Unmarshal(buffer.Bytes(), &fields)
		})

		It("should log a message", func() {
			Expect(len(buffer.Bytes())).To(BeNumerically(">", 0))
		})

		It("should format the log in JSON format", func() {
			Expect(err).To(BeNil())
		})

		It("should log the correct message", func() {
			Expect(fields["msg"]).To(Equal("A log message"))
		})

		It("should log the default field", func() {
			Expect(fields["defaultKey"]).To(Equal("defaultValue"))
		})
	})
})

var BeforeAll = func(beforeAllFunc func()) {
	var once sync.Once

	BeforeEach(func() {
		once.Do(func() {
			beforeAllFunc()
		})
	})
}
