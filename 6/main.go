package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	df "github.com/jwhett/advent22/6/daysix"
)

const (
	inputFile = "input"
)

var (
	markerEnv    string
	markerLength int
	err          error
)

func init() {
	var ok bool
	if markerEnv, ok = os.LookupEnv("MARKER_LENGTH"); !ok {
		markerLength = 4
	} else {
		if markerLength, err = strconv.Atoi(markerEnv); err != nil {
			fmt.Printf("%q is not a proper marker length. Please set MARKER_LENGTH to an integer and try again.\n", markerEnv)
			os.Exit(1)
		}
	}
}

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
		if result := df.IdentifyTransmissionBit(line, markerLength); result.Index > 0 {
			fmt.Printf("Found marker %q after processing %d characters.\n", result.Marker, result.Index)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}
}
