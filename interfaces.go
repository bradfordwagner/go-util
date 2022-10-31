package bwutil

import "time"

// Now - helps us to mock the time now function for unit testing
type Now interface {
	Now() time.Time
}
