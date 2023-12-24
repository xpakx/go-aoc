package main

import (
	"bufio"
	"fmt"
	"math"
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
		board = append(board, Line{Pos{p[0], p[1], p[2]}, Pos{v[0], v[1], v[2]}})
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
	return !math.Signbit(float64(line.vector.x))
}

func (line Line) UpwardY() bool {
	return !math.Signbit(float64(line.vector.y))
}

func (line Line) InThePast(intersection Intersection) bool {
	xInThePast := false
	epsilon := 1E-03
	if line.UpwardX() {
		xInThePast = intersection.x < float64(line.start.x)+epsilon
	} else {
		xInThePast = intersection.x > float64(line.start.x)-epsilon
	}
	yInThePast := false
	if line.UpwardY() {
		yInThePast = intersection.y < float64(line.start.y)+epsilon
	} else {
		yInThePast = intersection.y > float64(line.start.y)-epsilon
	}
	return xInThePast || yInThePast
}

func Solve(lines []Line) int {
	result := 0
	epsilon := 1E-03
	for i:=0; i<len(lines)-1; i++{
		for j:=i+1; j<len(lines); j++ {
			intersection := IntersectSlope(lines[i], lines[j])
			if intersection.x == -1 {
				continue
			}
			if intersection.x < MIN+epsilon || intersection.y < MIN+epsilon || intersection.x > MAX-epsilon || intersection.y > MAX-epsilon  {
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
	start Pos
	vector Pos
}

type Intersection struct {
	x float64 
	y float64
}

func IntersectSlope(l, m Line) Intersection {
	epsilon := 1E-03
	slopeL := float64(l.vector.y)/float64(l.vector.x)
	slopeM := float64(m.vector.y)/float64(m.vector.x)
	if slopeL >= slopeM - epsilon && slopeL <= slopeM + epsilon {
		return Intersection{-1, -1}
	}

	x0 := (float64(l.start.y)-float64(m.start.y)  - float64(l.start.x)*slopeL + float64(m.start.x)*slopeM)/(slopeM-slopeL)
	y0 := float64(l.start.y) + slopeL*(x0-float64(l.start.x))
	return Intersection{x0, y0}
}

func Intersect(l, m Line) Intersection {
	x1 := l.start.x
	y1 := l.start.y
	x2 := l.start.x+l.vector.x
	y2 := l.start.y+l.vector.y
	x3 := m.start.x
	y3 := m.start.y
	x4 := m.start.x+m.vector.x
	y4 := m.start.y+m.vector.y

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
