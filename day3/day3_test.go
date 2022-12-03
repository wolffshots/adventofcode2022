package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriority(t *testing.T) {
	input := []int32{'a', 'b', 'c', 'd', 'e', 'A', 'L', 'P', 'Z'}
	want := []int{1, 2, 3, 4, 5, 27, 38, 42, 52}
	for testIndex, testOutput := range want {
		result := Priority(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestSplitRucksackCompartments(t *testing.T) {
	input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}
	want := [][]string{{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, {"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}}
	for testIndex, testOutput := range want {
		result := SplitRucksackCompartments(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestFindDuplicateItem(t *testing.T) {
	input := [][]string{{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, {"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}, {"aaaaaaa", "bbbbbbb"}}
	want := []int32{'p', 'L', -1}
	for testIndex, testOutput := range want {
		result := FindDuplicateItem(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestFindCommonItem(t *testing.T) {
    input := [][]string{
        {"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
        {"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
        {"aaaaaaa", "bbbbbbb", "ccccccc"},
    }
    want := []int32{'r', 'Z', -1}
    for testIndex, testOutput := range want {
        result := FindCommonItem(input[testIndex])
        assert.Equal(t, testOutput, result)
    }
}
