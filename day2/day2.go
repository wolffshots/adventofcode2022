package main

import (
	"fmt"     // console io
	"os"      // file io
	"strings" // string splitting
)

var shapeValues = map[string]int{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

// Paper > Rock > Scissors
// Rock > Scissors > Paper
// Scissors > Paper > Rock

// win  6
// draw 3
// loss 0

func calculateScore1(opponentShape string, selfShape string) int {
	score := shapeValues[selfShape]
	switch selfShape {
	case "X": // Rock
		switch opponentShape {
		case "A": // Rock
			score += 3 // draw
		case "B": // Paper
			score += 0 // loss
		case "C": // Scissors
			score += 6 // win
		}
	case "Y": // Paper
		switch opponentShape {
		case "A": // Rock
			score += 6 // win
		case "B": // Paper
			score += 3 // draw
		case "C": // Scissors
			score += 0 // loss
		}
	case "Z": // Scissors
		switch opponentShape {
		case "A": // Rock
			score += 0 // loss
		case "B": // Paper
			score += 6 // win
		case "C": // Scissors
			score += 3 // draw
		}
	}
	return score
}

func calculateScore2(opponentShape string, selfShape string) int {
	score := 0
	switch selfShape {
	case "X": // lose
		score += 0 // loss
		switch opponentShape {
		case "A": // Rock
			score += 3 // Scissors
		case "B": // Paper
			score += 1 // Rock
		case "C": // Scissors
			score += 2 // Paper
		}
	case "Y": // draw
		score += 3 // draw
		switch opponentShape {
		case "A": // Rock
			score += 1 // Rock
		case "B": // Paper
			score += 2 // Paper
		case "C": // Scissors
			score += 3 // Scissors
		}
	case "Z": // win
		score += 6 // win
		switch opponentShape {
		case "A": // Rock
			score += 2 // Paper
		case "B": // Paper
			score += 3 // Scissors
		case "C": // Scissors
			score += 1 // Rock
		}
	}
	return score
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	rounds := strings.Split(string(data), "\n")
	totalScore := 0
	for roundIndex, roundValue := range rounds {
		chosenShapes := strings.Split(roundValue, " ")
		roundScore := calculateScore2(chosenShapes[0], chosenShapes[1])
		totalScore = totalScore + roundScore
		fmt.Printf("round %d had %s as opponent and %s as self and a score of: %d\n", roundIndex, chosenShapes[0], chosenShapes[1], roundScore)
	}
	fmt.Printf("Total score: %d", totalScore)
}
