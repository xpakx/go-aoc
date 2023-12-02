package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	first_star()
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
			if test_string(chars, "one", i) {
				first = 1
			}
			if test_string(chars, "two", i) {
				first = 2
			}
			if test_string(chars, "three", i) {
				first = 3
			}
			if test_string(chars, "four", i) {
				first = 4
			}
			if test_string(chars, "five", i) {
				first = 5
			}
			if test_string(chars, "six", i) {
				first = 6
			}
			if test_string(chars, "seven", i) {
				first = 7
			}
			if test_string(chars, "eight", i) {
				first = 8
			}
			if test_string(chars, "nine", i) {
				first = 9
			}
			if chars[i] <= 57 && chars[i] > 48 {
				first = chars[i] - 48
			}
			i++
		}
		var second int32 = 0
		j := len(chars) - 1
		for second == 0 {
			if test_string_reverse(chars, "one", j) {
				second = 1
			}
			if test_string_reverse(chars, "two", j) {
				second = 2
			}
			if test_string_reverse(chars, "three", j) {
				second = 3
			}
			if test_string_reverse(chars, "four", j) {
				second = 4
			}
			if test_string_reverse(chars, "five", j) {
				second = 5
			}
			if test_string_reverse(chars, "six", j) {
				second = 6
			}
			if test_string_reverse(chars, "seven", j) {
				second = 7
			}
			if test_string_reverse(chars, "eight", j) {
				second = 8
			}
			if test_string_reverse(chars, "nine", j) {
				second = 9
			}
			if chars[j] <= 57 && chars[j] > 48 {
				second = chars[j] - 48
			}
			j--
		}
		number := 10*first + second
		result += number
	}
	fmt.Println(result)

	readFile.Close()
}

func test_string(str []int32, substr string, start int) bool {
	if start < 0 {
		return false
	}
	if start + len(substr) > len(str) { 
		return false
	}
	chars := []int32(substr)
	for i:=0; i < len(chars); i++ {
		if chars[i] != str[start+i] {
			return false
		}
	}
	return true
}

func test_string_reverse(str []int32, substr string, start int) bool {
	newStart := start - len(substr) + 1
	return test_string(str, substr, newStart)
}
