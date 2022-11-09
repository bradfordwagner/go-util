package bwutil

// MathMin - gives the lower value
func MathMin[T Number](a, b T) (res T) {
	res = a
	if b < a {
		res = b
	}
	return
}

func MathMax[T Number](a, b T) (res T) {
	res = a
	if b > a {
		res = b
	}
	return
}
