package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	inputFile  = "input"
	priorities = "!abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Split will split a given string in half and return the two halves.
func Split(line string) (front, back string) {
	front = line[:len(line)/2]
	back = line[len(line)/2:]
	return
}

// FindDupe will return the rune that appears in both strings.
func FindDupe(first, second string) (rune, error) {
	var found rune
	if len(first) == 0 || len(second) == 0 {
		return found, errors.New("Empty list")
	}
	for _, c := range first {
		if strings.Contains(second, string(c)) {
			found = c
			break
		}
	}
	return found, nil
}

func Process(line string) (dupe rune) {
	front, back := Split(line)
	dupe, err := FindDupe(front, back)
	if err != nil {
		fmt.Printf("Error when processing: %v", err)
		return
	}
	return
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var totalPriority int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalPriority = totalPriority + strings.IndexRune(priorities, Process(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}

	fmt.Printf("Total priority: %d\n", totalPriority)
}
