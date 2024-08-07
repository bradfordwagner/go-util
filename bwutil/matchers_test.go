package bwutil_test

import (
	"github.com/bradfordwagner/go-util/bwutil"
	"github.com/bradfordwagner/go-util/mocks/bwutil"

	. "github.com/onsi/ginkgo/v2"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Matchers", func() {
	var ctrl *gomock.Controller
	BeforeEach(func() { ctrl = gomock.NewController(GinkgoT()) })

	Context("MatcherOneOf", func() {
		type args struct {
			set        []int
			invokeWith []int
		}
		type res struct {
		}
		var test = func(a args, r res) {
			invoke := mock_bwutil.NewMockStubOneOf(ctrl)
			matcher := bwutil.NewMatcherOneOf(a.set)

			invoke.EXPECT().Invoke(matcher).Times(len(a.invokeWith))

			for _, v := range a.invokeWith {
				invoke.Invoke(v)
			}

			ctrl.Finish()
		}
		It("invokes with everything expected", func() {
			test(args{
				set:        []int{1, 2, 3},
				invokeWith: []int{1, 2, 3},
			}, res{})
		})
		It("invokes only some of the set", func() {
			test(args{
				set:        []int{1, 2, 3},
				invokeWith: []int{1, 3},
			}, res{})
		})
	})

	Context("MatcherConversionOneOf", func() {
		type orig struct {
			i int
		}
		type args struct {
			set        []int
			invokeWith []orig
		}
		type res struct {
		}
		var test = func(a args, r res) {
			invoke := mock_bwutil.NewMockStubConversionOneOf(ctrl)
			matcher := bwutil.NewMatcherConversionOneOf(a.set, func(o orig) int {
				return o.i
			})

			invoke.EXPECT().Invoke(matcher).Times(len(a.invokeWith))

			for _, v := range a.invokeWith {
				invoke.Invoke(v)
			}

			ctrl.Finish()
		}
		It("", func() {
			test(args{
				set: []int{1, 2},
				invokeWith: []orig{
					{i: 1},
					{i: 2},
				},
			}, res{})
		})
	})

	It("NewMatcherConversionExploderOneOf", func() {
		type value struct {
			arr []int
		}
		type args struct {
			values []value
			set    []int
		}
		var test = func(a args) {
			invoke := mock_bwutil.NewMockStubConversionOneOf(ctrl)
			matcher := bwutil.NewMatcherConversionExploderOneOf(a.set, func(v value) []int {
				return v.arr
			})

			invoke.EXPECT().Invoke(matcher).Times(len(a.values))

			for _, v := range a.values {
				invoke.Invoke(v)
			}

			ctrl.Finish()
		}

		// run tests
		test(args{
			values: []value{
				{arr: []int{1, 2}},
			},
			set: []int{1, 2},
		})
		test(args{
			values: []value{
				{arr: []int{1, 2}},
				{arr: []int{2, 1}},
			},
			set: []int{1, 2},
		})
		test(args{
			values: []value{
				{arr: []int{1, 2}},
				{arr: []int{3, 4}},
			},
			set: []int{1, 2, 3, 4},
		})

	})

})
