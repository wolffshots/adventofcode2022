package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var verbose bool

type Tree struct {
	Height int
}

func FormatGrid(trees [][]Tree) string {
	formatted := ""
	for rowIndex, row := range trees {
		if rowIndex != 0 {
			formatted = fmt.Sprintf("%s\n", formatted)
		}
		for columnIndex, tree := range row {
			if columnIndex != 0 {
				formatted = fmt.Sprintf("%s %d", formatted, tree.Height)
			} else {
				formatted = fmt.Sprintf("%s%d", formatted, tree.Height)
			}
		}
	}
	return formatted
}

func FormatGridVisibility(trees [][]Tree) string {
	formatted := ""
	for rowIndex, row := range trees {
		if rowIndex != 0 {
			formatted = fmt.Sprintf("%s\n", formatted)
		} else {
			for columnIndex := range row {
				formatted = fmt.Sprintf("%s\t%d        ", formatted, columnIndex)
			}
			formatted = fmt.Sprintf("%s\n", formatted)
		}
		for columnIndex, tree := range row {
			if columnIndex != 0 {
				formatted = fmt.Sprintf("%s\t%d - %t", formatted, tree.Height, IsVisible(trees, rowIndex, columnIndex))
			} else {
				formatted = fmt.Sprintf("%s%d\t%d - %t", formatted, rowIndex, tree.Height, IsVisible(trees, rowIndex, columnIndex))
			}
		}
	}
	return formatted
}

func FormatGridViewing(trees [][]Tree) string {
	formatted := ""
	for rowIndex, row := range trees {
		if rowIndex != 0 {
			formatted = fmt.Sprintf("%s\n", formatted)
		} else {
			formatted = fmt.Sprintf("%s\t", formatted)
			for columnIndex := range row {
				formatted = fmt.Sprintf("%s%-3d   %-6s ", formatted, columnIndex, "")
			}
			formatted = fmt.Sprintf("%s\n", formatted)
		}
		for columnIndex, tree := range row {
			scores := ScenicScores(trees, rowIndex, columnIndex)
			score := 1
			for scoreIndex := range scores {
				score *= scores[scoreIndex]
			}
			if columnIndex != 0 {
				formatted = fmt.Sprintf("%s%-3d - %-6d ", formatted, tree.Height, score)
			} else {
				formatted = fmt.Sprintf("%s%-3d\t%-3d - %-6d ", formatted, rowIndex, tree.Height, score)
			}
		}
	}
	return formatted
}

func ParseGrid(input string) [][]Tree {
	var trees [][]Tree
	rows := strings.Split(input, "\n")
	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		trees = append(trees, []Tree{})
		columns := strings.Split(rows[rowIndex], "")
		for columnIndex := 0; columnIndex < len(columns); columnIndex++ {
			height, err := strconv.ParseInt(columns[columnIndex], 10, 0)
			if err != nil {
				log.Fatalf("Error converting input string to height: %v", err)
			}
			trees[rowIndex] = append(trees[rowIndex], Tree{Height: int(height)})
		}
	}

	return trees
}

func IsVisibleHorizontal(trees [][]Tree, treeRowIndex int, treeColumnIndex int) bool {
	// visible left
	visibleLeft := true
	visibleRight := true
	for columnIndex := 0; columnIndex < treeColumnIndex; columnIndex++ {
		if trees[treeRowIndex][columnIndex].Height >= trees[treeRowIndex][treeColumnIndex].Height {
			visibleLeft = false
		}
	}
	// visible right
	for columnIndex := treeColumnIndex + 1; columnIndex < len(trees[treeRowIndex]); columnIndex++ {
		if trees[treeRowIndex][columnIndex].Height >= trees[treeRowIndex][treeColumnIndex].Height {
			visibleRight = false
		}
	}
	return visibleLeft || visibleRight
}

