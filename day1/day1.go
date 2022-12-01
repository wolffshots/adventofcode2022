package main

import (
	"fmt"     // console io
    "log"     // console io
	"os"      // file io
	"strconv" // parsing string to int
	"strings" // sub stringing and splitting
)

func main() {
	data, err := os.ReadFile("day1data.txt")
	if err != nil {
        fmt.Println(err)
	}
	elves := strings.Split(string(data), "\n\n")
	var max = []int{0, 0, 0}
	var maxElf = []int{0, 0, 0}
	for elfIndex, elf := range elves {
		cals := strings.Split(elf, "\n")
		current := 0
		// covert input calories and sum for the elf
		for _, cal := range cals {
			// convert string to int
			cal, err := strconv.Atoi(cal)
			if err != nil {
				log.Fatalf("str conv failed with: %v", err)
			}
			current = current + cal // running sum for elf
		}
        fmt.Printf("Total for elf #%d: %d\n", elfIndex+1, current)
		// check if we have a new max
		for i, val := range max {
			if current > val {
				// shift down current values
                if i < len(max)-2 {
                    max[i+2] = max[i+1]
                    maxElf[i+2] = maxElf[i+1]
                }
                if i < len(max)-1 {
                    max[i+1] = max[i]
                    maxElf[i+1] = maxElf[i]
                }
				// insert new value
				max[i] = current
				maxElf[i] = elfIndex
				// break to prevent re-insertion
				break
			}
		}

	}
    fmt.Println("Maxes:")
    topThree := 0
	for i, val := range max {
        fmt.Printf("Max[%d] for elf #%d: %d\n", i, maxElf[i]+1, val)
        topThree = topThree + val
	}

    fmt.Printf("Top three are carrying: %d\n", topThree)
//    2022/12/01 15:03:04 Max[0] for elf #150: 71471
//    2022/12/01 15:03:04 Max[1] for elf #100: 70523
//    2022/12/01 15:03:04 Max[2] for elf #50: 68194
}
