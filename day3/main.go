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
	last := make([]bool, 140)
	currLine := make([]bool, 140)
	for i := range last {
		last[i] = false
	}
	numbersLast := make([]Number, 0)
	numbers := make([]Number, 0)
	result := 0

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
				numbers = append(numbers, num)
			}
			if (char > 57 || char < 48) && char != 46 {
				currLine[i] = true
			} else {
				currLine[i] = false
			}
		}
		fmt.Println("last")
		for _, num := range numbersLast {
			start := num.start
			if num.start > 0 {
				start -= 1
			}
			end := num.end
			if num.end+1 < len(chars) {
				end += 1
			}
			test := false
			for i:=start; i<=end; i++ {
				if currLine[i] {
					test = true
				}
			}
			if test {
				result += num.value
				fmt.Println("added", num.value)
			}
		}
		numbersLast = nil
		fmt.Println("current")

		for _, num := range numbers {
			start := num.start
			if num.start > 0 {
				start -= 1
			}
			end := num.end
			if num.end+1 < len(chars) {
				end += 1
			}
			test := false
			for i:=start; i<=end; i++ {
				if last[i] {
					test = true
				}
			}

			if test || (num.start > 0 && currLine[num.start-1]) || (num.end+1 < len(chars) && currLine[num.end+1]) {
				result += num.value
				fmt.Println("added", num.value)

			} else {
				numbersLast = append(numbersLast, num)
			}

		}
		numbers = nil
		
		for i := range last {
			last[i] = currLine[i]
		}
	}
	readFile.Close()
	return result
}

type Number struct {
	value int
	start int
	end int
}
