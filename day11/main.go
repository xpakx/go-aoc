package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 11")
	fmt.Println("=====================")
	galaxyMap := LoadMap("input.txt")
	fmt.Println(galaxyMap)
	first := CalcDistances(galaxyMap)
	second := solveSecond(galaxyMap)
	fmt.Print("*  ")
	fmt.Println(first)
	fmt.Print("** ")
	fmt.Println(second)
}

func solveSecond(galaxyMap []Galaxy) int {
	return 0
}


func LoadMap(filename string) []Galaxy {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make([]Galaxy, 0)
	emptyColumns := make([]bool, 0)
	y := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, "")
		empty := true
		for x := range values {
			if y == 0 {
				emptyColumns = append(emptyColumns, true)
			}
			if values[x] == "#" {
				galaxy := Galaxy{x, y}
				result = append(result, galaxy)
				empty = false
				emptyColumns[x] = false
			}
		}
		if empty {
			y++
		}
		y++
	}
	readFile.Close()

	for i := range result {
		n := result[i].x
		for j:=0; j<result[i].x; j++ {
			if emptyColumns[j] {
				n += 1
			}
		}
		result[i].x = n
	}

	return result
}


type Galaxy struct {
	x int
	y int
}

func CalcDistances(nodeMap []Galaxy) int {
	dist := 0
	for i := range nodeMap {
		for j:=i+1; j<len(nodeMap); j++{
			dist += Abs(nodeMap[i].x - nodeMap[j].x)+Abs(nodeMap[i].y - nodeMap[j].y)
		}
	}

	return dist
}

func Abs(x int) int {
   if x < 0 {
      return -x
   }
   return x
}
