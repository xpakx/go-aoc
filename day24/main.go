package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 24")
	fmt.Println("=====================")
	input := Parse("input.txt")

	first := Solve(input)
	fmt.Print("*  ")
	fmt.Println(first)
}

func Parse(filename string) []Line {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	board := make([]Line, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		points := strings.Split(line, " @ ")
		p := ParseList(points[0])
		v := ParseList(points[1])
		board = append(board, Line{Pos{p[0], p[1], p[2]}, Pos{p[0]+v[0], p[1]+v[1], p[2]+v[2]}})
	}
	readFile.Close()
	return board
}

func ParseList(list string) []int {
	splitted := strings.Split(list, ", ")
	listFin := []int{}
	for _, x := range splitted {
		n, _ := strconv.Atoi(strings.TrimSpace(x))
		listFin = append(listFin, n)
	}
	return listFin
}


const (
	MAX = 400_000_000_000_000
	MIN = 200_000_000_000_000
)

func (line Line) UpwardX() bool {
	return line.coordA.x < line.coordB.x
}

func (line Line) UpwardY() bool {
	return line.coordA.y < line.coordB.y
}

func (line Line) InThePast(intersection Intersection) bool {
	xInThePast := false
	if line.UpwardX() {
		xInThePast = intersection.x < float64(line.coordA.x)
	} else {
		xInThePast = intersection.x > float64(line.coordA.x)
	}
	yInThePast := false
	if line.UpwardY() {
		yInThePast = intersection.y < float64(line.coordA.y)
	} else {
		yInThePast = intersection.y > float64(line.coordA.y)
	}
	return xInThePast || yInThePast
}

func Solve(lines []Line) int {
	result := 0
	for i:=0; i<len(lines)-1; i++{
		for j:=i+1; j<len(lines); j++ {
			intersection := Intersect(lines[i].coordA, lines[i].coordB, lines[j].coordA, lines[j].coordB)
			if intersection.x == -1 {
				continue
			}
			if intersection.x < MIN || intersection.y < MIN || intersection.x > MAX || intersection.y > MAX  {
				continue
			}
			if lines[i].InThePast(intersection) {
				continue
			}
			if lines[j].InThePast(intersection) {
				continue
			}
			result++
		}
	}
	return result
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
	x float64 
	y float64
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

	return Intersection{x0, y0}
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
