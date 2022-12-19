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

var (
	inputReader   *InputReader
	mapDimensions MapDimensions
)

func init() {
	inputReader = NewInputReader(bufio.NewScanner(PrepareTestInput()))
	mapDimensions = MapDimensions{testMapLength, testMapHeight, testMapCols}
}

// PrepareTestInput wraps the testInput in an
// io.Reader for consumption.
func PrepareTestInput() io.Reader {
	return strings.NewReader(testInput)
}

func TestParseMap(t *testing.T) {
	stacks := inputReader.ParseMap(mapDimensions)
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
	count, moves := inputReader.ParseMoves()
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
	stacks, moves, err := ScanInput(PrepareTestInput(), mapDimensions)
	if err != nil {
		t.Errorf("Got an error when scanning test input: %v", err)
	}
	if len(stacks) != 3 && len(moves) != 4 {
		t.Errorf("Unexpected number of stacks or moves parsed.\nStacks: Got %d, wanted 3\nMoves: Got %d, wanted 4\n", len(stacks), len(moves))
	}
}

func TestMove(t *testing.T) {
	stacks, moves, err := ScanInput(PrepareTestInput(), mapDimensions)
	if err != nil {
		t.Errorf("Got an error when scanning test input: %v", err)
	}
	mover := Mover{stacks, moves}

	// Test a single Move()
	mover.Move(Standard)
	lasts := mover.Lasts()
	testLasts := []CrateID{'D', 'C', 'P'}
	if lasts[0] != testLasts[0] {
		t.Errorf("Move didn't result in expected configuration.\nGot: %v, Wanted: %v", lasts, testLasts)
	}
}

func TestMoveAll(t *testing.T) {
	stacks, moves, err := ScanInput(PrepareTestInput(), mapDimensions)
	if err != nil {
		t.Errorf("Got an error when scanning test input: %v", err)
	}

	// Perform all (remaining) moves
	mover := Mover{stacks, moves}
	mover.MoveAll(Standard)

	lasts := mover.Lasts()
	testLasts := []CrateID{'C', 'M', 'Z'}
	if lasts[0] != testLasts[0] {
		t.Errorf("MoveAll didn't result in expected configuration.\nGot: %v, Wanted: %v\nStacks: %+v\n", lasts, testLasts, mover.Stacks)
	}
}
