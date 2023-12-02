package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
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
