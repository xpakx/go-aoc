package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 17")
	fmt.Println("=====================")
	input := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(Dijkstra(input, 0, 3))
	fmt.Print("** ")
	fmt.Println(Dijkstra(input, 4, 10))
}

func Parse(filename string) [][]int {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make([][]int, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums := make([]int, 0)
		for i := range line {
			num := int(line[i] - '0')
			nums = append(nums, num)
		}
		result = append(result, nums)
	}
	readFile.Close()
	return result
}

type Key struct {
	pos  Pos
	dir int
	steps int
}

type Pos struct {
	x int
	y int
}

const (
	Right = 0
	Left = 1
	Up = 2
	Down = 3
)

func Dijkstra(input [][]int, minSteps int, maxSteps int) int {
	list := make([]Elem, 0)
	dist := make(map[Key]int)
	prev := make(map[Key]Key)
	visited := make(map[Key]struct{})
	for i := range input {
		for j := range input[i] {
			for dir:=0; dir<4; dir++ {
				for distance:=1; distance<=maxSteps; distance++ {
					key := Key{Pos{j,i}, dir, distance}
					dist[key] = math.MaxInt
					if i != 0 || j != 0 {
						list = append(list, Elem{key, math.MaxInt})
					} else {
						dist[key] = input[0][0]
					}
				}
			}
		}
	}

	key := Key{Pos{0,0}, -1, 0}
	dist[key] = 0
	list = append(list, Elem{key, 0})
	queue := buildHeap(list)

	for len(*queue) > 0 {
		minElem := heap.Pop(queue)
		elem := minElem.(Elem)
		minNode := elem.key
		minDist := elem.dist
		if minDist == math.MaxInt {
			break
		}
		if _, ok := visited[minNode]; ok {
			continue
		}
		visited[minNode] = struct{}{}

		neighbours := GetNeighbours(minNode, len(input[0]), len(input), minSteps, maxSteps)
		if len(neighbours) == 0 {
			dist[minNode] = math.MaxInt
		}
		for _, n := range neighbours {
			nDist := dist[n]
			alt := minDist + input[n.pos.y][n.pos.x]
			if minDist == math.MaxInt {
				alt = math.MaxInt
			}
			if alt < nDist {
				dist[n] = alt
				heap.Push(queue, Elem{n, alt})
				prev[n] = minNode
			}
		}
	}

	minValue := math.MaxInt
	for dir:=0; dir<4; dir++ {
		for distance:=1; distance<=maxSteps; distance++ {
			key := Key{Pos{len(input[0])-1,len(input)-1}, dir, distance}
			if dist[key] < minValue {
				minValue = dist[key]
			}
		}

	}
	return minValue
}

func GetNeighbours(node Key, width int, height int, minSteps int, maxSteps int) []Key {
	n := make([]Key, 0)
	if node.steps < minSteps && node.steps > 0 {
		newX := node.pos.x
		x := node.pos.x
		newY := node.pos.y
		y := node.pos.y
		dist := minSteps-node.steps
		if node.dir == Right {
			newX+=dist
			x++
		} else if node.dir == Left {
			newX-=dist
			x--
		} else if node.dir == Up {
			newY-=dist
			y--
		} else if node.dir == Down {
			newY+=dist
			y++
		}
		if newX >= 0 && newX < width && newY >=0 && newY < height {
			n = append(n, Key{Pos{x, y}, node.dir, node.steps+1})
		}
		return n
	}
	if node.pos.x > 0 && (node.dir != Left || node.steps < maxSteps) && node.dir != Right {
		steps := node.steps + 1
		if node.dir != Left {
			steps = 1
		}
		n = append(n, Key{Pos{node.pos.x-1, node.pos.y}, Left, steps})
	}
	if node.pos.x < width-1 && (node.dir != Right || node.steps < maxSteps) && node.dir != Left {
		steps := node.steps + 1
		if node.dir != Right {
			steps = 1
		}
		n = append(n, Key{Pos{node.pos.x+1, node.pos.y}, Right, steps})
	}
	if node.pos.y > 0 && (node.dir != Up || node.steps < maxSteps) && node.dir != Down {
		steps := node.steps + 1
		if node.dir != Up {
			steps = 1
		}
		n = append(n, Key{Pos{node.pos.x, node.pos.y-1}, Up, steps})
	}
	if node.pos.y < height-1 && (node.dir != Down || node.steps < maxSteps) && node.dir != Up {
		steps := node.steps + 1
		if node.dir != Down {
			steps = 1
		}
		n = append(n, Key{Pos{node.pos.x, node.pos.y+1}, Down, steps})
	}
	return n 
}

type Elem struct {
	key Key
	dist int
}

type MinHeap []Elem

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Elem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func buildHeap(array []Elem) *MinHeap {
	minHeap := &MinHeap{}
	for _, elem := range array {
		heap.Push(minHeap, elem)
	}
	return minHeap
}
