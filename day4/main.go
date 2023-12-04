package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code, day 4")
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
	result := 0
	reg := regexp.MustCompile("Card (?P<Id>\\d+): (?P<Winning>.*) \\| (?P<Numbers>.*)")

	for fileScanner.Scan() {
		line := fileScanner.Text()
		match := reg.FindStringSubmatch(line)
		id := match[1]
		idValue, _ := strconv.Atoi(id)
		winning := match[2]
		numbers := match[3]
		winning = strings.Join(strings.Fields(winning), " ")
		winningList := strings.Split(winning, " ")
		winningFin := []int{}
		for _, x := range winningList {
			n, _ := strconv.Atoi(x)
			winningFin = append(winningFin, n)
		}
		numbers = strings.Join(strings.Fields(numbers), " ")
		numbersList := strings.Split(numbers, " ")
		numbersFin := []int{}
		for _, x := range numbersList {
			n, _ := strconv.Atoi(x)
			numbersFin = append(numbersFin, n)
		}
		numbers = strings.Join(strings.Fields(numbers), " ")
		game := Game{idValue, numbersFin, winningFin}
		fmt.Printf("%#v\n", game)

	}
	readFile.Close()
	return result
}

type Game struct {
	id int
	numbers []int
	winning []int
}
