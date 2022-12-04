package main

import (
	"fmt"     // console io
	"os"      // file io
	"strconv" // string parsing
	"strings" // string formatting
)

// RangeContains takes string and substr as input and parses them to numbers and checks if the substr
// range is contained in the string range and returns a 2 if so, otherwise a 1 for partial overlap
// or a 0 for no overlap
func RangeContains(string string, substr string) int {
	rangeOne := strings.Split(string, "-")
	rangeOneStart, _ := strconv.Atoi(rangeOne[0])
	rangeOneEnd, _ := strconv.Atoi(rangeOne[1])
	rangeTwo := strings.Split(substr, "-")
	rangeTwoStart, _ := strconv.Atoi(rangeTwo[0])
	rangeTwoEnd, _ := strconv.Atoi(rangeTwo[1])

	if (rangeOneStart <= rangeTwoStart) && (rangeOneEnd >= rangeTwoEnd) { // one fully contains two
		return 2
	} else if rangeOneEnd < rangeTwoStart { // exclusive ranges
		return 0
	} else if rangeTwoEnd < rangeOneStart { // exclusive ranges
		return 0
	}
	return 1 // can't be exclusive therefore is inclusive
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 4\n===============-======\n\nPart One:")

	pairs := strings.Split(string(data), "\n")
	assignmentsWithRangeContaining := 0
	assignmentsWithRangeOverlapping := 0
	for pairIndex, pairValue := range pairs {
		ranges := strings.Split(pairValue, ",")

		fmt.Printf("Pair %d\t\tRange 1: %s \t\tRange 2: %s \n",
			pairIndex,
			ranges[0],
			ranges[1],
		)
		rangeOneContainsRangeTwo := RangeContains(ranges[0], ranges[1])
		rangeTwoContainsRangeOne := RangeContains(ranges[1], ranges[0])
		if rangeOneContainsRangeTwo == 2 || rangeTwoContainsRangeOne == 2 {
			assignmentsWithRangeContaining += 1
			assignmentsWithRangeOverlapping += 1
		} else if rangeOneContainsRangeTwo == 1 || rangeTwoContainsRangeOne == 1 {
			assignmentsWithRangeOverlapping += 1
		}
	}

	fmt.Println("\nPart Two:")

	fmt.Println("\nResults:")
	fmt.Printf("Assignments with one range fully containing the other: %d\n", assignmentsWithRangeContaining) // 475
	fmt.Printf("Assignments with one range overlapping the other: %d\n", assignmentsWithRangeOverlapping)     //
}
