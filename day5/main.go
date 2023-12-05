package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Advent of Code, day 5")
	fmt.Println("=====================")
	first:= solve()
	fmt.Print("*  ")
	fmt.Println(first)
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
	}
	readFile.Close()

	return 0
}
