package main

import (
	"github.com/stretchr/testify/assert" // boolean assertions for pass and fail
	"testing"                            // testing library
)

func TestIsCommand(t *testing.T) {
	input := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}
	want := []bool{
		true,
		true,
		false,
		false,
		false,
		false,
		true,
		true,
		false,
		false,
		false,
		false,
		true,
		true,
		false,
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		false,
	}
	for testIndex, testOutput := range want {
		result := IsCommand(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestList(t *testing.T) {
	input := []map[string]Node{
		{"/": {Name: "/", Path: "", Size: 1111, Type: "d"}},
	}
	want := []string{
		"d           1111\t                   / - \n",
	}
	for testIndex, testOutput := range want {
		result := List(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}

func TestSmallerThanByType(t *testing.T) {
	input := []map[string]Node{
		{"/": {Name: "/", Path: "", Size: 2000, Type: "d"},
			"/other":   {Name: "other", Path: "/", Size: 1500, Type: "d"},
			"/another": {Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
		{"/": {Name: "/", Path: "", Size: 2000, Type: "d"},
			"/other":   {Name: "other", Path: "/", Size: 1500, Type: "d"},
			"/another": {Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
		{"/": {Name: "/", Path: "", Size: 2000, Type: "d"},
			"/other":   {Name: "other", Path: "/", Size: 1500, Type: "d"},
			"/another": {Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
	}
	thresholds := []int{
		1700,
		1300,
		500,
	}
	want := [][]Node{
		{
			{Name: "other", Path: "/", Size: 1500, Type: "d"},
			{Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
		{
			{Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
		nil,
	}
	for testIndex, testOutput := range want {
		result := SmallerThanByType(input[testIndex], thresholds[testIndex], "d")
		assert.Equal(t, testOutput, result)
	}
}

func TestNearestGreatestByType(t *testing.T) {
	input := []map[string]Node{
		{"/": {Name: "/", Path: "", Size: 2000, Type: "d"},
			"/other":   {Name: "other", Path: "/", Size: 1500, Type: "d"},
			"/another": {Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
		{"/": {Name: "/", Path: "", Size: 2000, Type: "d"},
			"/other":   {Name: "other", Path: "/", Size: 1500, Type: "d"},
			"/another": {Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
		{"/": {Name: "/", Path: "", Size: 2000, Type: "d"},
			"/other":   {Name: "other", Path: "/", Size: 1500, Type: "d"},
			"/another": {Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
		{"/": {Name: "/", Path: "", Size: 2000, Type: "d"},
			"/other":   {Name: "other", Path: "/", Size: 1500, Type: "d"},
			"/another": {Name: "another", Path: "/other", Size: 1000, Type: "d"},
		},
	}
	thresholds := []int{
		1700,
		1300,
		500,
		2500,
	}
	want := []Node{
		{Name: "/", Path: "", Size: 2000, Type: "d"},
		{Name: "other", Path: "/", Size: 1500, Type: "d"},
		{Name: "another", Path: "/other", Size: 1000, Type: "d"},
		{Name: "", Path: "", Size: 0, Type: ""},
	}
	for testIndex, testOutput := range want {
		result := NearestGreatestByType(input[testIndex], thresholds[testIndex], "d")
		assert.Equal(t, testOutput, result)
	}
}

func TestRun(t *testing.T) {
	input := [][]string{
		{
			"$ cd /",
			"$ ls",
			"dir a",
			"14848514 b.txt",
			"8504156 c.dat",
			"dir d",
			"$ cd a",
			"$ ls",
			"dir e",
			"29116 f",
			"2557 g",
			"62596 h.lst",
			"$ cd e",
			"$ ls",
			"584 i",
			"$ cd ..",
			"$ cd ..",
			"$ cd d",
			"$ ls",
			"4060174 j",
			"8033020 d.log",
			"5626152 d.ext",
			"7214296 k",
		},
	}
	want := []map[string]Node{
		{"/": {Name: "/", Path: "", Size: 48381165, Type: "d"},
			"//a":       {Name: "a", Path: "/", Size: 94853, Type: "d"},
			"//a/e":     {Name: "e", Path: "//a", Size: 584, Type: "d"},
			"//a/e/i":   {Name: "i", Path: "//a/e", Size: 584, Type: "-"},
			"//a/f":     {Name: "f", Path: "//a", Size: 29116, Type: "-"},
			"//a/g":     {Name: "g", Path: "//a", Size: 2557, Type: "-"},
			"//a/h.lst": {Name: "h.lst", Path: "//a", Size: 62596, Type: "-"},
			"//b.txt":   {Name: "b.txt", Path: "/", Size: 14848514, Type: "-"},
			"//c.dat":   {Name: "c.dat", Path: "/", Size: 8504156, Type: "-"},
			"//d":       {Name: "d", Path: "/", Size: 24933642, Type: "d"},
			"//d/d.ext": {Name: "d.ext", Path: "//d", Size: 5626152, Type: "-"},
			"//d/d.log": {Name: "d.log", Path: "//d", Size: 8033020, Type: "-"},
			"//d/j":     {Name: "j", Path: "//d", Size: 4060174, Type: "-"},
			"//d/k":     {Name: "k", Path: "//d", Size: 7214296, Type: "-"},
		},
	}
	for testIndex, testOutput := range want {
		result := Run(input[testIndex])
		assert.Equal(t, testOutput, result)
	}
}
