package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var verbose bool
var showVisited bool

type Position struct {
	X int
	Y int
}

func Move(current Position, move string) Position {
	tempPosition := current
	if move == "" {

		return tempPosition
	}
	directions := strings.Split(move, " ")[0]
	distance, err := strconv.ParseInt(strings.Split(move, " ")[1], 0, 0)
	if err != nil {
		log.Fatalf("Couldn't process distance: %v", err)
	}

	for _, direction := range directions {
		switch direction {
		case 'R':
			tempPosition.X = tempPosition.X + int(distance)
		case 'L':
			tempPosition.X = tempPosition.X - int(distance)
		case 'U':
			tempPosition.Y = tempPosition.Y + int(distance)
		case 'D':
			tempPosition.Y = tempPosition.Y - int(distance)
		}
	}
	return tempPosition
}

func StraightMoves(tail Position, head Position) []string {
	xDistance := head.X - tail.X
	yDistance := head.Y - tail.Y
	yChange := int(math.Abs(float64(yDistance))) - 1
	xChange := int(math.Abs(float64(xDistance))) - 1
	var moves []string
	if head.Y > tail.Y+1 && head.X == tail.X {
		for i := 0; i < yChange; i++ {
			moves = append(moves, "U 1")
		}
	} else if head.Y < tail.Y-1 && head.X == tail.X {
		for i := 0; i < yChange; i++ {
			moves = append(moves, "D 1")
		}
	} else if head.X > tail.X+1 && head.Y == tail.Y {
		for i := 0; i < xChange; i++ {
			moves = append(moves, "R 1")
		}
	} else if head.X < tail.X-1 && head.Y == tail.Y {
		for i := 0; i < xChange; i++ {
			moves = append(moves, "L 1")
		}
	}
	return moves
}

func DiagMove(tail Position, head Position) string {
	move := "X 0"
	// Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves one step diagonally
	if head.Y > tail.Y+1 && head.X >= tail.X+1 || head.Y >= tail.Y+1 && head.X > tail.X+1 {
		move = "UR 1"
	} else if head.Y > tail.Y+1 && head.X <= tail.X-1 || head.Y >= tail.Y+1 && head.X < tail.X-1 {
		move = "UL 1"
	} else if head.Y < tail.Y-1 && head.X >= tail.X+1 || head.Y <= tail.Y-1 && head.X > tail.X+1 {
		move = "DR 1"
	} else if head.Y < tail.Y-1 && head.X <= tail.X-1 || head.Y <= tail.Y-1 && head.X < tail.X-1 {
		move = "DL 1"
	}
	return move
}

func FormatGrid(minX int, minY int, maxX int, maxY int, knots []Position, visitedPositions map[Position]int) string {
	grid := ""
	for yIndex := maxY; yIndex >= minY; yIndex-- {
		for xIndex := minX; xIndex <= maxX; xIndex++ {
			if xIndex == minX {
				grid = fmt.Sprintf("%s %-3d", grid, yIndex)
			}
			knotHere := -1
			for knotIndex := range knots {
				if xIndex == knots[knotIndex].X && yIndex == knots[knotIndex].Y {
					knotHere = knotIndex
				}
			}
			if xIndex == knots[0].X && yIndex == knots[0].Y {
				if knots[0].X == knots[len(knots)-1].X && knots[0].Y == knots[len(knots)-1].Y {
					grid = fmt.Sprintf("%s %-3s", grid, "0")

				} else {
					grid = fmt.Sprintf("%s %-3s", grid, "H")

				}
			} else if knotHere > 0 && knotHere < len(knots)-1 {
				grid = fmt.Sprintf("%s %-3d", grid, knotHere)
			} else if xIndex == knots[len(knots)-1].X && yIndex == knots[len(knots)-1].Y {
				grid = fmt.Sprintf("%s %-3s", grid, "T")
			} else {
				_, exists := visitedPositions[Position{xIndex, yIndex}]
				if exists {
					grid = fmt.Sprintf("%s %-3s", grid, "#")
				} else {
					grid = fmt.Sprintf("%s %-3s", grid, ".")

				}
			}
		}

		grid = fmt.Sprintf("%s %s", grid, "\n")
		if yIndex == minY {
			grid = fmt.Sprintf("%s %-3s", grid, "")
			for xIndex := minX; xIndex <= maxX; xIndex++ {
				grid = fmt.Sprintf("%s %-3d", grid, xIndex)
			}
		}
	}
	return grid
}

func FormatGridVisited(minX int, minY int, maxX int, maxY int, visitedPositions map[Position]int) string {
	grid := ""
	for yIndex := maxY; yIndex >= minY; yIndex-- {
		for xIndex := minX; xIndex <= maxX; xIndex++ {
			if xIndex == minX {
				grid = fmt.Sprintf("%s %-3d", grid, yIndex)
			}
			num, exists := visitedPositions[Position{xIndex, yIndex}]
			if exists && num > 0 {
				grid = fmt.Sprintf("%s %-3s", grid, "#")
			} else {
				grid = fmt.Sprintf("%s %-3s", grid, ".")
			}
		}
		grid = fmt.Sprintf("%s %s", grid, "\n")
		if yIndex == minY {
			grid = fmt.Sprintf("%s %-3s", grid, "")
			for xIndex := minX; xIndex <= maxX; xIndex++ {
				grid = fmt.Sprintf("%s %-3d", grid, xIndex)
			}
		}
	}
	return grid
}

