package bwutil_test

import (
	"github.com/bradfordwagner/go-util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("Map", func() {
	It("test", func() { Expect(true).To(BeTrue()) })
	Context("MapCopy", func() {
		It("passes deep equals", func() {
			orig := map[int]int{
				1: 1,
				2: 2,
			}
			cp := bwutil.MapCopy(orig)
			Expect(reflect.DeepEqual(orig, cp)).To(BeTrue())

			// altering the cp doesn't affect the orig
			cp[3] = 3
			Expect(reflect.DeepEqual(orig, cp)).To(BeFalse())
			_, exists := orig[3]
			Expect(exists).To(BeFalse())
		})
	})
})
