package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 16")
	fmt.Println("=====================")
	input := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(SolveFirst(input))
	fmt.Print("*  ")
	fmt.Println(SolveSecond(input))
}

func SolveFirst(input[][] rune) int {
	return Solve(input, Beam{Pos{-1, 0}, Right})
}

func SolveSecond(input [][]rune) int {
	results := make([]int, 0)
	for i:=0; i<len(input); i++ {
		results = append(results, Solve(input, Beam{Pos{-1, i}, Right}))
		results = append(results, Solve(input, Beam{Pos{len(input[i]), i}, Left}))
	}
	for i:=0; i<len(input[0]); i++ {
		results = append(results, Solve(input, Beam{Pos{i, -1}, Down}))
		results = append(results, Solve(input, Beam{Pos{i, len(input)}, Up}))
	}
	result := 0 
	for _, num := range results {
		if num > result {
			result = num
		}
	}
	return result
}

func Solve(input [][]rune, startBeam Beam) int {
	beams := make([]Beam, 0)
	beams = append(beams, startBeam)
	hashMap := make(map[Beam]struct{})
	for len(beams) > 0 {
		nextBeams := make([]Beam, 0)
		for _, beam := range beams {
			if _, ok := hashMap[beam]; !ok {
				hashMap[beam] = struct{}{}
				newPos := beam.step()
				if newPos.x >= 0 && newPos.y >= 0 && newPos.y < len(input) && newPos.x < len(input[newPos.y]) {
					symbol := input[newPos.y][newPos.x]
					beam.pos = newPos
					nextBeams = append(nextBeams, GetNewBeams(beam, symbol)...)
				}
			}
		}
		beams = nil
		beams = append(beams, nextBeams...)
	}

	positions := make(map[Pos]struct{})
	for key := range hashMap {
		positions[key.pos] = struct{}{}
	}

	return len(positions)-1
}


func GetNewBeams(beam Beam, symbol rune) []Beam {
	nextBeams := make([]Beam, 0)
	if symbol == '.' {
		nextBeams = append(nextBeams, beam)
	} else if symbol == '/' {
		if beam.dir == Right {
			beam.dir = Up
		} else if beam.dir == Down {
			beam.dir = Left
		} else if beam.dir == Left {
			beam.dir = Down
		} else if beam.dir == Up {
			beam.dir = Right
		}
		nextBeams = append(nextBeams, beam)
	} else if symbol == '\\' {
		if beam.dir == Right {
			beam.dir = Down
		} else if beam.dir == Down {
			beam.dir = Right
		} else if beam.dir == Left {
			beam.dir = Up
		} else if beam.dir == Up {
			beam.dir = Left
		}
		nextBeams = append(nextBeams, beam)
	} else if symbol == '|' {
		if beam.dir == Up || beam.dir == Down {
			nextBeams = append(nextBeams, beam)
		} else {
			beam.dir = Up
			secondBeam := Beam{beam.pos, Down}
			nextBeams = append(nextBeams, beam)
			nextBeams = append(nextBeams, secondBeam)
		}
	} else if symbol == '-' {
		if beam.dir == Left || beam.dir == Right {
			nextBeams = append(nextBeams, beam)
		} else {
			beam.dir = Left
			secondBeam := Beam{beam.pos, Right}
			nextBeams = append(nextBeams, beam)
			nextBeams = append(nextBeams, secondBeam)
		}
	}
	return nextBeams
}

func Parse(filename string) [][]rune {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make([][]rune, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result = append(result, []rune(line))
	}
	readFile.Close()
	return result
}

type Beam struct {
	pos  Pos
	dir int
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

func (beam Beam) step() Pos {
	if beam.dir == Right {
		return Pos{beam.pos.x + 1, beam.pos.y}
	}
	if beam.dir == Left {
		return Pos{beam.pos.x - 1, beam.pos.y}
	}
	if beam.dir == Up {
		return Pos{beam.pos.x, beam.pos.y - 1}
	}
	if beam.dir == Down {
		return Pos{beam.pos.x, beam.pos.y + 1}
	}
	return beam.pos
}
