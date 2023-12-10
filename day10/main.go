package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 10")
	fmt.Println("=====================")
	pipeMap := LoadMap("input.txt")
	for _, line := range pipeMap {
		fmt.Println(line)
	}
	fmt.Println()
	// fmt.Print("*  ")
	// fmt.Println(first)
}

func LoadMap(filename string) [][]Node {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make([][]Node, 0)
	startX := 0
	startY := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		values := strings.Split(line, "")
		nodes := toNodeList(values, len(result))
		for i := range nodes {
			if nodes[i].start {
				startX = i
				startY = len(result)
			}
		}
		result = append(result, nodes)
	}
	readFile.Close()

	startNode := result[startY][startX]
	if startX > 0 && result[startY][startX-1].east {
		startNode.west = true
	}
	if startX < len(result[startY])-1 && result[startY][startX+1].west {
		startNode.east = true
	}
	if startY > 0 && result[startY-1][startX].south {
		startNode.north = true
	}
	if startY < len(result)-1 && result[startY+1][startX].north {
		startNode.south = true
	}
	result[startY][startX] = startNode


	return result
}

func toNodeList(values []string, i int) []Node {
	result := make([]Node, 0)
	for j, value := range values {
		result = append(result, toNode(value, i, j))
	}
	return result
}

func toNode(value string, i int, j int) Node {
	if value == "L" {
		return Node{true, false, true, false, false, Coord{i, j}}
	}
	if value == "J" {
		return Node{false, true, true, false, false, Coord{i, j}}
	}
	if value == "7" {
		return Node{false, true, false, true, false, Coord{i, j}}
	}
	if value == "F" {
		return Node{true, false, false, true, false, Coord{i, j}}
	}
	if value == "|" {
		return Node{false, false, true, true, false, Coord{i, j}}
	}
	if value == "-" {
		return Node{true, true, false, false, false, Coord{i, j}}
	}
	if value == "S" {
		return Node{false, false, false, false, true, Coord{i, j}}
	}

	return Node{false, false, false, false, false, Coord{i, j}}
}

type Node struct {
	east bool
	west bool
	north bool
	south bool
	start bool
	coord Coord
}

type Coord struct {
	i int
	j int
}
