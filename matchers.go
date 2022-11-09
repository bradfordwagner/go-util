package bwutil

import "fmt"

// NewMatcherOneOf - creates a matcher where we expect the value invoked to be OneOf the inputs of type T
func NewMatcherOneOf[T comparable](a []T) MatcherOneOf[T] {
	return MatcherOneOf[T]{
		set: NewSetFromSlice(a),
	}
}

type MatcherOneOf[T comparable] struct {
	set *Set[T]
}

func (o MatcherOneOf[T]) Matches(x interface{}) (ok bool) {
	t, ok := x.(T)
	if !ok {
		return
	}
	return o.set.Exists(t)
}

func (o MatcherOneOf[T]) String() string {
	return fmt.Sprintf("is one of %v", o.set.Keyset())
}

// end MatcherOneOf

// NewMatcherConversionOneOf - allows converting into a type K from O
// then matching with an expected set of values of type K
func NewMatcherConversionOneOf[K comparable, O any](set []K, convert func(O) K) MatcherConversionOneOf[K, O] {
	return MatcherConversionOneOf[K, O]{
		convert: convert,
		set:     NewSetFromSlice(set),
	}
}

type MatcherConversionOneOf[K comparable, O any] struct {
	convert func(O) K
	set     *Set[K]
}

func (c MatcherConversionOneOf[K, O]) Matches(x interface{}) (ok bool) {
	o, ok := x.(O)
	if !ok {
		return
	}

	return c.set.Exists(c.convert(o))
}

func (c MatcherConversionOneOf[K, O]) String() string {
	return fmt.Sprintf("is one of %v", c.set.Keyset())
}

// end MatcherConversionOneOf
