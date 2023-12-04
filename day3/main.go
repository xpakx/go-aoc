package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 3")
	fmt.Println("=====================")
	result := solve()
	fmt.Print("*  ")
	fmt.Println(result)
	result = solveSecond()
	fmt.Print("** ")
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
			if i + 1 == len(chars) && current > 0 {
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
			}
		}
		numbersLast = nil

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

type Gear struct {
	position int
	adjacent int
	ratio int
}

func solveSecond() int {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	last := make([]int, 140)
	currLine := make([]int, 140)
	for i := range last {
		last[i] = 0
	}
	gearsLast := make([]Gear, 0)
	gears := make([]Gear, 0)
	result := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		chars := []int32(line)
		current := 0
		start := 0
		for i, char := range chars {
			currLine[i] = 0
			if char <= 57 && char >= 48 {
				if current == 0 {
					start = i
				}
				digit := int(char-48)
				current *= 10
				current += digit
			} else if current > 0 {
				for j:=start; j<i; j++ {
					currLine[j] = current
				}
				current = 0
			}
			if i + 1 == len(chars) && current > 0 {
				for j:=start; j<i; j++ {
					currLine[j] = current
				}
				current = 0
			}
			if (char > 57 || char < 48) && char != 46 {
				gear := Gear{i, 0, 1}
				gears = append(gears, gear)
			}
		}

		for _, gear := range gearsLast {
			if currLine[gear.position] != 0 {
				gear.adjacent += 1
				gear.ratio *= currLine[gear.position]
			} else {
				if gear.position > 0 && currLine[gear.position-1] != 0 {
					gear.adjacent += 1
					gear.ratio *= currLine[gear.position-1]
				}
				if gear.position + 1 < len(chars) && currLine[gear.position+1] != 0 {
					gear.adjacent += 1
					gear.ratio *= currLine[gear.position+1]
				}
			}
			if gear.adjacent == 2 {
				result += gear.ratio
			}
		}
		gearsLast = nil

		for _, gear := range gears {
			if gear.position > 0 && currLine[gear.position-1] != 0 {
				gear.adjacent += 1
				gear.ratio *= currLine[gear.position-1]
			}
			if gear.position + 1 < len(chars) && currLine[gear.position+1] != 0 {
				gear.adjacent += 1
				gear.ratio *= currLine[gear.position+1]
			}

			if last[gear.position] != 0 {
				gear.adjacent += 1
				gear.ratio *= last[gear.position]
			} else {
				if gear.position > 0 && last[gear.position-1] != 0 {
					gear.adjacent += 1
					gear.ratio *= last[gear.position-1]
				}
				if gear.position + 1 < len(chars) && last[gear.position+1] != 0 {
					gear.adjacent += 1
					gear.ratio *= last[gear.position+1]
				}
			}
			gearsLast = append(gearsLast, gear)
		}
		gears = nil
		
		for i := range last {
			last[i] = currLine[i]
		}
	}
	readFile.Close()
	return result
}
