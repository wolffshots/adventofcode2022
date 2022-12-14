package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// via https://www.reddit.com/u/Multipl in essence
func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("couldn't close file")
		}
	}(file)
	packets := parseInput(file)
	fmt.Println("Advent of Code - Day 13\n===============-=======")

	fmt.Println("Results:")
	fmt.Printf("\t%-14s: %6d\n", "Sum of indices", solvePart1(packets))
	fmt.Printf("\t%-14s: %6d\n", "Decoder key", solvePart2(packets))
}

func parseInput(r io.Reader) [][]interface{} {
	scanner := bufio.NewScanner(r)
	var packets [][]interface{}
	var pair []interface{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			packets = append(packets, pair)
			pair = []interface{}{}
		} else {
			var p interface{}
			err := json.Unmarshal([]byte(scanner.Text()), &p)
			if err != nil {
				log.Fatal(err)
			}
			pair = append(pair, p)
		}
	}

	return append(packets, pair)
}

func solvePart1(packets [][]interface{}) int {
	result := 0
	for i, pair := range packets {
		if compare(pair[0], pair[1]) <= 0 {
			result += i + 1
		}
	}

	return result
}

func compare(p1 interface{}, p2 interface{}) int {
	_, ok1 := p1.(float64)
	_, ok2 := p2.(float64)
	if ok1 && ok2 {
		return int(p1.(float64)) - int(p2.(float64))
	}

	if ok1 {
		p1 = []interface{}{p1}
	}
	if ok2 {
		p2 = []interface{}{p2}
	}

	if len(p1.([]interface{})) == 0 || len(p2.([]interface{})) == 0 {
		return len(p1.([]interface{})) - len(p2.([]interface{}))
	}

	result := compare(p1.([]interface{})[0], p2.([]interface{})[0])
	if result == 0 {
		next1 := p1.([]interface{})[1:]
		next2 := p2.([]interface{})[1:]

		if len(next1) == 0 || len(next2) == 0 {
			return len(next1) - len(next2)
		}
		return compare(next1, next2)
	}

	return result
}

func solvePart2(packets [][]interface{}) int {
	var newPacket []interface{}
	for _, pair := range packets {
		newPacket = append(newPacket, pair...)
	}

	var divider1 interface{}
	err := json.Unmarshal([]byte("[[2]]"), &divider1)
	if err != nil {
		log.Fatal(err)
	}

	var divider2 interface{}
	err = json.Unmarshal([]byte("[[6]]"), &divider2)
	if err != nil {
		log.Fatal(err)
	}

	newPacket = append(newPacket, []interface{}{divider1, divider2}...)
	sort.Slice(newPacket, func(i, j int) bool {
		return compare(newPacket[i], newPacket[j]) <= 0
	})

	result := 1
	for i, packet := range newPacket {
		packet, err := json.Marshal(packet)
		if err != nil {
			log.Fatal(err)
		}
		packetString := string(packet)
		if packetString == "[[2]]" || packetString == "[[6]]" {
			result *= i + 1
		}
	}

	return result
}
