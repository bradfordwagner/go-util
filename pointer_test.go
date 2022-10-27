package util_test

import (
	"github.com/bradfordwagner/go-util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pointer", func() {
	It("strings", func() {
		v := "hi friends"
		p := util.Pointer(v)
		Expect(*p).To(Equal(v))
	})
	It("ints", func() {
		v := 1234
		p := util.Pointer(v)
		Expect(*p).To(Equal(v))
	})
})
