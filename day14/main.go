package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 14")
	fmt.Println("=====================")
	input := Parse("input.txt")
	Print(input)
	fmt.Print("*  ")
	fmt.Println(SolveFirst(input))
	fmt.Print("*  ")
	fmt.Println(SolveSecond(input))
}

func SolveFirst(input [][]rune) int {
	input = TiltNorth(input)
	return Calculate(input)
}

func SolveSecond(input [][]rune) int {
	input = Cycle(input)
	fmt.Println()
	Print(input)
	return Calculate(input)
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

func TiltNorth(input [][]rune) [][]rune {
	for i:=0; i<len(input[0]); i++ {
		last_free := 0
		for j:=0; j<len(input); j++ {
			if input[j][i] == '#' {
				last_free = j+1
			} else if input[j][i] == 'O' {
				input[j][i] = '.'
				input[last_free][i] = 'O'
				last_free++
			}
		}
	}
	return input
}

func TiltSouth(input [][]rune) [][]rune {
	for i:=0; i<len(input[0]); i++ {
		last_free := len(input)-1
		for j:=len(input)-1; j>=0; j-- {
			if input[j][i] == '#' {
				last_free = j-1
			} else if input[j][i] == 'O' {
				input[j][i] = '.'
				input[last_free][i] = 'O'
				last_free--
			}
		}
	}
	return input
}

func TiltWest(input [][]rune) [][]rune {
	for j:=0; j<len(input); j++ {
		last_free := 0
		for i:=0; i<len(input[j]); i++ {
			if input[j][i] == '#' {
				last_free = i+1
			} else if input[j][i] == 'O' {
				input[j][i] = '.'
				input[j][last_free] = 'O'
				last_free++
			}
		}
	}
	return input
}

func TiltEast(input [][]rune) [][]rune {
	for j:=0; j<len(input); j++ {
		last_free := len(input[j])-1
		for i:=len(input[j])-1; i>=0; i-- {
			if input[j][i] == '#' {
				last_free = i-1
			} else if input[j][i] == 'O' {
				input[j][i] = '.'
				input[j][last_free] = 'O'
				last_free--
			}
		}
	}
	return input
}

func Cycle(input [][]rune) [][]rune {
	return TiltEast(TiltSouth(TiltWest(TiltNorth(input))))
}

func Print(input [][]rune) {
	for i:= range input {
		for j:=range input[i] {
			fmt.Print(string(input[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func Calculate(input [][]rune) int {
	result := 0
	for i:=0; i<len(input[0]); i++ {
		for j:=0; j<len(input); j++ {
			if input[j][i] == 'O' {
				result += len(input)-j
			}
		}
	}
	return result
}
