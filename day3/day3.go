package main

import (
	"fmt"
	"os"
	"strings"
)

func SplitRucksackCompartments(rucksackContents string) []string {
	return []string{rucksackContents[:len(rucksackContents)/2], rucksackContents[len(rucksackContents)/2:]}
}
func Priority(item int32) int {
	if item > 90 { // lowercase
		return int(item - 96)
	}
	return int(item - 38) // uppercase
}

func FindDuplicateItem(compartments []string) int32 {
	for _, item := range compartments[0] {
		if strings.Contains(compartments[1], string(item)) {
			return item
		}
	}
	return -1
}

func FindCommonItem(compartments []string) int32 {
    for _, item := range compartments[0] {
        if strings.Contains(compartments[1], string(item)) {
            if strings.Contains(compartments[2], string(item)) {
                return item
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
    fmt.Println("Advent of Code - Day 3\n===============-======\n\nPart One:")
	rucksacks := strings.Split(string(data), "\n")
    sumOfDuplicateItemPriorities := 0
	for rucksackIndex, rucksackContents := range rucksacks {
		rucksackCompartments := SplitRucksackCompartments(rucksackContents)
		duplicateItem := FindDuplicateItem(rucksackCompartments)
        itemPriority := Priority(duplicateItem)
        sumOfDuplicateItemPriorities += itemPriority
        fmt.Printf("Rucksack %d contains %s and has %c duplicated with a priority of %d\n", rucksackIndex+1, rucksackCompartments, duplicateItem, itemPriority)
	}
    fmt.Println("\nPart Two:")
    groupIndex := 0
    sumOfCommonItemPriorities := 0
    for rucksackIndex := 0; rucksackIndex < len(rucksacks); rucksackIndex+=3 {
        groupRucksacks := rucksacks[rucksackIndex:rucksackIndex+3]
        groupIndex+=1
        commonItem := FindCommonItem(groupRucksacks)
        itemPriority := Priority(commonItem)
        sumOfCommonItemPriorities += itemPriority
        fmt.Printf("Group %d: contains %v and has %c common with a priority of %d\n", groupIndex, groupRucksacks, commonItem, itemPriority)
    }
    fmt.Println("\nResults:")
    fmt.Printf("\tPart One:\n\t\tSum Of Item Priorities: %d\n", sumOfDuplicateItemPriorities)
    fmt.Printf("\tPart Two:\n\t\tSum Of Item Priorities: %d\n", sumOfCommonItemPriorities)
}
