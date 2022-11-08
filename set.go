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

// NewSetFromSlice - builds a set from a slice
func NewSetFromSlice[T comparable](s []T) *Set[T] {
	set := NewSet[T]()
	for _, t := range s {
		set.Add(t)
	}
	return set
}

// Add - adds a value to the set
func (s *Set[T]) Add(t T) {
	s.v[t] = true
}

// Remove - removes a value from the set
func (s *Set[T]) Remove(t T) {
	delete(s.v, t)
}

// RemoveAll - removes all elements of a slice from the set
func (s *Set[T]) RemoveAll(t []T) {
	for _, v := range t {
		s.Remove(v)
	}
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

// Copy - create a new copy of the set
func (s *Set[T]) Copy() *Set[T] {
	return NewSetFromSlice(s.Keyset())
}

// Difference - returns the elements which reside in the universal set, but not b
func (s *Set[T]) Difference(b *Set[T]) (res *Set[T]) {
	res = s.Copy()
	for k := range s.v {
		if b.Exists(k) {
			res.Remove(k)
		}
	}
	return
}

// Equals - checks two sets for equality
func (s *Set[T]) Equals(b *Set[T]) (equals bool) {
	equals = len(s.v) == len(b.v)
	if !equals {
		return
	}

	for k := range s.v {
		equals = b.Exists(k)
		if !equals {
			return
		}
	}
	return true
}
