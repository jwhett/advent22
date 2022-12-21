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
}

func TestLast(t *testing.T) {
	if init, last := Last(helperSlice); init[0] != 1 || last != 5 {
		t.Errorf("Pop failed. Expected init: [1 2 3 4], got %v. Expected last: 5, got %d\n", init, last)
	}
}

func TestTakeN(t *testing.T) {
	if remaining, taken := TakeN(2, helperSlice); remaining[0] != 1 || taken[0] != 4 {
		t.Errorf("Pop failed. Expected remaining: [1 2 3], got %v. Expected taken: [4 5], got %v\n", remaining, taken)
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

}
