// Solve Day 1 of AOC www.adventofcode.com
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"sync"
)

// Source data for the puzzle.
const inputFile = "input"

// Elf modes an Elf an its inventory.
type Elf struct {
	Calories int
}

// Increase the calorie count in Elf inventory.
func (e *Elf) AddToCalories(cal int) {
	e.Calories = e.Calories + cal
}

// Did you know a group of elves is called a grip? Who knew?
type GripOfElves []Elf

// Keep main() simple.
func (e GripOfElves) SumCalories() int {
	var result int
	for _, elf := range e {
		result = result + elf.Calories
	}
	return result
}

// ByCalories implements sort.Interface for a grip of Elves
// based on Calories of rations for each Elf.
type ByCalories GripOfElves

// Implementation for sorting by calorie count.
func (a ByCalories) Len() int           { return len(a) }
func (a ByCalories) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCalories) Less(i, j int) bool { return a[i].Calories < a[j].Calories }

// Result represents the product of parsing lines in inputFile.
type Result struct {
	Elf   Elf
	Error error
}

// Reads from r and sends each line through its output channel.
func generate(reader io.Reader) <-chan []string {
	output := make(chan []string)
	lines := make([]string, 0)
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			// Empty lines delimit inventories.
			if len(line) == 0 {
				output <- lines
				lines = make([]string, 0)
				continue
			}
			lines = append(lines, line)
		}
		close(output)
	}()
	return output
}

// Reads from upstream and sends on results.
func consume(done <-chan struct{}, upstream <-chan []string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)
		for lines := range upstream {
			currentElf := Elf{}
			// Sum the lines of this Elf's inventory.
			for _, line := range lines {
				// We want to keep integers of Calories, not strings.
				cal, err := strconv.Atoi(line)
				if err != nil {
					results <- Result{currentElf, err}
				}
				currentElf.AddToCalories(cal)
			}
			select {
			case results <- Result{currentElf, nil}:
			case <-done:
				return
			}
		}
	}()
	return results
}

// Merge all result channels into a single result channel
// to enable a variable pool of workers.
func merge(done <-chan struct{}, resultChans ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	out := make(chan Result)

	// Start an output goroutine for each input channel in resultChans.
	// output copies values from r to out until r is closed, then calls wg.Done.
	output := func(r <-chan Result) {
		defer wg.Done()
		for n := range r {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(resultChans))
	for _, r := range resultChans {
		go output(r)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	const maxWorkers = 5
	elves := make(GripOfElves, 0)

	// Open input file.
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Spawn the generator.
	output := generate(file)

	// Our indicator when we're done processing, likely due to an error.
	done := make(chan struct{})
	defer close(done)

	// Spawn consumers/workers.
	resultChannels := make([]<-chan Result, 0)
	for i := 0; i < maxWorkers; i++ {
		resultChannels = append(resultChannels, consume(done, output))
	}

	// Merge the Result channels from all consumers and build our
	// list of Elves.
	for result := range merge(done, resultChannels...) {
		if result.Error != nil {
			panic(fmt.Sprintf("Filure reported by a result: %v\n", result.Error))
		}
		elves = append(elves, result.Elf)
	}

	// Sort and reveal the most loaded Elf
	sort.Sort(ByCalories(elves))
	loadedElf := elves[len(elves)-1]
	fmt.Printf("// The loaded Elf has the most calories with %d\n", loadedElf.Calories)

	topThreeElves := elves[len(elves)-3:]
	fmt.Printf("// Top three Elves have %d calories between them.\n", topThreeElves.SumCalories())
}
