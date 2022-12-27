package daysix

const chunkSize = 4

type Result struct {
	Index  int
	Marker string
}

func Unique(input string) bool {
	// guards
	switch {
	case len(input) == 0:
		// empty strings do not have unique charaters
		return false
	case len(input) == 1:
		// a string with a single charater is unique
		return true
	case len(input) == 2:
		// simple check for strings with two charaters
		return input[0] != input[1]
	}

	matches := make(map[rune]int)
	// take inventory
	for _, r := range input {
		matches[r]++
	}
	// there should be unique keys equal to the length
	// of chunkSize.
	return len(matches) == chunkSize
}

func IdentifyTransmissionBit(freq string) Result {
	maxLen := len(freq)
	for i := 0; i+chunkSize <= maxLen; i++ {
		chunk := freq[i : i+chunkSize]
		if Unique(chunk) {
			// we found a marker
			return Result{Index: i + chunkSize, Marker: chunk}

		}
	}
	return Result{Index: -1, Marker: ""}
}
