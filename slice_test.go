package bwutil_test

import (
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bwutil "github.com/bradfordwagner/go-util"
)

var _ = Describe("Slice", func() {
	It("maintains order", func() {
		a, expected := []string{"1", "2", "3"}, []int{1, 2, 3}
		res := bwutil.SliceMap(a, func(a string) (b int) {
			b, _ = strconv.Atoi(a)
			return
		})
		Expect(res).To(Equal(expected))
	})

	It("takes a struct and returns a map", func() {
		type testStruct struct {
			A string
			B int
		}
		s := []testStruct{{"1", 1}, {"2", 2}, {"3", 3}}
		expected := map[string]int{"1": 1, "2": 2, "3": 3}
		res := bwutil.SliceToMap[testStruct, string, int](s, func(a testStruct) (k string, v int) {
			return a.A, a.B
		})

		Expect(res).To(Equal(expected))
	})
})
