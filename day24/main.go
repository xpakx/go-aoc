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
	second := SolveSecond(input)
	fmt.Print("** ")
	fmt.Println(second)
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

func SolveSecond(lines []Line) int {
	velocityCandidateX := make([]int, 0)
	velocityCandidateY := make([]int, 0)
	velocityCandidateZ := make([]int, 0)
	for i:=0; i<len(lines)-1; i++{
		for j:=i+1; j<len(lines); j++ {
			if lines[i].vector.x == lines[j].vector.x {
				// they're always equidistant in coordinate x 
				dist := Abs(lines[i].start.x-lines[j].start.x)
				newVelos := make([]int, 0)
				for a:=-300; a<=300; a++ {
					if a != lines[i].vector.x && dist%(a-lines[i].vector.x)==0 {
						newVelos = append(newVelos, a)
					}
				}
				if len(velocityCandidateX) > 0 {
					velocityCandidateX = SetIntersection(velocityCandidateX, newVelos)
				} else {
					velocityCandidateX = newVelos
				}
			}
			if lines[i].vector.y == lines[j].vector.y {
				dist := Abs(lines[i].start.y-lines[j].start.y)
				newVelos := make([]int, 0)
				for a:=-300; a<=300; a++ {
					if a != lines[i].vector.y && dist%(a-lines[i].vector.y)==0 {
						newVelos = append(newVelos, a)
					}
				}
				if len(velocityCandidateY) > 0 {
					velocityCandidateY = SetIntersection(velocityCandidateY, newVelos)
				} else {
					velocityCandidateY = newVelos
				}
			}
			if lines[i].vector.z == lines[j].vector.z {
				dist := Abs(lines[i].start.z-lines[j].start.z)
				newVelos := make([]int, 0)
				for a:=-300; a<=300; a++ {
					if a != lines[i].vector.z && dist%(a-lines[i].vector.z)==0 {
						newVelos = append(newVelos, a)
					}
				}
				if len(velocityCandidateZ) > 0 {
					velocityCandidateZ = SetIntersection(velocityCandidateZ, newVelos)
				} else {
					velocityCandidateZ = newVelos
				}
			}
		}
	}

	// TODO: multiple candidates (and zero candidates?)
	rockVelo := Pos{velocityCandidateX[0], velocityCandidateY[0], velocityCandidateZ[0]}

	// now let's suppose that rock stands in place (its velocity vector would be 0); 
	// if we adjust lines by its velocity, they should all interesect in that point
	line1 := lines[0]
	line1.vector = Pos{line1.vector.x-rockVelo.x, line1.vector.y-rockVelo.y, line1.vector.z-rockVelo.z }
	line2 := lines[1]
	line2.vector = Pos{line2.vector.x-rockVelo.x, line2.vector.y-rockVelo.y, line2.vector.z-rockVelo.z }

	intersection := Intersect(line1, line2)
	line1.start.y = line1.start.z
	line1.vector.y = line1.vector.z
	line2.start.y = line2.start.z
	line2.vector.y = line2.vector.z
	intersectionZ := Intersect(line1, line2)
	// TOD: fix. it calculates result correctly for this input, but intersectSlope is a little too imprecise, 
	// and Intersect have a strange bug sometimes

	return int(intersectionZ.x) + int(intersection.y) + int(intersectionZ.y)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

	x0 := (float64(l.start.y-m.start.y) - float64(l.start.x)*slopeL + float64(m.start.x)*slopeM)/(slopeM-slopeL)
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

func SetIntersection(list []int, list2 []int) []int {
	set := make([]int, 0)
	hash := make(map[int]struct{})
	for _, num := range list {
		hash[num] = struct{}{}
	}
	for _, num := range list2 {
		if _, ok := hash[num]; ok {
			set = append(set, num)
		}
	}
	return set
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
