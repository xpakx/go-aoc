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
	first := solve()
	fmt.Print("*  ")
	fmt.Println(first)
	second := solveSecond()
	fmt.Print("** ")
	fmt.Println(second)
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
	newSeeds := make([]bool, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !seedsLoaded {
			seeds = ParseSeeds(line)
			newSeeds = make([]bool, len(seeds))
			for i := range seeds {
				newSeeds[i] = false
			}
			seedsLoaded = true
		} else if IsHeader(line) {
			for i := range seeds {
				newSeeds[i] = false
			}
		} else if line != "" {
			rng := ParseRange(line)
			for i := range seeds {
				if !newSeeds[i] && rng.InRange(seeds[i]) {
					seeds[i] = rng.Transform(seeds[i])
					newSeeds[i] = true
				}
			}
		}
	}
	readFile.Close()

	return getMin(seeds)
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

func getMinRange(list []SeedRange) int {
	min := list[0].start
	for i:=1; i<len(list); i++ {
		if list[i].start < min {
			min = list[i].start
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

func ParseSeedRanges(line string) []SeedRange {
	numbers := ParseSeeds(line)
	seedRanges := make([]SeedRange, 0)
	for i:=0; i<len(numbers); i+=2 {
		seedRanges = append(seedRanges, SeedRange{numbers[i], numbers[i+1]})
	}
	return seedRanges
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
	return number - rng.fromStart + rng.toStart
}

func (rng Range) TransformRange(seeds SeedRange) SeedRange {
	return SeedRange{seeds.start - rng.fromStart + rng.toStart, seeds.length}
}

func (rng Range) InRange(number int) bool {
	return number >= rng.fromStart && number <= rng.FromEnd() 
}

type SeedRange struct {
	start int
	length int
}

func (rng SeedRange) End() int {
	return rng.start + rng.length - 1
}

func solveSecond() int {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	seedsLoaded := false
	seedRanges := make([]SeedRange, 0)
	newRanges := make([]SeedRange, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !seedsLoaded {
			seedRanges = ParseSeedRanges(line)
			seedsLoaded = true
		} else if IsHeader(line) {
			for _, r := range newRanges {
				seedRanges = append(seedRanges, r)
			}
			newRanges = nil
		} else if line != "" {
			rng := ParseRange(line)
			transformed, untransformed := SplitRanges(rng, seedRanges)
			for _, r := range transformed {
				newRanges = append(newRanges, r)
			}
			seedRanges = nil
			for _, r := range untransformed {
				seedRanges = append(seedRanges, r)
			}
		}
	}
	readFile.Close()
	for _, r := range newRanges {
		seedRanges = append(seedRanges, r)
	}

	return getMinRange(seedRanges)
}

func SplitRanges(rng Range, seedRanges []SeedRange) ([]SeedRange, []SeedRange) {
	transformed := make([]SeedRange, 0)
	untransformed := make([]SeedRange, 0)
	for _, r  := range seedRanges {
		intersection := GetIntersection(rng, r)
		if intersection.length != 0 {
			newRange := rng.TransformRange(intersection) 
			transformed = append(transformed, newRange)
			left := GetLeft(rng, r)
			right := GetLeft(rng, r)
			if left.length != 0 {
				untransformed = append(untransformed, left)
			}
			if right.length != 0 {
				untransformed = append(untransformed, right)
			}
		} else {
			untransformed = append(untransformed, r)
		}
	}
	return transformed, untransformed
}

func GetIntersection(rng Range, seedRange SeedRange) SeedRange {
	if rng.fromStart > seedRange.End() || rng.FromEnd() < seedRange.start {
		return SeedRange{0, 0}
	}
	if rng.fromStart <= seedRange.start && rng.FromEnd() >= seedRange.End() {
		return seedRange
	}
	if rng.fromStart > seedRange.start && rng.FromEnd() < seedRange.End() {
		return SeedRange{rng.fromStart, rng.length}
	}
	if rng.fromStart >= seedRange.start {
		return SeedRange{rng.fromStart, seedRange.End() - rng.fromStart + 1}
	}
	return SeedRange{seedRange.start, rng.FromEnd() - seedRange.start + 1}
}

func GetLeft(rng Range, seedRange SeedRange) SeedRange {
	if rng.fromStart <= seedRange.start {
		return SeedRange{0, 0}
	}
	return SeedRange{seedRange.start, rng.fromStart - seedRange.start}
}

func GetRight(rng Range, seedRange SeedRange) SeedRange {
	if rng.FromEnd() >= seedRange.End() {
		return SeedRange{0, 0}
	}
	return SeedRange{rng.FromEnd()+1, seedRange.End() - rng.FromEnd()}
}
