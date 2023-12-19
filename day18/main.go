package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code, day 18")
	fmt.Println("=====================")
	input := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(Solve(input))
	fmt.Print("** ")
	fmt.Println(SolveSecond(input))
}

func Parse(filename string) []Instruction {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make([]Instruction, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result = append(result, ParseLine(line))
	}
	readFile.Close()
	return result
}

type Instruction struct {
	dir  int
	length int
	color string
}

const (
	Right = 0
	Left = 2
	Up = 3
	Down = 1
)

func ParseLine(line string) Instruction {
	reg := regexp.MustCompile("^([LRUD]) (\\d+) \\((.+)\\)$")
	match := reg.FindStringSubmatch(line)
	dir := ToDir(match[1])
	length, _ := strconv.Atoi(match[2])
	color := match[3]
	return Instruction{dir, length, color}
}

func ToDir(letter string) int {
	if letter == "R" {
		return Right
	}
	if letter == "L" {
		return Left
	}
	if letter == "U" {
		return Up
	}
	if letter == "D" {
		return Down
	}
	return -1
}

func (instr Instruction) DirToDelta() (int, int) {
	if instr.dir == Right {
		return instr.length, 0
	} else if instr.dir == Left {
		return -instr.length, 0
	} else if instr.dir == Up {
		return 0, instr.length
	} else if instr.dir == Down {
		return 0, -instr.length
	}
	return 0, 0 
}

func Solve(input []Instruction) int {
	length := 0
	points := make([]Point, 0)
	points = append(points, Point{0, len(input)})
	for _, instr := range input {
		last := points[len(points)-1]
		dx, dy := instr.DirToDelta()
		points = append(points, Point{last.x+dx, last.y+dy})
		length += instr.length
	}

	area := 0
	for i:=0; i<len(points)-1; i++ {
		area += (points[i].y+points[i+1].y)*(points[i].x-points[i+1].x)
	}
	area = Abs(area)/2
	interior := area + 1 - length/2
	return length + interior
}

func SolveSecond(input []Instruction) int {
	newInstr := make([]Instruction, 0)
	for _, inst := range input {
		newInstr = append(newInstr, inst.Transform())
	}
	return Solve(newInstr)
}

type Point struct {
	x int
	y int
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (instr Instruction) Transform() Instruction {
	distStr := instr.color[1:6]
	dir := int(instr.color[6] - '0')
	dist, _ := strconv.ParseUint(distStr, 16, 64)
	return Instruction{dir, int(dist), ""}	
}
