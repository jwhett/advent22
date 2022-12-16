package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	inputFile = "input"
)

type Stack []rune

type Move struct {
	Count, From, To int
}

func ScanInput(r io.Reader) (s []Stack, m []Move) {
	s = make([]Stack, 0)
	m = make([]Move, 0)
	return
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}
}
