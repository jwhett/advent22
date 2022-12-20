package daysix

import "testing"

type FreqTest struct {
	Freq           string
	ExpectedMarker string
	ExpectedCount  int
}

func TestTransmissionBitIdentification(t *testing.T) {
	tests := []FreqTest{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", "vwbj", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", "pdvj", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "fntj", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "qfrl", 11},
	}

	for _, test := range tests {
		if marker, count := IdentifyTransmissionBit(test.Freq); count != test.ExpectedCount || marker != test.ExpectedMarker {
			t.Errorf("Failed to identify transmission bit. Got: marker=> %s count=> %d, Wanted marker=> %s count => %d\n", marker, count, test.ExpectedMarker, test.ExpectedCount)
		}
	}
}
