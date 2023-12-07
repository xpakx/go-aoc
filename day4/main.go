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
	first, second := solve()
	fmt.Print("*  ")
	fmt.Println(first)
	fmt.Print("** ")
	fmt.Println(second)
}

func solve() (int, int) {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	first := 0
	second := 0
	numbers := make([]int, 213)
	for i:=0; i<len(numbers); i++ {
		numbers[i] = 1
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		game := ParseLine(line)
		common := Intersection(game.numbers, game.winning)
		first += CalculatePoints(len(common))
		for i:=1; i<=len(common); i++ {
			id := game.id - 1 + i
			if id < len(numbers) {
				numbers[id] += numbers[game.id-1]
			}
		}
	}
	readFile.Close()

	for _, num := range numbers {
		second += num
	}
	return first, second
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
	splitted := strings.Split(strings.Join(strings.Fields(list), " "), " ")
	listFin := []int{}
	for _, x := range splitted {
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
