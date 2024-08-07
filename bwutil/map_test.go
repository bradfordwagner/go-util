package bwutil_test

import (
	"reflect"

	bwutil "github.com/bradfordwagner/go-util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map", func() {
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

	Context("ToSortedSliceByKey", func() {
		It("sorts by key", func() {
			m := map[int]int{
				5: 5,
				4: 4,
				3: 3,
				2: 2,
				1: 1,
				0: 0,
			}
			s := bwutil.ToSortedSliceByKey(m)
			Expect(s[0]).To(Equal(0))
			Expect(s[1]).To(Equal(1))
			Expect(s[2]).To(Equal(2))
			Expect(s[3]).To(Equal(3))
			Expect(s[4]).To(Equal(4))
			Expect(s[5]).To(Equal(5))
		})
	})
})
