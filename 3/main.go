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

// TODO Structs?
// Maybe a struct to hold each inventory. Like,
// type Inventory struct { Inv (the combined inv), firstHalf, secondHalf }
// then maybe the following become methods on Inventory.

// Split will split a given string in half
// and return the two halves.
func Split(line string) (front, back string) {
	front = line[:len(line)/2]
	back = line[len(line)/2:]
	return
}

// Unique will take a string and return
// a string with only the unique values.
func Unique(line string) (unique string) {
	for _, c := range line {
		if !strings.ContainsRune(unique, c) {
			unique = unique + string(c)
		}
	}
	return
}

// FindDupe will find and return the
// rune that appears in both strings.
func FindDupe(first, second string) (rune, error) {
	var found rune
	if len(first) == 0 {
		return found, errors.New("Empty list")
	} else if len(second) == 0 {
		return found, errors.New("Empty list")
	}
	for _, c := range first {
		if strings.Contains(first, string(c)) {
			found = c
			break
		}
	}
	return found, nil
}

func main() {
	// read input
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// for each line in input file...
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// TODO Option 1
		// Split() the input line
		// Unique() both halves, less to search
		// FindDupe()
		// Increment the priority tracker by
		//+ priority of dupe found
	}

	// TODO Option 2
	// Instead of loop above, we could read
	// parse the inventories in a goroutine
	// and have a channel that takes the
	// priority of the duplicate item.

	// errors?
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// TODO Testing
	// This looks useful. Though maybe we could use
	// strings.IndexFunc(s, f()) instead and pack the
	// logic into f().
	indexOfa := strings.Index(priorities, "a")
	indexOfA := strings.Index(priorities, "A")
	fmt.Printf("Index of a: %d\nIndex of A: %d\n", indexOfa, indexOfA)
}