func Min(one int, two int) int {
	if one < two {
		return one
	} else {
		return two
	}
}

func Max(one int, two int) int {
	if one > two {
		return one
	} else {
		return two
	}
}

func RecordTailPosition(tail Position, tailVisitedPositions map[Position]int) map[Position]int {
	entry, exists := tailVisitedPositions[Position{tail.X, tail.Y}]
	if exists {
		tailVisitedPositions[Position{tail.X, tail.Y}] = entry + 1
	} else {
		tailVisitedPositions[Position{tail.X, tail.Y}] = 1
	}
	return tailVisitedPositions
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Sets the output to verbose")
	flag.BoolVar(&showVisited, "visited", false, "Show the visited grid")
	flag.Parse()
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 9\n===============-======")
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	var tailVisitedPositions = map[Position]int{{0, 0}: 1}
	knots := []Position{{0, 0}, {0, 0}}
	moves := strings.Split(string(data), "\n")
	if verbose {
		fmt.Println("\nPart One:")
	}
	for _, move := range moves {
		knots[0] = Move(knots[0], move)
		knots[len(knots)-1] = Move(knots[len(knots)-1], DiagMove(knots[len(knots)-1], knots[0]))
		RecordTailPosition(knots[len(knots)-1], tailVisitedPositions)
		for _, tailMove := range StraightMoves(knots[len(knots)-1], knots[0]) {
			knots[len(knots)-1] = Move(knots[len(knots)-1], tailMove)
			RecordTailPosition(knots[len(knots)-1], tailVisitedPositions)
		}
		minX = Min(minX, Min(knots[len(knots)-1].X, knots[0].X))
		minY = Min(minY, Min(knots[len(knots)-1].Y, knots[0].Y))
		maxX = Max(maxX, Max(knots[len(knots)-1].X, knots[0].X))
		maxY = Max(maxY, Max(knots[len(knots)-1].Y, knots[0].Y))

		if verbose {
			fmt.Println(FormatGrid(minX, minY, maxX, maxY, knots, tailVisitedPositions))
			fmt.Printf("Head {%d:%d} and tail {%d,%d} after %s \n", knots[0].X, knots[0].Y, knots[len(knots)-1].X, knots[len(knots)-1].Y, move)
			fmt.Println()
		}
	}

	if verbose || showVisited {
		fmt.Println("Visited positions")
		fmt.Println(FormatGridVisited(minX, minY, maxX, maxY, tailVisitedPositions))
		fmt.Println()
	}

	partOneVisited := len(tailVisitedPositions)

	if verbose {
		fmt.Println("\nPart Two:")
	}
	tailVisitedPositions = map[Position]int{}
	knots = []Position{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	for _, move := range moves {
		// move head
		distance, err := strconv.ParseInt(strings.Split(move, " ")[1], 0, 0)
		if err != nil {
			log.Fatalf("couldn't get distance %v", err)
		}
		for i := 0; i < int(distance); i++ {
			knots[0] = Move(knots[0], fmt.Sprintf("%s %d", strings.Split(move, " ")[0], 1))
			// move rest
			for knotIndex := 1; knotIndex < len(knots); knotIndex++ {
				knots[knotIndex] = Move(knots[knotIndex], DiagMove(knots[knotIndex], knots[knotIndex-1]))
				RecordTailPosition(knots[len(knots)-1], tailVisitedPositions)
				for _, tailMove := range StraightMoves(knots[knotIndex], knots[knotIndex-1]) {
					knots[knotIndex] = Move(knots[knotIndex], tailMove)
					RecordTailPosition(knots[len(knots)-1], tailVisitedPositions)
				}
				minX = Min(minX, Min(knots[knotIndex].X, knots[knotIndex-1].X))
				minY = Min(minY, Min(knots[knotIndex].Y, knots[knotIndex-1].Y))
				maxX = Max(maxX, Max(knots[knotIndex].X, knots[knotIndex-1].X))
				maxY = Max(maxY, Max(knots[knotIndex].Y, knots[knotIndex-1].Y))
			}
		}

		if verbose {
			fmt.Println(FormatGrid(minX, minY, maxX, maxY, knots, tailVisitedPositions))
			fmt.Printf("Head {%d:%d} and tail {%d,%d} after %s \n", knots[0].X, knots[0].Y, knots[len(knots)-1].X, knots[len(knots)-1].Y, move)
			fmt.Println()
		}
	}

	if verbose || showVisited {
		fmt.Println("Visited positions")
		fmt.Println(FormatGridVisited(minX, minY, maxX, maxY, tailVisitedPositions))
		fmt.Println()
	}

	partTwoVisited := len(tailVisitedPositions)

	fmt.Println("\nResults:")
	fmt.Printf("\tVisited Positions of tail with 2 knots: %d\n", partOneVisited)  // 5902
	fmt.Printf("\tVisited Positions of tail with 10 knots: %d\n", partTwoVisited) // 2445
}
