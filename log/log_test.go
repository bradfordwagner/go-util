package log

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Log", func() {
	It("sets up a logger with no errors", func() {
		// outputs: 2024-03-26T13:04:23.927-0400 info  log/log_test.go:9 This is a test log message
		Log().Info("This is a test log message")
	})
})
