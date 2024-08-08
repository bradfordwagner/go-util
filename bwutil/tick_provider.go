package bwutil

import (
	"time"
)

/*
TickerProviderImpl - start
*/
type TickerProviderImpl struct {
}

// NewTickerProvider - provides a way to create tickers
func NewTickerProvider() TickerProvider {
	return &TickerProviderImpl{}
}

// Create - creates a new ticker
func (t *TickerProviderImpl) Create(duration time.Duration, f TickFunction) Ticker {
	return &TickerImpl{t: time.NewTicker(duration), f: f}
}

/* end TickProviderImpl */

/*
TickerImpl - start
*/
// TickerImpl wraps time.Ticker to allow for mock calls
type TickerImpl struct {
	t *time.Ticker
	f TickFunction
}

func (t *TickerImpl) Start() Ticker {
	go t.f(t)
	return t
}

// enforce interface
var _ Ticker = &TickerImpl{}

// Stop - stops the ticker
func (t *TickerImpl) Stop() {
	t.t.Stop()
}
func (t *TickerImpl) Chan() <-chan time.Time {
	return t.t.C
}

/* end TickerImpl */
