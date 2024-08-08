package bwutil_test

import (
	"github.com/bradfordwagner/go-util/bwutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	It("adds and removes from the set", func() {
		s := bwutil.NewSet[string]()

		// initially no items
		Expect(s.IsEmpty()).To(BeTrue())
		Expect(s.Size()).To(BeZero())
		Expect(s.Keyset()).To(ContainElements([]string{}))

		// test add
		Expect(s.Exists("a")).To(BeFalse())
		s.Add("a")
		Expect(s.Exists("a")).To(BeTrue())
		Expect(s.Size()).To(Equal(1))
		Expect(s.IsEmpty()).To(BeFalse())
		Expect(s.Keyset()).To(ContainElements([]string{"a"}))

		// second add
		Expect(s.Exists("b")).To(BeFalse())
		s.Add("b")
		Expect(s.Exists("b")).To(BeTrue())
		Expect(s.Size()).To(Equal(2))
		Expect(s.IsEmpty()).To(BeFalse())
		Expect(s.Keyset()).To(ContainElements([]string{"a", "b"}))

		// test remove of a, leaving b
		s.Remove("a")
		Expect(s.Exists("a")).To(BeFalse())
		Expect(s.IsEmpty()).To(BeFalse())
		Expect(s.Keyset()).To(ContainElements([]string{"b"}))
	})

	It("creates a set from slice", func() {
		s := []int{1, 2, 3, 4}
		set := bwutil.NewSetFromSlice(s)
		Expect(set.Size()).To(Equal(len(s)))

		s = []int{1, 1, 2, 3}
		set = bwutil.NewSetFromSlice(s)
		Expect(set.Size()).To(Equal(3))
	})

	It("difference", func() {
		type args struct {
			a, b, res []int
		}
		test := func(arg args) {
			a, b, expected := bwutil.NewSetFromSlice(arg.a), bwutil.NewSetFromSlice(arg.b), bwutil.NewSetFromSlice(arg.res)
			res := a.Difference(b)
			Expect(res.Keyset()).To(ConsistOf(expected.Keyset()))
		}

		// run tests
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{2, 3},
			res: []int{1},
		})
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{1, 2},
			res: []int{3},
		})
		test(args{
			a:   []int{},
			b:   []int{1, 2},
			res: []int{},
		})
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{},
			res: []int{1, 2, 3},
		})
	})

	It("equals", func() {
		type args struct {
			a, b []int
			res  bool
		}
		test := func(arg args) {
			a, b := bwutil.NewSetFromSlice(arg.a), bwutil.NewSetFromSlice(arg.b)
			eq := a.Equals(b)
			Expect(eq).To(Equal(arg.res))
		}

		// run tests
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{2},
			res: false,
		})
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{2, 3, 4},
			res: false,
		})
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{1, 2, 3},
			res: true,
		})
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{1, 2, 3, 4},
			res: false,
		})
	})

	It("remove all", func() {
		type args struct {
			a, b, res []int
		}
		test := func(arg args) {
			a := bwutil.NewSetFromSlice(arg.a)
			a.RemoveAll(arg.b)
			Expect(a.Keyset()).To(ConsistOf(arg.res))
		}

		// run tests
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{2, 3},
			res: []int{1},
		})
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{},
			res: []int{1, 2, 3},
		})
		test(args{
			a:   []int{1, 2, 3},
			b:   []int{4, 5},
			res: []int{1, 2, 3},
		})
		test(args{
			a:   []int{},
			b:   []int{4, 5},
			res: []int{},
		})
	})

	It("exists all", func() {
		type args struct {
			values []int
			set    []int
		}
		type res struct {
			existsAll bool
		}
		var test = func(a args, r res) {
			set := bwutil.NewSetFromSlice(a.set)
			res := set.ExistsAll(a.values)
			Expect(res).To(Equal(r.existsAll))
		}

		// run tests - truthy
		test(args{
			values: []int{1, 2, 3},
			set:    []int{1, 2, 3},
		}, res{existsAll: true})
		test(args{
			values: []int{1, 2, 2, 3},
			set:    []int{1, 2, 3},
		}, res{existsAll: true})
		test(args{
			values: []int{1},
			set:    []int{1, 2, 3},
		}, res{existsAll: true})
		test(args{
			values: []int{-1, 1, 2, 3},
			set:    []int{-1, 1, 2, 3},
		}, res{existsAll: true})
		test(args{
			values: []int{},
			set:    []int{-1, 1, 2, 3},
		}, res{existsAll: true})
		test(args{
			values: []int{},
			set:    []int{},
		}, res{existsAll: true})

		// run tests - false
		test(args{
			values: []int{4},
			set:    []int{1, 2, 3},
		}, res{existsAll: false})
		test(args{
			values: []int{1},
			set:    []int{},
		}, res{existsAll: false})
	})

	It("add all", func() {
		// testbed
		type args struct {
			initial, toAdd, res []int
		}
		test := func(arg args) {
			a := bwutil.NewSetFromSlice(arg.initial)
			a.AddAll(arg.toAdd)
			Expect(a.Keyset()).To(ConsistOf(arg.res))
		}

		// assert
		test(args{
			initial: []int{1, 2, 3},
			toAdd:   []int{4, 5, 6},
			res:     []int{1, 2, 3, 4, 5, 6},
		})
		test(args{
			initial: []int{1, 2, 3},
			toAdd:   []int{1, 2, 3},
			res:     []int{1, 2, 3},
		})
		test(args{
			initial: []int{},
			toAdd:   []int{},
			res:     []int{},
		})
		test(args{
			initial: []int{},
			toAdd:   []int{1, 2, 3},
			res:     []int{1, 2, 3},
		})
	})
})
