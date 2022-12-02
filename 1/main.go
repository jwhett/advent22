package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	inputFile = "input"
)

type Elf struct {
	Id       int
	Calories int
}

func (e *Elf) AddToCalories(cal int) {
	e.Calories = e.Calories + cal
}

// ByCalories implements sort.Interface for []Elf
// based on Calories of rations for each Elf.
type ByCalories []Elf

func (a ByCalories) Len() int           { return len(a) }
func (a ByCalories) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCalories) Less(i, j int) bool { return a[i].Calories < a[j].Calories }

func main() {
	elves := make([]Elf, 0)

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	currentId := 1                  // our IDs will start at 1 instead of 0
	currentElf := Elf{currentId, 0} // initialize our first Elf

	// for each line in input file...
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCalories := scanner.Text()
		if len(lineCalories) == 0 {
			// empty line delimits Elf inventories.
			// store current Elf and start on a new inventory.
			elves = append(elves, currentElf)
			currentId = currentId + 1
			currentElf = Elf{currentId, 0}
			continue
		}
		// we want to keep integers of Calories
		cal, err := strconv.Atoi(lineCalories)
		if err != nil {
			fmt.Println(err)
		}
		currentElf.AddToCalories(cal)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// sort and reveal the most loaded Elf
	sort.Sort(ByCalories(elves))
	loadedElf := elves[len(elves)-1]
	fmt.Printf("// Elf ID %d has the most calories with %d\n", loadedElf.Id, loadedElf.Calories)
}
