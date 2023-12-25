package main

import (
	"bufio"
	"math/rand"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 25")
	fmt.Println("=====================")
	input := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(Solve(input))
}

func Parse(filename string) []Edge {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	edges := make([]Edge, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitted := strings.Split(line, ": ")
		edges = append(edges, ParseList(splitted[0], splitted[1])...)
	}
	readFile.Close()
	return edges
}

func ParseList(id string, list string) []Edge {
	splitted := strings.Split(list, " ")
	listFin := []Edge{}
	for _, x := range splitted {
		listFin = append(listFin, Edge{id, x})
	}
	return listFin
}


func Solve(edges []Edge) int {
	kargerResult := 0
	a, b := 0, 0
	for kargerResult != 3 {
		kargerResult, a, b = Karger(edges)
	}
	return a * b
}

type Edge struct {
	from string
	to string
}

func Karger(input []Edge) (int, int, int) {
	edges := make([]Edge, len(input))
	nodes := make(map[string]int, 0)
	for i, edge := range input {
		edges[i] = input[i]
		nodes[edge.from] = 1
		nodes[edge.to] = 1
	}
	for len(nodes) > 2 {
		toDelete := rand.Intn(len(edges))
		edge := edges[toDelete]
		edge2 := Edge{edge.to, edge.from}
		n1 := nodes[edge.from]
		n2 := nodes[edge.to]
		delete(nodes, edge.from)
		delete(nodes, edge.to)
		newEdge := edge.to + edge.from
		nodes[newEdge] = n1 + n2

		newEdges := make([]Edge, 0)
		for i := range edges {
			if edges[i] == edge || edges[i] == edge2 {
				continue
			} else if edges[i].to == edge.to || edges[i].to == edge.from {
				newEdges = append(newEdges, Edge{edges[i].from, newEdge})
			} else if edges[i].from == edge.to || edges[i].from == edge.from {
				newEdges = append(newEdges, Edge{newEdge, edges[i].to})
			} else {
				newEdges = append(newEdges, edges[i])
			}
		}
		edges = nil
		edges = newEdges
	}

	return len(edges), nodes[edges[0].from], nodes[edges[0].to]
}
