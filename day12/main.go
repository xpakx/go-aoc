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
	first := SolveFirst("input.txt")
	fmt.Print("*  ")
	fmt.Println(first)
}

func SolveFirst(filename string) int {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		row :=  ParseLine(line);
		fmt.Println(row)
		result += row
	}
	readFile.Close()
	return result
}

func ParseLine(line string) int {
	fmt.Println("New Line")
	fmt.Println(line)
	values := strings.Split(line, " ")
	rows := values[0]
	vals := ParseList(values[1])

	return TryToPlace(rows, vals)
}

func TryToPlace(rows string, vals []int) int {
	possibilities := make([]int, 0)
	possibilities = append(possibilities, 0)

	for _, v := range vals {
		newPossibilities := make([]int, 0)
		for _, start := range possibilities {
			i := start
			for i<len(rows) && rows[i] != '#' {
				if TestPlacement(rows, i, v) {
					newPossibilities = append(newPossibilities, i+v-1+2)
				}
				i++
			}
			if TestPlacement(rows, i, v) {
				newPossibilities = append(newPossibilities, i+v-1+2)
			}
		}
		possibilities = nil
		possibilities = append(possibilities, newPossibilities...)
	}
	result := 0
	for _, p := range possibilities {
		if TestOmittedPlaces(rows, p) {
			result++
		}
	}

	return result
}

func TestPlacement(row string, start int, length int) bool {
	if start + length - 1 >= len(row) {
		return false
	}
	for i:=start; i<start+length; i++ {
		if row[i] == '.' {
			return false
		}
	}
	if start+length<len(row) && row[start+length] == '#' {
		return false
	}
	if start>0 && row[start-1] == '#' {
		return false
	}
	return true
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

func TestOmittedPlaces(row string, end int) bool {
	for i:=end; i<len(row); i++ {
		if row[i] == '#' {
			return false
		}
	}
	return true
}
