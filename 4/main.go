package main

import (
	"fmt"
	"io"
	"os"
)

const (
	inputFile = "input"
)

// Duty contains a range of work duty. Mostly an
// initialization shortcut and isn't strictly
// necessary.
type Duty struct {
	Start, Stop int
}

// A Pair represents Duty assigned to a pair of elves.
type Pair struct {
	First, Second Duty
}

// FullyRedundantWorkDuty will return true when either Duty
// fully contians the other.
func (p Pair) FullyRedundantWorkDuty() bool {
	// There's surely a better way to do this.
	switch {
	case p.First.Start <= p.Second.Start && p.First.Stop >= p.Second.Stop:
		// p.First fully contains p.Second Duty
		return true
	case p.First.Start >= p.Second.Start && p.First.Stop <= p.Second.Stop:
		// p.Second fully contains p.First Duty
		return true
	default:
		return false
	}
}

// PartiallyRedundantWorkDuty will return true when either Duty
// overlaps partially overlaps the other.
func (p Pair) PartiallyRedundantWorkDuty() bool {
	// There's surely a better way to do this.
	switch {
	case p.FullyRedundantWorkDuty():
		return true
	case p.First.Start <= p.Second.Start && p.First.Stop >= p.Second.Start:
		return true
	case p.Second.Start <= p.First.Start && p.Second.Stop >= p.First.Stop:
		return true
	default:
		return false
	}
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var dutyCounter, fullyRedundantCounter, partiallyRedundantCounter int
	var first, second Duty
	for {
		if _, err := fmt.Fscanf(file, "%d-%d,%d-%d\n", &first.Start, &first.Stop, &second.Start, &second.Stop); err == io.EOF {
			// nothing more to process; end of file.
			break
		}

		p := Pair{first, second}
		if p.FullyRedundantWorkDuty() {
			fullyRedundantCounter++
		}
		if p.PartiallyRedundantWorkDuty() {
			partiallyRedundantCounter++
		}
		dutyCounter++
	}
	fmt.Printf("Out of %d work duties, there are:\n", dutyCounter)
	fmt.Printf("- Fully redundant work duties: %d\n", fullyRedundantCounter)
	fmt.Printf("- Partially redundant work duties: %d\n", partiallyRedundantCounter)
}
