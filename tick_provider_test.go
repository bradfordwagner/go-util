package bwutil

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sync"
	"time"
)

var _ = Describe("TickProvider", func() {
	var provider = NewTickProvider()

	It("should bootstrap", func() { Expect(provider).NotTo(BeNil()) })

	When("creating a ticker", func() {
		It("ticker should tick", func() {
			t := time.Millisecond * 20
			ticker := provider.Create(t)
			Expect(ticker).ShouldNot(BeNil())

			wg := new(sync.WaitGroup)
			wg.Add(1)
			go func() {
				select {
				case <-ticker.Chan():
					wg.Done()
				}
			}()
			wg.Wait()
		})
		It("closed", func() {
			t := time.Millisecond * 20
			ticker := provider.Create(t)
			Expect(ticker).ShouldNot(BeNil())

			wg := new(sync.WaitGroup)
			// force a failure if wg.Done is called more than once (we canceled the ticker)
			wg.Add(1)
			ctx, _ := context.WithTimeout(context.TODO(), time.Millisecond*50)
			go func() {
				for {
					select {
					case <-ticker.Chan():
						ticker.Stop()
						// if this is called twice it will panic
						wg.Done()
					case <-ctx.Done():
						return
					}
				}
			}()
			wg.Wait()
			Eventually(ctx.Done).Should(BeClosed())
		})

	})
})
