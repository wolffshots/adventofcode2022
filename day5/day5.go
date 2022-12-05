package main

import (
	"fmt"     // console io
	"os"      // file io
	"strconv" // string parsing
	"strings" // string formatting
)

func PrintBoard(board [][]string) string {
	maxHeight := 0
	output := ""
	for _, stackValue := range board {
		if len(stackValue) > maxHeight {
			maxHeight = len(stackValue)
		}
	}

	for height := maxHeight - 1; height >= 0; height-- {
		for _, stackValue := range board {
			if height < len(stackValue) {
				output += fmt.Sprintf(" %v ", stackValue[height])
			} else {
				output += fmt.Sprint("  e  ")
			}
		}
		output += fmt.Sprintln()
	}
	for stackIndex := 0; stackIndex < len(board); stackIndex++ {
		output += fmt.Sprintf("  %d  ", stackIndex+1)
	}
	return output
}

func ParseBoard(input string) [][]string {
	var tempBoard [][]string
	boardLines := strings.Split(input, "\n")
	stacks := strings.Split(strings.Trim(boardLines[len(boardLines)-1], " "), "   ")

	// setup board structure
	for range stacks {
		tempBoard = append(tempBoard, []string{})
	}

	// parse the board into the board slice of slices
	for boardInputIndex := len(boardLines) - 2; boardInputIndex >= 0; boardInputIndex-- {
		for stackIndex := 0; stackIndex < len(stacks); stackIndex++ {
			if ((stackIndex+1)*4)-1 <= len(boardLines[boardInputIndex]) {
				cell := boardLines[boardInputIndex][(stackIndex * 4) : ((stackIndex+1)*4)-1]
				if cell != "   " {
					tempBoard[stackIndex] = append(tempBoard[stackIndex], cell)
				}
			}
		}
	}
	return tempBoard
}

func TopOfBoard(board [][]string) string {
	tempTop := ""
	for _, stackValue := range board {
		if len(stackValue) > 0 {
			tempTop = fmt.Sprintf("%s%s", tempTop, stackValue[len(stackValue)-1:][0][1:2])
		}
	}
	return tempTop
}

func Move(board [][]string, move string, singular bool) [][]string {
	slicedMove := strings.Split(move, " ")
	number, err := strconv.ParseInt(slicedMove[1], 10, 0)
	if err != nil {
		fmt.Printf("Error parsing to int %v\n", err)
	}
	origin, err := strconv.ParseInt(slicedMove[3], 10, 0)
	origin-- // convert to index
	if err != nil {
		fmt.Printf("Error parsing to int %v\n", err)
	}
	destination, err := strconv.ParseInt(slicedMove[5], 10, 0)
	destination-- // convert to index
	if err != nil {
		fmt.Printf("Error parsing to int %v\n", err)
	}

	var tempBoard [][]string
	// setup board structure
	for range board {
		tempBoard = append(tempBoard, []string{})
	}

	// move all the untouched stacks
	for stackIndex, _ := range board {
		if int64(stackIndex) != origin {
			tempBoard[stackIndex] = board[stackIndex]
		}
	}

	for index := 0; index < int(number); index++ {
		if singular {
			tempBoard[destination] = append(tempBoard[destination], board[origin][len(board[origin])-1-index])
		} else {
			tempBoard[destination] = append(tempBoard[destination], board[origin][len(board[origin])-int(number)+index])
		}

	}
	tempBoard[origin] = board[origin][:(len(board[origin]) - int(number))]

	return tempBoard
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 5\n===============-======\n\nPart One:")
	fmt.Println()

	input := strings.Split(string(data), "\n\n")
	board1 := ParseBoard(input[0])

	fmt.Println("Original Board:")
	fmt.Println(PrintBoard(board1))
	fmt.Println()

	moves := strings.Split(input[1], "\n")
	for _, move := range moves {
		fmt.Printf("Move: %s\n", move)
		board1 = Move(board1, move, true)
		fmt.Println(PrintBoard(board1))
		fmt.Println()
	}

	fmt.Println("\nPart Two:")

	board2 := ParseBoard(input[0])
	for _, move := range moves {
		fmt.Printf("Move: %s\n", move)
		board2 = Move(board2, move, false)
		fmt.Println(PrintBoard(board2))
		fmt.Println()
	}

	fmt.Println("\nResults:")
	fmt.Printf("\tTop of board after moves: %s\n", TopOfBoard(board1))
	fmt.Printf("\tTop of board after moves: %s\n", TopOfBoard(board2))

}
