// Template file only!
// Saving some time each day to ensure I
// actually write tests.
package main

import "testing"

type SpecialTest struct {
	Field    string // this will likely be a Struct to be tested.
	Expected bool   // expected test result.
}

// Test only exists to prevent errors below.
// These should be implemented in the non-test
// package.
func (st SpecialTest) Test() bool {
	return true
}

func TestYourFunctions(t *testing.T) {
	tests := []SpecialTest{
		{"State expected to be true", true},
		{"State expected to be false", false},
	}

	for _, test := range tests {
		if test.Test() != test.Expected {
			t.Errorf("input: %+v - wanted %t", test.Field, test.Expected)
		}
	}
}
