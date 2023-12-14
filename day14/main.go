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
	fmt.Print("*  ")
	fmt.Println(SolveFirst(input))
	fmt.Print("*  ")
	fmt.Println(SolveSecond(input))
}

func SolveFirst(input [][]rune) int {
	TiltNorth(input)
	return Calculate(input)
}

func SolveSecond(input [][]rune) int {
	cycleStart, cycleEnd := FindCycle(input)
	cycleLen := cycleEnd - cycleStart + 1
        n := 1_000_000_000 - (cycleStart - 1)
	n = n % cycleLen
	for i:= 0; i<n-1; i++ {
		Cycle(input)
	}
	return Calculate(input)
}

func FindCycle(input [][]rune) (int, int) {
	last := make(map[string]int, 0)
	last[ToKey(input)] = 0
	for i:= 1; i<1000; i++ {
		Cycle(input)
		key := ToKey(input)
		if value, ok := last[key]; ok {
			return value, i-1
		} else {
			last[key] = i
		}
	}
	return 0, 0
}

func ToKey(input [][]rune) string {
	key := ""
	for i := range input {
		row := 0
		for j := range input[i] {
			row = row << 1
			if input[i][j] != '.' {
				row += 1
			}
		}
		key += fmt.Sprint(row)
	}
	return key
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

func TiltNorth(input [][]rune) {
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
}

func TiltSouth(input [][]rune) {
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
}

func TiltWest(input [][]rune) {
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
}

func TiltEast(input [][]rune) {
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
}

func Cycle(input [][]rune) {
	TiltNorth(input)
	TiltWest(input)
	TiltSouth(input)
	TiltEast(input)
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
