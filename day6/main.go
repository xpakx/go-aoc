package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
	"strconv"
	"math"
)

func main() {
	fmt.Println("Advent of Code, day 6")
	fmt.Println("=====================")
	input := GetInput("input.txt")
	first := solveFirst(input)
	fmt.Print("*  ")
	fmt.Println(first)
}

func solveFirst(lines []string) int {
	races := ParseRaces(lines[0], lines[1])
	result := 1
	for _, race := range races {
		fmt.Println("Race with time", race.time, ", currect record:", race.distance)
		solution := SolveRace(race)
		fmt.Println(solution, "ways to win")
		result *= solution
	}
	return result
}

func GetInput(filename string) []string {
	readFile, err := os.Open(filename)

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
	return lines
}


func SolveRace(race Race) int {
	x1, x2 := SolveQuadraticFunction(-1, race.time, -race.distance)
	return x1 - x2 - 1
}

// T - time for race, S - current record
// S < V*(T-t), where t is time for holding and V is equal to t
// S < t*(T-t)
// 0 < -t^2 + Tt - S
// a < 0
func SolveQuadraticFunction(a int, b int, c int) (int, int) {
	aa := float64(a)
	bb := float64(b)
	cc := float64(c)
	delta := bb*bb - 4*aa*cc

	if delta == 0 {
		x := -b/(2*a)
		return int(x), int(x)
	}
	x1 := (-bb - math.Sqrt(delta))/(2*aa)
	x1 = math.Ceil(x1)
	x2 := (-bb + math.Sqrt(delta))/(2*aa)
	x2 = math.Floor(x2)
	return int(x1), int(x2)
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
