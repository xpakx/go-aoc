package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 10")
	fmt.Println("=====================")
	pipeMap := LoadMap("input.txt")
	dist := CalcDistances(pipeMap)
	first := solveFirst(dist)
	second := solveSecond(pipeMap, dist)
	PrintMap(pipeMap, dist)
	fmt.Print("*  ")
	fmt.Println(first)
	fmt.Print("** ")
	fmt.Println(second)
}

func solveFirst(dist map[Coord]int) int {
	maxValue := 0
	for _, v := range dist {
		if v != math.MaxInt && v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// using nonzero winding rule
func solveSecond(pipeMap [][]Node, dist map[Coord]int) int {
	counter := 0
	inside := 0
	for _, line := range pipeMap {
		for _, n := range line {
			if CheckChange(n.coord, pipeMap, dist) {
				if DoesPathIntersectUpward(n.coord, pipeMap, dist) {
					counter += 1
				} else {
					counter -= 1
				}
			}
			if distance, ok := dist[n.coord]; !ok || distance == math.MaxInt {
				if counter % 2 != 0 {
					inside += 1
				}
			}
		}
	}
	return inside
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

func CalcDistances(nodeMap [][]Node) map[Coord]int {
	start := make([]Node, 0)
	dist := make(map[Coord]int)
	for i := range nodeMap {
		for j := range nodeMap[i] {
			if nodeMap[i][j].IsPipe() {
				coord := nodeMap[i][j].coord
				dist[coord] = math.MaxInt
				if nodeMap[i][j].start {
					start = append(start, nodeMap[i][j])
					dist[coord] = 0
				}
			}
		}
	}

	n := start[0]

	neighbours := n.GetNeighboursCoord()
	for _, ne := range neighbours {
		prev := n.coord
		currNode := nodeMap[ne.i][ne.j]
		next := currNode.GetOtherCoord(prev)
		distance := 0
		for next != n.coord {
			prev = currNode.coord
			currNode = nodeMap[next.i][next.j]
			next = currNode.GetOtherCoord(prev)
			distance += 1
			if dist[prev] > distance {
				dist[prev] = distance
			}
		}
	}

	return dist
}

func (node Node) GetOtherCoord(enter Coord) Coord {
	coords := node.GetNeighboursCoord()
	for _, coord := range coords {
		if coord.i != enter.i || coord.j != enter.j {
			return coord
		}
	}
	return Coord{-1, -1}
}

func (node Node) GetNeighboursCoord() []Coord {
	coords := make([]Coord, 0)
	if node.north {
		coords = append(coords, Coord{node.coord.i-1, node.coord.j})
	}
	if node.south {
		coords = append(coords, Coord{node.coord.i+1, node.coord.j})
	}
	if node.west {
		coords = append(coords, Coord{node.coord.i, node.coord.j-1})
	}
	if node.east {
		coords = append(coords, Coord{node.coord.i, node.coord.j+1})
	}
	return coords
}

func (node Node) IsPipe() bool {
	return node.north || node.south || node.west || node.east
}

func PrintMap(pipeMap [][]Node, dist map[Coord]int) {
	for _, a := range pipeMap {
		for _, n := range a {
			if distance, ok := dist[n.coord]; ok && distance < math.MaxInt {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func CheckChange(coord Coord, pipeMap [][]Node, dist map[Coord]int) bool {
	if distance, ok := dist[coord]; !ok || distance == math.MaxInt {
		return false
	}
	if coord.i == 0 {
		return false
	}
	if !pipeMap[coord.i][coord.j].south {
		return false
	}
	coordBelow := Coord{coord.i+1, coord.j} 
	if distance, ok := dist[coordBelow]; !ok || distance == math.MaxInt {
		return false
	}
	return true
}

func DoesPathIntersectUpward(coord Coord, pipeMap [][]Node, dist map[Coord]int) bool {
	nodeDist := dist[coord]
	coordBelow := Coord{coord.i+1, coord.j} 
	belowDist := dist[coordBelow]
	if nodeDist + 1 == belowDist {
		return true
	}
	if nodeDist - 1 == belowDist {
		return false
	}
	return pipeMap[coord.i][coord.j].start
}
