package helpers

import "golang.org/x/exp/constraints"

// SliceEqual will attempt to test each slice element
// for equality in order.
func SliceEqual[V constraints.Ordered](first, second []V) bool {
	// two empty slices of the same type are equal
	if len(first) == 0 && len(second) == 0 {
		return true
	}
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

// Pop removes and returns the first element of collection
// and the resulting slice. The head of a single item
// is itself and an empty slice.
func Pop[T any](collection []T) (head T, tail []T) {
	if len(collection) == 0 {
		return
	}
	if len(collection) == 1 {
		return collection[0], tail
	}
	head, tail = collection[0], collection[1:]
	return
}

// Last removes the Last element from collection and
// returns the resulting slice and the final element.
// The last of a single item is an empty slice and
// itself.
func Last[T any](collection []T) (init []T, last T) {
	if len(collection) == 0 {
		return
	}
	if len(collection) == 1 {
		return init, collection[0]
	}
	init, last = collection[:len(collection)-1], collection[len(collection)-1]
	return
}

// TakeN will remove count elements from the end of
// from and return slices of both the remaining elements
// and the taken elements. The taken elements retain
// the order of the original slice. Will return empty
// slices if the from slice is empty. Attempting to
// take a negative number of elements will result
// in returning the given slice unaltered and an
// empty taken slice. Taking more elements than that
// which exists in the given collection will result
// in an empty remaining slice and the given collection.
func TakeN[T any](count int, from []T) (remaining []T, taken []T) {
	if len(from) == 0 {
		return
	}
	if count < 0 {
		return from, taken
	}
	if count > len(from) {
		return remaining, from
	}
	remaining, taken = from[:len(from)-count], from[len(from)-count:]
	return
}