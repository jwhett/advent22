package helpers

import "golang.org/x/exp/constraints"

// SliceEqual does the same as the above, but
// isn't restricted to just ints.
func SliceEqual[V constraints.Ordered](first, second []V) bool {
	// slices that differ in length are not equal
	if len(first) != len(second) {
		return false
	}
	// check each item for equality in order
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

// Pop removes and returns the first element in a
// slice and the rest of the slice.
func Pop[T any](collection []T) (head T, tail []T) {
	head, tail = collection[0], collection[1:]
	return
}

// Last removes the Last element from a slice and
// returns the remaining slice and the final element.
func Last[T any](collection []T) (init []T, last T) {
	init, last = collection[:len(collection)-1], collection[len(collection)-1]
	return
}

// TakeN will remove N elements from the end of
// xs and return both the remaining xs and the
// taken elements. The taken elements retain
// the order they were in the original slice.
func TakeN[T any](count int, from []T) (remaining []T, taken []T) {
	remaining, taken = from[:len(from)-count], from[len(from)-count:]
	return
}