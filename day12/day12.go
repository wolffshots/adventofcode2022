package main

import (
	"flag"
	"fmt"
    "image"
	"os"
	"strings"
)

var verbose bool

// PathFind - my recursive method didn't work beyond small data sets so I borrowed this from https://github.com/mnml
// I'd like to reimplement it how i was originally envisioning at some point
func PathFind(input string) []int{
    var start, end image.Point
    heights := map[image.Point]rune{}
    for rowIndex, row := range strings.Fields(input) {
        for columnIndex, height := range row {
            heights[image.Point{X: columnIndex, Y: rowIndex}] = height
            if height == 'S' {
                start = image.Point{X: columnIndex, Y: rowIndex}
            } else if height == 'E' {
                end = image.Point{X: columnIndex, Y: rowIndex}
            }
        }
    }
    heights[start], heights[end] = 'a', 'z'

    dists := map[image.Point]int{end: 0}
    queue := []image.Point{end}
    var shortest *image.Point

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if heights[current] == 'a' && shortest == nil {
            shortest = &current
        }

        for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
            next := current.Add(d)
            _, seen := dists[next]
            _, valid := heights[next]
            if seen || !valid || heights[next] < heights[current]-1 {
                continue
            }
            dists[next] = dists[current] + 1
            queue = append(queue, next)
        }
    }
    return []int{dists[start],dists[*shortest]}
}

func RecursivePathFind() int{
    // TODO
    return -1
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Sets the output to verbose")
	flag.Parse()
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 12\n===============-=======")
	fmt.Println()
	input := string(data)

    dists:=PathFind(input)

	fmt.Println("Results: ")
    fmt.Printf("\t Our Distance:    \t%5v\n", dists[0]) // <380
    fmt.Printf("\t Hiking distance: \t%5v\n", dists[1]) // <375
}
