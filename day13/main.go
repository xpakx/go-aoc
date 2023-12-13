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
	fmt.Println(FindSymmetry(input[0]))
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
	return 0
}
