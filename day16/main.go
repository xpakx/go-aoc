package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Advent of Code, day 16")
	fmt.Println("=====================")
	input := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(SolveFirst(input))
}

func SolveFirst(input [][]rune) int {
	beams := make([]Beam, 0)
	beams = append(beams, Beam{Pos{-1,0}, Right})
	nextBeams := make([]Beam, 0)
	hashMap := make(map[Beam]struct{})
	for len(beams) > 0 {
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
		nextBeams = nil
	}

	positions := make(map[Pos]struct{})
	for key := range hashMap {
		positions[key.pos] = struct{}{}
	}
	fmt.Println()
	for i := range input {
		for j := range input[i] {
			if _, ok := positions[Pos{j,i}]; ok {
				fmt.Print("\033[30m", string(input[i][j]), "\033[0m")
			} else {
				fmt.Print(string(input[i][j]))
			}
		}
		fmt.Println()
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
