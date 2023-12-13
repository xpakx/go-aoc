package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 13")
	fmt.Println("=====================")
	input := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(Solve(input))
}

func Solve(input [][][]rune) int {
	result := 0
	for _, a := range input {
		result += SolveSingle(a)
	}
	return result
}

func SolveSingle(input [][]rune) int {
	result := FindSymmetry(input)
	if result == -1 {
		res := FindSymmetry(Transpose(input))
		fmt.Println("transposed")
		fmt.Println(res)
	if res == -1 {
		Print(input)
	}
		return 100*res
	}
		fmt.Println("not transposed")
	fmt.Println(result)
	return result
}

func Parse(filename string) [][][]rune {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make([][][]rune, 0)
	curr := make([][]rune, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			result = append(result, curr)
			curr = nil
			curr = make([][]rune, 0)
		} else {
			curr = append(curr, []rune(line))
		}
	}
	result = append(result, curr)
	readFile.Close()
	return result
}

func FindSymmetry(input [][]rune) int {
	i := 0
	j := 1
	for i < len(input[0]) {
		if j == len(input[0])-2 {
			j++
			i++
		} else if j == len(input[0])-1 {
			i += 2
			if i >= len(input[0]) {
				break
			}
		} else {
			j += 2
		}

		foundDiff := false
		start := i
		end := j
		for start < end && !foundDiff {
			for n := range input {
				if input[n][i] != input[n][j] {
					foundDiff = true
					break
				}
			}
			start++
			end--
		}
		if !foundDiff {
			return i + (j-i+1)/2
		}
	}
	return -1
}

func Transpose(input [][]rune) [][]rune {
	result := make([][]rune, len(input[0]))
	for i := range result {
		result[i] = make([]rune, len(input))
	}

	for i := range input[0] {
		for j := range input {
			result[i][j] = input[j][i]
		}
		
	}
	return result 
}

func Print(input [][]rune) {
	for i := range input {
		for j := range input[i] {
			fmt.Print(string(input[i][j]))
		}
		fmt.Println()
	}
}
