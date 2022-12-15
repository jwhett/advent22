package main

import (
	"io"
	"strings"
	"testing"
)

const testInput = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

// PrepareTestInput wraps the testInput in an
// io.Reader for consumption.
func PrepareTestInput() io.Reader {
	return strings.NewReader(testInput)
}

func TestScanInput(t *testing.T) {
	inReader := PrepareTestInput()
	stacks, moves := ScanInput(inReader)
	if len(stacks) != 3 {
		t.Errorf("Wrong number of stacks collected. Wanted 3 got %d", len(stacks))
	}
	if len(moves) != 4 {
		t.Errorf("Wrong number of moves collected. Wanted 4 got %d", len(stacks))
	}
}
