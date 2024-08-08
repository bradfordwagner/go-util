package bwutil_test

import (
	"github.com/bradfordwagner/go-util/bwutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pointer", func() {
	It("strings", func() {
		v := "hi friends"
		p := bwutil.Pointer(v)
		Expect(*p).To(Equal(v))
	})
	It("ints", func() {
		v := 1234
		p := bwutil.Pointer(v)
		Expect(*p).To(Equal(v))
	})
})
