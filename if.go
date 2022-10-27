package util

func If[T any](check bool, t, f T) (v T) {
	v = f
	if check {
		v = t
	}
	return
}
