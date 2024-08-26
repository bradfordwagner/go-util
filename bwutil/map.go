package bwutil

import (
	"cmp"
	"slices"
)

// MapCopy - duplicate a map into a new one
func MapCopy[k comparable, v any](m map[k]v) (res map[k]v) {
	res = make(map[k]v)
	for k, v := range m {
		res[k] = v
	}
	return
}

// ToSortedSliceByKey - sorts the map by key and returns a slice of sorted values
func ToSortedSliceByKey[k cmp.Ordered, v any](m map[k]v) (res []v) {
	keys := make([]k, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	// use sorted keys to get values in order
	res = make([]v, 0, len(m))
	for _, k := range keys {
		res = append(res, m[k])
	}

	return
}

// MapKeys - return a slice of keys from a map
func MapKeys[k comparable, v any](m map[k]v) (res []k) {
	res = make([]k, 0, len(m))
	for k, _ := range m {
		res = append(res, k)
	}
	return
}
