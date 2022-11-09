package bwutil

// SliceMap - helper function to convert from A to B
func SliceMap[A any, B any](sa []A, f func(a A) (b B)) (sb []B) {
	sb = make([]B, len(sa))
	for i, a := range sa {
		sb[i] = f(a)
	}
	return
}
