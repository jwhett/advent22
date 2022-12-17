package main

import (
	"io"
	"strings"
	"testing"
)

const (
	testInput = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	testMapLength = 12
	testMapHeight = 3
	testMapCols   = 3
)

// PrepareTestInput wraps the testInput in an
// io.Reader for consumption.
func PrepareTestInput() io.Reader {
	return strings.NewReader(testInput)
}

func TestParseMap(t *testing.T) {
	inReader := PrepareTestInput()
	ir := InputReader{inReader}
	stacks := ir.ParseMap(testMapLength, testMapHeight, testMapCols)
	if len(stacks) != testMapCols {
		t.Errorf("Wrong number of stacks collected. Wanted 3 got %d\n-> %v", len(stacks), stacks)
	} else {
		t.Logf("Passed: %+v", stacks)
	}
}
