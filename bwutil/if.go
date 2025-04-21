package bwutil

func If[T any](check bool, t, f T) (v T) {
	v = f
	if check {
		v = t
	}
	return
}

func IfDo(check bool, f func()) {
	if check {
		f()
	}
}
