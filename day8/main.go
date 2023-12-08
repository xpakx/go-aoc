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
	first, second := Solve(directions, nodes)
	fmt.Print("*  ")
	fmt.Println(first)
	fmt.Print("** ")
	fmt.Println(second)
}

func Solve(directionString string, nodeList []string) (int, int) {
	nodes := ParseNodes(nodeList)
	directions := ParseDirections(directionString)
	nodeMap := constructMap(nodes)

	first := Walk(directions, nodeMap)
	second := WalkSecond(directions, nodeMap)
	return first, second
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

type MapKey struct {
	nodeName string
	dirPos int
}

type Cycle struct {
	ends []int
	cycle_start int
	cycle_end int
}

func FindCycle(start string, dirs []bool, nodes map[string]Node) Cycle {
	hash := make(map[MapKey]int)
	steps := 0
	curr := start
	ends := make([]int, 0)
	for true {
		node := nodes[curr]
		dir := steps % len(dirs)
		if dirs[dir] {
			curr = node.left
		} else {
			curr = node.right
		}
		if node.name[len(node.name)-1:] == "Z" {
			ends = append(ends, steps)
		}
		key := MapKey{node.name, dir}
		if value, ok := hash[key]; ok {
			return Cycle{ends, value, steps}
		}
		hash[key] = steps
		steps++
	}
	return Cycle{ends, -1, -1}
}

func WalkSecond(dirs []bool, nodes map[string]Node) int {
	cycles := make([]Cycle, 0)
	for k := range nodes {
		if k[len(k)-1:] == "A" {
			fmt.Println("Key", k)
			cycle := FindCycle(k, dirs, nodes)
			fmt.Println("Cycle:", cycle.cycle_start, cycle.cycle_end)
			fmt.Println("Ends:", cycle.ends)
			cycles = append(cycles, cycle)
		}
	}
	fmt.Println()
	// there is only one end node in each cycle in input data
	// and distance from ??A to ??Z is equal to distance from ??Z to end of the cycle
	// so the solution is equivalent to finding LCM
	// nonetheless, it would be nice to generalize this to other scenarios
	ends := make([]int, 0)
	for _, cycle := range cycles {
		ends = append(ends, cycle.ends[0])
	}
	
	return LCM(ends)
}

func GCD(a int, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

func LCM(nums []int) int {
      result := nums[0] * nums[1] / GCD(nums[0], nums[1])
      for i := 2; i < len(nums); i++ {
              result = LCM([]int{result, nums[i]})
      }
      return result
}
