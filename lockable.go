package bwutil

import "sync"

// Lockable - Structure to help with creating a thread safe object
type Lockable[T any] struct {
	object T
	m      *sync.RWMutex
}

// NewLockable - creates a new lockable object of type t
func NewLockable[T any](object T) *Lockable[T] {
	return &Lockable[T]{
		object: object,
		m:      new(sync.RWMutex),
	}
}

// Set - sets the value of Lockable
func (l *Lockable[T]) Set(obj T) {
	l.m.Lock()
	defer l.m.Unlock()
	l.object = obj
}

// Get - gets the value from Lockable
func (l *Lockable[T]) Get() T {
	l.m.RLock()
	defer l.m.RUnlock()
	return l.object
}