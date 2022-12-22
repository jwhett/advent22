package helpers

import "testing"

var helperSlice []int

func init() {
	helperSlice = []int{1, 2, 3, 4, 5}
}

func TestPop(t *testing.T) {
	if first, rest := Pop(helperSlice); first != 1 || rest[0] != 2 {
		t.Errorf("Pop failed. Expected first: 1, got %d. Expected rest: [2 3 4 5], got %v\n", first, rest)
	}
	if emptyF, emptyR := Pop([]int{}); emptyF != 0 || len(emptyR) != 0 {
		t.Error("Popping an empty collection didn't result in empty return values.")
	}
	if singleItem, alsoEmptyR := Pop([]int{9}); singleItem != 9 || len(alsoEmptyR) != 0 {
		t.Error("Popping a collection with a single item didn't return a singel item and/or empty remaining slice.")
	}
}

func TestLast(t *testing.T) {
	if init, last := Last(helperSlice); init[0] != 1 || last != 5 {
		t.Errorf("Last failed. Expected init: [1 2 3 4], got %v. Expected last: 5, got %d\n", init, last)
	}
	if emptyF, emptyR := Last([]int{}); len(emptyF) != 0 || emptyR != 0 {
		t.Error("Last of an empty collection didn't result in empty return values.")
	}
	if alsoEmptyF, singleItem := Last([]int{9}); len(alsoEmptyF) != 0 || singleItem != 9 {
		t.Error("Last for single-item collection didn't result in empty init or single item.")
	}
}

func TestTakeN(t *testing.T) {
	if remaining, taken := TakeN(2, helperSlice); remaining[0] != 1 || taken[0] != 4 {
		t.Errorf("TakeN failed. Expected remaining: [1 2 3], got %v. Expected taken: [4 5], got %v\n", remaining, taken)
	}
	if emptyRemaining, zeroValue := TakeN(3, []int{}); len(emptyRemaining) != 0 || len(zeroValue) != 0 {
		t.Error("TakeN from an empty collection didn't result in empty return values.")
	}
	if sameSlice, alsoEmptyR := TakeN(-2, helperSlice); !SliceEqual(helperSlice, sameSlice) || len(alsoEmptyR) != 0 {
		t.Error("TakeN number of negative elements didn't result in empty return values.")
	}
	if alsoEmptyRemaining, alsoSameSlice := TakeN(7, helperSlice); len(alsoEmptyRemaining) != 0 || !SliceEqual(helperSlice, alsoSameSlice) {
		t.Error("TakeN of more elements than collection contains doesn't return the same slice or empty remaining slice.")
	}
}

func TestSliceEqual(t *testing.T) {
	// Setup our vars two ways. First from
	// a []int literal, second by appending
	// elements to an initially empty slice.
	first := []int{1, 2, 3, 4, 5}
	second := make([]int, 0)
	for i := 1; i < 6; i++ {
		second = append(second, i)
	}

	firstString, secondString := []string{"one", "two"}, []string{"one", "two"}
	firstFloat, secondFloat := []float32{1.2, 3.4, 5.6}, []float32{1.2, 3.4, 5.6}
	if !SliceEqual(first, second) || !SliceEqual(firstString, secondString) || !SliceEqual(firstFloat, secondFloat) {
		t.Error("Equal slices failed to match.")
	}

	failIntContent := []int{2, 3, 4, 5, 6}
	if SliceEqual(first, failIntContent) {
		t.Error("Different slices with equal length matched.")
	}
	failIntLength := []int{1, 2, 3}
	if SliceEqual(first, failIntLength) {
		t.Error("Slices with different lengths matched.")
	}

	if !SliceEqual([]int{}, []int{}) {
		t.Error("Two empty slices of the same type weren't equal.")
	}

}
