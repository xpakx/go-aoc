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
	fmt.Println(input)
	fmt.Print("*  ")
	fmt.Println(Solve(input))
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
	Left = 1
	Up = 2
	Down = 3
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

func Solve(input []Instruction) int {
	length := 0
	points := make([]Point, 0)
	points = append(points, Point{0, len(input), ""})
	for _, instr := range input {
		last := points[len(points)-1]
		dx := 0
		dy :=0
		if instr.dir == Right {
			dx = instr.length
		} else if instr.dir == Left {
			dx = -instr.length
		} else if instr.dir == Up {
			dy = instr.length
		} else if instr.dir == Down {
			dy = -instr.length
		}
		points = append(points, Point{last.x+dx, last.y+dy, instr.color})
		length += instr.length
		
	}

	area := 0
	for i:=0; i<len(points)-1; i++ {
		area += (points[i].y+points[i+1].y)*(points[i].x-points[i+1].x)
	}
	area = Abs(area)/2

	interior := area + 1 - length/2

	fmt.Println(area)
	fmt.Println(length)
	fmt.Println(interior)

	return length + interior
}

type Point struct {
	x int
	y int
	color string
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
