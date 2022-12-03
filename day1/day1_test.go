package main

import (
	"github.com/stretchr/testify/assert" // boolean assertions for pass and fail
	"testing"                            // testing library
)

func TestCalculateElfCalories(t *testing.T) {
	input := []string{
		"8417\n8501\n5429\n2112\n6482\n7971\n9636\n4003", "",
	}
	want := []int{52551, 0}
	for testIndex, testOutput := range want {
		result := CalculateElfCalories(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestInsertNewMax(t *testing.T) {
	input := []int{
		0,
		50624,
		50623,
		80624,
		40624,
		50624,
		90655,
	}

	want := [][]int{
		{0, 0, 0},
		{50624, 0, 0},
		{50624, 50623, 0},
		{80624, 50624, 50623},
		{80624, 50624, 50623},
		{80624, 50624, 50624},
		{90655, 80624, 50624},
	}

	for testIndex, testOutput := range want {
		InsertNewMax(testIndex, input[testIndex])
		assert.Equal(t, testOutput, max)
	}
}
