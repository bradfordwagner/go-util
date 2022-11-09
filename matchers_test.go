package bwutil_test

import (
	"github.com/bradfordwagner/go-util"
	mock_bwutil "github.com/bradfordwagner/go-util/mocks/bwutil"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Matchers", func() {
	var ctrl *gomock.Controller
	BeforeEach(func() { ctrl = gomock.NewController(GinkgoT()) })

	Context("OneOfMatcher", func() {
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

	Context("ConversionOneOfMatcher", func() {
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
			matcher := bwutil.NewConversionOneOfMatcher(a.set, func(o orig) int {
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
})
