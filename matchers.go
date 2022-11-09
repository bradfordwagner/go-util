package bwutil

import "fmt"

func NewMatcherOneOf[T comparable](a []T) OneOfMatcher[T] {
	return OneOfMatcher[T]{
		set: NewSetFromSlice(a),
	}
}

type OneOfMatcher[T comparable] struct {
	set *Set[T]
}

func (o OneOfMatcher[T]) Matches(x interface{}) (ok bool) {
	t, ok := x.(T)
	if !ok {
		return
	}
	return o.set.Exists(t)
}

func (o OneOfMatcher[T]) String() string {
	return fmt.Sprintf("is one of %v", o.set.Keyset())
}

// end OneOfMatcher

// NewConversionOneOfMatcher - allows converting into a type K from O
// then matching with an expected set of values of type K
func NewConversionOneOfMatcher[K comparable, O any](set []K, convert func(O) K) ConversionOneOfMatcher[K, O] {
	return ConversionOneOfMatcher[K, O]{
		convert: convert,
		set:     NewSetFromSlice(set),
	}
}

type ConversionOneOfMatcher[K comparable, O any] struct {
	convert func(O) K
	set     *Set[K]
}

func (c ConversionOneOfMatcher[K, O]) Matches(x interface{}) (ok bool) {
	o, ok := x.(O)
	if !ok {
		return
	}

	return c.set.Exists(c.convert(o))
}

func (c ConversionOneOfMatcher[K, O]) String() string {
	return fmt.Sprintf("is one of %v", c.set.Keyset())
}
