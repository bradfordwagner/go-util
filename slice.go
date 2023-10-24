package bwutil

// SliceMap - helper function to convert from A to B
func SliceMap[A any, B any](sa []A, f func(a A) (b B)) (sb []B) {
	sb = make([]B, len(sa))
	for i, a := range sa {
		sb[i] = f(a)
	}
	return
}

// SliceToMap - helper function to modify a slice into a Map
func SliceToMap[A any, K comparable, V any](sa []A, f func(a A) (k K, v V)) (m map[K]V) {
	m = make(map[K]V)
	for _, a := range sa {
		k, v := f(a)
		m[k] = v
	}
	return
}
