package bwutil

import "time"

type defaultNow struct{}

// enforce interface
var _ Now = (*defaultNow)(nil)

// Now - gives the default unix time
func (d defaultNow) Now() time.Time {
	return time.Now()
}

// NewUnixNow - creates a new now with default unix now time
func NewUnixNow() Now {
	return new(defaultNow)
}
