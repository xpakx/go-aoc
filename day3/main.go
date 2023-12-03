package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	// "strconv"
)

func main() {
	fmt.Println("Advent of Code, day 3")
	fmt.Println("=====================")
	result := solve()
	fmt.Print("*  ")
	fmt.Println(result)
}

func solve() int {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		chars := []int32(line)
		current := 0
		start := 0
		for i, char := range chars {
			if char <= 57 && char >= 48 {
				if current == 0 {
					start = i
				}
				digit := int(char-48)
				current *= 10
				current += digit
			} else if current > 0 {
				num := Number{current, start, i-1}
				current = 0
				fmt.Print(num.value)
				fmt.Print("(", num.start, "-", num.end, ")")
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
	readFile.Close()
	return 0
}

type Number struct {
	value int
	start int
	end int
}
