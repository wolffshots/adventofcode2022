package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var verbose bool

func Process(input string, register int) int {
	tempRegister := register
	instruction := strings.Split(input, " ")
	switch instruction[0] {
	case "noop":
		{
			return tempRegister
		}
	case "addx":
		{
			value, err := strconv.ParseInt(instruction[1], 0, 0)
			if err != nil {
				log.Fatalf("Error converting value to int for '%s': %v", input, err)
			}
			return tempRegister + int(value)
		}
	default:
		return tempRegister
	}
}

var Duration = map[string]int{
	"addx": 2,
	"noop": 1,
}

func RunInstructions(input []string) (int, string, int) {
	cycle := 1
	register := 1
	screen := "\t    "
	sumOfStrengths := 0
	for pixelIndex := 0; pixelIndex < 40; pixelIndex++ {
		screen = fmt.Sprintf("%s %3d", screen, pixelIndex)
	}
	screen = fmt.Sprintf("%s %s", screen, "\n\t")
	for _, instruction := range input {
		operation := strings.Split(instruction, " ")
		for instructionCycleIndex := 0; instructionCycleIndex < Duration[operation[0]]; instructionCycleIndex++ {
			pixel := " "
			if register == ((cycle-1)%40)-1 {
				pixel = "L"
			} else if register == ((cycle - 1) % 40) {
				pixel = "#"
			} else if register == ((cycle-1)%40)+1 {
				pixel = "R"
			}

			if instructionCycleIndex == Duration[operation[0]]-1 {
				register = Process(instruction, register)
			}

			if cycle%40 == 0 {
				screen = fmt.Sprintf("%s %3s %3d %s", screen, pixel, cycle, "\n\t")
			} else if cycle%40 == 1 {
				screen = fmt.Sprintf("%s %3d %3s", screen, cycle, pixel)
			} else {
				screen = fmt.Sprintf("%s %3s", screen, pixel)
			}
			cycle++
			if cycle == 20 || (cycle-20)%40 == 0 {
				if verbose {
					fmt.Printf("\t%-6d\t%-8d\t%-7d\n", cycle, register, register*cycle)
				}
				sumOfStrengths += register * cycle
			}
		}
	}
	screen = fmt.Sprintf("%s    ", screen)
	for pixelIndex := 0; pixelIndex < 40; pixelIndex++ {
		screen = fmt.Sprintf("%s %3d", screen, pixelIndex)
	}
	return register, screen, sumOfStrengths
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Sets the output to verbose")
	flag.Parse()
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 10\n===============-=======")
	fmt.Println()

	input := strings.Split(string(data), "\n")
	if verbose {
		fmt.Println("\tcycle\tregister\tstrength")
	}
	register, screen, sumOfStrengths := RunInstructions(input)

	fmt.Println("\nResults")
	fmt.Printf("\tFinal register value:    %15d\n", register)       // 9
	fmt.Printf("\tSum of signal strengths: %15d\n", sumOfStrengths) // 11960
	fmt.Printf("\tScreen:\n%s\n", screen)                           // EJCFPGLH

}
