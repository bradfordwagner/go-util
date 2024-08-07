package bwutil_test

import (
	"github.com/bradfordwagner/go-util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Now", func() {
	It("is around now", func() {
		n := bwutil.NewUnixNow()
		t := n.Now()
		Expect(t).To(BeTemporally("~", time.Now(), time.Millisecond*25))
	})
})
