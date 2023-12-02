package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	first_star()
	second_star()

}

func first_star() {
	fmt.Println("Advent of Code, day 2*")
	fmt.Println("----------")

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var result int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		valid, id := testLine(line)
		if valid {
			result += id
		}
	}
	fmt.Println(result)

	readFile.Close()
}

func testLine(line string) (bool, int) {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	splitLine := strings.Split(line, ":")
	id := strings.Split(splitLine[0], " ")[1]
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return false, 0
	}
	sets := strings.Split(splitLine[1], ";")
	for _, set := range sets {
		stones := strings.Split(set, ",")
		for _, stone := range stones {
			stoneTrim := strings.TrimSpace(stone)
			pair := strings.Split(stoneTrim, " ")
			color := pair[1]
			value, err := strconv.Atoi(pair[0])
			if err != nil {
				return false, idNum
			}
			if color == "blue" && value > maxBlue {
				return false, idNum
			}
			if color == "green" && value > maxGreen {
				return false, idNum 
			}
			if color == "red" && value > maxRed {
				return false, idNum
			}
		}
	}
	return true, idNum
}

func second_star() {
	fmt.Println("Advent of Code, day 2**")
	fmt.Println("----------")

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var result int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		red, green, blue := getMax(line)
		power := red*green*blue
		result += power
	}
	fmt.Println(result)

	readFile.Close()
}

func getMax(line string) (int, int, int) {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	splitLine := strings.Split(line, ":")
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
	return maxRed, maxGreen, maxBlue
}