func IsVisibleVertical(trees [][]Tree, treeRowIndex int, treeColumnIndex int) bool {
	// visible left
	visibleUp := true
	visibleDown := true
	for rowIndex := 0; rowIndex < treeRowIndex; rowIndex++ {
		if trees[rowIndex][treeColumnIndex].Height >= trees[treeRowIndex][treeColumnIndex].Height {
			visibleUp = false
		}
	}
	// visible right
	for rowIndex := treeRowIndex + 1; rowIndex < len(trees); rowIndex++ {
		if trees[rowIndex][treeColumnIndex].Height >= trees[treeRowIndex][treeColumnIndex].Height {
			visibleDown = false
		}
	}
	return visibleUp || visibleDown
}

func IsVisible(trees [][]Tree, treeRowIndex int, treeColumnIndex int) bool {
	// if is edge return true
	if treeRowIndex == 0 || treeColumnIndex == 0 || treeRowIndex == len(trees)-1 || treeColumnIndex == len(trees[treeRowIndex])-1 {
		return true
	}
	return IsVisibleHorizontal(trees, treeRowIndex, treeColumnIndex) || IsVisibleVertical(trees, treeRowIndex, treeColumnIndex)
}

func ScenicScores(trees [][]Tree, treeRowIndex int, treeColumnIndex int) []int {
	scenicScoreLeft := 0
	for columnIndex := treeColumnIndex - 1; columnIndex >= 0; columnIndex-- {
		if trees[treeRowIndex][columnIndex].Height >= trees[treeRowIndex][treeColumnIndex].Height {
			scenicScoreLeft += 1
			break
		}
		scenicScoreLeft += 1
	}
	scenicScoreRight := 0
	for columnIndex := treeColumnIndex + 1; columnIndex < len(trees[treeRowIndex]); columnIndex++ {
		if trees[treeRowIndex][columnIndex].Height >= trees[treeRowIndex][treeColumnIndex].Height {
			scenicScoreRight += 1
			break
		}
		scenicScoreRight += 1
	}
	scenicScoreUp := 0
	for rowIndex := treeRowIndex - 1; rowIndex >= 0; rowIndex-- {
		if trees[rowIndex][treeColumnIndex].Height >= trees[treeRowIndex][treeColumnIndex].Height {
			scenicScoreUp += 1
			break
		}
		scenicScoreUp += 1
	}
	scenicScoreDown := 0
	for rowIndex := treeRowIndex + 1; rowIndex < len(trees[treeRowIndex]); rowIndex++ {
		if trees[rowIndex][treeColumnIndex].Height >= trees[treeRowIndex][treeColumnIndex].Height {
			scenicScoreDown += 1
			break
		}
		scenicScoreDown += 1
	}
	return []int{scenicScoreLeft, scenicScoreRight, scenicScoreUp, scenicScoreDown}
}

func CountVisible(trees [][]Tree) int {
	count := 0
	for rowIndex := range trees {
		for columnIndex := range trees {
			if IsVisible(trees, rowIndex, columnIndex) {
				count++
			}
		}
	}
	return count
}

func MaxScenicScore(trees [][]Tree) int {
	max := 0
	for rowIndex := range trees {
		for columnIndex := range trees {
			scores := ScenicScores(trees, rowIndex, columnIndex)
			score := 1
			for scoreIndex := range scores {
				score *= scores[scoreIndex]
			}
			if score > max {
				max = score
			}
		}
	}
	return max
}

func main() {
	flag.BoolVar(&verbose, "v", false, "Sets the output to verbose")
	flag.Parse()
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 8\n===============-======")
	trees := ParseGrid(string(data))

	if verbose {
		fmt.Println("\nPart One:")

		fmt.Println("\nGrid: ")
		fmt.Println(FormatGrid(trees))

		fmt.Println("\nVisibility: ")
		fmt.Println(FormatGridVisibility(trees))

		fmt.Println("\nPart Two:")

		fmt.Println("\nViewing: ")
		fmt.Println(FormatGridViewing(trees))
	}
	visibleTrees := CountVisible(trees)
	maxScenicScore := MaxScenicScore(trees)

	fmt.Println("\nResults:")
	fmt.Printf("\tTrees visible from the outside: %d\n", visibleTrees)
	fmt.Printf("\tMax scenic score: %d\n\n", maxScenicScore)
}
