package bwutil

import (
	"time"
)

type TickProviderImpl struct {
}

// TickerImpl wraps time.Ticker to allow for mock calls
type TickerImpl struct {
	t *time.Ticker
}

// Chan - returns the channel for the ticker
func (t *TickerImpl) Chan() <-chan time.Time {
	return t.t.C
}

// Stop - stops the ticker
func (t *TickerImpl) Stop() {
	t.t.Stop()
}

// TickProvider - provides a way to create tickers
func NewTickProvider() TickProvider {
	return &TickProviderImpl{}
}

// Create - creates a new ticker
func (t TickProviderImpl) Create(duration time.Duration) Ticker {
	return &TickerImpl{t: time.NewTicker(duration)}
}
