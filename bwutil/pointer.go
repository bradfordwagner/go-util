package bwutil

// Pointer - returns a pointer version of any object
func Pointer[T any](t T) *T {
	return &t
}
