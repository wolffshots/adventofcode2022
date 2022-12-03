package main

import (
	"github.com/stretchr/testify/assert" // boolean assertions for pass and fail
	"testing"                            // testing library
)

func TestCalculateScoreGivenInput(t *testing.T) {
	input := [][]string{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
	want := []int{8, 1, 6}
	for testIndex, testOutput := range want {
		result := CalculateScoreGivenInput(input[testIndex][0], input[testIndex][1])
		assert.Equal(t, testOutput, result)
	}
}

func TestCalculateScoreGivenOutcome(t *testing.T) {
	input := [][]string{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
	want := []int{4, 1, 7}
	for testIndex, testOutput := range want {
		result := CalculateScoreGivenOutcome(input[testIndex][0], input[testIndex][1])
		assert.Equal(t, testOutput, result)
	}
}
