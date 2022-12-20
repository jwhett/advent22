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
