package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
func (ir InputReader) ParseMap() Stacks {
	buffer := make([]byte, mapLength)
	stacks := make(Stacks, mapCols)
	for i := 1; i < mapHeight; i++ {
		_, err := ir.Read(buffer)
		if err != nil {
			fmt.Println(err)
		}
		// TODO: Parse every 4 bytes as a crate
		// and insert at the head of each Stack
		// because we're reading this from the
		// top down.
	}
	return stacks
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
