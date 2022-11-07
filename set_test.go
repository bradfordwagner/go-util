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
		Expect(s.Keyset()).To(BeEquivalentTo([]string{}))

		// test add
		Expect(s.Exists("a")).To(BeFalse())
		s.Add("a")
		Expect(s.Exists("a")).To(BeTrue())
		Expect(s.Size()).To(Equal(1))
		Expect(s.IsEmpty()).To(BeFalse())
		Expect(s.Keyset()).To(BeEquivalentTo([]string{"a"}))

		// second add
		Expect(s.Exists("b")).To(BeFalse())
		s.Add("b")
		Expect(s.Exists("b")).To(BeTrue())
		Expect(s.Size()).To(Equal(2))
		Expect(s.IsEmpty()).To(BeFalse())
		Expect(s.Keyset()).To(BeEquivalentTo([]string{"a", "b"}))

		// test remove of a, leaving b
		s.Remove("a")
		Expect(s.Exists("a")).To(BeFalse())
		Expect(s.IsEmpty()).To(BeFalse())
		Expect(s.Keyset()).To(BeEquivalentTo([]string{"b"}))
	})

	It("creates a set from slice", func() {
		s := []int{1, 2, 3, 4}
		set := bwutil.NewSetFromSlice(s)
		Expect(set.Size()).To(Equal(len(s)))

		s = []int{1, 1, 2, 3}
		set = bwutil.NewSetFromSlice(s)
		Expect(set.Size()).To(Equal(3))
	})
})
