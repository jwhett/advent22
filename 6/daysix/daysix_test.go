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
		ExpectedIndex  int
	}{
		{Freq: "", ExpectedMarker: "", ExpectedIndex: -1},
		{Freq: "aabbccddeeffgghhiijjkkllmmnn", ExpectedMarker: "", ExpectedIndex: -1},
		{Freq: "bvwbjplbgvbhsrlpgdmjqwftvncz", ExpectedMarker: "vwbj", ExpectedIndex: 5},
		{Freq: "nppdvjthqldpwncqszvftbrmjlhg", ExpectedMarker: "pdvj", ExpectedIndex: 6},
		{Freq: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", ExpectedMarker: "rfnt", ExpectedIndex: 10},
		{Freq: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", ExpectedMarker: "zqfr", ExpectedIndex: 11},
	}

	for _, test := range tests {
		if result := IdentifyTransmissionBit(test.Freq); result.Index != test.ExpectedIndex || result.Marker != test.ExpectedMarker {
			t.Errorf("Failed to identify transmission bit.\nGot: marker=> %q count => %d\nWanted: marker=> %q count => %d\n", result.Marker, result.Index, test.ExpectedMarker, test.ExpectedIndex)
		}
	}
}
