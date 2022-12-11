package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var verbose bool

func Operate(input string, old uint64) uint64 {
	var sides []string
	newVal := uint64(0)
	if strings.Contains(input, "+") {
		sides = strings.Split(input, " + ")
	} else if strings.Contains(input, "-") {
		sides = strings.Split(input, " - ")
	} else if strings.Contains(input, "*") {
		sides = strings.Split(input, " * ")
	} else if strings.Contains(input, "/") {
		sides = strings.Split(input, " / ")
	} else {
		log.Fatalf("Couldn't eval")
	}
	lhs, err := strconv.ParseInt(sides[0], 0, 0)
	if err != nil && sides[0] == "old" {
		lhs = int64(old)
	}
	rhs, err := strconv.ParseInt(sides[1], 0, 0)
	if err != nil && sides[1] == "old" {
		rhs = int64(old)
	}
	if strings.Contains(input, "+") {
		newVal = uint64(lhs + rhs)
	} else if strings.Contains(input, "-") {
		newVal = uint64(lhs - rhs)
	} else if strings.Contains(input, "*") {
		newVal = uint64(lhs * rhs)
	} else if strings.Contains(input, "/") {
		newVal = uint64(lhs / rhs)
	} else {
		log.Fatalf("Couldn't eval")
	}
	return newVal
}

type Monkey struct {
	Items            []string
	Operation        string
	Test             string
	TestPass         string
	TestFail         string
	InspectedCounter int
}

type Num interface {
	~int | ~int64 | ~float64
}

func IsDivisibleBy(by uint64, old uint64) bool {
	return (old % by) == 0
}

func Inspect(operation string, old uint64) uint64 {
	return Operate(operation, old)
}

func Relief(old uint64) uint64 {
	return uint64(math.Floor(float64(old / 3)))
}

func Test(test string, old uint64) bool {
	by, err := strconv.ParseInt(test, 0, 0)
	if err != nil {
		log.Fatalf("Error converting monkey test to int: %v", err)
	}
	return IsDivisibleBy(uint64(by), old)
}

func Throw(monkeys []Monkey, originMonkeyIndex uint64, oldItem string, newItem uint64, destinationMonkey string) []Monkey {
	if verbose {
		fmt.Printf("Removing %s from monkey %d and giving it to %s as %d\n", oldItem, originMonkeyIndex, destinationMonkey, newItem)
	}
	tempMonkeys := monkeys
	destinationMonkeyIndex, err := strconv.ParseInt(destinationMonkey, 0, 0)
	if err != nil {
		log.Fatalf("Error converting destination monkey index to int: %v", err)
	}
	tempMonkeys[originMonkeyIndex].Items = tempMonkeys[originMonkeyIndex].Items[1:]
	tempMonkeys[destinationMonkeyIndex].Items = append(tempMonkeys[destinationMonkeyIndex].Items, fmt.Sprint(newItem))
	return tempMonkeys
}

func Load(input string) []Monkey {
	monkeyInputs := strings.Split(input, "\n\n")
	var monkeys []Monkey
	for monkeyIndex, monkeyInput := range monkeyInputs {
		if verbose {
			fmt.Println(monkeyInput)
		}
		monkeyInputIndex, err := strconv.ParseInt(fmt.Sprintf("%c", strings.Split(strings.Split(monkeyInput, "\n")[0], " ")[1][0]), 0, 0)
		if err != nil {
			log.Fatalf("Error converting monkey number to int: %v", err)
		} else if monkeyInputIndex != int64(monkeyIndex) {
			log.Fatalf("Input does not match expected: %d != %d", monkeyInputIndex, monkeyIndex)
		}
		startingItems := strings.Split(strings.Trim(strings.Split(strings.Split(monkeyInput, "\n")[1], ": ")[1], " "), ", ")
		operation := strings.Split(strings.Trim(strings.Split(strings.Split(monkeyInput, "\n")[2], ":")[1], " "), " = ")[1]
		test := strings.Split(strings.Trim(strings.Split(strings.Split(monkeyInput, "\n")[3], ":")[1], " "), "divisible by ")[1]
		testPass := strings.Split(strings.Trim(strings.Split(strings.Split(monkeyInput, "\n")[4], ":")[1], " "), "throw to monkey ")[1]
		testFail := strings.Split(strings.Trim(strings.Split(strings.Split(monkeyInput, "\n")[5], ":")[1], " "), "throw to monkey ")[1]
		monkeys = append(monkeys, Monkey{startingItems, operation, test, testPass, testFail, 0})
		if verbose {
			fmt.Printf("\n%d\n\t", monkeyIndex)
			fmt.Printf("starting: %s\n\t", startingItems)
			fmt.Printf("new = %s\n\t", operation)
			fmt.Printf("if divisible by %s ", test)
			fmt.Printf("-> %s\n\t", testPass)
			fmt.Printf("else -> %s\n", testFail)
		}
	}
	return monkeys
}

