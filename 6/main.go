package main

import (
	"bufio"
	"fmt"
	"os"

	df "github.com/jwhett/advent22/6/daysix"
)

const (
	inputFile = "input"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if result := df.IdentifyTransmissionBit(line); result.Index > 0 {
			fmt.Printf("Found marker %q after processing %d characters.\n", result.Marker, result.Index)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}
}
