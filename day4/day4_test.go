package main

import (
	"github.com/stretchr/testify/assert" // boolean assertions for pass and fail
	"testing"                            // testing library
)

func TestRangeContains(t *testing.T) {
	input := [][]string{{"2-4", "6-8"}, {"2-3", "4-5"}, {"5-7", "7-9"}, {"2-8", "3-7"}}
	want := [][]int{{0, 0}, {0, 0}, {1, 1}, {2, 1}}
	for testIndex, testOutput := range want {
		result := RangeContains(input[testIndex][0], input[testIndex][1])
		assert.Equal(t, testOutput[0], result)
		result = RangeContains(input[testIndex][1], input[testIndex][0])
		assert.Equal(t, testOutput[1], result)
	}
}
