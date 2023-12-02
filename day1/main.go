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

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