func FormatMonkeys(monkeys []Monkey) string {
	result := ""
	for monkeyIndex, monkey := range monkeys {
		result = fmt.Sprintf("%sMonkey %d:", result, monkeyIndex)
		for itemIndex, item := range monkey.Items {
			if itemIndex == 0 {
				result = fmt.Sprintf("%s %s", result, item)
			} else {
				result = fmt.Sprintf("%s, %s", result, item)
			}
		}
		result = fmt.Sprintf("%s\n", result)
	}
	return result
}

func Round(monkeys []Monkey, reliefEnsues bool) []Monkey {
	tempMonkeys := monkeys
	var tests []int
	for _, m := range monkeys {
		testInput, err := strconv.ParseInt(m.Test, 0, 0)
		if err != nil {
			log.Fatalf("Error parsing monkey test %s to int: %v", m.Test, err)
		}
		tests = append(tests, int(testInput))
	}
	lowestCommonMultiple := 1
	for _, test := range tests {
		lowestCommonMultiple = (lowestCommonMultiple * test) / GreatestCommonDivisor(lowestCommonMultiple, test)
	}
	for monkeyIndex, monkey := range tempMonkeys {
		for _, item := range monkey.Items {
			old64, err := strconv.ParseInt(item, 0, 0)
			if err != nil {
				log.Fatalf("Error parsing item %s to int: %v", item, err)
			}
			if verbose {
				fmt.Printf("Monkey inspects an item with a worry level of %d\n", old64)
			}
			newItem := Inspect(monkey.Operation, uint64(old64))
			tempMonkeys[monkeyIndex].InspectedCounter = tempMonkeys[monkeyIndex].InspectedCounter + 1
			if verbose {
				fmt.Printf("Worry level is changed to %d\n", newItem)
			}
			if reliefEnsues {
				newItem = Relief(newItem)
			} else {
				newItem = newItem % uint64(lowestCommonMultiple)
			}
			if verbose {
				fmt.Printf("Monkey gets bored with item. Worry level is divided by 3 to %d\n", newItem)
			}
			if Test(monkey.Test, newItem) {
				if verbose {
					fmt.Printf("Current worry level is divisible by %s\n", monkey.Test)
				}
				tempMonkeys = Throw(tempMonkeys, uint64(monkeyIndex), item, newItem, monkey.TestPass)
			} else {
				if verbose {
					fmt.Printf("Current worry level is not divisible by %s\n", monkey.Test)
				}
				tempMonkeys = Throw(tempMonkeys, uint64(monkeyIndex), item, newItem, monkey.TestFail)
			}
		}
	}
	return tempMonkeys
}

func Absolute[T Num](input T) T {
	if input < 0 {
		return -input
	}
	return input
}

func GreatestCommonDivisor(lcm int, div int) int {
    lcm, div = Absolute(lcm), Absolute(div)
    for div != 0 {
        lcm, div = div, lcm%div
	}
    return lcm
}

func FormatMonkeysInspectedCounters(monkeys []Monkey) string {
	result := ""
	for monkeyIndex, monkey := range monkeys {
		result = fmt.Sprintf("%sMonkey %d inspected items %d times.\n", result, monkeyIndex, monkey.InspectedCounter)
	}
	return result
}

func Max(one int, two int) int {
	if one > two {
		return one
	}
	return two
}

func MonkeyBusiness(monkeys []Monkey) int {
	biggest := 0
	secondBiggest := 0
	for _, monkey := range monkeys {
		biggest = Max(biggest, monkey.InspectedCounter)
	}
	for _, monkey := range monkeys {
		if monkey.InspectedCounter != biggest {
			secondBiggest = Max(secondBiggest, monkey.InspectedCounter)
		}
	}
	return biggest * secondBiggest
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Sets the output to verbose")
	flag.Parse()
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 11\n===============-=======")
	fmt.Println()

	input := string(data)

	monkeys := Load(input)
	if verbose {
		fmt.Println("Part 1")
	}
	for roundIndex := 0; roundIndex < 20; roundIndex++ {
		monkeys = Round(monkeys, true)
		if verbose {
			fmt.Printf("Round %d\n", roundIndex+1)
			fmt.Println(FormatMonkeysInspectedCounters(monkeys))
			fmt.Println(FormatMonkeys(monkeys))
		}
	}
	monkeyBusiness1 := MonkeyBusiness(monkeys)
	if verbose {
		fmt.Println("Part 2")
	}
	monkeys = Load(input)
	for roundIndex := 0; roundIndex < 10000; roundIndex++ {
		monkeys = Round(monkeys, false)
		if (((roundIndex+1)%1000) == 0 || (roundIndex+1) == 20) && verbose {
			fmt.Printf("Round %d\n", roundIndex+1)
			fmt.Println(FormatMonkeysInspectedCounters(monkeys))
			fmt.Println(FormatMonkeys(monkeys))
		}
	}

	monkeyBusiness2 := MonkeyBusiness(monkeys)

	fmt.Printf("\tMonkey business 1: %d\n", monkeyBusiness1)   // = 62491
	fmt.Printf("\tMonkey business 2: %d\n\n", monkeyBusiness2) // = 17408399184
}
