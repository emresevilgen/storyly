package go_log_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	suppressLogs := os.Getenv("SUPRRESS_LOGS")
	os.Setenv("SUPPRESS_LOGS", "0")
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log Test Suite")
	os.Setenv("SUPPRESS_LOGS", suppressLogs)
}
