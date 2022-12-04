package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// My types
type Shape int
type Result int
type ShapeRules map[string]Rule
type CheatResults map[string]Result

type Rule struct {
	Shape   Shape
	Beats   Shape
	LosesTo Shape
}

const (
	inputFile             string = "input"
	ROCK, PAPER, SCISSORS Shape  = 1, 2, 3 // Shape values
	LOSE, DRAW, WIN       Result = 0, 3, 6 // Result values
)

// duel returns the result of two Rules.
func duel(theirShape, myShape Rule) Result {
	switch {
	case theirShape.LosesTo == myShape.Shape:
		return Result(myShape.Shape) + WIN
	case theirShape.Beats == myShape.Shape:
		return Result(myShape.Shape) + LOSE
	default:
		return Result(myShape.Shape) + DRAW
	}
}

func evalPartTwo(theirShape Rule, rules ShapeRules, cheats CheatResults) Result {
	// TODO
	return WIN
}

func main() {
	// read input
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// build a mapping of shapes to rules
	shapeRules := make(ShapeRules)
	shapeRules["A"] = Rule{ROCK, SCISSORS, PAPER}
	shapeRules["X"] = Rule{ROCK, SCISSORS, PAPER}
	shapeRules["B"] = Rule{PAPER, ROCK, SCISSORS}
	shapeRules["Y"] = Rule{PAPER, ROCK, SCISSORS}
	shapeRules["C"] = Rule{SCISSORS, PAPER, ROCK}
	shapeRules["Z"] = Rule{SCISSORS, PAPER, ROCK}

	// pt 2 changes the meaning of the second column
	cheatResults := make(CheatResults)
	cheatResults["X"] = LOSE
	cheatResults["Y"] = DRAW
	cheatResults["Z"] = WIN

	var partOneResults Result
	// var partTwoResults Result
	// for each line in input file...
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// parse the shapes in the study guide
		shapes := strings.Split(scanner.Text(), " ")
		theirShape := shapes[0]
		myShape := shapes[1]
		// add the result
		partOneResults = partOneResults + duel(shapeRules[theirShape], shapeRules[myShape])
		// TODO
		// partTwoResults = partTwoResults + evalPartTwo(shapeRules[theirShape], shapeRules, cheatResults)
	}

	// errors?
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("// Part 1 results: %d\n", partOneResults)
}
