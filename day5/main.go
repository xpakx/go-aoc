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

	seedsLoaded := false
	seeds := make([]int, 0)
	newSeeds := make([]int, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !seedsLoaded {
			seeds = ParseSeeds(line)
			newSeeds = make([]int, len(seeds))
			for i := range seeds {
				newSeeds[i] = seeds[i]
			}
			seedsLoaded = true
		} else if IsHeader(line) {
			for i := range seeds {
				seeds[i] = newSeeds[i]
			}
		} else if line != "" {
			rng := ParseRange(line)
			for i := range seeds {
				if newSeeds[i] == seeds[i] {
					newSeeds[i] = rng.Transform(seeds[i])
				}
			}
		}
	}
	readFile.Close()

	return getMin(newSeeds)
}

func getMin(list []int) int {
	min := list[0]
	for i:=1; i<len(list); i++ {
		if list[i] < min {
			min = list[i]
		}
	}
	return min
}

func ParseSeeds(line string) []int {
	reg := regexp.MustCompile("seeds: (.*)$")
	match := reg.FindStringSubmatch(line)
	numbers := ParseList(match[1])
	return numbers
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

func ParseListHeader(line string) Header {
	reg := regexp.MustCompile("([a-z]+)-to-([a-z]+) map:$")
	match := reg.FindStringSubmatch(line)
	return Header{match[1], match[2]}
}

func IsHeader(line string) bool {
	reg := regexp.MustCompile("([a-z]+)-to-([a-z]+) map:$")
	return reg.MatchString(line)
}

func ParseRange(line string) Range {
	list := ParseList(line)
	return Range{list[1], list[0], list[2]}
}

type Header struct {
	from string
	to string
}

type Range struct {
	fromStart int
	toStart int
	length int
}

func (rng Range) FromEnd() int {
	return rng.fromStart + rng.length - 1
}

func (rng Range) ToEnd() int {
	return rng.toStart + rng.length - 1
}

func (rng Range) Transform(number int) int {
	if number < rng.fromStart || number > rng.FromEnd() {
		return number
	}
	return number - rng.fromStart + rng.toStart
}
