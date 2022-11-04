package bwutil

import "sync"

// Lockable - Structure to help with creating a thread safe object
type Lockable[T any] struct {
	object T
	m      *sync.RWMutex
}

// NewLockableWithValue - creates a new lockable object of type t
func NewLockableWithValue[T any](object T) *Lockable[T] {
	return &Lockable[T]{
		object: object,
		m:      new(sync.RWMutex),
	}
}

// NewLockable - creates a new lockable object of type t
func NewLockable[T any]() *Lockable[T] {
	return &Lockable[T]{
		m: new(sync.RWMutex),
	}
}

// Set - sets the value of Lockable
func (l *Lockable[T]) Set(obj T) {
	l.m.Lock()
	defer l.m.Unlock()
	l.object = obj
}

// SetF - sets the value of the lockable by invoking a function
func (l *Lockable[T]) SetF(f func() T) {
	l.m.Lock()
	defer l.m.Unlock()
	l.object = f()
}

// Get - gets the value from Lockable
func (l *Lockable[T]) Get() T {
	l.m.RLock()
	defer l.m.RUnlock()
	return l.object
}
