package bwutil_test

import (
	"github.com/bradfordwagner/go-util/bwutil"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Math", func() {
	It("MathMax", func() {
		type args struct {
			a, b, exp int
		}
		test := func(arg args) {
			res := bwutil.MathMax(arg.a, arg.b)
			Expect(res).To(Equal(arg.exp))
		}

		// run tests
		test(args{
			a:   1,
			b:   2,
			exp: 2,
		})
		test(args{
			a:   2,
			b:   1,
			exp: 2,
		})
		test(args{
			a:   1,
			b:   1,
			exp: 1,
		})
		test(args{
			a:   -1,
			b:   1,
			exp: 1,
		})
		test(args{
			a:   -10000,
			b:   -1,
			exp: -1,
		})
	})
	It("MathMin", func() {
		type args struct {
			a, b, exp int
		}
		test := func(arg args) {
			res := bwutil.MathMin(arg.a, arg.b)
			Expect(res).To(Equal(arg.exp))
		}

		// run tests
		test(args{
			a:   1,
			b:   2,
			exp: 1,
		})
		test(args{
			a:   2,
			b:   1,
			exp: 1,
		})
		test(args{
			a:   1,
			b:   1,
			exp: 1,
		})
		test(args{
			a:   -1,
			b:   1,
			exp: -1,
		})
	})

})
