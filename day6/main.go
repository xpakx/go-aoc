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
	fmt.Println("Advent of Code, day 6")
	fmt.Println("=====================")
	first := solve()
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
	var lines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	readFile.Close()

	races := ParseRaces(lines[0], lines[1])
	for _, race := range races {
		fmt.Println("Race with time", race.time, ", currect record:", race.distance)
	}

	return 0
}

func ParseRaces(times string, distances string) []Race {
	timesList := ParseLine(times)
	distancesList := ParseLine(distances)
	
	var races []Race
	for i := range timesList {
		race := Race{timesList[i], distancesList[i]}
		races = append(races, race)
	}
	return races
}

func ParseLine(line string) []int {
	reg := regexp.MustCompile("(Time|Distance):\\s+(.*)")
	match := reg.FindStringSubmatch(line)
	list := match[2]
	return ParseList(list)
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

type Race struct {
	time int
	distance int
}
