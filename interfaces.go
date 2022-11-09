package bwutil

import "time"

// Now - helps us to mock the time now function for unit testing
type Now interface {
	Now() time.Time
}

// Number - sets an interface for number types
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}
