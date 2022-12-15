package main

import "testing"

type PairTest struct {
	Pair
	Expected bool
}

func TestFullyRedundantWorkDuty(t *testing.T) {
	tests := []PairTest{
		{Pair{Duty{2, 4}, Duty{6, 8}}, false},
		{Pair{Duty{2, 3}, Duty{4, 5}}, false},
		{Pair{Duty{5, 7}, Duty{7, 9}}, false},
		{Pair{Duty{2, 8}, Duty{3, 7}}, true},
		{Pair{Duty{6, 6}, Duty{4, 6}}, true},
		{Pair{Duty{2, 6}, Duty{4, 8}}, false},
		{Pair{Duty{4, 6}, Duty{2, 5}}, false},
	}

	for _, test := range tests {
		if test.Pair.FullyRedundantWorkDuty() != test.Expected {
			t.Errorf("input: %+v - wanted %t", test.Pair, test.Expected)
		}
	}
}

func TestPartiallyRedundantWorkDuty(t *testing.T) {
	tests := []PairTest{
		{Pair{Duty{2, 4}, Duty{6, 8}}, false},
		{Pair{Duty{2, 3}, Duty{4, 5}}, false},
		{Pair{Duty{5, 7}, Duty{7, 9}}, true},
		{Pair{Duty{2, 8}, Duty{3, 7}}, true},
		{Pair{Duty{6, 6}, Duty{4, 6}}, true},
		{Pair{Duty{2, 6}, Duty{4, 8}}, true},
		{Pair{Duty{4, 6}, Duty{2, 5}}, true},
	}

	for _, test := range tests {
		if test.Pair.PartiallyRedundantWorkDuty() != test.Expected {
			t.Errorf("input: %+v - wanted %t", test.Pair, test.Expected)
		}
	}
}
