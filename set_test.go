package bwutil_test

import (
	bwutil "github.com/bradfordwagner/go-util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	It("adds and removes from the set", func() {
		s := bwutil.NewSet[string]()
		// test add
		Expect(s.Exists("a")).To(BeFalse())
		s.Add("a")
		Expect(s.Exists("a")).To(BeTrue())

		// test remove
		s.Remove("a")
		Expect(s.Exists("a")).To(BeFalse())
	})
})
