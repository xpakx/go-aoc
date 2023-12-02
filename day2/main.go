package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code, day 2*")
	fmt.Println("==========")
	fmt.Println()
	fmt.Println("-= * =-")
	result := first_star()
	fmt.Println(result)
	fmt.Println("-= ** =-")
	secondResult := second_star()
	fmt.Println(secondResult)

}

func first_star() int {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	fileScanner.Split(bufio.ScanLines)
	var result int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		id, red, green, blue := getMax(line)
		if red <= maxRed && green <= maxGreen && blue <= maxBlue {
			result += id
		}
	}
	readFile.Close()
	return result
}

func second_star() int {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var result int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		_, red, green, blue := getMax(line)
		power := red*green*blue
		result += power
	}
	readFile.Close()
	return result 
}

func getMax(line string) (int, int, int, int) {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	splitLine := strings.Split(line, ":")
	id := strings.Split(splitLine[0], " ")[1]
	idNum, _ := strconv.Atoi(id)
	sets := strings.Split(splitLine[1], ";")
	for _, set := range sets {
		stones := strings.Split(set, ",")
		for _, stone := range stones {
			stoneTrim := strings.TrimSpace(stone)
			pair := strings.Split(stoneTrim, " ")
			color := pair[1]
			value, _ := strconv.Atoi(pair[0])
			if color == "red" && value > maxRed {
				maxRed = value
			}
			if color == "green" && value > maxGreen {
				maxGreen = value
			}
			if color == "blue" && value > maxBlue {
				maxBlue = value
			}
		}
	}
	return idNum, maxRed, maxGreen, maxBlue
}
