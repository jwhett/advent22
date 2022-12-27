package daysix

import "testing"

func TestUnique(t *testing.T) {
	tests := []struct {
		Input     string
		Expected  bool
		ChunkSize int
	}{
		{Input: "aabc", Expected: false, ChunkSize: 4},
		{Input: "abad", Expected: false, ChunkSize: 4},
		{Input: "", Expected: false, ChunkSize: 4},
		{Input: "a", Expected: true, ChunkSize: 4},
		{Input: "ca", Expected: true, ChunkSize: 4},
		{Input: "fedc", Expected: true, ChunkSize: 4},
	}

	for _, test := range tests {
		if Unique(test.Input, test.ChunkSize) != test.Expected {
			t.Errorf("Unique failed. %q was expected to be %t\n", test.Input, test.Expected)
		}
	}
}

func TestTransmissionBitIdentification(t *testing.T) {
	tests := []struct {
		Freq                     string
		ExpectedMarker           string
		ExpectedIndex, ChunkSize int
	}{
		{Freq: "", ExpectedMarker: "", ExpectedIndex: -1, ChunkSize: 4},
		{Freq: "aabbccddeeffgghhiijjkkllmmnn", ExpectedMarker: "", ExpectedIndex: -1, ChunkSize: 4},
		{Freq: "bvwbjplbgvbhsrlpgdmjqwftvncz", ExpectedMarker: "vwbj", ExpectedIndex: 5, ChunkSize: 4},
		{Freq: "nppdvjthqldpwncqszvftbrmjlhg", ExpectedMarker: "pdvj", ExpectedIndex: 6, ChunkSize: 4},
		{Freq: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", ExpectedMarker: "rfnt", ExpectedIndex: 10, ChunkSize: 4},
		{Freq: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", ExpectedMarker: "zqfr", ExpectedIndex: 11, ChunkSize: 4},
	}

	for _, test := range tests {
		if result := IdentifyTransmissionBit(test.Freq, test.ChunkSize); result.Index != test.ExpectedIndex || result.Marker != test.ExpectedMarker {
			t.Errorf("Failed to identify transmission bit.\nGot: marker=> %q count => %d\nWanted: marker=> %q count => %d\n", result.Marker, result.Index, test.ExpectedMarker, test.ExpectedIndex)
		}
	}
}
