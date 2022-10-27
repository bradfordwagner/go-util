package util

// MapCopy - duplicate a map into a new one
func MapCopy[k comparable, v any](m map[k]v) (res map[k]v) {
	res = make(map[k]v)
	for k, v := range m {
		res[k] = v
	}
	return
}
