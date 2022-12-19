package dayfive

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// CrateID represents the rune stamped on each crate.
type CrateID rune

// A Stack is list of CrateIDs. The higher the index,
// the higher in the stack; last item is the top.
type Stack []CrateID

// Stacks map a column or "stack" number to a Stack
// of crates.
type Stacks map[int]Stack

// A Move represents a single instruction
// for moving a set of crates between Stacks.
type Move struct {
	Count, From, To int
}
type Moves []Move

// InputReader wraps an io.Reader to perform
// input-specific parsing of that io.Reader.
type InputReader struct {
	*bufio.Scanner
}

func NewInputReader(scanner *bufio.Scanner) *InputReader {
	return &InputReader{scanner}
}

type MapDimensions struct {
	MapLength int
	MapHeight int
	MapCols   int
}

// ParseMap will parse the ASCII art map showing
// the location of all crates. This results in
// a full set of Stacks.
func (ir InputReader) ParseMap(md MapDimensions) (stacks Stacks) {
	const crateWidth = 4
	maxStackLength := md.MapHeight
	stacks = make(Stacks)
	// initialize the stacks
	for i := 1; i <= md.MapCols; i++ {
		stacks[i] = make(Stack, maxStackLength)
	}

	// scan rows top to bottom
	row := maxStackLength - 1
	// for each row of the ASCII map of crates...
ScannerLoop:
	for ir.Scan() {
		line := ir.Text()
		if len(line) == 0 {
			// empty line
			continue
		}
		for cursorPosition, column := 0, 1; cursorPosition < md.MapLength; cursorPosition, column = cursorPosition+crateWidth, column+1 {
			// look for a crate at a given slice of the row..
			var substring string
			if cursorPosition+crateWidth >= len(line)-1 {
				substring = line[cursorPosition:]
			} else {
				substring = line[cursorPosition : cursorPosition+crateWidth]
			}
			for _, c := range substring {
				// store the crate ID if we find one in this slice!
				if unicode.IsLetter(c) {
					// found the crate in this range
					stacks[column][row] = CrateID(c)
					break
				}
			}
		}
		// descending; decrement our row counter/location
		row--
		if row < 0 {
			break ScannerLoop
		}
	}
	for key, stack := range stacks {
		filterStack := make(Stack, 0)
		for _, v := range stack {
			if v > 0 {
				filterStack = append(filterStack, v)
			}
		}
		stacks[key] = filterStack
	}
	return
}

// ParseMoves will parse the io.Reader for move instructions and return the
// number of moves that were read as well as the list of moves. It is expected
// that moveCount is greater than zero on a successful parse.
func (ir InputReader) ParseMoves() (moveCount int, moves Moves) {
	moveCount = 0
	moves = make(Moves, 0)
	var count, from, to int
	for ir.Scan() {
		line := ir.Text()
		if !strings.HasPrefix(line, "move") {
			// not a move instruction
			continue
		}
		fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
		moves = append(moves, Move{count, from, to})
		moveCount++
	}
	return
}

// ScanInput is the entrypoint for parsing the input
// file for the puzzle.
func ScanInput(r io.Reader, md MapDimensions) (stacks Stacks, moves Moves, err error) {
	var count int
	ir := NewInputReader(bufio.NewScanner(r))
	stacks = ir.ParseMap(md)
	count, moves = ir.ParseMoves()
	if count == 0 {
		err = errors.New("No moves were parsed.")
	}
	return
}

// A Mover is in charge of moving.
type Mover struct {
	Stacks
	Moves
}

// Lasts returns the last crate in each stack.
func (m Mover) Lasts() []CrateID {
	lasts := make([]CrateID, 0)
	for i := 1; i <= len(m.Stacks); i++ {
		if len(m.Stacks[i]) == 0 {
			// This stack is empty.
			lasts = append(lasts, CrateID(' '))
			continue
		}
		lasts = append(lasts, m.Stacks[i][len(m.Stacks[i])-1])
	}
	return lasts
}

// Move performs a single move instruction
func (m *Mover) Move() {
	var move Move
	var init Stack
	var last CrateID

	if len(m.Moves) == 0 {
		// No moves to perform
		return
	}

	move, m.Moves = m.Moves[0], m.Moves[1:]
	for i := 0; i < move.Count; i++ {
		if len(m.Stacks[move.From]) == 1 {
			init, last = make(Stack, 0), m.Stacks[move.From][0]
		} else {
			init, last = m.Stacks[move.From][:len(m.Stacks[move.From])-1], m.Stacks[move.From][len(m.Stacks[move.From])-1]
		}
		m.Stacks[move.From] = init
		m.Stacks[move.To] = append(m.Stacks[move.To], last)
	}
}

// Perform all moves
func (m *Mover) MoveAll() {
	// len(m.Moves) changes each loop iteration so
	// we have to save it before starting the loop.
	maxMoves := len(m.Moves)
	for i := 0; i < maxMoves; i++ {
		m.Move()
	}
}
