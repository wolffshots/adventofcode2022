package main

import (
	"fmt"     // console io
	"os"      // file io
	"strings" // string manipulation
)

func IsMarker(input string) bool {
	for _, char := range input {
		if strings.Count(input, string(char)) > 1 {
			return false
		}
	}
	return true
}

func FindMarkerPosition(input string, uniques int) int {
	for charIndex := range input {
		if charIndex+uniques <= len(input) {
			if IsMarker(input[charIndex : charIndex+uniques]) {
				return charIndex + uniques
			}
		}
	}
	return -1
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 5\n===============-======")
	fmt.Println()

	input := strings.Split(string(data), "\n")
	var results [][]int
	for range input {
		results = append(results, []int{})
	}
	for lineIndex, line := range input {
		markerPosition4 := FindMarkerPosition(line, 4)
		markerPosition14 := FindMarkerPosition(line, 14)
		results[lineIndex] = []int{markerPosition4, markerPosition14}
		if len(line) < 60 {
			fmt.Println("\nPart One:")
			fmt.Printf("Input #%d: \n%s*%s\n%s*%s\nMarker at %d\n\n",
				lineIndex,
				line[:markerPosition4],
				line[markerPosition4:],
				strings.Repeat("----=", len(line)+1/5)[:len(line)+1][:markerPosition4],
				strings.Repeat("----=", len(line)+1/5)[:len(line)+1][markerPosition4:],
				markerPosition4,
			)
			fmt.Println("Part Two:")
			fmt.Printf("Input #%d: \n%s*%s\n%s*%s\nMarker at %d\n\n",
				lineIndex,
				line[:markerPosition14],
				line[markerPosition14:],
				strings.Repeat("----=", len(line)+1/5)[:len(line)+1][:markerPosition14],
				strings.Repeat("----=", len(line)+1/5)[:len(line)+1][markerPosition14:],
				markerPosition14,
			)
		}
	}

	fmt.Println("\nResults:")
	for resultIndex, result := range results {
		fmt.Printf("\tResult #%d \t- 4: \t%d, \t14: \t%d", resultIndex, result[0], result[1])
		if len(input[resultIndex]) < 60 {
			fmt.Printf(" - %s\n", input[resultIndex])
		} else {
			fmt.Println()
		}

	}
}
