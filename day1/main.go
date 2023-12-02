package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// first_star()
	second_star()

}

func first_star() {
	fmt.Println("Advent of Code, day 1")
	fmt.Println("----------")

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var result int32 = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		chars := []int32(line)
		i := 0
		for chars[i] > 57 || chars[i] < 48 {
			i++
		}
		j := len(chars) - 1
		for chars[j] > 57 || chars[j] < 48 {
			j--
		}
		first := chars[i] - 48
		second := chars[j] - 48
		number := 10*first + second
		result += number
	}
	fmt.Println(result)

	readFile.Close()
}

func second_star() {
	fmt.Println("Advent of Code, day 1")
	fmt.Println("----------")

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var result int32 = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		chars := []int32(line)
		i := 0
		var first int32 = 0
		for first == 0 {
			if i + 2 < len(chars) {
				if chars[i] == int32('o') && chars[i+1] == int32('n') && chars[i+2] == int32('e') {
					first = 1
				} else if chars[i] == int32('t') && chars[i+1] == int32('w') && chars[i+2] == int32('o') {
					first = 2
				} else if chars[i] == int32('s') && chars[i+1] == int32('i') && chars[i+2] == int32('x') {
					first = 6
				}
			} 
			if i + 3 < len(chars) {
				if chars[i] == int32('f') && chars[i+1] == int32('o') && chars[i+2] == int32('u') && chars[i+3] == int32('r') {
					first = 4
				} else if chars[i] == int32('f') && chars[i+1] == int32('i') && chars[i+2] == int32('v') && chars[i+3] == int32('e') {
					first = 5
				} else if chars[i] == int32('n') && chars[i+1] == int32('i') && chars[i+2] == int32('n') && chars[i+3] == int32('e') {
					first = 9
				}
			} 
			if i + 4 < len(chars) {
				if chars[i] == int32('t') && chars[i+1] == int32('h') && chars[i+2] == int32('r') && chars[i+3] == int32('e') && chars[i+4] == int32('e') {
					first = 3
				} else if chars[i] == int32('s') && chars[i+1] == int32('e') && chars[i+2] == int32('v') && chars[i+3] == int32('e') && chars[i+4] == int32('n') {
					first = 7
				} else if chars[i] == int32('e') && chars[i+1] == int32('i') && chars[i+2] == int32('g') && chars[i+3] == int32('h') && chars[i+4] == int32('t') {
					first = 8
				}
			} 
			if chars[i] <= 57 && chars[i] > 48 {
				first = chars[i] - 48
			}
			i++
		}
		var second int32 = 0
		j := len(chars) - 1
		for second == 0 {
			if j - 2 >= 0 {
				if chars[j-2] == int32('o') && chars[j-1] == int32('n') && chars[j] == int32('e') {
					second = 1
				} else if chars[j-2] == int32('t') && chars[j-1] == int32('w') && chars[j] == int32('o') {
					second = 2
				} else if chars[j-2] == int32('s') && chars[j-1] == int32('i') && chars[j] == int32('x') {
					second = 6
				}
			}
			if j - 3 >= 0 {
				if chars[j-3] == int32('f') && chars[j-2] == int32('o') && chars[j-1] == int32('u') && chars[j] == int32('r') {
					second = 4
				} else if chars[j-3] == int32('f') && chars[j-2] == int32('i') && chars[j-1] == int32('v') && chars[j] == int32('e') {
					second = 5
				} else if chars[j-3] == int32('n') && chars[j-2] == int32('i') && chars[j-1] == int32('n') && chars[j] == int32('e') {
					second = 9
				}
			} 
			if j - 4 >= 0 {
				if chars[j-4] == int32('t') && chars[j-3] == int32('h') && chars[j-2] == int32('r') && chars[j-1] == int32('e') && chars[j] == int32('e') {
					second = 3
				} else if chars[j-4] == int32('s') && chars[j-3] == int32('e') && chars[j-2] == int32('v') && chars[j-1] == int32('e') && chars[j] == int32('n') {
					second = 7
				} else if chars[j-4] == int32('e') && chars[j-3] == int32('i') && chars[j-2] == int32('g') && chars[j-1] == int32('h') && chars[j] == int32('t') {
					second = 8
				}
			} 
			if chars[j] <= 57 && chars[j] > 48 {
				second = chars[j] - 48
			}
			j--
		}
		number := 10*first + second
		result += number
		fmt.Println(number)
	}
	fmt.Println(result)

	readFile.Close()
}
