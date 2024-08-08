package bwutil

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sync"
	"time"
)

var _ = Describe("TickerProvider", func() {
	var provider = NewTickerProvider()

	It("should bootstrap", func() { Expect(provider).NotTo(BeNil()) })

	It("waits one tick for it to be completed", func() {
		wg := new(sync.WaitGroup)
		wg.Add(1)
		provider.Create(time.Millisecond*20, func(t Ticker) {
			for {
				select {
				case <-t.Chan():
					wg.Done()
					t.Stop()
				}
			}
		}).Start()
		wg.Wait()
	})

	It("stops when a context is canceled", func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*40)
		ticker := provider.Create(time.Millisecond*20, func(t Ticker) {
			for {
				select {
				case <-ctx.Done():
					t.Stop()
					return
				}
			}
		}).Start()
		// wait for context cancel
		<-ticker.Chan()
	})
})
