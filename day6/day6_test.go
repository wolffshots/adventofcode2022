package main

import (
	"github.com/stretchr/testify/assert" // boolean assertions for pass and fail
	"testing"                            // testing library
)

//"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
//"bvwbjplbgvbhsrlpgdmjqwftvncz",
//"nppdvjthqldpwncqszvftbrmjlhg",
//"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
//"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",

func TestIsMarker(t *testing.T) {
	input := []string{
		"mjqj",
		"jqjp",
		"qjpq",
		"jpqm",
	}
	want := []bool{
		false,
		false,
		false,
		true,
	}
	for testIndex, testOutput := range want {
		result := IsMarker(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestFindMarkerPositionFour(t *testing.T) {
	input := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
	want := []int{
		7,
		5,
		6,
		10,
		11,
	}
	for testIndex, testOutput := range want {
		result := FindMarkerPosition(input[testIndex], 4)
		assert.Equal(t, testOutput, result)
	}
}

func TestFindMarkerPositionFourteen(t *testing.T) {
	input := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
	want := []int{
		19,
		23,
		23,
		29,
		26,
	}
	for testIndex, testOutput := range want {
		result := FindMarkerPosition(input[testIndex], 14)
		assert.Equal(t, testOutput, result)
	}
}
