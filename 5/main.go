package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

const (
	inputFile = "input"
	// mapLength is the maximum line length in bytes
	// including the newline.
	mapLength = 36
	// mapHeight is the maximum height of each "stack"
	// in the input.
	mapHeight = 8
	// mapCols represents the number of columns or
	// "stacks" of crates.
	mapCols = 9
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
	io.Reader
}

// ParseMap will parse the ASCII art map showing
// the location of all crates. This results in
// a full set of Stacks.
func (ir InputReader) ParseMap(mapLength, mapHeight, mapCols int) (stacks Stacks) {
	const crateWidth = 4
	maxStackLength := mapHeight
	stacks = make(Stacks)
	// initialize the stacks
	for i := 1; i <= mapCols; i++ {
		stacks[i] = make(Stack, maxStackLength)
	}

	// scan rows top to bottom
	row := maxStackLength - 1
	scanner := bufio.NewScanner(ir)
	// for each row of the ASCII map of crates...
ScannerLoop:
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			// empty line
			continue
		}
		for cursorPosition, column := 0, 1; cursorPosition < mapLength; cursorPosition, column = cursorPosition+crateWidth, column+1 {
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
	return
}

// ParseMoves will parse the io.Reader for move instructions and return the
// number of moves that were read as well as the list of moves. It is expected
// that moveCount is greater than zero on a successful parse.
func (ir InputReader) ParseMoves() (moveCount int, moves Moves) {
	moveCount = 0
	moves = make(Moves, 0)
	scanner := bufio.NewScanner(ir)
	var count, from, to int
	for scanner.Scan() {
		line := scanner.Text()
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
func ScanInput(r io.Reader) (s []Stack, m []Move) {
	s = make([]Stack, 0)
	m = make([]Move, 0)

	return
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}
}
