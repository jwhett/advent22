package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// My types
type shape int
type result int

const (
	inputFile             string = "input"
	ROCK, PAPER, SCISSORS shape  = 1, 2, 3 // Shape values
	LOSE, DRAW, WIN       result = 0, 3, 6 // Result values
)

// parseShapes takes a string slice and returns a
// parsed slice of shapes.
func parseShapes(rawShapes []string) ([]shape, error) {
	shapes := make([]shape, 0)
	for _, rs := range rawShapes {
		parsed, err := parseShape(rs)
		if err != nil {
			return shapes, err
		}
		shapes = append(shapes, parsed)
	}
	return shapes, nil
}

// parseShape takes a string, assumed to be a single character,
// and returns its shape followed by any errors.
func parseShape(rawShape string) (shape, error) {
	var parsedShape shape
	switch strings.ToUpper(rawShape) {
	case "A", "X":
		parsedShape = ROCK
	case "B", "Y":
		parsedShape = PAPER
	case "C", "Z":
		parsedShape = SCISSORS
	default:
		return 0, errors.New(fmt.Sprintf("Invalid shape: %s", rawShape))
	}
	return parsedShape, nil
}

// duel takes two shapes and returns the rock/paper/scissors result
// of those two shapes. A "result" is the result of the duel,
// win/loss/draw, plus the shape that "we" played.
func duel(theirShape, myShape shape) result {
	switch {
	case theirShape == myShape:
		return DRAW + result(myShape)
	case myShape == ROCK && theirShape == SCISSORS:
		return WIN + result(myShape)
	case myShape == PAPER && theirShape == ROCK:
		return WIN + result(myShape)
	case myShape == SCISSORS && theirShape == PAPER:
		return WIN + result(myShape)
	default:
		return LOSE + result(myShape)
	}
}

// total sums all results given a slice of result.
func total(results []result) int {
	var final int
	for _, i := range results {
		final = final + int(i)
	}
	return final
}

func main() {
	// read input
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	results := make([]result, 0)
	// for each line in input file...
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// parse the shapes in the study guide
		parsedShapes, err := parseShapes(strings.Split(scanner.Text(), " "))
		if err != nil {
			fmt.Println(err)
		}
		theirShape := parsedShapes[0]
		myShape := parsedShapes[1]
		// add the result
		results = append(results, duel(theirShape, myShape))
	}

	// errors?
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// output the final results!
	fmt.Printf("// Final score: %d\n", total(results))
}
