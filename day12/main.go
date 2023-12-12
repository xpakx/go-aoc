package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code, day 12")
	fmt.Println("=====================")
	input := Parse("input.txt")
	fmt.Print("*  ")
	fmt.Println(Solve(input, false))
	fmt.Print("** ")
	fmt.Println(Solve(input, true))
}

func Parse(filename string) []Row {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := make([]Row, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result = append(result, ParseLine(line))
	}
	readFile.Close()
	return result
}

type Row struct {
	row string
	values []int
}

func Solve(input []Row, second bool) int {
	result := 0
	for i := range input {
		posMap := make(map[Pos]int, 0)
		if second {
			row := input[i].row
			row = row + "?" + row + "?" + row + "?" + row + "?" + row
			old := input[i].values
			vals := make([]int, 5*len(old))
			for j := range old {
				vals[j] = old[j]
				vals[j+len(old)] = old[j]
				vals[j+2*len(old)] = old[j]
				vals[j+3*len(old)] = old[j]
				vals[j+4*len(old)] = old[j]
			}
			result += TryToPlaceRec(posMap, row, vals, 0, 0, 0)
		} else {
			result += TryToPlaceRec(posMap, input[i].row, input[i].values, 0, 0, 0)
		}
	}
	return result
}

func ParseLine(line string) Row {
	values := strings.Split(line, " ")
	rows := values[0]
	vals := ParseList(values[1])
	return Row{rows, vals}

}

func ParseList(list string) []int {
	splitted := strings.Split(list, ",")
	listFin := []int{}
	for _, x := range splitted {
		n, _ := strconv.Atoi(x)
		listFin = append(listFin, n)
	}
	return listFin
}

type Pos struct {
	pos int
	value int
	found int
}

func TryToPlaceRec(hash map[Pos]int, row string, values []int, pos int, value int, found int) int {
	key := Pos{pos, value, found}
	if value, ok := hash[key]; ok {
		return value
	}
	if pos == len(row) {
		if value == len(values) && found == 0 {
			return 1
		} else if value == len(values)-1 && values[value]== found {
			return 1
		} else {
			return 0
		}
	}

	result := 0
	if row[pos] == '.' || row[pos] == '?' {
		if found == 0 {
			result += TryToPlaceRec(hash, row, values, pos+1, value, 0)
		} else if found > 0 && value<len(values) && values[value]==found {
			result += TryToPlaceRec(hash, row, values, pos+1, value+1, 0)
		}
	}
	if row[pos] == '#' || row[pos] == '?' {
		result += TryToPlaceRec(hash, row, values, pos+1, value, found+1)
	}
	hash[key] = result
	return result
}
