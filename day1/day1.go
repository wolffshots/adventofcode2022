package main

import (
	"fmt"     // console io
	"log"     // console io
	"os"      // file io
	"strconv" // parsing string to int
	"strings" // sub stringing and splitting
)

var max = []int{0, 0, 0}
var maxElfIndex = []int{0, 0, 0}

func CalculateElfCalories(elf string) int {
	cals := strings.Split(elf, "\n")
	current := 0
	if cals[0] == "" {
		return 0
	}
	// covert input calories and sum for the elf

	for _, cal := range cals {
		// convert string to int
		cal, err := strconv.Atoi(cal)
		if err != nil {
			log.Fatalf("str conv failed with: %v", err)
		}
		current = current + cal // running sum for elf
	}
	return current
}

func InsertNewMax(elfIndex, newValue int) {
	for maxIndex, currentMax := range max {
		if newValue > currentMax {
			// shift down current values
			if maxIndex < len(max)-2 {
				max[maxIndex+2] = max[maxIndex+1]
				maxElfIndex[maxIndex+2] = maxElfIndex[maxIndex+1]
			}
			if maxIndex < len(max)-1 {
				max[maxIndex+1] = max[maxIndex]
				maxElfIndex[maxIndex+1] = maxElfIndex[maxIndex]
			}
			// insert new value
			max[maxIndex] = newValue
			maxElfIndex[maxIndex] = elfIndex
			// break to prevent re-insertion
			break
        }
	}
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	elves := strings.Split(string(data), "\n\n")

	for elfIndex, elf := range elves {
		newValue := CalculateElfCalories(elf)
		fmt.Printf("Total for elf #%d: %d\n", elfIndex+1, newValue)
		// check if we have a new max
		InsertNewMax(elfIndex, newValue)
	}
    
	fmt.Println("\nMaxes:")
	topThree := 0
	for i, val := range max {
		fmt.Printf("\tMax[%d] for elf #%d: %d\n", i, maxElfIndex[i]+1, val)
		topThree = topThree + val
	}

	fmt.Printf("\n\tTop three are carrying: %d\n", topThree)
	// Maxes:
	// Max[0] for elf #150: 71471
	// Max[1] for elf #100: 70523
	// Max[2] for elf #119: 69195
	// Top three are carrying: 211189
}
