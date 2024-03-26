package flag_helper

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFlagHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FlagHelper Suite")
}
