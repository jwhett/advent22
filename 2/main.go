package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type shape int
type result int

const (
	inputFile             string = "input"
	ROCK, PAPER, SCISSORS shape  = 1, 2, 3 // Shape values
	LOSE, DRAW, WIN       result = 0, 3, 6 // Result values
)

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

func duel(theirShape, myShape shape) result {
	switch {
	case theirShape == myShape:
		return DRAW
	case myShape == ROCK && theirShape == SCISSORS:
		return WIN
	case myShape == PAPER && theirShape == ROCK:
		return WIN
	case myShape == SCISSORS && theirShape == PAPER:
		return WIN
	default:
		return LOSE
	}
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

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
		fmt.Printf("Their shape: %d, My shape: %d\n", theirShape, myShape)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
