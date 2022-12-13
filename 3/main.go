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

func ProcessPartTwo(lines []string) (badgePriority int) {
	for _, c := range lines[0] {
		if strings.ContainsRune(lines[1], c) && strings.ContainsRune(lines[2], c) {
			badgePriority = strings.IndexRune(priorities, c)
			break
		}
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
	var badgePriority int
	var line string

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		// part 1
		totalPriority = totalPriority + strings.IndexRune(priorities, Process(line))
		// part 2
		lines = append(lines, line)
		if len(lines) == 3 {
			badgePriority = badgePriority + ProcessPartTwo(lines)
			lines = make([]string, 0)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}

	fmt.Printf("Total priority: %d\n", totalPriority)
	fmt.Printf("Badge priority: %d\n", badgePriority)
}
