package bwutil_test

import (
	bwutil "github.com/bradfordwagner/go-util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	It("adds and removes from the set", func() {
		s := bwutil.NewSet[string]()

		// initially no items
		Expect(s.IsEmpty()).To(BeTrue())
		Expect(s.Size()).To(BeZero())

		// test add
		Expect(s.Exists("a")).To(BeFalse())
		s.Add("a")
		Expect(s.Exists("a")).To(BeTrue())
		Expect(s.Size()).To(Equal(1))
		Expect(s.IsEmpty()).To(BeFalse())

		// test remove
		s.Remove("a")
		Expect(s.Exists("a")).To(BeFalse())
		Expect(s.IsEmpty()).To(BeTrue())
	})
})
