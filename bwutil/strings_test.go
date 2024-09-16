package bwutil

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strings", func() {
	It("StringsSplitAndTrim", func() {
		type args struct {
			input  string
			sep    string
			output []string
		}
		var test = func(args args) {
			Expect(StringsSplitAndTrim(args.input, args.sep)).To(Equal(args.output))
		}
		test(args{"a,b,c", ",", []string{"a", "b", "c"}})
		test(args{"a, b, c", ",", []string{"a", "b", "c"}})
		test(args{"a, b, c", " ", []string{"a,", "b,", "c"}})
		test(args{"a, b, c", ",", []string{"a", "b", "c"}})
		test(args{"   a,   b, c  ", ",", []string{"a", "b", "c"}})
	})
})
