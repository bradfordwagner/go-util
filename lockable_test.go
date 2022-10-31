package bwutil_test

import (
	"github.com/bradfordwagner/go-util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lockable", func() {
	It("creates a lockable", func() {
		l := bwutil.NewLockable("abcd")
		Expect(l).ToNot(BeNil())
		Expect(l.Get()).To(Equal("abcd"))
	})
	It("sets the value", func() {
		l := bwutil.NewLockable("abcd")
		Expect(l).ToNot(BeNil())
		Expect(l.Get()).To(Equal("abcd"))

		l.Set("hi")
		Expect(l.Get()).To(Equal("hi"))
	})
})
