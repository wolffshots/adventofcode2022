package main

import (
	"github.com/stretchr/testify/assert" // boolean assertions for pass and fail
	"testing"                            // testing library
)

func TestFormatGrid(t *testing.T) {
	input := [][][]Tree{
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
			{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
		},
		{{}, {}, {}, {}, {}},
	}
	want := []string{
		"3 0 3 7 3\n2 5 5 1 2\n6 5 3 3 2\n3 3 5 4 9\n3 5 3 9 0",
		"\n\n\n\n",
	}

	for testIndex, testOutput := range want {
		result := FormatGrid(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestFormatGridVisibility(t *testing.T) {
	input := [][][]Tree{
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
			{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
		},
		{{}, {}, {}, {}, {}},
	}
	want := []string{
		"\t0        \t1        \t2        \t3        \t4        \n" +
			"0\t3 - true\t0 - true\t3 - true\t7 - true\t3 - true\n" +
			"1\t2 - true\t5 - true\t5 - true\t1 - false\t2 - true\n" +
			"2\t6 - true\t5 - true\t3 - false\t3 - true\t2 - true\n" +
			"3\t3 - true\t3 - false\t5 - true\t4 - false\t9 - true\n" +
			"4\t3 - true\t5 - true\t3 - true\t9 - true\t0 - true",
		"\n\n\n\n\n",
	}
	for testIndex, testOutput := range want {
		result := FormatGridVisibility(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestFormatGridViewing(t *testing.T) {
	input := [][][]Tree{
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
			{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
		},
		{{}, {}, {}, {}, {}},
	}
	want := []string{
		"\t0            1            2            3            4            \n" +
			"0  \t3   - 0      0   - 0      3   - 0      7   - 0      3   - 0      \n" +
			"1  \t2   - 0      5   - 1      5   - 4      1   - 1      2   - 0      \n" +
			"2  \t6   - 0      5   - 6      3   - 1      3   - 2      2   - 0      \n" +
			"3  \t3   - 0      3   - 1      5   - 8      4   - 3      9   - 0      \n" +
			"4  \t3   - 0      5   - 0      3   - 0      9   - 0      0   - 0      ",
		"\t\n\n\n\n\n",
	}
	for testIndex, testOutput := range want {
		result := FormatGridViewing(input[testIndex])
		assert.Equal(t, testOutput, result)
	}

}

func TestParseGrid(t *testing.T) {
	input := []string{
		"30373\n25512\n65332\n33549\n35390",
		"\n\n\n\n",
	}
	want := [][][]Tree{
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
			{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
		},
		{{}, {}, {}, {}, {}},
	}
	for testIndex, testOutput := range want {
		result := ParseGrid(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestIsVisibleHorizontal(t *testing.T) {
	trees := [][]Tree{
		{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
		{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
		{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
		{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
		{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
	}
	position := [][]int{
		{0, 0},
		{2, 2},
		{3, 1},
		{2, 3},
		{4, 4},
	}
	want := []bool{
		true,
		false,
		false,
		true,
		true,
	}
	for testIndex, testOutput := range want {
		result := IsVisibleHorizontal(trees, position[testIndex][0], position[testIndex][1])
		assert.Equal(t, testOutput, result)
	}
}

func TestIsVisibleVertical(t *testing.T) {
	trees := [][]Tree{
		{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
		{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
		{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
		{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
		{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
	}
	position := [][]int{
		{0, 0},
		{2, 2},
		{3, 1},
		{2, 3},
		{4, 4},
	}
	want := []bool{
		true,
		false,
		false,
		false,
		true,
	}
	for testIndex, testOutput := range want {
		result := IsVisibleVertical(trees, position[testIndex][0], position[testIndex][1])
		assert.Equal(t, testOutput, result)
	}
}

func TestIsVisible(t *testing.T) {
	trees := [][]Tree{
		{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
		{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
		{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
		{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
		{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
	}
	position := [][]int{
		{0, 0},
		{2, 2},
		{3, 1},
		{2, 3},
		{4, 4},
	}
	want := []bool{
		true,
		false,
		false,
		true,
		true,
	}
	for testIndex, testOutput := range want {
		result := IsVisible(trees, position[testIndex][0], position[testIndex][1])
		assert.Equal(t, testOutput, result)
	}
}

func TestScenicScores(t *testing.T) {
	trees := [][]Tree{
		{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
		{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
		{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
		{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
		{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
	}
	position := [][]int{
		{0, 0},
		{2, 1},
	}
	want := [][]int{
		{0, 2, 0, 2},
		{1, 3, 1, 2},
	}
	for testIndex, testOutput := range want {
		result := ScenicScores(trees, position[testIndex][0], position[testIndex][1])
		assert.Equal(t, testOutput, result)
	}
}

func TestCountVisible(t *testing.T) {
	trees := [][][]Tree{
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
			{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
		},
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 3}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 6}, {Height: 3}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 3}, {Height: 3}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 5}, {Height: 9}, {Height: 0}},
		},
	}
	want := []int{
		21,
		19,
	}
	for testIndex, testOutput := range want {
		result := CountVisible(trees[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestMaxScenicScore(t *testing.T) {
	trees := [][][]Tree{
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
			{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 3}, {Height: 9}, {Height: 0}},
		},
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
			{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 5}, {Height: 4}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 5}, {Height: 9}, {Height: 0}},
		},
		{
			{{Height: 3}, {Height: 0}, {Height: 3}, {Height: 7}, {Height: 3}},
			{{Height: 2}, {Height: 5}, {Height: 5}, {Height: 1}, {Height: 2}},
			{{Height: 6}, {Height: 5}, {Height: 3}, {Height: 3}, {Height: 2}},
			{{Height: 3}, {Height: 3}, {Height: 3}, {Height: 4}, {Height: 9}},
			{{Height: 3}, {Height: 5}, {Height: 5}, {Height: 9}, {Height: 0}},
		},
	}
	want := []int{
		8,
		8,
		9,
	}
	for testIndex, testOutput := range want {
		result := MaxScenicScore(trees[testIndex])
		assert.Equal(t, testOutput, result)
	}
}
