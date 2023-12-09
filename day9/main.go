package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code, day 9")
	fmt.Println("=====================")
	first := solve("input.txt")
	fmt.Print("*  ")
	fmt.Println(first)
}

func solve(filename string) int {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums := ParseList(line)
		fmt.Println(nums)

	}
	readFile.Close()

	return 0
}

func ParseList(list string) []int {
	splitted := strings.Split(strings.Join(strings.Fields(list), " "), " ")
	listFin := []int{}
	for _, x := range splitted {
		n, _ := strconv.Atoi(x)
		listFin = append(listFin, n)
	}
	return listFin
}
