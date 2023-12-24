package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 24")
	fmt.Println("=====================")
	input := Parse("input.txt")

	first := Solve(input)
	fmt.Print("*  ")
	fmt.Println(first)
}

func Parse(filename string) [][]rune {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	board := make([][]rune, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		board = append(board, []rune(line))
	}
	readFile.Close()
	return board
}

func Solve(input [][]rune) int {
	return 0
}

type Pos struct {
	x int 
	y int 
	z int
}

type Line struct {
	coordA Pos
	coordB Pos
}

type Intersection struct {
	x int 
	y int
}

func Intersect(a1, a2, b1, b2 Pos) Intersection {
	x1 := a1.x
	y1 := a1.y
	x2 := a2.x
	y2 := a2.y
	x3 := b1.x
	y3 := b1.y
	x4 := b2.x
	y4 := b2.y

	denominator := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denominator == 0 {
		return Intersection{-1, -1}
	}

	a := (x1*y2 - y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)
	b := (x1*y2 - y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)
	
	x0 := float64(a)/float64(denominator)
	y0 := float64(b)/float64(denominator)

	if Between(x0, x1, x2) && Between(y0, y1, y2) && Between(x0, x3, x4) && Between(y0, y3, y4) {
		return Intersection{int(x0), int(y0)}
	}
	
	return Intersection{-1, -1}
}

func Between(x float64, x1, x2 int) bool {
	epsilon := 1E-03
	lower := Min(x1, x2)
	higher := Max(x1, x2)
	return x <= float64(higher) + epsilon && x >= float64(lower) - epsilon
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
