package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	inputFile = "input"
)

// Split will split a given string in half
// and return the two halves.
func Split(line string) (front, back string) {
	// TODO
	// split at len(line)/2
	return
}

// FindDupe will find and return the
// character that appears in both strings.
func FindDupe(first, second string) (found byte, err error) {
	// TODO
	// error on no match? or empty/nil instead?
	// sort the strings?
	// for i in first, attempt match against j in second?
	return
}

func main() {
	// read input
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	priorities := []byte("!abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// for each line in input file...
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// TODO
		// Split() the input line
		// FindDupe()
		// Increment the priority tracker
	}

	// Testing
	indexOfa := bytes.IndexByte(priorities, byte('a'))
	indexOfA := bytes.IndexByte(priorities, byte('A'))
	fmt.Printf("Index of a: %d\nIndex of A: %d\n", indexOfa, indexOfA)

	// errors?
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
