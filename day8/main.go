package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Advent of Code, day 8")
	fmt.Println("=====================")
	directions, nodes := GetInput("input.txt")
	first := Solve(directions, nodes)
	fmt.Print("*  ")
	fmt.Println(first)
}

func Solve(directionString string, nodeList []string) int {
	nodes := ParseNodes(nodeList)
	directions := ParseDirections(directionString)
	nodeMap := constructMap(nodes)
	fmt.Println(nodeMap)
	return Walk(directions, nodeMap)
}

func GetInput(filename string) (string, []string) {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var directions string
	var nodes []string
	i := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if i == 0 {
			directions = line
		}
		if i > 1 {
			nodes = append(nodes, line)
		}
		i++;
	}
	readFile.Close()
	return directions, nodes
}

func ParseNodes(lines []string) []Node {
	var nodes []Node
	for _, line := range lines {
		node := ParseNode(line)
		nodes = append(nodes, node)
	}
	return nodes
}

func ParseNode(line string) Node {
	reg := regexp.MustCompile("([A-Z]+) = \\(([A-Z]+), ([A-Z]+)\\)$")
	match := reg.FindStringSubmatch(line)
	name := match[1]
	left := match[2]
	right := match[3]
	return Node{name, left, right}
}

func ParseDirections(line string) []bool {
	chars := []int32(line)
	result := make([]bool, 0)
	left := int32('L');
	right := int32('R');
	for _, char := range chars{
		if char == left {
			result = append(result, true)
		} else if char == right {
			result = append(result, false)
		}

	}
	return result
}

func constructMap(nodes []Node) map[string]Node {
	hash := make(map[string]Node)
	for _, node := range nodes {
		hash[node.name] = node
	}
	return hash
}

func Walk(dirs []bool, nodes map[string]Node) int {
	steps := 0
	curr := "AAA"
	for curr != "ZZZ" {
		node := nodes[curr]
		dir := steps % len(dirs)
		if dirs[dir] {
			curr = node.left
		} else {
			curr = node.right
		}
		steps++
	}
	return steps
}

type Node struct {
	name string
	left string
	right string
}
