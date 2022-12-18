package dayfive

import (
	"bufio"
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
	ir := NewInputReader(bufio.NewScanner(inReader))
	stacks := ir.ParseMap(MapDimensions{testMapLength, testMapHeight, testMapCols})
	// Ensure we receive the proper number of stacks from
	// the ASCII map.
	if len(stacks) != testMapCols {
		t.Errorf("Wrong number of stacks collected. Wanted 3 got %d\n-> %v", len(stacks), stacks)
	}
	// Ensure we have creates in the right order.
	bottomRow := []rune{'Z', 'M', 'P'}
	for stackCount, rowCount := 1, 0; stackCount < len(stacks); stackCount, rowCount = stackCount+1, rowCount+1 {
		if stacks[stackCount][0] != CrateID(bottomRow[rowCount]) {
			t.Errorf("Crate in the wrong place. Got %q, expected %q", stacks[stackCount][0], bottomRow[rowCount])
		}
	}
}

func TestParseMoves(t *testing.T) {
	ir := NewInputReader(bufio.NewScanner(PrepareTestInput()))
	count, moves := ir.ParseMoves()
	if count == 0 {
		t.Error("Failed to parse moves from the test input.")
	}
	testMoves := Moves{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}
	for i, move := range moves {
		if move != testMoves[i] {
			t.Errorf("Unexpected move found! Got %+v, wanted %+v\n", move, testMoves[i])
		}
	}
}

func TestScanInput(t *testing.T) {
	stacks, moves, err := ScanInput(PrepareTestInput(), MapDimensions{testMapLength, testMapHeight, testMapCols})
	if err != nil {
		t.Errorf("Got an error when scanning test input: %v", err)
	}
	if len(stacks) != 3 && len(moves) != 4 {
		t.Errorf("Unexpected number of stacks or moves parsed.\nStacks: Got %d, wanted 3\nMoves: Got %d, wanted 4\n", len(stacks), len(moves))
	}
}
