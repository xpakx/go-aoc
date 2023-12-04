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

	for fileScanner.Scan() {
		line := fileScanner.Text()
		game := ParseLine(line)
		fmt.Printf("%#v\n", game)
		common := Intersection(game.numbers, game.winning)
		fmt.Println(common)
		result += CalculatePoints(len(common))
	}
	readFile.Close()
	return result
}

type Game struct {
	id int
	numbers []int
	winning []int
}

func ParseLine(line string) Game {
	reg := regexp.MustCompile("Card\\s+(?P<Id>\\d+): (?P<Winning>.*) \\| (?P<Numbers>.*)")
	match := reg.FindStringSubmatch(line)
	id, _ := strconv.Atoi(match[1])
	winning := ParseList(match[2])
	numbers := ParseList(match[3])
	return Game{id, numbers, winning}
}

func ParseList(list string) []int {
	listWithoutDuplicatedSpaces := strings.Join(strings.Fields(list), " ")
	listOfStrings := strings.Split(listWithoutDuplicatedSpaces, " ")
	listFin := []int{}
	for _, x := range listOfStrings {
		n, _ := strconv.Atoi(x)
		listFin = append(listFin, n)
	}
	return listFin
}

func Intersection(list []int, list2 []int) []int {
	set := make([]int, 0)
	hash := make(map[int]struct{})

	for _, num := range list {
		hash[num] = struct{}{}
	}

	for _, num := range list2 {
		if _, ok := hash[num]; ok {
			set = append(set, num)
		}
	}

	return set
}

func CalculatePoints(numbers int) int {
	if numbers <= 1 {
		return numbers
	} 
	value := 1
	for i:=1; i<numbers; i++ {
		value *= 2
	}
	return value
}
