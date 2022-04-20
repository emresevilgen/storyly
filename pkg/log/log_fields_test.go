package go_log_test

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	log "storyly/pkg/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log fields tests:", func() {
	Describe("when logging a message with default fields", func() {
		var (
			buffer = new(bytes.Buffer)
			fields logrus.Fields
			err    error
		)

		BeforeAll(func() {
			logger := log.NewLoggerByName("testlogger", map[string]interface{}{"defaultKey": "defaultValue"})
			logger.SetOut(buffer)

			logger.Info("A log message")

			err = json.Unmarshal(buffer.Bytes(), &fields)
		})

		It("should log a message", func() {
			Expect(len(buffer.Bytes())).To(BeNumerically(">", 0))
		})

		It("should format the log in JSON format", func() {
			Expect(err).To(BeNil())
		})

		It("should log the default field", func() {
			Expect(fields["defaultKey"]).To(Equal("defaultValue"))
		})
	})

	Describe("when logging a message with \"sensitive\" default field", func() {
		var (
			buffer = new(bytes.Buffer)
			fields logrus.Fields
			err    error
		)

		BeforeAll(func() {
			logger := log.NewLoggerByName("testlogger", map[string]interface{}{"Authorization": "XXX"})
			logger.SetOut(buffer)

			logger.Info("A log message")

			err = json.Unmarshal(buffer.Bytes(), &fields)
		})

		It("should log a message", func() {
			Expect(len(buffer.Bytes())).To(BeNumerically(">", 0))
		})

		It("should format the log in JSON format", func() {
			Expect(err).To(BeNil())
		})

		It("should not log the \"sensitive\" default field ", func() {
			Expect(fields).NotTo(HaveKey("Authorization"))
		})
	})

	Describe("when logging a message with additional fields", func() {
		var (
			buffer = new(bytes.Buffer)
			fields logrus.Fields
			err    error
		)

		BeforeAll(func() {
			logger := log.NewLoggerByName("testlogger", map[string]interface{}{"defaultKey": "defaultValue"})
			logger.SetOut(buffer)

			logger.InfoWithFields(map[string]interface{}{"additionalKey": "value", "Authorization": "XXX"}, "A log message")

			err = json.Unmarshal(buffer.Bytes(), &fields)
		})

		It("should log a message", func() {
			Expect(len(buffer.Bytes())).To(BeNumerically(">", 0))
		})

		It("should format the log in JSON format", func() {
			Expect(err).To(BeNil())
		})

		It("should log the default field", func() {
			Expect(fields["defaultKey"]).To(Equal("defaultValue"))
		})

		It("should log the additional field", func() {
			Expect(fields["additionalKey"]).To(Equal("value"))
		})
	})

	Describe("when logging a message with \"sensitive\" additional field", func() {
		var (
			buffer = new(bytes.Buffer)
			fields logrus.Fields
			err    error
		)

		BeforeAll(func() {
			logger := log.NewLoggerByName("testlogger", map[string]interface{}{"defaultKey": "defaultValue"})
			logger.SetOut(buffer)

			logger.InfoWithFields(map[string]interface{}{"Authorization": "XXX"}, "A log message")

			err = json.Unmarshal(buffer.Bytes(), &fields)
		})

		It("should log a message", func() {
			Expect(len(buffer.Bytes())).To(BeNumerically(">", 0))
		})

		It("should format the log in JSON format", func() {
			Expect(err).To(BeNil())
		})

		It("should log the default field", func() {
			Expect(fields["defaultKey"]).To(Equal("defaultValue"))
		})

		It("should not log the the \"sensitive\" additional Authorization field", func() {
			Expect(fields).NotTo(HaveKey("Authorization"))
		})
	})
})
