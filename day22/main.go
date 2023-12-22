package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 22")
	fmt.Println("=====================")
	blocks := Parse("input.txt")
	fmt.Println(blocks)


	fmt.Print("*  ")
	fmt.Println(SolveFirst(blocks))
}

func Parse(filename string) map[int][]Block {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make(map[int][]Block, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		block := GetBlock(line)
		lowerZ := Min(block.coordA.z, block.coordB.z)
		if _, ok := result[lowerZ]; ok {
			result[lowerZ] = append(result[lowerZ], block)
		} else {
			result[lowerZ] = make([]Block, 0)
			result[lowerZ] = append(result[lowerZ], block)
		}
	}
	readFile.Close()

	return result
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetBlock(line string) Block {
	coords := strings.Split(line, "~")
	coordA := strings.Split(coords[0], ",")
	coordB := strings.Split(coords[1], ",")
	return Block{
		ListToPos(coordA),
		ListToPos(coordB),
		make([]Block, 0),
		make([]Block, 0),
	}
}

func ListToPos(pos []string) Pos {
	x, _ := strconv.Atoi(pos[0])
	y, _ := strconv.Atoi(pos[1])
	z, _ := strconv.Atoi(pos[2])
	return Pos{x, y, z}
}

type Pos struct {
	x int
	y int
	z int
}

type Block struct {
	coordA Pos
	coordB Pos
	supports []Block
	supportedBy []Block
}

func SolveFirst(blocks map[int][]Block) int {
	return 0
}
