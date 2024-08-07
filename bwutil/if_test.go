package bwutil_test

import (
	"github.com/bradfordwagner/go-util/bwutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("If", func() {
	It("true, returns first value", func() {
		a, b := 1, 2
		res := bwutil.If(true, a, b)
		Expect(res).To(Equal(a))
	})
	It("false, returns second value", func() {
		a, b := 1, 2
		res := bwutil.If(false, a, b)
		Expect(res).To(Equal(b))
	})
})
