package bwutil_test

import (
	"errors"

	"github.com/bradfordwagner/go-util/bwutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lockable", func() {
	It("creates a lockable", func() {
		l := bwutil.NewLockableWithValue("abcd")
		Expect(l).ToNot(BeNil())
		Expect(l.Get()).To(Equal("abcd"))
	})
	It("sets the value", func() {
		l := bwutil.NewLockableWithValue("abcd")
		Expect(l).ToNot(BeNil())
		Expect(l.Get()).To(Equal("abcd"))

		l.Set("hi")
		Expect(l.Get()).To(Equal("hi"))
	})

	It("creates a lockable without a value set (string)", func() {
		l := bwutil.NewLockable[string]()
		Expect(l).ToNot(BeNil())
		Expect(l.Get()).To(Equal(""))
	})

	type mystruct struct {
		s string
	}
	It("creates a lockable without a value set (struct)", func() {
		l := bwutil.NewLockable[mystruct]()
		Expect(l).ToNot(BeNil())
		Expect(l.Get()).To(Equal(mystruct{}))

		v := mystruct{s: "abcd"}
		l.Set(v)
		Expect(l.Get()).To(Equal(v))
	})
	It("creates a lockable without a value set (*struct)", func() {
		l := bwutil.NewLockable[*mystruct]()
		Expect(l).ToNot(BeNil())
		Expect(l.Get()).To(BeNil())
	})

	Context("SetF", func() {
		It("sets using function, without error", func() {
			l := bwutil.NewLockable[string]()
			Expect(l.Get()).To(Equal(""))
			err := l.SetF(func(v string) (string, error) {
				return "hi friends", nil
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(l.Get()).To(Equal("hi friends"))
		})
		It("sets using function, and errors", func() {
			l := bwutil.NewLockable[string]()
			Expect(l.Get()).To(Equal(""))
			err := l.SetF(func(v string) (string, error) {
				return "no val", errors.New("expected error")
			})
			Expect(err.Error()).Should(Equal("expected error"))
			Expect(l.Get()).To(Equal(""))
		})
	})
})
