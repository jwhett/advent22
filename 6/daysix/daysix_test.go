package daysix

import "testing"

func TestUnique(t *testing.T) {
	tests := []struct {
		Input    string
		Expected bool
	}{
		{Input: "aabc", Expected: false},
		{Input: "abad", Expected: false},
		{Input: "", Expected: false},
		{Input: "a", Expected: true},
		{Input: "ca", Expected: true},
		{Input: "fedc", Expected: true},
	}

	for _, test := range tests {
		if Unique(test.Input) != test.Expected {
			t.Errorf("Unique failed. %q was expected to be %t\n", test.Input, test.Expected)
		}
	}
}

func TestTransmissionBitIdentification(t *testing.T) {
	tests := []struct {
		Freq           string
		ExpectedMarker string
		ExpectedCount  int
	}{
		{Freq: "", ExpectedMarker: "", ExpectedCount: -1},
		{Freq: "aabbccddeeffgghhiijjkkllmmnn", ExpectedMarker: "", ExpectedCount: -1},
		{Freq: "bvwbjplbgvbhsrlpgdmjqwftvncz", ExpectedMarker: "vwbj", ExpectedCount: 5},
		{Freq: "nppdvjthqldpwncqszvftbrmjlhg", ExpectedMarker: "pdvj", ExpectedCount: 6},
		{Freq: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", ExpectedMarker: "fntj", ExpectedCount: 10},
		{Freq: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", ExpectedMarker: "qfrl", ExpectedCount: 11},
	}

	for _, test := range tests {
		if marker, count := IdentifyTransmissionBit(test.Freq); count != test.ExpectedCount || marker != test.ExpectedMarker {
			t.Errorf("Failed to identify transmission bit.\nGot: marker=> %s count => %d\nWanted: marker=> %s count => %d\n", marker, count, test.ExpectedMarker, test.ExpectedCount)
		}
	}
}
