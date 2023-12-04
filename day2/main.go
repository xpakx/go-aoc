package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code, day 2")
	fmt.Println("=====================")
	result := solve()
	fmt.Print("*  ")
	fmt.Println(result.firstStar)
	fmt.Print("** ")
	fmt.Println(result.secondStar)
}

type Results struct {
	firstStar int
	secondStar int
}

func solve() Results {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	fileScanner.Split(bufio.ScanLines)
	var first int = 0
	var second int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		game := getMax(line)

		if game.red <= maxRed && game.green <= maxGreen && game.blue <= maxBlue {
			first += game.id
		}

		power := game.red * game.green * game.blue
		second += power
	}
	readFile.Close()
	return Results{first, second}
}

type GameData struct {
	id int
	red int
	green int
	blue int
}

func getMax(line string) GameData {
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
	return GameData{id: idNum, red: maxRed, green: maxGreen, blue: maxBlue}
}
