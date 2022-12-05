package main

import (
	"github.com/stretchr/testify/assert" // boolean assertions for pass and fail
	"testing"                            // testing library
)

func TestPrintBoard(t *testing.T) {
	input := [][][]string{
		{
			{"[Z]", "[N]"},
			{"[M]", "[C]", "[D]"},
			{"[P]"},
		},
	}
	want := []string{
		"  e   [D]   e  \n [N]  [C]   e  \n [Z]  [M]  [P] \n  1    2    3  ",
	}
	for testIndex, testOutput := range want {
		result := PrintBoard(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestParseBoard(t *testing.T) {
	input := []string{
		"    [D]\n[N] [C]\n[Z] [M] [P]\n1   2   3",
	}
	want := [][][]string{
		{
			{"[Z]", "[N]"},
			{"[M]", "[C]", "[D]"},
			{"[P]"},
		},
	}

	for testIndex, testOutput := range want {
		result := ParseBoard(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestMoveSingular(t *testing.T) {
	boards := [][][]string{
		{
			{"[Z]", "[N]"},
			{"[M]", "[C]", "[D]"},
			{"[P]"},
		},
		{
			{"[Z]", "[N]", "[D]"},
			{"[M]", "[C]"},
			{"[P]"},
		},
		{
			{},
			{"[M]", "[C]"},
			{"[P]", "[D]", "[N]", "[Z]"},
		},
		{
			{"[C]", "[M]"},
			{},
			{"[P]", "[D]", "[N]", "[Z]"},
		},
	}
	moves := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	want := [][][]string{
		{
			{"[Z]", "[N]", "[D]"},
			{"[M]", "[C]"},
			{"[P]"},
		},
		{
			{},
			{"[M]", "[C]"},
			{"[P]", "[D]", "[N]", "[Z]"},
		},
		{
			{"[C]", "[M]"},
			{},
			{"[P]", "[D]", "[N]", "[Z]"},
		},
		{
			{"[C]"},
			{"[M]"},
			{"[P]", "[D]", "[N]", "[Z]"},
		},
	}

	for testIndex, testOutput := range want {
		result := Move(boards[testIndex], moves[testIndex], true)
		assert.Equal(t, testOutput, result)
	}
}

func TestMoveMultiple(t *testing.T) {
	boards := [][][]string{
		{
			{"[Z]", "[N]"},
			{"[M]", "[C]", "[D]"},
			{"[P]"},
		},
		{
			{"[Z]", "[N]", "[D]"},
			{"[M]", "[C]"},
			{"[P]"},
		},
		{
			{},
			{"[M]", "[C]"},
			{"[P]", "[Z]", "[N]", "[D]"},
		},
		{
			{"[M]", "[C]"},
			{},
			{"[P]", "[Z]", "[N]", "[D]"},
		},
	}
	moves := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	want := [][][]string{
		{
			{"[Z]", "[N]", "[D]"},
			{"[M]", "[C]"},
			{"[P]"},
		},
		{
			{},
			{"[M]", "[C]"},
			{"[P]", "[Z]", "[N]", "[D]"},
		},
		{
			{"[M]", "[C]"},
			{},
			{"[P]", "[Z]", "[N]", "[D]"},
		},
		{
			{"[M]"},
			{"[C]"},
			{"[P]", "[Z]", "[N]", "[D]"},
		},
	}

	for testIndex, testOutput := range want {
		result := Move(boards[testIndex], moves[testIndex], false)
		assert.Equal(t, testOutput, result)
	}
}

func TestTopOfBoard(t *testing.T) {
	boards := [][][]string{
		{
			{"[C]"},
			{"[M]"},
			{"[P]", "[D]", "[N]", "[Z]"},
		},
	}
	want := []string{
		"CMZ",
	}

	for testIndex, testOutput := range want {
		result := TopOfBoard(boards[testIndex])
		assert.Equal(t, testOutput, result)
	}
}
