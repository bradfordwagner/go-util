package bwutil_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strconv"

	"github.com/bradfordwagner/go-util"
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
})
