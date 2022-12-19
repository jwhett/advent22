package main

import (
	"bufio"
	"fmt"
	df "github.com/jwhett/advent22/5/dayfive"
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

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	stacks, moves, err := df.ScanInput(bufio.NewReader(file), df.MapDimensions{MapLength: mapLength, MapHeight: mapHeight, MapCols: mapCols})
	if err != nil {
		fmt.Printf("Got an error scanning input: %v", err)
	}

	mover := df.Mover{Stacks: stacks, Moves: moves}
	mover.MoveAll(df.Standard)

	// Build the answer
	var answer string
	for _, last := range mover.Lasts() {
		answer += string(last)
	}
	fmt.Printf("Top of the stacks: %s\n", answer)
}
