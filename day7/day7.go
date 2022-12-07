package main

import (
	"fmt" // console io
	"math"
	"os"      // file io
	"strconv" // int parsing
	"strings" // string manipulation
)

// Node is a file or directory
type Node struct {
	Name string // name of node on disk
	Path string // path to the current node excluding it's name
	Size int    // either of the file or containing files for directories
	Type string // "d" or "-"
}

func List(input map[string]Node) string {
	tree := ""
	for _, node := range input {
		tree += fmt.Sprintf("%s%15d\t%20s - %s\n", node.Type, node.Size, node.Name, node.Path)
	}
	return tree
}

func SmallerThanByType(input map[string]Node, threshold int, t string) []Node {
	var nodes []Node
	for _, node := range input {
		if node.Type == t && node.Size <= threshold {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func NearestGreatestByType(input map[string]Node, threshold int, t string) Node {
	var nearest Node
	distance := math.MaxInt

	for _, node := range input {
		if node.Type == t && node.Size >= threshold && node.Size-threshold < distance {
			distance = node.Size - threshold
			nearest = node
		}
	}
	return nearest
}

func IsCommand(input string) bool {
	return len(input) > 2 && input[0:2] == "$ "
}

func Run(input []string) map[string]Node {
	var workingDirectory []string // start empty
	nodes := map[string]Node{"/": {Path: "", Name: "/", Size: 0, Type: "d"}}
	for _, line := range input {
		if IsCommand(line) {
			// map command
			switch line[:4] {
			case "$ cd":
				{
					// update working directory
					if line[5:] == ".." {
						workingDirectory = workingDirectory[:len(workingDirectory)-1]
					} else {
						workingDirectory = append(workingDirectory, line[5:])
					}
				}
			case "$ ls":
				{
					// we just need to process following lines
				}
			}
		} else {
			// add to file/dir list
			details := strings.Split(line, " ")
			// check for existing node in list
			key := fmt.Sprintf("%s/%s", strings.Join(workingDirectory, "/"), details[1])
			// insert
			if details[0] == "dir" {
				// size is unknown until we traverse
				nodes[key] = Node{Name: details[1], Path: strings.Join(workingDirectory, "/"), Size: 0, Type: "d"}
			} else {
				size, _ := strconv.ParseInt(details[0], 10, 0)
				nodes[key] = Node{Name: details[1], Path: strings.Join(workingDirectory, "/"), Size: int(size), Type: "-"}
				// dir entries should be updated incrementally (traverse up)
				for levelNumber := range workingDirectory {
					key = strings.Join(workingDirectory[:len(workingDirectory)-levelNumber], "/")
					entry, exists := nodes[key]
					if exists {
						entry.Size += int(size)
						nodes[key] = entry
					} else {
						fmt.Printf("no entry for: %s\n", key)
					}
				}
			}
		}
	}
	return nodes
}

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Advent of Code - Day 5\n===============-======")
	lines := strings.Split(string(data), "\n")
	nodes := Run(lines)
	fmt.Println(List(nodes))
	fmt.Println("\nPart One:")
	sum1 := 0
	for _, node := range SmallerThanByType(nodes, 100000, "d") {
		sum1 += node.Size
		fmt.Printf("%60s/%s\t%d\n", node.Path, node.Name, node.Size)
	}
	fmt.Println("\nPart Two:")
	fmt.Printf("Used space: %d\n", nodes["/"].Size)
	fmt.Printf("Free space: %d\n", 70000000-nodes["/"].Size)
	fmt.Printf("Required free space: %d\n", 30000000)
	deficit := 30000000 - (70000000 - nodes["/"].Size)
	fmt.Printf("Deficit free space: %d\n", deficit)
	nearest := NearestGreatestByType(nodes, deficit, "d")
	fmt.Printf("Directory nearest deficit: %v\n", nearest)

	fmt.Println("\nResults:")
	fmt.Printf("\tSum of directories less than 100000: %d\n", sum1)           // 1501149
	fmt.Printf("\tSize of directory nearest %d: %d\n", deficit, nearest.Size) // 10096985

}
