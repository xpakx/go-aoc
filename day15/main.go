package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of Code, day 15")
	fmt.Println("=====================")
	input := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(SolveFirst(input))
}

func SolveFirst(input [][]rune) int {

	result := 0
	for _, code := range input {
		c := 0
		for i := range code {
			c = (c+int(code[i]))*17%256
		}
		result += c
	}
	return result
}

func Parse(filename string) [][]rune {
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}
	contentString := string(content)
	contentString  = strings.TrimSuffix(contentString, "\n")
	codes := strings.Split(contentString, ",")
	input := make([][]rune, 0)
	for i := range codes {
		input = append(input, []rune(codes[i]))
	}
	
	return input
}
