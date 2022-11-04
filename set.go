package bwutil

// Set - set data structure
type Set[T comparable] struct {
	v map[T]bool
}

// NewSet - creates a new instance of the set
func NewSet[T comparable]() *Set[T] {
	v := make(map[T]bool)
	return &Set[T]{v: v}
}

// Add - adds a value to the set
func (s *Set[T]) Add(t T) {
	s.v[t] = true
}

// Remove - removes a value from the set
func (s *Set[T]) Remove(t T) {
	delete(s.v, t)
}

// Exists - checks if a value exists in the set
func (s *Set[T]) Exists(t T) bool {
	ok, _ := s.v[t]
	return ok
}

// Size - returns the num of elements in the set
func (s *Set[T]) Size() int {
	return len(s.v)
}

// IsEmpty - tells if the set is Empty
func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Keyset - returns the keyset as a slice of t
func (s *Set[T]) Keyset() (res []T) {
	res = make([]T, len(s.v))
	var i int
	for k := range s.v {
		res[i] = k
		i++
	}
	return
}
